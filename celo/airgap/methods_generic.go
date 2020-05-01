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
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

func genericAirGapMethodFactory(registry *wrapper.RegistryWrapper, abi *abi.ABI, method *CeloMethod) airGapServerMethod {
	return func(ctx context.Context, args []interface{}) ([]byte, common.Address, error) {
		data, err := abi.Pack(method.Name, args...)
		if err != nil {
			return nil, common.ZeroAddress, err
		}
		address, err := registry.GetAddressForString(ctx, nil, method.Contract.String())
		if err != nil {
			return nil, common.ZeroAddress, err
		}
		return data, address, nil
	}
}

type abiFactory func() (*abi.ABI, error)
type airGapMethodGroupFactory func(registry *wrapper.RegistryWrapper) (airGapMethodGroup, error)
type airGapMethodGroup interface {
	register(server *airgGapServerImpl)
}

type genericAirGapMethodGroup struct {
	abi      *abi.ABI
	registry *wrapper.RegistryWrapper
	methods  []*CeloMethod
}

func newGenericAirGapMethodGroup(registry *wrapper.RegistryWrapper, abiFactory abiFactory, methods []*CeloMethod) (airGapMethodGroup, error) {
	abi, err := abiFactory()
	if err != nil {
		return nil, err
	}

	return &genericAirGapMethodGroup{
		abi:      abi,
		registry: registry,
		methods:  methods,
	}, nil
}

func (eag *genericAirGapMethodGroup) register(server *airgGapServerImpl) {
	for _, method := range eag.methods {
		server.registerMethod(method, genericAirGapMethodFactory(eag.registry, eag.abi, method))
	}
}

////////////////////////////////////////////////////////////////////////////////////////
// Argument Parsers
////////////////////////////////////////////////////////////////////////////////////////

func argTypeErr(arg interface{}, index int, expected string) error {
	return fmt.Errorf("Argument %v provided at index %d needs type %s", arg, index, expected)
}

func parseAddress(args []interface{}, index int) (common.Address, error) {
	addr, ok := args[index].(common.Address)
	if !ok {
		return addr, argTypeErr(args, index, "common.Address")
	}
	return addr, nil
}

func parseBigInt(args []interface{}, index int) (*big.Int, error) {
	bigInt, ok := args[index].(*big.Int)
	if !ok {
		return nil, argTypeErr(args, index, "*big.Int")
	}
	return bigInt, nil
}

func parseBytes(args []interface{}, index int) ([]byte, error) {
	bytes, ok := args[index].([]byte)
	if !ok {
		return nil, argTypeErr(args, index, "[]byte")
	}
	return bytes, nil
}

func validateArgLength(args []interface{}, length int) error {
	if len(args) != length {
		return fmt.Errorf("Received %d args; expected %d", len(args), length)
	}
	return nil
}
