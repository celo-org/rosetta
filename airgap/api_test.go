package airgap

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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
		sig := valuesToSignature(tx.ChainId, v, r, s)
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
