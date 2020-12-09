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

package rpc

import (
	"context"
	"errors"
	"fmt"
	"math/big"

	"github.com/celo-org/kliento/client"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/contracts/helpers"
	"github.com/celo-org/kliento/registry"
	"github.com/celo-org/kliento/utils/chain"
	"github.com/celo-org/rosetta/airgap"
	"github.com/celo-org/rosetta/airgap/server"
	"github.com/celo-org/rosetta/analyzer"
	"github.com/celo-org/rosetta/db"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Servicer is a service that implents the logic for the Servicerr
// This service should implement the business logic for every endpoint for the AccountApi API.
// Include any external packages or services that will be required by this service.
type Servicer struct {
	cc          *client.CeloClient
	db          db.RosettaDBReader
	chainParams *chain.ChainParameters
	airgap      airgap.Server
}

// NewServicer creates a default api service
func NewServicer(celoClient *client.CeloClient, db db.RosettaDBReader, cp *chain.ChainParameters) (*Servicer, error) {
	srvCtx, err := server.NewServerContext(celoClient)
	if err != nil {
		return nil, err
	}

	airgap, err := server.NewAirgapServer(cp.ChainId, srvCtx)
	if err != nil {
		return nil, err
	}

	return &Servicer{
		cc:          celoClient,
		db:          db,
		chainParams: cp,
		airgap:      airgap,
	}, nil
}

// Mempool - Get All Mempool Transactions
func (s *Servicer) Mempool(ctx context.Context, request *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {

	content, err := s.cc.TxPool.Content(ctx)
	if err != nil {
		return nil, LogErrCeloClient("TxPoolContent", err)
	}

	allTransactionIds := append(TxIdsFromTxAccountMap((*content)["pending"]), TxIdsFromTxAccountMap((*content)["queued"])...)

	response := types.MempoolResponse{
		TransactionIdentifiers: allTransactionIds,
	}
	return &response, nil
}

// MempoolTransaction - Get a Mempool Transaction
func (s *Servicer) MempoolTransaction(ctx context.Context, request *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	return nil, LogErrUnimplemented("/mempool/transaction")
}

func (s *Servicer) NetworkList(ctx context.Context, request *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	response := types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{{
			Blockchain: BlockchainName,
			Network:    s.chainParams.ChainId.String(),
		}},
	}

	return &response, nil
}

func (s *Servicer) NetworkOptions(ctx context.Context, request *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	response := types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion:    RosettaVersion,
			NodeVersion:       NodeVersion,
			MiddlewareVersion: &MiddlewareVersion,
		},
		Allow: &types.Allow{
			OperationStatuses: []*types.OperationStatus{
				OperationFailed.ToOperationStatus(),
				OperationSuccess.ToOperationStatus(),
			},
			OperationTypes: analyzer.AllOperationTypesString(),
			Errors: []*types.Error{
				ErrValidation,
				ErrUnimplemented,
				ErrInternal,
				ErrCeloClient,
			},
		},
	}

	return &response, nil
}

