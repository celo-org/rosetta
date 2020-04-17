package monitor

import (
	"context"
	"errors"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/celo-org/rosetta/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var ErrMultipleGasPriceMinimumUpdates = errors.New("Error multiple GasPriceMinimumUpdated events emmitted in same block")

func BlockProcessor(ctx context.Context, headers <-chan *types.Header, changes chan<- *db.BlockChangeSet, cc *client.CeloClient, db_ db.RosettaDBReader, logger log.Logger) error {
	logger = logger.New("pipe", "processor")

	registry, err := wrapper.NewRegistry(cc)
	if err != nil {
		return err
	}
	lastProcessedBlock, err := db_.LastPersistedBlock(ctx)
	if err != nil {
		return err
	}
	gpmAddress, err := db_.RegistryAddressOn(ctx, lastProcessedBlock, 0, "GasPriceMinimum")
	if err != nil && err != db.ErrContractNotFound {
		return err
	}
	gpm, err := db_.GasPriceMinimumOn(ctx, lastProcessedBlock)
	if err != nil {
		return err
	}

	var h *types.Header

	for {

		select {
		case <-ctx.Done():
			return ctx.Err()
		case h = <-headers:
		}

		bcs := db.BlockChangeSet{
			BlockNumber: h.Number,
		}

		blockNumber := h.Number.Uint64()

		iter, err := registry.Contract().FilterRegistryUpdated(&bind.FilterOpts{
			End:     &blockNumber,
			Start:   blockNumber,
			Context: ctx,
		}, nil, nil)
		if err != nil {
			return err
		}

		registryChanges := make([]db.RegistryChange, 0)
		for iter.Next() {
			if iter.Event.Identifier == "GasPriceMinimum" {
				gpmAddress = iter.Event.Addr
			}
			registryChanges = append(registryChanges, db.RegistryChange{
				TxIndex:    iter.Event.Raw.TxIndex,
				Contract:   iter.Event.Identifier,
				NewAddress: iter.Event.Addr,
			})
			logger.Info("Core Contract Address Changed", "name", iter.Event.Identifier, "newAddress", iter.Event.Addr.Hex(), "txIndex", iter.Event.Raw.TxIndex)
		}
		if err := iter.Error(); err != nil {
			return err
		}

		bcs.RegistryChanges = registryChanges

		if gpmAddress != common.ZeroAddress {
			gpmContract, err := contract.NewGasPriceMinimum(gpmAddress, cc.Eth)
			if err != nil {
				return err
			}

			iter, err := gpmContract.FilterGasPriceMinimumUpdated(&bind.FilterOpts{
				End:     &blockNumber,
				Start:   blockNumber,
				Context: ctx,
			})
			if err != nil {
				return err
			}

			// iter should only have 1 event, as gpm can only be updated once per block.
			for multipleUpdates := false; iter.Next(); {
				if multipleUpdates {
					return ErrMultipleGasPriceMinimumUpdates
				}
				gpmNew := iter.Event.GasPriceMinimum

				// We only add GasPriceMinimum to bcs if there's a change.
				if gpm.Cmp(gpmNew) != 0 {
					logger.Info("Gas Price Minimum Updated", "from", gpm, "to", gpmNew, "block", h.Number)
					gpm = gpmNew
					bcs.GasPriceMinimum = gpm
				}
				multipleUpdates = true
			}

			if err := iter.Error(); err != nil {
				return err
			}
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case changes <- &bcs:
		}
	}
}

func getGasPriceMinimumFromEthCall(ctx context.Context, cc *client.CeloClient, gpmAddress common.Address, block *big.Int) (*big.Int, error) {
	gpmContract, err := contract.NewGasPriceMinimum(gpmAddress, cc.Eth)
	if err != nil {
		return nil, err
	}

	callOpts := &bind.CallOpts{
		BlockNumber: block,
		Context:     ctx,
	}

	gpm, err := gpmContract.GetGasPriceMinimum(callOpts, common.ZeroAddress)
	if err != nil {
		return nil, err
	}

	return gpm, nil
}
