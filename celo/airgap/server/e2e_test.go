package server

import (
	"context"
	"encoding/json"
	"math/big"
	"testing"

	"github.com/celo-org/rosetta/celo/airgap"
	"github.com/celo-org/rosetta/celo/airgap/client"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func simulateWire(txArgs *airgap.TxArgs) (*airgap.TxArgs, error) {
	data, err := json.Marshal(txArgs)
	if err != nil {
		return nil, err
	}

	var parsed airgap.TxArgs
	err = json.Unmarshal(data, &parsed)
	if err != nil {
		return nil, err
	}

	if parsed.Method != nil {
		parsed.Method, err = airgap.MethodFromString(parsed.Method.String())
		if err != nil {
			return nil, err
		}
	}

	return &parsed, nil
}

func TestClientServer(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	client := client.NewAirgapClient()
	server, err := NewAirgapServer(&serverContextStub{})
	Ω(err).ShouldNot(HaveOccurred())

	t.Run("Vote()", func(t *testing.T) {
		RegisterTestingT(t)

		txArgs, err := client.Vote(common.HexToAddress("A"), common.HexToAddress("C"), big.NewInt(4))
		Ω(err).ShouldNot(HaveOccurred())

		txArgs, err = simulateWire(txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		txMetadata, err := server.ObtainMetadata(ctx, txArgs)
		Ω(err).ShouldNot(HaveOccurred())

		tx, err := client.TxFromMetadata(txMetadata)
		Ω(err).ShouldNot(HaveOccurred())

		_ = tx

	})

}

// func ClientServerSerialization(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	cc, err := celoClient.Dial("http://localhost:8545")
// 	if err != nil {
// 		return
// 	}

// 	ctx := context.Background()
// 	chainId, err := cc.Net.ChainId(ctx)

// 	privKey, _ := client.Keygen()
// 	_, addr, _ := client.Derive(privKey)
// 	signer := client.NewSigner(chainId)

// 	server, err := NewAirgapServer(cc)
// 	g.Expect(err).ToNot(HaveOccurred())

// 	voteTxArgs := &TxArgs{
// 		From:   common.ZeroAddress,
// 		Method: Vote,
// 		Args:   []interface{}{addr, big.NewInt(100)},
// 	}

// 	voteTxMetadata, err := server.ObtainMetadata(ctx, voteTxArgs)
// 	g.Expect(err).ToNot(HaveOccurred())

// 	tx := voteTxMetadata.AsTransaction()
// 	signedTx, err := client.SignTransaction(tx, privKey, signer)
// 	serializedTx, err := client.EncodeTransaction(signedTx)
// 	//hey follow me
// 	// okay
// 	parsedTx, err := client.DecodeTransaction(serializedTx)

// 	status, err := server.SubmitTx(ctx, signedTx)
// 	g.Expect(err).ToNot(HaveOccurred())
// 	g.Expect(status).To(Equal(true))

// }
