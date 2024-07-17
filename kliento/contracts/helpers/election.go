// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://wwe.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package helpers

import (
	"fmt"
	"math/big"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/rosetta/kliento/contracts"
	"github.com/celo-org/rosetta/kliento/utils/bn"
)

// Election helper methods
type Election struct{ *contracts.Election }

type VotesByGroup map[common.Address]*big.Int
type ElectionVotes struct {
	Active  VotesByGroup
	Pending VotesByGroup
}

func (e *Election) GetAccountElectionVotes(opts *bind.CallOpts, account common.Address) (*ElectionVotes, error) {
	groups, err := e.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var pending VotesByGroup = make(map[common.Address]*big.Int)
	var active VotesByGroup = make(map[common.Address]*big.Int)
	for _, groupAddr := range groups {
		// TODO(yorke): dedup
		pendingAmt, err := e.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if pendingAmt.Cmp(bn.Big0) == 1 {
			pending[groupAddr] = pendingAmt
		}

		activeAmt, err := e.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if activeAmt.Cmp(bn.Big0) == 1 {
			active[groupAddr] = activeAmt
		}
	}

	return &ElectionVotes{
		Active:  active,
		Pending: pending,
	}, nil
}

func (e *Election) GetAccountPendingVotes(opts *bind.CallOpts, account common.Address) (VotesByGroup, error) {
	groups, err := e.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes VotesByGroup = make(map[common.Address]*big.Int)
	for _, groupAddr := range groups {
		pendingAmt, err := e.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if pendingAmt.Cmp(bn.Big0) == 1 {
			votes[groupAddr] = pendingAmt
		}
	}

	return votes, nil
}

func (e *Election) GetAccountActiveVotes(opts *bind.CallOpts, account common.Address) (VotesByGroup, error) {
	groups, err := e.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes VotesByGroup = make(map[common.Address]*big.Int)
	for _, groupAddr := range groups {
		activeAmt, err := e.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if activeAmt.Cmp(bn.Big0) == 1 {
			votes[groupAddr] = activeAmt
		}
	}

	return votes, nil
}

func (e *Election) GetVotesForGroupByAccount(opts *bind.CallOpts, groupAddr, account common.Address) (*big.Int, error) {
	activeAmt, err := e.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
	if err != nil {
		return nil, err
	}
	pendingAmt, err := e.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Add(activeAmt, pendingAmt), nil
}

type AddressLesserGreater struct {
	Lesser  common.Address
	Greater common.Address
}

func (e *Election) VoteMetadata(opts *bind.CallOpts, group common.Address, value *big.Int) (*AddressLesserGreater, error) {
	votes, err := e.GetTotalVotesForEligibleValidatorGroups(opts)
	if err != nil {
		return nil, err
	}

	var voteTotal *big.Int
	for idx, currGroup := range votes.Groups {
		if group == currGroup {
			voteTotal = bn.Add(value, votes.Values[idx])
			break
		}
	}

	var result AddressLesserGreater
	// relies on votes.Groups being sorted
	for idx, currGroup := range votes.Groups {
		if group != currGroup {
			currValue := votes.Values[idx]
			if bn.IsLte(currValue, voteTotal) {
				result.Lesser = currGroup
				break
			}
			result.Greater = currGroup
		}
	}

	return &result, nil
}

func (e *Election) Vote(opts *bind.TransactOpts, group common.Address, value *big.Int) (*types.Transaction, error) {
	meta, err := e.VoteMetadata(callOptsFromTxOpts(opts), group, value)
	if err != nil {
		return nil, err
	}

	return e.Election.Vote(opts, group, value, meta.Lesser, meta.Greater)
}

type RevokeMetadata struct {
	Index *big.Int
	Value *big.Int
	*AddressLesserGreater
}

func addressIndexOf(slice []common.Address, item common.Address) (*big.Int, error) {
	for idx, currItem := range slice {
		if currItem == item {
			return big.NewInt(int64(idx)), nil
		}
	}
	return nil, fmt.Errorf("Item not found in slice")
}

