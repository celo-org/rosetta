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
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewSigner(chainId *big.Int) types.EIP155Signer {
	return types.NewEIP155Signer(chainId)
}

func SignTransaction(transaction *types.Transaction, privateKey *ecdsa.PrivateKey, signer types.Signer) (*types.Transaction, error) {
	return types.SignTx(transaction, signer, privateKey)
}

func ParseSignedTransaction(transaction *types.Transaction, signer *types.EIP155Signer) (*common.Address, *common.Address, *big.Int, error) {
	from, err := signer.Sender(transaction)
	if err != nil {
		return nil, nil, nil, err
	}
	return &from, transaction.To(), transaction.Value(), nil
}

func HashTransaction(transaction *types.Transaction) common.Hash {
	return transaction.Hash()
}
