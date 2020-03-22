package admin

import (
	"context"

	"github.com/ethereum/go-ethereum/p2p"
	"github.com/ethereum/go-ethereum/rpc"
)

// AdminClient defines typed wrappers for the Debug RPC API.
type AdminClient struct {
	c *rpc.Client
}

// Dial connects a client to the given URL.
func Dial(rawurl string) (*AdminClient, error) {
	return DialContext(context.Background(), rawurl)
}

func DialContext(ctx context.Context, rawurl string) (*AdminClient, error) {
	c, err := rpc.DialContext(ctx, rawurl)
	if err != nil {
		return nil, err
	}
	return NewClient(c), nil
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *AdminClient {
	return &AdminClient{c}
}

func (ac *AdminClient) Close() {
	ac.c.Close()
}

func (ac *AdminClient) Peers(ctx context.Context) ([]p2p.PeerInfo, error) {
	var peerInfos []p2p.PeerInfo
	err := ac.c.CallContext(ctx, &peerInfos, "admin_peers")
	if err != nil {
		return []p2p.PeerInfo{}, err
	}
	return peerInfos, nil
}
