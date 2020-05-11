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

	"github.com/celo-org/rosetta/internal/signals"
	"github.com/celo-org/rosetta/service"
	"github.com/celo-org/rosetta/service/geth"
	"github.com/ethereum/go-ethereum/log"
)

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
