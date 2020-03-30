package celo

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type TxTracer struct {
	blockHeader *types.Header
	tx          *types.Transaction
	receipt     *types.Receipt

	celoClient *client.CeloClient
	ctx        context.Context
	logger     log.Logger
}

func NewTxTracer(
	ctx context.Context,
	celoClient *client.CeloClient,
	blockHeader *types.Header,
	tx *types.Transaction,
	receipt *types.Receipt,
) *TxTracer {
	logger := log.New("context", "TxTracer", "txHash", tx.Hash().Hex(), "txIndex", receipt.TransactionIndex, "blockHash", blockHeader.Hash().Hex(), "blockNumber", blockHeader.Number)
	return &TxTracer{
		celoClient:  celoClient,
		blockHeader: blockHeader,
		tx:          tx,
		receipt:     receipt,
		ctx:         ctx,
		logger:      logger,
	}
}

func (tc *TxTracer) ObtainRegistryAddresses(identifiers ...[32]byte) (map[[32]byte]common.Address, error) {
	addresses := make(map[[32]byte]common.Address)
	registry, err := wrapper.NewRegistry(tc.celoClient)
	if err != nil {
		return addresses, err
	}

	// Check the Address on the previous block
	callOpts := &bind.CallOpts{
		BlockNumber: new(big.Int).Sub(tc.blockHeader.Number, big.NewInt(1)),
		Context:     tc.ctx,
	}

	for _, identifier := range identifiers {
		address, err := registry.GetAddressFor(callOpts, identifier)
		if err == client.ErrContractNotDeployed {
			// Ignore this error. Attempt to obtain others
			continue
		} else if err != nil {
			return nil, err
		}
		addresses[identifier] = address
	}

	// Check if address changed between last block and previous transaction
	// on the same block
	updates, err := registry.GetUpdatesOnBlock(tc.ctx, tc.blockHeader.Number.Uint64(), &tc.receipt.TransactionIndex, identifiers)
	if err != nil {
		return nil, err
	}
	for id, address := range updates {
		addresses[id] = address
	}

	return addresses, nil
}

func (tc *TxTracer) GasPriceMinimum(gpmAddress common.Address) (*big.Int, error) {
	gpm, err := contract.NewGasPriceMinimum(gpmAddress, tc.celoClient.Eth)
	if err != nil {
		return nil, err
	}

	// Check the price on the previous block on the previous block
	callOpts := &bind.CallOpts{
		BlockNumber: new(big.Int).Sub(tc.blockHeader.Number, big.NewInt(1)),
		Context:     tc.ctx,
	}
	gasPrice, err := gpm.GetGasPriceMinimum(callOpts, common.ZeroAddress)
	if err != nil {
		return nil, err
	}

	// GasPriceMinimun is updated at the end of each block
	// So no need to check previuos tx within the same block

	return gasPrice, nil
}

func (tc *TxTracer) GasDetail() (map[common.Address]*big.Int, error) {
	balanceChanges := make(map[common.Address]*big.Int)

	registryAddresses, err := tc.ObtainRegistryAddresses(params.GasPriceMinimumRegistryId, params.GovernanceRegistryId)
	if err != nil {
		return nil, err
	}

	var gasPriceMinimum *big.Int
	if gpmAddress, ok := registryAddresses[params.GasPriceMinimumRegistryId]; ok {
		gasPriceMinimum, err = tc.GasPriceMinimum(gpmAddress)
		// TODO - What happens when there's no gasPrice Minimun
		if err != nil {
			fmt.Println("Failed go get GasPriceMinum")
		}
	}

	gasUsed := new(big.Int).SetUint64(tc.receipt.GasUsed)

	// baseTxFee is what goes to the community fund (if any)
	baseTxFee := new(big.Int).Mul(gasPriceMinimum, gasUsed)
	totalTxFee := new(big.Int).Mul(tc.tx.GasPrice(), gasUsed)

	// The "tip" goes to the coinbase address
	balanceChanges[tc.blockHeader.Coinbase] = new(big.Int).Sub(totalTxFee, baseTxFee)

	if governanceAddress, ok := registryAddresses[params.GovernanceRegistryId]; ok {
		// The baseTxFee goes to the community fund
		balanceChanges[governanceAddress] = baseTxFee
	} else {
		// No community fund, we won't charge the user
		totalTxFee.Sub(totalTxFee, baseTxFee)
	}

	if tc.tx.GatewayFeeRecipient() != nil {
		balanceChanges[*tc.tx.GatewayFeeRecipient()] = tc.tx.GatewayFee()
		totalTxFee.Add(totalTxFee, tc.tx.GatewayFee())
	}

	// TODO find a better way to do this?
	from, err := tc.celoClient.Eth.TransactionSender(tc.ctx, tc.tx, tc.blockHeader.Hash(), tc.receipt.TransactionIndex)
	if err != nil {
		return balanceChanges, err
	}
	balanceChanges[from] = new(big.Int).Neg(totalTxFee)

	return balanceChanges, nil
}

