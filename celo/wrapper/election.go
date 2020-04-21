package wrapper

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ElectionWrapper struct {
	*contract.Election
}

func NewElection(celoClient *client.CeloClient, registryWrapper *RegistryWrapper) (*ElectionWrapper, error) {
	election, err := registryWrapper.GetElection(nil, celoClient.Eth)
	if err != nil {
		return nil, err
	}

	return &ElectionWrapper{
		election,
	}, nil
}

type VotesByGroup map[common.Address]*big.Int
type ElectionVotes struct {
	Active  VotesByGroup
	Pending VotesByGroup
}

func (w *ElectionWrapper) GetAccountElectionVotes(opts *bind.CallOpts, account common.Address) (*ElectionVotes, error) {
	groups, err := w.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes *ElectionVotes
	for _, groupAddr := range groups {
		// TODO(yorke): dedup
		pendingAmt, err := w.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if pendingAmt.Cmp(utils.Big0) == 1 {
			votes.Pending[groupAddr] = pendingAmt
		}

		activeAmt, err := w.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if activeAmt.Cmp(utils.Big0) == 1 {
			votes.Active[groupAddr] = activeAmt
		}
	}

	return votes, nil
}

func (w *ElectionWrapper) GetAccountPendingVotes(opts *bind.CallOpts, account common.Address) (VotesByGroup, error) {
	groups, err := w.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes VotesByGroup
	for _, groupAddr := range groups {
		pendingAmt, err := w.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if pendingAmt.Cmp(utils.Big0) == 1 {
			votes[groupAddr] = pendingAmt
		}
	}

	return votes, nil
}

func (w *ElectionWrapper) GetAccountActiveVotes(opts *bind.CallOpts, account common.Address) (VotesByGroup, error) {
	groups, err := w.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes VotesByGroup
	for _, groupAddr := range groups {
		activeAmt, err := w.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		if activeAmt.Cmp(utils.Big0) == 1 {
			votes[groupAddr] = activeAmt
		}
	}

	return votes, nil
}

func (w *ElectionWrapper) GetVotesForGroupByAccount(opts *bind.CallOpts, groupAddr, account common.Address) (*big.Int, error) {
	activeAmt, err := w.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
	if err != nil {
		return nil, err
	}
	pendingAmt, err := w.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
	if err != nil {
		return nil, err
	}
	return new(big.Int).Add(activeAmt, pendingAmt), nil
}
