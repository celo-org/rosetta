package debug

import (
	"context"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/rpc"
)

// DebugClient defines typed wrappers for the Debug RPC API.
type DebugClient struct {
	c *rpc.Client
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *DebugClient {
	return &DebugClient{c}
}

// TraceTransactions performs a tracer over a transaction. Can use a custom tracer or default one
// result type depends on the tracer, and it's the caller reponsability to use the proper one
func (dc *DebugClient) TraceTransaction(ctx context.Context, result interface{}, txhash common.Hash, traceConfig *eth.TraceConfig) error {
	return dc.c.CallContext(ctx, result, "debug_traceTransaction", txhash, traceConfig)
}

// ReadTracer reads a tracer file from the filesystem
func ReadTracer(path string) (*string, error) {
	tracerBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	tracerString := string(tracerBytes)
	return &tracerString, nil
}
