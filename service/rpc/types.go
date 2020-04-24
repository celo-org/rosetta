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

package rpc

import (
	"github.com/coinbase/rosetta-sdk-go/types"
)

var (
	CeloGold = &types.Currency{
		Symbol:   "cGLD",
		Decimals: 18,
	}
	CeloDollar = &types.Currency{
		Symbol:   "cUSD",
		Decimals: 18,
	}
)

type OperationResult string

const (
	OperationSuccess OperationResult = "success"
	OperationFailed  OperationResult = "failed"
)

func (or OperationResult) String() string { return string(or) }

func (or OperationResult) ToOperationStatus() *types.OperationStatus {
	return &types.OperationStatus{
		Successful: or == OperationSuccess,
		Status:     string(or),
	}
}

func GetOperationStatus(success bool) OperationResult {
	if success {
		return OperationSuccess
	} else {
		return OperationFailed
	}
}
