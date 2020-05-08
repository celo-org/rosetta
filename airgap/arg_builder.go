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
	return CreateAccount.CreateTxArgs(signer, nil)
}

func (c *airgapArgBuilderImpl) AuthorizeVoteSigner(account common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeVoteSigner.CreateTxArgs(account, nil, account, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeAttestationSigner(account common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeAttestationSigner.CreateTxArgs(account, nil, account, popSignature)
}

func (c *airgapArgBuilderImpl) AuthorizeValidatorSigner(account common.Address, popSignature []byte) (*TxArgs, error) {
	return AuthorizeValidatorSigner.CreateTxArgs(account, nil, account, popSignature)
}

func (c *airgapArgBuilderImpl) LockGold(signer common.Address, amount *big.Int) (*TxArgs, error) {
	return LockGold.CreateTxArgs(signer, amount)
}
func (c *airgapArgBuilderImpl) UnlockGold(signer common.Address, value *big.Int) (*TxArgs, error) {
	return UnlockGold.CreateTxArgs(signer, nil, value)
}
func (c *airgapArgBuilderImpl) RelockGold(signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error) {
	return RelockGold.CreateTxArgs(signer, nil, index, value)
}
func (c *airgapArgBuilderImpl) WithdrawGold(signer common.Address, index *big.Int) (*TxArgs, error) {
	return WithdrawGold.CreateTxArgs(signer, nil, index)
}
func (c *airgapArgBuilderImpl) Vote(signer common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return Vote.CreateTxArgs(signer, nil, group, value)
}
func (c *airgapArgBuilderImpl) ActivateVotes(signer common.Address, account common.Address, group common.Address) (*TxArgs, error) {
	return ActivateVotes.CreateTxArgs(signer, nil, account, group)
}
func (c *airgapArgBuilderImpl) RevokePendingVotes(signer common.Address, account common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return RevokePendingVotes.CreateTxArgs(signer, nil, account, group, value)
}

func (c *airgapArgBuilderImpl) RevokeActiveVotes(signer common.Address, account common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	return RevokeActiveVotes.CreateTxArgs(signer, nil, account, group, value)
}

func (c *airgapArgBuilderImpl) ReleaseGoldCreateAccount(releaseGold common.Address, signer common.Address) (*TxArgs, error) {
	txArgs, err := ReleaseGoldCreateAccount.CreateTxArgs(signer, nil)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}

func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeVoteSigner(releaseGold common.Address, account common.Address, popSignature []byte) (*TxArgs, error) {
	txArgs, err := ReleaseGoldAuthorizeVoteSigner.CreateTxArgs(account, nil, account, popSignature)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeAttestationSigner(releaseGold common.Address, account common.Address, popSignature []byte) (*TxArgs, error) {
	txArgs, err := ReleaseGoldAuthorizeAttestationSigner.CreateTxArgs(account, nil, account, popSignature)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldAuthorizeValidatorSigner(releaseGold common.Address, account common.Address, popSignature []byte) (*TxArgs, error) {
	txArgs, err := ReleaseGoldAuthorizeValidatorSigner.CreateTxArgs(account, nil, account, popSignature)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldLockGold(releaseGold common.Address, signer common.Address, amount *big.Int) (*TxArgs, error) {
	txArgs, err := ReleaseGoldLockGold.CreateTxArgs(signer, amount)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldUnlockGold(releaseGold common.Address, signer common.Address, value *big.Int) (*TxArgs, error) {
	txArgs, err := ReleaseGoldUnlockGold.CreateTxArgs(signer, nil, value)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldRelockGold(releaseGold common.Address, signer common.Address, index *big.Int, value *big.Int) (*TxArgs, error) {
	txArgs, err := ReleaseGoldRelockGold.CreateTxArgs(signer, nil, index, value)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldWithdrawGold(releaseGold common.Address, signer common.Address, index *big.Int) (*TxArgs, error) {
	txArgs, err := ReleaseGoldWithdrawGold.CreateTxArgs(signer, nil, index)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokePendingVotes(releaseGold common.Address, signer common.Address, account common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	txArgs, err := ReleaseGoldRevokePendingVotes.CreateTxArgs(signer, nil, account, group, value)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
func (c *airgapArgBuilderImpl) ReleaseGoldRevokeActiveVotes(releaseGold common.Address, signer common.Address, account common.Address, group common.Address, value *big.Int) (*TxArgs, error) {
	txArgs, err := ReleaseGoldRevokeActiveVotes.CreateTxArgs(signer, nil, account, group, value)
	if err != nil {
		return nil, err
	}
	txArgs.To = &releaseGold
	return txArgs, nil
}
