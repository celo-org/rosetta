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

package helpers

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/common/hexutil"
)

func callOptsFromTxOpts(txOpts *bind.TransactOpts) *bind.CallOpts {
	return &bind.CallOpts{
		From:    txOpts.From,
		Context: txOpts.Context,
	}
}

func FromFixed(number *big.Int) float64 {
	var fixed1, _ = new(big.Float).SetString("1000000000000000000000000")
	ret := new(big.Float)
	ret.Quo(new(big.Float).SetInt(number), fixed1)
	retF, _ := ret.Float64()
	return retF
}

func isSuffixEqual(s string, suffix string) bool {
	return len(suffix) <= len(s) && s[len(s)-len(suffix):] == suffix
}

func isFixidity(name string) bool {
	return isSuffixEqual(name, "Factor") || isSuffixEqual(name, "Fraction") || isSuffixEqual(name, "Multiplier") || isSuffixEqual(name, "Ratio") || isSuffixEqual(name, "Rate")
}

// EventToSlice transforms an abigen event struct binding to a string/int64/float64 encoded slice (ideal for BigQuery).
//
//	type AttestationsAttestationsRequested struct {
//		Identifier                 [32]byte
//		Account                    common.Address
//		AttestationsRequested      *big.Int
//		AttestationRequestFeeToken common.Address
//		Raw                        types.Log // Blockchain specific contextual infos
//	}
func EventToSlice(event interface{}) ([]interface{}, error) {
	v := reflect.ValueOf(event)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("eventToSlice only accepts structs; got %T", v)
	}

	slice := make([]interface{}, (v.NumField()-1)*2) // exclude raw, include field names

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fi := typ.Field(i)

		if fi.Name == "Raw" { // skip raw
			continue
		}

		var out interface{}
		switch v := v.Field(i).Interface().(type) {
		case common.Address:
			out = v.Hex()
		case *big.Int:
			if isFixidity(fi.Name) {
				out = FromFixed(v)
			} else {
				out = v.String()
			}
		case []byte:
			out = hexutil.Encode(v)
		case [32]byte:
			out = hexutil.Encode(v[:])
		case [4]byte:
			out = hexutil.Encode(v[:])
		default:
			out = v
		}

		slice[2*i] = fi.Name
		slice[2*i+1] = out
	}

	return slice, nil
}
