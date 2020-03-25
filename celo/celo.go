package celo

import (
	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
)

func GetRegistry(eth *ethclient.Client) (*contract.Registry, error) {
	return contract.NewRegistry(params.RegistrySmartContractAddress, eth)
}
