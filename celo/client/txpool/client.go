// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package txpool

import (
	"context"

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

type TxNonceMap map[string]*rpc.RPCTransaction
type TxAccountMap map[string]TxNonceMap
type TxPoolContent map[string]TxAccountMap

/*
{
  pending: {
    0x0216d5032f356960cd3749c31ab34eeff21b3395: {
      806: RPCTransaction,
    },
    0x24d407e5a0b506e1cb2fae163100b5de01f5193c: {
      34: RPCTransaction,
    }
  },
  queued: {
    0x976a3fc5d6f7d259ebfb4cc2ae75115475e9867c: {
      3: RPCTransaction,
    },
    0x9b11bf0459b0c4b2f87f8cebca4cfc26f294b63a: {
      2: RPCTransaction,
      6: RPCTransaction,
    }
  }
}
*/
func (tpc *TxPoolClient) Content(ctx context.Context) (*TxPoolContent, error) {
	var result TxPoolContent
	if err := tpc.c.CallContext(ctx, &result, "txpool_content"); err != nil {
		return nil, err
	}
	return &result, nil
}
