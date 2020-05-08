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
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func TestArgParsers(t *testing.T) {
	RegisterTestingT(t)

	type testCase struct {
		input  interface{}
		output interface{}
		fails  bool
	}

	parserTest := func(name string, parser argParser, testCases []testCase) {
		t.Run(name, func(t *testing.T) {
			for _, _case := range testCases {
				testCase := _case
				t.Run(fmt.Sprintf("input %#v => (%t", testCase.input, !testCase.fails), func(t *testing.T) {
					RegisterTestingT(t)
					ret, err := parser(testCase.input)
					if testCase.fails {
						Ω(err).Should(HaveOccurred())
					} else {
						Ω(err).ShouldNot(HaveOccurred())
						Ω(ret).Should(Equal(testCase.output))
					}

				})
			}
		})
	}

	parserTest("address", addressParser, []testCase{
		{
			input:  "0x111",
			output: common.HexToAddress("0x111"),
		},
		{
			input:  "1111111",
			output: common.HexToAddress("0x1111111"),
		},
		{
			input:  common.HexToAddress("0x1111111"),
			output: common.HexToAddress("0x1111111"),
		},
		{
			input: 999999,
			fails: true,
		},
	})

	parserTest("bigInt", bigIntParser, []testCase{
		{
			input:  big.NewInt(1000),
			output: big.NewInt(1000),
		},
		{
			input:  1000,
			output: big.NewInt(1000),
		},
		{
			input:  float64(1000),
			output: big.NewInt(1000),
		},
		{
			input:  uint64(1000),
			output: big.NewInt(1000),
		},
		{
			input:  int64(1000),
			output: big.NewInt(1000),
		},
		{
			input:  "1000",
			output: big.NewInt(1000),
		},
	})

	parserTest("bytes", bytesParser, []testCase{
		{
			input:  []byte{1, 2, 3},
			output: []byte{1, 2, 3},
		},
		{
			input:  []uint8{1, 2, 3},
			output: []byte{1, 2, 3},
		},
		{
			input:  "0x010203",
			output: []byte{1, 2, 3},
		},
		{
			input:  "010203",
			output: []byte{1, 2, 3},
		},
	})
}
