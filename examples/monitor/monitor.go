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
	gethBinary  = os.ExpandEnv("${CELO_BLOCKCHAIN_PATH}/build/bin/geth")
	genesis     = "./envs/alfajores/genesis.json"
	datadir     = "./envs/alfajores/celo"
	sqlitepath  = "./envs/alfajores/rosetta.db"
	staticNodes = []string{
		"enode://e99a883d0b7d0bacb84cde98c4729933b49adbc94e718b77fdb31779c7ed9da6c49236330a9ae096f42bcbf6e803394229120201704b7a4a3ae8004993fa0876@35.247.115.108:30303",
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

	gethOpts := &geth.GethOpts{
		GethBinary:  gethBinary,
		GenesisPath: genesis,
		Datadir:     datadir,
		StaticNodes: staticNodes,
	}

	gethSrv := geth.NewGethService(gethOpts)
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
