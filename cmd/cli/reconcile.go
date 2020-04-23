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
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"

	"github.com/celo-org/rosetta/cmd/internal/utils"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// reconcileCmd represents the getblock command
var reconcileCmd = &cobra.Command{
	Use:   "reconcile",
	Short: "Fetch & Reconcile a block",
	Run:   runReconciler,
}

var blockNum int64
var fromBlockNum int64
var toBlockNum int64
var batchSize int64

func init() {
	reconcileCmd.Flags().Int64Var(&blockNum, "block", -1, "block to reconcile")
	reconcileCmd.Flags().Int64Var(&fromBlockNum, "from", -1, "from block to reconcile")
	reconcileCmd.Flags().Int64Var(&toBlockNum, "to", -1, "to block to reconcile")
	reconcileCmd.Flags().Int64Var(&batchSize, "batchSize", 5000, "to block to reconcile")
}

func runReconciler(cmd *cobra.Command, args []string) {
	logger := log.New()
	ctx := context.Background()

	fetcher, network, _ := getFetcher()

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

	checkDifferences := func(id string, changes *AccountBalanceSet, from, to *types.BlockIdentifier) {
		logger := log.New("id", id)

		accountsChanged := changes.Accounts()
		if len(accountsChanged) == 0 {
			logger.Debug("No balance changes, skipping..")
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
				logger.Error("Balance Difference", "acc", fmt.Sprintf("%v", acc), "realchange", diff, "computed", changes.Get(acc))
			}
		}
		w.Flush()
	}

	if blockNum >= 0 {
		fromBlockNum = blockNum
		toBlockNum = blockNum
	}

	for curr := fromBlockNum; curr <= toBlockNum; {
		var blocks []*types.Block

		to := curr + batchSize - 1
		if to > toBlockNum {
			to = toBlockNum
		}

		logger.Info("Fetching block range (might take a while)", "from", curr, "to", to)
		blockMap, err := fetcher.BlockRange(ctx, network, curr, to)
		utils.ExitOnError(err)

		blocks = make([]*types.Block, to-curr+1)
		for num, block := range blockMap {
			blocks[num-curr] = block
		}

		reconcileRange(blocks, checkDifferences)
		if to == toBlockNum {
			break
		}
		curr = to
	}

}

func reconcileRange(blocks []*types.Block, checkDifferences func(id string, changes *AccountBalanceSet, from, to *types.BlockIdentifier)) {
	logger := log.New()

	rangeChanges := NewAccountBalanceSet()
	for _, block := range blocks {
		if block == nil {
			panic("NIL BLOCK")
		}
		blockLogger := logger.New("block", block.BlockIdentifier.Index)

		blockChanges := NewAccountBalanceSet()

		if len(block.Transactions) == 0 {
			blockLogger.Debug("No transactions, skipping..")
			continue
		}

		for _, tx := range block.Transactions {
			for _, op := range tx.Operations {
				val, _ := new(big.Int).SetString(op.Amount.Value, 10)
				blockChanges.Add(op.Account, val)
				rangeChanges.Add(op.Account, val)
			}
		}

		checkDifferences(fmt.Sprintf("block %d", block.BlockIdentifier.Index), blockChanges, block.ParentBlockIdentifier, block.BlockIdentifier)
	}

	logger.Info("Range differences")
	checkDifferences("range", rangeChanges, blocks[0].ParentBlockIdentifier, blocks[len(blocks)-1].BlockIdentifier)

}

func HashAccount(acc *types.AccountIdentifier) common.Hash {
	h := sha256.New()
	h.Write([]byte(acc.Address))
	if acc.SubAccount != nil {
		h.Write([]byte(acc.SubAccount.Address))

		if val, ok := acc.SubAccount.Metadata["group"]; ok {
			valStr := val.(string)
			h.Write([]byte("group:"))
			h.Write([]byte(valStr))
		}

	}
	var hash []byte
	hash = h.Sum(hash)
	return common.BytesToHash(hash)
}

type AccountBalanceSet struct {
	accounts map[common.Hash]*types.AccountIdentifier
	balances map[common.Hash]*big.Int
}

func NewAccountBalanceSet() *AccountBalanceSet {
	return &AccountBalanceSet{
		accounts: make(map[common.Hash]*types.AccountIdentifier),
		balances: make(map[common.Hash]*big.Int),
	}
}

func (bs *AccountBalanceSet) Get(acc *types.AccountIdentifier) *big.Int {
	return bs.balances[HashAccount(acc)]
}

func (bs *AccountBalanceSet) Add(acc *types.AccountIdentifier, value *big.Int) *AccountBalanceSet {
	hash := HashAccount(acc)

	if _, ok := bs.accounts[hash]; ok {
		bs.balances[hash] = new(big.Int).Add(bs.balances[hash], value)
	} else {
		bs.accounts[hash] = acc
		bs.balances[hash] = value
	}
	return bs
}

func (bs *AccountBalanceSet) Accounts() []*types.AccountIdentifier {
	accounts := make([]*types.AccountIdentifier, 0, len(bs.accounts))
	for _, acc := range bs.accounts {
		accounts = append(accounts, acc)
	}
	return accounts
}
