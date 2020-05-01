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

package utils

import (
	"errors"
	"math/big"
)

var ErrNotImplemented = errors.New("Not implemented")

// 10^(24) == 1000000000000000000000000
var TobinTaxDenominator = new(big.Int).SetBytes([]byte{211, 194, 27, 206, 204, 237, 161, 0, 0, 0})

func CalcTobinTaxAmount(transferAmount, tobinTaxNumerator *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(transferAmount, tobinTaxNumerator), TobinTaxDenominator)
}
