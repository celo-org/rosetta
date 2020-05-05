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

	"github.com/celo-org/rosetta/celo/airgap"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type argumentsParser func(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error)
type contractMethods struct {
	abiFactory func() (*abi.ABI, error)
	methods    map[*airgap.CeloMethod]argumentsParser
}

var serverMethods = []contractMethods{
	{
		abiFactory: contract.ParseAccountsABI,
		methods: map[*airgap.CeloMethod]argumentsParser{
			airgap.AuthorizeVoteSigner: authorizeVoteSignerParser,
			airgap.CreateAccount:       zeroArgumentsParser,
		},
	},
	{
		abiFactory: contract.ParseElectionABI,
		methods: map[*airgap.CeloMethod]argumentsParser{
			airgap.Vote:               voteMethodParser,
			airgap.RevokeActiveVotes:  revokeParser,
			airgap.RevokePendingVotes: revokeParser,
			airgap.ActivateVotes:      genericParser(addressParser),
		},
	},
	{
		abiFactory: contract.ParseLockedGoldABI,
		methods: map[*airgap.CeloMethod]argumentsParser{
			airgap.LockGold:     zeroArgumentsParser,
			airgap.UnlockGold:   genericParser(bigIntParser),
			airgap.RelockGold:   genericParser(bigIntParser, bigIntParser),
			airgap.WithdrawGold: genericParser(bigIntParser),
		},
	},
}

func hydrateMethods(srvCtx ServerContext, definitions []contractMethods) (map[*airgap.CeloMethod]airGapServerMethod, error) {
	serverMethods := make(map[*airgap.CeloMethod]airGapServerMethod)

	for _, cm := range definitions {
		abi, err := cm.abiFactory()
		if err != nil {
			return nil, err
		}

		for method, argParser := range cm.methods {
			serverMethods[method] = airgapMethodFactory(srvCtx, abi, argParser, method)
		}
	}
	return serverMethods, nil
}

func airgapMethodFactory(srvCtx ServerContext, abi *abi.ABI, argsParser argumentsParser, method *airgap.CeloMethod) airGapServerMethod {
	return func(ctx context.Context, args []interface{}) ([]byte, common.Address, error) {

		args, err := argsParser(ctx, srvCtx, args)
		if err != nil {
			return nil, common.ZeroAddress, err
		}

		data, err := abi.Pack(method.Name, args...)
		if err != nil {
			return nil, common.ZeroAddress, err
		}

		address, err := srvCtx.addressFor(ctx, method.Contract)
		if err != nil {
			return nil, common.ZeroAddress, err
		}

		return data, address, nil
	}
}
