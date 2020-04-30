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
		if pendingAmt.Cmp(utils.Big0) == 1 {
			votes.Pending[groupAddr] = pendingAmt
		}

		activeAmt, err := w.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if activeAmt.Cmp(utils.Big0) == 1 {
			votes.Active[groupAddr] = activeAmt
		}
	}

	return votes, nil
}
