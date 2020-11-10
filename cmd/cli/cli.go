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
	"strconv"
	"strings"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/rosetta/cmd/internal/utils"
	"github.com/celo-org/rosetta/db"
	"github.com/coinbase/rosetta-sdk-go/fetcher"
	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/spf13/cobra"
)

// cliCmd represents the cli command
var CliCmd = &cobra.Command{
	Use:   "cli",
	Short: "Group commands to query rosetta-rpc",
}

var serverUrl string
var dbPath string

func init() {
	CliCmd.AddCommand(blockCmd)
	CliCmd.AddCommand(reconcileCmd)

	CliCmd.PersistentFlags().StringVar(&serverUrl, "url", "http://localhost:8080", "Base url for rosetta rpc")
	CliCmd.PersistentFlags().StringVar(&dbPath, "db", "./envs/alfajores/rosetta.db", "RosettaDb path")
	CliCmd.PersistentFlags().StringVar(&dbPath, "nodeUrl", "http://localhost:8545", "Geth Node url")
}

func getFetcher() (*fetcher.Fetcher, *types.NetworkIdentifier, *types.NetworkStatusResponse) {
	ctx := context.Background()
	fetcher := fetcher.New(serverUrl)

	// Step 2: Initialize the fetcher's asserter
	//
	// Behind the scenes this makes a call to get the
	// network status and uses the response to inform
	// the asserter what are valid responses.
	primaryNetwork, networkStatus, err := fetcher.InitializeAsserter(ctx, nil)
	utils.ExitOnError(err.Err)

	return fetcher, primaryNetwork, networkStatus
}

func getDb() db.RosettaDB {
	celoStore, err := db.NewSqliteDb("./envs/alfajores/rosetta.db")
	utils.ExitOnError(err)
	return celoStore
}

func getCeloClient() *client.CeloClient {
	cc, err := client.Dial("http://localhost:8545")
	utils.ExitOnError(err)
	return cc
}

func toBlockIdentifier(arg string) (blockIdentifier *types.PartialBlockIdentifier) {
	if strings.HasPrefix("0x", arg) {
		blockIdentifier = &types.PartialBlockIdentifier{
			Hash: &arg,
		}
	} else {
		number, err := strconv.ParseInt(arg, 10, 0)
		utils.ExitOnError(err)
		blockIdentifier = &types.PartialBlockIdentifier{
			Index: &number,
		}
	}
	return
}
