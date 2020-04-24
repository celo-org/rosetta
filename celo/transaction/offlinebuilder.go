package transaction

import (
	"fmt"

	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

type OfflineBuilder struct {
	abiCache map[*wrapper.RegistryKey]*abi.ABI
}

func NewOfflineBuilder() *OfflineBuilder {
	return &OfflineBuilder{
		abiCache: make(map[*wrapper.RegistryKey]*abi.ABI),
	}
}

func (b *OfflineBuilder) getAbi(key *wrapper.RegistryKey) (*abi.ABI, error) {
	var abi *abi.ABI
	abi, present := b.abiCache[key]

	if !present {
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
			return nil, err
		}
	}

	b.abiCache[key] = abi
	return abi, nil
}

func (b *OfflineBuilder) getData(options *TransactionOptions) ([]byte, error) {
	if options.Method == nil {
		return nil, fmt.Errorf("'Method' required for building tx data")
	}

	contractABI, err := b.getAbi(CeloMethodToRegistryKey[options.Method])
	if err != nil {
		return nil, err
	}

	packedData, err := contractABI.Pack(options.Method.String(), options.Args...)
	if err != nil {
		return nil, err
	}

	return packedData, nil
}
