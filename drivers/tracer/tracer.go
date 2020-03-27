package main

import (
	"context"
	"os"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/k0kubun/pp"
)

var ctx = context.Background()

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	DriveTracer()
}

func CeloClient() *client.CeloClient {
	celo, err := client.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	return celo
}

func DriveTracer() {
	cc := CeloClient()
	//fromAddress := "0x5409ed021d9299bf6814279a6a1411a7e866a631"
	//toAddress := "0xbBae99F0E1EE565404465638d40827b54D343638"

	// TestContract.selfDestruct
	txHash := common.HexToHash("0x3506ee7d435786e3630c0fcee80d08015c51221675ca72c349fc23f64c56cf20")
	transfers, err := cc.Debug.TransactionTransfers(ctx, txHash)
	if err != nil {
		panic(err)
	}

	pp.Print(transfers)
}
