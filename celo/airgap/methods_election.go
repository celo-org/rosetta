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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type argumentsParser func(ctx context.Context, registry *wrapper.RegistryWrapper, args []interface{}) ([]interface{}, error)

func electionMethodGroup(registry *wrapper.RegistryWrapper) (map[*CeloMethod]airGapServerMethod, error) {
	abi, err := contract.ParseElectionABI()
	if err != nil {
		return nil, err
	}

	methods := make(map[*CeloMethod]airGapServerMethod)
	methods[Vote] = airgapMethodFactory(registry, abi, voteMethodParser, Vote)
	methods[RevokeActiveVotes] = airgapMethodFactory(registry, abi, revokeParser, RevokeActiveVotes)
	methods[RevokePendingVotes] = airgapMethodFactory(registry, abi, revokeParser, RevokePendingVotes)
	methods[ActivateVotes] = airgapMethodFactory(registry, abi, nil, ActivateVotes)
	return methods, nil
}

func voteMethodParser(ctx context.Context, registry *wrapper.RegistryWrapper, args []interface{}) ([]interface{}, error) {
	err := validateArgLength(args, 2)
	if err != nil {
		return nil, err
	}

	group, err := parseAddress(args, 0)
	if err != nil {
		return nil, err
	}

	value, err := parseBigInt(args, 1)
	if err != nil {
		return nil, err
	}

	election, err := registry.GetElectionWrapper(ctx, nil)
	if err != nil {
		return nil, err
	}

	keys, err := election.VoteMetadata(&bind.CallOpts{
		Context: ctx,
	}, group, value)
	if err != nil {
		return nil, err
	}

	return []interface{}{group, value, keys.Lesser, keys.Greater}, nil
}

func revokeParser(ctx context.Context, registry *wrapper.RegistryWrapper, args []interface{}) ([]interface{}, error) {

	err := validateArgLength(args, 3)
	if err != nil {
		return nil, err
	}

	account, err := parseAddress(args, 0)
	if err != nil {
		return nil, err
	}

	group, err := parseAddress(args, 1)
	if err != nil {
		return nil, err
	}

	value, err := parseBigInt(args, 2)
	if err != nil {
		return nil, err
	}

	election, err := registry.GetElectionWrapper(ctx, nil)
	if err != nil {
		return nil, err
	}

	keys, err := election.RevokeMetadata(&bind.CallOpts{
		Context: ctx,
	}, account, group, value)
	if err != nil {
		return nil, err
	}

	return []interface{}{group, value, keys.Lesser, keys.Greater}, nil
}