// NetworkStatus - Get Network Status
func (s *Servicer) NetworkStatus(ctx context.Context, request *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {

	lastPersitedBlock, err := s.db.LastPersistedBlock(ctx)
	if err != nil {
		return nil, LogErrInternal(err)
	}

	latestHeader, err := s.cc.Eth.HeaderByNumber(ctx, lastPersitedBlock) // nil == latest
	if err != nil {
		return nil, LogErrCeloClient("HeaderByNumber", err)
	}

	genesisHeader, err := s.cc.Eth.HeaderByNumber(ctx, big.NewInt(0)) // 0 == genesis
	if err != nil {
		return nil, LogErrCeloClient("HeaderByNumber", err)
	}

	peersInfo, err := s.cc.Admin.Peers(ctx)
	if err != nil {
		return nil, LogErrCeloClient("AdminPeers", err)
	}

	response := types.NetworkStatusResponse{
		CurrentBlockIdentifier: HeaderToBlockIdentifier(latestHeader),
		CurrentBlockTimestamp:  int64(latestHeader.Time * 1000),
		GenesisBlockIdentifier: HeaderToBlockIdentifier(genesisHeader),
		Peers:                  PeersFromInfo(peersInfo),
	}
	return &response, nil
}

// AccountBalance - Get an Account Balance
func (s *Servicer) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {

	accountAddr := common.HexToAddress(request.AccountIdentifier.Address)

	blockHeader, errRsp := s.blockHeader(ctx, request.BlockIdentifier)
	if errRsp != nil {
		return nil, errRsp
	}

	requestedBlockOpts := &bind.CallOpts{
		BlockNumber: blockHeader.Number,
		Context:     ctx,
	}

	emptyResponse := &types.AccountBalanceResponse{
		BlockIdentifier: HeaderToBlockIdentifier(&blockHeader.Header),
	}
	createResponse := func(amount *types.Amount) *types.AccountBalanceResponse {
		return &types.AccountBalanceResponse{
			BlockIdentifier: HeaderToBlockIdentifier(&blockHeader.Header),
			Balances:        []*types.Amount{amount},
		}
	}

	subAccount := request.AccountIdentifier.SubAccount

	if subAccount == nil || subAccount.Address == string(analyzer.AccSigner) {
		// Main Account or Signer => just cGLD
		goldAmt, err := s.cc.Eth.BalanceAt(ctx, accountAddr, blockHeader.Number)
		if err != nil {
			return nil, LogErrCeloClient("BalanceAt", err)
		}
		return createResponse(NewAmount(goldAmt, CeloGold)), nil
	}

	lenRg := len("ReleaseGold")
	if lenRg <= len(subAccount.Address) && subAccount.Address[:lenRg] == "ReleaseGold" {
		releaseGold, err := contracts.NewReleaseGold(accountAddr, s.cc.Eth)
		if err != nil {
			return nil, LogErrValidation(errors.New("Account address must be a ReleaseGold instance"))
		}

		beneficiary, err := releaseGold.Beneficiary(requestedBlockOpts)
		if err != nil {
			return nil, LogErrCeloClient("ReleaseGold.beneficiary", err)
		}

		if subAccount.Metadata != nil {
			if addrStr, ok := subAccount.Metadata["beneficiary"]; ok {
				providedBeneficiary := common.HexToAddress(addrStr.(string))
				if beneficiary != providedBeneficiary {
					return nil, LogErrValidation(errors.New("Provided beneficiary address does not match ReleaseGold contract"))
				}
			}
		}

		var celoGoldAmount *big.Int
		switch subAccount.Address {
		case string(analyzer.AccReleaseGoldVested):
			celoGoldAmount, err = releaseGold.GetCurrentReleasedTotalAmount(requestedBlockOpts)
		case string(analyzer.AccReleaseGoldUnvestedUnLocked):
			celoGoldAmount, err = releaseGold.GetRemainingUnlockedBalance(requestedBlockOpts)
		case string(analyzer.AccReleaseGoldUnvestedLocked):
			celoGoldAmount, err = releaseGold.GetRemainingLockedBalance(requestedBlockOpts)
		default:
			return nil, LogErrValidation(fmt.Errorf("Subaccount must be %s, %s, or %s",
				string(analyzer.AccReleaseGoldVested),
				string(analyzer.AccReleaseGoldUnvestedLocked),
				string(analyzer.AccReleaseGoldUnvestedUnLocked),
			))
		}
		if err != nil {
			return nil, LogErrCeloClient(subAccount.Address, err)
		}
		return createResponse(NewAmount(celoGoldAmount, CeloGold)), nil
	}

	registry, err := registry.New(s.cc)
	if err == client.ErrContractNotDeployed {
		// Nothing is deployed => ignore lockedGold & election balances
		return emptyResponse, nil
	} else if err != nil {
		return nil, LogErrCeloClient("NewRegistry", err)
	}

	if subAccount.Address == string(analyzer.AccLockedGoldNonVoting) {
		// Fetch LockedGold Balances
		lockedGold, err := registry.GetLockedGoldContract(ctx, nil)
		if err == client.ErrContractNotDeployed {
			// Nothing is deployed => ignore lockedGold & election balances
			return emptyResponse, nil
		} else if err != nil {
			return nil, LogErrCeloClient("NewLockedGold", err)
		}

		nonVotingLockedGold, err := lockedGold.GetAccountNonvotingLockedGold(requestedBlockOpts, accountAddr)
		if err != nil {
			return nil, LogErrCeloClient("GetAccountNonvotingLockedGold", err)
		}

		return createResponse(NewAmount(nonVotingLockedGold, CeloGold)), nil
	}

	if subAccount.Address == string(analyzer.AccLockedGoldPending) {
		// Fetch LockedGold Balances
		lockedGold, err := registry.GetLockedGoldContract(ctx, nil)
		if err == client.ErrContractNotDeployed {
			// Nothing is deployed => ignore lockedGold & election balances
			return emptyResponse, nil
		} else if err != nil {
			return nil, LogErrCeloClient("NewLockedGold", err)
		}

		totalPending, err := lockedGold.GetTotalPendingWithdrawals(requestedBlockOpts, accountAddr)
		if err != nil {
			return nil, LogErrCeloClient("GetTotalPendingWithdrawals", err)
		}

		return createResponse(NewAmount(totalPending, CeloGold)), nil
	}

	// If we are here need to be election based

	// Fetch Election (Votes) Balances
	_election, err := registry.GetElectionContract(ctx, nil)
	if err == client.ErrContractNotDeployed {
		// Nothing is deployed => ignore lockedGold & election balances
		return emptyResponse, nil
	} else if err != nil {
		return nil, LogErrCeloClient("NewElection", err)
	}
	election := helpers.Election{_election}

	sumVotes := func(targetVotes helpers.VotesByGroup) *big.Int {
		sum := big.NewInt(0)
		for _, amount := range targetVotes {
			sum.Add(sum, amount)
		}
		return sum
	}

	var groupAddr common.Address
	if subAccount.Metadata != nil {
		if groupStr, ok := subAccount.Metadata["group"]; ok {
			groupAddr = common.HexToAddress(groupStr.(string))
		}
	}

	var voteBalance *big.Int
	if subAccount.Address == string(analyzer.AccLockedGoldVotingPending) {
		if groupAddr != common.ZeroAddress {
			voteBalance, err = election.GetPendingVotesForGroupByAccount(requestedBlockOpts, groupAddr, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetPendingVotesForGroupByAccount", err)
			}
		} else {
			votes, err := election.GetAccountPendingVotes(requestedBlockOpts, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetPendingVotesForGroupByAccount", err)
			}
			voteBalance = sumVotes(votes)
		}
	} else if subAccount.Address == string(analyzer.AccLockedGoldVotingActive) {
		if groupAddr != common.ZeroAddress {
			voteBalance, err = election.GetActiveVotesForGroupByAccount(requestedBlockOpts, groupAddr, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetActiveVotesForGroupByAccount", err)
			}
		} else {
			votes, err := election.GetAccountActiveVotes(requestedBlockOpts, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetActiveVotesForGroupByAccount", err)
			}
			voteBalance = sumVotes(votes)
		}
	} else {
		return nil, LogErrCeloClient("InvalidAccountIdentifier", nil)
	}

	return createResponse(NewAmount(voteBalance, CeloGold)), nil
}

// Block - Get a Block
func (s *Servicer) Block(ctx context.Context, request *types.BlockRequest) (*types.BlockResponse, *types.Error) {

	blockHeader, err := s.blockHeader(ctx, request.BlockIdentifier)
	if err != nil {
		return nil, err
	}

	transactions := MapTxHashesToTransaction(blockHeader.Transactions)

	// If it's the last block of the Epoch, add a transaction for the block Finalize()
	if s.chainParams.IsLastBlockOfEpoch(blockHeader.Number.Uint64()) {
		transactions = append(transactions, &types.TransactionIdentifier{Hash: blockHeader.Hash().Hex()})
	}

	return &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier:       HeaderToBlockIdentifier(&blockHeader.Header),
			ParentBlockIdentifier: HeaderToParentBlockIdentifier(&blockHeader.Header),
			Timestamp:             int64(blockHeader.Time * 1000), // TODO unsafe casting from uint to int 64
		},
		OtherTransactions: transactions,
	}, nil

}

