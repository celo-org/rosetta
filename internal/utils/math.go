package utils

import (
	"math/big"
)

var Big0 = big.NewInt(0)

func Sum(values ...*big.Int) *big.Int {
	acc := new(big.Int)
	for _, val := range values {
		acc.Add(acc, val)
	}
	return acc
}
