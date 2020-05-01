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
)

type accountsAirGapMethods struct {
	abi      *abi.ABI
	registry *wrapper.RegistryWrapper
}

func accountsMethodGroup(registry *wrapper.RegistryWrapper) (map[*CeloMethod]airGapServerMethod, error) {
	abi, err := contract.ParseAccountsABI()
	if err != nil {
		return nil, err
	}

	methods := make(map[*CeloMethod]airGapServerMethod)
	methods[AuthorizeVoteSigner] = airgapMethodFactory(registry, abi, authorizeVoteSignerParser, AuthorizeVoteSigner)
	methods[CreateAccount] = airgapMethodFactory(registry, abi, nil, CreateAccount)
	return methods, nil
}

func authorizeVoteSignerParser(ctx context.Context, registry *wrapper.RegistryWrapper, args []interface{}) ([]interface{}, error) {
	err := validateArgLength(args, 2)
	if err != nil {
		return nil, err
	}

	signer, err := parseAddress(args, 0)
	if err != nil {
		return nil, err
	}

	pop, err := parseBytes(args, 1)
	if err != nil {
		return nil, err
	}

	accounts, err := registry.GetAccountsWrapper(ctx, nil)
	if err != nil {
		return nil, err
	}
	encodedSig, err := accounts.AuthorizeMetadata(pop)
	if err != nil {
		return nil, err
	}

	return []interface{}{signer, encodedSig.V, encodedSig.R, encodedSig.S}, nil
}
