package db

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrContractNotFound = errors.New("db: contract record not found")
	ErrFutureBlock      = errors.New("db: block number is greater than last persisted block")
)

type RosettaDBReader interface {
	// LastPersistedBlock will return the last block that was persisted
	// In case of not block, it will return 0
	LastPersistedBlock(ctx context.Context) (*big.Int, error)

	// GasPriceMinimumOn return the gasPriceMinimum registered for that block
	// In case of no value, will return with fallbackValue which is 0
	GasPriceMinimumFor(ctx context.Context, block *big.Int) (*big.Int, error)

	// RegistryAddressOn returns the address of the contract at the point in history (block, txIndex)
	// In case there's no record for that contract it will fail with ErrContractNotFound
	RegistryAddressOn(ctx context.Context, block *big.Int, txIndex uint, contractName string) (common.Address, error)

	// RegistryAddressesOn returns the address of the contracts at the point in history (block, txIndex)
	// For the case a contract is not yet deployed, that contract won't be in the result map
	RegistryAddressesOn(ctx context.Context, block *big.Int, txIndex uint, contractName ...string) (map[string]common.Address, error)
}

type RosettaDBWriter interface {
	ApplyChanges(ctx context.Context, changeSet *BlockChangeSet) error
}

type RosettaDB interface {
	RosettaDBReader
	RosettaDBWriter
}

type RegistryChange struct {
	TxIndex    uint
	Contract   string
	NewAddress common.Address
}

type BlockChangeSet struct {
	BlockNumber     *big.Int
	GasPriceMinimum *big.Int
	RegistryChanges []RegistryChange
}
