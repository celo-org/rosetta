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
	"fmt"

	"github.com/celo-org/rosetta/airgap"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var abiFactoryMap = map[wrapper.RegistryKey]func() (*abi.ABI, error){
	wrapper.AccountsRegistryId:   contract.ParseAccountsABI,
	wrapper.ElectionRegistryId:   contract.ParseElectionABI,
	wrapper.LockedGoldRegistryId: contract.ParseLockedGoldABI,
}

type argsPreProcessor func(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error)

var serverMethodsDefinitions = map[*airgap.CeloMethod]argsPreProcessor{
	airgap.CreateAccount:       noopArgsPreProcessor,
	airgap.AuthorizeVoteSigner: authorizeVoteSignerParser,
	airgap.LockGold:            noopArgsPreProcessor,
	airgap.UnlockGold:          noopArgsPreProcessor,
	airgap.RelockGold:          noopArgsPreProcessor,
	airgap.WithdrawGold:        noopArgsPreProcessor,

	airgap.Vote:          voteMethodParser,
	airgap.ActivateVotes: noopArgsPreProcessor,

	airgap.RevokePendingVotes: revokeParser,
	airgap.RevokeActiveVotes:  revokeParser,
}

func noopArgsPreProcessor(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error) {
	return args, nil
}

func hydrateMethods(srvCtx ServerContext) (map[*airgap.CeloMethod]airGapServerMethod, error) {
	abis := make(map[wrapper.RegistryKey]*abi.ABI)
	for id, abiFactory := range abiFactoryMap {
		abi, err := abiFactory()
		if err != nil {
			return nil, err
		}
		abis[id] = abi
	}

	serverMethods := make(map[*airgap.CeloMethod]airGapServerMethod)
	for method, preProcessor := range serverMethodsDefinitions {
		abi, ok := abis[method.Contract]
		if !ok {
			return nil, fmt.Errorf("Missing abi mapping for %s", method.Contract)
		}

		serverMethods[method] = airgapMethodFactory(srvCtx, abi, preProcessor, method)
	}
	return serverMethods, nil
}

func airgapMethodFactory(srvCtx ServerContext, abi *abi.ABI, argsParser argsPreProcessor, method *airgap.CeloMethod) airGapServerMethod {
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
