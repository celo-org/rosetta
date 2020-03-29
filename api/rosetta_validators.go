package api

import (
	"fmt"

	"github.com/celo-org/rosetta/celo"
)

func ValidateNetworkId(id *NetworkIdentifier, cp *celo.ChainParameters) error {
	if id.Blockchain != BlockchainName {
		return ErrBadNetworkIdentifier(fmt.Errorf("Mismatched blockchain identifiers: %s != %s", id.Blockchain, BlockchainName))
	}

	if cp.NetworkName() != id.Network {
		return ErrBadNetworkIdentifier(fmt.Errorf("Mismatched network identifiers: %s != %s", id.Network, cp.NetworkName()))
	}
	return nil
}
