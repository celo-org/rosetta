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

package server

import (
	"context"
	"math/big"

	"github.com/celo-org/celo-blockchain/common"
)

func preprocessAuthorizeSigner(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error) {
	encodedSig, err := srvCtx.authorizeMetadata(ctx, args[1].([]byte))
	if err != nil {
		return nil, err
	}

	return []interface{}{args[0], encodedSig.V, encodedSig.R, encodedSig.S}, nil
}

func preprocessVote(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error) {
	keys, err := srvCtx.voteMetadata(ctx, args[0].(common.Address), args[1].(*big.Int))
	if err != nil {
		return nil, err
	}

	return []interface{}{args[0], args[1], keys.Lesser, keys.Greater}, nil
}

func preprocessRevoke(ctx context.Context, srvCtx ServerContext, args []interface{}) ([]interface{}, error) {
	keys, err := srvCtx.revokeMetadata(ctx, args[0].(common.Address), args[1].(common.Address), args[2].(*big.Int))
	if err != nil {
		return nil, err
	}

	return []interface{}{args[1], args[2], keys.Lesser, keys.Greater, keys.Index}, nil
}
