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

type topicParser func(interface{}) (common.Hash, error)

func addressTopicParser(value interface{}) (common.Hash, error) {
	switch v := value.(type) {
	case string:
		return common.HexToHash(v), nil
	case common.Address:
		return v.Hash(), nil
	default:
		return common.ZeroAddress.Hash(), fmt.Errorf("Not a valid address: %v", v)
	}
}

func bytesTopicParser(value interface{}) (common.Hash, error) {
	var res common.Hash
	switch v := value.(type) {
	case string:
		res = common.HexToHash(v)
	case []uint8:
		res = common.BytesToHash(v)
	default:
		return common.Hash{}, fmt.Errorf("Not a valid []byte: %v", v)
	}
	return res, nil
}

func bigIntTopicParser(value interface{}) (common.Hash, error) {
	var ret *big.Int
	switch v := value.(type) {
	case string:
		val, ok := new(big.Int).SetString(v, 10)
		if !ok {
			return common.Hash{}, fmt.Errorf("Not a big.Int: %s", val)
		}
		return common.BigToHash(val), nil
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
		return common.Hash{}, fmt.Errorf("Not a big.Int: %v", v)
	}
	return common.BigToHash(ret), nil
}
