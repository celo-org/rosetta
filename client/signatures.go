package client

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func Sign(message []byte, privateKey *ecdsa.PrivateKey) *[]byte {
	digest := crypto.Keccak256(message)
	sig, err := crypto.Sign(digest, privateKey)
	if err != nil {
		log.Fatal("error generating signature")
	}
	return &sig
}

func Verify(message []byte, publicKey *ecdsa.PublicKey, signature *[]byte) bool {
	digest := crypto.Keccak256(message)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return crypto.VerifySignature(publicKeyBytes, digest, *signature)
}
