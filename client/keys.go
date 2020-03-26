package client

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func Keygen() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

func Derive(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, *common.Address) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return publicKeyECDSA, &address
}
