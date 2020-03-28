package celo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

type ChainParameters struct {
	ChainId   *big.Int
	EpochSize uint64
}

func (cp *ChainParameters) IsLastBlockOfEpoch(blockNumber uint64) bool {
	return istanbul.IsLastBlockOfEpoch(blockNumber, cp.EpochSize)
}
