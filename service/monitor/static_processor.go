// Copyright 2023 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package monitor

import (
	"context"
	"math/big"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/log"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/kliento/client"
	"github.com/celo-org/rosetta/kliento/contracts"
	"github.com/celo-org/rosetta/kliento/registry"
)

type staticProcessor struct {
	callOpts      *bind.CallOpts
	boundRegistry *contracts.Registry
	cc            *client.CeloClient
	logger        log.Logger
}

func newStaticProcessor(
	ctx context.Context,
	blockNumber *big.Int,
	cc *client.CeloClient,
	logger log.Logger,
) (*staticProcessor, error) {
	boundRegistry, err := contracts.NewRegistry(registry.RegistryAddress, cc.Eth)
	if err != nil {
		return nil, err
	}
	callOpts := &bind.CallOpts{
		Context:     ctx,
		BlockNumber: blockNumber,
	}
	return &staticProcessor{
		callOpts:      callOpts,
		boundRegistry: boundRegistry,
		logger:        logger.New("pipe", "static_processor"),
		cc:            cc,
	}, nil
}

func (sf *staticProcessor) fetchContractAddresses() ([]db.RegistryChange, error) {
	contractAddresses := []db.RegistryChange{}

	for _, contractId := range registry.RegisteredContractIDs {
		contractStrName := contractId.String()
		contractAddress, err := sf.boundRegistry.GetAddressForString(sf.callOpts, contractStrName)
		if err != nil {
			return nil, err
		}
		contractAddresses = append(contractAddresses, db.RegistryChange{
			TxIndex:    0,
			Contract:   contractStrName,
			NewAddress: contractAddress,
		})
		sf.logger.Info("Found Core Contract address", "name", contractStrName, "newAddress", contractAddress.Hex())
	}
	return contractAddresses, nil
}

func (sf *staticProcessor) fetchGasPriceMinimum() (*big.Int, error) {
	gpmAddress, err := sf.boundRegistry.GetAddressForString(sf.callOpts, registry.GasPriceMinimumContractID.String())
	if err != nil || gpmAddress == common.ZeroAddress {
		return nil, err
	}
	gpmContract, err := contracts.NewGasPriceMinimum(gpmAddress, sf.cc.Eth)
	if err != nil {
		return nil, err
	}
	return gpmContract.GetGasPriceMinimum(sf.callOpts, common.ZeroAddress)
}

// Fetches relevant state from block and updates the Rosetta DB accordingly.
// This is necessary for running Rosetta on networks that have Core Contracts
// already deployed at genesis (e.g. myCelo networks).
func UpdateDBForBlock(
	ctx context.Context,
	blockNumber *big.Int,
	cc *client.CeloClient,
	rosettaDB db.RosettaDB,
	logger log.Logger,
) error {
	sf, err := newStaticProcessor(ctx, blockNumber, cc, logger)
	if err != nil {
		return err
	}
	sf.logger.Info("Fetching contract state", "block", blockNumber)
	contractAddresses, err := sf.fetchContractAddresses()
	if err != nil {
		return err
	}
	gpm, err := sf.fetchGasPriceMinimum()
	if err != nil {
		return err
	}
	sf.logger.Info("Found GasPriceMinimum", "block", blockNumber, "GPM", gpm)

	return rosettaDB.ApplyChanges(ctx, &db.BlockChangeSet{
		BlockNumber:     blockNumber,
		RegistryChanges: contractAddresses,
		GasPriceMinimum: gpm,
	})
}