// BlockTransaction - Get a Block Transaction
func (s *Servicer) BlockTransaction(ctx context.Context, request *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {

	blockHeader, err := s.blockHeader(ctx, FullToPartialBlockIdentifier(request.BlockIdentifier))
	if err != nil {
		return nil, err
	}

	txHash := common.HexToHash(request.TransactionIdentifier.Hash)

	var operations []*types.Operation
	// Check If it's block transaction (imaginary transaction)
	if s.chainParams.IsLastBlockOfEpoch(blockHeader.Number.Uint64()) && txHash == blockHeader.Hash() {
		rewards, err := analyzer.ComputeEpochRewards(ctx, s.cc, s.db, &blockHeader.Header)
		if err != nil {
			return nil, LogErrCeloClient("ComputeEpochRewards", err)
		}
		operations = OperationsFromAnalyzer(rewards, 0)
	} else {
		// Normal transaction

		if !HeaderContainsTx(blockHeader, txHash) {
			return nil, LogErrInternal(ErrMissingTxInBlock, "blockNumber", blockHeader.Number, "txHash", txHash.Hex())
		}

		tx, _, err := s.cc.Eth.TransactionByHash(ctx, txHash)
		if err != nil {
			return nil, LogErrCeloClient("TransactionByHash", err)
		}

		receipt, err := s.cc.Eth.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, LogErrCeloClient("TransactionReceipt", err)
		}

		tracer := analyzer.NewTracer(ctx, s.cc, s.db)

		ops, err := tracer.TraceTransaction(&blockHeader.Header, tx, receipt)
		if err != nil {
			return nil, LogErrCeloClient("TraceTransaction", err)
		}

		for _, aop := range ops {
			// TODO - revisit
			// nolint:gosec
			transferOps := OperationsFromAnalyzer(&aop, int64(len(operations)))
			operations = append(operations, transferOps...)
		}
	}

	return &types.BlockTransactionResponse{
		Transaction: &types.Transaction{
			TransactionIdentifier: &types.TransactionIdentifier{Hash: txHash.Hex()},
			Operations:            operations,
		},
	}, nil
}

