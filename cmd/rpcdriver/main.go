package main

import (
	"context"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

var (
	RegistrySmartContractAddress = common.HexToAddress("0x000000000000000000000000000000000000ce10")
)

func main() {
	// _, err := rpc.Dial("https://alfajores-forno.celo-testnet.org/")
	// rpc, err := rpc.Dial("https://rc0-forno.celo-testnet.org/")
	// if err != nil {
	// 	log.Fatalf("Can't connect, %s", err)
	// }

	// debugClient := debug.NewClient(rpc)

	// var result []interface{}
	// err = debugClient.TraceTransaction(context.Background(), &result, common.HexToHash("0x82864c43454f7b3dc993315cbbe5f18807c2998f4a58c0561448f38929e4f263"), tracerConfig)
	// if err != nil {
	// 	log.Fatalf("Error calling debug_traceTransaction %s", err)
	// }

	// tx, pending, err := client.TransactionByHash(ctx, common.HexToHash("0x7c5b0a06c8f138a8d48ba2ebb34304ef17c4cec6294cddd42d7b1753f0e335c4"))
	// if err != nil {
	// 	log.Fatalf("Error with TransactionByHash, %s\n", err)
	// }

	// fmt.Printf("tx: %v %s\n", tx, pending)

	// isListening, err := client.NetworkListening(ctx)
	// if err != nil {
	// 	log.Fatalf("Error with IsListening, %s", err)
	// }
	// fmt.Printf("IsListening: %s\n", isListening)

	// common.HexToAddress("0x21e6fc92f93c8a1bb41e2be64b4e1f88a54d3576")

	// stableTokenAddress, err := registry.GetAddressForString(nil, "StableToken")
	// if err != nil {
	// 	log.Fatalf("Error %s", err)
	// }

	// fmt.Println(stableTokenAddress.Hex())
	// _, err = client.HeaderByNumber(context.Background(), nil)
	// if err != nil {
	// 	log.Fatalf("Error with HeaderByNumber, %s", err)
	// }
	// fmt.Printf("Header %v", header)
}

func printGoldLocked(lockedGold *contract.LockedGold) {
	var end uint64 = 500
	filterOpts := &bind.FilterOpts{
		Start:   0,
		End:     &end,
		Context: context.Background(),
	}

	iter, err := lockedGold.FilterGoldLocked(filterOpts, nil)

	w := tabwriter.NewWriter(os.Stdout, 20, 0, 7, ' ', 0)
	fmt.Fprintf(w, "Address\tLockedGold\tTxHash\n")
	for iter.Next() {
		fmt.Fprintf(w, "%s\t%s\t%s\n", iter.Event.Account.Hex(), iter.Event.Value.String(), iter.Event.Raw.TxHash.Hex())

	}
	w.Flush()
	if err = iter.Error(); err != nil {
		fmt.Printf("Found error, err=%s", err)
	}
	if err = iter.Close(); err != nil {
		fmt.Printf("Can't close, err=%s", err)
	}
}

func printRegistryUpdated(registry *contract.Registry) {
	var end uint64 = 500
	filterOpts := &bind.FilterOpts{
		Start:   0,
		End:     &end,
		Context: context.Background(),
	}

	iter, err := registry.RegistryFilterer.FilterRegistryUpdated(filterOpts, nil)

	w := tabwriter.NewWriter(os.Stdout, 20, 0, 7, ' ', 0)
	_, err = fmt.Fprintf(w, "Contract\tAddress\tHash\n")
	for iter.Next() {
		fmt.Fprintf(w, "%s\t%s\t%s\n", iter.Event.Identifier, iter.Event.Addr.Hex(), common.Hash(iter.Event.IdentifierHash).Hex())
	}
	w.Flush()
	if err = iter.Error(); err != nil {
		fmt.Printf("Found error, err=%s", err)
	}
	if err = iter.Close(); err != nil {
		fmt.Printf("Can't close, err=%s", err)
	}
}
