/*
Copyright © 2020 Celo Org

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
package cmd

import (
	"context"

	"github.com/celo-org/rosetta/api"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/internal/signals"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/spf13/cobra"
)

// remoteCmd represents the remote command
var remoteCmd = &cobra.Command{
	Use:   "remote",
	Short: "Launch Rosetta using a remote celo-blockchain node",
	Long: `On this mode, rosetta won't start its own celo-blockchain instance.
Instead it will connect to a given node.

This mode should be used only during development`,
	Args: cobra.NoArgs,
	Run:  runRemoteCmd,
}

var nodeUri string
var epochSize uint64

func init() {
	serveCmd.AddCommand(remoteCmd)

	remoteCmd.Flags().StringVar(&nodeUri, "nodeUri", "", "Connection URI for the celo-blockchain node")
	exitOnError(remoteCmd.MarkFlagRequired("nodeUri"))

	remoteCmd.Flags().Uint64Var(&epochSize, "epoch", 0, "Epoch Size on the network")
}

func runRemoteCmd(cmd *cobra.Command, args []string) {

	rpcClient, err := rpc.Dial(nodeUri)
	if err != nil {
		log.Crit("Can't connect to node", "err", err)
	}

	cc := client.NewCeloClient(rpcClient)

	log.Debug("Obtaining chain id")
	chainId, err := cc.Net.ChainId(context.Background())
	if err != nil {
		log.Crit("Error fetching chainId", "err", err)
	}

	chainParams := &celo.ChainParameters{
		ChainId:   chainId,
		EpochSize: epochSize,
	}

	log.Info("Initializing Rosetta In Remote Mode..", "chainId", chainId, "epochSize", epochSize, "nodeUri", nodeUri)
	rosettaServer := api.NewRosettaServer(cc, getRosettaServerConfig(), chainParams)

	go rosettaServer.Start()
	defer rosettaServer.Stop()

	gotExitSignal := signals.WatchForExitSignals()
	<-gotExitSignal
}
