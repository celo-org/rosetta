package server

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ServerContext interface {
	addressFor(ctx context.Context, identifier wrapper.RegistryKey) (common.Address, error)

	authorizeMetadata(ctx context.Context, popSignature []byte) (*wrapper.EncodedSignature, error)
	voteMetadata(ctx context.Context, group common.Address, value *big.Int) (*wrapper.AddressLesserGreater, error)
	revokeMetadata(ctx context.Context, account common.Address, group common.Address, value *big.Int) (*wrapper.RevokeMetadata, error)

	// From EthClient

	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)
	SendRawTransaction(ctx context.Context, data []byte) (*common.Hash, error)
}

type serverContextImpl struct {
	*ethclient.Client

	registry *wrapper.RegistryWrapper
}

func NewServerContext(cc *client.CeloClient) (ServerContext, error) {
	registry, err := wrapper.NewRegistry(cc)
	if err != nil {
		return nil, err
	}

	return &serverContextImpl{
		Client:   cc.Eth,
		registry: registry,
	}, nil
}

func (sc *serverContextImpl) suggestGasPrice(ctx context.Context, identifier wrapper.RegistryKey) (common.Address, error) {
	return sc.registry.GetAddressForString(ctx, nil, identifier.String())
}

func (sc *serverContextImpl) addressFor(ctx context.Context, identifier wrapper.RegistryKey) (common.Address, error) {
	return sc.registry.GetAddressForString(ctx, nil, identifier.String())
}

func (sc *serverContextImpl) authorizeMetadata(ctx context.Context, popSignature []byte) (*wrapper.EncodedSignature, error) {
	accounts, err := sc.registry.GetAccountsWrapper(ctx, nil)
	if err != nil {
		return nil, err
	}
	return accounts.AuthorizeMetadata(popSignature)
}

func (sc *serverContextImpl) voteMetadata(ctx context.Context, group common.Address, value *big.Int) (*wrapper.AddressLesserGreater, error) {
	elections, err := sc.registry.GetElectionWrapper(ctx, nil)
	if err != nil {
		return nil, err
	}
	return elections.VoteMetadata(&bind.CallOpts{Context: ctx}, group, value)
}

func (sc *serverContextImpl) revokeMetadata(ctx context.Context, account common.Address, group common.Address, value *big.Int) (*wrapper.RevokeMetadata, error) {
	elections, err := sc.registry.GetElectionWrapper(ctx, nil)
	if err != nil {
		return nil, err
	}
	return elections.RevokeMetadata(&bind.CallOpts{Context: ctx}, account, group, value)
}
