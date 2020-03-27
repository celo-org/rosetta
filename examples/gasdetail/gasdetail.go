package main

import (
	"context"
	"fmt"
	"math/big"
	"os"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/internal/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
)

var ctx = context.Background()

func PrintUsage(msg string) {
	fmt.Println(msg)
	fmt.Println("Usage: <cmd> <blockNumber> <txhash>")
	os.Exit(1)
}
func main() {
	if len(os.Args) != 3 {
		PrintUsage("Incorrect Arguments")
	}

	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	blockNumber, ok := new(big.Int).SetString(os.Args[1], 10)
	if !ok {
		PrintUsage("Invalid Block Number")
	}

	txHash := common.HexToHash(os.Args[2])

	cc := CeloClient()

	header, err := cc.Eth.HeaderByNumber(ctx, blockNumber)
	if err != nil {
		panic(err)
	}

	tx, _, err := cc.Eth.TransactionByHash(ctx, txHash)
	if err != nil {
		panic(err)
	}

	receipt, err := cc.Eth.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		panic(err)
	}

	fromAddr, err := cc.Eth.TransactionSender(ctx, tx, header.Hash(), receipt.TransactionIndex)
	if err != nil {
		panic(err)
	}

	log.Info("TxInfo", "usedGas", receipt.GasUsed, "gasPrice", tx.GasPrice(), "gatewayFee", tx.GatewayFee(), "gatewayFeeRecipient", tx.GatewayFeeRecipient(),
		"from", fromAddr, "amount", tx.Value(), "to", tx.To())

	totalCost := new(big.Int).Mul(new(big.Int).SetUint64(receipt.GasUsed), tx.GasPrice())
	totalCost.Add(totalCost, tx.Value())
	log.Info("Total Cost", "cost", totalCost)
	txTracer := celo.NewTxTracer(
		ctx,
		cc,
		header,
		tx,
		receipt,
	)

	gasDetail, err := txTracer.GasDetail()
	if err != nil {
		panic(err)
	}

	log.Info("Got Gasdetail")
	for addr, value := range gasDetail {
		prevBalance, err := cc.Eth.BalanceAt(ctx, addr, new(big.Int).Sub(blockNumber, big.NewInt(1)))
		if err != nil {
			panic(err)
		}
		afterBalance, err := cc.Eth.BalanceAt(ctx, addr, blockNumber)
		if err != nil {
			panic(err)
		}

		diff := new(big.Int).Sub(afterBalance, prevBalance)
		if fromAddr == addr {
			diff.Add(diff, tx.Value())
		} else if *tx.To() == addr {
			diff.Add(diff, tx.Value())
		}

		log.Info("Operation", "account", addr.Hex(), "amount", value, "diff", diff, "balanceBefore", prevBalance, "balanceAfter", afterBalance)
	}
}

func CeloClient() *client.CeloClient {
	celo, err := client.Dial(config.FornoRC0Url)
	if err != nil {
		panic(err)
	}
	return celo
}
