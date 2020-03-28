package api

import (
	"fmt"

	"github.com/celo-org/rosetta/internal/config"
)

func ValidateNetworkId(id *NetworkIdentifier) error {
	if id.Blockchain != BlockchainName {
		return ErrBadNetworkIdentifier(fmt.Errorf("Mismatched blockchain identifiers: %s != %s", id.Blockchain, BlockchainName))
	}

	networkName := NetworkNameFromId(config.Chain.ChainId)
	if networkName != id.Network {
		return ErrBadNetworkIdentifier(fmt.Errorf("Mismatched network identifiers: %s != %s", id.Network, networkName))
	}
	return nil
}
