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

package helpers

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/rosetta/kliento/contracts"
	"github.com/celo-org/rosetta/kliento/utils/bn"
)

// LockedGold helper methods
type LockedGold struct{ *contracts.LockedGold }

type PendingWithdrawal struct {
	Amount    *big.Int
	Timestamp *big.Int
}

type NonVotingLockedGold struct {
	Amount             *big.Int
	PendingWithdrawals []PendingWithdrawal
}

func (lg *LockedGold) GetTotalPendingWithdrawals(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	values, _, err := lg.LockedGold.GetPendingWithdrawals(opts, account)
	if err != nil {
		return nil, err
	}

	return bn.Sum(values...), nil
}

func (lg *LockedGold) GetPendingWithdrawals(opts *bind.CallOpts, account common.Address) ([]PendingWithdrawal, error) {
	values, timestamps, err := lg.LockedGold.GetPendingWithdrawals(opts, account)
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
