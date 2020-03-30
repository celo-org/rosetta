package client

import (
	"crypto/ecdsa"

	"github.com/celo-org/rosetta/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func EncodeTransaction(transaction *types.Transaction) (*api.TransactionSubmitRequest, error) {
	bytes, err := rlp.EncodeToBytes(transaction)
	if err != nil {
		return nil, err
	}
	return &api.TransactionSubmitRequest{
		SignedTransaction: string(bytes),
	}, nil
}

func DecodeTransaction(transaction *api.TransactionSubmitRequest) (*types.Transaction, error) {
	bytes := []byte(transaction.SignedTransaction)
	var tx types.Transaction
	err := rlp.DecodeBytes(bytes, tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func EncodeAccount(publicKey *ecdsa.PublicKey, address *common.Address) *api.AccountIdentifier {
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return &api.AccountIdentifier{
		Address: address.String(),
		Metadata: map[string]interface{}{
			"publicKey": string(publicKeyBytes),
		},
	}
}
