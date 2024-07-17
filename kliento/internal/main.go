package main

import (
	"log"
	"os"

	"github.com/celo-org/rosetta/kliento/client"
	"github.com/celo-org/rosetta/kliento/internal/examples"
)

func main() {
	args := os.Args[1:]
	nodeURL := args[0]
	cc, err := client.Dial(nodeURL)
	if err != nil {
		log.Fatalf("dialing node failed %e", err)
	}

	recipient := args[1]
	examples.WatchStableTokenTransfers(cc, recipient)
}
