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
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
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