type CallMethod string

const CeloCall CallMethod = "celo_call"
const CeloGetLogs CallMethod = "celo_getLogs"

func (cm CallMethod) String() string { return string(cm) }

func AllCallMethods() []string {
	return []string{CeloCall.String(), CeloGetLogs.String()}
}

func (s *Servicer) Call(ctx context.Context, request *types.CallRequest) (*types.CallResponse, *types.Error) {

	switch request.Method {
	case CeloCall.String():
		var callParams airgap.CallParams
		if err := airgap.UnmarshallFromMap(request.Parameters, &callParams); err != nil {
			return nil, LogErrValidation(err)
		}

		data, err := s.airgap.CallData(ctx, &callParams)
		if err != nil {
			return nil, LogErrCeloClient(request.Method, err)
		}

		return &types.CallResponse{
			Result: map[string]interface{}{
				"raw": data,
			},
			Idempotent: false,
		}, nil
	case CeloGetLogs.String():
		var filterQueryParams airgap.FilterQueryParams
		if err := airgap.UnmarshallFromMap(request.Parameters, &filterQueryParams); err != nil {
			return nil, LogErrValidation(err)
		}

		logs, err := s.airgap.FilterQuery(ctx, &filterQueryParams)
		if err != nil {
			return nil, LogErrCeloClient(request.Method, err)
		}

		return &types.CallResponse{
			Result: map[string]interface{}{
				"logs": logs,
			},
			Idempotent: true,
		}, nil
	}

	return nil, LogErrValidation(fmt.Errorf("unsupported method '%s'", request.Method))
}