func (tc *TxTracer) TransferDetail() ([]Transfer, error) {
	if tc.receipt.Status == types.ReceiptStatusFailed {
		return nil, nil
	}

	internalTransfers, err := tc.celoClient.Debug.TransactionTransfers(tc.ctx, tc.tx.Hash())
	if err != nil {
		return nil, err
	}

	transfers := make([]Transfer, len(internalTransfers))
	for i, it := range internalTransfers {
		transfers[i] = Transfer{
			From: Account{
				Address:    it.From,
				SubAccount: Main,
			},
			To: Account{
				Address:    it.To,
				SubAccount: Main,
			},
			Value: it.Value,
		}
	}
	return transfers, nil
}

func (tc *TxTracer) LockedGoldTransferDetail() ([]Transfer, error) {
	registryAddresses, err := tc.ObtainRegistryAddresses(params.LockedGoldRegistryId, params.GovernanceRegistryId)
	if err != nil {
		return nil, err
	}
	lockedGoldAddr, ok := registryAddresses[params.LockedGoldRegistryId]

	// TODO deal with Error

	// LockedGold not found (not deployed) => no transfers
	if !ok {
		return nil, nil
	}

	lockedGold, err := contract.NewLockedGold(lockedGoldAddr, tc.celoClient.Eth)
	if err != nil {
		return nil, err
	}

	transfers := make([]Transfer, 0, len(tc.receipt.Logs))
	for _, eventLog := range tc.receipt.Logs {
		eventName, eventRaw, ok, err := lockedGold.TryParseLog(*eventLog)
		if err != nil {
			return nil, err
		}
		if !ok {
			continue
		}

		switch eventName {
		case "AccountSlashed":
			governanceAddr, ok := registryAddresses[params.GovernanceRegistryId]
			if !ok {
				return nil, fmt.Errorf("Can't slash before governance is deployed")
			}

			event := eventRaw.(contract.LockedGoldAccountSlashed)
			transfers = append(transfers,
				Transfer{
					From: Account{
						Address:    event.Slashed,
						SubAccount: LockedGoldLocked,
					},
					To: Account{
						Address:    event.Reporter,
						SubAccount: LockedGoldLocked,
					},
					Value: event.Reward,
				},
				Transfer{
					From: Account{
						Address:    event.Slashed,
						SubAccount: LockedGoldLocked,
					},
					To: Account{
						Address:    governanceAddr,
						SubAccount: Main,
					},
					Value: new(big.Int).Sub(event.Penalty, event.Reward),
				},
			)

		case "GoldLocked":
			event := eventRaw.(contract.LockedGoldGoldLocked)
			transfers = append(transfers, Transfer{
				From: Account{
					Address:    event.Account,
					SubAccount: Main,
				},
				To: Account{
					Address:    event.Account,
					SubAccount: LockedGoldLocked,
				},
				Value: event.Value,
			})

		case "GoldUnlocked":
			event := eventRaw.(contract.LockedGoldGoldUnlocked)
			transfers = append(transfers, Transfer{
				From: Account{
					Address:    event.Account,
					SubAccount: LockedGoldLocked,
				},
				To: Account{
					Address:    event.Account,
					SubAccount: LockedGoldPending,
				},
				Value: event.Value,
			})
		case "GoldWithdrawn":
			event := eventRaw.(contract.LockedGoldGoldWithdrawn)
			transfers = append(transfers, Transfer{
				From: Account{
					Address:    event.Account,
					SubAccount: LockedGoldPending,
				},
				To: Account{
					Address:    event.Account,
					SubAccount: Main,
				},
				Value: event.Value,
			})
		}

	}

	return transfers, nil
}

/*

compute

Compute Operations
 - compute Gas Operations
 - compute Transfer Operations

Compute Gas Operations:
	- If feeCurrency != cGLD => EXIT
	- From Block => coinbase, number
	- From tx => gatewayFee, gasPrice, gatewayFeeRecipient, index
	- From receipt => gasUsed, status
	- From Registry: Governance, GasPriceMinimun
	- From GasPriceMinimun: gasPriceMinimum
	- Error Conditions:
		- Fail to fetch Governance => No communityFund Fee
		- Fail to fetch gasPriceMinimun => ?
		- Fail to obtain block, tx, receipt => error

Compute Transfer Operations;
	- From receipt => status
	- If status == failure => EXIT (no transfers)
	- Call debug_traceTransaction
	- From Registry: LockedGold
	- From LockedGold => All Events
	- Error Conditions:
		- Fail to fetch LockedGold.add => not deployed, ignore lockedGold subaccount
		- Fail to debug_traceTRansactions =>


{
	From: {
		account common.Address
		subaccount *SubAccountKind
	}
	To: {
		account common.Address
		subaccount *SubAccountKind
	}
	Value: *big.Int


}

*/
