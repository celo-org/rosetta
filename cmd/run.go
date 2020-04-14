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
	"os"
	"path/filepath"
	"time"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/internal/fileutils"
	"github.com/celo-org/rosetta/internal/signals"
	"github.com/celo-org/rosetta/service"
	"github.com/celo-org/rosetta/service/geth"
	"github.com/celo-org/rosetta/service/rpc"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "run",
	Short: "Start rosetta server",
	Args:  cobra.NoArgs,
	Run:   runRunCmd,
}

var rosettaRpcConfig rpc.RosettaServerConfig

type ConfigPaths string

var gethBinary string
var staticNodes []string

func init() {
	rootCmd.AddCommand(serveCmd)

	flagSet := serveCmd.Flags()

	// Common Flags
	flagSet.String("datadir", "", "datadir to use")
	exitOnError(viper.BindPFlag("datadir", flagSet.Lookup("datadir")))
	exitOnError(serveCmd.MarkFlagDirname("datadir"))

	// RPC Service Flags
	flagSet.UintVar(&rosettaRpcConfig.Port, "port", 8080, "Listening port for http server")
	flagSet.StringVar(&rosettaRpcConfig.Interface, "address", "", "Listening address for http server")
	flagSet.DurationVar(&rosettaRpcConfig.RequestTimeout, "reqTimeout", 25*time.Second, "Timeout when serving a request")

	// Geth Service Flags
	flagSet.String("geth", "", "Path to the celo-blockchain binary")
	exitOnError(viper.BindPFlag("geth", flagSet.Lookup("geth")))
	exitOnError(serveCmd.MarkFlagFilename("geth"))

	flagSet.String("genesis", "", "path to the genesis.json")
	exitOnError(viper.BindPFlag("genesis", flagSet.Lookup("genesis")))
	exitOnError(serveCmd.MarkFlagFilename("genesis", "json"))

	flagSet.StringArrayVar(&staticNodes, "staticNode", []string{}, "StaticNode to use (can be repeated many times")
	exitOnError(serveCmd.MarkFlagRequired("staticNode"))

	// Monitor Service Flags

}

func getDatadir(cmd *cobra.Command) string {
	exitOnMissingConfig(cmd, "datadir")

	absDatadir, err := filepath.Abs(viper.GetString("datadir"))
	if err != nil {
		log.Crit("Can't resolve datadir path", "datadir", absDatadir, "err", err)
	}

	isDir, err := fileutils.IsDirectory(absDatadir)
	if err != nil {
		log.Crit("Can't access datadir", "datadir", absDatadir, "err", err)
	} else if !isDir {
		log.Crit("Datadir is not a directory", "datadir", absDatadir)
	}

	return absDatadir
}

func runRunCmd(cmd *cobra.Command, args []string) {
	datadir := getDatadir(cmd)
	gethDataDir := filepath.Join(datadir, "celo")

	exitOnMissingConfig(cmd, "geth")
	exitOnMissingConfig(cmd, "genesis")

	gethBinary = viper.GetString("geth")
	genesisPath := viper.GetString("genesis")

	// TODO - create context that encapsulate Stop on Signal behaviour
	srvCtx, stopServices := context.WithCancel(context.Background())
	defer stopServices()

	gotExitSignal := signals.WatchForExitSignals()
	go func() {
		<-gotExitSignal
		stopServices()
	}()

	sm := service.NewServiceManager(srvCtx)

	gethSrv := geth.NewGethService(
		gethBinary,
		gethDataDir,
		genesisPath,
		staticNodes,
	)

	if err := gethSrv.Setup(); err != nil {
		log.Error("Error on geth setup", "err", err)
		os.Exit(1)
	}

	sm.Add(gethSrv)

	chainParams := gethSrv.ChainParameters()
	log.Info("Detected Chain Parameters", "chainId", chainParams.ChainId, "epochSize", chainParams.EpochSize)

	nodeUri := gethSrv.IpcFilePath()
	log.Debug("celo nodes ipc file", "filepath", nodeUri)

	time.Sleep(5 * time.Second)
	cc, err := client.Dial(nodeUri)
	if err != nil {
		log.Error("Error on client connection to geth", "err", err)
		os.Exit(1)
	}

	sm.Add(rpc.NewRosettaServer(cc, &rosettaRpcConfig, chainParams))

	if err := sm.Wait(); err != nil {
		log.Error("Error running services", "err", err)
		os.Exit(1)
	}
}
