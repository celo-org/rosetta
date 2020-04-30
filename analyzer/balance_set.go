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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceSet struct {
	balances map[common.Address]*big.Int
}

func NewBalanceSet() *BalanceSet {
	return &BalanceSet{
		balances: make(map[common.Address]*big.Int),
	}
}

func (bs *BalanceSet) Add(addr common.Address, value *big.Int) *BalanceSet {
	if oldValue, ok := bs.balances[addr]; ok {
		bs.balances[addr] = new(big.Int).Add(oldValue, value)
	} else {
		bs.balances[addr] = value
	}
	return bs
}

func (bs *BalanceSet) ToMap() map[common.Address]*big.Int {
	return bs.balances
}
