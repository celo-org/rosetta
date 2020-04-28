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

package wrapper

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type ElectionWrapper struct {
	*contract.Election
}

func NewElection(celoClient *client.CeloClient, registryWrapper *RegistryWrapper) (*ElectionWrapper, error) {
	election, err := registryWrapper.GetElection(nil, celoClient.Eth)
	if err != nil {
		return nil, err
	}

	return &ElectionWrapper{
		election,
	}, nil
}

type VotesByGroup map[common.Address]*big.Int
type ElectionVotes struct {
	Active  VotesByGroup
	Pending VotesByGroup
}

func (w *ElectionWrapper) GetAccountElectionVotes(opts *bind.CallOpts, account common.Address) (*ElectionVotes, error) {
	groups, err := w.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes *ElectionVotes
	for _, groupAddr := range groups {
		// TODO(yorke): dedup
		pendingAmt, err := w.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if utils.IsNonZero(pendingAmt) {
			votes.Pending[groupAddr] = pendingAmt
		}

		activeAmt, err := w.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if utils.IsNonZero(activeAmt) {
			votes.Active[groupAddr] = activeAmt
		}
	}

	return votes, nil
}

type AddressLesserGreater struct {
	Lesser  common.Address
	Greater common.Address
}

func (w *ElectionWrapper) VoteMetadata(opts *bind.CallOpts, group common.Address, value *big.Int) (*AddressLesserGreater, error) {
	votes, err := w.Election.GetTotalVotesForEligibleValidatorGroups(opts)
	if err != nil {
		return nil, err
	}

	var voteTotal *big.Int = value
	for idx, currGroup := range votes.Groups {
		if group == currGroup {
			voteTotal = utils.Add(voteTotal, votes.Values[idx])
			break
		}
	}

	var result AddressLesserGreater
	// relies on votes.Groups being sorted
	for idx, currGroup := range votes.Groups {
		if group != currGroup {
			currValue := votes.Values[idx]
			if utils.IsLte(currValue, voteTotal) {
				result.Lesser = currGroup
				break
			}
			result.Greater = currGroup
		}
	}

	return &result, nil
}

func (w *ElectionWrapper) Vote(opts *bind.TransactOpts, group common.Address, value *big.Int) (*types.Transaction, error) {
	meta, err := w.VoteMetadata(CallOptsFromTxOpts(opts), group, value)
	if err != nil {
		return nil, err
	}

	return w.Election.Vote(opts, group, value, meta.Lesser, meta.Greater)
}

type RevokeMetadata struct {
	Index *big.Int
	Value *big.Int
	*AddressLesserGreater
}

func (w *ElectionWrapper) RevokeMetadata(opts *bind.CallOpts, account common.Address, group common.Address, value *big.Int) (*RevokeMetadata, error) {
	groups, err := w.Election.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	voteMeta, err := w.VoteMetadata(opts, group, utils.Neg(value))
	if err != nil {
		return nil, err
	}

	idx, err := utils.AddressIndexOf(groups, group)
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

func (w *ElectionWrapper) RevokeBothMetadata(opts *bind.CallOpts, account common.Address, group common.Address, value *big.Int) (*RevokeBothMetadata, error) {
	votesPending, err := w.Election.GetPendingVotesForGroupByAccount(opts, group, account)
	if err != nil {
		return nil, err
	}

	var result RevokeBothMetadata
	pendingValue := utils.Min(votesPending, value)
	if utils.IsNonZero(pendingValue) {
		result.Pending, err = w.RevokeMetadata(opts, account, group, pendingValue)
		if err != nil {
			return nil, err
		}
	}

	if utils.IsLt(pendingValue, value) {
		activeValue := utils.Sub(value, pendingValue)
		result.Active, err = w.RevokeMetadata(opts, account, group, activeValue)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

func (w *ElectionWrapper) RevokeAllMetadata(opts *bind.CallOpts, account common.Address, group common.Address) (*RevokeBothMetadata, error) {
	totalVotes, err := w.Election.GetTotalVotesForGroupByAccount(opts, group, account)
	if err != nil {
		return nil, err
	}

	return w.RevokeBothMetadata(opts, account, group, totalVotes)
}

type RevokeAll struct {
	Pending *types.Transaction
	Active  *types.Transaction
}

func (w *ElectionWrapper) RevokeAll(opts *bind.TransactOpts, account common.Address, group common.Address) (*RevokeAll, error) {
	meta, err := w.RevokeAllMetadata(CallOptsFromTxOpts(opts), account, group)
	if err != nil {
		return nil, err
	}

	pendingTx, err := w.Election.RevokePending(opts, group, meta.Pending.Value, meta.Pending.Lesser, meta.Pending.Greater, meta.Pending.Index)
	if err != nil {
		return nil, err
	}

	activeTx, err := w.Election.RevokeActive(opts, group, meta.Active.Value, meta.Active.Lesser, meta.Active.Greater, meta.Active.Index)
	if err != nil {
		return nil, err
	}

	result := RevokeAll{
		Pending: pendingTx,
		Active:  activeTx,
	}
	return &result, nil
}

func (w *ElectionWrapper) ActivateAllMetadata(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	groups, err := w.Election.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var activatableGroups []common.Address
	for _, group := range groups {
		activatable, err := w.Election.HasActivatablePendingVotes(opts, account, group)
		if err != nil {
			return nil, err
		}
		if activatable {
			activatableGroups = append(activatableGroups, group)
		}
	}
	return activatableGroups, nil
}

func (w *ElectionWrapper) ActivateAll(opts *bind.TransactOpts) ([]*types.Transaction, error) {
	groups, err := w.ActivateAllMetadata(CallOptsFromTxOpts(opts), opts.From)
	if err != nil {
		return nil, err
	}

	activateTransactions := make([]*types.Transaction, len(groups))
	for idx, group := range groups {
		activateTransactions[idx], err = w.Election.Activate(opts, group)
		if err != nil {
			return nil, err
		}
	}
	return activateTransactions, nil
}
