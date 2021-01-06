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
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type airgapArgBuilderImpl struct {
}

func NewArgBuilder() ArgBuilder {
	return &airgapArgBuilderImpl{}
}

func (c *airgapArgBuilderImpl) FillTxArgs(options *TxArgs, methodArgs ...interface{}) (*TxArgs, error) {
	if options == nil || options.Method == nil {
		return nil, fmt.Errorf("options and options.method cannot be nil")
	}
	parsedArgs, err := options.Method.SerializeArguments(methodArgs...)
	options.Args = parsedArgs
	return options, err
}

func (c *airgapArgBuilderImpl) TransferGold(from common.Address, to common.Address, value *big.Int) (*TxArgs, error) {
	return &TxArgs{
		From:  from,
		To:    &to,
		Value: value,
	}, nil
}

func (c *airgapArgBuilderImpl) CreateAccount(signer common.Address) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: CreateAccount, From: signer})
}

func (c *airgapArgBuilderImpl) AuthorizeVoteSigner(account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: AuthorizeVoteSigner, From: account}, voteSigner, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeAttestationSigner(account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: AuthorizeAttestationSigner, From: account}, attestationSigner, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeValidatorSigner(account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: AuthorizeValidatorSigner, From: account}, validatorSigner, popSignature)
}

func (c *airgapArgBuilderImpl) LockGold(signer common.Address, value *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: LockGold, From: signer, Value: value})
}
func (c *airgapArgBuilderImpl) UnlockGold(signer common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: UnlockGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) RelockGold(signer common.Address, index *big.Int, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: RelockGold, From: signer}, index, amount)
}
func (c *airgapArgBuilderImpl) WithdrawGold(signer common.Address, index *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: WithdrawGold, From: signer}, index)
}
func (c *airgapArgBuilderImpl) Vote(signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: Vote, From: signer}, group, amount)
}
func (c *airgapArgBuilderImpl) ActivateVotes(signer common.Address, group common.Address) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ActivateVotes, From: signer}, group)
}
func (c *airgapArgBuilderImpl) RevokePendingVotes(signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: RevokePendingVotes, From: signer}, signer, group, amount)
}
func (c *airgapArgBuilderImpl) RevokeActiveVotes(signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: RevokeActiveVotes, From: signer}, signer, group, amount)
}

func (c *airgapArgBuilderImpl) ReleaseGoldCreateAccount(releaseGold common.Address, signer common.Address) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldCreateAccount, To: &releaseGold, From: signer})
}
func (c *airgapArgBuilderImpl) ReleaseGoldWithdraw(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldWithdraw, To: &releaseGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeVoteSigner(releaseGold common.Address, account common.Address, voteSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldAuthorizeVoteSigner, To: &releaseGold, From: account}, voteSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeAttestationSigner(releaseGold common.Address, account common.Address, attestationSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldAuthorizeAttestationSigner, To: &releaseGold, From: account}, attestationSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeValidatorSigner(releaseGold common.Address, account common.Address, validatorSigner common.Address, popSignature []byte) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldAuthorizeValidatorSigner, To: &releaseGold, From: account}, validatorSigner, popSignature)
}
func (c *airgapArgBuilderImpl) ReleaseGoldLockGold(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldLockGold, To: &releaseGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldUnlockGold(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldUnlockGold, To: &releaseGold, From: signer}, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRelockGold(releaseGold common.Address, signer common.Address, index *big.Int, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldRelockGold, To: &releaseGold, From: signer}, index, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldWithdrawGold(releaseGold common.Address, signer common.Address, index *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldWithdrawGold, To: &releaseGold, From: signer}, index)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokePendingVotes(releaseGold common.Address, signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldRevokePendingVotes, To: &releaseGold, From: signer}, signer, group, amount)
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokeActiveVotes(releaseGold common.Address, signer common.Address, group common.Address, amount *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: ReleaseGoldRevokeActiveVotes, To: &releaseGold, From: signer}, signer, group, amount)
}

func (c *airgapArgBuilderImpl) StableTokenTransfer(to common.Address, value *big.Int) (*TxArgs, error) {
	return c.FillTxArgs(&TxArgs{Method: StableTokenTransfer, To: &to, Value: value})
}
