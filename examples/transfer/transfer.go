// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"time"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/log"
	"github.com/celo-org/celo-blockchain/mycelo/hdwallet"
	"github.com/celo-org/rosetta/internal/config"
	"github.com/celo-org/rosetta/kliento/client"
)

var beneficiary = common.HexToAddress("0x671d520ae3e89ea5383a5d7162bced79fd25cdee")
var ctx = context.Background()

func MustGetAccount(wallet *hdwallet.Wallet, index int) *accounts.Account {
	path := hdwallet.MustParseDerivationPath("m/44'/52752'/0'/0/" + strconv.Itoa(index))
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Crit("Error deriving path", "err", err)
		return nil

	}
	return &account
}

func MustNewWallet(mnemonic string) *hdwallet.Wallet {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Crit("Error creating wallet", "err", err)
		return nil
	}
	return wallet
}

func GoldTransferTx(cc *client.CeloClient, from common.Address, to common.Address, value *big.Int) *types.Transaction {
	gasPrice, err := cc.Eth.SuggestGasPrice(ctx)
	if err != nil {
		log.Crit("Error getting gas price", "err", err)
	}
	log.Info("Got gasPrice", "gasPrice", gasPrice)

	nonce, err := cc.Eth.NonceAt(ctx, from, nil)
	if err != nil {
		log.Crit("Error getting nonce", "err", err)
	}
	log.Info("Got nonce", "nonce", nonce)

	gasLimit, err := cc.Eth.EstimateGas(ctx, ethereum.CallMsg{
		From:     from,
		To:       &to,
		GasPrice: gasPrice,
		Value:    value,
	})
	if err != nil {
		log.Crit("Error getting nonce", "err", err)
	}
	log.Info("Got gasLimit", "gasLimit", gasLimit)

	return types.NewTransaction(nonce, to, value, gasLimit, gasPrice, nil)
}

func CeloClient() *client.CeloClient {
	celo, err := client.Dial(config.FornoRC0Url)
	if err != nil {
		panic(err)
	}
	return celo
}

func PrintUsage(msg string) {
	fmt.Println(msg)
	fmt.Println("Usage: <cmd> <mnemonic> ")
	os.Exit(1)
}

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))
	if len(os.Args) != 2 {
		PrintUsage("Incorrect Arguments")
	}

	mnemonic := os.Args[1]

	cc := CeloClient()
	wallet := MustNewWallet(mnemonic)
	account := MustGetAccount(wallet, 0)

	log.Info("Got Account", "address", account.Address.Hex())

	balance, err := cc.Eth.BalanceAt(ctx, account.Address, nil)
	if err != nil {
		log.Crit("Error on BalanceAt", "err", err)
	}
	log.Info("Account Balance", "address", account.Address.Hex(), "balance", balance)

	tenGwei := big.NewInt(10000000000)
	tx := GoldTransferTx(cc, account.Address, beneficiary, tenGwei)

	signedTx, err := wallet.SignTx(*account, tx, nil)
	if err != nil {
		log.Crit("Error signing tx", "err", err)
	}

	txHash := signedTx.Hash()
	log.Info("Transaction Signed", "txHash", txHash.Hex())

	err = cc.Eth.SendTransaction(ctx, signedTx)
	if err != nil {
		log.Crit("Error Sending Transaction", "err", err)
	}

	log.Info("Waiting for receipt")
	for {
		// Wait For 10 seconds
		<-time.After(10 * time.Second)

		receipt, err := cc.Eth.TransactionReceipt(ctx, txHash)
		if err == ethereum.NotFound {
			continue
		} else if err != nil {
			log.Crit("Error Getting receipt", "err", err)
		} else {
			fmt.Println("Tx Mined:", receipt.BlockNumber, txHash.Hex())
			break
		}
	}

}
