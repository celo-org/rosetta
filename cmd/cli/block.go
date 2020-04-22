/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/cmd/internal/utils"
	"github.com/celo-org/rosetta/db"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

// blockCmd represents the getblock command
var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Fetchs a block",
	Long:  `Fettchs the given block`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.Background()

		fetcher, network, _ := getFetcher()

		blockIdentifier := toBlockIdentifier(args[0])
		block, err := fetcher.BlockRetry(ctx, network, blockIdentifier)
		utils.ExitOnError(err)

		// printBlockContext(block)

		printBlockContext(block)
	},
}

func printTitle(title string) {
	fmt.Printf("\n%s\n-----------------------\n", title)
}

func printBlockContext(rosettabBlock *types.Block) {
	ctx := context.Background()
	celoStore, err := db.NewSqliteDb("./envs/alfajores/rosetta.db")
	utils.ExitOnError(err)

	cc, err := client.Dial("http://localhost:8545")
	utils.ExitOnError(err)

	blockNumber := big.NewInt(rosettabBlock.BlockIdentifier.Index)

	gpm, err := celoStore.GasPriceMinimumFor(ctx, blockNumber)
	utils.ExitOnError(err)

	addresses, err := celoStore.RegistryAddressesStartOf(ctx, blockNumber, 0, "Governance", "GasPriceMinimum", "LockedGold", "Election", "EpochRewards")

	block, err := cc.Eth.BlockByNumber(ctx, blockNumber)
	utils.ExitOnError(err)

	printTitle("Block Context")
	w := tabwriter.NewWriter(os.Stdout, 20, 5, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "GasPriceMinimum:\t%s\n", gpm)
	fmt.Fprintf(w, "Coinbase:\t%s\n", block.Coinbase().Hex())
	if addr, ok := addresses["Governance"]; ok {
		fmt.Fprintf(w, "Governance:\t%s\n", addr.Hex())
	} else {
		fmt.Fprintf(w, "Governance:\t%s\n", "Not Deployed")
	}
	if addr, ok := addresses["GasPriceMinimum"]; ok {
		fmt.Fprintf(w, "GasPriceMinimum:\t%s\n", addr.Hex())
	} else {
		fmt.Fprintf(w, "GasPriceMinimum:\t%s\n", "Not Deployed")
	}
	if addr, ok := addresses["LockedGold"]; ok {
		fmt.Fprintf(w, "LockedGold:\t%s\n", addr.Hex())
	} else {
		fmt.Fprintf(w, "LockedGold:\t%s\n", "Not Deployed")
	}
	if addr, ok := addresses["Election"]; ok {
		fmt.Fprintf(w, "Election:\t%s\n", addr.Hex())
	} else {
		fmt.Fprintf(w, "Election:\t%s\n", "Not Deployed")
	}
	w.Flush()

	printTitle("Block Transactions")
	for i, rtx := range rosettabBlock.Transactions {
		fmt.Printf("Transaction %d: %s\n", i, rtx.TransactionIdentifier.Hash)
		if rtx.TransactionIdentifier.Hash != rosettabBlock.BlockIdentifier.Hash {
			txHash := common.HexToHash(rtx.TransactionIdentifier.Hash)
			tx := block.Transaction(txHash)
			receipt, err := cc.Eth.TransactionReceipt(ctx, txHash)
			utils.ExitOnError(err)

			fmt.Printf("GasPrice: %s\tGasUsed: %d\tStatus:%d\n", tx.GasPrice(), receipt.GasUsed, receipt.Status)
		}
		fmt.Println("Operations")
		for _, op := range rtx.Operations {
			fmt.Printf("\tAddr: %s  SubAccount: %v\t  Amount: %s\tType: %s\n", op.Account.Address, op.Account.SubAccount, op.Amount.Value, op.Type)
		}
	}

	// fmt.Printf("Coinbase:\t%s\n", gpm)
}
