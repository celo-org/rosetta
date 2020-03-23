package admin

import (
	"context"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// AdminClient defines typed wrappers for the Admin RPC API.
type AdminClient struct {
	c *rpc.Client
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *AdminClient {
	return &AdminClient{c}
}

func (ac *AdminClient) Peers(ctx context.Context) (*[]p2p.PeerInfo, error) {
	var peerInfos []p2p.PeerInfo
	err := ac.c.CallContext(ctx, &peerInfos, "admin_peers")
	if err != nil {
		return nil, err
	}
	return &peerInfos, nil
}
