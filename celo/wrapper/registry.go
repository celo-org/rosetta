package wrapper

import (
	"context"
	"errors"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

type RegistryWrapper struct {
	contract *contract.Registry
}

var (
	ErrRegistryNotDeployed = errors.New("Registry Not Deployed")
)

func NewRegistry(celoClient *client.CeloClient) (*RegistryWrapper, error) {
	registry, err := contract.NewRegistry(params.RegistrySmartContractAddress, celoClient.Eth)
	err = client.WrapRpcError(err)
	if err != nil {
		return nil, err
	}

	return &RegistryWrapper{
		contract: registry,
	}, nil
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

func (w *RegistryWrapper) GetAddressForString(opts *bind.CallOpts, identifier string) (common.Address, error) {
	address, err := w.contract.GetAddressForString(opts, identifier)
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

func (w *RegistryWrapper) GetLockedGold(opts *bind.CallOpts, backend bind.ContractBackend) (*contract.LockedGold, error) {
	addr, err := w.GetAddressForString(opts, "LockedGold")
	if err != nil {
		return nil, err
	}

	lockedGold, err := contract.NewLockedGold(addr, backend)
	if err != nil {
		return nil, err
	}

	return lockedGold, nil
}

func (w *RegistryWrapper) GetElection(opts *bind.CallOpts, backend bind.ContractBackend) (*contract.Election, error) {
	addr, err := w.GetAddressForString(opts, "Election")
	if err != nil {
		return nil, err
	}

	election, err := contract.NewElection(addr, backend)
	if err != nil {
		return nil, err
	}

	return election, nil
}

func (w *RegistryWrapper) GetUpdatesOnBlock(ctx context.Context, blockNumber uint64, maxTxIndex *uint, identifiers [][32]byte) (map[[32]byte]common.Address, error) {
	addresses := make(map[[32]byte]common.Address)

	// Get Iterator for events on given block
	iter, err := w.contract.FilterRegistryUpdated(&bind.FilterOpts{
		Start:   blockNumber,
		End:     &blockNumber,
		Context: ctx,
	}, identifiers)

	if err != nil {
		return nil, err
	}

	for iter.Next() {
		if maxTxIndex != nil && iter.Event.Raw.TxIndex >= *maxTxIndex {
			break
		}
		addresses[iter.Event.IdentifierHash] = iter.Event.Addr
	}

	err = iter.Close()
	if err != nil {
		return addresses, err
	}

	return addresses, nil
}
