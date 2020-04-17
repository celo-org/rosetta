package analyzer

import (
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
)

type SubAccountType string

const (
	AccMain                    SubAccountType = "Main"
	AccLockedGoldNonVoting     SubAccountType = "LockedGoldNonVoting"
	AccLockedGoldVotingActive  SubAccountType = "LockedGoldVotingActive"
	AccLockedGoldVotingPending SubAccountType = "LockedGoldVotingPending"
	AccLockedGoldPending       SubAccountType = "LockedGoldPending"
)

type SubAccount struct {
	Identifier SubAccountType
	Metadata   map[string]interface{}
}

type Account struct {
	Address    common.Address
	SubAccount SubAccount
}

type OperationType string

const (
	OpFee                OperationType = "fee"
	OpTransfer           OperationType = "transfer"
	OpLockGold           OperationType = "lockGold"
	OpUnlockGold         OperationType = "unlockGold"
	OpRelockGold         OperationType = "relockGold"
	OpWithdrawGold       OperationType = "withdrawGold"
	OpVote               OperationType = "vote"
	OpActiveVotes        OperationType = "activateVotes"
	OpRevokePendingVotes OperationType = "revokePendingVotes"
	OpRevokeActiveVotes  OperationType = "revokeActiveVotes"
	OpSlash              OperationType = "slash"
	OpEpochRewards       OperationType = "epochRewards"
)

func (ot OperationType) String() string { return string(ot) }

var AllOperationTypes = []OperationType{
	OpFee,
	OpTransfer,
	OpLockGold,
	OpUnlockGold,
	OpRelockGold,
	OpWithdrawGold,
	OpVote,
	OpActiveVotes,
	OpRevokePendingVotes,
	OpRevokeActiveVotes,
	OpSlash,
	OpEpochRewards,
}

func AllOperationTypesString() []string {
	strs := make([]string, len(AllOperationTypes))
	for i, opType := range AllOperationTypes {
		strs[i] = string(opType)
	}
	return strs
}

type BalanceChange struct {
	Account Account
	Amount  *big.Int
}

type Operation struct {
	Type       OperationType
	Changes    []BalanceChange
	Successful bool
}

// ---------------------------------------------------------------------------------------------------
// Account Factories
// ---------------------------------------------------------------------------------------------------

func NewAccount(addr common.Address, subAccount SubAccountType) Account {
	return Account{
		Address:    addr,
		SubAccount: SubAccount{Identifier: subAccount},
	}
}

func NewVotingAccount(addr common.Address, subAccount SubAccountType, group common.Address) Account {
	return Account{
		Address: addr,
		SubAccount: SubAccount{
			Identifier: subAccount,
			Metadata: map[string]interface{}{
				"group": group,
			},
		},
	}
}

// ---------------------------------------------------------------------------------------------------
// Operation Factories
// ---------------------------------------------------------------------------------------------------

func NewEpochRewards(changes map[common.Address]*big.Int) *Operation {
	return &Operation{
		Type:       OpEpochRewards,
		Changes:    mapToBalanceChanges(changes),
		Successful: true,
	}
}

func NewFee(changes map[common.Address]*big.Int) *Operation {
	return &Operation{
		Type:       OpFee,
		Changes:    mapToBalanceChanges(changes),
		Successful: true,
	}
}

func NewTransfer(from common.Address, to common.Address, value *big.Int, successful bool) *Operation {
	return &Operation{
		Type:       OpTransfer,
		Successful: successful,
		Changes: []BalanceChange{
			{Account: NewAccount(from, AccMain), Amount: negate(value)},
			{Account: NewAccount(to, AccMain), Amount: value},
		},
	}
}

func NewSlash(slashed common.Address, slasher common.Address, communityFund common.Address, lockedGoldAddr common.Address, penalty *big.Int, reward *big.Int) *Operation {
	communityFundReward := new(big.Int).Sub(penalty, reward)
	return &Operation{
		Type:       OpSlash,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(slashed, AccLockedGoldNonVoting), Amount: negate(penalty)},
			{Account: NewAccount(slasher, AccLockedGoldNonVoting), Amount: reward},
			{Account: NewAccount(communityFund, AccMain), Amount: communityFundReward},
			// The actual cGLD transfer
			{Account: NewAccount(lockedGoldAddr, AccMain), Amount: negate(communityFundReward)},
		},
	}
}

func NewLockGold(addr common.Address, lockedGoldAddr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpLockGold,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(addr, AccMain), Amount: negate(value)},
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
			// The actual cGLD transfer
			{Account: NewAccount(lockedGoldAddr, AccMain), Amount: value},
		},
	}
}

func NewWithdrawGold(addr common.Address, lockedGoldAddr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpWithdrawGold,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(addr, AccLockedGoldPending), Amount: negate(value)},
			{Account: NewAccount(addr, AccMain), Amount: value},
			// The actual cGLD transfer
			{Account: NewAccount(lockedGoldAddr, AccMain), Amount: negate(value)},
		},
	}
}

func NewUnlockGold(addr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpUnlockGold,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: negate(value)},
			{Account: NewAccount(addr, AccLockedGoldPending), Amount: value},
		},
	}
}

func NewRelockGold(addr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpRelockGold,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(addr, AccLockedGoldPending), Amount: negate(value)},
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
		},
	}
}

func NewVote(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpVote,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: negate(value)},
			{Account: NewVotingAccount(addr, AccLockedGoldVotingPending, group), Amount: big.NewInt(0)},
		},
	}
}

func NewActiveVotes(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpActiveVotes,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewVotingAccount(addr, AccLockedGoldVotingPending, group), Amount: big.NewInt(0)},
			{Account: NewVotingAccount(addr, AccLockedGoldVotingActive, group), Amount: big.NewInt(0)},
		},
	}
}

func NewRevokePendingVotes(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpRevokePendingVotes,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewVotingAccount(addr, AccLockedGoldVotingPending, group), Amount: big.NewInt(0)},
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
		},
	}
}

func NewRevokeActiveVotes(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpRevokeActiveVotes,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewVotingAccount(addr, AccLockedGoldVotingActive, group), Amount: big.NewInt(0)},
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
		},
	}
}

// ---------------------------------------------------------------------------------------------------
// Helpers
// ---------------------------------------------------------------------------------------------------

func FilterChangesBySubAccount(op *Operation, subAccountType SubAccountType) map[common.Address]*big.Int {
	changes := make(map[common.Address]*big.Int)
	for _, change := range op.Changes {
		if change.Account.SubAccount.Identifier == subAccountType {
			changes[change.Account.Address] = change.Amount
		}
	}
	return changes
}

func MatchChangesOnSubAccount(op1 *Operation, op2 *Operation, subAccountType SubAccountType) bool {
	op1Changes := FilterChangesBySubAccount(op1, subAccountType)
	op2Changes := FilterChangesBySubAccount(op2, subAccountType)
	return reflect.DeepEqual(op1Changes, op2Changes)
}

func negate(value *big.Int) *big.Int { return new(big.Int).Neg(value) }

func mapToBalanceChanges(changes map[common.Address]*big.Int) []BalanceChange {
	balanceChanges := make([]BalanceChange, 0, len(changes))
	for addr, amount := range changes {
		balanceChanges = append(balanceChanges, BalanceChange{
			Account: NewAccount(addr, AccMain),
			Amount:  amount,
		})
	}
	return balanceChanges
}
