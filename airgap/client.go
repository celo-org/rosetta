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
	return c.SignHash(crypto.Keccak256(message), privateKey)
}

func (c *clientImpl) SignHash(digest []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(digest, privateKey)
	if err != nil {
		// see https://github.com/celo-org/celo-blockchain/blob/0792a7189b531e22b97b81b6d6aa29301c3ebb8e/internal/ethapi/api.go#L1613
		sig[crypto.RecoveryIDOffset] += 27
	}
	return sig, err
}

func (c *clientImpl) SignWithPrefix(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	return c.SignHash(accounts.TextHash(message), privateKey)
}

// Verify the signature of an arbitrary message
func (c *clientImpl) Verify(message []byte, publicKey *ecdsa.PublicKey, signature []byte) bool {
	return c.VerifyHash(crypto.Keccak256(message), publicKey, signature)
}

func (c *clientImpl) VerifyHash(digest []byte, publicKey *ecdsa.PublicKey, signature []byte) bool {
	if signature[crypto.RecoveryIDOffset] != 27 && signature[crypto.RecoveryIDOffset] != 28 {
		return false
	}
	signature[crypto.RecoveryIDOffset] -= 27

	publicKeyBytes := crypto.FromECDSAPub(publicKey)
	return crypto.VerifySignature(publicKeyBytes, digest, signature)
}

func (c *clientImpl) VerifyWithPrefix(message []byte, publicKey *ecdsa.PublicKey, signature []byte) bool {
	return c.VerifyHash(accounts.TextHash(message), publicKey, signature)
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
	return c.SignWithPrefix(address.Bytes(), privateKey)
}

var abiParsers = map[string]func() (*abi.ABI, error){
	ReleaseGold:                            contracts.ParseReleaseGoldABI,
	registry.AccountsContractID.String():   contracts.ParseAccountsABI,
	registry.LockedGoldContractID.String(): contracts.ParseLockedGoldABI,
	registry.ElectionContractID.String():   contracts.ParseElectionABI,
}

func (c *clientImpl) ParseMethodAndArgs(data []byte) (*CeloMethod, []interface{}, error) {
	if len(data) == 0 {
		return nil, nil, nil
	}

	methodId, methodData := data[:4], data[4:]

	for contractName, abiParser := range abiParsers {
		abi, err := abiParser()
		if err != nil {
			continue
		}

		abiMethod, err := abi.MethodById(methodId)
		if err != nil {
			continue
		}

		args, err := abiMethod.Inputs.UnpackValues(methodData)
		if err != nil {
			return nil, nil, err
		}

		method, err := MethodFromString(fmt.Sprintf("%s.%s", contractName, abiMethod.Name))
		return method, args, err
	}

	return nil, nil, fmt.Errorf("data does not match any abi parsers")
}

func (c *clientImpl) ParseTxArgs(metadata *TxMetadata) (*TxArgs, error) {
	method, args, err := c.ParseMethodAndArgs(metadata.Data)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse method and args from %s with %s", metadata.Data, err.Error())
	}

	return &TxArgs{
		From:   metadata.From,
		To:     &metadata.To,
		Value:  metadata.Value,
		Method: method,
		Args:   args,
	}, nil
}
