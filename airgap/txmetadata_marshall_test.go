package airgap

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
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
				From:       common.HexToAddress("0x55555"),
				Nonce:      70,
				GasPrice:   big.NewInt(5000),
				GatewayFee: big.NewInt(5000),
				To:         common.HexToAddress("0x33333"),
				Data:       []byte{},
				Value:      big.NewInt(3000),
				Gas:        98,
				ChainId:    big.NewInt(2000),
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
