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
	abiCache map[*wrapper.RegistryKey]*abi.ABI
}

func NewAbiBuilder() *AbiBuilder {
	return &AbiBuilder{
		abiCache: make(map[*wrapper.RegistryKey]*abi.ABI),
	}
}

func (b *AbiBuilder) getAbi(key *wrapper.RegistryKey) (*abi.ABI, error) {
	var abi *abi.ABI
	abi, present := b.abiCache[key]

	if !present {
		var err error
		switch key {
		case &wrapper.AccountsRegistryId:
			abi, err = contract.ParseAccountsABI()
		case &wrapper.ElectionRegistryId:
			abi, err = contract.ParseElectionABI()
		case &wrapper.LockedGoldRegistryId:
			abi, err = contract.ParseLockedGoldABI()
		default:
			err = fmt.Errorf("Unimplemented ABI parsing for %s", key)
		}
		if err != nil {
			return nil, err
		}
	}

	b.abiCache[key] = abi
	return abi, nil
}

func (b *AbiBuilder) GetData(method *CeloMethod, args ...interface{}) ([]byte, error) {
	if method == nil {
		return nil, fmt.Errorf("'Method' required for building tx data")
	}

	contractABI, err := b.getAbi(CeloMethodToRegistryKey[method])
	if err != nil {
		return nil, err
	}

	packedData, err := contractABI.Pack(method.String(), args...)
	if err != nil {
		return nil, err
	}

	return packedData, nil
}
