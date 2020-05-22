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
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type Tracer struct {
	ctx    context.Context
	cc     *client.CeloClient
	db     db.RosettaDBReader
	logger log.Logger
}

func NewTracer(ctx context.Context, cc *client.CeloClient, db db.RosettaDBReader) *Tracer {
	logger := log.New("module", "tracer")
	return &Tracer{
		ctx:    ctx,
		cc:     cc,
		db:     db,
		logger: logger,
	}
}

func (tr *Tracer) GetRegistryAddresses(receipt *types.Receipt, contractNames ...string) (map[string]common.Address, error) {
	contracts, err := tr.db.RegistryAddressesStartOf(tr.ctx, receipt.BlockNumber, receipt.TransactionIndex, contractNames...)
	if err != nil {
		return nil, fmt.Errorf("Error fetching registry addresses: %w", err)
	}

	return contracts, nil
}

func (tr *Tracer) GetTobinTax(blockNumber *big.Int, reserve common.Address) (*TobinTax, error) {
	numerator, err := tr.db.TobinTaxFor(tr.ctx, blockNumber)
	if err != nil {
		return nil, err
	}

	if numerator.Cmp(utils.Big0) > 0 && reserve == common.ZeroAddress {
		// Error: If we find tobinTax in the db, we should also find reserveAddr
		return nil, db.ErrContractNotFound
	}

	return NewTobinTax(numerator, reserve), nil
}

func (tr *Tracer) TraceTransaction(blockHeader *types.Header, tx *types.Transaction, receipt *types.Receipt) ([]Operation, error) {
	ops := make([]Operation, 0)

	if receipt.Status == types.ReceiptStatusSuccessful {
		contracts, err := tr.GetRegistryAddresses(receipt, "Reserve", "LockedGold", "Election", "Governance")
		if err != nil {
			return nil, err
		}

		tobinTax, err := tr.GetTobinTax(receipt.BlockNumber, contracts["Reserve"])
		if err != nil {
			return nil, err
		}

		lockedGoldOps, err := tr.TxLockedGoldTransfers(tx, receipt, tobinTax, contracts)
		if err != nil {
			return nil, err
		}

		transferOps, err := tr.TxTransfers(tx, receipt, tobinTax)
		if err != nil {
			return nil, err
		}

		ops, err = ReconcileLockedGoldTransfers(lockedGoldOps, transferOps)
		if err != nil {
			return nil, err
		}
	}

	if tx.FeeCurrency() == nil { // nil implies cGLD
		gasOp, err := tr.TxGasDetails(blockHeader.Coinbase, tx, receipt)
		if err != nil {
			return nil, err
		}
		ops = append(ops, *gasOp)
	}

	return ops, nil
}

func (tr *Tracer) TxGasDetails(coinbase common.Address, tx *types.Transaction, receipt *types.Receipt) (*Operation, error) {
	balanceChanges := NewBalanceSet()

	gpm, err := tr.db.GasPriceMinimumFor(tr.ctx, receipt.BlockNumber)
	if err != nil {
		return nil, fmt.Errorf("can't get gasPriceMinimun: %w", err)
	}

	gasUsed := new(big.Int).SetUint64(receipt.GasUsed)

	// baseTxFee is what goes to the community fund (if any)
	baseTxFee := new(big.Int).Mul(gpm, gasUsed)
	totalTxFee := new(big.Int).Mul(tx.GasPrice(), gasUsed)

	// The "tip" goes to the coinbase address
	balanceChanges.Add(coinbase, new(big.Int).Sub(totalTxFee, baseTxFee))

	// We want to get state AFTER the tx, since gas fees are processed by the end of the TX
	governanceAddress, err := tr.db.RegistryAddressStartOf(tr.ctx, receipt.BlockNumber, receipt.TransactionIndex+1, "Governance")
	if err == db.ErrContractNotFound {
		// No community fund, we won't charge the user
		totalTxFee.Sub(totalTxFee, baseTxFee)
	} else if err == nil {
		// The baseTxFee goes to the community fund
		balanceChanges.Add(governanceAddress, baseTxFee)
	} else {
		return nil, fmt.Errorf("can't get governanceAddress: %w", err)
	}

	if tx.GatewayFeeRecipient() != nil {
		balanceChanges.Add(*tx.GatewayFeeRecipient(), tx.GatewayFee())
		totalTxFee.Add(totalTxFee, tx.GatewayFee())
	}

	// TODO find a better way to do this?
	from, err := tr.cc.Eth.TransactionSender(tr.ctx, tx, receipt.BlockHash, receipt.TransactionIndex)
	if err != nil {
		return nil, fmt.Errorf("can't get transaction sender: %w", err)
	}
	balanceChanges.Add(from, new(big.Int).Neg(totalTxFee))
	return NewFee(balanceChanges.ToMap()), nil
}

