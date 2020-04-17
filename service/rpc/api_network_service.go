package rpc

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/analyzer"
	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)

// NetworkApiService is a service that implents the logic for the NetworkApiServicer
// This service should implement the business logic for every endpoint for the NetworkApi API.
// Include any external packages or services that will be required by this service.
type NetworkApiService struct {
	celoClient  *client.CeloClient
	chainParams *celo.ChainParameters
}

// NewNetworkApiService creates a default api service
func NewNetworkApiService(celoClient *client.CeloClient, chainParams *celo.ChainParameters) server.NetworkAPIServicer {
	return &NetworkApiService{
		celoClient:  celoClient,
		chainParams: chainParams,
	}
}

func (s *NetworkApiService) NetworkList(ctx context.Context, networkStatusRequest *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	response := types.NetworkListResponse{
		NetworkIdentifiers: []*types.NetworkIdentifier{{
			Blockchain: BlockchainName,
			Network:    s.chainParams.ChainId.String(),
		}},
	}

	return &response, nil
}

func (s *NetworkApiService) NetworkOptions(ctx context.Context, networkRequest *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
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
			Errors:         []*types.Error{
				// TODO: fill
			},
		},
	}

	return &response, nil
}

// NetworkStatus - Get Network Status
func (s *NetworkApiService) NetworkStatus(ctx context.Context, networkStatusRequest *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {

	latestHeader, err := s.celoClient.Eth.HeaderByNumber(ctx, nil) // nil == latest
	if err != nil {
		return nil, NewCeloClientError("HeaderByNumber", err)
	}

	genesisHeader, err := s.celoClient.Eth.HeaderByNumber(ctx, big.NewInt(0)) // 0 == genesis
	if err != nil {
		return nil, NewCeloClientError("HeaderByNumber", err)
	}

	peersInfo, err := s.celoClient.Admin.Peers(ctx)
	if err != nil {
		return nil, NewCeloClientError("AdminPeers", err)
	}

	response := types.NetworkStatusResponse{
		CurrentBlockIdentifier: HeaderToBlockIdentifier(latestHeader),
		CurrentBlockTimestamp:  int64(latestHeader.Time),
		GenesisBlockIdentifier: HeaderToBlockIdentifier(genesisHeader),
		Peers:                  PeersFromInfo(peersInfo),
	}
	return &response, nil
}
