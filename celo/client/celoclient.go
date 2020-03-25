package client

import (
	"context"

	"github.com/celo-org/rosetta/celo/client/admin"
	"github.com/celo-org/rosetta/celo/client/debug"
	"github.com/celo-org/rosetta/celo/client/network"
	"github.com/celo-org/rosetta/celo/client/txpool"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type CeloClient struct {
	Rpc    *rpc.Client
	Eth    *ethclient.Client
	Debug  *debug.DebugClient
	Admin  *admin.AdminClient
	Net    *network.NetworkClient
	TxPool *txpool.TxPoolClient
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*CeloClient, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*CeloClient, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewCeloClient(c), nil
}

func NewCeloClient(rpcClient *rpc.Client) *CeloClient {
	return &CeloClient{
		Rpc:    rpcClient,
		Eth:    ethclient.NewClient(rpcClient),
		Debug:  debug.NewClient(rpcClient),
		Admin:  admin.NewClient(rpcClient),
		Net:    network.NewClient(rpcClient),
		TxPool: txpool.NewClient(rpcClient),
	}
}

func (c *CeloClient) Close() {
	c.Rpc.Close()
}
