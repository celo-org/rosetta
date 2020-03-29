package celo

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

type rewardsContext struct {
	ctx    context.Context
	cc     *client.CeloClient
	header *types.Header
}

func ComputeEpochRewards(ctx context.Context, cc *client.CeloClient, header *types.Header) (map[common.Address]*big.Int, error) {
	rewards := make(map[common.Address]*big.Int)

	rctx := &rewardsContext{
		ctx:    ctx,
		cc:     cc,
		header: header,
	}

	// Obtain the address of all the contract we need for the computation
	addresses, err := rctx.getEpochRewardsAddresses()
	if err == wrapper.ErrRegistryNotDeployed || err == client.ErrContractNotDeployed {
		// We assume rewards are active AFTER migration. So, we return an empty map
		return rewards, nil
	} else if err != nil {
		return nil, err
	}

	// Voters rewards are deposited to the LockedGold contract
	rewards[addresses[params.LockedGoldRegistryId]], err = rctx.getTotalVoterRewards(addresses)
	if err != nil {
		return nil, err
	}

	// Rewards are deposited in: Reserve, Governance, CarbonOffsetAddress (if any	)
	// We can't compute exactly the rewards.
	// For now the approach is to compute the balance difference with the previous block
	// but this can have errors since it doesn't consider transfer that happened within the block in other txs

	// First obtain carbonOffsettingPartner
	epochRewards, err := contract.NewEpochRewards(addresses[params.EpochRewardsRegistryId], rctx.cc.Eth)
	if err != nil {
		// Error here means a Bug. Failed to parse ABI
		return nil, err
	}
	carbonOffsettingPartner, err := epochRewards.CarbonOffsettingPartner(&bind.CallOpts{
		BlockNumber: rctx.blockNumber(),
		Context:     rctx.ctx,
	})
	if err != nil {
		// Error here means some unexpected error, abort
		return nil, err
	}

	for _, addr := range []common.Address{addresses[params.ReserveRegistryId], addresses[params.GovernanceRegistryId], carbonOffsettingPartner} {
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

	return rewards, nil
}

func (rctx *rewardsContext) blockNumber() *big.Int {
	return rctx.header.Number
}

func (rctx *rewardsContext) prevBlockNumber() *big.Int {
	return new(big.Int).Sub(rctx.header.Number, big.NewInt(1))
}

func (rctx *rewardsContext) contractsIdentifiers() [][32]byte {
	return [][32]byte{
		params.EpochRewardsRegistryId,
		params.LockedGoldRegistryId,
		params.GovernanceRegistryId,
		params.ReserveRegistryId,
	}
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

func (rctx *rewardsContext) getTotalVoterRewards(addresses map[[32]byte]common.Address) (*big.Int, error) {
	election, err := contract.NewElection(addresses[params.ElectionRegistryId], rctx.cc.Eth)
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

func (rctx *rewardsContext) getEpochRewardsAddresses() (map[[32]byte]common.Address, error) {
	addresses := make(map[[32]byte]common.Address)

	registry, err := wrapper.NewRegistry(rctx.cc)
	if err != nil {
		return nil, err
	}

	// Since there are not registry changes in the Finalize tx
	// we can query the state just after that (block = header.Number)
	callOpts := &bind.CallOpts{
		BlockNumber: rctx.blockNumber(),
		Context:     rctx.ctx,
	}

	for _, id := range rctx.contractsIdentifiers() {
		addr, err := registry.GetAddressFor(callOpts, id)
		if err != nil {
			return nil, err
		}
		addresses[id] = addr
	}

	return addresses, nil
}
