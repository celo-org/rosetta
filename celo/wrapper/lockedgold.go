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

type LockedGoldWrapper struct {
	*contract.LockedGold
}

func NewLockedGold(celoClient *client.CeloClient, registryWrapper *RegistryWrapper) (*LockedGoldWrapper, error) {
	lockedgold, err := registryWrapper.GetLockedGold(nil, celoClient.Eth)
	return &LockedGoldWrapper{lockedgold}, err
}

type PendingWithdrawal struct {
	Amount    *big.Int
	Timestamp *big.Int
}

type NonVotingLockedGold struct {
	Amount             *big.Int
	PendingWithdrawals []PendingWithdrawal
}

func (w *LockedGoldWrapper) GetTotalPendingWithdrawals(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	values, _, err := w.LockedGold.GetPendingWithdrawals(opts, account)
	if err != nil {
		return nil, err
	}

	return utils.Sum(values...), nil
}

func (w *LockedGoldWrapper) GetPendingWithdrawals(opts *bind.CallOpts, account common.Address) ([]PendingWithdrawal, error) {
	values, timestamps, err := w.LockedGold.GetPendingWithdrawals(opts, account)
	if err != nil {
		return nil, err
	}

	withdrawals := make([]PendingWithdrawal, len(values))
	for idx, val := range values {
		withdrawals[idx].Amount = val
		withdrawals[idx].Timestamp = timestamps[idx]
	}
	return withdrawals, nil
}
