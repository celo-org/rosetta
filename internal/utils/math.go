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
	"math/big"
)

var Big0 = big.NewInt(0)
var Big1 = big.NewInt(1)

func Sum(values ...*big.Int) *big.Int {
	acc := new(big.Int)
	for _, val := range values {
		acc.Add(acc, val)
	}
	return acc
}

func Inc(val *big.Int) *big.Int {
	return new(big.Int).Add(val, Big1)
}

func Dec(val *big.Int) *big.Int {
	return new(big.Int).Sub(val, Big1)
}
