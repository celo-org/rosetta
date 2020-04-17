package client

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
)

func Sign(message []byte, privateKey *ecdsa.PrivateKey) (*[]byte, error) {
	digest := crypto.Keccak256(message)
	sig, err := crypto.Sign(digest, privateKey)
	return &sig, err
}

func Verify(message []byte, publicKey *ecdsa.PublicKey, signature *[]byte) bool {
	digest := crypto.Keccak256(message)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return crypto.VerifySignature(publicKeyBytes, digest, *signature)
}
