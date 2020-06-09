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

	"github.com/celo-org/kliento/client/debug"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/ethereum/go-ethereum/common"
)

type SubAccountType string

const (
	AccMain                    SubAccountType = "Main"
	AccSigner                  SubAccountType = "AccountsAuthorizedSigner"
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
		Changes:    getTransferChangesWithTobinTax(from, to, value, tobinTax),
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
			{Account: NewAccount(from, AccMain), Amount: nil},
			{Account: NewAccount(signer, AccSigner), Amount: nil},
		},
	}
}

// Ex. Slash(penalty=110, reward=100 cGlD), tobinTax == 10%
// Transfer Operation:
// 	lockedGoldContractAccMain        -9
// 	communityFundAccMain              9
// 	lockedGoldContractAccMain        -1
// 	tobinRecipientAccMain             1
// LockedGold Operation (created from `AccountSlashed(slashed, -100, reporter, 110)` event):
// 	slashedAccLockedGoldNonVoting    -110
// 	reporterAccLockedGoldNonVoting    100
// 	lockedGoldContractAccMain        -9
// 	communityFundAccMain              9
// 	lockedGoldContractAccMain        -1
// 	tobinRecipientAccMain             1
func NewSlash(slashed, slasher, communityFund, lockedGoldAddr common.Address, penalty, reward *big.Int, tobinTax *TobinTax) *Operation {
	communityFundReward := new(big.Int).Sub(penalty, reward)
	return &Operation{
		Type:       OpSlash,
		Successful: true,
		Changes: append([]BalanceChange{
			{Account: NewAccount(slashed, AccLockedGoldNonVoting), Amount: negate(penalty)},
			{Account: NewAccount(slasher, AccLockedGoldNonVoting), Amount: reward},
		}, getTransferChangesWithTobinTax(lockedGoldAddr, communityFund, communityFundReward, tobinTax)...),
	}
}

// Ex. lock(100 cGlD), tobinTax == 10%
// Transfer Operation:
// 	fromAccMain         	 	-90
// 	lockedGoldContractAccMain    90
// 	fromAccMain			    	-10
// 	tobinRecipientAccMain        10
// LockedGold Operation (created from `GoldLocked(fromAccount, 90)` event):
// 	fromAccMain                 -90
// 	lockedGoldContractAccMain    90
//  fromAccLockedNonVoting       90
// Reconciled Operation:
// 	fromAccMain         	 	-90
// 	lockedGoldContractAccMain    90
// 	fromAccMain			    	-10
// 	tobinRecipientAccMain        10
//  fromAccLockedNonVoting       90
func NewLockGold(addr, lockedGoldAddr common.Address, value *big.Int) *Operation {
	return &Operation{
		Type:       OpLockGold,
		Successful: true,
		Changes: append(
			// We don't apply tobinTax here, since it's already applied
			getTransferChanges(addr, lockedGoldAddr, value),
			BalanceChange{Account: NewAccount(addr, AccLockedGoldNonVoting), Amount: value},
		),
	}
}

