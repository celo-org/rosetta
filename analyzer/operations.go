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

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/kliento/client/debug"
)

type SubAccountType string

const (
	AccMain                        SubAccountType = "Main"
	AccSigner                      SubAccountType = "AccountsAuthorizedSigner"
	AccLockedGoldNonVoting         SubAccountType = "LockedGoldNonVoting"
	AccLockedGoldVotingActive      SubAccountType = "LockedGoldVotingActive"
	AccLockedGoldVotingPending     SubAccountType = "LockedGoldVotingPending"
	AccLockedGoldPending           SubAccountType = "LockedGoldPending"
	AccReleaseGoldVested           SubAccountType = "ReleaseGoldVested"
	AccReleaseGoldUnvestedLocked   SubAccountType = "ReleaseGoldUnvestedLocked"
	AccReleaseGoldUnvestedUnLocked SubAccountType = "ReleaseGoldUnvestedUnlocked"
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
	OpFee                        OperationType = "fee"
	OpTransfer                   OperationType = "transfer"
	OpCreateAccount              OperationType = "createAccount"
	OpAuthorizeVoteSigner        OperationType = "authorizeVoteSigner"
	OpAuthorizeValidatorSigner   OperationType = "authorizeValidatorSigner"
	OpAuthorizeAttestationSigner OperationType = "authorizeAttestationSigner"
	OpLockGold                   OperationType = "lockGold"
	OpUnlockGold                 OperationType = "unlockGold"
	OpRelockGold                 OperationType = "relockGold"
	OpWithdrawGold               OperationType = "withdrawGold"
	OpVote                       OperationType = "vote"
	OpActiveVotes                OperationType = "activateVotes"
	OpRevokePendingVotes         OperationType = "revokePendingVotes"
	OpRevokeActiveVotes          OperationType = "revokeActiveVotes"
	OpSlash                      OperationType = "slash"
	OpEpochRewards               OperationType = "epochRewards"
)

func (ot OperationType) String() string { return string(ot) }

func (ot OperationType) requiresTransfer() bool {
	return ot == OpLockGold || ot == OpWithdrawGold || ot == OpSlash
}

