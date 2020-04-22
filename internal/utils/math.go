package utils

import "math/big"

var Big0 = big.NewInt(0)
var Big1 = big.NewInt(1)

func Sum(values ...*big.Int) *big.Int {
	acc := new(big.Int)
	for _, val := range values {
		acc.Add(acc, val)
	}
	return acc
}

func Inc(val *big.Int) *big.Int {
	return new(big.Int).Add(val, Big1)
}

func Dec(val *big.Int) *big.Int {
	return new(big.Int).Sub(val, Big1)
}
