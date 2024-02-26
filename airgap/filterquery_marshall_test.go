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

	. "github.com/onsi/gomega"
)

func TestFilterQueryMarshalling(t *testing.T) {
	RegisterTestingT(t)

	samples := []struct {
		name   string
		sample FilterQueryParams
	}{
		{
			name: "Complete",
			sample: FilterQueryParams{
				Event:     EpochRewardsDistributedToVoters,
				Topics:    [][]interface{}{{"0x00000011"}}, // these will be simple types always
				FromBlock: big.NewInt(1000),
				ToBlock:   big.NewInt(2000),
			},
		},
		{
			name: "Nil Values",
			sample: FilterQueryParams{
				Event: EpochRewardsDistributedToVoters,
			},
		},
	}

	for _, sample := range samples {
		input := sample.sample

		t.Run(sample.name, func(t *testing.T) {
			t.Run("Json", func(t *testing.T) {
				RegisterTestingT(t)

				var output FilterQueryParams
				err := simulateWire(input, &output)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(output).Should(Equal(input))
			})

			t.Run("JsonWithPointer", func(t *testing.T) {
				RegisterTestingT(t)

				var output FilterQueryParams
				err := simulateWire(&input, &output)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(output).Should(Equal(input))
			})

			t.Run("Map", func(t *testing.T) {
				RegisterTestingT(t)

				var output FilterQueryParams

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

				var output FilterQueryParams
				err = UnmarshallFromMap(overTheWireMap, &output)
				Ω(err).ShouldNot(HaveOccurred())

				Ω(output).Should(Equal(input))

			})

		})
	}
}
