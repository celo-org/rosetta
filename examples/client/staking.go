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

package client

import (
	"context"
	"log"
	"math/big"
	"strconv"

	"github.com/celo-org/rosetta/celo/airgap"
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
	fetcherInstance := fetcher.New("http://localhost:8080")
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
	submitSigned := func(txOptions *airgap.TxArgs) error {
		// step 1: decide options OFFLINE
		optionsMap := map[string]interface{}{
			"from":   txOptions.From,
			"to":     txOptions.To,
			"value":  txOptions.Value,
			"method": txOptions.Method,
			"args":   txOptions.Args,
		}
		// optionsMap := client.SerializeTransactionOptions(txOptions)

		// step 2: fetch metadata ONLINE
		metadataMap, err := fetcherInstance.ConstructionMetadata(ctx, networkId, optionsMap)
		if err != nil {
			return err
		}

		// step 3: sign transaction OFFLINE
		txMetadata := metadataMap["tx"].(airgap.TxMetadata)
		tx := txMetadata.AsTransaction()
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

	submitSigned(&airgap.TxArgs{
		From:   *addr,
		Method: airgap.CreateAccount,
	})

	submitSigned(&airgap.TxArgs{
		From:   *addr,
		Method: airgap.LockGold,
		Value:  big.NewInt(100),
	})

	submitSigned(&airgap.TxArgs{
		From:   *addr,
		Method: airgap.UnlockGold,
		Value:  big.NewInt(50),
	})

	submitSigned(&airgap.TxArgs{
		From:   *addr,
		Method: airgap.WithdrawGold,
		Args:   []interface{}{big.NewInt(0)}, // withdrawal index 0
	})
}
