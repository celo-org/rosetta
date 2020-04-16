package wrapper

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type LockedGoldWrapper struct {
	contract *contract.LockedGold
}

func NewLockedGold(celoClient *client.CeloClient, registryWrapper *RegistryWrapper) (*LockedGoldWrapper, error) {
	lockedgold, err := registryWrapper.GetLockedGold(nil, celoClient.Eth)
	if err != nil {
		return nil, err
	}

	return &LockedGoldWrapper{
		contract: lockedgold,
	}, nil
}

type PendingWithdrawal struct {
	Amount    *big.Int
	Timestamp *big.Int
}

type NonVotingLockedGold struct {
	Amount             *big.Int
	PendingWithdrawals []PendingWithdrawal
}

func (w *LockedGoldWrapper) GetAccountNonVotingLockedGold(account common.Address, opts *bind.CallOpts) (*NonVotingLockedGold, error) {
	lockedGoldNonVotingAmt, err := w.contract.GetAccountNonvotingLockedGold(opts, account)
	if err != nil {
		return nil, err
	}

	values, timestamps, err := w.contract.GetPendingWithdrawals(opts, account)
	if err != nil {
		return nil, err
	}

	withdrawals := make([]PendingWithdrawal, len(values))
	for idx, val := range values {
		withdrawals[idx].Amount = val
		withdrawals[idx].Timestamp = timestamps[idx]
	}

	result := NonVotingLockedGold{
		Amount:             lockedGoldNonVotingAmt,
		PendingWithdrawals: withdrawals,
	}
	return &result, nil
}
