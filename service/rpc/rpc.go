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

package rpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/service"
	"github.com/coinbase/rosetta-sdk-go/asserter"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/felixge/httpsnoop"
)

type RosettaServerConfig struct {
	Port           uint
	Interface      string
	RequestTimeout time.Duration
}

func (hs *RosettaServerConfig) ListenAddress() string {
	return fmt.Sprintf("%s:%d", hs.Interface, hs.Port)
}

type rosettaServer struct {
	cc          *client.CeloClient
	cfg         *RosettaServerConfig
	chainParams *celo.ChainParameters

	running service.RunningLock
	server  *http.Server
}

func NewRosettaServer(cc *client.CeloClient, db db.RosettaDBReader, cfg *RosettaServerConfig, chainParams *celo.ChainParameters) (*rosettaServer, error) {
	var mainHandler http.Handler
	var err error

	mainHandler, err = createRouter(cc, db, chainParams)
	if err != nil {
		return nil, err
	}

	mainHandler = requestLogHandler(mainHandler)
	mainHandler = http.TimeoutHandler(mainHandler, cfg.RequestTimeout, "Request Timed out")

	server := &http.Server{
		Addr:         cfg.ListenAddress(),
		Handler:      mainHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// TODO set ErrorLog: ,
	}

	return &rosettaServer{
		cc:          cc,
		cfg:         cfg,
		server:      server,
		chainParams: chainParams,
	}, nil
}

func (rs *rosettaServer) Name() string {
	return "rosetta-rpc"
}

func (rs *rosettaServer) Running() bool {
	return rs.running.Running()
}

func (rs *rosettaServer) Start(ctx context.Context) error {
	if err := rs.running.EnableOrFail(); err != nil {
		return err
	}
	defer rs.running.Disable()

	go func() {
		<-ctx.Done()
		if err := rs.server.Close(); err != nil {
			panic("Failed to Stop the service")
		}
	}()

	log.Info("Starting httpServer", "listen_address", rs.cfg.ListenAddress())
	err := rs.server.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}
	return err
}

func requestLogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, w, r)
		log.Info("RequestServed", "method", r.Method, "uri", r.URL, "code", m.Code, "duration", m.Duration, "bytes", m.Written)
	})
}

func createRouter(celoClient *client.CeloClient, db db.RosettaDBReader, chainParams *celo.ChainParameters) (http.Handler, error) {
	servicer, err := NewServicer(celoClient, db, chainParams)
	if err != nil {
		return nil, err
	}

	network := &types.NetworkIdentifier{
		Blockchain: BlockchainName,
		Network:    chainParams.ChainId.String(),
	}

	asserter, err := asserter.NewServer([]*types.NetworkIdentifier{network})
	if err != nil {
		return nil, err
	}

	AccountApiController := server.NewAccountAPIController(servicer, asserter)
	BlockApiController := server.NewBlockAPIController(servicer, asserter)
	ConstructionApiController := server.NewConstructionAPIController(servicer, asserter)
	MempoolApiController := server.NewMempoolAPIController(servicer, asserter)
	NetworkApiController := server.NewNetworkAPIController(servicer, asserter)

	router := server.NewRouter(AccountApiController, BlockApiController, ConstructionApiController, MempoolApiController, NetworkApiController)
	return router, nil
}
