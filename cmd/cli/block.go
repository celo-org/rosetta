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

	"github.com/celo-org/rosetta/cmd/internal/utils"
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

		blockIdentifier := toBlockIndentifier(args[0])
		block, err := fetcher.BlockRetry(ctx, network, blockIdentifier)
		utils.ExitOnError(err)

		utils.PrettyPrint(block)

	},
}
