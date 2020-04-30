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
