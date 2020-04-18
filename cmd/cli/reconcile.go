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
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// reconcileCmd represents the getblock command
var reconcileCmd = &cobra.Command{
	Use:   "reconcile",
	Short: "Fetch & Reconcile a block",
	Run:   runReconciler,
}

var blockStr string
var fromBlockNum int64
var toBlockNum int64

func init() {
	reconcileCmd.Flags().StringVar(&blockStr, "block", "", "block to reconcile")
	reconcileCmd.Flags().Int64Var(&fromBlockNum, "from", -1, "from block to reconcile")
	reconcileCmd.Flags().Int64Var(&toBlockNum, "to", -1, "to block to reconcile")
}

func runReconciler(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	fetcher, network, _ := getFetcher()

	var blocks []*types.Block
	if len(blockStr) > 0 {
		block, err := fetcher.BlockRetry(ctx, network, toBlockIndentifier(blockStr))
		utils.ExitOnError(err)
		blocks = []*types.Block{block}
	} else {
		blocks = make([]*types.Block, toBlockNum-fromBlockNum+1)
		for i := fromBlockNum; i <= toBlockNum; i++ {
			curr := &types.PartialBlockIdentifier{Index: &i}
			currBlock, err := fetcher.BlockRetry(ctx, network, curr)
			utils.ExitOnError(err)
			blocks[i-fromBlockNum] = currBlock
		}

	}

	getBalance := func(acc *types.AccountIdentifier, block *types.BlockIdentifier) (*big.Int, error) {
		_, amounts, _, err := fetcher.AccountBalance(ctx, network, acc, types.ConstructPartialBlockIdentifier(block))
		if err != nil {
			return nil, err
		}
		if len(amounts) != 1 {
			return nil, fmt.Errorf("Invalid number of amounts %d", len(amounts))
		}
		val, ok := new(big.Int).SetString(amounts[0].Value, 10)
		if !ok {
			return nil, fmt.Errorf("Invalid amounts format %s", amounts[0].Value)
		}
		return val, nil
	}

	checkDifferences := func(changes *AccountBalanceSet, from, to *types.BlockIdentifier) {
		accountsChanged := changes.Accounts()
		if len(accountsChanged) == 0 {
			log.Debug("No balance changes, skipping..")
			return
		}

		w := tabwriter.NewWriter(os.Stdout, 20, 5, 3, ' ', tabwriter.TabIndent)

		for _, acc := range accountsChanged {
			before, err := getBalance(acc, from)
			utils.ExitOnError(err)

			after, err := getBalance(acc, to)
			utils.ExitOnError(err)

			diff := new(big.Int).Sub(after, before)

			ok := diff.Cmp(changes.Get(acc)) == 0
			if !ok {
				fmt.Fprintf(w, "%v\t%s\t%s\n", acc, diff, changes.Get(acc))
			}
		}
		w.Flush()
	}

	rangeChanges := NewAccountBalanceSet()
	for _, block := range blocks {
		log.Info("Processing block", "block", block.BlockIdentifier.Index)

		blockChanges := NewAccountBalanceSet()

		if len(block.Transactions) == 0 {
			log.Debug("No transsactions, skipping..")
			continue
		}

		for _, tx := range block.Transactions {
			for _, op := range tx.Operations {
				val, _ := new(big.Int).SetString(op.Amount.Value, 10)
				blockChanges.Add(op.Account, val)
				rangeChanges.Add(op.Account, val)
			}
		}

		checkDifferences(blockChanges, block.ParentBlockIdentifier, block.BlockIdentifier)
	}

	checkDifferences(rangeChanges, blocks[0].ParentBlockIdentifier, blocks[len(blocks)-1].BlockIdentifier)

}

type comparableAccountIdentifier struct {
	address    string
	subAccount *string
	group      *string
}

func (cai *comparableAccountIdentifier) ToAccount() *types.AccountIdentifier {
	var acc types.AccountIdentifier

	acc.Address = cai.address
	if cai.subAccount != nil {
		acc.SubAccount = &types.SubAccountIdentifier{
			Address: *cai.subAccount,
		}

		if cai.group != nil {
			acc.SubAccount.Metadata = &map[string]interface{}{
				"group": *cai.group,
			}
		}
	}
	return &acc
}

func (cai *comparableAccountIdentifier) fromAccount(acc *types.AccountIdentifier) {
	cai.address = acc.Address
	if acc.SubAccount != nil {
		cai.subAccount = &acc.SubAccount.Address

		if acc.SubAccount.Metadata != nil {
			val, ok := (*acc.SubAccount.Metadata)["group"]
			if ok {
				valStr := val.(string)
				cai.group = &valStr
			}
		}
	}
}

type AccountBalanceSet struct {
	balances map[comparableAccountIdentifier]*big.Int
}

func NewAccountBalanceSet() *AccountBalanceSet {
	return &AccountBalanceSet{
		balances: make(map[comparableAccountIdentifier]*big.Int),
	}
}

func (bs *AccountBalanceSet) Get(acc *types.AccountIdentifier) *big.Int {
	var cai comparableAccountIdentifier
	cai.fromAccount(acc)
	return bs.balances[cai]
}

func (bs *AccountBalanceSet) Add(acc *types.AccountIdentifier, value *big.Int) *AccountBalanceSet {
	var cai comparableAccountIdentifier
	cai.fromAccount(acc)

	if oldValue, ok := bs.balances[cai]; ok {
		bs.balances[cai] = new(big.Int).Add(oldValue, value)
	} else {
		bs.balances[cai] = value
	}
	return bs
}

func (bs *AccountBalanceSet) Accounts() []*types.AccountIdentifier {
	accounts := make([]*types.AccountIdentifier, 0, len(bs.balances))
	for acc, _ := range bs.balances {
		accounts = append(accounts, acc.ToAccount())
	}
	return accounts
}
