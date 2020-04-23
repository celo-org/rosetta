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

type processor struct {
	ctx                 context.Context
	headers             <-chan *types.Header
	changes             chan<- *db.BlockChangeSet
	cc                  *client.CeloClient
	registry            *wrapper.RegistryWrapper
	epochRewardsAddress common.Address
	gpmAddress          common.Address
	gpm                 *big.Int
	logger              log.Logger
}

var ErrMultipleGasPriceMinimumUpdates = errors.New("Error multiple GasPriceMinimumUpdated events emmitted in same block")

func BlockProcessor(ctx context.Context, headers <-chan *types.Header, changes chan<- *db.BlockChangeSet, cc *client.CeloClient, db_ db.RosettaDBReader, logger log.Logger) error {
	bp, err := newProcessor(ctx, headers, changes, cc, db_, logger)
	if err != nil {
		return err
	}

	for {
		h, err := bp.nextHeader()
		if err != nil {
			return err
		}

		bcs := &db.BlockChangeSet{
			BlockNumber: h.Number,
		}

		if err := bp.registryChanges(bcs); err != nil {
			return err
		}

		if err := bp.gasPriceMinimum(bcs); err != nil {
			return err
		}

		if err := bp.carbonOffsetPartner(bcs); err != nil {
			return err
		}

		if err := bp.writeChanges(bcs); err != nil {
			return err
		}
	}
}

func newProcessor(ctx context.Context, headers <-chan *types.Header, changes chan<- *db.BlockChangeSet, cc *client.CeloClient, db_ db.RosettaDBReader, logger log.Logger) (*processor, error) {
	registry, err := wrapper.NewRegistry(cc)
	if err != nil {
		return nil, err
	}

	lastProcessedBlock, err := db_.LastPersistedBlock(ctx)
	if err != nil {
		return nil, err
	}

	epochRewardsAddress, err := db_.RegistryAddressStartOf(ctx, lastProcessedBlock, 0, "EpochRewards")
	if err != nil && err != db.ErrContractNotFound {
		return nil, err
	}

	gpmAddress, err := db_.RegistryAddressStartOf(ctx, lastProcessedBlock, 0, "GasPriceMinimum")
	if err != nil && err != db.ErrContractNotFound {
		return nil, err
	}

	gpm, err := db_.GasPriceMinimumFor(ctx, new(big.Int).Add(lastProcessedBlock, big.NewInt(1)))
	if err != nil {
		return nil, err
	}

	return &processor{
		ctx:                 ctx,
		headers:             headers,
		changes:             changes,
		cc:                  cc,
		registry:            registry,
		epochRewardsAddress: epochRewardsAddress,
		gpmAddress:          gpmAddress,
		gpm:                 gpm,
		logger:              logger.New("pipe", "processor"),
	}, nil
}

func (bp *processor) writeChanges(bcs *db.BlockChangeSet) error {
	select {
	case <-bp.ctx.Done():
		return bp.ctx.Err()
	case bp.changes <- bcs:
		return nil
	}
}

func (bp *processor) nextHeader() (*types.Header, error) {
	select {
	case <-bp.ctx.Done():
		return nil, bp.ctx.Err()
	case h := <-bp.headers:
		return h, nil
	}
}

func (bp *processor) registryChanges(bcs *db.BlockChangeSet) error {
	blockNumber := bcs.BlockNumber.Uint64()

	iter, err := bp.registry.Contract().FilterRegistryUpdated(&bind.FilterOpts{
		End:     &blockNumber,
		Start:   blockNumber,
		Context: bp.ctx,
	}, nil, nil)
	if err != nil {
		return err
	}

	registryChanges := make([]db.RegistryChange, 0)

	for iter.Next() {
		if iter.Event.Identifier == "GasPriceMinimum" {
			bp.gpmAddress = iter.Event.Addr
		}
		if iter.Event.Identifier == "EpochRewards" {
			bp.epochRewardsAddress = iter.Event.Addr
		}
		registryChanges = append(registryChanges, db.RegistryChange{
			TxIndex:    iter.Event.Raw.TxIndex,
			Contract:   iter.Event.Identifier,
			NewAddress: iter.Event.Addr,
		})
		bp.logger.Info("Core Contract Address Changed", "name", iter.Event.Identifier, "newAddress", iter.Event.Addr.Hex(), "txIndex", iter.Event.Raw.TxIndex)
	}
	if err := iter.Error(); err != nil {
		return err
	}

	bcs.RegistryChanges = registryChanges

	return nil
}

func (bp *processor) gasPriceMinimum(bcs *db.BlockChangeSet) error {
	if bp.gpmAddress == common.ZeroAddress {
		return nil
	}

	blockNumber := bcs.BlockNumber.Uint64()

	gpmContract, err := contract.NewGasPriceMinimum(bp.gpmAddress, bp.cc.Eth)
	if err != nil {
		return err
	}

	iter, err := gpmContract.FilterGasPriceMinimumUpdated(&bind.FilterOpts{
		End:     &blockNumber,
		Start:   blockNumber,
		Context: bp.ctx,
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
		if bp.gpm.Cmp(gpmNew) != 0 {
			bp.logger.Info("Gas Price Minimum Updated", "from", bp.gpm, "to", gpmNew, "block", blockNumber)
			bcs.GasPriceMinimum = new(big.Int).Set(gpmNew)
			bp.gpm = new(big.Int).Set(gpmNew)
		}
		multipleUpdates = true
	}

	if err := iter.Error(); err != nil {
		return err
	}

	return nil
}

func (bp *processor) carbonOffsetPartner(bcs *db.BlockChangeSet) error {
	if bp.epochRewardsAddress == common.ZeroAddress {
		return nil
	}

	blockNumber := bcs.BlockNumber.Uint64()

	epochRewards, err := contract.NewEpochRewards(bp.epochRewardsAddress, bp.cc.Eth)
	if err != nil {
		return err
	}

	iter, err := epochRewards.FilterCarbonOffsettingFundSet(&bind.FilterOpts{
		End:     &blockNumber,
		Start:   blockNumber,
		Context: bp.ctx,
	}, nil)
	if err != nil {
		return err
	}

	for iter.Next() {
		bcs.CarbonOffsetPartnerChange = db.CarbonOffsetPartnerChange{
			Address: iter.Event.Partner,
			TxIndex: iter.Event.Raw.TxIndex,
		}
	}

	if err := iter.Error(); err != nil {
		return err
	}

	return nil
}
