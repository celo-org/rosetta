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

package analyzer

import (
	"context"
	"math/big"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/rosetta/db"
)

type rewardsContext struct {
	ctx    context.Context
	cc     *client.CeloClient
	db     db.RosettaDBReader
	header *types.Header
}

func ComputeEpochRewards(ctx context.Context, cc *client.CeloClient, db db.RosettaDBReader, header *types.Header) (*Operation, error) {
	rctx := &rewardsContext{
		ctx:    ctx,
		cc:     cc,
		db:     db,
		header: header,
	}

	rewards := make(map[common.Address]*big.Int)
	// Epoch rewards are distributed in the last tx of each block.
	txIndex, err := rctx.cc.Eth.TransactionCount(rctx.ctx, rctx.header.Hash())
	if err != nil {
		return nil, err
	}

	// Ensure that we fetch the goldToken contract event from the current
	goldToken, err := rctx.getGoldToken(txIndex)
	if err != nil {
		return nil, err
	}
	if goldToken == nil {
		// we don't have the GoldToken address =>
		// We assume rewards are active AFTER migration. So, we return an empty map
		return NewEpochRewards(rewards), nil
	}

	if err := rctx.computeRewards(rewards, goldToken, txIndex); err != nil {
		return nil, err
	}

	// NOTE: Validator & Validator Groups rewards are in CUSD. We don't cover them for now

	return NewEpochRewards(rewards), nil
}

func (rctx *rewardsContext) blockNumber() *big.Int {
	return rctx.header.Number
}

func (rctx *rewardsContext) computeRewards(rewardsMap map[common.Address]*big.Int, token *contracts.GoldToken, txIndex uint) error {
	blockNumber := rctx.blockNumber().Uint64()

	iter, err := token.FilterTransfer(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: rctx.ctx,
	}, []common.Address{common.ZeroAddress}, nil)
	if err != nil {
		return err
	}

	for iter.Next() {
		if iter.Event.Raw.TxIndex == txIndex {
			rewardsMap[iter.Event.To] = iter.Event.Value
		}
	}

	return nil
}

func (rctx *rewardsContext) getGoldToken(txIndex uint) (*contracts.GoldToken, error) {
	address, err := rctx.db.RegistryAddressStartOf(rctx.ctx, rctx.blockNumber(), txIndex, "GoldToken")
	if err != nil && err != db.ErrContractNotFound {
		return nil, err
	}
	if address == common.ZeroAddress || err == db.ErrContractNotFound {
		// we don't have the address  =>
		// We assume rewards are active AFTER migration. So, we return nil
		return nil, nil
	}

	goldToken, err := contracts.NewGoldToken(address, rctx.cc.Eth)
	if err != nil {
		return nil, err
	}

	return goldToken, nil
}
