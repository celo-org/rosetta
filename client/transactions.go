package client

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func NewSigner(chainId *big.Int) types.EIP155Signer {
	return types.NewEIP155Signer(chainId)
}

func SignTransaction(transaction *types.Transaction, privateKey *ecdsa.PrivateKey, signer *types.EIP155Signer) (*types.Transaction, error) {
	return types.SignTx(transaction, signer, privateKey)
}

func ParseSignedTransaction(transaction *types.Transaction, signer *types.EIP155Signer) (*common.Address, *common.Address, *big.Int, error) {
	from, err := signer.Sender(transaction)
	if err != nil {
		return nil, nil, nil, err
	}
	return &from, transaction.To(), transaction.Value(), nil
}

func HashTransaction(transaction *types.Transaction) common.Hash {
	return transaction.Hash()
}
