package rpc

import "github.com/celo-org/rosetta/service"

var _ service.Service = (*rosettaServer)(nil)
