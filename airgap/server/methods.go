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

	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/registry"
	"github.com/celo-org/rosetta/airgap"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var abiFactoryMap = map[string]func() (*abi.ABI, error){
	registry.AccountsContractID.String():   contracts.ParseAccountsABI,
	registry.ElectionContractID.String():   contracts.ParseElectionABI,
	registry.LockedGoldContractID.String(): contracts.ParseLockedGoldABI,
	airgap.ReleaseGold:                     contracts.ParseReleaseGoldABI,
}

type argsPreProcessor func(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error)

var serverMethodsDefinitions = map[*airgap.CeloMethod]argsPreProcessor{
	airgap.CreateAccount:              noopArgsPreProcessor,
	airgap.AuthorizeVoteSigner:        preprocessAuthorizeSigner,
	airgap.AuthorizeAttestationSigner: preprocessAuthorizeSigner,
	airgap.AuthorizeValidatorSigner:   preprocessAuthorizeSigner,

	airgap.LockGold:     noopArgsPreProcessor,
	airgap.UnlockGold:   noopArgsPreProcessor,
	airgap.RelockGold:   noopArgsPreProcessor,
	airgap.WithdrawGold: noopArgsPreProcessor,

	airgap.Vote:               preprocessVote,
	airgap.ActivateVotes:      noopArgsPreProcessor,
	airgap.RevokePendingVotes: preprocessRevoke,
	airgap.RevokeActiveVotes:  preprocessRevoke,

	airgap.ReleaseGoldWithdraw:                   noopArgsPreProcessor,
	airgap.ReleaseGoldCreateAccount:              noopArgsPreProcessor,
	airgap.ReleaseGoldLockGold:                   noopArgsPreProcessor,
	airgap.ReleaseGoldUnlockGold:                 noopArgsPreProcessor,
	airgap.ReleaseGoldRelockGold:                 noopArgsPreProcessor,
	airgap.ReleaseGoldWithdrawGold:               noopArgsPreProcessor,
	airgap.ReleaseGoldAuthorizeVoteSigner:        preprocessAuthorizeSigner,
	airgap.ReleaseGoldAuthorizeAttestationSigner: preprocessAuthorizeSigner,
	airgap.ReleaseGoldAuthorizeValidatorSigner:   preprocessAuthorizeSigner,
	airgap.ReleaseGoldRevokePendingVotes:         preprocessRevoke,
	airgap.ReleaseGoldRevokeActiveVotes:          preprocessRevoke,
}

func noopArgsPreProcessor(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error) {
	return args, nil
}

func hydrateMethods(srvCtx ServerContext) (map[*airgap.CeloMethod]airGapServerMethod, error) {
	abis := make(map[string]*abi.ABI)
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
	return func(ctx context.Context, args []interface{}) ([]byte, error) {

		args, err := argsParser(ctx, srvCtx, args)
		if err != nil {
			return nil, err
		}

		data, err := abi.Pack(method.Name, args...)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
}
