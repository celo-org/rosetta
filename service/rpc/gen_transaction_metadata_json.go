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

// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package rpc

import (
	"encoding/json"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// MarshalJSON marshals as JSON.
func (t TransactionMetadata) MarshalJSON() ([]byte, error) {
	type TransactionMetadata struct {
		Nonce               uint64          `json:"nonce"    gencodec:"required"`
		GasPrice            *big.Int        `json:"gasPrice" gencodec:"required"`
		GasLimit            uint64          `json:"gasLimit"      gencodec:"required"`
		GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient" rlp:"nil"`
		GatewayFee          *big.Int        `json:"gatewayFee" rlp:"nil"`
	}
	var enc TransactionMetadata
	enc.Nonce = t.Nonce
	enc.GasPrice = t.GasPrice
	enc.GasLimit = t.GasLimit
	enc.GatewayFeeRecipient = t.GatewayFeeRecipient
	enc.GatewayFee = t.GatewayFee
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (t *TransactionMetadata) UnmarshalJSON(input []byte) error {
	type TransactionMetadata struct {
		Nonce               *uint64         `json:"nonce"    gencodec:"required"`
		GasPrice            *big.Int        `json:"gasPrice" gencodec:"required"`
		GasLimit            *uint64         `json:"gasLimit"      gencodec:"required"`
		GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient" rlp:"nil"`
		GatewayFee          *big.Int        `json:"gatewayFee" rlp:"nil"`
	}
	var dec TransactionMetadata
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Nonce == nil {
		return errors.New("missing required field 'nonce' for TransactionMetadata")
	}
	t.Nonce = *dec.Nonce
	if dec.GasPrice == nil {
		return errors.New("missing required field 'gasPrice' for TransactionMetadata")
	}
	t.GasPrice = dec.GasPrice
	if dec.GasLimit == nil {
		return errors.New("missing required field 'gasLimit' for TransactionMetadata")
	}
	t.GasLimit = *dec.GasLimit
	if dec.GatewayFeeRecipient != nil {
		t.GatewayFeeRecipient = dec.GatewayFeeRecipient
	}
	if dec.GatewayFee != nil {
		t.GatewayFee = dec.GatewayFee
	}
	return nil
}
