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
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
)

type Server interface {
	ObtainMetadata(ctx context.Context, txOpts *TxArgs) (*TxMetadata, error)
	SubmitTx(ctx context.Context, rawTx []byte) (*common.Hash, error)
}

type ArgBuilder interface {
	CreateAccount(signer common.Address) (*TxArgs, error)
	AuthorizeVoteSigner(account common.Address, popSignature []byte) (*TxArgs, error)
	LockGold(signer common.Address, amount *big.Int) (*TxArgs, error)
	UnlockGold(signer common.Address, value *big.Int) (*TxArgs, error)
	RelockGold(signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error)
	WithdrawGold(signer common.Address, index *big.Int) (*TxArgs, error)
	Vote(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error)
	ActivateVotes(signer common.Address, account common.Address, group common.Address) (*TxArgs, error)
	RevokePendingVotes(signer common.Address, account common.Address, group common.Address, value *big.Int) (*TxArgs, error)
	RevokeActiveVotes(signer common.Address, account common.Address, group common.Address, value *big.Int) (*TxArgs, error)
	TransferGold(from common.Address, to common.Address, value *big.Int) (*TxArgs, error)
}

type Client interface {
	// Keygen generates a private key
	Keygen() (*ecdsa.PrivateKey, error)
	// Derive the cryptographic public key and on-chain address from the private key
	Derive(privateKey *ecdsa.PrivateKey) (*ecdsa.PublicKey, *common.Address, error)
	// Sign an arbitrary message with the private key
	Sign(message []byte, privateKey *ecdsa.PrivateKey) ([]byte, error)
	// Verify the signature of an arbitrary message
	Verify(message []byte, publicKey *ecdsa.PublicKey, signature []byte) bool

	// ConstructTxFromMetadata creates a new transaction using given Metadata
	ConstructTxFromMetadata(txMetadata *TxMetadata) (*Transaction, error)
	// SignTx signs an unsignedTx using the private seed and return a signedTx that can be submitted to the node
	SignTx(transaction *Transaction, privateKey *ecdsa.PrivateKey) (*Transaction, error)

	// GenerateProofOfPossessionSignature generates a PoP needed for Authorize calls
	GenerateProofOfPossessionSignature(privateKey *ecdsa.PrivateKey, address *common.Address) ([]byte, error)
}

type TxArgs struct {
	From  common.Address
	Value *big.Int
	// non-nil means exclusively cGLD transfer
	To *common.Address
	// non-nil means celo registry contract invokation
	Method *CeloMethod
	Args   []interface{}
}

type TxMetadata struct {
	From                common.Address
	Nonce               uint64
	GasPrice            *big.Int
	GatewayFeeRecipient *common.Address
	GatewayFee          *big.Int
	FeeCurrency         *common.Address
	To                  common.Address
	Data                []byte
	Value               *big.Int
	Gas                 uint64
	ChainId             *big.Int
}

func (tm *TxMetadata) AsCallMessage() ethereum.CallMsg {
	return ethereum.CallMsg{
		From:                tm.From,
		GatewayFee:          tm.GatewayFee,
		GatewayFeeRecipient: tm.GatewayFeeRecipient,
		GasPrice:            tm.GasPrice,
		To:                  &tm.To,
		Data:                tm.Data,
		Value:               tm.Value,
		FeeCurrency:         tm.FeeCurrency,
	}
}

type Transaction struct {
	*TxMetadata `json:"metadata"`

	// Signature values
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (t *Transaction) Signed() bool {
	return t.V != nil && t.R != nil && t.S != nil
}

func (tx *Transaction) AsGethTransaction() *types.Transaction {
	return types.NewTransaction(
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
}

func (tx *Transaction) Hash() common.Hash {
	return tx.AsGethTransaction().Hash()
}

func (tx *Transaction) Serialize() ([]byte, error) {
	gethTx := tx.AsGethTransaction()
	return rlp.EncodeToBytes(gethTx)
}
