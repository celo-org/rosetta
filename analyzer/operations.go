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

package analyzer

import (
	"fmt"
	"math/big"
	"reflect"

	"github.com/celo-org/rosetta/celo/client/debug"
	"github.com/celo-org/rosetta/internal/utils"
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

func (ot OperationType) requiresTransfer() bool {
	return ot == OpLockGold || ot == OpWithdrawGold || ot == OpSlash
}

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

func NewTransfer(from, to common.Address, value *big.Int, tobinTax *TobinTax, successful bool) *Operation {
	return &Operation{
		Type:       OpTransfer,
		Successful: successful,
		Changes:    GetTransferChangesWithTobinTax(from, to, value, tobinTax),
	}
}

func NewSlash(slashed, slasher, communityFund, lockedGoldAddr common.Address, penalty, reward *big.Int, tobinTax *TobinTax) *Operation {
	communityFundReward := new(big.Int).Sub(penalty, reward)
	return &Operation{
		Type:       OpSlash,
		Successful: true,
		Changes: append([]BalanceChange{
			{Account: NewAccount(slashed, AccLockedGoldNonVoting), Amount: negate(penalty)},
			{Account: NewAccount(slasher, AccLockedGoldNonVoting), Amount: reward},
		}, GetTransferChangesWithTobinTax(communityFund, lockedGoldAddr, communityFundReward, tobinTax)...),
	}
}

func NewLockGold(addr, lockedGoldAddr common.Address, value *big.Int, tobinTax *TobinTax) *Operation {
	return &Operation{
		Type:       OpLockGold,
		Successful: true,
		Changes: append(
			GetTransferChangesWithTobinTax(addr, lockedGoldAddr, value, tobinTax),
			BalanceChange{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: tobinTax.Apply(value).AfterTaxAmount},
		),
	}
}

func NewWithdrawGold(addr, lockedGoldAddr common.Address, value *big.Int, tobinTax *TobinTax) *Operation {
	return &Operation{
		Type:       OpWithdrawGold,
		Successful: true,
		Changes: append(
			GetTransferChangesWithTobinTax(lockedGoldAddr, addr, value, tobinTax),
			BalanceChange{Account: NewAccount(addr, AccLockedGoldPending), Amount: negate(value)},
		),
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
			{Account: NewVotingAccount(addr, AccLockedGoldVotingPending, group), Amount: nil},
		},
	}
}

func NewActiveVotes(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpActiveVotes,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewVotingAccount(addr, AccLockedGoldVotingPending, group), Amount: nil},
			{Account: NewVotingAccount(addr, AccLockedGoldVotingActive, group), Amount: nil},
		},
	}
}

func NewRevokePendingVotes(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpRevokePendingVotes,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewVotingAccount(addr, AccLockedGoldVotingPending, group), Amount: nil},
			{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
		},
	}
}

func NewRevokeActiveVotes(addr common.Address, group common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpRevokeActiveVotes,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewVotingAccount(addr, AccLockedGoldVotingActive, group), Amount: nil},
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

func GetTransferChangesWithTobinTax(from, to common.Address, value *big.Int, tobinTax *TobinTax) []BalanceChange {
	res := tobinTax.Apply(value)

	changes := GetTransferChanges(from, to, res.AfterTaxAmount)

	if res.TaxAmount.Cmp(utils.Big0) > 0 {
		changes = append(changes, GetTransferChanges(from, tobinTax.Recipient, res.TaxAmount)...)
	}

	return changes
}

func GetTransferChanges(from, to common.Address, value *big.Int) []BalanceChange {
	return []BalanceChange{
		{Account: NewAccount(from, AccMain), Amount: negate(value)},
		{Account: NewAccount(to, AccMain), Amount: value},
	}
}

// ReconcileLockedGoldTransfers merges transfer and lockedGold operations
// together into one slice that contains no duplicate operations.
//
// Ex. lock(100 cGlD), tobinTax == 10%
// Transfer Operation
// 	fromAccMain         	 	-90
// 	lockedGoldContractAccMain    90
// 	fromAccMain			    	-10
// 	tobinRecipientAccMain        10
// LockedGold Operation (created from `GoldLocked(fromAccount, 90)` event):
// 	fromAccMain                 -90
// 	lockedGoldContractAccMain    90
// 	fromAccMain		     		-10
// 	tobinRecipientAccMain        10
//  fromAccLockedNonVoting       90
// Since these Operations match on all AccMain Balance transfers, we can discard
// the Transfer Operation in favor of the more detailed LockedGold Operation.
func ReconcileLockedGoldTransfers(lockedGoldOps, transferOps []Operation) ([]Operation, error) {
	ops := make([]Operation, 0, len(transferOps)+len(lockedGoldOps))

	findMatch := func(transferOps []Operation, lgOp *Operation, i int) (int, error) {
		for ; i < len(transferOps); i++ {
			if MatchChangesOnSubAccount(lgOp, &transferOps[i], AccMain) {
				return i, nil
			}
		}
		return i, fmt.Errorf("BUG: Can't find matching transfer for LockedGold op")
	}

	// We assume both arrays are in order and remove rendundant transfer
	// operations as follows:
	// 		for each lgOp
	// 			if lgOp should have a matching transferOp
	//				append transferOps until we find one matching lgOp
	// 				skip matching transfer op
	// 			append lgOp
	// 		append remaining transferOps
	i := 0
	for _, lgOp := range lockedGoldOps {
		if lgOp.Type.requiresTransfer() {
			matchIndex, err := findMatch(transferOps, &lgOp, i)
			if err != nil {
				return nil, err
			}
			ops = append(ops, transferOps[i:matchIndex]...)
			i = matchIndex + 1 // skip the matching transfer operation
		}
		ops = append(ops, lgOp)
	}
	// add the rest of the transfers
	if i < len(transferOps) {
		ops = append(ops, transferOps[i:]...)
	}

	return ops, nil
}

func InternalTransfersToOperations(transfers []debug.Transfer, tobinTax *TobinTax) []Operation {
	transferOps := make([]Operation, len(transfers))
	for i, t := range transfers {
		transferOps[i] = *NewTransfer(t.From, t.To, t.Value, tobinTax, t.Status.String() == debug.TransferStatusSuccess.String())
	}
	return transferOps
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
