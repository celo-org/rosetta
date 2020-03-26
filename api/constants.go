package api

import (
	"github.com/ethereum/go-ethereum/params"
)

// VERSIONING
const (
	RosettaVersion    = "1.2.4"
	MiddlewareVersion = "1.0.0"
	BlockchainName    = "celo"
)

var (
	NodeVersion      = params.Version
	ChainIdToNetwork = map[uint64]string{
		44786:  "alfajores",
		200110: "baklava",
		200312: "rc0",
	}
)

// API
const (
	TransferMethod string = "transfer"
)

var (
	GasUpperBound = map[string]string{
		TransferMethod: "40000",
	}
	CeloGold = Currency{
		Symbol:   "cGLD",
		Decimals: 18,
	}
	CeloDollar = Currency{
		Symbol:   "cUSD",
		Decimals: 18,
	}
)
