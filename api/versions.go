package api

import (
	"github.com/ethereum/go-ethereum/params"
)

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
