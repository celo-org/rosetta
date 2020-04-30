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
	gpmAddress, err := db_.RegistryAddressStartOf(ctx, lastProcessedBlock, 0, "GasPriceMinimum")
	if err != nil && err != db.ErrContractNotFound {
		return err
	}
	gpm, err := db_.GasPriceMinimumFor(ctx, new(big.Int).Add(lastProcessedBlock, big.NewInt(1)))
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
		}, nil)
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
		if err != nil {
			return err
		}

		bcs.RegistryChanges = registryChanges

		if gpmAddress != common.ZeroAddress {
			gpmNew, err := getGasPriceMinimumFromEthCall(ctx, cc, gpmAddress, h.Number)
			if err != nil {
				return err
			}
			// We only add GasPriceMinimum to bcs if there's a change.
			if gpm.Cmp(gpmNew) != 0 {
				bcs.GasPriceMinimum = gpmNew
			}
			gpm = gpmNew
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
