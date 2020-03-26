package api

import (
	"net/http"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/internal/config"
	"github.com/ethereum/go-ethereum/log"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
)

func requestLogHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(handler, w, r)
		log.Info("RequestServed", "method", r.Method, "uri", r.URL, "code", m.Code, "duration", m.Duration, "bytes", m.Written)
	})
}

func addMiddlewares(handler http.Handler) http.Handler {
	handler = requestLogHandler(handler)
	handler = http.TimeoutHandler(handler, config.HttpServer.RequestTimeout, "Request Timed out")
	return handler
}

func StartHttpServer(celoClient *client.CeloClient) {
	var mainHandler http.Handler
	mainHandler = CreateRouter(celoClient)
	mainHandler = addMiddlewares(mainHandler)

	log.Info("Starting server", "listen_address", config.HttpServer.ListenAddress())

	if err := http.ListenAndServe(config.HttpServer.ListenAddress(), mainHandler); err != nil {
		log.Crit("Failed to start server", "err", err)
	}
}

func CreateRouter(celoClient *client.CeloClient) *mux.Router {
	AccountApiService := NewAccountApiService(celoClient)
	AccountApiController := NewAccountApiController(AccountApiService)

	BlockApiService := NewBlockApiService(celoClient)
	BlockApiController := NewBlockApiController(BlockApiService)

	ConstructionApiService := NewConstructionApiService(celoClient)
	ConstructionApiController := NewConstructionApiController(ConstructionApiService)

	MempoolApiService := NewMempoolApiService(celoClient)
	MempoolApiController := NewMempoolApiController(MempoolApiService)

	NetworkApiService := NewNetworkApiService(celoClient)
	NetworkApiController := NewNetworkApiController(NetworkApiService)

	router := NewRouter(AccountApiController, BlockApiController, ConstructionApiController, MempoolApiController, NetworkApiController)
	return router
}

// BaddRequest replies to the request with status 400 and the error message
// It does not otherwise end the request; the caller should ensure no further
func BadRequest(w http.ResponseWriter, err error) {
	payload := BuildErrorResponse(int32(http.StatusBadRequest), err)
	if err = EncodeJSONResponse(payload, http.StatusBadRequest, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
