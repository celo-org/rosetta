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
