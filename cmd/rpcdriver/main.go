package main

import (
	"context"
	"fmt"
	"log"

	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	RegistrySmartContractAddress = common.HexToAddress("0x000000000000000000000000000000000000ce10")
)

func main() {
	// client, err := ethclient.Dial("ws://localhost:8546")
	client, err := ethclient.Dial("https://alfajores-forno.celo-testnet.org")
	if err != nil {
		log.Fatalf("Can't connect, %s", err)
	}

	ctx := context.Background()

	tx, pending, err := client.TransactionByHash(ctx, common.HexToHash("0x7c5b0a06c8f138a8d48ba2ebb34304ef17c4cec6294cddd42d7b1753f0e335c4"))
	if err != nil {
		log.Fatalf("Error with TransactionByHash, %s\n", err)
	}

	fmt.Printf("tx: %v %t\n", tx, pending)

	isListening, err := client.NetworkListening(ctx)
	if err != nil {
		log.Fatalf("Error with IsListening, %s", err)
	}
	fmt.Printf("IsListening: %t\n", isListening)

	common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576")

	registry, err := contract.NewRegistry(RegistrySmartContractAddress, client)
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	stableTokenAddress, err := registry.GetAddressForString(nil, "StableToken")
	if err != nil {
		log.Fatalf("Error %s", err)
	}

	fmt.Println(stableTokenAddress.Hex())
	// _, err = client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("Error with HeaderByNumber, %s", err)
	// }
	// fmt.Printf("Header %v", header)
}
