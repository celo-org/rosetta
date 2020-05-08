package airgap

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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
