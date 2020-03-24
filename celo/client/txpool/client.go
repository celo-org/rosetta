package txpool

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rpc"
)

// TxPoolClient defines typed wrappers for the TxPool RPC API.
type TxPoolClient struct {
	c *rpc.Client
}

// NewClient creates a client that uses the given RPC client.
func NewClient(c *rpc.Client) *TxPoolClient {
	return &TxPoolClient{c}
}

type TxPoolContent struct {
	Pending map[common.Address]types.Transactions
	Queued  map[common.Address]types.Transactions
}

func (tpc *TxPoolClient) Content(ctx context.Context) (*TxPoolContent, error) {
	var result TxPoolContent
	if err := tpc.c.CallContext(ctx, &result, "txpool_content"); err != nil {
		return nil, err
	}
	return &result, nil
}
