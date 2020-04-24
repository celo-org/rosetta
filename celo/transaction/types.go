package transaction

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CeloMethod string

// values taken from contract method names for ABI usage
var (
	CreateAccount      CeloMethod = "createAccount"
	LockGold           CeloMethod = "lock"
	UnlockGold         CeloMethod = "unlock"
	RelockGold         CeloMethod = "relock"
	WithdrawGold       CeloMethod = "withdraw"
	Vote               CeloMethod = "vote"
	ActivateVotes      CeloMethod = "activate"
	RevokePendingVotes CeloMethod = "revokePending"
	RevokeActiveVotes  CeloMethod = "revokeActive"
	// RevokeAllActiveVotes TransactionType = "revokeAllActiveVotes"
	// Slash                TransactionType = "slash"
	// EpochRewards         TransactionType = "epochRewards"
)

func (tt CeloMethod) String() string { return string(tt) }

var (
	CeloMethodToRegistryKey = map[*CeloMethod]*wrapper.RegistryKey{
		&CreateAccount:      &wrapper.AccountsRegistryId,
		&LockGold:           &wrapper.LockedGoldRegistryId,
		&UnlockGold:         &wrapper.LockedGoldRegistryId,
		&RelockGold:         &wrapper.LockedGoldRegistryId,
		&WithdrawGold:       &wrapper.LockedGoldRegistryId,
		&Vote:               &wrapper.ElectionRegistryId,
		&ActivateVotes:      &wrapper.ElectionRegistryId,
		&RevokePendingVotes: &wrapper.ElectionRegistryId,
		&RevokeActiveVotes:  &wrapper.ElectionRegistryId,
	}
)

type TransactionOptions struct {
	From   common.Address
	To     *common.Address // non-nil means exclusively cGLD transfer
	Value  *big.Int
	Method *CeloMethod // non-nil means celo registry contract invokation
	Args   []interface{}
}

//go:generate gencodec -type GenericMetadata -out gen_transaction_metadata_json.go

// [note]: non cGLD fee currencies currently unsupported
type GenericMetadata struct {
	From                common.Address
	Nonce               uint64          `json:"nonce"    gencodec:"required"`
	GasPrice            *big.Int        `json:"gasPrice" gencodec:"required"`
	GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient" rlp:"nil"` // nil means no gateway fee is paid
	GatewayFee          *big.Int        `json:"gatewayFee" rlp:"nil"`          // nil means no gateway fee is paid
}

//go:generate gencodec -type TransactionMetadata -out gen_transaction_metadata_json.go
type TransactionMetadata struct {
	Generic *GenericMetadata
	To      *common.Address
	Value   *big.Int
	Gas     uint64
	Data    []byte
}

func (tm *TransactionMetadata) AsTransaction() *types.Transaction {
	return types.NewTransaction(
		tm.Generic.Nonce,
		*tm.To,
		tm.Value,
		tm.Gas,
		tm.Generic.GasPrice,
		nil, // non-cGLD fees not supported
		tm.Generic.GatewayFeeRecipient,
		tm.Generic.GatewayFee,
		tm.Data,
	)
}
