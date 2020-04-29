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

package transaction

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CeloMethod string

// values taken from contract method names for ABI usage
var (
	CreateAccount       CeloMethod = "createAccount"
	AuthorizeVoteSigner CeloMethod = "authorizeVoteSigner"
	LockGold            CeloMethod = "lock"
	UnlockGold          CeloMethod = "unlock"
	RelockGold          CeloMethod = "relock"
	WithdrawGold        CeloMethod = "withdraw"
	Vote                CeloMethod = "vote"
	ActivateVotes       CeloMethod = "activate"
	RevokePendingVotes  CeloMethod = "revokePending"
	RevokeActiveVotes   CeloMethod = "revokeActive"
)

func (tt CeloMethod) String() string { return string(tt) }

var (
	CeloMethodToRegistryKey = map[*CeloMethod]*wrapper.RegistryKey{
		&CreateAccount:       &wrapper.AccountsRegistryId,
		&AuthorizeVoteSigner: &wrapper.AccountsRegistryId,
		&LockGold:            &wrapper.LockedGoldRegistryId,
		&UnlockGold:          &wrapper.LockedGoldRegistryId,
		&RelockGold:          &wrapper.LockedGoldRegistryId,
		&WithdrawGold:        &wrapper.LockedGoldRegistryId,
		&Vote:                &wrapper.ElectionRegistryId,
		&ActivateVotes:       &wrapper.ElectionRegistryId,
		&RevokePendingVotes:  &wrapper.ElectionRegistryId,
		&RevokeActiveVotes:   &wrapper.ElectionRegistryId,
	}
)

type TransactionOptions struct {
	From   *common.Address
	To     *common.Address // non-nil means exclusively cGLD transfer
	Value  *big.Int
	Method *CeloMethod // non-nil means celo registry contract invokation
	Args   []interface{}
}

// [note]: non cGLD fee currencies currently unsupported
type AccountMetadata struct {
	From  common.Address
	Nonce uint64
}

type NodeMetadata struct {
	GasPrice            *big.Int
	GatewayFeeRecipient *common.Address
	GatewayFee          *big.Int
}

type MethodMetadata struct {
	To   *common.Address
	Data []byte
}

type TransactionMetadata struct {
	*AccountMetadata
	*NodeMetadata
	*MethodMetadata
	Value *big.Int
	Gas   uint64
}

func (tm *TransactionMetadata) AsTransaction() *types.Transaction {
	return types.NewTransaction(
		tm.Nonce,
		*tm.To,
		tm.Value,
		tm.Gas,
		tm.GasPrice,
		nil, // non-cGLD fees not supported
		tm.GatewayFeeRecipient,
		tm.GatewayFee,
		tm.Data,
	)
}
