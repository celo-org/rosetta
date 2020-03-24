package api

import (
	"context"
	"errors"
	"fmt"

	network "github.com/celo-org/rosetta/celo/client/network"
)

func ValidateNetworkId(id *NetworkIdentifier, client *network.NetworkClient, ctx context.Context) error {
	if id.Blockchain != BlockchainName {
		return errors.New(fmt.Sprintf("Mismatched blockchain identifiers: %s != %s", id.Blockchain, BlockchainName))
	}

	chainId, err := client.ChainId(ctx)
	if err != nil {
		return err
	}

	networkName := NetworkNameFromId(chainId)
	if networkName != id.Network {
		return errors.New(fmt.Sprintf("Mismatched network identifiers: %s != %s", id.Network, networkName))
	}
	return nil
}
