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
	"math/big"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/kliento/client/debug"
	"github.com/celo-org/rosetta/internal/utils"
	. "github.com/onsi/gomega"
	gs "github.com/onsi/gomega/gstruct"
	gtypes "github.com/onsi/gomega/types"
)

var (
	address1      = common.HexToAddress("0x1111")
	address2      = common.HexToAddress("0x2222")
	address3      = common.HexToAddress("0x3333")
	address4      = common.HexToAddress("0x4444")
	amount1       = big.NewInt(10)
	amount2       = big.NewInt(20)
	success       = debug.TransferStatusSuccess
	fail          = debug.TransferStatusRevert
	emptyTobinTax = NewTobinTax(utils.Big0, common.ZeroAddress)
)

func NewInternalTransfer(from, to common.Address, value *big.Int, status debug.TransferStatus) debug.Transfer {
	return debug.Transfer{
		From:   from,
		To:     to,
		Value:  value,
		Status: status,
	}
}

func NewTestTobinTax(percent int64, recipient common.Address) *TobinTax {
	numerator := new(big.Int).Div(TobinTaxDenominator, big.NewInt(percent))
	return NewTobinTax(numerator, recipient)
}

func MatchTransferOps(transfers []debug.Transfer, tobinTax *TobinTax) gtypes.GomegaMatcher {
	transferOpMatchers := make([]interface{}, len(transfers))
	for i, tr := range transfers {
		transferOpMatchers[i] = MatchTransferOp(tr, tobinTax)
	}
	return And(HaveLen(len(transfers)), ConsistOf(transferOpMatchers...))
}

func MatchTransferOp(transfer debug.Transfer, tobinTax *TobinTax) gtypes.GomegaMatcher {
	return gs.MatchAllFields(gs.Fields{
		"Type":       Equal(OpTransfer),
		"Successful": Equal(transfer.Status.String() == debug.TransferStatusSuccess.String()),
		"Changes":    MatchTransferBalanceChanges(transfer, tobinTax),
	})
}

func MatchTransferBalanceChanges(transfer debug.Transfer, tobinTax *TobinTax) gtypes.GomegaMatcher {
	expectedLen := 2
	balanceChangeMatchers := GetTransferBalanceChangeMatchers(transfer.From, transfer.To, transfer.Value)

	if tobinTax.IsDefined() {
		expectedLen = 4
		taxAmount, afterTaxAmount := tobinTax.Apply(transfer.Value)
		balanceChangeMatchers = append(
			GetTransferBalanceChangeMatchers(transfer.From, transfer.To, afterTaxAmount),
			GetTransferBalanceChangeMatchers(transfer.From, tobinTax.Recipient, taxAmount)...)
	}

	return And(HaveLen(expectedLen), ConsistOf(balanceChangeMatchers...))
}

func GetTransferBalanceChangeMatchers(from, to common.Address, value *big.Int) []interface{} {
	return []interface{}{
		MatchBalanceChange(from, new(big.Int).Neg(value), AccMain),
		MatchBalanceChange(to, value, AccMain),
	}
}

func MatchBalanceChange(address common.Address, value *big.Int, subAccount SubAccountType) gtypes.GomegaMatcher {
	return gs.MatchAllFields(gs.Fields{
		"Amount":  gs.PointTo(Equal(*value)),
		"Account": MatchAccount(address, subAccount),
	})
}

func MatchAccount(address common.Address, subAccount SubAccountType) gtypes.GomegaMatcher {
	return gs.MatchAllFields(gs.Fields{
		"Address": Equal(address),
		"SubAccount": gs.MatchFields(gs.IgnoreExtras, gs.Fields{
			"Identifier": Equal(subAccount),
		}),
	})
}

