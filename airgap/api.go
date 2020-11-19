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
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

type Server interface {
	ObtainMetadata(ctx context.Context, txOpts *TxArgs) (*TxMetadata, error)
	SubmitTx(ctx context.Context, rawTx []byte) (*common.Hash, error)
	CallData(ctx context.Context, callOpts *CallParams) ([]byte, error)
	FilterQuery(ctx context.Context, filterQueryOps *FilterQueryParams) ([]types.Log, error)
}

type ArgBuilder interface {
	TransferGold(from common.Address, to common.Address, value *big.Int) (*TxArgs, error)

	CreateAccount(signer common.Address) (*TxArgs, error)
	AuthorizeVoteSigner(account common.Address, signer common.Address, popSignature []byte) (*TxArgs, error)
	LockGold(signer common.Address, value *big.Int) (*TxArgs, error)
	UnlockGold(signer common.Address, value *big.Int) (*TxArgs, error)
	RelockGold(signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error)
	WithdrawGold(signer common.Address, index *big.Int) (*TxArgs, error)
	Vote(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error)
	ActivateVotes(signer common.Address, group common.Address) (*TxArgs, error)
	RevokePendingVotes(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error)
	RevokeActiveVotes(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error)

	ReleaseGoldCreateAccount(releaseGold common.Address, signer common.Address) (*TxArgs, error)
	ReleaseGoldWithdraw(releaseGold common.Address, signer common.Address, value *big.Int) (*TxArgs, error)
	ReleaseGoldAuthorizeVoteSigner(releaseGold common.Address, account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error)
	ReleaseGoldAuthorizeAttestationSigner(releaseGold common.Address, account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error)
	ReleaseGoldAuthorizeValidatorSigner(releaseGold common.Address, account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error)
	ReleaseGoldLockGold(releaseGold common.Address, signer common.Address, value *big.Int) (*TxArgs, error)
	ReleaseGoldUnlockGold(releaseGold common.Address, signer common.Address, value *big.Int) (*TxArgs, error)
	ReleaseGoldRelockGold(releaseGold common.Address, signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error)
	ReleaseGoldWithdrawGold(releaseGold common.Address, signer common.Address, index *big.Int) (*TxArgs, error)
	ReleaseGoldRevokePendingVotes(releaseGold common.Address, signer common.Address, group common.Address, value *big.Int) (*TxArgs, error)
	ReleaseGoldRevokeActiveVotes(releaseGold common.Address, signer common.Address, group common.Address, value *big.Int) (*TxArgs, error)
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

	// Parse transaction arguments from transaction metadata
	ParseTxArgs(txMetadata *TxMetadata) (*TxArgs, error)

	// Parse transaction data to celo method and args
	ParseMethodAndArgs(data []byte) (*CeloMethod, []interface{}, error)

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

type CallParams struct {
	TxArgs
	BlockNumber *big.Int
}

type FilterQueryParams struct {
	Event     *CeloEvent
	Topics    [][]interface{}
	FromBlock *big.Int
	ToBlock   *big.Int
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
	Signature   []byte `json:"signature"`
}

func (tx *Transaction) Signed() bool {
	return len(tx.Signature) > 0
}

func (tx *Transaction) AsGethTransaction() (*types.Transaction, error) {
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
	if tx.Signed() {
		signer := types.NewEIP155Signer(tx.ChainId)
		signedGethTx, err := gethTx.WithSignature(signer, tx.Signature)
		if err != nil {
			return nil, err
		}
		return signedGethTx, nil
	}
	return gethTx, nil
}

type SignatureValues struct {
	V *big.Int `json:"v"`
	R *big.Int `json:"r"`
	S *big.Int `json:"s"`
}

func (tx *Transaction) GetSignatureValues() (*SignatureValues, error) {
	gethTx, err := tx.AsGethTransaction()
	if err != nil {
		return nil, err
	}
	v, r, s := gethTx.RawSignatureValues()
	return &SignatureValues{v, r, s}, nil
}

func (tx *Transaction) Hash() (common.Hash, error) {
	gethTx, err := tx.AsGethTransaction()
	if err != nil {
		return common.Hash{}, err
	}
	return gethTx.Hash(), nil
}

func (tx *Transaction) Serialize() ([]byte, error) {
	gethTx, err := tx.AsGethTransaction()
	if err != nil {
		return nil, err
	}
	return rlp.EncodeToBytes(gethTx)
}

func (tx *Transaction) Deserialize(data []byte, chainId *big.Int) error {
	var err error

	gethTx := new(types.Transaction)
	if err = rlp.DecodeBytes(data, gethTx); err != nil {
		return err
	}
	tx.TxMetadata = &TxMetadata{}
	tx.Nonce = gethTx.Nonce()
	tx.GasPrice = gethTx.GasPrice()
	tx.GatewayFee = gethTx.GatewayFee()
	tx.GatewayFeeRecipient = gethTx.GatewayFeeRecipient()
	tx.FeeCurrency = gethTx.FeeCurrency()
	tx.To = *gethTx.To()
	tx.Data = gethTx.Data()
	tx.Value = gethTx.Value()
	tx.Gas = gethTx.Gas()
	tx.ChainId = chainId

	v, r, s := gethTx.RawSignatureValues()
	if v == nil || r == nil || s == nil {
		return errors.New("can't deserialize unsigned transactions")
	}
	tx.Signature = ValuesToSignature(chainId, v, r, s)
	tx.From, err = types.Sender(types.NewEIP155Signer(chainId), gethTx)
	if err != nil {
		return err
	}

	return nil
}

func ValuesToSignature(chainId, v, r, s *big.Int) []byte {
	sig := make([]byte, crypto.SignatureLength)

	// doing the inverse of signer.SignatureValues()
	copy(sig, r.Bytes())
	copy(sig[32:64], s.Bytes())

	// v = byte[64] + 35 + chainId * 2
	// byte[64] = v - 35 - chainId * 2
	chainMul := new(big.Int).Mul(chainId, big.NewInt(2))
	sig[64] = byte(new(big.Int).Sub(v, chainMul).Int64() - 35)

	return sig
}
