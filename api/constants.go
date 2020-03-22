package api

import (
	"github.com/ethereum/go-ethereum/params"
)

var (
	NodeVersion = params.Version
)

// TODO(yorke): consider reading version info from disk
const (
	RosettaVersion    = "1.2.3"
	MiddlewareVersion = "1.0.0"
	BlockchainName    = "celo"
)
