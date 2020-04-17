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
	"log"

	"github.com/celo-org/rosetta/client"
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

	_ = addr

	// step 2: query rosetta for network/api information ONLINE
	// ctx := context.Background()
	// fetcherInstance := fetcher.New(ctx, "http://localhost:8080", "transferAgent", http.DefaultClient, fetcher.DefaultBlockConcurrency, fetcher.DefaultTransactionConcurrency)
	// networkStatusResp, err := fetcherInstance.InitializeAsserter(ctx)
	// if err != nil {
	// 	log.Fatalf("Error initializing ")
	// }
	// networkIdentifier := client.PartialToFullNetworkIdentifier(networkStatusResp.NetworkStatus.NetworkIdentifier)

	// // step 2: query rosetta for transfer transaction metadata ONLINE
	// transferMethod := "transfer" // TODO: source from consts
	// fee, meta, err := fetcherInstance.ConstructionMetadata(ctx, networkIdentifier, accountIdentifier, &transferMethod)

	// transferMetadata, ok := (*meta)[transferMethod].(*api.TransferMetadata)
	// if !ok {
	// 	log.Fatalf("Failed to cast received metadata[transfer] to TransferMetadata")
	// }

	// // step 3: construct and sign transfer transaction OFFLINE
	// transferTx := client.ConstructTransferTransaction(addr, &api.DummyAddress, transferMetadata.Balance, transferMetadata)
	// // TODO: apply suggestedFee to transferTx
	// _ = fee

	// chainIDInt, err := strconv.ParseInt(networkIdentifier.Network, 10, 64)
	// if err != nil {
	// 	log.Fatalf("Failed to cast network to chainId %s", err)
	// }

	// signer := client.NewSigner(big.NewInt(chainIDInt))
	// signedTransferTx, err := client.SignTransaction(transferTx, privKey, &signer)
	// if err != nil {
	// 	log.Fatalf("Failed to sign transfer transaction %s", err)
	// }

	// // step 4: submit signed transfer transaction ONLINE
	// rawSignedTx, err := client.EncodeTransaction(signedTransferTx)
	// if err != nil {
	// 	log.Fatalf("Failed to encode transaction %s", err)
	// }
	// txID, submitStatus, _, err := fetcherInstance.ConstructionSubmit(ctx, rawSignedTx)
	// if err != nil {
	// 	log.Fatalf("Failed to submit transaction to node %s", err)
	// }

	// log.Printf("Successfully subbmited transfer transaction %s with status %s", txID.Hash, *submitStatus)
}