func (s *Servicer) ConstructionCombine(
	context.Context,
	*types.ConstructionCombineRequest,
) (*types.ConstructionCombineResponse, *types.Error) {
	return nil, LogErrUnimplemented("ConstructionCombine")
}

func (s *Servicer) ConstructionDerive(
	context.Context,
	*types.ConstructionDeriveRequest,
) (*types.ConstructionDeriveResponse, *types.Error) {
	return nil, LogErrUnimplemented("ConstructionDerive")
}

func (s *Servicer) ConstructionHash(
	context.Context,
	*types.ConstructionHashRequest,
) (*types.TransactionIdentifierResponse, *types.Error) {
	return nil, LogErrUnimplemented("ConstructionHash")
}

func (s *Servicer) ConstructionParse(
	context.Context,
	*types.ConstructionParseRequest,
) (*types.ConstructionParseResponse, *types.Error) {
	return nil, LogErrUnimplemented("ConstructionParse")
}

func (s *Servicer) ConstructionPayloads(
	context.Context,
	*types.ConstructionPayloadsRequest,
) (*types.ConstructionPayloadsResponse, *types.Error) {
	return nil, LogErrUnimplemented("ConstructionPayloads")
}

func (s *Servicer) ConstructionPreprocess(
	context.Context,
	*types.ConstructionPreprocessRequest,
) (*types.ConstructionPreprocessResponse, *types.Error) {
	return nil, LogErrUnimplemented("ConstructionPreprocess")
}

func (s *Servicer) ConstructionMetadata(ctx context.Context, request *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {

	var txArgs airgap.TxArgs
	if err := airgap.UnmarshallFromMap(request.Options, &txArgs); err != nil {
		return nil, LogErrValidation(err)
	}

	txMetadata, err := s.airgap.ObtainMetadata(ctx, &txArgs)
	if err != nil {
		return nil, LogErrInternal(err)
	}

	metadata, err := airgap.MarshallToMap(txMetadata)
	if err != nil {
		return nil, LogErrInternal(err)
	}

	response := types.ConstructionMetadataResponse{
		Metadata: metadata,
	}
	return &response, nil
}

func (s *Servicer) ConstructionSubmit(ctx context.Context, request *types.ConstructionSubmitRequest) (*types.TransactionIdentifierResponse, *types.Error) {
	rawTx := common.Hex2Bytes(request.SignedTransaction)

	txhash, err := s.airgap.SubmitTx(ctx, rawTx)
	if err != nil {
		return nil, LogErrCeloClient("SendRawTx", err)
	}

	response := types.TransactionIdentifierResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: txhash.String(),
		},
	}
	return &response, nil
}

// ----------------------------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------------------------

func (s *Servicer) blockHeader(ctx context.Context, blockIdentifier *types.PartialBlockIdentifier) (*ethclient.HeaderAndTxnHashes, *types.Error) {
	if blockIdentifier == nil || blockIdentifier.Hash == nil {
		var number *big.Int
		if blockIdentifier != nil && blockIdentifier.Index != nil {
			number = big.NewInt(*blockIdentifier.Index)
		}

		blockHeader, err := s.cc.Eth.HeaderAndTxnHashesByNumber(ctx, number)
		if err != nil {
			return nil, LogErrCeloClient("HeaderAndTxnHashesByNumber", err)
		}
		return blockHeader, nil
	}

	hash := common.HexToHash(*blockIdentifier.Hash)
	blockHeader, err := s.cc.Eth.HeaderAndTxnHashesByHash(ctx, hash)
	if err != nil {
		return nil, LogErrCeloClient("HeaderAndTxnHashesByHash", err)
	}
	// If both Index and Hash were specified check the result matches
	if blockIdentifier.Index != nil && blockHeader.Number.Cmp(big.NewInt(*blockIdentifier.Index)) != 0 {
		return nil, LogErrValidation(ErrBadBlockIdentifier)
	}
	return blockHeader, nil
}
