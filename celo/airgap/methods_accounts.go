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
	"github.com/ethereum/go-ethereum/common"
)

type accountsAirGapMethods struct {
	abi      *abi.ABI
	registry *wrapper.RegistryWrapper
}

func newAccountsAirGapMethods(registry *wrapper.RegistryWrapper) (airGapMethodGroup, error) {
	abi, err := contract.ParseAccountsABI()
	if err != nil {
		return nil, err
	}

	return &accountsAirGapMethods{
		abi:      abi,
		registry: registry,
	}, nil

}

func (eag *accountsAirGapMethods) register(server *airgGapServerImpl) {
	server.registerMethod(AuthorizeVoteSigner, eag.authorizeVoteSigner)
	server.registerMethod(CreateAccount, genericAirGapMethodFactory(eag.registry, eag.abi, CreateAccount))
}

func (eag *accountsAirGapMethods) address(ctx context.Context) (common.Address, error) {
	return eag.registry.GetAddressForString(ctx, nil, wrapper.AccountsRegistryId.String())
}

func (eag *accountsAirGapMethods) authorizeVoteSigner(ctx context.Context, args []interface{}) ([]byte, common.Address, error) {
	err := validateArgLength(args, 2)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	signer, err := parseAddress(args, 0)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	pop, err := parseBytes(args, 1)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	accounts, err := eag.registry.GetAccountsWrapper(ctx, nil)
	if err != nil {
		return nil, common.ZeroAddress, err
	}
	encodedSig, err := accounts.AuthorizeMetadata(pop)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	data, err := eag.abi.Pack(Vote.Name, signer, encodedSig.V, encodedSig.R, encodedSig.S)
	if err != nil {
		return nil, common.ZeroAddress, err
	}

	address, err := eag.address(ctx)
	if err != nil {
		return nil, common.ZeroAddress, err
	}
	return data, address, nil
}
