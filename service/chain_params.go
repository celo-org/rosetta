package service

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/consensus/istanbul"
	"github.com/celo-org/celo-blockchain/params"
)

type ChainParameters struct {
	ChainId       *big.Int
	EpochSize     uint64
	IsGingerbread func(*big.Int) bool
}

func NewChainParametersFromConfig(config *params.ChainConfig) *ChainParameters {
	return &ChainParameters{
		ChainId:       config.ChainID,
		EpochSize:     config.Istanbul.Epoch,
		IsGingerbread: config.IsGingerbread,
	}
}

func (cp *ChainParameters) IsLastBlockOfEpoch(blockNumber uint64) bool {
	return istanbul.IsLastBlockOfEpoch(blockNumber, cp.EpochSize)
}
