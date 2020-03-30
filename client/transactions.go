package client

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/api"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func ConstructTransferTransaction(from *common.Address, to *common.Address, amount *big.Int, metadata *api.TransferMetadata) (*types.Transaction, error) {
	if metadata.Balance.Cmp(amount) == -1 {
		return nil, fmt.Errorf("Provided amount %d exceeds balance %d of account %s", amount.Uint64(), metadata.Balance.Uint64(), from)
	}

	transaction := types.NewTransaction(
		metadata.Txdata.AccountNonce,
		*to,
		amount,
		metadata.Txdata.GasLimit,
		metadata.Txdata.Price,
		nil, // guarantee cGLD
		metadata.Txdata.GatewayFeeRecipient,
		metadata.Txdata.GatewayFee,
		[]byte{},
	)

	return transaction, nil
}

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
