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
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type airgapArgBuilderImpl struct {
}

func NewArgBuilder() ArgBuilder {
	return &airgapArgBuilderImpl{}
}

func (c *airgapArgBuilderImpl) TransferGold(from common.Address, to common.Address, value *big.Int) (*TxArgs, error) {
	return &TxArgs{
		From:  from,
		To:    &to,
		Value: value,
	}, nil
}

func (c *airgapArgBuilderImpl) CreateAccount(signer common.Address) (*TxArgs, error) {
	return CreateAccount.CreateTxArgs(nil, signer, nil)
}

func (c *airgapArgBuilderImpl) AuthorizeVoteSigner(account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeVoteSigner.CreateTxArgs(nil, account, nil, voteSigner, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeAttestationSigner(account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeAttestationSigner.CreateTxArgs(nil, account, nil, attestationSigner, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeValidatorSigner(account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeValidatorSigner.CreateTxArgs(nil, account, nil, validatorSigner, popSignature)
}

func (c *airgapArgBuilderImpl) LockGold(signer common.Address, amount *big.Int) (*TxArgs, error) {
	return LockGold.CreateTxArgs(nil, signer, amount)
}
func (c *airgapArgBuilderImpl) UnlockGold(signer common.Address, value *big.Int) (*TxArgs, error) {
	return UnlockGold.CreateTxArgs(nil, signer, nil, value)
}
func (c *airgapArgBuilderImpl) RelockGold(signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error) {
	return RelockGold.CreateTxArgs(nil, signer, nil, index, value)
}
func (c *airgapArgBuilderImpl) WithdrawGold(signer common.Address, index *big.Int) (*TxArgs, error) {
	return WithdrawGold.CreateTxArgs(nil, signer, nil, index)
}
func (c *airgapArgBuilderImpl) Vote(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return Vote.CreateTxArgs(nil, signer, nil, group, value)
}
func (c *airgapArgBuilderImpl) ActivateVotes(signer common.Address, group common.Address) (*TxArgs, error) {
	return ActivateVotes.CreateTxArgs(nil, signer, nil, group)
}
func (c *airgapArgBuilderImpl) RevokePendingVotes(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return RevokePendingVotes.CreateTxArgs(nil, signer, nil, signer, group, value)
}
func (c *airgapArgBuilderImpl) RevokeActiveVotes(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return RevokeActiveVotes.CreateTxArgs(nil, signer, nil, signer, group, value)
}

func (c *airgapArgBuilderImpl) ReleaseGoldCreateAccount(releaseGold common.Address, signer common.Address) (*TxArgs, error) {
	return ReleaseGoldCreateAccount.CreateTxArgs(&releaseGold, signer, nil)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeVoteSigner(releaseGold common.Address, account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return ReleaseGoldAuthorizeVoteSigner.CreateTxArgs(&releaseGold, account, nil, voteSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeAttestationSigner(releaseGold common.Address, account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return ReleaseGoldAuthorizeAttestationSigner.CreateTxArgs(&releaseGold, account, nil, attestationSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeValidatorSigner(releaseGold common.Address, account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return ReleaseGoldAuthorizeValidatorSigner.CreateTxArgs(&releaseGold, account, nil, validatorSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldLockGold(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldLockGold.CreateTxArgs(&releaseGold, signer, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldUnlockGold(releaseGold common.Address, signer common.Address, value *big.Int) (*TxArgs, error) {
	return ReleaseGoldUnlockGold.CreateTxArgs(&releaseGold, signer, nil, value)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRelockGold(releaseGold common.Address, signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error) {
	return ReleaseGoldRelockGold.CreateTxArgs(&releaseGold, signer, nil, index, value)
}
func (c *airgapArgBuilderImpl) ReleaseGoldWithdrawGold(releaseGold common.Address, signer common.Address, index *big.Int) (*TxArgs, error) {
	return ReleaseGoldWithdrawGold.CreateTxArgs(&releaseGold, signer, nil, index)

}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokePendingVotes(releaseGold common.Address, signer common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return ReleaseGoldRevokePendingVotes.CreateTxArgs(&releaseGold, signer, nil, signer, group, value)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokeActiveVotes(releaseGold common.Address, signer common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return ReleaseGoldRevokeActiveVotes.CreateTxArgs(&releaseGold, signer, nil, signer, group, value)
}