func (tr *Tracer) TxTransfers(tx *types.Transaction, receipt *types.Receipt, tobinTax *TobinTax) ([]Operation, error) {
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("Cannot trace transfers on failed tx: %v  blockNumber: %v", receipt.TransactionIndex, receipt.BlockNumber)
	}

	internalTransfers, err := tr.cc.Debug.TransactionTransfers(tr.ctx, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("can't run celo-rpc tx-tracer: %w", err)
	}

	return InternalTransfersToOperations(internalTransfers, tobinTax), nil
}

func (tr *Tracer) TxLockedGoldTransfers(tx *types.Transaction, receipt *types.Receipt, tobinTax *TobinTax, contracts map[string]common.Address) ([]Operation, error) {
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, fmt.Errorf("Cannot trace LockedGold transfers on failed tx: %v  blockNumber: %v", receipt.TransactionIndex, receipt.BlockNumber)
	}

	lockedGoldAddr, ok := contracts["LockedGold"]
	if !ok {
		// LockedGold not found (not deployed) => no transfers
		return nil, nil
	}
	lockedGold, err := contract.NewLockedGold(lockedGoldAddr, tr.cc.Eth)
	if err != nil {
		return nil, fmt.Errorf("can't initialize LockedGold contract: %w", err)
	}

	electionAddr, ok := contracts["Election"]
	if !ok {
		// Election not found (not deployed) => no transfers
		return nil, nil
	}
	election, err := contract.NewElection(electionAddr, tr.cc.Eth)
	if err != nil {
		return nil, fmt.Errorf("can't initialize Election contract: %w", err)
	}

	// We only need governace for slashing and you can't slash if there's no governance contract, so we ignore if not found
	governanceAddr, _ := contracts["Governance"]

	logs := utils.RemoveProxyLogs(receipt.Logs)

	transfers := make([]Operation, 0, len(logs))
	for _, eventLog := range logs {
		if eventLog.Address == electionAddr {
			eventName, eventRaw, ok, err := election.TryParseLog(*eventLog)
			if err != nil {
				return nil, fmt.Errorf("can't parse Election event: %w", err)
			}
			if !ok {
				continue
			}

			// Election:

			switch eventName {
			case "ValidatorGroupVoteCast":
				// vote() [ValidatorGroupVoteCast] => lockNonVoting->lockVotingPending
				event := eventRaw.(*contract.ElectionValidatorGroupVoteCast)
				transfers = append(transfers, *NewVote(event.Account, event.Group, event.Value))
			case "ValidatorGroupVoteActivated":
				// activate() [ValidatorGroupVoteActivated] => lockVotingPending->lockVotingActive
				event := eventRaw.(*contract.ElectionValidatorGroupVoteActivated)
				transfers = append(transfers, *NewActiveVotes(event.Account, event.Group, event.Value))
			case "ValidatorGroupPendingVoteRevoked":
				// revokePending() [ValidatorGroupPendingVoteRevoked] => lockVotingPending->lockNonVoting
				event := eventRaw.(*contract.ElectionValidatorGroupPendingVoteRevoked)
				transfers = append(transfers, *NewRevokePendingVotes(event.Account, event.Group, event.Value))
			case "ValidatorGroupActiveVoteRevoked":
				// revokeActive() [ValidatorGroupActiveVoteRevoked] => lockVotingActive->lockNonVoting
				event := eventRaw.(contract.ElectionValidatorGroupActiveVoteRevoked)
				transfers = append(transfers, *NewRevokeActiveVotes(event.Account, event.Group, event.Value))
			}

		} else if eventLog.Address == lockedGoldAddr {
			eventName, eventRaw, ok, err := lockedGold.TryParseLog(*eventLog)
			if err != nil {
				return nil, fmt.Errorf("can't parse LockedGold event: %w", err)
			}
			if !ok {
				continue
			}

			switch eventName {
			case "GoldLocked":
				// lock() [GoldLocked + transfer] => main->lockNonVoting
				event := eventRaw.(*contract.LockedGoldGoldLocked)
				transfers = append(transfers, *NewLockGold(event.Account, lockedGoldAddr, event.Value, tobinTax))
				// relock() [GoldLocked] => lockPending->lockNonVoting

			case "GoldUnlocked":
				// unlock() [GoldUnlocked] => lockNonVoting->lockPending
				event := eventRaw.(*contract.LockedGoldGoldUnlocked)
				transfers = append(transfers, *NewUnlockGold(event.Account, event.Value))

			case "GoldWithdrawn":
				// withdraw() [GoldWithdrawn + transfer] => lockPending->main
				event := eventRaw.(*contract.LockedGoldGoldWithdrawn)
				transfers = append(transfers, *NewWithdrawGold(event.Account, lockedGoldAddr, event.Value, tobinTax))

			case "AccountSlashed":
				// slash() [AccountSlashed + transfer] => account:lockNonVoting -> beneficiary:lockNonVoting + governance:main
				event := eventRaw.(*contract.LockedGoldAccountSlashed)
				transfers = append(transfers, *NewSlash(event.Slashed, event.Reporter, governanceAddr, lockedGoldAddr, event.Penalty, event.Reward, tobinTax))

			}
		}

	}

	return transfers, nil
}
