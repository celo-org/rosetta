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
	"strings"
	"time"

	"github.com/celo-org/celo-blockchain/log"
	"github.com/celo-org/kliento/client"
	"github.com/celo-org/rosetta/cmd/internal/utils"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/internal/fileutils"
	"github.com/celo-org/rosetta/internal/signals"
	"github.com/celo-org/rosetta/service"
	"github.com/celo-org/rosetta/service/geth"
	"github.com/celo-org/rosetta/service/monitor"
	"github.com/celo-org/rosetta/service/rpc"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "run",
	Short: "Start rosetta server",
	Args:  cobra.NoArgs,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("ROSETTA")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()

		return viper.BindPFlags(cmd.Flags())
	},
	Run: runRunCmd,
}

type ConfigPaths string

func init() {
	flagSet := serveCmd.Flags()

	// Common Flags
	flagSet.String("datadir", "", "datadir to use")
	utils.ExitOnError(serveCmd.MarkFlagDirname("datadir"))

	// RPC Service Flags
	flagSet.Uint("rpc.port", 8080, "Listening port for http server")
	flagSet.String("rpc.address", "", "Listening address for http server")
	flagSet.Duration("rpc.reqTimeout", 25*time.Second, "Timeout for requests to this service, this also controls the timeout sent to the blockchain node for trace transaction requests")

	// Geth Service Flags
	flagSet.String("geth.binary", "", "Path to the celo-blockchain binary")
	utils.ExitOnError(serveCmd.MarkFlagFilename("geth.binary"))

	flagSet.String("geth.logfile", "", "Path to logs file")
	utils.ExitOnError(serveCmd.MarkFlagDirname("geth.logfile"))

	flagSet.String("geth.ipcpath", "", "Path to the geth ipc file")
	utils.ExitOnError(serveCmd.MarkFlagFilename("geth.ipcpath"))

	flagSet.String("geth.genesis", "", "(Optional) path to the genesis.json, for use with custom chains")
	utils.ExitOnError(serveCmd.MarkFlagFilename("geth.genesis", "json"))
	flagSet.String("geth.network", "", "Network to use, either 'mainnet', 'alfajores', or 'baklava'")

	flagSet.String("geth.staticnodes", "", "StaticNode to use (separated by ,)")
	flagSet.String("geth.bootnodes", "", "Bootnodes to use (separated by ,)")
	flagSet.String("geth.verbosity", "", "Geth log verbosity (number between [1-5])")
	flagSet.String("geth.publicip", "", "Public Ip to configure geth (sometimes required for discovery)")
	flagSet.String("geth.cache", "", "Memory (in MB) allocated to geth's internal caching")

	flagSet.String("geth.rpcaddr", "127.0.0.1", "Geth HTTP-RPC server listening interface")
	flagSet.String("geth.rpcport", "8545", "Geth HTTP-RPC server listening port")
	flagSet.String("geth.rpcvhosts", "localhost", "Geth comma separated list of virtual hostnames from which to accept requests")

	flagSet.String("geth.syncmode", "fast", "Geth blockchain sync mode (fast, full, light)")
	flagSet.String("geth.gcmode", "full", "Geth garbage collection mode (full, archive)")
	flagSet.String("geth.maxpeers", "1100", "Maximum number of network peers (network disabled if set to 0) (default: 1100)")
}

func getDatadir(cmd *cobra.Command) string {
	exitOnMissingConfig(cmd, "datadir")

	absDatadir, err := filepath.Abs(viper.GetString("datadir"))
	if err != nil {
		printUsageAndExit(cmd, fmt.Sprintf("Can't resolve datadir path: %s, error: %s", absDatadir, err))
	}

	return absDatadir
}

func readGethOption(cmd *cobra.Command, datadir string) *geth.GethOpts {
	opts := &geth.GethOpts{
		GethBinary:  viper.GetString("geth.binary"),
		GenesisPath: viper.GetString("geth.genesis"),
		Network:     viper.GetString("geth.network"),
		Datadir:     filepath.Join(datadir, "celo"),
		LogsPath:    viper.GetString("geth.logfile"),
		IpcPath:     viper.GetString("geth.ipcpath"),
		Bootnodes:   viper.GetString("geth.bootnodes"),
		Verbosity:   viper.GetString("geth.verbosity"),
		StaticNodes: viper.GetString("geth.staticnodes"),
		PublicIp:    viper.GetString("geth.publicip"),
		Cache:       viper.GetString("geth.cache"),
		RpcAddr:     viper.GetString("geth.rpcaddr"),
		RpcPort:     viper.GetString("geth.rpcport"),
		RpcVHosts:   viper.GetString("geth.rpcvhosts"),
		SyncMode:    viper.GetString("geth.syncmode"),
		GcMode:      viper.GetString("geth.gcmode"),
		MaxPeers:    viper.GetString("geth.maxpeers"),
	}

	if opts.GethBinary == "" {
		printUsageAndExit(cmd, "Missing config option for 'geth.binary'")
	}
	if opts.GenesisPath == "" && opts.Network == "" {
		printUsageAndExit(cmd, "Missing config option for 'geth.genesis' or 'geth.network'")
	} else if opts.GenesisPath != "" && opts.Network != "" {
		printUsageAndExit(cmd, "Must provide exactly one of 'geth.genesis' or 'geth.network'")
	}

	return opts
}

func runRunCmd(cmd *cobra.Command, args []string) {
	datadir := getDatadir(cmd)
	gethOpts := readGethOption(cmd, datadir)
	sqlitePath := filepath.Join(datadir, "rosetta.db")

	rpcConfig :=
		&rpc.RosettaServerConfig{
			Interface:      viper.GetString("rpc.address"),
			Port:           viper.GetUint("rpc.port"),
			RequestTimeout: viper.GetDuration("rpc.reqTimeout"),
		}

	// TODO - create context that encapsulate Stop on Signal behaviour
	srvCtx, stopServices := context.WithCancel(context.Background())
	defer stopServices()

	gotExitSignal := signals.WatchForExitSignals()
	go func() {
		<-gotExitSignal
		stopServices()
	}()

	if err := runAllServices(srvCtx, sqlitePath, gethOpts, rpcConfig); err != nil {
		log.Error("Rosetta run failed", "err", err)
		os.Exit(1)
	}
}

func runAllServices(ctx context.Context, sqlitePath string, gethOpts *geth.GethOpts, rpcConfig *rpc.RosettaServerConfig) error {
	ctx, stopServices := context.WithCancel(ctx)
	defer stopServices()

	sm := service.NewServiceManager(ctx)

	gethSrv := geth.NewGethService(gethOpts)

	if err := gethSrv.Setup(); err != nil {
		return fmt.Errorf("error on geth setup: %w", err)
	}

	celoStore, err := db.NewSqliteDb(sqlitePath)
	if err != nil {
		return fmt.Errorf("can't open rosetta.db: %w", err)
	}

	sm.Add(gethSrv)

	gethStarted := utils.WaitUntil(500*time.Millisecond, 5*time.Minute, func() bool {
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

	rpcService, err := rpc.NewRosettaServer(cc, celoStore, rpcConfig, chainParams)
	if err != nil {
		return fmt.Errorf("can't create rpc server: %w", err)
	}

	sm.Add(rpcService)
	sm.Add(monitor.NewMonitorService(cc, celoStore))
	return sm.Wait()
}
