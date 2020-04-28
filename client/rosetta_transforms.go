package client

import (
	"crypto/ecdsa"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func NewAccountIdentifier(addr *common.Address) *types.AccountIdentifier {
	return &types.AccountIdentifier{
		Address: addr.Hex(),
	}
}

func EncodeTransaction(transaction *types.Transaction) (string, error) {
	bytes, err := rlp.EncodeToBytes(transaction)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func DecodeTransaction(transaction *types.ConstructionSubmitRequest) (*types.Transaction, error) {
	bytes := []byte(transaction.SignedTransaction)
	var tx types.Transaction
	err := rlp.DecodeBytes(bytes, tx)
	if err != nil {
		return nil, err
	}
	return &tx, nil
}

func EncodeAccount(publicKey *ecdsa.PublicKey, address *common.Address) *types.AccountIdentifier {
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return &types.AccountIdentifier{
		Address: address.String(),
		Metadata: map[string]interface{}{
			"publicKey": string(publicKeyBytes),
		},
	}
}
