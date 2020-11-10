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
	"encoding/json"
	"fmt"
	"math/big"
)

func MarshallToMap(input interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var output map[string]interface{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UnmarshallFromMap(input map[string]interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, output)
}

func bigIntToString(input *big.Int) *string {
	if input == nil {
		return nil
	}
	out := input.String()
	return &out
}

func stringToBigInt(input *string) (*big.Int, error) {
	if input == nil {
		return nil, nil
	}

	out, ok := new(big.Int).SetString(*input, 10)
	if !ok {
		return nil, fmt.Errorf("Invalid bigInt format: %s", *input)
	}
	return out, nil
}

func stringToMethod(input *string) (*CeloMethod, error) {
	if input == nil {
		return nil, nil
	}
	out, err := MethodFromString(*input)
	if err != nil {
		return nil, err
	}
	return out, nil
}
