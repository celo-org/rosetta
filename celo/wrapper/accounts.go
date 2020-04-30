package wrapper

import (
	"fmt"

	"github.com/celo-org/rosetta/celo/contract"
	"github.com/ethereum/go-ethereum/crypto"
)

type AccountsWrapper struct {
	*contract.Accounts
}

func NewAccounts(contract *contract.Accounts) *AccountsWrapper {
	return &AccountsWrapper{contract}
}

type EncodedSignature struct {
	R [32]byte
	S [32]byte
	V uint8
}

func (w *AccountsWrapper) AuthorizeMetadata(popSignature []byte) (*EncodedSignature, error) {
	if len(popSignature) != crypto.SignatureLength {
		return nil, fmt.Errorf("Invalid signature")
	}

	var r [32]byte
	copy(r[:], popSignature[:32])
	var s [32]byte
	copy(s[:], popSignature[32:64])
	v := popSignature[64]

	encodedSig := EncodedSignature{
		R: r,
		S: s,
		V: v,
	}
	return &encodedSig, nil
}
