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

package wrapper

import (
	"context"
	"errors"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

type RegistryWrapper struct {
	cc       *client.CeloClient
	contract *contract.Registry
}

type RegistryKey string

var (
	LockedGoldRegistryId RegistryKey = "LockedGold"
	ElectionRegistryId   RegistryKey = "Election"
	AccountsRegistryId   RegistryKey = "Accounts"
)

func (rk RegistryKey) String() string { return string(rk) }

var (
	ErrRegistryNotDeployed = errors.New("Registry Not Deployed")
)

func NewRegistry(cc *client.CeloClient) (*RegistryWrapper, error) {
	registry, err := contract.NewRegistry(params.RegistrySmartContractAddress, cc.Eth)
	err = client.WrapRpcError(err)
	return &RegistryWrapper{
		cc:       cc,
		contract: registry,
	}, err
}

func (w *RegistryWrapper) Contract() *contract.Registry {
	return w.contract
}

// GetAddressFor is a free data retrieval call binding the contract method 0xdd927233.
//
// Solidity: function getAddressFor(bytes32 identifierHash) constant returns(address)
func (w *RegistryWrapper) GetAddressFor(opts *bind.CallOpts, identifierHash [32]byte) (common.Address, error) {
	address, err := w.contract.GetAddressFor(opts, identifierHash)
	err = client.WrapRpcError(err)

	if err != nil {
		return common.ZeroAddress, err
	} else if err == client.ErrContractNotDeployed {
		return common.ZeroAddress, ErrRegistryNotDeployed
	}

	if address == common.ZeroAddress {
		return common.ZeroAddress, client.ErrContractNotDeployed
	}

	return address, nil
}

func (w *RegistryWrapper) GetAddressForString(ctx context.Context, block *big.Int, identifier string) (common.Address, error) {
	address, err := w.contract.GetAddressForString(&bind.CallOpts{
		BlockNumber: block,
		Context:     ctx,
	}, identifier)
	err = client.WrapRpcError(err)

	if err == client.ErrContractNotDeployed {
		return common.ZeroAddress, ErrRegistryNotDeployed
	}

	if err != nil {
		return common.ZeroAddress, err
	}

	if address == common.ZeroAddress {
		return common.ZeroAddress, client.ErrContractNotDeployed
	}

	return address, nil
}

func (w *RegistryWrapper) GetLockedGold(ctx context.Context, block *big.Int) (*contract.LockedGold, error) {
	addr, err := w.GetAddressForString(ctx, block, LockedGoldRegistryId.String())
	if err != nil {
		return nil, err
	}

	lockedGold, err := contract.NewLockedGold(addr, w.cc.Eth)
	if err != nil {
		return nil, err
	}

	return lockedGold, nil
}

func (w *RegistryWrapper) GetLockedGoldWrapper(ctx context.Context, block *big.Int) (*LockedGoldWrapper, error) {
	contract, err := w.GetLockedGold(ctx, block)
	if err != nil {
		return nil, err
	}
	return NewLockedGold(contract), nil
}

func (w *RegistryWrapper) GetAccounts(ctx context.Context, block *big.Int) (*contract.Accounts, error) {
	addr, err := w.GetAddressForString(ctx, block, AccountsRegistryId.String())
	if err != nil {
		return nil, err
	}

	accounts, err := contract.NewAccounts(addr, w.cc.Eth)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (w *RegistryWrapper) GetAccountsWrapper(ctx context.Context, block *big.Int) (*AccountsWrapper, error) {
	contract, err := w.GetAccounts(ctx, block)
	if err != nil {
		return nil, err
	}
	return NewAccounts(contract), nil
}

func (w *RegistryWrapper) GetElection(ctx context.Context, block *big.Int) (*contract.Election, error) {
	addr, err := w.GetAddressForString(ctx, block, ElectionRegistryId.String())
	if err != nil {
		return nil, err
	}

	election, err := contract.NewElection(addr, w.cc.Eth)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (w *RegistryWrapper) GetElectionWrapper(ctx context.Context, block *big.Int) (*ElectionWrapper, error) {
	contract, err := w.GetElection(ctx, block)
	if err != nil {
		return nil, err
	}
	return NewElection(contract), nil
}
