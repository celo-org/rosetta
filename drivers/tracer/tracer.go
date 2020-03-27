package main

import (
	"context"
	"fmt"
	"os"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	//"github.com/k0kubun/pp"
)

var ctx = context.Background()

// Started by https://github.com/celo-org/celo-monorepo/blob/jfoutts/test-custody/packages/celotool/src/e2e-tests/tracer_tests.ts#L728
func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	DriveTracer()
	fmt.Println("success")
}

func CeloClient() *client.CeloClient {
	celo, err := client.Dial("http://localhost:8545")
	if err != nil {
		panic(err)
	}
	return celo
}

func assertIntEqual(x int, y int) {
	if x != y {
		panic(fmt.Sprintf("%d != %d", x, y))
	}
}

func assertStringEqual(x string, y string) {
	if x != y {
		panic(x + " != " + y)
	}
}

func assertAddressEqual(x common.Address, y common.Address) {
	if x != y {
		panic(x.Hex() + " != " + y.Hex())
	}
}

func DriveTracer() {
	cc := CeloClient()
	fromAddress := common.HexToAddress("0x5409ed021d9299bf6814279a6a1411a7e866a631")
	toAddress := common.HexToAddress("0xbBae99F0E1EE565404465638d40827b54D343638")

	// https://github.com/celo-org/celo-monorepo/blob/jfoutts/test-custody/packages/celotool/src/e2e-tests/tracer_tests.ts#L25
	testContractAddress := common.HexToAddress("0x07f96aa816c1f244cbc6ef114bb2b023ba54a2eb")

	// TestContract.selfDestruct
	txHash := common.HexToHash("0x3506ee7d435786e3630c0fcee80d08015c51221675ca72c349fc23f64c56cf20")
	transfers, err := cc.Debug.TransactionTransfers(ctx, txHash)
	if err != nil {
		panic(err)
	}
	assertIntEqual(len(transfers), 2)
	assertAddressEqual(transfers[0].From, fromAddress)
	assertAddressEqual(transfers[0].To, testContractAddress)
	assertStringEqual(transfers[0].Status, "success")
	assertAddressEqual(transfers[1].From, testContractAddress)
	assertAddressEqual(transfers[1].To, toAddress)
	assertStringEqual(transfers[1].Status, "success")
}
