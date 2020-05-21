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

package server

import (
	"context"
	"math/big"

	"github.com/celo-org/kliento/contracts/helpers"
	"github.com/celo-org/kliento/registry"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type serverContextStub struct{}

var _ ServerContext = &serverContextStub{}

func (sc *serverContextStub) addressFor(ctx context.Context, identifier registry.ContractID) (common.Address, error) {
	return common.ZeroAddress, nil
}

func (sc *serverContextStub) authorizeMetadata(ctx context.Context, popSignature []byte) (*helpers.SignatureValues, error) {
	return &helpers.SignatureValues{
		R: [32]byte{},
		S: [32]byte{},
		V: 31,
	}, nil
}

func (sc *serverContextStub) voteMetadata(ctx context.Context, group common.Address, value *big.Int) (*helpers.AddressLesserGreater, error) {
	return &helpers.AddressLesserGreater{
		Greater: common.HexToAddress("0x01"),
		Lesser:  common.HexToAddress("0x02"),
	}, nil
}

func (sc *serverContextStub) revokeMetadata(ctx context.Context, account common.Address, group common.Address, value *big.Int) (*helpers.RevokeMetadata, error) {
	return &helpers.RevokeMetadata{
		Index: big.NewInt(1),
		AddressLesserGreater: &helpers.AddressLesserGreater{
			Greater: common.HexToAddress("0x01"),
			Lesser:  common.HexToAddress("0x02"),
		},
		Value: big.NewInt(4),
	}, nil
}

func (sc *serverContextStub) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return big.NewInt(678), nil
}
func (sc *serverContextStub) EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	return 666, nil
}
func (sc *serverContextStub) NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error) {
	return 7, nil
}
func (sc *serverContextStub) SendRawTransaction(ctx context.Context, data []byte) (*common.Hash, error) {
	hash := common.HexToHash("0x666777888")
	return &hash, nil
}
