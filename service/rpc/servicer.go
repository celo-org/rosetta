package rpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/analyzer"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
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
	chainParams *celo.ChainParameters
}

// NewServicer creates a default api service
func NewServicer(celoClient *client.CeloClient, db db.RosettaDBReader, cp *celo.ChainParameters) *Servicer {
	return &Servicer{
		cc:          celoClient,
		db:          db,
		chainParams: cp,
	}
}

// Mempool - Get All Mempool Transactions
func (s *Servicer) Mempool(ctx context.Context, request *types.MempoolRequest) (*types.MempoolResponse, *types.Error) {
	if errResp := s.validateNetworkId(request.NetworkIdentifier); errResp != nil {
		return nil, errResp
	}

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
		CurrentBlockTimestamp:  int64(latestHeader.Time),
		GenesisBlockIdentifier: HeaderToBlockIdentifier(genesisHeader),
		Peers:                  PeersFromInfo(peersInfo),
	}
	return &response, nil
}

// AccountBalance - Get an Account Balance
func (s *Servicer) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	if errResp := s.validateNetworkId(request.NetworkIdentifier); errResp != nil {
		return nil, errResp
	}

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

	if subAccount == nil {
		// Main Account, just cGLD
		goldAmt, err := s.cc.Eth.BalanceAt(ctx, accountAddr, blockHeader.Number)
		if err != nil {
			return nil, LogErrCeloClient("BalanceAt", err)
		}
		return createResponse(NewAmount(goldAmt, CeloGold)), nil
	}

	registryWrapper, err := wrapper.NewRegistry(s.cc)
	if err == client.ErrContractNotDeployed {
		// Nothing is deployed => ignore lockedGold & election balances
		return emptyResponse, nil
	} else if err != nil {
		return nil, LogErrCeloClient("NewRegistry", err)
	}

	if subAccount.Address == string(analyzer.AccLockedGoldNonVoting) {
		// Fetch LockedGold Balances
		lockedGoldWrapper, err := wrapper.NewLockedGold(s.cc, registryWrapper)
		if err == client.ErrContractNotDeployed {
			// Nothing is deployed => ignore lockedGold & election balances
			return emptyResponse, nil
		} else if err != nil {
			return nil, LogErrCeloClient("NewLockedGold", err)
		}

		nonVotingLockedGold, err := lockedGoldWrapper.GetAccountNonvotingLockedGold(requestedBlockOpts, accountAddr)
		if err != nil {
			return nil, LogErrCeloClient("GetAccountNonvotingLockedGold", err)
		}

		return createResponse(NewAmount(nonVotingLockedGold, CeloGold)), nil
	}

	if subAccount.Address == string(analyzer.AccLockedGoldPending) {
		// Fetch LockedGold Balances
		lockedGoldWrapper, err := wrapper.NewLockedGold(s.cc, registryWrapper)
		if err == client.ErrContractNotDeployed {
			// Nothing is deployed => ignore lockedGold & election balances
			return emptyResponse, nil
		} else if err != nil {
			return nil, LogErrCeloClient("NewLockedGold", err)
		}

		totalPending, err := lockedGoldWrapper.GetTotalPendingWithdrawals(requestedBlockOpts, accountAddr)
		if err != nil {
			return nil, LogErrCeloClient("GetTotalPendingWithdrawals", err)
		}

		return createResponse(NewAmount(totalPending, CeloGold)), nil
	}

	// If we are here need to be election based

	// Fetch Election (Votes) Balances
	electionWrapper, err := wrapper.NewElection(s.cc, registryWrapper)
	if err == client.ErrContractNotDeployed {
		// Nothing is deployed => ignore lockedGold & election balances
		return emptyResponse, nil
	} else if err != nil {
		return nil, LogErrCeloClient("NewElection", err)
	}

	sumVotes := func(targetVotes wrapper.VotesByGroup) *big.Int {
		sum := big.NewInt(0)
		for _, amount := range targetVotes {
			sum.Add(sum, amount)
		}
		return sum
	}

	var groupAddr common.Address
	if subAccount.Metadata != nil {
		if groupStr, ok := (*subAccount.Metadata)["group"]; ok {
			groupAddr = common.HexToAddress(groupStr.(string))
		}
	}

	voteBalance := big.NewInt(0)
	if subAccount.Address == string(analyzer.AccLockedGoldVotingPending) {
		if groupAddr != common.ZeroAddress {
			voteBalance, err = electionWrapper.GetPendingVotesForGroupByAccount(requestedBlockOpts, groupAddr, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetPendingVotesForGroupByAccount", err)
			}
		} else {
			votes, err := electionWrapper.GetAccountPendingVotes(requestedBlockOpts, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetPendingVotesForGroupByAccount", err)
			}
			voteBalance = sumVotes(votes)
		}
	} else if subAccount.Address == string(analyzer.AccLockedGoldVotingActive) {
		if groupAddr != common.ZeroAddress {
			voteBalance, err = electionWrapper.GetActiveVotesForGroupByAccount(requestedBlockOpts, groupAddr, accountAddr)
			if err != nil {
				return nil, LogErrCeloClient("GetActiveVotesForGroupByAccount", err)
			}
		} else {
			votes, err := electionWrapper.GetAccountActiveVotes(requestedBlockOpts, accountAddr)
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
	if errResp := s.validateNetworkId(request.NetworkIdentifier); errResp != nil {
		return nil, errResp
	}

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
			Timestamp:             int64(blockHeader.Time), // TODO unsafe casting from uint to int 64
		},
		OtherTransactions: transactions,
	}, nil

}

