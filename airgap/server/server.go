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
	"log"
	"math/big"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/kliento/registry"
	"github.com/celo-org/rosetta/airgap"
)

// airGapServerMethod is a function that returns the tx.data for that method + parameters
type airGapServerMethod func(context.Context, []interface{}) ([]byte, error)

// airGapServerEvent is a function that returns the topic[0] for that event name + parameters
type airGapServerEvent func(context.Context, [][]common.Hash) [][]common.Hash

type airGapServerImpl struct {
	srvCtx             ServerContext
	chainId            *big.Int
	transactionMethods map[*airgap.CeloMethod]airGapServerMethod
	callMethods        map[*airgap.CeloMethod]airGapServerMethod
	callEvents         map[*airgap.CeloEvent]airGapServerEvent
}

func NewAirgapServer(chainId *big.Int, srvCtx ServerContext) (airgap.Server, error) {
	transactionMethods, err := hydrateMethods(srvCtx, serverTransactionMethodDefinitions)
	if err != nil {
		return nil, err
	}

	callMethods, err := hydrateMethods(srvCtx, serverCallMethodDefinitions)
	if err != nil {
		return nil, err
	}

	callEvents, err := hydrateEvents(serverCallEventDefinitions)
	if err != nil {
		return nil, err
	}

	return &airGapServerImpl{
		srvCtx,
		chainId,
		transactionMethods,
		callMethods,
		callEvents,
	}, nil
}

func (b *airGapServerImpl) SubmitTx(ctx context.Context, rawTx []byte) (*common.Hash, error) {
	tx := new(types.Transaction)
	if err := tx.UnmarshalBinary(rawTx); err != nil {
		return nil, err
	}

	hash := tx.Hash()
	return &hash, b.srvCtx.SendTransaction(ctx, tx)
}

func (b *airGapServerImpl) CallData(ctx context.Context, options *airgap.CallParams) ([]byte, error) {
	if options.Method == nil {
		return nil, fmt.Errorf("'Method' must be provided as options")
	}

	serverMethod, ok := b.callMethods[options.Method]
	if !ok {
		return nil, fmt.Errorf("Unsupported method: %s", options.Method)
	}

	hydratedArgs, err := options.Method.DeserializeArguments(options.Args...)
	if err != nil {
		return nil, err
	}

	data, err := serverMethod(ctx, hydratedArgs)
	if err != nil {
		return nil, err
	}

	var to common.Address
	if options.To != nil {
		to = *options.To
	} else {
		to, err = b.srvCtx.addressFor(ctx, registry.ContractID(options.Method.Contract), options.BlockNumber)
		if err != nil {
			return nil, fmt.Errorf("'Contract' not a valid registry ID")
		}
	}

	return b.srvCtx.CallContract(ctx, ethereum.CallMsg{
		From: options.From,
		To:   &to,
		Data: data,
		Gas:  0,
	}, options.BlockNumber)
}

func (b *airGapServerImpl) ObtainMetadata(ctx context.Context, options *airgap.TxArgs) (*airgap.TxMetadata, error) {
	nonce, err := b.srvCtx.NonceAt(ctx, options.From, nil) // nil == latest
	if err != nil {
		return nil, err
	}

	gasPrice, err := b.srvCtx.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	txMetadata := airgap.TxMetadata{
		From:     options.From,
		Nonce:    nonce,
		GasPrice: gasPrice,
		Value:    options.Value,
		ChainId:  b.chainId,
	}

	if options.To != nil {
		txMetadata.To = *options.To
	} else {
		if options.Method == nil {
			return nil, fmt.Errorf("'To' or 'Method' must be provided as options")
		}
	}

	if options.Method != nil {
		if options.To == nil { // 'to' is implicit from registry
			addr, err := b.srvCtx.addressFor(ctx, registry.ContractID(options.Method.Contract), nil)
			if err != nil {
				return nil, fmt.Errorf("'To' not provided and 'Contract' not a valid registry ID")
			}
			txMetadata.To = addr
		}

		var data []byte
		if serverMethod, ok := b.transactionMethods[options.Method]; ok {
			// This method is registered; use the registry to parse arguments
			log.Printf("Building metadata for celo method %s", options.Method.String())
			hydratedArgs, err := options.Method.DeserializeArguments(options.Args...)
			if err != nil {
				return nil, err
			}
			data, err = serverMethod(ctx, hydratedArgs)
			if err != nil {
				return nil, fmt.Errorf("Airgap server method failed: %s", err)
			}
		} else {
			// This method is unregistered; calculate the call data based on whether the args are already encoded
			methodSig := options.Method.Name
			log.Printf("Building metadata for contract call method %s", methodSig)

			var methodArgs interface{}
			if options.ArgsEncoded {
				// options.Args is a string array and the first element is the encoded args
				if len(options.Args) > 0 {
					if encodedArgsData, ok := options.Args[0].(string); ok {
						methodArgs = encodedArgsData
					} else {
						return nil, fmt.Errorf("Invalid encoded args: %s for method: %s", options.Args[0], options.Method)
					}
				} else {
					return nil, fmt.Errorf("Args not provided for method: %s", options.Method)
				}
			} else {
				// options.Args is a []interface{} and can be processed by ConstructContractCallDataGeneric as-is
				methodArgs = options.Args
			}

			data, err = constructContractCallDataGeneric(methodSig, methodArgs)
			if err != nil {
				return nil, fmt.Errorf("Failed to construct call data for method: %s using args: %v due to %s", options.Method, options.Args, err)
			}
		}
		txMetadata.Data = data
	}

	estimatedGas, err := b.srvCtx.EstimateGas(ctx, txMetadata.AsCallMessage())
	if err != nil {
		result, callErr := b.srvCtx.CallContract(ctx, txMetadata.AsCallMessage(), nil)
		if callErr != nil {
			return nil, fmt.Errorf("Gas estimation failed: %s \n Call for revert reason failed: %s", err, callErr)
		}
		return nil, fmt.Errorf("Gas estimation failed: %s \n Revert reason: %s", err, result)
	}
	txMetadata.Gas = estimatedGas

	return &txMetadata, nil
}

func (b *airGapServerImpl) FilterQuery(ctx context.Context, options *airgap.FilterQueryParams) ([]types.Log, error) {
	if options.Event == nil {
		return nil, fmt.Errorf("'Event' must be provided as options")
	}

	serverEvent, ok := b.callEvents[options.Event]
	if !ok {
		return nil, fmt.Errorf("Unsupported event: %v", options.Event)
	}

	hydratedTopics, err := options.Event.DeserializeTopics(options.Topics...)
	if err != nil {
		return nil, err
	}

	topics := serverEvent(ctx, hydratedTopics)

	address, err := b.srvCtx.addressFor(ctx, registry.ContractID(options.Event.Contract), options.FromBlock)
	if err != nil {
		return nil, fmt.Errorf("'%s' not a valid registry ID", options.Event.Contract)
	}

	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
		FromBlock: options.FromBlock,
		ToBlock:   options.ToBlock,
		Topics:    topics,
	}
	return b.srvCtx.FilterLogs(ctx, query)
}
