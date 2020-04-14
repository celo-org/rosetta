package main

import (
	"context"
	"os"

	"github.com/celo-org/rosetta/internal/signals"
	"github.com/celo-org/rosetta/service"
	"github.com/celo-org/rosetta/service/geth"
	"github.com/ethereum/go-ethereum/log"
)

var (
	gethBinary  = os.ExpandEnv("${CELO_BLOCKCHAIN_PATH}/build/bin/geth")
	genesis     = "./envs/rc0/genesis.json"
	datadir     = "./envs/rc0/celo"
	sqlitepath  = "./envs/rc0/rosetta.db"
	staticNodes = []string{
		"enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303",
	}
)

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlDebug, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	ctx, stopFn := context.WithCancel(context.Background())
	defer stopFn()

	gethSrv := geth.NewGethService(gethBinary, datadir, genesis, staticNodes)

	sm := service.NewServiceManager(ctx)
	sm.Add(gethSrv)

	go func() {
		gotExitSignal := signals.WatchForExitSignals()
		<-gotExitSignal
		stopFn()
	}()

	if err := sm.Wait(); err != nil {
		log.Error("Error running Services", "err", err)
		os.Exit(1)
	}
}
