package main

import (
	"context"
	"os"
	"time"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/internal/signals"
	"github.com/celo-org/rosetta/service"
	"github.com/celo-org/rosetta/service/geth"
	mservice "github.com/celo-org/rosetta/service/monitor"
	"github.com/ethereum/go-ethereum/log"
)

// Configuration Parameters
var (
	//gethBinary  = "/Users/alec/celo/go/src/github.com/ethereum/go-ethereum/build/bin/geth"
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

	celoStore, err := db.NewSqliteDb(sqlitepath)
	if err != nil {
		log.Error("Error opening CeloStore", "err", err)
		os.Exit(1)
	}

	ctx, stopFn := context.WithCancel(context.Background())
	defer stopFn()

	go func() {
		gotExitSignal := signals.WatchForExitSignals()
		<-gotExitSignal
		stopFn()
	}()

	sm := service.NewServiceManager(ctx)
	gethSrv := geth.NewGethService(gethBinary, datadir, genesis, staticNodes)
	sm.Add(gethSrv)

	// Wait for geth to start and connect to it
	time.Sleep(5 * time.Second)
	// cc, err := client.Dial("http://localhost:8545")
	// cc, err := client.Dial("ws://localhost:8546")
	cc, err := client.Dial(gethSrv.IpcFilePath())
	if err != nil {
		log.Error("Error on client connection to geth", "err", err)
		os.Exit(1)
	}

	sm.Add(mservice.NewMonitorService(cc, celoStore))

	if err = sm.Wait(); err != nil {
		log.Error("Error running Services", "err", err)
		os.Exit(1)
	}
}
