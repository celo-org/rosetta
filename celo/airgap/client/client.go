package client

import (
	"math/big"

	"github.com/celo-org/rosetta/celo/airgap"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type airgapClientImpl struct {
}

func NewAirgapClient() airgap.AirGapClient {
	return &airgapClientImpl{}
}

func (c *airgapClientImpl) CreateAccount(signer common.Address) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.CreateAccount,
		Args:   []interface{}{},
	}, nil
}
func (c *airgapClientImpl) AuthorizeVoteSigner(account common.Address, popSignature []byte) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   account,
		Method: airgap.AuthorizeVoteSigner,
		Args:   []interface{}{account, popSignature},
	}, nil
}

func (c *airgapClientImpl) LockGold(signer common.Address, amount *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.LockGold,
		Args:   []interface{}{},
		Value:  amount,
	}, nil
}
func (c *airgapClientImpl) UnlockGold(signer common.Address, value *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.UnlockGold,
		Args:   []interface{}{value},
	}, nil
}
func (c *airgapClientImpl) RelockGold(signer common.Address, index *big.Int, value *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.RelockGold,
		Args:   []interface{}{index, value},
	}, nil
}
func (c *airgapClientImpl) WithdrawGold(signer common.Address, index *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.WithdrawGold,
		Args:   []interface{}{index},
	}, nil
}
func (c *airgapClientImpl) Vote(signer common.Address, group common.Address, value *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.Vote,
		Args:   []interface{}{group, value},
	}, nil
}
func (c *airgapClientImpl) ActivateVotes(signer common.Address, account common.Address, group common.Address) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.ActivateVotes,
		Args:   []interface{}{account, group},
	}, nil
}
func (c *airgapClientImpl) RevokePendingVotes(signer common.Address, account common.Address, group common.Address, value *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.RevokePendingVotes,
		Args:   []interface{}{account, group, value},
	}, nil
}
func (c *airgapClientImpl) RevokeActiveVotes(signer common.Address, account common.Address, group common.Address, value *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:   signer,
		Method: airgap.RevokeActiveVotes,
		Args:   []interface{}{account, group, value},
	}, nil
}
func (c *airgapClientImpl) TransferGold(from common.Address, to common.Address, value *big.Int) (*airgap.TxArgs, error) {
	return &airgap.TxArgs{
		From:  from,
		To:    &to,
		Value: value,
	}, nil
}

func (c *airgapClientImpl) TxFromMetadata(tm *airgap.TxMetadata) (*types.Transaction, error) {
	return types.NewTransaction(
		tm.Nonce,
		tm.To,
		tm.Value,
		tm.Gas,
		tm.GasPrice,
		tm.FeeCurrency,
		tm.GatewayFeeRecipient,
		tm.GatewayFee,
		tm.Data,
	), nil
}
