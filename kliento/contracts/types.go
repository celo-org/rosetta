package contracts

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
)

// Contract includes basic functions to call and transact with the contract
type Contract interface {
	Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error
	Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error)
}

// InitializableContract includes functions from the Initializable Solidity interface
type InitializableContract interface {
	Initialized(opts *bind.CallOpts) (bool, error)
	Initialize(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error)
}

// CeloVersionedContract includes functions from the ICeloVersionedContract Solidity interface
type CeloVersionedContract interface {
	GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error)
}

// OwnableContract includes functions from the Ownable Solidity interface
type OwnableContract interface {
	IsOwner(opts *bind.CallOpts) (bool, error)
	Owner(opts *bind.CallOpts) (common.Address, error)
	RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error)
	TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error)
}

// Erc20Contract includes functions from the IERC20 Solidity interface.
// This only includes the required functions by ERC-20.
type Erc20Contract interface {
	TotalSupply(opts *bind.CallOpts) (*big.Int, error)
	BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error)
	Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error)
	Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error)
	Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error)
	TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error)
}

// CeloTokenContract includes functions from the IERC20 Solidity interface and
// the ICeloToken Solidity interface. While the ICeloToken interface does not
// necessarily include the IERC20 interface, ICeloToken is intended to be used
// alongside IERC20. Both are included here for ease of use.
type CeloTokenContract interface {
	Erc20Contract
	TransferWithComment(opts *bind.TransactOpts, to common.Address, value *big.Int, comment string) (*types.Transaction, error)
	Name(opts *bind.CallOpts) (string, error)
	Symbol(opts *bind.CallOpts) (string, error)
	Decimals(opts *bind.CallOpts) (uint8, error)
}
