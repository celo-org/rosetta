package rpc

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/analyzer"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/db"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// BlockApiService is a service that implents the logic for the BlockApiServicer
// This service should implement the business logic for every endpoint for the BlockApi API.
// Include any external packages or services that will be required by this service.
type BlockApiService struct {
	celoClient  *client.CeloClient
	db          db.RosettaDBReader
	chainParams *celo.ChainParameters
}

// NewBlockApiService creates a default api service
func NewBlockApiService(celoClient *client.CeloClient, db db.RosettaDBReader, cp *celo.ChainParameters) server.BlockAPIServicer {
	return &BlockApiService{
		celoClient:  celoClient,
		db:          db,
		chainParams: cp,
	}
}

func (b *BlockApiService) BlockHeader(ctx context.Context, blockIdentifier *types.PartialBlockIdentifier) (*ethclient.HeaderAndTxnHashes, *types.Error) {
	var err error
	var blockHeader *ethclient.HeaderAndTxnHashes

	if blockIdentifier.Hash != nil {
		hash := common.HexToHash(*blockIdentifier.Hash)
		blockHeader, err = b.celoClient.Eth.HeaderAndTxnHashesByHash(ctx, hash)
		if err != nil {
			return nil, LogErrFetchBlockHeader(err)
		}

		// If both were specified check the result matches
		if blockIdentifier.Index != nil && blockHeader.Number.Cmp(big.NewInt(*blockIdentifier.Index)) != 0 {
			return nil, LogErrValidation(ErrBadBlockIdentifier)
		}

	} else if blockIdentifier.Index != nil {
		blockHeader, err = b.celoClient.Eth.HeaderAndTxnHashesByNumber(ctx, big.NewInt(*blockIdentifier.Index))
		if err != nil {
			return nil, LogErrFetchBlockHeader(err)
		}
	} else {
		blockHeader, err = b.celoClient.Eth.HeaderAndTxnHashesByNumber(ctx, nil)
		if err != nil {
			return nil, LogErrFetchBlockHeader(err)
		}
	}

	return blockHeader, nil

}

// Block - Get a Block
func (b *BlockApiService) Block(ctx context.Context, blockRequest *types.BlockRequest) (*types.BlockResponse, *types.Error) {

	err := ValidateNetworkId(blockRequest.NetworkIdentifier, b.chainParams)
	if err != nil {
		return nil, err
	}

	blockHeader, err := b.BlockHeader(ctx, blockRequest.BlockIdentifier)
	if err != nil {
		return nil, err
	}

	transactions := MapTxHashesToTransaction(blockHeader.Transactions)

	// If it's the last block of the Epoch, add a transaction for the block Finalize()
	if b.chainParams.IsLastBlockOfEpoch(blockHeader.Number.Uint64()) {
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
func (s *BlockApiService) BlockTransaction(ctx context.Context, blockTransactionRequest *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {

	err := ValidateNetworkId(blockTransactionRequest.NetworkIdentifier, s.chainParams)
	if err != nil {
		return nil, err
	}

	blockHeader, err := s.BlockHeader(ctx, FullToPartialBlockIdentifier(blockTransactionRequest.BlockIdentifier))
	if err != nil {
		return nil, err
	}

	txHash := common.HexToHash(blockTransactionRequest.BlockIdentifier.Hash)

	var operations []*types.Operation
	// Check If it's block transaction (imaginary transaction)
	if s.chainParams.IsLastBlockOfEpoch(blockHeader.Number.Uint64()) && txHash == blockHeader.Hash() {
		rewards, err := analyzer.ComputeEpochRewards(ctx, s.celoClient, s.db, &blockHeader.Header)
		if err != nil {
			return nil, LogErrCeloClient("ComputeEpochRewards", err)
		}
		operations = OperationsFromAnalyzer(rewards, 0)
	} else {
		// Normal transaction

		if !HeaderContainsTx(blockHeader, txHash) {
			return nil, LogErrInternal(ErrMissingTxInBlock)
		}

		tx, _, err := s.celoClient.Eth.TransactionByHash(ctx, txHash)
		if err != nil {
			return nil, LogErrCeloClient("TransactionByHash", err)
		}

		receipt, err := s.celoClient.Eth.TransactionReceipt(ctx, tx.Hash())
		if err != nil {
			return nil, LogErrCeloClient("TransactionReceipt", err)
		}

		tracer := analyzer.NewTracer(ctx, s.celoClient, s.db)

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
