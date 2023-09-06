// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitor

import (
	"context"
	"errors"
	"math/big"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/log"
	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/registry"
	"github.com/celo-org/rosetta/db"
)

//nolint:unused
type processor struct {
	ctx                 context.Context
	headers             <-chan *types.Header
	changes             chan<- *db.BlockChangeSet
	cc                  *client.CeloClient
	registry            *contracts.Registry
	epochRewardsAddress common.Address
	gpmAddress          common.Address
	reserveAddress      common.Address
	gpm                 *big.Int
	logger              log.Logger
}

var ErrMultipleGasPriceMinimumUpdates = errors.New("Error multiple GasPriceMinimumUpdated events emitted in same block")

func BlockProcessor(ctx context.Context, headers <-chan *types.Header, changes chan<- *db.BlockChangeSet, cc *client.CeloClient, db_ db.RosettaDBReader, isGingerbread func(*big.Int) bool, logger log.Logger) error {
	bp, err := newProcessor(ctx, headers, changes, cc, db_, logger)
	if err != nil {
		return err
	}

	for {
		bcs, err := bp.initNextBlockChangeSet()
		if err != nil {
			return err
		}

		// TODO current impl doesn't catch case where contract changed mid-tx and has events before & after that

		if err := bp.registryChanges(bcs); err != nil {
			return err
		}

		// After Gingerbread HF, GPM must be fetched directly from the block header
		// (changes at the start of the block as opposed to at the end)
		if !isGingerbread(bcs.BlockNumber) {
			if err := bp.gasPriceMinimum(bcs); err != nil {
				return err
			}
		}

		if err := bp.writeChanges(bcs); err != nil {
			return err
		}
	}
}

func newProcessor(ctx context.Context, headers <-chan *types.Header, changes chan<- *db.BlockChangeSet, cc *client.CeloClient, db_ db.RosettaDBReader, logger log.Logger) (*processor, error) {
	registry, err := contracts.NewRegistry(registry.RegistryAddress, cc.Eth)
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

	// GasPriceMinimum is updated at the end of each block and applied to the following block.
	// So, to get the gpm that was SET at the end of the lastProcessedBlock we query the gpm
	// used FOR the next block.
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

func (bp *processor) initNextBlockChangeSet() (*db.BlockChangeSet, error) {
	h, err := bp.nextHeader()
	if err != nil {
		return nil, err
	}

	return &db.BlockChangeSet{
		BlockNumber: h.Number,
	}, nil
}

func (bp *processor) registryChanges(bcs *db.BlockChangeSet) error {
	blockNumber := bcs.BlockNumber.Uint64()

	iter, err := bp.registry.FilterRegistryUpdated(&bind.FilterOpts{
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

	gpmContract, err := contracts.NewGasPriceMinimum(bp.gpmAddress, bp.cc.Eth)
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
