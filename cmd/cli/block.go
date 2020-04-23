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

	"github.com/celo-org/rosetta/cmd/internal/utils"
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
	db := getDb()
	cc := getCeloClient()

	blockNumber := big.NewInt(rosettabBlock.BlockIdentifier.Index)

	gpm, err := db.GasPriceMinimumFor(ctx, blockNumber)
	utils.ExitOnError(err)

	names := []string{"Governance", "GasPriceMinimum", "LockedGold", "Election", "StableToken", "Validators", "EpochRewards"}
	addresses, err := db.RegistryAddressesStartOf(ctx, blockNumber, 0, names...)

	block, err := cc.Eth.BlockByNumber(ctx, blockNumber)
	utils.ExitOnError(err)

	printTitle("Block Context")
	w := tabwriter.NewWriter(os.Stdout, 20, 5, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintf(w, "GasPriceMinimum:\t%s\n", gpm)
	fmt.Fprintf(w, "Coinbase:\t%s\n", block.Coinbase().Hex())

	for _, name := range names {
		if addr, ok := addresses[name]; ok {
			fmt.Fprintf(w, "%s:\t%s\n", name, addr.Hex())
		} else {
			fmt.Fprintf(w, "%s:\t%s\n", name, "Not Deployed")
		}
	}
	w.Flush()

	printTitle("Block Transactions")
	for i, rtx := range rosettabBlock.Transactions {
		fmt.Printf("Transaction %d: %s\n", i, rtx.TransactionIdentifier.Hash)
		txHash := common.HexToHash(rtx.TransactionIdentifier.Hash)
		tx := block.Transaction(txHash)
		receipt, err := cc.Eth.TransactionReceipt(ctx, txHash)
		utils.ExitOnError(err)

		fmt.Printf("GasPrice: %s\tGasUsed: %d\tStatus:%d\n", tx.GasPrice(), receipt.GasUsed, receipt.Status)
		fmt.Println("Operations")
		w = tabwriter.NewWriter(os.Stdout, 20, 5, 3, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "Account\tAmount\tType\tStatus\n")
		for _, op := range rtx.Operations {
			var acc string
			if op.Account.SubAccount == nil {
				acc = op.Account.Address
			} else if group, ok := op.Account.SubAccount.Metadata["group"]; ok {
				acc = fmt.Sprintf("%s - %s(%s)", op.Account.Address, op.Account.SubAccount.Address, group)
			} else {
				acc = fmt.Sprintf("%s - %s", op.Account.Address, op.Account.SubAccount.Address)
			}
			fmt.Fprintf(w, "%s\t%s\t%s\t%s\n", acc, op.Amount.Value, op.Type, op.Status)
		}
		w.Flush()
	}

	// fmt.Printf("Coinbase:\t%s\n", gpm)

}
