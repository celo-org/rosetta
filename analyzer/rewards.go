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
	header *types.Header
}

func ComputeEpochRewards(ctx context.Context, cc *client.CeloClient, db db.RosettaDBReader, header *types.Header) (*Operation, error) {
	rewards := make(map[common.Address]*big.Int)

	rctx := &rewardsContext{
		ctx:    ctx,
		cc:     cc,
		header: header,
	}

	// We use start of next block for queries since the won't change during epoch rewards
	nextBlock := new(big.Int).Add(header.Number, big.NewInt(1))

	// Obtain the address of all the contract we need for the computation
	addresses, err := db.RegistryAddressesStartOf(ctx, nextBlock, 0, "EpochRewards", "LockedGold", "Governance", "Reserve", "GoldToken")
	if err != nil {
		return nil, err
	}
	if len(addresses) < 5 {
		// we don't have all addresses =>
		// We assume rewards are active AFTER migration. So, we return an empty map
		return NewEpochRewards(rewards), nil
	}

	goldToken, err := contract.NewGoldToken(addresses["GoldToken"], cc.Eth)
	if err != nil {
		return nil, err
	}
	delete(addresses, "GoldToken")

	addresses["CarbonOffset"], err = rctx.getCarbonOffsettingPartner(addresses["EpochRewards"])
	if err != nil {
		return nil, err
	}
	delete(addresses, "EpochRewards")

	for contract, address := range addresses {
		rewardsForAddr, err := rctx.getRewardsForAddr(goldToken, address)
		if err != nil {
			return nil, err
		}

		if rewardsForAddr.Cmp(big.NewInt(0)) > 0 {
			rewards[addresses[contract]] = rewardsForAddr
		}
	}

	// NOTE: Validator & Validator Groups rewards are in CUSD. We don't cover them for now

	return NewEpochRewards(rewards), nil
}

func (rctx *rewardsContext) blockNumber() *big.Int {
	return rctx.header.Number
}

func (rctx *rewardsContext) prevBlockNumber() *big.Int {
	return new(big.Int).Sub(rctx.header.Number, big.NewInt(1))
}

func (rctx *rewardsContext) getRewardsForAddr(token *contract.GoldToken, addr common.Address) (*big.Int, error) {
	blockNumber := rctx.blockNumber().Uint64()
	iter, err := token.FilterTransfer(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: rctx.ctx,
	}, []common.Address{common.ZeroAddress}, []common.Address{addr})

	if err != nil {
		return nil, err
	}

	rewards := big.NewInt(0)
	for iter.Next() {
		rewards.Add(rewards, iter.Event.Value)
	}
	return rewards, nil
}

func (rctx *rewardsContext) getCarbonOffsettingPartner(epochRewardsAddr common.Address) (common.Address, error) {
	epochRewards, err := contract.NewEpochRewards(epochRewardsAddr, rctx.cc.Eth)
	if err != nil {
		// Error here means a Bug. Failed to parse ABI
		return common.ZeroAddress, err
	}
	return epochRewards.CarbonOffsettingPartner(&bind.CallOpts{
		BlockNumber: rctx.blockNumber(), // by the end of the block
		Context:     rctx.ctx,
	})
}
