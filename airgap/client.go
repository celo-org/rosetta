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

	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/registry"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
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
	gethTx, err := tx.AsGethTransaction()
	if err != nil {
		return nil, err
	}

	h := signer.Hash(gethTx)
	sig, err := crypto.Sign(h[:], privateKey)
	if err != nil {
		return nil, err
	}

	tx.Signature = sig
	return tx, nil
}

func (c *clientImpl) GenerateProofOfPossessionSignature(privateKey *ecdsa.PrivateKey, address *common.Address) ([]byte, error) {
	msg := accounts.TextHash(address.Bytes())
	signature, err := c.Sign(msg, privateKey)
	return signature, err
}

func (c *clientImpl) ParseTxArgs(method *CeloMethod, metadata *TxMetadata) (*TxArgs, error) {
	var abi *abi.ABI
	var err error

	switch method.Contract {
	case ReleaseGold:
		abi, err = contracts.ParseReleaseGoldABI()
	case registry.AccountsContractID.String():
		abi, err = contracts.ParseAccountsABI()
	case registry.LockedGoldContractID.String():
		abi, err = contracts.ParseLockedGoldABI()
	case registry.ElectionContractID.String():
		abi, err = contracts.ParseElectionABI()
	default:
		err = fmt.Errorf("Unknown contract %s", method.Contract)
	}
	if err != nil {
		return nil, err
	}

	abiMethod, ok := abi.Methods[method.Name]
	if !ok {
		return nil, fmt.Errorf("Method %s not found on ABI for contract %s", method.Name, method.Contract)
	}

	args, err := abiMethod.Inputs.UnpackValues(metadata.Data)
	if err != nil {
		return nil, err
	}

	return &TxArgs{
		From:   metadata.From,
		To:     &metadata.To,
		Value:  metadata.Value,
		Method: method,
		Args:   args,
	}, nil
}
