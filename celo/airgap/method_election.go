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

package airgap

import (
	"context"

	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type electionAirGapMethods struct {
	abi      *abi.ABI
	registry *wrapper.RegistryWrapper
}

func newElectionAirGapMethodGroup(registry *wrapper.RegistryWrapper) (airGapMethodGroup, error) {
	abi, err := contract.ParseElectionABI()
	if err != nil {
		return nil, err
	}

	return &electionAirGapMethods{
		abi:      abi,
		registry: registry,
	}, nil

}

func (eag *electionAirGapMethods) register(server *airgGapServerImpl) {
	server.registerMethod(Vote, eag.voteMethod)
	server.registerMethod(RevokeActiveVotes, eag.revokeActiveVotes)
	server.registerMethod(RevokePendingVotes, eag.revokePendingVotes)
	server.registerMethod(ActivateVotes, genericAirGapMethodFactory(eag.registry, eag.abi, ActivateVotes))
}

func (eag *electionAirGapMethods) address(ctx context.Context) (common.Address, error) {
	return eag.registry.GetAddressForString(ctx, nil, wrapper.ElectionRegistryId.String())
}

func (eag *electionAirGapMethods) voteMethod(ctx context.Context, args []interface{}) ([]byte, common.Address, error) {
	err := validateArgLength(args, 2)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	group, err := parseAddress(args, 0)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	value, err := parseBigInt(args, 1)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	election, err := eag.registry.GetElectionWrapper(ctx, nil)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	keys, err := election.VoteMetadata(&bind.CallOpts{
		Context: ctx,
	}, group, value)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	data, err := eag.abi.Pack(Vote.Name, group, value, keys.Lesser, keys.Greater)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	address, err := eag.address(ctx)
	if err != nil {
		return nil, common.ZeroAddress, err
	}
	return data, address, nil
}

func (eag *electionAirGapMethods) revokeActiveVotes(ctx context.Context, args []interface{}) ([]byte, common.Address, error) {
	return eag.revokeMethod(ctx, RevokeActiveVotes.Name, args)
}

func (eag *electionAirGapMethods) revokePendingVotes(ctx context.Context, args []interface{}) ([]byte, common.Address, error) {
	return eag.revokeMethod(ctx, RevokePendingVotes.Name, args)
}

func (eag *electionAirGapMethods) revokeMethod(ctx context.Context, methodName string, args []interface{}) ([]byte, common.Address, error) {

	err := validateArgLength(args, 3)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	account, err := parseAddress(args, 0)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	group, err := parseAddress(args, 1)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	value, err := parseBigInt(args, 2)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	election, err := eag.registry.GetElectionWrapper(ctx, nil)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	keys, err := election.RevokeMetadata(&bind.CallOpts{
		Context: ctx,
	}, account, group, value)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	data, err := eag.abi.Pack(methodName, group, value, keys.Lesser, keys.Greater)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	address, err := eag.address(ctx)
	if err != nil {
		return nil, common.ZeroAddress, err
	}
	return data, address, nil
}
