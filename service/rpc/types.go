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
	"math/big"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
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

type Method = string

const (
	TransferMethod Method = "transfer"
)

var (
	DummyAddress = common.HexToAddress("abc")
)

//go:generate gencodec -type TransactionMetadata -out gen_transaction_metadata_json.go

type TransactionMetadata struct {
	Nonce               uint64          `json:"nonce"    gencodec:"required"`
	GasPrice            *big.Int        `json:"gasPrice" gencodec:"required"`
	GasLimit            uint64          `json:"gasLimit"      gencodec:"required"`
	GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient" rlp:"nil"` // nil means no gateway fee is paid
	GatewayFee          *big.Int        `json:"gatewayFee" rlp:"nil"`          // nil means no gateway fee is paid
}

func (txm *TransactionMetadata) asMessage() *ethereum.CallMsg {
	return &ethereum.CallMsg{
		GasPrice:            txm.GasPrice,
		GatewayFee:          txm.GatewayFee,
		GatewayFeeRecipient: txm.GatewayFeeRecipient,
	}
}

//go:generate gencodec -type TransferMetadata -out gen_transfer_json.go

type TransferMetadata struct {
	Balance *big.Int             `json:"balance" gencodec:"required"`
	Tx      *TransactionMetadata `json:"tx"`
}