func (e *Election) RevokeMetadata(opts *bind.CallOpts, account common.Address, group common.Address, value *big.Int) (*RevokeMetadata, error) {
	groups, err := e.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	// Only provide lesser & greater indices if the group is eligible,
	// since ineligible validators are excluded in total votes,
	// but we need the correct indices in this list for eligible validators.
	eligibleGroups, err := e.GetEligibleValidatorGroups(opts)
	if err != nil {
		return nil, err
	}
	groupIsEligible := false
	for _, currGroup := range eligibleGroups {
		if currGroup == group {
			groupIsEligible = true
		}
	}
	voteMeta := &AddressLesserGreater{}
	if groupIsEligible {
		voteMeta, err = e.VoteMetadata(opts, group, bn.Neg(value))
		if err != nil {
			return nil, err
		}
	}

	idx, err := addressIndexOf(groups, group)
	if err != nil {
		return nil, err
	}

	result := RevokeMetadata{
		Value:                value,
		Index:                idx,
		AddressLesserGreater: voteMeta,
	}
	return &result, nil
}

type RevokeBothMetadata struct {
	Pending *RevokeMetadata
	Active  *RevokeMetadata
}

func (e *Election) RevokeBothMetadata(opts *bind.CallOpts, account common.Address, group common.Address, value *big.Int) (*RevokeBothMetadata, error) {
	votesPending, err := e.GetPendingVotesForGroupByAccount(opts, group, account)
	if err != nil {
		return nil, err
	}

	var result RevokeBothMetadata
	pendingValue := bn.Min(votesPending, value)
	if bn.IsNonZero(pendingValue) {
		result.Pending, err = e.RevokeMetadata(opts, account, group, pendingValue)
		if err != nil {
			return nil, err
		}
	}

	if bn.IsLt(pendingValue, value) {
		activeValue := bn.Sub(value, pendingValue)
		result.Active, err = e.RevokeMetadata(opts, account, group, activeValue)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (e *Election) RevokeAllMetadata(opts *bind.CallOpts, account common.Address, group common.Address) (*RevokeBothMetadata, error) {
	totalVotes, err := e.GetTotalVotesForGroupByAccount(opts, group, account)
	if err != nil {
		return nil, err
	}

	return e.RevokeBothMetadata(opts, account, group, totalVotes)
}

type RevokeAll struct {
	Pending *types.Transaction
	Active  *types.Transaction
}

func (e *Election) RevokeAll(opts *bind.TransactOpts, account common.Address, group common.Address) (*RevokeAll, error) {
	meta, err := e.RevokeAllMetadata(callOptsFromTxOpts(opts), account, group)
	if err != nil {
		return nil, err
	}

	pendingTx, err := e.RevokePending(opts, group, meta.Pending.Value, meta.Pending.Lesser, meta.Pending.Greater, meta.Pending.Index)
	if err != nil {
		return nil, err
	}

	activeTx, err := e.RevokeActive(opts, group, meta.Active.Value, meta.Active.Lesser, meta.Active.Greater, meta.Active.Index)
	if err != nil {
		return nil, err
	}

	result := RevokeAll{
		Pending: pendingTx,
		Active:  activeTx,
	}
	return &result, nil
}

func (e *Election) ActivateAllMetadata(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	groups, err := e.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var activatableGroups []common.Address
	for _, group := range groups {
		activatable, err := e.HasActivatablePendingVotes(opts, account, group)
		if err != nil {
			return nil, err
		}
		if activatable {
			activatableGroups = append(activatableGroups, group)
		}
	}
	return activatableGroups, nil
}

func (e *Election) ActivateAll(opts *bind.TransactOpts) ([]*types.Transaction, error) {
	groups, err := e.ActivateAllMetadata(callOptsFromTxOpts(opts), opts.From)
	if err != nil {
		return nil, err
	}

	activateTransactions := make([]*types.Transaction, len(groups))
	for idx, group := range groups {
		activateTransactions[idx], err = e.Activate(opts, group)
		if err != nil {
			return nil, err
		}
	}
	return activateTransactions, nil
}
