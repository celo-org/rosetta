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
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
)

type argParser func(interface{}) (interface{}, error)

func addressParser(value interface{}) (interface{}, error) {
	switch v := value.(type) {
	case string:
		return common.HexToAddress(v), nil
	case common.Address:
		return v, nil
	default:
		return common.ZeroAddress, fmt.Errorf("Not a valid address: %v", v)
	}
}

func bytesParser(value interface{}) (interface{}, error) {
	var res []byte
	switch v := value.(type) {
	case string:
		res = common.FromHex(v)
	case []uint8:
		res = v
	default:
		return nil, fmt.Errorf("Not a valid []byte: %v", v)
	}
	return res, nil
}

func bigIntParser(value interface{}) (interface{}, error) {
	var ret *big.Int
	switch v := value.(type) {
	case string:
		val, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return nil, fmt.Errorf("Not a big.Int: %s", val)
		}
		return val, nil
	case int:
		ret = new(big.Int).SetInt64(int64(v))
	case int64:
		ret = new(big.Int).SetInt64(v)
	case uint:
		ret = new(big.Int).SetUint64(uint64(v))
	case uint64:
		ret = new(big.Int).SetUint64(v)
	case float32:
		ret = new(big.Int).SetInt64(int64(v))
	case float64:
		ret = new(big.Int).SetInt64(int64(v))
	case *big.Int:
		ret = v
	default:
		return nil, fmt.Errorf("Not a big.Int: %v", v)
	}
	return ret, nil
}
