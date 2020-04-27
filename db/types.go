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
	// LastPersistedBlock returns the last block that was persisted
	// In case of not block, it will return 0
	LastPersistedBlock(ctx context.Context) (*big.Int, error)

	// GasPriceMinimumFor returns the gasPriceMinimum registered for that block
	// In case of no value, will return with fallbackValue which is 0
	GasPriceMinimumFor(ctx context.Context, block *big.Int) (*big.Int, error)

	// TobinTaxFor returns the tobinTax applied to txs in that block.
	// More specifically, it returns the numerator of the tobin tax.
	// In case of no value, will return with fallbackValue which is 0
	TobinTaxFor(ctx context.Context, block *big.Int) (*big.Int, error)

	// RegistryAddressStartOf returns the address of the contract at the start of (block, tx)
	// In case there's no record for that contract it will fail with ErrContractNotFound
	RegistryAddressStartOf(ctx context.Context, block *big.Int, txIndex uint, contractName string) (common.Address, error)

	// RegistryAddressesStartOf returns the address of the contracts at the start of (block, tx)
	// For the case a contract is not yet deployed, that contract won't be in the result map
	RegistryAddressesStartOf(ctx context.Context, block *big.Int, txIndex uint, contractName ...string) (map[string]common.Address, error)

	// CarbonOffsetPartnerStartOf returns the address of the contract at the start of (block, tx)
	// In case of no value, will return with fallbackValue which is common.ZeroAddress
	CarbonOffsetPartnerStartOf(ctx context.Context, block *big.Int, txIndex uint) (common.Address, error)
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

type CarbonOffsetPartnerChange struct {
	TxIndex uint
	Address common.Address
}

type BlockChangeSet struct {
	BlockNumber               *big.Int
	GasPriceMinimum           *big.Int
	TobinTax                  *big.Int
	RegistryChanges           []RegistryChange
	CarbonOffsetPartnerChange CarbonOffsetPartnerChange
}
