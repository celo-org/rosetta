package client

import (
	"context"
	"log"
	"math/big"
	"strconv"

	"github.com/celo-org/rosetta/celo/transaction"
	"github.com/celo-org/rosetta/client"
	"github.com/coinbase/rosetta-sdk-go/fetcher"
)

func main() {
	// step 1: create a new account OFFLINE
	privKey, err := client.Keygen()
	if err != nil {
		log.Fatalf("Error generating private key %s", err)
	}

	pubKey, addr, err := client.Derive(privKey)
	if err != nil {
		log.Fatalf("Error deriving pubkey & address %s", err)
	}
	_ = pubKey

	// step 2: query rosetta for network/api information ONLINE
	ctx := context.Background()
	fetcherInstance := fetcher.New(ctx, "http://localhost:8080")
	networkId, _, err := fetcherInstance.InitializeAsserter(ctx)
	if err != nil {
		log.Fatalf("Error initializing ")
	}

	chainIDInt, err := strconv.ParseInt(networkId.Network, 10, 64)
	if err != nil {
		log.Fatalf("Failed to cast network to chainId %s", err)
	}

	signer := client.NewSigner(big.NewInt(chainIDInt))

	// generalizes tx flow
	submitSigned := func(txOptions *transaction.TransactionOptions) error {
		// step 1: decide options OFFLINE
		optionsMap := client.SerializeTransactionOptions(txOptions)

		// step 2: fetch metadata ONLINE
		metadataMap, err := fetcherInstance.ConstructionMetadata(ctx, networkId, &optionsMap)
		if err != nil {
			return err
		}

		// step 3: sign transaction OFFLINE
		txMetadata := (*metadataMap)["tx"].(transaction.TransactionMetadata)
		tx := client.ConstructTxFromMetadata(&txMetadata)
		signedTx, err := client.SignTransaction(tx, privKey, &signer)
		if err != nil {
			return err
		}

		signedTxRaw, err := client.EncodeTransaction(signedTx)
		if err != nil {
			return err
		}

		// step 4: submit transaction ONLINE
		txId, _, err := fetcherInstance.ConstructionSubmit(ctx, networkId, signedTxRaw)
		if err != nil {
			return err
		}
		log.Printf("'%s' tx submitted successfully with hash '%s'", txOptions.Method, txId.Hash)
		return nil
	}

	submitSigned(&transaction.TransactionOptions{
		From:   *addr,
		Method: &transaction.CreateAccount,
	})

	submitSigned(&transaction.TransactionOptions{
		From:   *addr,
		Method: &transaction.LockGold,
		Value:  big.NewInt(100),
	})

	submitSigned(&transaction.TransactionOptions{
		From:   *addr,
		Method: &transaction.UnlockGold,
		Value:  big.NewInt(50),
	})

	submitSigned(&transaction.TransactionOptions{
		From:   *addr,
		Method: &transaction.WithdrawGold,
		Args:   []interface{}{big.NewInt(0)}, // withdrawal index 0
	})
}
