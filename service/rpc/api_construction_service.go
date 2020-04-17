package rpc

import (
	"context"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
)

// ConstructionApiService is a service that implents the logic for the ConstructionApiServicer
// This service should implement the business logic for every endpoint for the ConstructionApi API.
// Include any external packages or services that will be required by this service.
type ConstructionApiService struct {
	celoClient  *client.CeloClient
	chainParams *celo.ChainParameters
}

// NewConstructionApiService creates a default api service
func NewConstructionApiService(celoClient *client.CeloClient, chainParams *celo.ChainParameters) server.ConstructionAPIServicer {
	return &ConstructionApiService{
		celoClient:  celoClient,
		chainParams: chainParams,
	}
}

func (s *ConstructionApiService) getTxMetadata(ctx context.Context, address common.Address) (*TransactionMetadata, *types.Error) {
	var txMetadata TransactionMetadata
	var err error

	txMetadata.Nonce, err = s.celoClient.Eth.NonceAt(ctx, address, nil) // nil == latest
	if err != nil {
		return nil, NewCeloClientError("NonceAt", err)
	}

	txMetadata.GatewayFeeRecipient, err = s.celoClient.Eth.Coinbase(ctx)
	if err != nil {
		return nil, NewCeloClientError("Coinbase", err)
	}

	txMetadata.GasPrice, err = s.celoClient.Eth.SuggestGasPrice(ctx)
	if err != nil {
		return nil, NewCeloClientError("SuggestGasPrice", err)
	}

	// TODO: consider fetching from node
	txMetadata.GatewayFee = big.NewInt(0)

	return &txMetadata, nil
}

func (s *ConstructionApiService) ConstructionMetadata(ctx context.Context, txConstructionRequest *types.ConstructionMetadataRequest) (*types.ConstructionMetadataResponse, *types.Error) {
	err := ValidateNetworkId(txConstructionRequest.NetworkIdentifier, s.chainParams)
	if err != nil {
		return nil, err
	}

	options := *txConstructionRequest.Options
	account := fmt.Sprintf("%v", options["account"])
	address := common.HexToAddress(account)

	var metadata = make(map[string]interface{})

	txMetadata, err := s.getTxMetadata(ctx, address)
	if err != nil {
		return nil, err
	}

	metadata["general"] = txMetadata

	// switch txConstructionRequest.Method {
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

func (s *ConstructionApiService) ConstructionSubmit(ctx context.Context, constructionSubmitRequest *types.ConstructionSubmitRequest) (*types.ConstructionSubmitResponse, *types.Error) {
	txhash, err := s.celoClient.Eth.SendRawTransaction(ctx, []byte(constructionSubmitRequest.SignedTransaction))
	if err != nil {
		return nil, NewCeloClientError("SendRawTx", err)
	}

	response := types.ConstructionSubmitResponse{
		TransactionIdentifier: &types.TransactionIdentifier{
			Hash: txhash.String(),
		},
	}
	return &response, nil
}
