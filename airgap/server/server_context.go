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

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/ethclient"
	"github.com/celo-org/rosetta/kliento/client"
	"github.com/celo-org/rosetta/kliento/contracts/helpers"
	"github.com/celo-org/rosetta/kliento/registry"
)

type ServerContext interface {
	addressFor(ctx context.Context, identifier registry.ContractID, blockNumber *big.Int) (common.Address, error)

	authorizeMetadata(ctx context.Context, popSignature []byte) (*helpers.SignatureValues, error)
	voteMetadata(ctx context.Context, group common.Address, value *big.Int) (*helpers.AddressLesserGreater, error)
	revokeMetadata(ctx context.Context, signerOrAccount common.Address, group common.Address, value *big.Int) (*helpers.RevokeMetadata, error)

	// From ethclient.Client
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	SendTransaction(ctx context.Context, tx *types.Transaction) error
}

type serverContextImpl struct {
	*ethclient.Client

	registry registry.Registry
}

func NewServerContext(cc *client.CeloClient) (ServerContext, error) {
	registry, err := registry.New(cc)
	if err != nil {
		return nil, err
	}

	return &serverContextImpl{
		Client:   cc.Eth,
		registry: registry,
	}, nil
}

func (sc *serverContextImpl) addressFor(ctx context.Context, identifier registry.ContractID, blockNumber *big.Int) (common.Address, error) {
	return sc.registry.GetAddressFor(ctx, blockNumber, identifier)
}

func (sc *serverContextImpl) authorizeMetadata(ctx context.Context, popSignature []byte) (*helpers.SignatureValues, error) {
	return helpers.GetSignatureValues(popSignature)
}

func (sc *serverContextImpl) voteMetadata(ctx context.Context, group common.Address, value *big.Int) (*helpers.AddressLesserGreater, error) {
	election, err := sc.registry.GetElectionContract(ctx, nil)
	if err != nil {
		return nil, err
	}
	he := helpers.Election{Election: election}
	return he.VoteMetadata(&bind.CallOpts{Context: ctx}, group, value)
}

func (sc *serverContextImpl) revokeMetadata(ctx context.Context, signerOrAccount common.Address, group common.Address, value *big.Int) (*helpers.RevokeMetadata, error) {
	callOpts := &bind.CallOpts{Context: ctx}
	accounts, err := sc.registry.GetAccountsContract(ctx, nil)
	if err != nil {
		return nil, err
	}
	account, err := accounts.SignerToAccount(callOpts, signerOrAccount)
	if err != nil {
		return nil, err
	}
	election, err := sc.registry.GetElectionContract(ctx, nil)
	if err != nil {
		return nil, err
	}
	he := helpers.Election{Election: election}
	return he.RevokeMetadata(callOpts, account, group, value)
}
