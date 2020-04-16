package rpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/celo-org/rosetta/api"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/service"
	"github.com/ethereum/go-ethereum/log"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
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

func NewRosettaServer(cc *client.CeloClient, db db.RosettaDBReader, cfg *RosettaServerConfig, chainParams *celo.ChainParameters) *rosettaServer {
	var mainHandler http.Handler
	mainHandler = createRouter(cc, db, chainParams)
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
	}
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

func createRouter(celoClient *client.CeloClient, db db.RosettaDBReader, chainParams *celo.ChainParameters) *mux.Router {
	AccountApiService := api.NewAccountApiService(celoClient, chainParams)
	AccountApiController := api.NewAccountApiController(AccountApiService)

	BlockApiService := api.NewBlockApiService(celoClient, db, chainParams)
	BlockApiController := api.NewBlockApiController(BlockApiService)

	ConstructionApiService := api.NewConstructionApiService(celoClient, chainParams)
	ConstructionApiController := api.NewConstructionApiController(ConstructionApiService)

	MempoolApiService := api.NewMempoolApiService(celoClient, chainParams)
	MempoolApiController := api.NewMempoolApiController(MempoolApiService)

	NetworkApiService := api.NewNetworkApiService(celoClient, chainParams)
	NetworkApiController := api.NewNetworkApiController(NetworkApiService)

	router := api.NewRouter(AccountApiController, BlockApiController, ConstructionApiController, MempoolApiController, NetworkApiController)
	return router
}
