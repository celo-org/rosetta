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
		//"enode://33ac194052ccd10ce54101c8340dbbe7831de02a3e7dcbca7fd35832ff8c53a72fd75e57ce8c8e73a0ace650dc2c2ec1e36f0440e904bc20a3cf5927f2323e85@34.83.199.225:30303",
		"enode://e99a883d0b7d0bacb84cde98c4729933b49adbc94e718b77fdb31779c7ed9da6c49236330a9ae096f42bcbf6e803394229120201704b7a4a3ae8004993fa0876@35.247.115.108:30303",
		// "enode://b3b42a9a6ef1125006f39b95850c624422eadb6340ac86e4023e47d5452914bb3d17340f9455cd1cdd0d246058b1fec2d3c500eeddbeafa49abd71f8f144b04e@34.83.73.169:30303", "enode://af5677afe5bf99a00bdb86d0f80f948b2e25f8978867b38cba8e860a6426507cbc37e15900f798305ceff6b7ac7f4057195827274d6b5f6a7e547ce512ff26ba@34.83.196.182:30303", "enode://02d18a52c4e097c12412ac3da9a9af24745cff182306e21fb1d3d691fe0c25f65c60586302b933bb9ec300b7fb00ed719d26a4d57af91d447691bac2b2d778af@34.83.107.15:30303", "enode://05977f6b7d3e16a99d27b714f8a029a006e41ec7732167d373dd920d31f72b3a1776650798d8763560854369d36867e9564dad13b4b60a90c347feeb491d83a9@34.83.42.50:30303", "enode://822366c6b9f80c3f3fdf7748209399ddd888bd54882958f6c75d05b8156c6274f54c8a1a6b468c3e85cade93a7dee2a2b701ccfc820b7669507d4bee855ebbf1@35.230.78.72:30303", "enode://5bb26439782fb6f9d0d997f907968f4ada618d49d83a2e6d202a107d7b17c67c680877ee733e2f92656714697e6f5eb0da25f26180c3e13d5dc39dc037160651@35.203.168.192:30303", "enode://29373f661cbea23f4f566d54470fae6ef5419c2a88aa52f306628df2d143f86807c02fd19b3d2d6d2e2a98d99a2db44643c6274e3aadd3632bd744a8be498768@34.83.1.123:30303", "enode://cc11ee6b035c8948dfaca5b09d676e28b0986585dac7a819376f12a8ee8b6b0ffd31907bb0d8bda27ef0a2296ee614d31773c7c5ea4333a121d80e0ba9bae801@35.185.234.63:30303", "enode://703cf979becdc501c4221090296fe75299cb9520f19a344098154c14c7133ebf6b649dad7f3f42947ad96312930bea5380a8ff86faa5a3795b0b6cc483adcfc8@34.82.82.249:30303",
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
	time.Sleep(10 * time.Second)
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
