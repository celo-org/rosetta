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

package client

import (
	"errors"
	"strings"

	"github.com/celo-org/celo-blockchain/rpc"
)

var (
	ErrMissingTrieNode     = errors.New("Block too Old: Missing Trie Node")
	ErrContractNotDeployed = errors.New("Contract Not Deployed")
)

func RpcErrorCode(err error) (int, bool) {
	rpcError, ok := err.(rpc.Error)
	if ok {
		return rpcError.ErrorCode(), true
	}
	return 0, false
}

func WrapRpcError(err error) error {
	if err == nil {
		return nil
	}

	if strings.Contains(err.Error(), "No Implementation set") {
		return ErrContractNotDeployed
	}

	code, isRpc := RpcErrorCode(err)
	if !isRpc {
		return err
	}

	if code == -32000 && strings.HasPrefix(err.Error(), "missing trie node") {
		return ErrMissingTrieNode
	}

	return err
}
