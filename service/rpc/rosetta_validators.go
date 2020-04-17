package rpc

import (
	"fmt"

	"github.com/celo-org/rosetta/celo"
	"github.com/coinbase/rosetta-sdk-go/types"
)

func ValidateNetworkId(id *types.NetworkIdentifier, cp *celo.ChainParameters) *types.Error {
	if id.Blockchain != BlockchainName {
		return LogErrValidation(fmt.Errorf("Expected blockchain id %s to be %s", id.Blockchain, BlockchainName))
	}

	if cp.ChainId.String() != id.Network {
		return LogErrValidation(fmt.Errorf("Expected network id %s to be %s", id.Network, cp.ChainId.String()))
	}

	return nil
}
