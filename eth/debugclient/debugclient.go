package debugclient

import (
	"context"
	"io/ioutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/eth"
	"github.com/ethereum/go-ethereum/rpc"
)

// Client defines typed wrappers for the Debug RPC API.
type Client struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*Client, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*Client, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *Client {
	return &Client{c}
}

func (ec *Client) Close() {
	ec.c.Close()
}

// TraceTransactions performs a tracer over a transaction. Can use a custom tracer or default one
// result type depends on the tracer, and it's the caller reponsability to use the proper one
func (ec *Client) TraceTransaction(ctx context.Context, result interface{}, txhash common.Hash, traceConfig *eth.TraceConfig) error {
	return ec.c.CallContext(ctx, result, "debug_traceTransaction", txhash, traceConfig)
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
