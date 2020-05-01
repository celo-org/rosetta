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

package airgap

import (
	"context"
	"fmt"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/common"
)

type airgGapServerImpl struct {
	cc               *client.CeloClient
	supportedMethods map[*CeloMethod]airGapServerMethod
}

func NewAirGapServer(cc *client.CeloClient) (AirGapServer, error) {
	server := &airgGapServerImpl{
		cc:               cc,
		supportedMethods: make(map[*CeloMethod]airGapServerMethod),
	}

	registry, err := wrapper.NewRegistry(cc)
	if err != nil {
		return nil, err
	}

	groupFactories := []airGapMethodGroupFactory{
		newAccountsAirGapMethods,
		newAccountsAirGapMethods,
		newLockedGoldAirGapMethods,
	}
	for _, factory := range groupFactories {
		group, err := factory(registry)
		if err != nil {
			return nil, err
		}
		group.register(server)
	}

	return server, nil
}

func (b *airgGapServerImpl) registerMethod(method *CeloMethod, executor airGapServerMethod) {
	b.supportedMethods[method] = executor
}

func (b *airgGapServerImpl) SubmitTx(ctx context.Context, rawTx []byte) (*common.Hash, error) {
	return b.cc.Eth.SendRawTransaction(ctx, rawTx)
}

func (b *airgGapServerImpl) ObtainMetadata(ctx context.Context, options *TxArgs) (*TxMetadata, error) {
	nonce, err := b.cc.Eth.NonceAt(ctx, options.From, nil) // nil == latest
	if err != nil {
		return nil, err
	}

	gasPrice, err := b.cc.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	txMetadata := TxMetadata{
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

		data, to, err := serverMethod(ctx, options.Args)
		if err != nil {
			return nil, err
		}
		txMetadata.To = to
		txMetadata.Data = data
	}

	estimatedGas, err := b.cc.Eth.EstimateGas(ctx, txMetadata.AsCallMessage())
	if err != nil {
		return nil, err
	}
	txMetadata.Gas = estimatedGas

	return &txMetadata, nil
}

// func (b *airgGapServerImpl) resolveMethod(ctx context.Context, method *CeloMethod, args []interface{}) ([]byte, common.Address, error) {
// 	resolvedArgs, err := b.resolvers.FindResolverFor(method)(ctx, args)
// 	if err != nil {
// 		return nil, common.ZeroAddress, err
// 	}

// 	data, err := b.abiBuilder.GetData(method, resolvedArgs)
// 	if err != nil {
// 		return nil, common.ZeroAddress, err
// 	}

// 	contractAddress, err := b.registry.GetAddressForString(ctx, nil, method.Contract.String())
// 	if err != nil {
// 		return nil, common.ZeroAddress, err
// 	}

// 	return data, contractAddress, err
// }
