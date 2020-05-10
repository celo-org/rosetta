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
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func simulateWire(input interface{}, output interface{}) error {
	data, err := json.MarshalIndent(input, "", " ")
	if err != nil {
		return err
	}

	return json.Unmarshal(data, output)
}

func TestTxArgsMarshalling(t *testing.T) {
	RegisterTestingT(t)

	address1 := common.HexToAddress("0x11111")
	address2 := common.HexToAddress("0x22222")

	samples := []struct {
		name   string
		sample TxArgs
	}{
		{
			name: "Complete",
			sample: TxArgs{
				From:   address1,
				To:     &address2,
				Value:  big.NewInt(5000),
				Method: LockGold,
				Args:   []interface{}{"0x00000011"}, // these will be simple types always
			},
		},
		{
			name: "Nil Values",
			sample: TxArgs{
				From: address1,
			},
		},
	}

	for _, sample := range samples {
		input := sample.sample

		t.Run(sample.name, func(t *testing.T) {
			t.Run("Json", func(t *testing.T) {
				RegisterTestingT(t)

				var output TxArgs
				err := simulateWire(input, &output)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(output).Should(Equal(input))
			})

			t.Run("JsonWithPointer", func(t *testing.T) {
				RegisterTestingT(t)

				var output TxArgs
				err := simulateWire(&input, &output)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(output).Should(Equal(input))
			})

			t.Run("Map", func(t *testing.T) {
				RegisterTestingT(t)

				var output TxArgs

				theMap, err := MarshallToMap(input)
				Ω(err).ShouldNot(HaveOccurred())

				err = UnmarshallFromMap(theMap, &output)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(output).Should(Equal(input))
			})

			t.Run("MapThenJson", func(t *testing.T) {
				RegisterTestingT(t)

				theMap, err := MarshallToMap(input)
				Ω(err).ShouldNot(HaveOccurred())

				var overTheWireMap map[string]interface{}
				err = simulateWire(theMap, &overTheWireMap)
				Ω(err).ShouldNot(HaveOccurred())

				var output TxArgs
				err = UnmarshallFromMap(overTheWireMap, &output)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(output).Should(Equal(input))

			})

		})
	}
}
