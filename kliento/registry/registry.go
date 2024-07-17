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

package registry

import (
	"context"
	"math/big"
	"strings"

	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	blockchainErrors "github.com/celo-org/celo-blockchain/contracts"
	"github.com/celo-org/celo-blockchain/contracts/config"
	"github.com/celo-org/celo-blockchain/core/types"

	"github.com/celo-org/rosetta/kliento/client"
	"github.com/celo-org/rosetta/kliento/contracts"
)

// ContractID represents the ID of a contract according to the Registry
type ContractID string

func (cid ContractID) String() string { return string(cid) }

// RegistryAddress is the address of the registry which is the same on any celo network
var RegistryAddress = config.RegistrySmartContractAddress

// Registry defines an interface to access all Celo core Contracts
type Registry interface {
	AllAddresses(ctx context.Context, blockNumber *big.Int) ([]common.Address, error)
	GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error)
	generatedRegistry
	TryParseLog(ctx context.Context, eventLog types.Log, blockNumber *big.Int) (*RegistryParsedLog, error)
	EnableCaching(ctx context.Context, blockNumber *big.Int) error
}

type RegistryParsedLog struct {
	Contract string
	Event    string
	Log      interface{}
}

type addressContainer struct {
	address common.Address
	dirty   bool
}

type registryCache struct {
	idToAddress map[string]*addressContainer
	addressToID map[common.Address]string
}

type registryImpl struct {
	cc               *client.CeloClient
	RegistryContract *contracts.Registry
	cache            *registryCache
	proxyAbi         *abi.ABI
	boundContracts
}

// New creates a new contracts Registry
func New(cc *client.CeloClient) (Registry, error) {
	registry, err := contracts.NewRegistry(RegistryAddress, cc.Eth)
	err = client.WrapRpcError(err)
	if err != nil {
		return nil, err
	}

	return &registryImpl{
		cc:               cc,
		RegistryContract: registry,
		cache:            nil,
	}, err
}

func (r *registryImpl) GetRegistryContract() *contracts.Registry {
	return r.RegistryContract
}

func (r *registryImpl) AllAddresses(ctx context.Context, blockNumber *big.Int) ([]common.Address, error) {
	addresses := make([]common.Address, len(RegisteredContractIDs))
	for idx, id := range RegisteredContractIDs {
		_address, err := r.GetAddressFor(ctx, blockNumber, id)
		if err != nil {
			return nil, err
		}
		addresses[idx] = _address
	}
	return addresses, nil
}

func (r *registryImpl) GetAddressFor(ctx context.Context, blockNumber *big.Int, contractID ContractID) (common.Address, error) {
	identifier := contractID.String()
	// fetch from cache and mark clean if successful
	ac := r.cache.getAddressContainer(identifier)
	if ac != nil {
		ac.dirty = false
		return ac.address, nil
	}

	// if uncached, fetch from contract state
	address, err := r.RegistryContract.GetAddressForString(&bind.CallOpts{BlockNumber: blockNumber, Context: ctx}, identifier)
	err = client.WrapRpcError(err)
	if err != nil {
		return common.ZeroAddress, err
	} else if address == common.ZeroAddress {
		return address, client.ErrContractNotDeployed
	}

	// update cache
	r.cache.put(identifier, address)
	return address, nil
}

// IsExpectedBeforeContractsDeployed checks for expected errors when contracts are not deployed yet
func IsExpectedBeforeContractsDeployed(err error) bool {
	return isErrSubset(err, blockchainErrors.ErrRegistryContractNotDeployed) || isErrSubset(err, blockchainErrors.ErrSmartContractNotDeployed) || isErrSubset(err, client.ErrContractNotDeployed)
}

func isErrSubset(err error, suberr error) bool {
	return strings.Contains(err.Error(), suberr.Error())
}

type logParser interface {
	TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error)
}

func (r *registryImpl) tryParseProxyLog(log types.Log) (eventName string, event interface{}, err error) {
	proxy, err := contracts.NewProxyFilterer(log.Address, r.cc.Eth)
	if err != nil {
		return "", nil, err
	}
	eventName, event, ok, err := proxy.TryParseLog(log)
	if !ok || err != nil {
		return "", nil, err
	}
	return eventName, event, nil
}

// TryParseLog parses an event log using a fully hydrated registry
func (r *registryImpl) TryParseLog(ctx context.Context, log types.Log, blockNumber *big.Int) (*RegistryParsedLog, error) {
	var eventName, id string
	var event interface{}
	var parseError error

	// determine contract identifier and parse log if address is known
	if log.Address == config.RegistrySmartContractAddress {
		id = "Registry"
		eventName, event, _, parseError = r.RegistryContract.TryParseLog(log)

		// update cache if registry mapping was changed
		switch v := event.(type) {
		case *contracts.RegistryRegistryUpdated:
			r.cache.put(v.Identifier, v.Addr)
		}
	} else if id = r.cache.getIdentifier(log.Address); id != "" {
		contract, err := r.GetContractByID(ctx, id, blockNumber)
		if err != nil {
			return nil, err
		}
		eventName, event, _, parseError = contract.(logParser).TryParseLog(log)
	}

	// if address is known but parsing failed, try using proxy ABI
	if parseError != nil {
		eventName, event, parseError = r.tryParseProxyLog(log)
		if parseError != nil {
			return nil, parseError
		}
		id += "Proxy"
	}

	if event != nil {
		return &RegistryParsedLog{
			Contract: id,
			Event:    eventName,
			Log:      event,
		}, nil
	}

	return nil, nil
}

// Enables address caching
func (r *registryImpl) EnableCaching(ctx context.Context, blockNumber *big.Int) error {
	r.cache = &registryCache{
		idToAddress: make(map[string]*addressContainer, len(RegisteredContractIDs)),
		addressToID: make(map[common.Address]string, len(RegisteredContractIDs)),
	}

	// Hydrate initially at provided block
	for _, id := range RegisteredContractIDs {
		_, err := r.GetContractByID(ctx, id.String(), blockNumber)
		if err != nil && !IsExpectedBeforeContractsDeployed(err) {
			return err
		}
	}
	return nil
}

func (rc *registryCache) put(identifier string, address common.Address) {
	if rc != nil {
		ac := rc.getAddressContainer(identifier)
		if ac != nil {
			ac.address = address
			rc.addressToID[address] = identifier
			ac.dirty = true
		} else {
			rc.idToAddress[identifier] = &addressContainer{address, true}
			rc.addressToID[address] = identifier
		}
	}
}

func (rc *registryCache) getAddressContainer(identifier string) *addressContainer {
	if rc != nil {
		if addressContainer, ok := rc.idToAddress[identifier]; ok {
			return addressContainer
		}
	}
	return nil
}

func (rc *registryCache) getIdentifier(address common.Address) string {
	if rc != nil {
		if identifier, ok := rc.addressToID[address]; ok {
			return identifier
		}
	}
	return ""
}

// used in generated registry for contract binding reinitialization
func (rc *registryCache) isDirty(identifier string) bool {
	ac := rc.getAddressContainer(identifier)
	if ac != nil {
		return ac.dirty
	}
	return false
}
