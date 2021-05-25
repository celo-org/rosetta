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

	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/registry"
	"github.com/celo-org/rosetta/airgap"
)

var abiFactoryMap = map[string]func() (*abi.ABI, error){
	registry.AccountsContractID.String():    contracts.ParseAccountsABI,
	registry.ElectionContractID.String():    contracts.ParseElectionABI,
	registry.LockedGoldContractID.String():  contracts.ParseLockedGoldABI,
	airgap.ReleaseGold:                      contracts.ParseReleaseGoldABI,
	registry.StableTokenContractID.String(): contracts.ParseStableTokenABI,
}

type argsPreProcessor func(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error)

var serverTransactionMethodDefinitions = map[*airgap.CeloMethod]argsPreProcessor{
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

	airgap.StableTokenTransfer: noopArgsPreProcessor,
}

var serverCallMethodDefinitions = map[*airgap.CeloMethod]argsPreProcessor{
	airgap.IsAccount:                             noopArgsPreProcessor,
	airgap.GetVoteSigner:                         noopArgsPreProcessor,
	airgap.CanReceiveVotes:                       noopArgsPreProcessor,
	airgap.GetEpochNumber:                        noopArgsPreProcessor,
	airgap.GetEpochSize:                          noopArgsPreProcessor,
	airgap.GetEpochNumberOfBlock:                 noopArgsPreProcessor,
	airgap.GetGroupsVotedForByAccount:            noopArgsPreProcessor,
	airgap.GetActiveVotesForGroup:                noopArgsPreProcessor,
	airgap.GetActiveVotesForGroupByAccount:       noopArgsPreProcessor,
	airgap.ReleaseGoldMaxDistribution:            noopArgsPreProcessor,
	airgap.ReleaseGoldTotalWithdrawn:             noopArgsPreProcessor,
	airgap.ReleaseGoldBeneficiary:                noopArgsPreProcessor,
	airgap.ReleaseGoldReleaseOwner:               noopArgsPreProcessor,
	airgap.ReleaseGoldRefundAddress:              noopArgsPreProcessor,
	airgap.ReleaseGoldRevocationInfo:             noopArgsPreProcessor,
	airgap.ReleaseGoldIsRevoked:                  noopArgsPreProcessor,
	airgap.ReleaseGoldCurrentReleasedTotalAmount: noopArgsPreProcessor,
	airgap.ReleaseGoldTotalBalance:               noopArgsPreProcessor,
	airgap.ReleaseGoldRemainingTotalBalance:      noopArgsPreProcessor,
	airgap.ReleaseGoldRemainingUnlockedBalance:   noopArgsPreProcessor,
	airgap.ReleaseGoldRemainingLockedBalance:     noopArgsPreProcessor,
	airgap.StableTokenBalanceOf:                  noopArgsPreProcessor,
}

func noopArgsPreProcessor(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error) {
	return args, nil
}

func hydrateMethods(srvCtx ServerContext, methods map[*airgap.CeloMethod]argsPreProcessor) (map[*airgap.CeloMethod]airGapServerMethod, error) {
	abis := make(map[string]*abi.ABI)
	for id, abiFactory := range abiFactoryMap {
		abi, err := abiFactory()
		if err != nil {
			return nil, err
		}
		abis[id] = abi
	}

	mappedMethods := make(map[*airgap.CeloMethod]airGapServerMethod)
	for method, preProcessor := range methods {
		abi, ok := abis[method.Contract]
		if !ok {
			return nil, fmt.Errorf("Missing abi mapping for %s", method.Contract)
		}

		mappedMethods[method] = airgapMethodFactory(srvCtx, abi, preProcessor, method)
	}
	return mappedMethods, nil
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
