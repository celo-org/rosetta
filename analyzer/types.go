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

	"github.com/celo-org/rosetta/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

type ChainParameters struct {
	ChainId   *big.Int
	EpochSize uint64
}

func (cp *ChainParameters) IsLastBlockOfEpoch(blockNumber uint64) bool {
	return istanbul.IsLastBlockOfEpoch(blockNumber, cp.EpochSize)
}

type Transfer struct {
	From   Account
	To     Account
	Value  *big.Int
	Status bool
}

// 10^(24) == 1000000000000000000000000
var TobinTaxDenominator *big.Int = new(big.Int).Exp(big.NewInt(10), big.NewInt(24), nil)

type TobinTax struct {
	Numerator *big.Int
	Recipient common.Address
}

func NewTobinTax(numerator *big.Int, recipient common.Address) *TobinTax {
	return &TobinTax{
		Numerator: numerator,
		Recipient: recipient,
	}
}

func (tt *TobinTax) Apply(value *big.Int) (taxAmount, afterTaxAmount *big.Int) {
	taxAmount = new(big.Int).Div(new(big.Int).Mul(value, tt.Numerator), TobinTaxDenominator)
	afterTaxAmount = new(big.Int).Sub(value, taxAmount)
	return
}

func (tt *TobinTax) IsDefined() bool {
	return tt.Numerator.Cmp(utils.Big0) > 0
}
