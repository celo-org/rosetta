package analyzer

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/client/debug"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/db"
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

func (tr *Tracer) TraceTransaction(blockHeader *types.Header, tx *types.Transaction, receipt *types.Receipt) ([]Operation, error) {
	var operations []Operation

	gasOperation, err := tr.TxGasDetails(blockHeader, tx, receipt)
	if err != nil {
		return nil, err
	}
	operations = append(operations, *gasOperation)

	lockedGoldOperations, err := tr.TxLockedGoldTransfers(blockHeader, tx, receipt)
	if err != nil {
		return nil, err
	}
	operations = append(operations, lockedGoldOperations...)

	transferOperations, err := tr.TxTransfers(blockHeader, tx, receipt)
	if err != nil {
		return nil, err
	}

	// Only add non-redundant transfers

	// We assume both arrays are in order
	// then look for the first lockedGold operation with a change in AccMain
	// and look for the matching transfer
	ti := 0
	for _, lgOp := range lockedGoldOperations {
		// only interested in lockedGold with credit/degits on AccMain
		if lgOp.Type == OpLockGold || lgOp.Type == OpWithdrawGold || lgOp.Type == OpSlash {
			// search for the next transfer that match this
			for ; ti < len(transferOperations) && !MatchChangesOnSubAccount(&lgOp, &transferOperations[ti], AccMain); ti++ {
				// if it doesn't match, it's good to add it to the operations
				operations = append(operations, transferOperations[ti])
			}
			if ti == len(transferOperations) {
				// we didn't find it... this means we have a bug
				tr.logger.Error("BUG: Can't find matching transfer for LockedGold op", "block", blockHeader.Number, "txHash", tx.Hash())
			}
			ti++ // we skip the matched operation and continue

		}
	}
	// add the rest of the transfers
	if ti < len(transferOperations) {
		operations = append(operations, transferOperations[ti:]...)
	}

	return operations, nil
}

func (tr *Tracer) TxGasDetails(blockHeader *types.Header, tx *types.Transaction, receipt *types.Receipt) (*Operation, error) {
	balanceChanges := NewBalanceSet()

	gpm, err := tr.db.GasPriceMinimumFor(tr.ctx, blockHeader.Number)
	if err != nil {
		return nil, fmt.Errorf("can't get gasPriceMinimun: %w", err)
	}

	gasUsed := new(big.Int).SetUint64(receipt.GasUsed)

	// baseTxFee is what goes to the community fund (if any)
	baseTxFee := new(big.Int).Mul(gpm, gasUsed)
	totalTxFee := new(big.Int).Mul(tx.GasPrice(), gasUsed)

	// The "tip" goes to the coinbase address
	balanceChanges.Add(blockHeader.Coinbase, new(big.Int).Sub(totalTxFee, baseTxFee))

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
	from, err := tr.cc.Eth.TransactionSender(tr.ctx, tx, blockHeader.Hash(), receipt.TransactionIndex)
	if err != nil {
		return nil, fmt.Errorf("can't get transaction sender: %w", err)
	}
	balanceChanges.Add(from, new(big.Int).Neg(totalTxFee))
	return NewFee(balanceChanges.ToMap()), nil
}

func (tr *Tracer) TxTransfers(blockHeader *types.Header, tx *types.Transaction, receipt *types.Receipt) ([]Operation, error) {
	if receipt.Status == types.ReceiptStatusFailed {
		return nil, nil
	}

	internalTransfers, err := tr.cc.Debug.TransactionTransfers(tr.ctx, tx.Hash())
	if err != nil {
		return nil, fmt.Errorf("can't run celo-rcp tx-tracer: %w", err)
	}

	transfers := make([]Operation, len(internalTransfers))
	for i, it := range internalTransfers {
		transfers[i] = *NewTransfer(it.From, it.To, it.Value, it.Status.String() == debug.TransferStatusSuccess.String())
	}
	return transfers, nil
}

func (tr *Tracer) TxLockedGoldTransfers(blockHeader *types.Header, tx *types.Transaction, receipt *types.Receipt) ([]Operation, error) {
	lockedGoldAddr, err := tr.db.RegistryAddressStartOf(tr.ctx, receipt.BlockNumber, receipt.TransactionIndex, "LockedGold")
	if err == db.ErrContractNotFound {
		// LockedGold not found (not deployed) => no transfers
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("can't get lockedGoldAddress: %w", err)
	}
	lockedGold, err := contract.NewLockedGold(lockedGoldAddr, tr.cc.Eth)
	if err != nil {
		return nil, fmt.Errorf("can't initialize LockedGold contract: %w", err)
	}

	electionAddr, err := tr.db.RegistryAddressStartOf(tr.ctx, receipt.BlockNumber, receipt.TransactionIndex, "Election")
	if err == db.ErrContractNotFound {
		// Election not found (not deployed) => no transfers
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("can't get electionAddress: %w", err)
	}
	election, err := contract.NewElection(electionAddr, tr.cc.Eth)
	if err != nil {
		return nil, fmt.Errorf("can't initialize Election contract: %w", err)
	}

	governanceAddr, err := tr.db.RegistryAddressStartOf(tr.ctx, receipt.BlockNumber, receipt.TransactionIndex, "Governance")
	if err != nil && err != db.ErrContractNotFound {
		// We only need governace for slashing and you can't slash if there no governance contract. So we ignore the error
		return nil, fmt.Errorf("can't get governanceAddress: %w", err)
	}

	transfers := make([]Operation, 0, len(receipt.Logs))
	for _, eventLog := range receipt.Logs {
		if eventLog.Address == electionAddr {
			eventName, eventRaw, ok, err := election.TryParseLog(*eventLog)
			if err != nil {
				tr.logger.Warn("Ignoring event: Error parsing election event", "err", err, "eventId", eventLog.Topics[0].Hex(), "tx", receipt.TxHash.Hex())
				// return nil, fmt.Errorf("can't parse Election event: %w", err)
				continue
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
				transfers = append(transfers, *NewLockGold(event.Account, lockedGoldAddr, event.Value))
				// relock() [GoldLocked] => lockPending->lockNonVoting

			case "GoldUnlocked":
				// unlock() [GoldUnlocked] => lockNonVoting->lockPending
				event := eventRaw.(*contract.LockedGoldGoldUnlocked)
				transfers = append(transfers, *NewUnlockGold(event.Account, event.Value))

			case "GoldWithdrawn":
				// withdraw() [GoldWithdrawn + transfer] => lockPending->main
				event := eventRaw.(*contract.LockedGoldGoldWithdrawn)
				transfers = append(transfers, *NewWithdrawGold(event.Account, lockedGoldAddr, event.Value))

			case "AccountSlashed":
				// slash() [AccountSlashed + transfer] => account:lockNonVoting -> beneficiary:lockNonVoting + governance:main
				event := eventRaw.(*contract.LockedGoldAccountSlashed)
				transfers = append(transfers, *NewSlash(event.Slashed, event.Reporter, governanceAddr, lockedGoldAddr, event.Penalty, event.Reward))

			}
		}

	}

	return transfers, nil
}
