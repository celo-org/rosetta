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
	return CreateAccount.CreateTxArgs(&TxArgs{From: signer})
}

func (c *airgapArgBuilderImpl) AuthorizeVoteSigner(account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeVoteSigner.CreateTxArgs(&TxArgs{From: account}, voteSigner, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeAttestationSigner(account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeAttestationSigner.CreateTxArgs(&TxArgs{From: account}, attestationSigner, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeValidatorSigner(account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeValidatorSigner.CreateTxArgs(&TxArgs{From: account}, validatorSigner, popSignature)
}

func (c *airgapArgBuilderImpl) LockGold(signer common.Address, value *big.Int) (*TxArgs, error) {
	return LockGold.CreateTxArgs(&TxArgs{From: signer, Value: value})
}
func (c *airgapArgBuilderImpl) UnlockGold(signer common.Address, amount *big.Int) (*TxArgs, error) {
	return UnlockGold.CreateTxArgs(&TxArgs{From: signer}, amount)
}
func (c *airgapArgBuilderImpl) RelockGold(signer common.Address, index *big.Int, amount *big.Int) (*TxArgs, error) {
	return RelockGold.CreateTxArgs(&TxArgs{From: signer}, index, amount)
}
func (c *airgapArgBuilderImpl) WithdrawGold(signer common.Address, index *big.Int) (*TxArgs, error) {
	return WithdrawGold.CreateTxArgs(&TxArgs{From: signer}, index)
}
func (c *airgapArgBuilderImpl) Vote(signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return Vote.CreateTxArgs(&TxArgs{From: signer}, group, amount)
}
func (c *airgapArgBuilderImpl) ActivateVotes(signer common.Address, group common.Address) (*TxArgs, error) {
	return ActivateVotes.CreateTxArgs(&TxArgs{From: signer}, group)
}
func (c *airgapArgBuilderImpl) RevokePendingVotes(signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return RevokePendingVotes.CreateTxArgs(&TxArgs{From: signer}, signer, group, amount)
}
func (c *airgapArgBuilderImpl) RevokeActiveVotes(signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return RevokeActiveVotes.CreateTxArgs(&TxArgs{From: signer}, signer, group, amount)
}

func (c *airgapArgBuilderImpl) ReleaseGoldCreateAccount(releaseGold common.Address, signer common.Address) (*TxArgs, error) {
	return ReleaseGoldCreateAccount.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer})
}
func (c *airgapArgBuilderImpl) ReleaseGoldWithdraw(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldWithdraw.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeVoteSigner(releaseGold common.Address, account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return ReleaseGoldAuthorizeVoteSigner.CreateTxArgs(&TxArgs{To: &releaseGold, From: account}, voteSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeAttestationSigner(releaseGold common.Address, account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return ReleaseGoldAuthorizeAttestationSigner.CreateTxArgs(&TxArgs{To: &releaseGold, From: account}, attestationSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeValidatorSigner(releaseGold common.Address, account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return ReleaseGoldAuthorizeValidatorSigner.CreateTxArgs(&TxArgs{To: &releaseGold, From: account}, validatorSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldLockGold(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldLockGold.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldUnlockGold(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldUnlockGold.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRelockGold(releaseGold common.Address, signer common.Address, index *big.Int, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldRelockGold.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, index, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldWithdrawGold(releaseGold common.Address, signer common.Address, index *big.Int) (*TxArgs, error) {
	return ReleaseGoldWithdrawGold.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, index)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokePendingVotes(releaseGold common.Address, signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldRevokePendingVotes.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, signer, group, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokeActiveVotes(releaseGold common.Address, signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return ReleaseGoldRevokeActiveVotes.CreateTxArgs(&TxArgs{To: &releaseGold, From: signer}, signer, group, amount)
}
