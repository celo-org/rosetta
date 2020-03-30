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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"

	"github.com/celo-org/rosetta/api"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/internal/fileutils"
	"github.com/celo-org/rosetta/internal/signals"
	"github.com/ethereum/go-ethereum/log"
	"github.com/spf13/cobra"
)

// localCmd represents the local command
var localCmd = &cobra.Command{
	Use:   "local",
	Short: "Launch Rosetta using a local celo-blockchain node",
	Long:  `On this mode, rosetta will spawn a celo-blockchain node`,
	Run:   runLocalCmd,
}

var genesisPath string
var gethBinary string
var staticNodes []string

func init() {
	serveCmd.AddCommand(localCmd)

	localCmd.Flags().StringVar(&gethBinary, "geth", "", "Path to the celo-blockchain binary")
	exitOnError(localCmd.MarkFlagRequired("geth"))
	exitOnError(localCmd.MarkFlagFilename("geth"))

	localCmd.Flags().StringVar(&genesisPath, "genesis", "", "path to the genesis.json")
	exitOnError(localCmd.MarkFlagRequired("genesis"))
	exitOnError(localCmd.MarkFlagFilename("genesis", "json"))

	localCmd.Flags().StringArrayVar(&staticNodes, "staticNode", []string{}, "StaticNode to use (can be repeated many times")
	exitOnError(localCmd.MarkFlagRequired("staticNode"))
}

func runLocalCmd(cmd *cobra.Command, args []string) {
	// Read Genesis to get chain parameters
	chainParams := chainParamsFromGenesisFile()

	if err := os.MkdirAll(datadir.GethDatadir(), os.ModePerm); err != nil {
		log.Crit("Can't create celo-blockchain datadir")
	}

	staticNodesRaw, err := json.Marshal(staticNodes)
	if err != nil {
		log.Crit("Can't serialize static nodes", "err", err)
	}
	err = ioutil.WriteFile(datadir.GethStaticNodesFile(), staticNodesRaw, 0644)
	if err != nil {
		log.Crit("Can't serialize static nodes", "err", err)
	}

	if len(staticNodes) == 0 {
		cmd.PrintErrln(cmd.UsageString())
		cmd.PrintErrln(cmd.UsageString())

	}

	ensureGethInit()

	log.Info("Starting local geth")
	gethCmd := startGeth(chainParams)
	cmdExit := make(chan struct{})
	go func() {
		err := gethCmd.Wait()
		if err != nil {
			log.Error("Error runnin geth", "err", err)
			close(cmdExit)
		}

	}()
	// give geth some time to start
	time.Sleep(5 * time.Second)

	log.Info("Initializing Rosetta In Local Mode..", "chainId", chainParams.ChainId, "epochSize", chainParams.EpochSize)
	cc, err := client.Dial(datadir.GethIpcFile())
	if err != nil {
		log.Crit("Error on client connection to geth", "err", err)
	}

	rosettaServer := api.NewRosettaServer(cc, getRosettaServerConfig(), chainParams)

	go rosettaServer.Start()
	defer rosettaServer.Stop()

	gotExitSignal := signals.WatchForExitSignals()
	select {
	case <-gotExitSignal:
		if err := gethCmd.Process.Signal(os.Interrupt); err != nil {
			log.Crit("Error sending kil signal to geth", "err", err)
		}
	case <-cmdExit:
	}
}

func chainParamsFromGenesisFile() *celo.ChainParameters {
	data, err := ioutil.ReadFile(genesisPath)
	if err != nil {
		log.Crit("Can't read genesis.json on DataDir", "genesisPath", genesisPath, "err", err)
	}

	// We only map the fields we need
	var genesis struct {
		Config struct {
			ChainId  uint64 `json:"chainId"`
			Isntabul struct {
				Epoch uint64 `json:"epoch"`
			} `json:"istanbul"`
		} `json:"config"`
	}

	if err = json.Unmarshal(data, &genesis); err != nil {
		log.Crit("Can't parse genesis.json on DataDir", "genesisPath", genesisPath, "err", err)
	}

	return &celo.ChainParameters{
		ChainId:   new(big.Int).SetUint64(genesis.Config.ChainId),
		EpochSize: genesis.Config.Isntabul.Epoch,
	}
}

func gethCmd(args ...string) *exec.Cmd {
	fullArgs := append([]string{"--datadir", datadir.GethDatadir()}, args...)
	return exec.Command(gethBinary, fullArgs...)
}

func ensureGethInit() {
	// Check if geth is initialized already
	flagFile := datadir.GethInitializedFile()

	if fileutils.FileExists(flagFile) {
		log.Info("Geth Already initialized... skipping init")
		return
	} else {
		log.Info("Running geth init")
		out, err := gethCmd("init", genesisPath).CombinedOutput()
		if err != nil {
			log.Error("Error running geth init", "err", err)
			fmt.Println(string(out))
			os.Exit(1)
		}
		if err = fileutils.TouchFile(flagFile); err != nil {
			log.Crit("Error creating marker file", "err", err)
		}
	}
}

func startGeth(chainParams *celo.ChainParameters) *exec.Cmd {
	gethArgs := []string{
		"--verbosity", "3",
		"--syncmode", "full",
		"--networkid", chainParams.ChainId.String(),
		"--rpc",
		"--rpcaddr", "127.0.0.1",
		"--rpcapi", "eth,net,web3,debug,admin,personal",
		"--light.serve", "0",
		"--light.maxpeers", "0",
		"--maxpeers", "1100",
		"--gcmode", "archive",
		"--consoleformat", "term",
		"--consoleoutput", "split",
	}

	fmt.Println("geth", strings.Join(gethArgs, " "))
	f, err := os.OpenFile(datadir.GethLogFile(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Crit("Can't open geth logfile", "err", err)
	}
	defer f.Close()

	cmd := gethCmd(gethArgs...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setsid: true}
	cmd.Stdout = f
	cmd.Stderr = os.Stdout

	if err = cmd.Start(); err != nil {
		log.Crit("Error starting geth", "err", err)
	}

	return cmd
}
