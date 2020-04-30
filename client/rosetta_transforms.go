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

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func NewAccountIdentifier(addr *common.Address) *types.AccountIdentifier {
	return &types.AccountIdentifier{
		Address: addr.Hex(),
	}
}

func EncodeTransaction(transaction *types.Transaction) (string, error) {
	bytes, err := rlp.EncodeToBytes(transaction)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func DecodeTransaction(transaction *types.ConstructionSubmitRequest) (*types.Transaction, error) {
	bytes := []byte(transaction.SignedTransaction)
	var tx types.Transaction
	err := rlp.DecodeBytes(bytes, tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func EncodeAccount(publicKey *ecdsa.PublicKey, address *common.Address) *types.AccountIdentifier {
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return &types.AccountIdentifier{
		Address: address.String(),
		Metadata: map[string]interface{}{
			"publicKey": string(publicKeyBytes),
		},
	}
}
