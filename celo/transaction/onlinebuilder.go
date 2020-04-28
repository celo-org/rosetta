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

package transaction

import (
	"context"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type OnlineBuilder struct {
	celoClient      *client.CeloClient
	registry        *wrapper.RegistryWrapper
	resolverLocator *ResolverLocator
	*OfflineBuilder
}

func NewOnlineBuilder(client *client.CeloClient) (*OnlineBuilder, error) {
	registry, err := wrapper.NewRegistry(client)
	if err != nil {
		return nil, err
	}

	locator, err := NewResolverLocator(registry, client)
	if err != nil {
		return nil, err
	}

	builder := &OnlineBuilder{
		registry:        registry,
		celoClient:      client,
		resolverLocator: locator,
		OfflineBuilder:  NewOfflineBuilder(),
	}

	return builder, nil
}

func (b *OnlineBuilder) FetchNodeMetadata(ctx context.Context) (*NodeMetadata, error) {
	gasPrice, err := b.celoClient.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	metadata := NodeMetadata{
		GatewayFeeRecipient: nil,
		GatewayFee:          nil,
		GasPrice:            gasPrice,
	}
	return &metadata, nil
}

func (b *OnlineBuilder) FetchAccountMetadata(ctx context.Context, address common.Address) (*AccountMetadata, error) {
	nonce, err := b.celoClient.Eth.NonceAt(ctx, address, nil) // nil == latest
	if err != nil {
		return nil, err
	}

	metadata := AccountMetadata{
		From:  address,
		Nonce: nonce,
	}
	return &metadata, nil
}

func (b *OnlineBuilder) FetchMethodMetadata(ctx context.Context, method *CeloMethod, args []interface{}) (*MethodMetadata, error) {
	argResolver := b.resolverLocator.GetResolver(method)
	resolvedArgs, err := argResolver(args, ctx)
	if err != nil {
		return nil, err
	}

	data, err := b.getData(method, resolvedArgs)
	if err != nil {
		return nil, err
	}

	contractAddress, err := b.registry.GetAddressForString(&bind.CallOpts{Context: ctx}, CeloMethodToRegistryKey[method].String())
	if err != nil {
		return nil, err
	}

	meta := MethodMetadata{
		Data: data,
		To:   &contractAddress,
	}
	return &meta, nil
}

func (b *OnlineBuilder) FetchTransactionMetadata(ctx context.Context, options *TransactionOptions) (*TransactionMetadata, error) {
	accountMetadata, err := b.FetchAccountMetadata(ctx, options.From)
	if err != nil {
		return nil, err
	}

	nodeMetadata, err := b.FetchNodeMetadata(ctx)
	if err != nil {
		return nil, err
	}

	methodMetadata := &MethodMetadata{
		Data: nil,
		To:   options.To,
	}

	if options.Method != nil {
		meta, err := b.FetchMethodMetadata(ctx, options.Method, options.Args)
		if err != nil {
			return nil, err
		}
		methodMetadata = meta
	}

	msg := ethereum.CallMsg{
		From:                accountMetadata.From,
		GatewayFee:          nodeMetadata.GatewayFee,
		GatewayFeeRecipient: nodeMetadata.GatewayFeeRecipient,
		GasPrice:            nodeMetadata.GasPrice,
		To:                  methodMetadata.To,
		Data:                methodMetadata.Data,
		Value:               options.Value,
		FeeCurrency:         nil, // only cGLD fees currently supported
	}

	estimatedGas, err := b.celoClient.Eth.EstimateGas(ctx, msg)
	if err != nil {
		return nil, err
	}

	txMetadata := TransactionMetadata{
		AccountMetadata: accountMetadata,
		NodeMetadata:    nodeMetadata,
		MethodMetadata:  methodMetadata,
		Value:           options.Value,
		Gas:             estimatedGas,
	}

	return &txMetadata, nil
}