var AllOperationTypes = []OperationType{
	OpFee,
	OpTransfer,
	OpCreateAccount,
	OpAuthorizeValidatorSigner,
	OpAuthorizeAttestationSigner,
	OpAuthorizeVoteSigner,
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

func NewSignerAccount(addr common.Address, main common.Address) Account {
	return Account{
		Address: addr,
		SubAccount: SubAccount{
			Identifier: AccSigner,
			Metadata: map[string]interface{}{
				"account": main,
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

func NewTransfer(from, to common.Address, value *big.Int, successful bool) *Operation {
	return &Operation{
		Type:       OpTransfer,
		Successful: successful,
		Changes:    getTransferChanges(from, to, value),
	}
}

func NewCreateAccount(from common.Address) *Operation {
	return &Operation{
		Type:       OpCreateAccount,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewAccount(from, AccMain), Amount: nil},
		},
	}
}

func NewAuthorizeSigner(from common.Address, signer common.Address, authorizeOp OperationType) *Operation {
	return &Operation{
		Type:       authorizeOp,
		Successful: true,
		Changes: []BalanceChange{
			{Account: NewSignerAccount(signer, from)},
		},
	}
}

// Ex. Slash(penalty=110, reward=100 CELO)
// Transfer Operation:
//
//	lockedGoldContractAccMain        -10
//	communityFundAccMain              10
//
// LockedGold Operation (created from `AccountSlashed(slashed, -100, reporter, 110)` event):
//
//	slashedAccLockedGoldNonVoting    -110
//	reporterAccLockedGoldNonVoting    100
//	lockedGoldContractAccMain        -10
//	communityFundAccMain              10
func NewSlash(slashed, slasher, communityFund, lockedGoldAddr common.Address, penalty, reward *big.Int) *Operation {
	communityFundReward := new(big.Int).Sub(penalty, reward)
	return &Operation{
		Type:       OpSlash,
		Successful: true,
		Changes: append([]BalanceChange{
			{Account: NewAccount(slashed, AccLockedGoldNonVoting), Amount: negate(penalty)},
			{Account: NewAccount(slasher, AccLockedGoldNonVoting), Amount: reward},
		}, getTransferChanges(lockedGoldAddr, communityFund, communityFundReward)...),
	}
}

// Ex. lock(100 CELO)
// Transfer Operation:
//
//	fromAccMain         	 	-100
//	lockedGoldContractAccMain    100
//
// LockedGold Operation (created from `GoldLocked(fromAccount, 100)` event):
//
//	fromAccMain                 -100
//	lockedGoldContractAccMain    100
//	fromAccLockedNonVoting       100
//
// Reconciled Operation:
//
//	fromAccMain         	 	-100
//	lockedGoldContractAccMain    100
//	fromAccLockedNonVoting       100
func NewLockGold(addr, lockedGoldAddr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpLockGold,
		Successful: true,
		Changes: append(
			getTransferChanges(addr, lockedGoldAddr, value),
			BalanceChange{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
		),
	}
}

// Ex. Withdraw(100 CELO)
// Transfer Operation:
//
//	lockedGoldContractAccMain   -100
//	toAccMain         	 	     100
//
// LockedGold Operation (created from `GoldWithdraw(toAccount, 100)` event):
//
//	lockedGoldContractAccMain   -100
//	toAccMain         	 	     100
//	AccLockedPending            -100
func NewWithdrawGold(addr, lockedGoldAddr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpWithdrawGold,
		Successful: true,
		Changes: append(
			getTransferChanges(lockedGoldAddr, addr, value),
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
			if val, ok := changes[change.Account.Address]; ok {
				changes[change.Account.Address] = new(big.Int).Add(val, change.Amount)
			} else {
				changes[change.Account.Address] = change.Amount
			}
		}
	}
	return changes
}

func MatchChangesOnSubAccount(op1 *Operation, op2 *Operation, subAccountType SubAccountType) bool {
	op1Changes := FilterChangesBySubAccount(op1, subAccountType)
	op2Changes := FilterChangesBySubAccount(op2, subAccountType)
	return reflect.DeepEqual(op1Changes, op2Changes)
}

func getTransferChanges(from, to common.Address, value *big.Int) []BalanceChange {
	return []BalanceChange{
		{Account: NewAccount(from, AccMain), Amount: negate(value)},
		{Account: NewAccount(to, AccMain), Amount: value},
	}
}

// ReconcileLogOpsWithTransfers merges operations derived from the tx logs with transfers
// derived from the tx tracer into one slice that contains no duplicate operations.
//
// See NewLockGold, NewWithdrawal, NewSlash for examples.
func ReconcileLogOpsWithTransfers(logOps, transferOps []Operation) ([]Operation, error) {
	ops := make([]Operation, 0, len(transferOps)+len(logOps))

	findMatchAndReconcile := func(transferOps []Operation, logOp *Operation, i int) (int, error) {
		for ; i < len(transferOps); i++ {
			if MatchChangesOnSubAccount(logOp, &transferOps[i], AccMain) {
				return i, nil
			}
		}
		return i, fmt.Errorf("BUG: Can't find matching transfer for log op")
	}

	// We assume both arrays are in order and remove rendundant transfer
	// operations as follows:
	// 		for each logOp
	// 			if logOp should have a matching transferOp
	//				append transferOps until we find one matching logOp
	//              point logOp to reconciled operation
	// 				skip matching transfer op
	// 			append logOp
	// 		append remaining transferOps
	i := 0
	for _, logOp := range logOps {
		if logOp.Type.requiresTransfer() {
			// TODO - revisit
			// nolint:gosec
			matchIndex, err := findMatchAndReconcile(transferOps, &logOp, i)
			if err != nil {
				return nil, err
			}
			ops = append(ops, transferOps[i:matchIndex]...)
			i = matchIndex + 1 // skip the matching transfer operation
		}
		ops = append(ops, logOp)
	}
	// add the rest of the transfers
	if i < len(transferOps) {
		ops = append(ops, transferOps[i:]...)
	}

	return ops, nil
}

func InternalTransfersToOperations(transfers []debug.Transfer) []Operation {
	transferOps := make([]Operation, len(transfers))
	for i, t := range transfers {
		transferOps[i] = *NewTransfer(t.From, t.To, t.Value, t.Status.String() == debug.TransferStatusSuccess.String())
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
