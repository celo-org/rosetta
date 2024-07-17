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

package bn

import (
	"math/big"
)

var Big0 = big.NewInt(0)
var Big1 = big.NewInt(1)

func Neg(val *big.Int) *big.Int {
	return new(big.Int).Neg(val)
}

func Sub(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Sub(a, b)
}

func Add(a *big.Int, b *big.Int) *big.Int {
	return new(big.Int).Add(a, b)
}

// a < b
func IsLt(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == -1
}

// a <= b
func IsLte(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) < 1
}

// a > b
func IsGt(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == 1
}

// a >= b
func IsGte(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) > -1
}

// a == b
func IsEq(a *big.Int, b *big.Int) bool {
	return a.Cmp(b) == 0
}

// a == 0
func IsZero(a *big.Int) bool {
	return IsEq(a, Big0)
}

func IsNonZero(a *big.Int) bool {
	return !IsZero(a)
}

func Min(a *big.Int, b *big.Int) *big.Int {
	if IsLt(a, b) {
		return a
	} else {
		return b
	}
}