func TestInternalTransfersToOperations(t *testing.T) {
	RegisterTestingT(t)

	transfers := []debug.Transfer{
		NewInternalTransfer(address1, address2, amount1, success),
		NewInternalTransfer(address2, address1, amount2, success),
		NewInternalTransfer(address2, address3, amount1, success),
		NewInternalTransfer(address3, address1, amount2, fail),
		NewInternalTransfer(address1, address2, amount1, success),
		NewInternalTransfer(address2, address3, amount1, success),
		NewInternalTransfer(address3, address1, amount2, fail),
	}

	t.Run("Without tobinTax", func(t *testing.T) {
		Ω(InternalTransfersToOperations(transfers, emptyTobinTax)).Should(MatchTransferOps(transfers, emptyTobinTax))
	})

	t.Run("With tobinTax", func(t *testing.T) {
		tobinTax := NewTestTobinTax(10, address3)
		Ω(InternalTransfersToOperations(transfers, tobinTax)).Should(MatchTransferOps(transfers, tobinTax))
	})
}

func TestReconcileLogOpWithTransfers(t *testing.T) {
	RegisterTestingT(t)

	slashPenalty := amount2
	slashReward := amount1
	slashRewardFrom := address3
	slashRewardTo := address4
	lockedGoldAdrr := address2

	getTestLogOps := func(tobinTax *TobinTax) []Operation {
		_, afterTaxAmount1 := tobinTax.Apply(amount1)
		return []Operation{
			*NewLockGold(address1, lockedGoldAdrr, afterTaxAmount1),
			*NewWithdrawGold(address1, lockedGoldAdrr, amount1, tobinTax),
			*NewSlash(address1, address2, slashRewardTo, slashRewardFrom, slashPenalty, slashReward, tobinTax),
			*NewUnlockGold(address1, amount1),                                   // does not require transfer
			*NewAuthorizeSigner(address1, address2, OpAuthorizeValidatorSigner), // does not require transfer
		}
	}

	getTestTransferOps := func(tobinTax *TobinTax) []Operation {
		return []Operation{
			*NewTransfer(address1, lockedGoldAdrr, amount1, tobinTax, true),                                           // matches LockGoldOp
			*NewTransfer(lockedGoldAdrr, address1, amount1, tobinTax, false),                                          // matches WithdrawGoldOp
			*NewTransfer(slashRewardFrom, slashRewardTo, new(big.Int).Sub(slashPenalty, slashReward), tobinTax, true), // matches SlashOp
			*NewTransfer(address1, address3, amount1, tobinTax, true),                                                 // does not match a logOp
		}
	}

	getReconciledOps := func(logOps []Operation, tobinTax *TobinTax) []Operation {
		if tobinTax.IsDefined() {
			// This is the operation that should result from calling reconcile() on the LockGoldOp
			logOps[0] = Operation{
				Type: OpLockGold,
				Changes: []BalanceChange{
					{
						Account: Account{
							Address:    address1,
							SubAccount: SubAccount{Identifier: AccMain, Metadata: nil},
						},
						Amount: big.NewInt(-9),
					},
					{
						Account: Account{
							Address:    lockedGoldAdrr,
							SubAccount: SubAccount{Identifier: AccMain, Metadata: nil},
						},
						Amount: big.NewInt(9),
					},
					{
						Account: Account{
							Address:    address1,
							SubAccount: SubAccount{Identifier: AccMain, Metadata: nil},
						},
						Amount: big.NewInt(-1),
					},
					{
						Account: Account{
							Address:    tobinTax.Recipient,
							SubAccount: SubAccount{Identifier: AccMain, Metadata: nil},
						},
						Amount: big.NewInt(1),
					},
					{
						Account: Account{
							Address:    address1,
							SubAccount: SubAccount{Identifier: AccLockedGoldNonVoting, Metadata: nil},
						},
						Amount: big.NewInt(9),
					},
				},
				Successful: true,
			}
		}
		return append(logOps,
			*NewTransfer(address1, address3, amount1, tobinTax, true), // does not match a logOp
		)
	}

	logOps := getTestLogOps(emptyTobinTax)
	transferOps := getTestTransferOps(emptyTobinTax)

	t.Run("Without tobinTax", func(t *testing.T) {
		Ω(ReconcileLogOpsWithTransfers(logOps, transferOps)).Should(ConsistOf(getReconciledOps(logOps, emptyTobinTax)))
	})
}
