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
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/cmd/internal/utils"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/internal/fileutils"
	"github.com/celo-org/rosetta/internal/signals"
	"github.com/celo-org/rosetta/service"
	"github.com/celo-org/rosetta/service/geth"
	"github.com/celo-org/rosetta/service/monitor"
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

var staticNodes []string

func init() {
	RootCmd.AddCommand(serveCmd)

	flagSet := serveCmd.Flags()

	// Common Flags
	flagSet.String("datadir", "", "datadir to use")
	utils.ExitOnError(viper.BindPFlag("datadir", flagSet.Lookup("datadir")))
	utils.ExitOnError(serveCmd.MarkFlagDirname("datadir"))

	flagSet.String("logs", "", "Path to logs file")
	utils.ExitOnError(viper.BindPFlag("logs", flagSet.Lookup("logs")))
	utils.ExitOnError(serveCmd.MarkFlagDirname("logs"))

	// RPC Service Flags
	flagSet.UintVar(&rosettaRpcConfig.Port, "port", 8080, "Listening port for http server")
	flagSet.StringVar(&rosettaRpcConfig.Interface, "address", "", "Listening address for http server")
	flagSet.DurationVar(&rosettaRpcConfig.RequestTimeout, "reqTimeout", 25*time.Second, "Timeout when serving a request")

	// Geth Service Flags
	flagSet.String("geth", "", "Path to the celo-blockchain binary")
	utils.ExitOnError(viper.BindPFlag("geth", flagSet.Lookup("geth")))
	utils.ExitOnError(serveCmd.MarkFlagFilename("geth"))

	flagSet.String("ipcpath", "", "Path to the geth ipc file")
	utils.ExitOnError(viper.BindPFlag("ipcpath", flagSet.Lookup("ipcpath")))
	utils.ExitOnError(serveCmd.MarkFlagFilename("ipcpath"))

	flagSet.String("genesis", "", "path to the genesis.json")
	utils.ExitOnError(viper.BindPFlag("genesis", flagSet.Lookup("genesis")))
	utils.ExitOnError(serveCmd.MarkFlagFilename("genesis", "json"))

	flagSet.StringArrayVar(&staticNodes, "staticNode", []string{}, "StaticNode to use (can be repeated many times")
	utils.ExitOnError(serveCmd.MarkFlagRequired("staticNode"))

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
	exitOnMissingConfig(cmd, "geth")
	exitOnMissingConfig(cmd, "genesis")

	gethOpts := &geth.GethOpts{
		GethBinary:  viper.GetString("geth"),
		GenesisPath: viper.GetString("genesis"),
		Datadir:     filepath.Join(datadir, "celo"),
		LogsPath:    viper.GetString("logs"),
		IpcPath:     viper.GetString("ipcpath"),
		StaticNodes: staticNodes,
	}

	sqlitePath := filepath.Join(datadir, "rosetta.db")

	// TODO - create context that encapsulate Stop on Signal behaviour
	srvCtx, stopServices := context.WithCancel(context.Background())
	defer stopServices()

	gotExitSignal := signals.WatchForExitSignals()
	go func() {
		<-gotExitSignal
		stopServices()
	}()

	if err := runAllServices(srvCtx, sqlitePath, gethOpts); err != nil {
		log.Error("Rosetta run failed", "err", err)
		os.Exit(1)
	}
}

func runAllServices(ctx context.Context, sqlitePath string, gethOpts *geth.GethOpts) error {
	ctx, stopServices := context.WithCancel(ctx)
	defer stopServices()

	celoStore, err := db.NewSqliteDb(sqlitePath)
	if err != nil {
		return fmt.Errorf("can't open rosetta.db: %w", err)
	}

	sm := service.NewServiceManager(ctx)

	gethSrv := geth.NewGethService(gethOpts)

	if err := gethSrv.Setup(); err != nil {
		return fmt.Errorf("error on geth setup: %w", err)
	}

	sm.Add(gethSrv)

	gethStarted := utils.WaitUntil(500*time.Millisecond, 30*time.Second, func() bool {
		return fileutils.FileExists(gethSrv.IpcFilePath())
	})

	if !gethStarted {
		return fmt.Errorf("geth service failed to start before timeout: %d", time.Millisecond)
	}

	chainParams := gethSrv.ChainParameters()
	log.Info("Detected Chain Parameters", "chainId", chainParams.ChainId, "epochSize", chainParams.EpochSize)

	cc, err := client.Dial(gethSrv.IpcFilePath())
	if err != nil {
		return fmt.Errorf("can't connect to geth: %w", err)
	}

	rpcService, err := rpc.NewRosettaServer(cc, celoStore, &rosettaRpcConfig, chainParams)
	if err != nil {
		return fmt.Errorf("can't create rpc server: %w", err)
	}

	sm.Add(rpcService)
	sm.Add(monitor.NewMonitorService(cc, celoStore))
	return sm.Wait()
}
