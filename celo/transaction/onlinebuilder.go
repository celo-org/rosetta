package transaction

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type TransactionBuilder struct {
	celoClient *client.CeloClient
	registry   *wrapper.RegistryWrapper
	abiCache   map[*wrapper.RegistryKey]*abi.ABI
}

func NewTransactionBuilder(client *client.CeloClient) (*TransactionBuilder, error) {
	registry, err := wrapper.NewRegistry(client)
	if err != nil {
		return nil, err
	}

	builder := &TransactionBuilder{
		registry:   registry,
		celoClient: client,
		abiCache:   make(map[*wrapper.RegistryKey]*abi.ABI),
	}

	return builder, nil
}

func (b *TransactionBuilder) populateCache(key *wrapper.RegistryKey) error {
	_, present := b.abiCache[key]
	if present {
		return nil
	}

	var abi *abi.ABI
	var err error

	switch key {
	case &wrapper.AccountsRegistryId:
		abi, err = contract.ParseAccountsABI()
	case &wrapper.ElectionRegistryId:
		abi, err = contract.ParseElectionABI()
	case &wrapper.LockedGoldRegistryId:
		abi, err = contract.ParseLockedGoldABI()
	default:
		err = fmt.Errorf("Unimplemented ABI parsing for %s", key)
	}

	if err != nil {
		return err
	}

	b.abiCache[key] = abi
	return nil
}

func (b *TransactionBuilder) fetchGenericMetadata(ctx context.Context, address common.Address) (*GenericMetadata, error) {
	nonce, err := b.celoClient.Eth.NonceAt(ctx, address, nil) // nil == latest
	if err != nil {
		return nil, err
	}

	gatewayFeeRecipient, err := b.celoClient.Eth.Coinbase(ctx)
	if err != nil {
		return nil, err
	}

	// TODO: consider fetching from node
	gatewayFee := big.NewInt(0)

	gasPrice, err := b.celoClient.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	metadata := GenericMetadata{
		From:                address,
		Nonce:               nonce,
		GatewayFeeRecipient: gatewayFeeRecipient,
		GatewayFee:          gatewayFee,
		GasPrice:            gasPrice,
	}
	return &metadata, nil
}

func (b *TransactionBuilder) FetchTransactionMetadata(ctx context.Context, options *TransactionOptions) (*TransactionMetadata, error) {
	generic, err := b.fetchGenericMetadata(ctx, options.From)
	if err != nil {
		return nil, err
	}

	var to *common.Address
	var data []byte
	if options.To != nil && options.Method == nil {
		to = options.To
		data = nil
	} else if options.Method != nil {
		registryKey, methodPresent := CeloMethodToContract[options.Method]
		if !methodPresent {
			return nil, errors.New("Celo Method has no corresponding registry contract key")
		}

		contractAddress, err := b.registry.GetAddressForString(nil, registryKey.String())
		if err != nil {
			return nil, err
		}

		contractABI, abiPresent := b.abiCache[registryKey]
		if !abiPresent {
			err := b.populateCache(registryKey)
			if err != nil {
				return nil, err
			}
			contractABI = b.abiCache[registryKey]
		}

		packedData, err := contractABI.Pack(options.Method.String(), options.Args...)
		if err != nil {
			return nil, err
		}

		to = &contractAddress
		data = packedData
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
