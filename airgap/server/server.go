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
	"fmt"

	"github.com/celo-org/rosetta/airgap"
	"github.com/ethereum/go-ethereum/common"
)

// airGapServerMethod is a function that returns the tx.data and tx.to for that method + parameters
type airGapServerMethod func(context.Context, []interface{}) ([]byte, common.Address, error)

type airgGapServerImpl struct {
	srvCtx           ServerContext
	supportedMethods map[*airgap.CeloMethod]airGapServerMethod
}

func NewAirgapServer(srvCtx ServerContext) (airgap.Server, error) {
	supportedMethods, err := hydrateMethods(srvCtx)
	if err != nil {
		return nil, err
	}

	return &airgGapServerImpl{
		srvCtx:           srvCtx,
		supportedMethods: supportedMethods,
	}, nil
}

func (b *airgGapServerImpl) SubmitTx(ctx context.Context, rawTx []byte) (*common.Hash, error) {
	return b.srvCtx.SendRawTransaction(ctx, rawTx)
}

func (b *airgGapServerImpl) ObtainMetadata(ctx context.Context, options *airgap.TxArgs) (*airgap.TxMetadata, error) {
	nonce, err := b.srvCtx.NonceAt(ctx, options.From, nil) // nil == latest
	if err != nil {
		return nil, err
	}

	gasPrice, err := b.srvCtx.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	txMetadata := airgap.TxMetadata{
		From:                options.From,
		Nonce:               nonce,
		GasPrice:            gasPrice,
		GatewayFee:          nil,
		GatewayFeeRecipient: nil,
		Value:               options.Value,
	}

	if options.To != nil && options.Method != nil {
		return nil, fmt.Errorf("Can't specify 'to' &  'method' at the same time (%s,%s)", options.To, options.Method)
	} else if options.To != nil {
		txMetadata.To = *options.To
	} else {
		serverMethod, ok := b.supportedMethods[options.Method]
		if !ok {
			return nil, fmt.Errorf("Unsupported method: %s", options.Method)
		}

		hydratedArgs, err := options.Method.DeserializeArguments(options.Args...)
		if err != nil {
			return nil, err
		}

		data, to, err := serverMethod(ctx, hydratedArgs)
		if err != nil {
			return nil, err
		}
		txMetadata.To = to
		txMetadata.Data = data
	}

	estimatedGas, err := b.srvCtx.EstimateGas(ctx, txMetadata.AsCallMessage())
	if err != nil {
		return nil, err
	}
	txMetadata.Gas = estimatedGas

	return &txMetadata, nil
}
