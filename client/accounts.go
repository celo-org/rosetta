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

package client

import (
	"errors"

	"github.com/celo-org/rosetta/service/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func ConstructCreateAccount(from *common.Address, metadata *rpc.TransactionMetadata) (*types.Transaction, error) {
	return nil, errors.New("Not implemented yet")
	// args, err := metadata.ABIMethod.Inputs.Pack()
	// if err != nil {
	// 	return nil, err
	// }

	// data := append(metadata.ABIMethod.ID(), args...)

	// return types.NewTransaction(
	// 	metadata.Nonce,
	// 	*metadata.To,
	// 	big.NewInt(0),
	// 	metadata.GasLimit,
	// 	metadata.GasPrice,
	// 	nil, // guarantee cGLD
	// 	metadata.GatewayFeeRecipient,
	// 	metadata.GatewayFee,
	// 	data,
	// ), nil
}
