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

	"github.com/celo-org/celo-blockchain/accounts"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/crypto"
	"github.com/celo-org/kliento/contracts"
	"github.com/celo-org/kliento/registry"
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

// Sign signs an arbitrary message with the private key the returned signature
// as 64 bytes long and in the [R || S] format.
func (c *clientImpl) Sign(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error) {
	// crypto.Sigh produces signatures in the [ R || S || V ] format, but the
	// signatures returned by this method are not used for recovery and so the
	// recovery byte is removed.
	sig, err := crypto.Sign(crypto.Keccak256(message), privateKey)
	if err != nil {
		return nil, err
	}
	return sig[:64], nil
}

// Verify verifies the signature of an arbitrary message, it expects a 64 byte signature in the [ R || S ] format.
func (c *clientImpl) Verify(message []byte, publicKey *ecdsa.PublicKey, signature []byte) bool {
	return crypto.VerifySignature(crypto.FromECDSAPub(publicKey), crypto.Keccak256(message), signature)
}

// ConstructTxFromMetadata creates a new transaction using given Metadata
func (c *clientImpl) ConstructTxFromMetadata(tm *TxMetadata) (*Transaction, error) {
	return &Transaction{
		TxMetadata: tm,
	}, nil
}

// SignTx signs an unsignedTx using the private key and returns a signedTx that can be submitted to the node.
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

// GenerateProofOfPossessionSignature generates a recoverable (65 byte) ecdsa
// signature over the given address using the given privateKey. The signature
// is used to authorize release gold operrations on chain. The signature
// returned is in the [ R || S || V ] format where V is 27 or 28. This is a
// stange hangover from Bitcoin and is required becuase the precompiled
// ecrecover contract in the evm expects this format.
//
// See explanation for this format here:
// https://github.com/ethereum/go-ethereum/issues/19751#issuecomment-504900739
func (c *clientImpl) GenerateProofOfPossessionSignature(privateKey *ecdsa.PrivateKey, address *common.Address) ([]byte, error) {

	sig, err := crypto.Sign(accounts.TextHash(address.Bytes()), privateKey)
	// if there was no error then adjust the recovery byte.
	if err == nil {
		sig[crypto.RecoveryIDOffset] += 27
	}
	return sig, err
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
