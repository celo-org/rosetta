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

package server

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type argParser func(value interface{}) (interface{}, error)

func genericParser(parsers ...argParser) argumentsParser {
	return func(ctx context.Context, srvCtx serverContext, values []interface{}) ([]interface{}, error) {
		parsedArgs := make([]interface{}, len(parsers))
		if len(values) != len(parsers) {
			return nil, fmt.Errorf("Received %d args; expected %d", len(values), len(parsers))
		}

		var err error
		for i, value := range values {
			parsedArgs[i], err = parsers[i](value)
			if err != nil {
				return nil, fmt.Errorf("bad argument: idx=%d error=%w", i, err)
			}
		}

		return parsedArgs, nil
	}
}

func zeroArgumentsParser(ctx context.Context, srvCtx serverContext, values []interface{}) ([]interface{}, error) {
	if len(values) != 0 {
		return nil, fmt.Errorf("Received %d args; expected %d", len(values), 0)
	}
	return nil, nil
}

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

func authorizeVoteSignerParser(ctx context.Context, srvCtx serverContext, args []interface{}) ([]interface{}, error) {
	parsedArgs, err := genericParser(addressParser, bytesParser)(ctx, srvCtx, args)
	if err != nil {
		return nil, err
	}

	encodedSig, err := srvCtx.authorizeMetadata(ctx, parsedArgs[1].([]byte))
	if err != nil {
		return nil, err
	}

	return []interface{}{parsedArgs[0], encodedSig.V, encodedSig.R, encodedSig.S}, nil
}

func voteMethodParser(ctx context.Context, srvCtx serverContext, args []interface{}) ([]interface{}, error) {
	parsedArgs, err := genericParser(addressParser, bigIntParser)(ctx, srvCtx, args)
	if err != nil {
		return nil, err
	}

	keys, err := srvCtx.voteMetadata(ctx, parsedArgs[0].(common.Address), parsedArgs[1].(*big.Int))
	if err != nil {
		return nil, err
	}

	return []interface{}{parsedArgs[0], parsedArgs[1], keys.Lesser, keys.Greater}, nil
}

func revokeParser(ctx context.Context, srvCtx serverContext, args []interface{}) ([]interface{}, error) {
	parsedArgs, err := genericParser(addressParser, addressParser, bigIntParser)(ctx, srvCtx, args)
	if err != nil {
		return nil, err
	}

	keys, err := srvCtx.revokeMetadata(ctx, parsedArgs[0].(common.Address), parsedArgs[1].(common.Address), parsedArgs[2].(*big.Int))
	if err != nil {
		return nil, err
	}

	return []interface{}{parsedArgs[1], parsedArgs[2], keys.Lesser, keys.Greater, keys.Index}, nil
}

func noopArgParser(ctx context.Context, srvCtx serverContext, args []interface{}) ([]interface{}, error) {
	return args, nil
}
