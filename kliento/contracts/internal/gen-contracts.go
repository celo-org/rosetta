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
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
)

const contractsPath = "contracts"

const baseBuildPathMonorepo = "packages/protocol/build/"
const coreContractsBuildDir = "contracts"
const coreContracts08BuildDir = "contracts-0.8"
const mentoContractsBuildDir = "contracts-mento"

var coreContractsToGenerate = []string{
	"Accounts",
	"Attestations",
	"BlockchainParameters",
	"DoubleSigningSlasher",
	"DowntimeSlasher",
	"Election",
	"EpochRewards",
	"Escrow",
	"FeeCurrencyWhitelist",
	"FeeHandler",
	"Freezer",
	"GoldToken",
	"Governance",
	"GovernanceSlasher",
	"GovernanceApproverMultiSig",
	"LockedGold",
	"MultiSig",
	"Proxy",
	"Random",
	"Registry",
	"ReleaseGold",
	"Validators",
}
var coreContracts08ToGenerate = []string{
	"GasPriceMinimum",
}

var mentoContractsToGenerate = []string{
	"Exchange",
	"Reserve",
	"ReserveSpenderMultiSig",
	"SortedOracles",
	"StableToken",
}

func main() {
	monorepoPath := flag.String("monorepo", "", "Path to celo-monorepo")
	celoBlockchainPath := flag.String("gcelo", "", "Path to celo-blockchain")

	flag.Parse()
	fmt.Println(*monorepoPath, *celoBlockchainPath)

	if *monorepoPath == "" || *celoBlockchainPath == "" {
		exitWithHelpMessage()
	}

	validatePathExists(*monorepoPath)
	validatePathExists(*celoBlockchainPath)

	if pathExists(contractsPath) {
		files, err := filepath.Glob(filepath.Join(contractsPath, "gen_*.go"))
		if err != nil {
			exitMessage("Error running glob: %s", err)
		}

		for _, file := range files {
			if err := os.Remove(file); err != nil {
				exitMessage("Error removing "+file+" directory: %s\n", err)
			}
		}

	}
	if err := os.MkdirAll(contractsPath, os.ModePerm); err != nil {
		exitMessage("Error creating "+contractsPath+" directory: %s\n", err)
	}

	abigen := path.Join(*celoBlockchainPath, "build/bin", "abigen")

	buildMonorepoContracts(abigen, *monorepoPath, coreContractsBuildDir, coreContractsToGenerate)
	buildMonorepoContracts(abigen, *monorepoPath, coreContracts08BuildDir, coreContracts08ToGenerate)
	buildMonorepoContracts(abigen, *monorepoPath, mentoContractsBuildDir, mentoContractsToGenerate)
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func validatePathExists(dirpath string) {
	if _, err := os.Stat(dirpath); err != nil {
		if os.IsNotExist(err) {
			fmt.Printf("Path %s does not exists", dirpath)
		} else {
			fmt.Printf("Can't access %s: %s", dirpath, err)
		}
		exitWithHelpMessage()
	}
}

func exitWithHelpMessage() {
	flag.PrintDefaults()
	os.Exit(1)
}

func exitMessage(msg string, a ...interface{}) {
	fmt.Printf(msg, a...)
	os.Exit(1)
}

// mustRun executes the given command and exits the host process for
// any error.
func mustRun(cmd *exec.Cmd) {
	fmt.Println(">>>", strings.Join(cmd.Args, " "))
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		log.Printf("Command failed \"%s\", err: \"%v\"", strings.Join(cmd.Args, " "), err)
		log.Fatal(fmt.Sprintf("Command failed \"%s\", err: \"%v\"", strings.Join(cmd.Args, " "), err))
	}
}

func mustRunCommand(cmd string, args ...string) {
	mustRun(exec.Command(cmd, args...))
}

func buildContract(abigenPath, contractBuildDir, contractName string) {
	contractTrufflePath := path.Join(contractBuildDir, contractName+".json")
	validatePathExists(contractTrufflePath)
	mustRunCommand(abigenPath, "--truffle", contractTrufflePath,
		"--pkg", "contracts", "--type", contractName,
		"--out", path.Join(contractsPath, "gen_"+strings.ToLower(contractName)+".go"))
}

func buildMonorepoContracts(abigenPath, monorepoPath, contractBuildDir string, contractNames []string) {
	fullContractDirPath := path.Join(monorepoPath, baseBuildPathMonorepo, contractBuildDir)
	for _, contract := range contractNames {
		buildContract(abigenPath, fullContractDirPath, contract)
	}
}
