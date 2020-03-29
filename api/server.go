package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
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

type RosettaServer struct {
	cc          *client.CeloClient
	cfg         *RosettaServerConfig
	chainParams *celo.ChainParameters
	server      *http.Server
}

func NewRosettaServer(cc *client.CeloClient, cfg *RosettaServerConfig, chainParams *celo.ChainParameters) *RosettaServer {
	var mainHandler http.Handler
	mainHandler = createRouter(cc, chainParams)
	mainHandler = requestLogHandler(mainHandler)
	mainHandler = http.TimeoutHandler(mainHandler, cfg.RequestTimeout, "Request Timed out")

	server := &http.Server{
		Addr:         cfg.ListenAddress(),
		Handler:      mainHandler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// TODO set ErrorLog: ,
	}

	return &RosettaServer{
		cc:          cc,
		cfg:         cfg,
		server:      server,
		chainParams: chainParams,
	}
}

func (rs *RosettaServer) Stop() {
	if err := rs.server.Close(); err != nil {
		log.Error("Error stoping httpServer", "err", err)
	}
}

func (rs *RosettaServer) Start() {
	log.Info("Starting httpServer", "listen_address", rs.cfg.ListenAddress())

	if err := rs.server.ListenAndServe(); err != nil {
		log.Crit("Error starting httpServer", "err", err)
	}
}

func requestLogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, w, r)
		log.Info("RequestServed", "method", r.Method, "uri", r.URL, "code", m.Code, "duration", m.Duration, "bytes", m.Written)
	})
}

func createRouter(celoClient *client.CeloClient, chainParams *celo.ChainParameters) *mux.Router {
	AccountApiService := NewAccountApiService(celoClient, chainParams)
	AccountApiController := NewAccountApiController(AccountApiService)

	BlockApiService := NewBlockApiService(celoClient, chainParams)
	BlockApiController := NewBlockApiController(BlockApiService)

	ConstructionApiService := NewConstructionApiService(celoClient, chainParams)
	ConstructionApiController := NewConstructionApiController(ConstructionApiService)

	MempoolApiService := NewMempoolApiService(celoClient, chainParams)
	MempoolApiController := NewMempoolApiController(MempoolApiService)

	NetworkApiService := NewNetworkApiService(celoClient, chainParams)
	NetworkApiController := NewNetworkApiController(NetworkApiService)

	router := NewRouter(AccountApiController, BlockApiController, ConstructionApiController, MempoolApiController, NetworkApiController)
	return router
}

// BaddRequest replies to the request with status 400 and the error message
// It does not otherwise end the request; the caller should ensure no further
func BadRequest(w http.ResponseWriter, err error) {
	payload := BuildErrorResponse(int32(http.StatusBadRequest), err)
	JSONResponse(payload, http.StatusBadRequest, w)
}
