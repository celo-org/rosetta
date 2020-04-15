package main

import (
	"flag"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/celo-org/rosetta/internal/build"
)

const contractsPath = "celo/contract"

var contractsToGenerate = []string{
	"Registry",
	"LockedGold",
	"Election",
	"StableToken",
	"GasPriceMinimum",
	"Election",
	"EpochRewards",
	"Validators",
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
		if err := os.RemoveAll(contractsPath); err != nil {
			exitMessage("Error removing "+contractsPath+" directory: %s\n", err)
		}
	}
	if err := os.MkdirAll(contractsPath, os.ModePerm); err != nil {
		exitMessage("Error creating "+contractsPath+" directory: %s\n", err)
	}

	abigen := path.Join(*celoBlockchainPath, "build/bin", "abigen")

	for _, contract := range contractsToGenerate {
		contractTrufflePath := path.Join(*monorepoPath, "packages/protocol/build/contracts/", contract+".json")
		validatePathExists(contractTrufflePath)
		build.MustRunCommand(abigen, "--truffle", contractTrufflePath,
			"--pkg", "contract", "--type", contract,
			"--out", path.Join(contractsPath, strings.ToLower(contract)+".go"))
	}
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
