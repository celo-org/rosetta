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
	"fmt"
	"strings"
)

var methodRegistry = make(map[string]map[string]*CeloMethod)

// FromString returns the CeloMethod that matches the given string
// Methods are represented as "Contract.Name"
func MethodFromString(celoMethodStr string) (*CeloMethod, error) {
	parts := strings.Split(celoMethodStr, ".")
	if len(parts) != 2 {
		// If there is no . present, try to parse this string as a generic EVM method signature
		unregisteredMethod, err := unregisteredMethodFromString(celoMethodStr)
		if err != nil {
			return nil, err
		}
		return unregisteredMethod, nil
	}
	m, ok := methodRegistry[parts[0]]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	cm, ok := m[parts[1]]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	return cm, nil
}

func registerMethod(contract string, name string, argParsers []argParser) *CeloMethod {
	cm := &CeloMethod{
		Name:       name,
		Contract:   contract,
		argParsers: argParsers,
	}
	// register the method in the methodRegistry
	if methodRegistry[contract] == nil {
		methodRegistry[contract] = make(map[string]*CeloMethod)
	}
	methodRegistry[contract][name] = cm
	return cm
}

// unregisteredMethodFromString returns a CeloMethod representing an arbitrary EVM-compatible method
// Unknown methods are represented by their function signature string like "transfer(bytes32,address)"
func unregisteredMethodFromString(methodSignature string) (*CeloMethod, error) {
	if err := validateMethodSignature(methodSignature); err != nil {
		return nil, err
	}
	unregisteredMethod := CeloMethod{
		Name: methodSignature,
		// Unregistered methods do not identify contract by name
		// The "to" address value should be set to the contract's address
		Contract: "",
		// Unregistered methods do not use argParsers
		// Args will be pre-encoded OR parsed based on the methodSignature
		argParsers: nil,
	}
	return &unregisteredMethod, nil
}

func validateMethodSignature(methodSig string) error {
	splitSigByLeadingParenthesis := strings.Split(methodSig, "(")
	if len(splitSigByLeadingParenthesis) < 2 {
		return nil
	}
	splitSigByTrailingParenthesis := strings.Split(splitSigByLeadingParenthesis[1], ")")
	if len(splitSigByTrailingParenthesis) < 1 || splitSigByTrailingParenthesis[0] == "" {
		return nil
	}
	methodTypes := strings.Split(splitSigByTrailingParenthesis[0], ",")

	for _, v := range methodTypes {
		switch {
		case v == "address":
			continue
		case strings.HasPrefix(v, "uint") || strings.HasPrefix(v, "int"):
			continue
		case strings.HasPrefix(v, "bytes"):
			continue
		case v == "string":
			continue
		case v == "bool":
			continue
		default:
			return fmt.Errorf("Invalid type %s in method signature: %s", v, methodSig)
		}
	}

	return nil
}
