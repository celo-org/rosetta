/*
 * Rosetta
 *
 * A standard for blockchain interaction
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

import (
	"context"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
)

// AccountApiService is a service that implents the logic for the AccountApiServicer
// This service should implement the business logic for every endpoint for the AccountApi API.
// Include any external packages or services that will be required by this service.
type AccountApiService struct {
	celoClient  *client.CeloClient
	chainParams *celo.ChainParameters
}

// NewAccountApiService creates a default api service
func NewAccountApiService(celoClient *client.CeloClient, chainParams *celo.ChainParameters) AccountApiServicer {
	return &AccountApiService{
		celoClient:  celoClient,
		chainParams: chainParams,
	}
}

func makeCeloGoldBalance(account string, subaccount string, amount string, metadata *map[string]interface{}) *Balance {
	return &Balance{
		AccountIdentifier: AccountIdentifier{
			Address: account,
			SubAccount: SubAccountIdentifier{
				SubAccount: subaccount,
				Metadata:   *metadata,
			},
		},
		Amounts: []Amount{{
			Value:    amount,
			Currency: CeloGold,
		}},
	}
}

func accountLockedGoldNonVotingBalance(lockedGold *contract.LockedGold, accountAddr common.Address, callOpts *bind.CallOpts) (*Balance, error) {
	lockedGoldNonVotingAmt, err := lockedGold.GetAccountNonvotingLockedGold(callOpts, accountAddr)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("GetAccountTotalLockedGold", err)
	}

	return makeCeloGoldBalance(accountAddr.String(), "LockedGoldNonVoting", lockedGoldNonVotingAmt.String(), nil), nil
}

func electionPendingVotesBalance(election *contract.Election, accountAddr common.Address, groupAddr common.Address, callOpts *bind.CallOpts) (*Balance, error) {
	pendingAmt, err := election.GetPendingVotesForGroupByAccount(callOpts, groupAddr, accountAddr)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("GetPendingVotesForGroupsByAccount", err)
	}

	return makeCeloGoldBalance(accountAddr.String(), "ElectionPendingVotes", pendingAmt.String(), &map[string]interface{}{"group": groupAddr.String()}), nil
}

func electionActiveVotesBalance(election *contract.Election, accountAddr common.Address, groupAddr common.Address, callOpts *bind.CallOpts) (*Balance, error) {
	activeAmt, err := election.GetActiveVotesForGroupByAccount(callOpts, groupAddr, accountAddr)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("GetActiveVotesForGroupsByAccount", err)
	}

	return makeCeloGoldBalance(accountAddr.String(), "ElectionActiveVotes", activeAmt.String(), &map[string]interface{}{"group": groupAddr.String()}), nil
}

// AccountBalance - Get an Account Balance
func (s *AccountApiService) AccountBalance(ctx context.Context, accountBalanceRequest AccountBalanceRequest) (interface{}, error) {

	err := ValidateNetworkId(&accountBalanceRequest.NetworkIdentifier, s.chainParams)
	if err != nil {
		return nil, err
	}

	accountAddr := common.HexToAddress(accountBalanceRequest.AccountIdentifier.Address)

	latestHeader, err := s.celoClient.Eth.HeaderByNumber(ctx, nil) // nil == latest
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("HeaderByNumber", err)
	}
	latestBlockOpt := &bind.CallOpts{
		BlockNumber: latestHeader.Number,
		Context:     ctx,
	}

	goldAmt, err := s.celoClient.Eth.BalanceAt(ctx, accountAddr, latestHeader.Number)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("BalanceAt", err)
	}

	goldBalance := Balance{
		AccountIdentifier: accountBalanceRequest.AccountIdentifier,
		Amounts: []Amount{{
			Value:    goldAmt.String(),
			Currency: CeloGold,
		}},
	}

	registryWrapper, err := wrapper.NewRegistry(s.celoClient)
	if err != nil {
		return nil, err
	}

	lockedGoldAddr, err := registryWrapper.GetAddressFor(latestBlockOpt, params.LockedGoldRegistryId)
	if err != nil {
		return nil, err
	}

	lockedGold, err := contract.NewLockedGold(lockedGoldAddr, s.celoClient.Eth)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("NewLockedGold", err)
	}

	lockedGoldNonVotingBalance, err := accountLockedGoldNonVotingBalance(lockedGold, accountAddr, latestBlockOpt)
	if err != nil {
		return nil, err
	}

	electionAddr, err := registryWrapper.GetAddressFor(latestBlockOpt, params.ElectionRegistryId)
	if err != nil {
		return nil, err
	}

	election, err := contract.NewElection(electionAddr, s.celoClient.Eth)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("NewLockedGold", err)
	}

	groups, err := election.GetGroupsVotedForByAccount(latestBlockOpt, accountAddr)
	if err != nil {
		err = client.WrapRpcError(err)
		return nil, ErrRpcError("GetGroupsVotedForByAccount", err)
	}

	balances := []Balance{goldBalance, *lockedGoldNonVotingBalance}
	for _, groupAddr := range groups {
		pendingBal, err := electionPendingVotesBalance(election, accountAddr, groupAddr, latestBlockOpt)
		if err != nil {
			return nil, err
		}

		activeBal, err := electionActiveVotesBalance(election, accountAddr, groupAddr, latestBlockOpt)
		if err != nil {
			return nil, err
		}

		balances = append(balances, *pendingBal, *activeBal)
	}

	response := AccountBalanceResponse{
		BlockIdentifier: *HeaderToBlockIdentifier(latestHeader),
		Balances:        balances,
	}
	return response, nil
}
