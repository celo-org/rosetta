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

package transaction

import (
	"fmt"

	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type AbiBuilder struct {
	abiMap map[wrapper.RegistryKey]*abi.ABI
}

func NewAbiBuilder() (*AbiBuilder, error) {
	abiMap := make(map[wrapper.RegistryKey]*abi.ABI)

	var abi *abi.ABI
	var err error

	abi, err = contract.ParseAccountsABI()
	if err != nil {
		return nil, err
	}
	abiMap[wrapper.AccountsRegistryId] = abi

	abi, err = contract.ParseElectionABI()
	if err != nil {
		return nil, err
	}
	abiMap[wrapper.ElectionRegistryId] = abi

	abi, err = contract.ParseLockedGoldABI()
	if err != nil {
		return nil, err
	}
	abiMap[wrapper.LockedGoldRegistryId] = abi

	return &AbiBuilder{
		abiMap: abiMap,
	}, nil
}

func (b *AbiBuilder) GetData(method CeloMethod, args ...interface{}) ([]byte, error) {
	contractABI, ok := b.abiMap[CeloMethodToRegistryKey[method]]
	if !ok {
		return nil, fmt.Errorf("No abi registered for method: %s", method)
	}

	return contractABI.Pack(method.String(), args...)
}
