package client

import (
	"math/big"

	"github.com/celo-org/rosetta/service/rpc"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func ConstructCreateAccount(from *common.Address, metadata *rpc.TransactionMetadata) (*types.Transaction, error) {
	args, err := metadata.ABIMethod.Inputs.Pack()
	if err != nil {
		return nil, err
	}

	data := append(metadata.ABIMethod.ID(), args...)

	return types.NewTransaction(
		metadata.Nonce,
		*metadata.To,
		big.NewInt(0),
		metadata.GasLimit,
		metadata.GasPrice,
		nil, // guarantee cGLD
		metadata.GatewayFeeRecipient,
		metadata.GatewayFee,
		data,
	), nil
}
