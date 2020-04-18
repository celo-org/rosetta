package rpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/analyzer"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	gethTypes "github.com/ethereum/go-ethereum/core/types"
)

// AccountApiService is a service that implents the logic for the AccountApiServicer
// This service should implement the business logic for every endpoint for the AccountApi API.
// Include any external packages or services that will be required by this service.
type AccountApiService struct {
	celoClient  *client.CeloClient
	chainParams *celo.ChainParameters
}

// NewAccountApiService creates a default api service
func NewAccountApiService(celoClient *client.CeloClient, chainParams *celo.ChainParameters) server.AccountAPIServicer {
	return &AccountApiService{
		celoClient:  celoClient,
		chainParams: chainParams,
	}
}

// AccountBalance - Get an Account Balance
func (s *AccountApiService) AccountBalance(ctx context.Context, accountBalanceRequest *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {

	errResp := ValidateNetworkId(accountBalanceRequest.NetworkIdentifier, s.chainParams)
	if errResp != nil {
		return nil, errResp
	}

	accountAddr := common.HexToAddress(accountBalanceRequest.AccountIdentifier.Address)

	var requestedHeader *gethTypes.Header
	var err error
	blockIdx := accountBalanceRequest.BlockIdentifier.Index
	if blockIdx != nil {
		blockNumber := big.NewInt(*blockIdx)
		requestedHeader, err = s.celoClient.Eth.HeaderByNumber(ctx, blockNumber)
	} else {
		blockHash := accountBalanceRequest.BlockIdentifier.Hash
		if blockHash == nil {
			return nil, NewValidationError(fmt.Errorf("No block hash or index provided"))
		}
		requestedHeader, err = s.celoClient.Eth.HeaderByHash(ctx, common.HexToHash(*blockHash))
	}
	if err != nil {
		return nil, ErrCantFetchBlockHeader(fmt.Errorf("Not found"))
	}

	requestedBlockOpts := &bind.CallOpts{
		BlockNumber: requestedHeader.Number,
		Context:     ctx,
	}

	emptyResponse := &types.AccountBalanceResponse{
		BlockIdentifier: HeaderToBlockIdentifier(requestedHeader),
	}
	createReponse := func(amount *types.Amount) *types.AccountBalanceResponse {
		return &types.AccountBalanceResponse{
			BlockIdentifier: HeaderToBlockIdentifier(requestedHeader),
			Balances:        []*types.Amount{amount},
		}
	}

	subAccount := accountBalanceRequest.AccountIdentifier.SubAccount

	if subAccount == nil {
		// Main Account, just cGLD
		goldAmt, err := s.celoClient.Eth.BalanceAt(ctx, accountAddr, requestedHeader.Number)
		if err != nil {
			return nil, NewCeloClientError("BalanceAt", err)
		}
		return createReponse(NewAmount(goldAmt, CeloGold)), nil
	}

	registryWrapper, err := wrapper.NewRegistry(s.celoClient)
	if err == client.ErrContractNotDeployed {
		// Nothing is deployed => ignore lockedGold & election balances
		return emptyResponse, nil
	} else if err != nil {
		return nil, NewCeloClientError("NewRegistry", err)
	}

	if subAccount.Address == string(analyzer.AccLockedGoldNonVoting) {
		// Fetch LockedGold Balances
		lockedGoldWrapper, err := wrapper.NewLockedGold(s.celoClient, registryWrapper)
		if err == client.ErrContractNotDeployed {
			// Nothing is deployed => ignore lockedGold & election balances
			return emptyResponse, nil
		} else if err != nil {
			return nil, NewCeloClientError("NewLockedGold", err)
		}

		nonVotingLockedGold, err := lockedGoldWrapper.GetAccountNonvotingLockedGold(requestedBlockOpts, accountAddr)
		if err != nil {
			return nil, NewCeloClientError("GetAccountNonvotingLockedGold", err)
		}

		return createReponse(NewAmount(nonVotingLockedGold, CeloGold)), nil
	}

	if subAccount.Address == string(analyzer.AccLockedGoldPending) {
		// Fetch LockedGold Balances
		lockedGoldWrapper, err := wrapper.NewLockedGold(s.celoClient, registryWrapper)
		if err == client.ErrContractNotDeployed {
			// Nothing is deployed => ignore lockedGold & election balances
			return emptyResponse, nil
		} else if err != nil {
			return nil, NewCeloClientError("NewLockedGold", err)
		}

		totalPending, err := lockedGoldWrapper.GetTotalPendingWithdrawals(requestedBlockOpts, accountAddr)
		if err != nil {
			return nil, NewCeloClientError("GetTotalPendingWithdrawals", err)
		}

		return createReponse(NewAmount(totalPending, CeloGold)), nil
	}

	// If we are here need to be election based

	// Fetch Election (Votes) Balances
	electionWrapper, err := wrapper.NewElection(s.celoClient, registryWrapper)
	if err == client.ErrContractNotDeployed {
		// Nothing is deployed => ignore lockedGold & election balances
		return emptyResponse, nil
	} else if err != nil {
		return nil, NewCeloClientError("NewElection", err)
	}

	// For now we only support "pending" votes

	electionVotes, err := electionWrapper.GetAccountElectionVotes(requestedBlockOpts, accountAddr)
	if err != nil {
		return nil, NewCeloClientError("GetAccountElectionVotes", err)
	}

	if subAccount.Metadata != nil {
		if groupStr, ok := (*subAccount.Metadata)["group"]; ok {
			groupAddr := common.HexToAddress(groupStr.(string))
			pendingAmt, ok := electionVotes.Pending[groupAddr]
			if !ok {
				pendingAmt = utils.Big0
			}
			activeAmt, ok := electionVotes.Active[groupAddr]
			if !ok {
				activeAmt = utils.Big0
			}
			total := new(big.Int).Add(pendingAmt, activeAmt)
			return createReponse(NewAmount(total, CeloGold)), nil
		}
	}

	// On RC0 we can't track pending vs active votes, so we sum them up as "pending"
	// TODO(rc1) fix
	// correct code
	// for groupAddr, activeAmt := range electionVotes.Active {
	// 	account := analyzer.NewVotingAccount(accountAddr, analyzer.AccLockedGoldVotingActive, groupAddr)
	// 	balances = append(balances, *NewCeloGoldBalance(account, activeAmt))
	// }

	// for groupAddr, pendingAmt := range electionVotes.Pending {
	// 	account := analyzer.NewVotingAccount(accountAddr, analyzer.AccLockedGoldVotingPending, groupAddr)
	// 	balances = append(balances, *NewCeloGoldBalance(account, pendingAmt))
	// }

	return nil, NewCeloClientError("InvalidAccountIdentifier", err)
}
