package client

import (
	"math/big"
	"strconv"

	"github.com/celo-org/rosetta/api"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewSigner(id *api.NetworkIdentifier) types.EIP155Signer {
	chainId, err := strconv.ParseInt(id.Network, 0, 64)
	if err != nil {
		// TODO: err
	}
	return newSigner(big.NewInt(chainId))
}

// func SignTransaction(tx *api.TransactionConstructionResponse) {
// 	v, ok = i.(T)
// 	switch tx.Metadata
// }
