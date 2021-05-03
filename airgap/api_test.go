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
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/crypto"
	. "github.com/onsi/gomega"
)

func TestTransaction(t *testing.T) {
	RegisterTestingT(t)

	privateKeyHex := "0c06d31f68640188d8baacdf2805aadf4d887b95a9f6e330b3cfcf6063ac010d"

	client := NewClient()

	privKey, err := crypto.ToECDSA(common.Hex2Bytes(privateKeyHex))
	Ω(err).ShouldNot(HaveOccurred())

	_, fromAddr, err := client.Derive(privKey)
	Ω(err).ShouldNot(HaveOccurred())

	addr2 := common.HexToAddress("0x2222")
	addr3 := common.HexToAddress("0x3333")
	newTx := func() *Transaction {
		return &Transaction{
			TxMetadata: &TxMetadata{
				From:                *fromAddr,
				Nonce:               3,
				GasPrice:            big.NewInt(111),
				GatewayFeeRecipient: &addr2,
				GatewayFee:          big.NewInt(222),
				FeeCurrency:         &addr3,
				To:                  common.HexToAddress("0x4444"),
				Data:                []byte{1, 2, 3, 4},
				Value:               big.NewInt(4444),
				Gas:                 4,
				ChainId:             big.NewInt(42220),
			},
		}
	}

	t.Run("Signature", func(t *testing.T) {
		RegisterTestingT(t)
		tx := newTx()
		signedTx, err := client.SignTx(tx, privKey)
		Ω(err).ShouldNot(HaveOccurred())

		gethTx, _ := signedTx.AsGethTransaction()

		v, r, s := gethTx.RawSignatureValues()
		sig := ValuesToSignature(tx.ChainId, v, r, s)
		Ω(sig).Should(Equal(signedTx.Signature))

		addr, err := types.NewEIP155Signer(tx.ChainId).Sender(gethTx)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(tx.From))

	})

	t.Run("Serialize", func(t *testing.T) {
		RegisterTestingT(t)
		tx := newTx()
		signedTx, err := client.SignTx(tx, privKey)
		Ω(err).ShouldNot(HaveOccurred())

		signedTxRaw, err := signedTx.Serialize()
		Ω(err).ShouldNot(HaveOccurred())

		var desrializedTx Transaction
		err = desrializedTx.Deserialize(signedTxRaw, tx.ChainId)
		Ω(err).ShouldNot(HaveOccurred())

		Ω(desrializedTx).Should(Equal(*signedTx))
	})

}
