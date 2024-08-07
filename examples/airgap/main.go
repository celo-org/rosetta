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

package main

import (
	"context"
	"log"

	// "math/big"

	"github.com/celo-org/celo-blockchain/common"
	// "github.com/celo-org/celo-blockchain/common/hexutil"
	// "github.com/celo-org/celo-blockchain/crypto"
	"github.com/celo-org/rosetta/airgap"

	// celoClient "github.com/celo-org/rosetta/kliento/client"
	// "github.com/celo-org/rosetta/kliento/registry"

	"github.com/coinbase/rosetta-sdk-go/fetcher"
)

func main() {
	client := airgap.NewClient()

	// step 1: create a new account OFFLINE
	privKey, err := client.Keygen()
	if err != nil {
		log.Fatalf("Error generating private key %s", err)
	}

	// (or use an existing key)
	// recommended for testing, since the account will need to have funds for gas fees
	// privKeyBytes, err := hexutil.Decode("0x8e14643e23a3a6aa027e6e01fad86f54e63c5186636ce4be01ea9042907b1ff5")
	// privKey, err := crypto.ToECDSA(privKeyBytes)

	pubKey, addr, err := client.Derive(privKey)
	if err != nil {
		log.Fatalf("Error deriving pubkey & address %s", err)
	}
	_ = pubKey

	// step 2: query rosetta for network/api information ONLINE
	ctx := context.Background()
	fetcherInstance := fetcher.New("http://localhost:8080")
	networkId, _, fetcherErr := fetcherInstance.InitializeAsserter(ctx, nil)
	if fetcherErr != nil {
		log.Fatalf("Error initializing ")
	}

	// generalizes tx flow
	submitSigned := func(txArgs *airgap.TxArgs) error {
		// step 1: decide options OFFLINE
		txArgsMap, err := airgap.MarshallToMap(txArgs)
		if err != nil {
			return err
		}
		// step 2: fetch metadata ONLINE
		txMetadataMap, _, fetcherErr := fetcherInstance.ConstructionMetadata(ctx, networkId, txArgsMap, nil)
		if fetcherErr != nil {
			return fetcherErr.Err
		}

		txMetadata := &airgap.TxMetadata{}
		if err := airgap.UnmarshallFromMap(txMetadataMap, &txMetadata); err != nil {
			return err
		}

		// step 3: sign transaction OFFLINE
		tx, err := client.ConstructTxFromMetadata(txMetadata)
		if err != nil {
			return err
		}

		signedTx, err := client.SignTx(tx, privKey)
		if err != nil {
			return err
		}

		signedTxRaw, err := signedTx.Serialize()
		if err != nil {
			return err
		}

		// step 4: submit transaction ONLINE
		txId, _, fetcherErr := fetcherInstance.ConstructionSubmit(ctx, networkId, common.Bytes2Hex(signedTxRaw))
		if fetcherErr != nil {
			return fetcherErr.Err
		}
		log.Printf("'%s' tx submitted successfully with hash '%s'", txArgs.Method, txId.Hash)
		return nil
	}

	argBuilder := airgap.NewArgBuilder()

	var txArgs *airgap.TxArgs

	txArgs, err = argBuilder.CreateAccount(*addr)
	if err != nil {
		log.Fatalf("Error build txArgs: %s", err)
	}
	if err := submitSigned(txArgs); err != nil {
		log.Fatalf("Error on submit Tx: %s", err)
	}

	// lockValue := big.NewInt(10000000000)

	// txArgs, err = argBuilder.LockGold(*addr, lockValue)
	// if err != nil {
	// 	log.Fatalf("Error build txArgs: %s", err)
	// }
	// if err = submitSigned(txArgs); err != nil {
	// 	log.Fatalf("Error on submit Tx: %s", err)
	// }

	// cc, err := celoClient.Dial("http://localhost:8545")
	// registry, err := registry.New(cc)
	// if err != nil {
	// 	log.Fatalf("Error creating Registry: %s", err)
	// }
	// election, err := registry.GetElectionContract(ctx, nil)
	// if err != nil {
	// 	log.Fatalf("Error getting Election Contract: %s", err)
	// }
	// groups, err := election.GetEligibleValidatorGroups(nil)
	// if err != nil {
	// 	log.Fatalf("Error getting eligible validator groups: %s", err)
	// }

	// txArgs, err = argBuilder.Vote(*addr, groups[0], lockValue)
	// if err != nil {
	// 	log.Fatalf("Error build txArgs: %s", err)
	// }
	// if err = submitSigned(txArgs); err != nil {
	// 	log.Fatalf("Error on submit Tx: %s", err)
	// }

	// votedForGroups, err := election.GetGroupsVotedForByAccount(nil, *addr)
	// log.Printf("votedForGroups %s:", votedForGroups)

	// // eligible
	// txArgs, err = argBuilder.RevokePendingVotes(*addr, groups[0], lockValue)
	// if err != nil {
	// 	log.Fatalf("Error build txArgs: %s", err)
	// }
	// if err = submitSigned(txArgs); err != nil {
	// 	log.Fatalf("Error on submit Tx: %s", err)
	// }
	// // not eligible
	// txArgs, err = argBuilder.RevokePendingVotes(*addr, *addr, lockValue)
	// if err != nil {
	// 	log.Fatalf("Error build txArgs: %s", err)
	// }
	// if err = submitSigned(txArgs); err != nil {
	// 	log.Fatalf("Error on submit Tx: %s", err)
	// }

	// txArgs, err = argBuilder.UnlockGold(*addr, big.NewInt(50))
	// if err != nil {
	// 	log.Fatalf("Error build txArgs: %s", err)
	// }
	// if err = submitSigned(txArgs); err != nil {
	// 	log.Fatalf("Error on submit Tx: %s", err)
	// }

	// txArgs, err = argBuilder.WithdrawGold(*addr, big.NewInt(0))
	// if err != nil {
	// 	log.Fatalf("Error build txArgs: %s", err)
	// }
	// if err = submitSigned(txArgs); err != nil {
	// 	log.Fatalf("Error on submit Tx: %s", err)
	// }
}
