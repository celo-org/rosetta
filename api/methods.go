package api

import (
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

var (
	CeloGold = Currency{
		Symbol:   "cGLD",
		Decimals: 18,
	}
	CeloDollar = Currency{
		Symbol:   "cUSD",
		Decimals: 18,
	}
)

type Method = string

//go:generate gencodec -type TransferMetadata -out gen_transfer_json.go

type TransferMetadata struct {
	Balance *big.Int      `json:"balance" gencodec:"required"`
	Txdata  *types.Txdata `json:"txdata"`
}

const (
	TransferMethod Method = "transfer"
)

var (
	GasUpperBound = map[Method]string{
		TransferMethod: "40000",
	}
)
