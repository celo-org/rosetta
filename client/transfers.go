package client

import (
	"math/big"

	"github.com/celo-org/rosetta/service/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func ConstructTransfer(from *common.Address, to *common.Address, amount *big.Int, metadata *rpc.TransactionMetadata) *types.Transaction {
	return types.NewTransaction(
		metadata.Nonce,
		*to,
		amount,
		metadata.GasLimit,
		metadata.GasPrice,
		nil, // guarantee cGLD
		metadata.GatewayFeeRecipient,
		metadata.GatewayFee,
		[]byte{},
	)
}
