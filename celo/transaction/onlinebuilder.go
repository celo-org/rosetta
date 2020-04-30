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
	"errors"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

type OnlineBuilder struct {
	celoClient *client.CeloClient
	registry   *wrapper.RegistryWrapper
	*OfflineBuilder
}

func NewOnlineBuilder(client *client.CeloClient) (*OnlineBuilder, error) {
	registry, err := wrapper.NewRegistry(client)
	if err != nil {
		return nil, err
	}

	builder := &OnlineBuilder{
		registry:       registry,
		celoClient:     client,
		OfflineBuilder: NewOfflineBuilder(),
	}

	return builder, nil
}

func (b *OnlineBuilder) fetchGenericMetadata(ctx context.Context, address common.Address) (*GenericMetadata, error) {
	nonce, err := b.celoClient.Eth.NonceAt(ctx, address, nil) // nil == latest
	if err != nil {
		return nil, err
	}

	gasPrice, err := b.celoClient.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	metadata := GenericMetadata{
		From:                address,
		Nonce:               nonce,
		GatewayFeeRecipient: nil,
		GatewayFee:          nil,
		GasPrice:            gasPrice,
	}
	return &metadata, nil
}

func (b *OnlineBuilder) FetchTransactionMetadata(ctx context.Context, options *TransactionOptions) (*TransactionMetadata, error) {
	generic, err := b.fetchGenericMetadata(ctx, options.From)
	if err != nil {
		return nil, err
	}

	var data []byte
	var to *common.Address
	if options.To != nil && options.Method == nil {
		data = nil
		to = options.To
	} else if options.Method != nil {
		data, err = b.getData(options)
		if err != nil {
			return nil, err
		}

		contractAddress, err := b.registry.GetAddressForString(nil, CeloMethodToRegistryKey[options.Method].String())
		if err != nil {
			return nil, err
		}
		to = &contractAddress
	} else {
		return nil, errors.New("Options 'To' and 'Method' required and mutually exclusive")
	}

	msg := ethereum.CallMsg{
		From:                generic.From,
		GatewayFee:          generic.GatewayFee,
		GatewayFeeRecipient: generic.GatewayFeeRecipient,
		GasPrice:            generic.GasPrice,
		FeeCurrency:         nil, // only cGLD fees currently supported
		Value:               options.Value,
		To:                  to,
		Data:                data,
	}

	estimatedGas, err := b.celoClient.Eth.EstimateGas(ctx, msg)
	if err != nil {
		return nil, err
	}

	txMetadata := TransactionMetadata{
		Generic: generic,
		To:      to,
		Value:   options.Value,
		Data:    data,
		Gas:     estimatedGas,
	}

	return &txMetadata, nil
}
