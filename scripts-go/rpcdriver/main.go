package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
		log.Fatalf("Error with TransactionByHash, %s", err)
	}

	fmt.Printf("tx: %v %s", tx, pending)

	isListening, err := client.NetworkListening(ctx)
	if err != nil {
		log.Fatalf("Error with IsListening, %s", err)
	}
	fmt.Printf("IsListening: %s", isListening)
	// _, err = client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("Error with HeaderByNumber, %s", err)
	// }
	// fmt.Printf("Header %v", header)
}
