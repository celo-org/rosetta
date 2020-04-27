package utils

import (
	"errors"
	"math/big"
)

var ErrNotImplemented = errors.New("Not implemented")

// 10^(24) == 1000000000000000000000000
var TobinTaxDenominator = new(big.Int).SetBytes([]byte{211, 194, 27, 206, 204, 237, 161, 0, 0, 0})

func CalcTobinTaxAmount(transferAmount, tobinTaxNumerator *big.Int) *big.Int {
	return new(big.Int).Div(new(big.Int).Mul(transferAmount, tobinTaxNumerator), TobinTaxDenominator)
}
