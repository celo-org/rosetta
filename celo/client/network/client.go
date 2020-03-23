package admin

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/rpc"
)

// NetworkClient defines typed wrappers for the Network RPC API.
type NetworkClient struct {
	c *rpc.Client
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *NetworkClient {
	return &NetworkClient{c}
}

func (nc *NetworkClient) NetworkId(ctx context.Context) (*big.Int, error) {
	version := new(big.Int)
	var ver string
	if err := nc.c.CallContext(ctx, &ver, "net_version"); err != nil {
		return nil, err
	}
	if _, ok := version.SetString(ver, 10); !ok {
		return nil, fmt.Errorf("invalid net_version result %q", ver)
	}
	return version, nil
}
