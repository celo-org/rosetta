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
	"math/big"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	. "github.com/onsi/gomega"
)

func TestTxMetadataMarshalling(t *testing.T) {
	RegisterTestingT(t)

	address1 := common.HexToAddress("0x11111")
	address2 := common.HexToAddress("0x22222")

	samples := []struct {
		name   string
		sample TxMetadata
	}{
		{
			name: "Complete",
			sample: TxMetadata{
				From:                common.HexToAddress("0x55555"),
				Nonce:               70,
				GasPrice:            big.NewInt(5000),
				GatewayFeeRecipient: &address1,
				GatewayFee:          big.NewInt(5000),
				FeeCurrency:         &address2,
				To:                  common.HexToAddress("0x33333"),
				Data:                []byte{1, 2, 3},
				Value:               big.NewInt(3000),
				Gas:                 98,
				ChainId:             big.NewInt(2000),
			},
		},
		{
			name: "Nil Values",
			sample: TxMetadata{
				From:     common.HexToAddress("0x55555"),
				Nonce:    70,
				GasPrice: big.NewInt(5000),
				To:       common.HexToAddress("0x33333"),
				Data:     []byte{},
				Value:    big.NewInt(3000),
				Gas:      98,
				ChainId:  big.NewInt(2000),
			},
		},
	}

	for _, sample := range samples {
		input := sample.sample

		t.Run(sample.name, func(t *testing.T) {
			t.Run("Json", func(t *testing.T) {
				RegisterTestingT(t)

				var output TxMetadata
				err := simulateWire(input, &output)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(output).Should(Equal(input))
			})

			t.Run("JsonWithPointer", func(t *testing.T) {
				RegisterTestingT(t)

				var output TxMetadata
				err := simulateWire(&input, &output)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(output).Should(Equal(input))
			})

			t.Run("Map", func(t *testing.T) {
				RegisterTestingT(t)

				var output TxMetadata

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

				var output TxMetadata
				err = UnmarshallFromMap(overTheWireMap, &output)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(output).Should(Equal(input))

			})

		})
	}
}
