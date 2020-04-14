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

var (
	ErrBlockProcessor = errors.New("Error In Block Processor")
)

func BlockProcessor(ctx context.Context, headers <-chan *types.Header, changes chan<- *db.BlockChangeSet, cc *client.CeloClient, db_ db.RosettaDBReader, logger log.Logger) error {
	logger = logger.New("pipe", "processor")

	registry, err := wrapper.NewRegistry(cc)
	if err != nil {
		return err
	}

	cache := make(map[string]*types.Header)

	lastProcessedBlock, err := db_.LastPersistedBlock(ctx)
	if err != nil {
		return err
	}

	var gpmAddress common.Address
	var gpmPrev *big.Int
	var h *types.Header

	for {

		nextExpectedBlock := new(big.Int).Add(lastProcessedBlock, big.NewInt(1))

		if val, ok := cache[nextExpectedBlock.String()]; ok {
			// If the next expected block is cached, assign it to h and process.
			if val.Number.Cmp(nextExpectedBlock) != 0 {
				log.Error("Block Processor Key/Value Mismatch", "key", nextExpectedBlock.Int64(), "value", val.Number.Int64()) // TODO(Alec): Return formatted error here.
				return ErrBlockProcessor
			}
			h = val
			delete(cache, nextExpectedBlock.String())
			log.Info("Cached Block is Processing", "block", h.Number.Int64())
		} else {
			// If the next expected block is not cached, wait for it on headers channel.
		selectLoop:
			for {
				select {
				case <-ctx.Done():
					return ctx.Err()
				case h = <-headers:
					switch h.Number.Cmp(nextExpectedBlock) {
					case 0: // We have the right block, process it.
						break selectLoop
					case -1: // We somehow got an old block, error.
						log.Error("Block(s) Repeated in Processing", "last processed", lastProcessedBlock.Int64(), "now receiving", h.Number.Int64()) // TODO(Alec): Return formatted error here.
						return ErrBlockProcessor
					case 1: // We aren't ready to process this block yet, add it to cache.
						cache[h.Number.String()] = h
					}
				}
			}
		}

		lastProcessedBlock = h.Number

		bcs := db.BlockChangeSet{
			BlockNumber: h.Number,
		}

		blockNumber := h.Number.Uint64()

		iter, err := registry.Contract().FilterRegistryUpdated(&bind.FilterOpts{
			End:     &blockNumber,
			Start:   blockNumber,
			Context: ctx,
		}, nil)
		if err != nil {
			return err
		}

		registryChanges := make([]db.RegistryChange, 0)
		for iter.Next() {
			if iter.Event.Identifier == "GasPriceMinimum" {
				gpmAddress = iter.Event.Addr
				//TODO(Alec): could gpm value change unexpectedly in this case?
			}
			registryChanges = append(registryChanges, db.RegistryChange{
				TxIndex:    iter.Event.Raw.TxIndex,
				Contract:   iter.Event.Identifier,
				NewAddress: iter.Event.Addr,
			})
			logger.Info("Core Contract Address Changed", "name", iter.Event.Identifier, "newAddress", iter.Event.Addr.Hex(), "txIndex", iter.Event.Raw.TxIndex)
		}
		if err != nil {
			return err
		}

		bcs.RegistryChanges = registryChanges

		// If gpmAddress has been defined, ask it for the gpm.
		if gpmAddress != common.ZeroAddress {
			gpm, err := getGasPriceMinimumFromEthCall(ctx, cc, gpmAddress, h.Number)
			if err != nil {
				return err
			}
			// We only add GasPriceMinimum to bcs if there's a change.
			if gpmPrev == nil || gpmPrev.Cmp(gpm) != 0 {
				bcs.GasPriceMinimum = gpm
			}
			gpmPrev = gpm
		}

		// TODO(Alec): add rc1 implementation (leave commented for now)

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