// Ex. Withdraw(100 cGlD), tobinTax == 10%
// Transfer Operation:
// 	lockedGoldContractAccMain   -90
// 	toAccMain         	 	     90
// 	lockedGoldContractAccMain   -10
// 	tobinRecipientAccMain        10
// LockedGold Operation (created from `GoldWithdraw(toAccount, 100)` event):
// 	lockedGoldContractAccMain   -90
// 	toAccMain         	 	     90
// 	lockedGoldContractAccMain   -10
// 	tobinRecipientAccMain        10
//  AccLockedPending            -100
func NewWithdrawGold(addr, lockedGoldAddr common.Address, value *big.Int, tobinTax *TobinTax) *Operation {
	return &Operation{
		Type:       OpWithdrawGold,
		Successful: true,
		Changes: append(
			getTransferChangesWithTobinTax(lockedGoldAddr, addr, value, tobinTax),
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

func getTransferChangesWithTobinTax(from, to common.Address, value *big.Int, tobinTax *TobinTax) []BalanceChange {
	taxAmount, afterTaxAmount := tobinTax.Apply(value)

	changes := getTransferChanges(from, to, afterTaxAmount)

	if taxAmount.Cmp(utils.Big0) > 0 {
		changes = append(changes, getTransferChanges(from, tobinTax.Recipient, taxAmount)...)
	}

	return changes
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
func ReconcileLogOpsWithTransfers(logOps, transferOps []Operation, tobinTax *TobinTax, lockedGold common.Address) ([]Operation, error) {
	ops := make([]Operation, 0, len(transferOps)+len(logOps))

	findMatchAndReconcile := func(transferOps []Operation, logOp *Operation, i int) (int, error) {
		for ; i < len(transferOps); i++ {
			if MatchAndReconcileLogOpWithTransfer(logOp, &transferOps[i], tobinTax, lockedGold) {
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

// MatchAndReconcileLogOpWithTransfer returns true if the operations match as is or
// when we factor in tobin tax
//
// If the operations match with tobin tax, then we reconcile the operations into one
// by combining their balance changes
//
// operations match with tobin tax iff
// 		logOp.Type == OpLockGold &&
// 		logOp.lockedGoldContractAccMain.diff == trOp.lockedGoldContractAccMain.diff &&
// 		logOp.fromAccMain.diff - trOp.fromAccMain.diff == trOp.tobinRecipientAccMain.diff
//
// Ex. lock(100 cGlD), tobinTax == 10%
// Transfer Operation:
// 	fromAccMain         	 	-90
// 	lockedGoldContractAccMain    90
// 	fromAccMain			    	-10 // diff == -100
// 	tobinRecipientAccMain        10
// LockedGold Operation (created from `GoldLocked(fromAccount, 90)` event):
// 	fromAccMain                 -90 // diff == -90
// 	lockedGoldContractAccMain    90
//  fromAccLockedNonVoting       90
// Reconciled Operation:
// 	fromAccMain         	 	-90
// 	lockedGoldContractAccMain    90
// 	fromAccMain			    	-10
// 	tobinRecipientAccMain        10
//  fromAccLockedNonVoting       90
func MatchAndReconcileLogOpWithTransfer(logOp, trOp *Operation, tobinTax *TobinTax, lockedGold common.Address) bool {

	// If the operations match as is then we can just return true
	// without checking if they match with tobin tax. Because tobin tax is applied
	// when we derive OpWithdraw and OpSlash, these should be matched here.
	if MatchChangesOnSubAccount(logOp, trOp, AccMain) {
		return true
	}

	// If the operations don't match as is and tobin tax is not defined, then
	// the operations don't match
	if !tobinTax.IsDefined() {
		return false
	}

	// Check if the operations match when we factor in the tobin tax.

	// NOTE: At this point the only type of operation we are checking for is OpLockGold.
	if logOp.Type != OpLockGold {
		return false
	}

	lgAccMainChanges := FilterChangesBySubAccount(logOp, AccMain)
	trAccMainChanges := FilterChangesBySubAccount(trOp, AccMain)

	// If any of these accounts don't have balance changes, the operations
	// don't match with tobin tax
	lgLockedGoldChange, ok := lgAccMainChanges[lockedGold]
	if !ok {
		return false
	}
	trLockedGoldChange, ok := trAccMainChanges[lockedGold]
	if !ok {
		return false
	}
	trTobinRecipientChange, ok := trAccMainChanges[tobinTax.Recipient]
	if !ok {
		return false
	}

	// Now that we know all the right accounts have balance changes, check if the amounts
	// match up

	// if logOp.lockedGoldContractAccMain.diff == trOp.lockedGoldContractAccMain.diff
	if lgLockedGoldChange.Cmp(trLockedGoldChange) == 0 {

		for address, lgChange := range lgAccMainChanges {
			if address != lockedGold {
				if trChange, ok := trAccMainChanges[address]; ok {

					// if logOp.fromAccMain.diff - trOp.fromAccMain.diff == trOp.tobinRecipientAccMain.diff
					if new(big.Int).Sub(lgChange, trChange).Cmp(trTobinRecipientChange) == 0 {
						// The operations match with tobin tax
						reconcileLogOpWithTransfer(logOp, trOp)
						return true
					}

				}
			}
		}

	}

	return false
}

func reconcileLogOpWithTransfer(logOp, trOp *Operation) {
	// start with the transferOp
	op := *trOp
	// change type to OpLockGold because we will be adding balance changes specific to the log op
	op.Type = logOp.Type
	// add all changes from logOp with SubAccountTypes other than AccMain
	for _, bc := range logOp.Changes {
		if bc.Account.SubAccount.Identifier != AccMain {
			op.Changes = append(op.Changes, bc)
		}
	}
	// replace logOp with reconciled op
	*logOp = op
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
