package analyzer

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/db"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

	goldToken, err := rctx.getGoldToken()
	if err != nil {
		return nil, err
	}
	if goldToken == nil {
		// we don't have the GoldToken address =>
		// We assume rewards are active AFTER migration. So, we return an empty map
		return NewEpochRewards(rewards), nil
	}

	if err := rctx.computeRewards(rewards, goldToken); err != nil {
		return nil, err
	}

	// NOTE: Validator & Validator Groups rewards are in CUSD. We don't cover them for now

	return NewEpochRewards(rewards), nil
}

func (rctx *rewardsContext) blockNumber() *big.Int {
	return rctx.header.Number
}

func (rctx *rewardsContext) nextBlockNumber() *big.Int {
	return new(big.Int).Add(rctx.header.Number, big.NewInt(1))
}

func (rctx *rewardsContext) prevBlockNumber() *big.Int {
	return new(big.Int).Sub(rctx.header.Number, big.NewInt(1))
}

func (rctx *rewardsContext) computeRewards(rewardsMap map[common.Address]*big.Int, token *contract.GoldToken) error {
	blockNumber := rctx.blockNumber().Uint64()

	iter, err := token.FilterTransfer(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: rctx.ctx,
	}, []common.Address{common.ZeroAddress}, nil)
	if err != nil {
		return err
	}

	// Epoch rewards are distributed in the last tx of each block.

	txCount, err := rctx.cc.Eth.TransactionCount(rctx.ctx, rctx.header.Hash())
	if err != nil {
		return err
	}

	for iter.Next() {
		if iter.Event.Raw.TxIndex == txCount-1 {
			rewardsMap[iter.Event.To] = iter.Event.Value
		}
	}

	return nil
}

func (rctx *rewardsContext) getGoldToken() (*contract.GoldToken, error) {
	address, err := rctx.db.RegistryAddressStartOf(rctx.ctx, rctx.nextBlockNumber(), 0, "GoldToken")
	if err != nil {
		return nil, err
	}
	if address == common.ZeroAddress {
		// we don't have the address  =>
		// We assume rewards are active AFTER migration. So, we return nil
		return nil, nil
	}

	goldToken, err := contract.NewGoldToken(address, rctx.cc.Eth)
	if err != nil {
		return nil, err
	}

	return goldToken, nil
}
