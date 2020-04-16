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
	addresses, err := db.RegistryAddressesStartOf(ctx, nextBlock, 0, "EpochRewards", "LockedGold", "Governance", "Reserve")
	if err != nil {
		return nil, err
	}
	if len(addresses) < 4 {
		// we don't have all addresses =>
		// We assume rewards are active AFTER migration. So, we return an empty map
		return NewEpochRewards(rewards), nil
	}

	// Voters rewards are deposited to the LockedGold contract
	voterRewards, err := rctx.getTotalVoterRewards(addresses)
	if err != nil {
		return nil, err
	}
	if voterRewards.Cmp(big.NewInt(0)) > 0 {
		rewards[addresses["LockedGold"]] = voterRewards
	}

	// Rewards are deposited in: Reserve, Governance, CarbonOffsetAddress (if any	)
	// We can't compute exactly the rewards.
	// For now the approach is to compute the balance difference with the previous block
	// but this can have errors since it doesn't consider transfer that happened within the block in other txs

	// First obtain carbonOffsettingPartner
	epochRewards, err := contract.NewEpochRewards(addresses["EpochRewards"], rctx.cc.Eth)
	if err != nil {
		// Error here means a Bug. Failed to parse ABI
		return nil, err
	}
	carbonOffsettingPartner, err := epochRewards.CarbonOffsettingPartner(&bind.CallOpts{
		BlockNumber: rctx.blockNumber(), // by the end of the block
		Context:     rctx.ctx,
	})
	if err != nil {
		// Error here means some unexpected error, abort
		return nil, err
	}

	for _, addr := range []common.Address{addresses["Reserve"], addresses["Governance"], carbonOffsettingPartner} {
		diff, err := rctx.balanceDifferenceInBlock(addr)
		if err != nil {
			return nil, err
		}
		// If balance is positive register it
		// If balance were negative, then we are surely doing something wrong. Related to our bad heuristic
		if diff.Cmp(new(big.Int)) > 0 {
			rewards[addr] = diff
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

func (rctx *rewardsContext) balanceDifferenceInBlock(address common.Address) (*big.Int, error) {
	currBalance, err := rctx.cc.Eth.BalanceAt(rctx.ctx, address, rctx.header.Number)
	if err != nil {
		return nil, err
	}
	prevBalance, err := rctx.cc.Eth.BalanceAt(rctx.ctx, address, rctx.prevBlockNumber())
	if err != nil {
		return nil, err
	}

	currBalance.Sub(currBalance, prevBalance)
	return currBalance, nil
}

func (rctx *rewardsContext) getTotalVoterRewards(addresses map[string]common.Address) (*big.Int, error) {
	election, err := contract.NewElection(addresses["Election"], rctx.cc.Eth)
	if err != nil {
		return nil, err
	}

	blockNumber := rctx.blockNumber().Uint64()
	iter, err := election.FilterEpochRewardsDistributedToVoters(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: rctx.ctx,
	}, nil)

	if err != nil {
		return nil, err
	}

	totalRewards := new(big.Int)
	for iter.Next() {
		totalRewards.Add(totalRewards, iter.Event.Value)
	}
	return totalRewards, nil
}
