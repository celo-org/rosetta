package celo

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

var ChainIdToNetwork = map[uint64]string{
	44786:  "alfajores",
	200110: "baklava",
	200312: "rc0",
}

type ChainParameters struct {
	ChainId   *big.Int
	EpochSize uint64
}

func (cp *ChainParameters) IsLastBlockOfEpoch(blockNumber uint64) bool {
	return istanbul.IsLastBlockOfEpoch(blockNumber, cp.EpochSize)
}

func (cp *ChainParameters) NetworkName() string {
	uid := cp.ChainId.Uint64()
	if name, ok := ChainIdToNetwork[uid]; ok {
		return name
	}
	return fmt.Sprintf("unknown %d", uid)
}
