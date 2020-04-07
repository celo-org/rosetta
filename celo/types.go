package celo

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
)

type ChainParameters struct {
	ChainId   *big.Int
	EpochSize uint64
}

func (cp *ChainParameters) IsLastBlockOfEpoch(blockNumber uint64) bool {
	return istanbul.IsLastBlockOfEpoch(blockNumber, cp.EpochSize)
}

type SubAccount string

const (
	Main              SubAccount = "Main"
	LockedGoldLocked  SubAccount = "LockedGoldLocked"
	LockedGoldPending SubAccount = "LockedGoldPending"
)

type Account struct {
	Address    common.Address
	SubAccount SubAccount
}

type Transfer struct {
	From   Account
	To     Account
	Value  *big.Int
	Status bool
}
