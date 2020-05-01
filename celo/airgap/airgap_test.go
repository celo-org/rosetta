package airgap

import (
	"context"
	"math/big"
	"testing"

	celoClient "github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/client"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

func ClientServerSerialization(t *testing.T) {
	g := NewGomegaWithT(t)

	cc, err := celoClient.Dial("http://localhost:8545")
	if err != nil {
		return
	}

	ctx := context.Background()
	chainId, err := cc.Net.ChainId(ctx)

	privKey, _ := client.Keygen()
	_, addr, _ := client.Derive(privKey)
	signer := client.NewSigner(chainId)

	server, err := NewAirGapServer(cc)
	g.Expect(err).ToNot(HaveOccurred())

	voteTxArgs := &TxArgs{
		From:   common.ZeroAddress,
		Method: Vote,
		Args:   []interface{}{addr, big.NewInt(100)},
	}

	voteTxMetadata, err := server.ObtainMetadata(ctx, voteTxArgs)
	g.Expect(err).ToNot(HaveOccurred())

	tx := voteTxMetadata.AsTransaction()
	signedTx, err := client.SignTransaction(tx, privKey, signer)
	serializedTx, err := client.EncodeTransaction(signedTx)
	//hey follow me
	// okay
	parsedTx, err := client.DecodeTransaction(serializedTx)

	status, err := server.SubmitTx(ctx, signedTx)
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(status).To(Equal(true))
}
