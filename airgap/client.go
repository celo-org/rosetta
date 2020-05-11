// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package airgap

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type clientImpl struct {
}

func NewClient() Client {
	return &clientImpl{}
}

// Keygen generates a private key
func (c *clientImpl) Keygen() (*ecdsa.PrivateKey, error) {
	return crypto.GenerateKey()
}

// Derive the cryptographic public key and on-chain address from the private key
func (c *clientImpl) Derive(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, *common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("error casting public key to ECDSA")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return publicKeyECDSA, &address, nil
}

// Sign an arbitrary message with the private key
func (c *clientImpl) Sign(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	digest := crypto.Keccak256(message)
	sig, err := crypto.Sign(digest, privateKey)
	return sig, err
}

// Verify the signature of an arbitrary message
func (c *clientImpl) Verify(message []byte, publicKey *ecdsa.PublicKey, signature []byte) bool {
	digest := crypto.Keccak256(message)
	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return crypto.VerifySignature(publicKeyBytes, digest, signature)
}

// ConstructTxFromMetadata creates a new transaction using given Metadata
func (c *clientImpl) ConstructTxFromMetadata(tm *TxMetadata) (*Transaction, error) {
	return &Transaction{
		TxMetadata: tm,
	}, nil
}

// SignTx signs an unsignedTx using the private seed and return a signedTx that can be submitted to the node
func (c *clientImpl) SignTx(tx *Transaction, privateKey *ecdsa.PrivateKey) (*Transaction, error) {
	signer := types.NewEIP155Signer(tx.ChainId)
	gethTx := types.NewTransaction(
		tx.Nonce,
		tx.To,
		tx.Value,
		tx.Gas,
		tx.GasPrice,
		tx.FeeCurrency,
		tx.GatewayFeeRecipient,
		tx.GatewayFee,
		tx.Data,
	)
	signedTx, err := types.SignTx(gethTx, signer, privateKey)
	if err != nil {
		return nil, err
	}

	tx.V, tx.R, tx.S = signedTx.RawSignatureValues()

	// TODO copy tx?
	return tx, nil
}

func (c *clientImpl) GenerateProofOfPossessionSignature(privateKey *ecdsa.PrivateKey, address *common.Address) ([]byte, error) {
	msg := accounts.TextHash(address.Bytes())
	signature, err := c.Sign(msg, privateKey)
	return signature, err
}
