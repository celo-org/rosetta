package rpc

import (
	"context"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// MempoolApiService is a service that implents the logic for the MempoolApiServicer
// This service should implement the business logic for every endpoint for the MempoolApi API.
// Include any external packages or services that will be required by this service.
type MempoolApiService struct {
	celoClient  *client.CeloClient
	chainParams *celo.ChainParameters
}

// NewMempoolApiService creates a default api service
func NewMempoolApiService(celoClient *client.CeloClient, chainParams *celo.ChainParameters) server.MempoolAPIServicer {
	return &MempoolApiService{
		celoClient:  celoClient,
		chainParams: chainParams,
	}
}

// Mempool - Get All Mempool Transactions
func (m *MempoolApiService) Mempool(ctx context.Context, mempoolRequest *types.MempoolRequest) (*types.MempoolResponse, *types.Error) {

	errResp := ValidateNetworkId(mempoolRequest.NetworkIdentifier, m.chainParams)
	if errResp != nil {
		return nil, errResp
	}

	content, err := m.celoClient.TxPool.Content(ctx)
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
func (m *MempoolApiService) MempoolTransaction(ctx context.Context, mempoolTransactionRequest *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	return nil, LogErrUnimplemented("/mempool/transaction")
}