// BlockTransaction - Get a Block Transaction
func (s *Servicer) BlockTransaction(ctx context.Context, request *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	if errResp := s.validateNetworkId(request.NetworkIdentifier); errResp != nil {
		return nil, errResp
	}

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

func (s *Servicer) getTxMetadata(ctx context.Context, address common.Address) (*TransactionMetadata, *types.Error) {
	var txMetadata TransactionMetadata
	var err error

	txMetadata.Nonce, err = s.cc.Eth.NonceAt(ctx, address, nil) // nil == latest
	if err != nil {
		return nil, LogErrCeloClient("NonceAt", err)
	}

	txMetadata.GatewayFeeRecipient, err = s.cc.Eth.Coinbase(ctx)
	if err != nil {
		return nil, LogErrCeloClient("Coinbase", err)
	}

	txMetadata.GasPrice, err = s.cc.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, LogErrCeloClient("SuggestGasPrice", err)
	}

	// TODO: consider fetching from node
	txMetadata.GatewayFee = big.NewInt(0)

	return &txMetadata, nil
}

func (s *Servicer) ConstructionMetadata(ctx context.Context, request *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	if errResp := s.validateNetworkId(request.NetworkIdentifier); errResp != nil {
		return nil, errResp
	}

	options := *request.Options
	account := fmt.Sprintf("%v", options["account"])
	address := common.HexToAddress(account)

	var metadata = make(map[string]interface{})

	txMetadata, err := s.getTxMetadata(ctx, address)
	if err != nil {
		return nil, err
	}

	metadata["general"] = txMetadata

	// switch request.Method {
	// case TransferMethod:
	// 	balance, err := s.celoClient.Eth.BalanceAt(ctx, address, nil) // nil == latest
	// 	if err != nil {
	// 		return nil, WrapError("Transfer: BalanceAt", err)
	// 	}

	// 	msg := txMetadata.asMessage()
	// 	msg.Value = balance
	// 	msg.To = &DummyAddress
	// 	gasLimit, err := s.celoClient.Eth.EstimateGas(ctx, *msg)
	// 	if err != nil {
	// 		return nil, WrapError("Transfer: EstimateGas", err)
	// 	}

	// 	txMetadata.GasLimit = gasLimit

	// 	metadata[TransferMethod] = TransferMetadata{
	// 		Balance: balance,
	// 		Tx:      txMetadata,
	// 	}

	// default:
	// 	return nil, WrapError("Unknown method", err)
	// }

	response := types.ConstructionMetadataResponse{
		Metadata: &metadata,
	}
	return &response, nil
}

func (s *Servicer) ConstructionSubmit(ctx context.Context, request *types.ConstructionSubmitRequest) (*types.ConstructionSubmitResponse, *types.Error) {
	txhash, err := s.cc.Eth.SendRawTransaction(ctx, []byte(request.SignedTransaction))
	if err != nil {
		return nil, LogErrCeloClient("SendRawTx", err)
	}

	response := types.ConstructionSubmitResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: txhash.String(),
		},
	}
	return &response, nil
}

// ----------------------------------------------------------------------------------------
// Private Functions
// ----------------------------------------------------------------------------------------

func (s *Servicer) validateNetworkId(id *types.NetworkIdentifier) *types.Error {
	if id.Blockchain != BlockchainName {
		return LogErrValidation(fmt.Errorf("Expected blockchain id %s to be %s", id.Blockchain, BlockchainName))
	}

	if s.chainParams.ChainId.String() != id.Network {
		return LogErrValidation(fmt.Errorf("Expected network id %s to be %s", id.Network, s.chainParams.ChainId.String()))
	}

	return nil
}

func (s *Servicer) blockHeader(ctx context.Context, blockIdentifier *types.PartialBlockIdentifier) (*ethclient.HeaderAndTxnHashes, *types.Error) {
	var err error
	var blockHeader *ethclient.HeaderAndTxnHashes

	if blockIdentifier.Hash != nil {
		hash := common.HexToHash(*blockIdentifier.Hash)
		blockHeader, err = s.cc.Eth.HeaderAndTxnHashesByHash(ctx, hash)
		if err != nil {
			return nil, LogErrFetchBlockHeader(err)
		}

		// If both were specified check the result matches
		if blockIdentifier.Index != nil && blockHeader.Number.Cmp(big.NewInt(*blockIdentifier.Index)) != 0 {
			return nil, LogErrValidation(ErrBadBlockIdentifier)
		}

	} else if blockIdentifier.Index != nil {
		blockHeader, err = s.cc.Eth.HeaderAndTxnHashesByNumber(ctx, big.NewInt(*blockIdentifier.Index))
		if err != nil {
			return nil, LogErrFetchBlockHeader(err)
		}
	} else {
		blockHeader, err = s.cc.Eth.HeaderAndTxnHashesByNumber(ctx, nil)
		if err != nil {
			return nil, LogErrFetchBlockHeader(err)
		}
	}

	return blockHeader, nil

}
