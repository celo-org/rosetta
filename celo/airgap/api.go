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
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type AirGapServer interface {
	ObtainMetadata(ctx context.Context, txOpts *TxArgs) (*TxMetadata, error)
	SubmitTx(ctx context.Context, rawTx []byte) (*common.Hash, error)
}

type AirGapClient interface {
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

	TxFromMetadata(*TxMetadata) (*types.Transaction, error)
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
