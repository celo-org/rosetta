package wrapper

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ElectionWrapper struct {
	contract *contract.Election
}

func NewElection(celoClient *client.CeloClient, registryWrapper *RegistryWrapper) (*ElectionWrapper, error) {
	election, err := registryWrapper.GetElection(nil, celoClient.Eth)
	if err != nil {
		return nil, err
	}

	return &ElectionWrapper{
		contract: election,
	}, nil
}

type VotesByGroup map[common.Address]*big.Int
type ElectionVotes struct {
	Active  VotesByGroup
	Pending VotesByGroup
}

func (w *ElectionWrapper) GetAccountElectionVotes(account common.Address, opts *bind.CallOpts) (*ElectionVotes, error) {
	groups, err := w.contract.GetGroupsVotedForByAccount(opts, account)
	if err != nil {
		return nil, err
	}

	var votes *ElectionVotes
	for _, groupAddr := range groups {
		pendingAmt, err := w.contract.GetPendingVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		votes.Pending[groupAddr] = pendingAmt

		activeAmt, err := w.contract.GetActiveVotesForGroupByAccount(opts, groupAddr, account)
		if err != nil {
			return nil, err
		}
		votes.Active[groupAddr] = activeAmt
	}

	return votes, nil
}
