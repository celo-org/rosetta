package analyzer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type BalanceSet struct {
	balances map[common.Address]*big.Int
}

func NewBalanceSet() *BalanceSet {
	return &BalanceSet{
		balances: make(map[common.Address]*big.Int),
	}
}

func (bs *BalanceSet) Add(addr common.Address, value *big.Int) *BalanceSet {
	if oldValue, ok := bs.balances[addr]; ok {
		bs.balances[addr] = new(big.Int).Add(oldValue, value)
	} else {
		bs.balances[addr] = value
	}
	return bs
}

func (bs *BalanceSet) ToMap() map[common.Address]*big.Int {
	return bs.balances
}
