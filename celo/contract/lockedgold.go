// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LockedGoldABI is the input ABI used to generate the binding from.
const LockedGoldABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"unlockingPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slashingWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalNonvoting\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"UnlockingPeriodSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GoldLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"available\",\"type\":\"uint256\"}],\"name\":\"GoldUnlocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GoldRelocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"GoldWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slasherIdentifier\",\"type\":\"string\"}],\"name\":\"SlasherWhitelistAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slasherIdentifier\",\"type\":\"string\"}],\"name\":\"SlasherWhitelistRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"slashed\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"penalty\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"reporter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"AccountSlashed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"name\":\"slasher\",\"type\":\"address\"}],\"name\":\"isSlasher\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"_unlockingPeriod\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setUnlockingPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"lock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"incrementNonvotingAccountBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decrementNonvotingAccountBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"relock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalLockedGold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNonvotingLockedGold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountTotalLockedGold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountNonvotingLockedGold\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPendingWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalPendingWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSlashingWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slasherIdentifier\",\"type\":\"string\"}],\"name\":\"addSlasher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"slasherIdentifier\",\"type\":\"string\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeSlasher\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"penalty\",\"type\":\"uint256\"},{\"name\":\"reporter\",\"type\":\"address\"},{\"name\":\"reward\",\"type\":\"uint256\"},{\"name\":\"lessers\",\"type\":\"address[]\"},{\"name\":\"greaters\",\"type\":\"address[]\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// LockedGold is an auto generated Go binding around an Ethereum contract.
type LockedGold struct {
	LockedGoldCaller     // Read-only binding to the contract
	LockedGoldTransactor // Write-only binding to the contract
	LockedGoldFilterer   // Log filterer for contract events
}

// LockedGoldCaller is an auto generated read-only Go binding around an Ethereum contract.
type LockedGoldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockedGoldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LockedGoldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockedGoldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LockedGoldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LockedGoldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LockedGoldSession struct {
	Contract     *LockedGold       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LockedGoldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LockedGoldCallerSession struct {
	Contract *LockedGoldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// LockedGoldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LockedGoldTransactorSession struct {
	Contract     *LockedGoldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// LockedGoldRaw is an auto generated low-level Go binding around an Ethereum contract.
type LockedGoldRaw struct {
	Contract *LockedGold // Generic contract binding to access the raw methods on
}

// LockedGoldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LockedGoldCallerRaw struct {
	Contract *LockedGoldCaller // Generic read-only contract binding to access the raw methods on
}

// LockedGoldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LockedGoldTransactorRaw struct {
	Contract *LockedGoldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLockedGold creates a new instance of LockedGold, bound to a specific deployed contract.
func NewLockedGold(address common.Address, backend bind.ContractBackend) (*LockedGold, error) {
	contract, err := bindLockedGold(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LockedGold{LockedGoldCaller: LockedGoldCaller{contract: contract}, LockedGoldTransactor: LockedGoldTransactor{contract: contract}, LockedGoldFilterer: LockedGoldFilterer{contract: contract}}, nil
}

// NewLockedGoldCaller creates a new read-only instance of LockedGold, bound to a specific deployed contract.
func NewLockedGoldCaller(address common.Address, caller bind.ContractCaller) (*LockedGoldCaller, error) {
	contract, err := bindLockedGold(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LockedGoldCaller{contract: contract}, nil
}

// NewLockedGoldTransactor creates a new write-only instance of LockedGold, bound to a specific deployed contract.
func NewLockedGoldTransactor(address common.Address, transactor bind.ContractTransactor) (*LockedGoldTransactor, error) {
	contract, err := bindLockedGold(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LockedGoldTransactor{contract: contract}, nil
}

// NewLockedGoldFilterer creates a new log filterer instance of LockedGold, bound to a specific deployed contract.
func NewLockedGoldFilterer(address common.Address, filterer bind.ContractFilterer) (*LockedGoldFilterer, error) {
	contract, err := bindLockedGold(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LockedGoldFilterer{contract: contract}, nil
}

// bindLockedGold binds a generic wrapper to an already deployed contract.
func bindLockedGold(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LockedGoldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseLockedGoldABI parses the ABI
func ParseLockedGoldABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(LockedGoldABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockedGold *LockedGoldRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockedGold.Contract.LockedGoldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockedGold *LockedGoldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedGold.Contract.LockedGoldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockedGold *LockedGoldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockedGold.Contract.LockedGoldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LockedGold *LockedGoldCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LockedGold.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LockedGold *LockedGoldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedGold.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LockedGold *LockedGoldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LockedGold.Contract.contract.Transact(opts, method, params...)
}

// GetAccountNonvotingLockedGold is a free data retrieval call binding the contract method 0x3f199b40.
//
// Solidity: function getAccountNonvotingLockedGold(address account) constant returns(uint256)
func (_LockedGold *LockedGoldCaller) GetAccountNonvotingLockedGold(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "getAccountNonvotingLockedGold", account)
	return *ret0, err
}

// GetAccountNonvotingLockedGold is a free data retrieval call binding the contract method 0x3f199b40.
//
// Solidity: function getAccountNonvotingLockedGold(address account) constant returns(uint256)
func (_LockedGold *LockedGoldSession) GetAccountNonvotingLockedGold(account common.Address) (*big.Int, error) {
	return _LockedGold.Contract.GetAccountNonvotingLockedGold(&_LockedGold.CallOpts, account)
}

// GetAccountNonvotingLockedGold is a free data retrieval call binding the contract method 0x3f199b40.
//
// Solidity: function getAccountNonvotingLockedGold(address account) constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) GetAccountNonvotingLockedGold(account common.Address) (*big.Int, error) {
	return _LockedGold.Contract.GetAccountNonvotingLockedGold(&_LockedGold.CallOpts, account)
}

// GetAccountTotalLockedGold is a free data retrieval call binding the contract method 0x30ec70f5.
//
// Solidity: function getAccountTotalLockedGold(address account) constant returns(uint256)
func (_LockedGold *LockedGoldCaller) GetAccountTotalLockedGold(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "getAccountTotalLockedGold", account)
	return *ret0, err
}

// GetAccountTotalLockedGold is a free data retrieval call binding the contract method 0x30ec70f5.
//
// Solidity: function getAccountTotalLockedGold(address account) constant returns(uint256)
func (_LockedGold *LockedGoldSession) GetAccountTotalLockedGold(account common.Address) (*big.Int, error) {
	return _LockedGold.Contract.GetAccountTotalLockedGold(&_LockedGold.CallOpts, account)
}

// GetAccountTotalLockedGold is a free data retrieval call binding the contract method 0x30ec70f5.
//
// Solidity: function getAccountTotalLockedGold(address account) constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) GetAccountTotalLockedGold(account common.Address) (*big.Int, error) {
	return _LockedGold.Contract.GetAccountTotalLockedGold(&_LockedGold.CallOpts, account)
}

// GetNonvotingLockedGold is a free data retrieval call binding the contract method 0x807876b7.
//
// Solidity: function getNonvotingLockedGold() constant returns(uint256)
func (_LockedGold *LockedGoldCaller) GetNonvotingLockedGold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "getNonvotingLockedGold")
	return *ret0, err
}

// GetNonvotingLockedGold is a free data retrieval call binding the contract method 0x807876b7.
//
// Solidity: function getNonvotingLockedGold() constant returns(uint256)
func (_LockedGold *LockedGoldSession) GetNonvotingLockedGold() (*big.Int, error) {
	return _LockedGold.Contract.GetNonvotingLockedGold(&_LockedGold.CallOpts)
}

// GetNonvotingLockedGold is a free data retrieval call binding the contract method 0x807876b7.
//
// Solidity: function getNonvotingLockedGold() constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) GetNonvotingLockedGold() (*big.Int, error) {
	return _LockedGold.Contract.GetNonvotingLockedGold(&_LockedGold.CallOpts)
}

// GetPendingWithdrawals is a free data retrieval call binding the contract method 0xf340c0d0.
//
// Solidity: function getPendingWithdrawals(address account) constant returns(uint256[], uint256[])
func (_LockedGold *LockedGoldCaller) GetPendingWithdrawals(opts *bind.CallOpts, account common.Address) ([]*big.Int, []*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _LockedGold.contract.Call(opts, out, "getPendingWithdrawals", account)
	return *ret0, *ret1, err
}

// GetPendingWithdrawals is a free data retrieval call binding the contract method 0xf340c0d0.
//
// Solidity: function getPendingWithdrawals(address account) constant returns(uint256[], uint256[])
func (_LockedGold *LockedGoldSession) GetPendingWithdrawals(account common.Address) ([]*big.Int, []*big.Int, error) {
	return _LockedGold.Contract.GetPendingWithdrawals(&_LockedGold.CallOpts, account)
}

// GetPendingWithdrawals is a free data retrieval call binding the contract method 0xf340c0d0.
//
// Solidity: function getPendingWithdrawals(address account) constant returns(uint256[], uint256[])
func (_LockedGold *LockedGoldCallerSession) GetPendingWithdrawals(account common.Address) ([]*big.Int, []*big.Int, error) {
	return _LockedGold.Contract.GetPendingWithdrawals(&_LockedGold.CallOpts, account)
}

// GetSlashingWhitelist is a free data retrieval call binding the contract method 0x08764ee2.
//
// Solidity: function getSlashingWhitelist() constant returns(bytes32[])
func (_LockedGold *LockedGoldCaller) GetSlashingWhitelist(opts *bind.CallOpts) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "getSlashingWhitelist")
	return *ret0, err
}

// GetSlashingWhitelist is a free data retrieval call binding the contract method 0x08764ee2.
//
// Solidity: function getSlashingWhitelist() constant returns(bytes32[])
func (_LockedGold *LockedGoldSession) GetSlashingWhitelist() ([][32]byte, error) {
	return _LockedGold.Contract.GetSlashingWhitelist(&_LockedGold.CallOpts)
}

// GetSlashingWhitelist is a free data retrieval call binding the contract method 0x08764ee2.
//
// Solidity: function getSlashingWhitelist() constant returns(bytes32[])
func (_LockedGold *LockedGoldCallerSession) GetSlashingWhitelist() ([][32]byte, error) {
	return _LockedGold.Contract.GetSlashingWhitelist(&_LockedGold.CallOpts)
}

// GetTotalLockedGold is a free data retrieval call binding the contract method 0x30a61d59.
//
// Solidity: function getTotalLockedGold() constant returns(uint256)
func (_LockedGold *LockedGoldCaller) GetTotalLockedGold(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "getTotalLockedGold")
	return *ret0, err
}

// GetTotalLockedGold is a free data retrieval call binding the contract method 0x30a61d59.
//
// Solidity: function getTotalLockedGold() constant returns(uint256)
func (_LockedGold *LockedGoldSession) GetTotalLockedGold() (*big.Int, error) {
	return _LockedGold.Contract.GetTotalLockedGold(&_LockedGold.CallOpts)
}

// GetTotalLockedGold is a free data retrieval call binding the contract method 0x30a61d59.
//
// Solidity: function getTotalLockedGold() constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) GetTotalLockedGold() (*big.Int, error) {
	return _LockedGold.Contract.GetTotalLockedGold(&_LockedGold.CallOpts)
}

// GetTotalPendingWithdrawals is a free data retrieval call binding the contract method 0xb6e1e49d.
//
// Solidity: function getTotalPendingWithdrawals(address account) constant returns(uint256)
func (_LockedGold *LockedGoldCaller) GetTotalPendingWithdrawals(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "getTotalPendingWithdrawals", account)
	return *ret0, err
}

// GetTotalPendingWithdrawals is a free data retrieval call binding the contract method 0xb6e1e49d.
//
// Solidity: function getTotalPendingWithdrawals(address account) constant returns(uint256)
func (_LockedGold *LockedGoldSession) GetTotalPendingWithdrawals(account common.Address) (*big.Int, error) {
	return _LockedGold.Contract.GetTotalPendingWithdrawals(&_LockedGold.CallOpts, account)
}

// GetTotalPendingWithdrawals is a free data retrieval call binding the contract method 0xb6e1e49d.
//
// Solidity: function getTotalPendingWithdrawals(address account) constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) GetTotalPendingWithdrawals(account common.Address) (*big.Int, error) {
	return _LockedGold.Contract.GetTotalPendingWithdrawals(&_LockedGold.CallOpts, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_LockedGold *LockedGoldCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_LockedGold *LockedGoldSession) Initialized() (bool, error) {
	return _LockedGold.Contract.Initialized(&_LockedGold.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_LockedGold *LockedGoldCallerSession) Initialized() (bool, error) {
	return _LockedGold.Contract.Initialized(&_LockedGold.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_LockedGold *LockedGoldCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_LockedGold *LockedGoldSession) IsOwner() (bool, error) {
	return _LockedGold.Contract.IsOwner(&_LockedGold.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_LockedGold *LockedGoldCallerSession) IsOwner() (bool, error) {
	return _LockedGold.Contract.IsOwner(&_LockedGold.CallOpts)
}

// IsSlasher is a free data retrieval call binding the contract method 0x57601c5d.
//
// Solidity: function isSlasher(address slasher) constant returns(bool)
func (_LockedGold *LockedGoldCaller) IsSlasher(opts *bind.CallOpts, slasher common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "isSlasher", slasher)
	return *ret0, err
}

// IsSlasher is a free data retrieval call binding the contract method 0x57601c5d.
//
// Solidity: function isSlasher(address slasher) constant returns(bool)
func (_LockedGold *LockedGoldSession) IsSlasher(slasher common.Address) (bool, error) {
	return _LockedGold.Contract.IsSlasher(&_LockedGold.CallOpts, slasher)
}

// IsSlasher is a free data retrieval call binding the contract method 0x57601c5d.
//
// Solidity: function isSlasher(address slasher) constant returns(bool)
func (_LockedGold *LockedGoldCallerSession) IsSlasher(slasher common.Address) (bool, error) {
	return _LockedGold.Contract.IsSlasher(&_LockedGold.CallOpts, slasher)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockedGold *LockedGoldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockedGold *LockedGoldSession) Owner() (common.Address, error) {
	return _LockedGold.Contract.Owner(&_LockedGold.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_LockedGold *LockedGoldCallerSession) Owner() (common.Address, error) {
	return _LockedGold.Contract.Owner(&_LockedGold.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_LockedGold *LockedGoldCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_LockedGold *LockedGoldSession) Registry() (common.Address, error) {
	return _LockedGold.Contract.Registry(&_LockedGold.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_LockedGold *LockedGoldCallerSession) Registry() (common.Address, error) {
	return _LockedGold.Contract.Registry(&_LockedGold.CallOpts)
}

// SlashingWhitelist is a free data retrieval call binding the contract method 0x6adcc938.
//
// Solidity: function slashingWhitelist(uint256 ) constant returns(bytes32)
func (_LockedGold *LockedGoldCaller) SlashingWhitelist(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "slashingWhitelist", arg0)
	return *ret0, err
}

// SlashingWhitelist is a free data retrieval call binding the contract method 0x6adcc938.
//
// Solidity: function slashingWhitelist(uint256 ) constant returns(bytes32)
func (_LockedGold *LockedGoldSession) SlashingWhitelist(arg0 *big.Int) ([32]byte, error) {
	return _LockedGold.Contract.SlashingWhitelist(&_LockedGold.CallOpts, arg0)
}

// SlashingWhitelist is a free data retrieval call binding the contract method 0x6adcc938.
//
// Solidity: function slashingWhitelist(uint256 ) constant returns(bytes32)
func (_LockedGold *LockedGoldCallerSession) SlashingWhitelist(arg0 *big.Int) ([32]byte, error) {
	return _LockedGold.Contract.SlashingWhitelist(&_LockedGold.CallOpts, arg0)
}

// TotalNonvoting is a free data retrieval call binding the contract method 0xc1867f6d.
//
// Solidity: function totalNonvoting() constant returns(uint256)
func (_LockedGold *LockedGoldCaller) TotalNonvoting(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "totalNonvoting")
	return *ret0, err
}

// TotalNonvoting is a free data retrieval call binding the contract method 0xc1867f6d.
//
// Solidity: function totalNonvoting() constant returns(uint256)
func (_LockedGold *LockedGoldSession) TotalNonvoting() (*big.Int, error) {
	return _LockedGold.Contract.TotalNonvoting(&_LockedGold.CallOpts)
}

// TotalNonvoting is a free data retrieval call binding the contract method 0xc1867f6d.
//
// Solidity: function totalNonvoting() constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) TotalNonvoting() (*big.Int, error) {
	return _LockedGold.Contract.TotalNonvoting(&_LockedGold.CallOpts)
}

// UnlockingPeriod is a free data retrieval call binding the contract method 0x20637d8e.
//
// Solidity: function unlockingPeriod() constant returns(uint256)
func (_LockedGold *LockedGoldCaller) UnlockingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LockedGold.contract.Call(opts, out, "unlockingPeriod")
	return *ret0, err
}

// UnlockingPeriod is a free data retrieval call binding the contract method 0x20637d8e.
//
// Solidity: function unlockingPeriod() constant returns(uint256)
func (_LockedGold *LockedGoldSession) UnlockingPeriod() (*big.Int, error) {
	return _LockedGold.Contract.UnlockingPeriod(&_LockedGold.CallOpts)
}

// UnlockingPeriod is a free data retrieval call binding the contract method 0x20637d8e.
//
// Solidity: function unlockingPeriod() constant returns(uint256)
func (_LockedGold *LockedGoldCallerSession) UnlockingPeriod() (*big.Int, error) {
	return _LockedGold.Contract.UnlockingPeriod(&_LockedGold.CallOpts)
}

// AddSlasher is a paid mutator transaction binding the contract method 0x64891198.
//
// Solidity: function addSlasher(string slasherIdentifier) returns()
func (_LockedGold *LockedGoldTransactor) AddSlasher(opts *bind.TransactOpts, slasherIdentifier string) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "addSlasher", slasherIdentifier)
}

// AddSlasher is a paid mutator transaction binding the contract method 0x64891198.
//
// Solidity: function addSlasher(string slasherIdentifier) returns()
func (_LockedGold *LockedGoldSession) AddSlasher(slasherIdentifier string) (*types.Transaction, error) {
	return _LockedGold.Contract.AddSlasher(&_LockedGold.TransactOpts, slasherIdentifier)
}

// AddSlasher is a paid mutator transaction binding the contract method 0x64891198.
//
// Solidity: function addSlasher(string slasherIdentifier) returns()
func (_LockedGold *LockedGoldTransactorSession) AddSlasher(slasherIdentifier string) (*types.Transaction, error) {
	return _LockedGold.Contract.AddSlasher(&_LockedGold.TransactOpts, slasherIdentifier)
}

// DecrementNonvotingAccountBalance is a paid mutator transaction binding the contract method 0x18a4ff8c.
//
// Solidity: function decrementNonvotingAccountBalance(address account, uint256 value) returns()
func (_LockedGold *LockedGoldTransactor) DecrementNonvotingAccountBalance(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "decrementNonvotingAccountBalance", account, value)
}

// DecrementNonvotingAccountBalance is a paid mutator transaction binding the contract method 0x18a4ff8c.
//
// Solidity: function decrementNonvotingAccountBalance(address account, uint256 value) returns()
func (_LockedGold *LockedGoldSession) DecrementNonvotingAccountBalance(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.DecrementNonvotingAccountBalance(&_LockedGold.TransactOpts, account, value)
}

// DecrementNonvotingAccountBalance is a paid mutator transaction binding the contract method 0x18a4ff8c.
//
// Solidity: function decrementNonvotingAccountBalance(address account, uint256 value) returns()
func (_LockedGold *LockedGoldTransactorSession) DecrementNonvotingAccountBalance(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.DecrementNonvotingAccountBalance(&_LockedGold.TransactOpts, account, value)
}

// IncrementNonvotingAccountBalance is a paid mutator transaction binding the contract method 0x6edf77a5.
//
// Solidity: function incrementNonvotingAccountBalance(address account, uint256 value) returns()
func (_LockedGold *LockedGoldTransactor) IncrementNonvotingAccountBalance(opts *bind.TransactOpts, account common.Address, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "incrementNonvotingAccountBalance", account, value)
}

// IncrementNonvotingAccountBalance is a paid mutator transaction binding the contract method 0x6edf77a5.
//
// Solidity: function incrementNonvotingAccountBalance(address account, uint256 value) returns()
func (_LockedGold *LockedGoldSession) IncrementNonvotingAccountBalance(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.IncrementNonvotingAccountBalance(&_LockedGold.TransactOpts, account, value)
}

// IncrementNonvotingAccountBalance is a paid mutator transaction binding the contract method 0x6edf77a5.
//
// Solidity: function incrementNonvotingAccountBalance(address account, uint256 value) returns()
func (_LockedGold *LockedGoldTransactorSession) IncrementNonvotingAccountBalance(account common.Address, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.IncrementNonvotingAccountBalance(&_LockedGold.TransactOpts, account, value)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address registryAddress, uint256 _unlockingPeriod) returns()
func (_LockedGold *LockedGoldTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _unlockingPeriod *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "initialize", registryAddress, _unlockingPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address registryAddress, uint256 _unlockingPeriod) returns()
func (_LockedGold *LockedGoldSession) Initialize(registryAddress common.Address, _unlockingPeriod *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Initialize(&_LockedGold.TransactOpts, registryAddress, _unlockingPeriod)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address registryAddress, uint256 _unlockingPeriod) returns()
func (_LockedGold *LockedGoldTransactorSession) Initialize(registryAddress common.Address, _unlockingPeriod *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Initialize(&_LockedGold.TransactOpts, registryAddress, _unlockingPeriod)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_LockedGold *LockedGoldTransactor) Lock(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "lock")
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_LockedGold *LockedGoldSession) Lock() (*types.Transaction, error) {
	return _LockedGold.Contract.Lock(&_LockedGold.TransactOpts)
}

// Lock is a paid mutator transaction binding the contract method 0xf83d08ba.
//
// Solidity: function lock() returns()
func (_LockedGold *LockedGoldTransactorSession) Lock() (*types.Transaction, error) {
	return _LockedGold.Contract.Lock(&_LockedGold.TransactOpts)
}

// Relock is a paid mutator transaction binding the contract method 0xb2fb30cb.
//
// Solidity: function relock(uint256 index, uint256 value) returns()
func (_LockedGold *LockedGoldTransactor) Relock(opts *bind.TransactOpts, index *big.Int, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "relock", index, value)
}

// Relock is a paid mutator transaction binding the contract method 0xb2fb30cb.
//
// Solidity: function relock(uint256 index, uint256 value) returns()
func (_LockedGold *LockedGoldSession) Relock(index *big.Int, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Relock(&_LockedGold.TransactOpts, index, value)
}

// Relock is a paid mutator transaction binding the contract method 0xb2fb30cb.
//
// Solidity: function relock(uint256 index, uint256 value) returns()
func (_LockedGold *LockedGoldTransactorSession) Relock(index *big.Int, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Relock(&_LockedGold.TransactOpts, index, value)
}

// RemoveSlasher is a paid mutator transaction binding the contract method 0x1fe2dfda.
//
// Solidity: function removeSlasher(string slasherIdentifier, uint256 index) returns()
func (_LockedGold *LockedGoldTransactor) RemoveSlasher(opts *bind.TransactOpts, slasherIdentifier string, index *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "removeSlasher", slasherIdentifier, index)
}

// RemoveSlasher is a paid mutator transaction binding the contract method 0x1fe2dfda.
//
// Solidity: function removeSlasher(string slasherIdentifier, uint256 index) returns()
func (_LockedGold *LockedGoldSession) RemoveSlasher(slasherIdentifier string, index *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.RemoveSlasher(&_LockedGold.TransactOpts, slasherIdentifier, index)
}

// RemoveSlasher is a paid mutator transaction binding the contract method 0x1fe2dfda.
//
// Solidity: function removeSlasher(string slasherIdentifier, uint256 index) returns()
func (_LockedGold *LockedGoldTransactorSession) RemoveSlasher(slasherIdentifier string, index *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.RemoveSlasher(&_LockedGold.TransactOpts, slasherIdentifier, index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LockedGold *LockedGoldTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LockedGold *LockedGoldSession) RenounceOwnership() (*types.Transaction, error) {
	return _LockedGold.Contract.RenounceOwnership(&_LockedGold.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_LockedGold *LockedGoldTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _LockedGold.Contract.RenounceOwnership(&_LockedGold.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_LockedGold *LockedGoldTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_LockedGold *LockedGoldSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _LockedGold.Contract.SetRegistry(&_LockedGold.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_LockedGold *LockedGoldTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _LockedGold.Contract.SetRegistry(&_LockedGold.TransactOpts, registryAddress)
}

// SetUnlockingPeriod is a paid mutator transaction binding the contract method 0x66f0633b.
//
// Solidity: function setUnlockingPeriod(uint256 value) returns()
func (_LockedGold *LockedGoldTransactor) SetUnlockingPeriod(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "setUnlockingPeriod", value)
}

// SetUnlockingPeriod is a paid mutator transaction binding the contract method 0x66f0633b.
//
// Solidity: function setUnlockingPeriod(uint256 value) returns()
func (_LockedGold *LockedGoldSession) SetUnlockingPeriod(value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.SetUnlockingPeriod(&_LockedGold.TransactOpts, value)
}

// SetUnlockingPeriod is a paid mutator transaction binding the contract method 0x66f0633b.
//
// Solidity: function setUnlockingPeriod(uint256 value) returns()
func (_LockedGold *LockedGoldTransactorSession) SetUnlockingPeriod(value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.SetUnlockingPeriod(&_LockedGold.TransactOpts, value)
}

// Slash is a paid mutator transaction binding the contract method 0x31993fc9.
//
// Solidity: function slash(address account, uint256 penalty, address reporter, uint256 reward, address[] lessers, address[] greaters, uint256[] indices) returns()
func (_LockedGold *LockedGoldTransactor) Slash(opts *bind.TransactOpts, account common.Address, penalty *big.Int, reporter common.Address, reward *big.Int, lessers []common.Address, greaters []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "slash", account, penalty, reporter, reward, lessers, greaters, indices)
}

// Slash is a paid mutator transaction binding the contract method 0x31993fc9.
//
// Solidity: function slash(address account, uint256 penalty, address reporter, uint256 reward, address[] lessers, address[] greaters, uint256[] indices) returns()
func (_LockedGold *LockedGoldSession) Slash(account common.Address, penalty *big.Int, reporter common.Address, reward *big.Int, lessers []common.Address, greaters []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Slash(&_LockedGold.TransactOpts, account, penalty, reporter, reward, lessers, greaters, indices)
}

// Slash is a paid mutator transaction binding the contract method 0x31993fc9.
//
// Solidity: function slash(address account, uint256 penalty, address reporter, uint256 reward, address[] lessers, address[] greaters, uint256[] indices) returns()
func (_LockedGold *LockedGoldTransactorSession) Slash(account common.Address, penalty *big.Int, reporter common.Address, reward *big.Int, lessers []common.Address, greaters []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Slash(&_LockedGold.TransactOpts, account, penalty, reporter, reward, lessers, greaters, indices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LockedGold *LockedGoldTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LockedGold *LockedGoldSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LockedGold.Contract.TransferOwnership(&_LockedGold.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_LockedGold *LockedGoldTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _LockedGold.Contract.TransferOwnership(&_LockedGold.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 value) returns()
func (_LockedGold *LockedGoldTransactor) Unlock(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "unlock", value)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 value) returns()
func (_LockedGold *LockedGoldSession) Unlock(value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Unlock(&_LockedGold.TransactOpts, value)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 value) returns()
func (_LockedGold *LockedGoldTransactorSession) Unlock(value *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Unlock(&_LockedGold.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 index) returns()
func (_LockedGold *LockedGoldTransactor) Withdraw(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _LockedGold.contract.Transact(opts, "withdraw", index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 index) returns()
func (_LockedGold *LockedGoldSession) Withdraw(index *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Withdraw(&_LockedGold.TransactOpts, index)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 index) returns()
func (_LockedGold *LockedGoldTransactorSession) Withdraw(index *big.Int) (*types.Transaction, error) {
	return _LockedGold.Contract.Withdraw(&_LockedGold.TransactOpts, index)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_LockedGold *LockedGoldFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _LockedGold.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AccountSlashed":
		event, err = _LockedGold.ParseAccountSlashed(log)
	case "GoldLocked":
		event, err = _LockedGold.ParseGoldLocked(log)
	case "GoldRelocked":
		event, err = _LockedGold.ParseGoldRelocked(log)
	case "GoldUnlocked":
		event, err = _LockedGold.ParseGoldUnlocked(log)
	case "GoldWithdrawn":
		event, err = _LockedGold.ParseGoldWithdrawn(log)
	case "OwnershipTransferred":
		event, err = _LockedGold.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _LockedGold.ParseRegistrySet(log)
	case "SlasherWhitelistAdded":
		event, err = _LockedGold.ParseSlasherWhitelistAdded(log)
	case "SlasherWhitelistRemoved":
		event, err = _LockedGold.ParseSlasherWhitelistRemoved(log)
	case "UnlockingPeriodSet":
		event, err = _LockedGold.ParseUnlockingPeriodSet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// LockedGoldAccountSlashedIterator is returned from FilterAccountSlashed and is used to iterate over the raw logs and unpacked data for AccountSlashed events raised by the LockedGold contract.
type LockedGoldAccountSlashedIterator struct {
	Event *LockedGoldAccountSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldAccountSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldAccountSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldAccountSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldAccountSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldAccountSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldAccountSlashed represents a AccountSlashed event raised by the LockedGold contract.
type LockedGoldAccountSlashed struct {
	Slashed  common.Address
	Penalty  *big.Int
	Reporter common.Address
	Reward   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAccountSlashed is a free log retrieval operation binding the contract event 0x7abcb995a115c34a67528d58d5fc5ce02c22cb835ce1685046163f7d366d7111.
//
// Solidity: event AccountSlashed(address indexed slashed, uint256 penalty, address indexed reporter, uint256 reward)
func (_LockedGold *LockedGoldFilterer) FilterAccountSlashed(opts *bind.FilterOpts, slashed []common.Address, reporter []common.Address) (*LockedGoldAccountSlashedIterator, error) {

	var slashedRule []interface{}
	for _, slashedItem := range slashed {
		slashedRule = append(slashedRule, slashedItem)
	}

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "AccountSlashed", slashedRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldAccountSlashedIterator{contract: _LockedGold.contract, event: "AccountSlashed", logs: logs, sub: sub}, nil
}

// WatchAccountSlashed is a free log subscription operation binding the contract event 0x7abcb995a115c34a67528d58d5fc5ce02c22cb835ce1685046163f7d366d7111.
//
// Solidity: event AccountSlashed(address indexed slashed, uint256 penalty, address indexed reporter, uint256 reward)
func (_LockedGold *LockedGoldFilterer) WatchAccountSlashed(opts *bind.WatchOpts, sink chan<- *LockedGoldAccountSlashed, slashed []common.Address, reporter []common.Address) (event.Subscription, error) {

	var slashedRule []interface{}
	for _, slashedItem := range slashed {
		slashedRule = append(slashedRule, slashedItem)
	}

	var reporterRule []interface{}
	for _, reporterItem := range reporter {
		reporterRule = append(reporterRule, reporterItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "AccountSlashed", slashedRule, reporterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldAccountSlashed)
				if err := _LockedGold.contract.UnpackLog(event, "AccountSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAccountSlashed is a log parse operation binding the contract event 0x7abcb995a115c34a67528d58d5fc5ce02c22cb835ce1685046163f7d366d7111.
//
// Solidity: event AccountSlashed(address indexed slashed, uint256 penalty, address indexed reporter, uint256 reward)
func (_LockedGold *LockedGoldFilterer) ParseAccountSlashed(log types.Log) (*LockedGoldAccountSlashed, error) {
	event := new(LockedGoldAccountSlashed)
	if err := _LockedGold.contract.UnpackLog(event, "AccountSlashed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldGoldLockedIterator is returned from FilterGoldLocked and is used to iterate over the raw logs and unpacked data for GoldLocked events raised by the LockedGold contract.
type LockedGoldGoldLockedIterator struct {
	Event *LockedGoldGoldLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldGoldLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldGoldLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldGoldLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldGoldLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldGoldLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldGoldLocked represents a GoldLocked event raised by the LockedGold contract.
type LockedGoldGoldLocked struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGoldLocked is a free log retrieval operation binding the contract event 0x0f0f2fc5b4c987a49e1663ce2c2d65de12f3b701ff02b4d09461421e63e609e7.
//
// Solidity: event GoldLocked(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) FilterGoldLocked(opts *bind.FilterOpts, account []common.Address) (*LockedGoldGoldLockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "GoldLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldGoldLockedIterator{contract: _LockedGold.contract, event: "GoldLocked", logs: logs, sub: sub}, nil
}

// WatchGoldLocked is a free log subscription operation binding the contract event 0x0f0f2fc5b4c987a49e1663ce2c2d65de12f3b701ff02b4d09461421e63e609e7.
//
// Solidity: event GoldLocked(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) WatchGoldLocked(opts *bind.WatchOpts, sink chan<- *LockedGoldGoldLocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "GoldLocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldGoldLocked)
				if err := _LockedGold.contract.UnpackLog(event, "GoldLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGoldLocked is a log parse operation binding the contract event 0x0f0f2fc5b4c987a49e1663ce2c2d65de12f3b701ff02b4d09461421e63e609e7.
//
// Solidity: event GoldLocked(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) ParseGoldLocked(log types.Log) (*LockedGoldGoldLocked, error) {
	event := new(LockedGoldGoldLocked)
	if err := _LockedGold.contract.UnpackLog(event, "GoldLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldGoldRelockedIterator is returned from FilterGoldRelocked and is used to iterate over the raw logs and unpacked data for GoldRelocked events raised by the LockedGold contract.
type LockedGoldGoldRelockedIterator struct {
	Event *LockedGoldGoldRelocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldGoldRelockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldGoldRelocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldGoldRelocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldGoldRelockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldGoldRelockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldGoldRelocked represents a GoldRelocked event raised by the LockedGold contract.
type LockedGoldGoldRelocked struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGoldRelocked is a free log retrieval operation binding the contract event 0xa823fc38a01c2f76d7057a79bb5c317710f26f7dbdea78634598d5519d0f7cb0.
//
// Solidity: event GoldRelocked(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) FilterGoldRelocked(opts *bind.FilterOpts, account []common.Address) (*LockedGoldGoldRelockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "GoldRelocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldGoldRelockedIterator{contract: _LockedGold.contract, event: "GoldRelocked", logs: logs, sub: sub}, nil
}

// WatchGoldRelocked is a free log subscription operation binding the contract event 0xa823fc38a01c2f76d7057a79bb5c317710f26f7dbdea78634598d5519d0f7cb0.
//
// Solidity: event GoldRelocked(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) WatchGoldRelocked(opts *bind.WatchOpts, sink chan<- *LockedGoldGoldRelocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "GoldRelocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldGoldRelocked)
				if err := _LockedGold.contract.UnpackLog(event, "GoldRelocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGoldRelocked is a log parse operation binding the contract event 0xa823fc38a01c2f76d7057a79bb5c317710f26f7dbdea78634598d5519d0f7cb0.
//
// Solidity: event GoldRelocked(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) ParseGoldRelocked(log types.Log) (*LockedGoldGoldRelocked, error) {
	event := new(LockedGoldGoldRelocked)
	if err := _LockedGold.contract.UnpackLog(event, "GoldRelocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldGoldUnlockedIterator is returned from FilterGoldUnlocked and is used to iterate over the raw logs and unpacked data for GoldUnlocked events raised by the LockedGold contract.
type LockedGoldGoldUnlockedIterator struct {
	Event *LockedGoldGoldUnlocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldGoldUnlockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldGoldUnlocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldGoldUnlocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldGoldUnlockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldGoldUnlockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldGoldUnlocked represents a GoldUnlocked event raised by the LockedGold contract.
type LockedGoldGoldUnlocked struct {
	Account   common.Address
	Value     *big.Int
	Available *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGoldUnlocked is a free log retrieval operation binding the contract event 0xb1a3aef2a332070da206ad1868a5e327f5aa5144e00e9a7b40717c153158a588.
//
// Solidity: event GoldUnlocked(address indexed account, uint256 value, uint256 available)
func (_LockedGold *LockedGoldFilterer) FilterGoldUnlocked(opts *bind.FilterOpts, account []common.Address) (*LockedGoldGoldUnlockedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "GoldUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldGoldUnlockedIterator{contract: _LockedGold.contract, event: "GoldUnlocked", logs: logs, sub: sub}, nil
}

// WatchGoldUnlocked is a free log subscription operation binding the contract event 0xb1a3aef2a332070da206ad1868a5e327f5aa5144e00e9a7b40717c153158a588.
//
// Solidity: event GoldUnlocked(address indexed account, uint256 value, uint256 available)
func (_LockedGold *LockedGoldFilterer) WatchGoldUnlocked(opts *bind.WatchOpts, sink chan<- *LockedGoldGoldUnlocked, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "GoldUnlocked", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldGoldUnlocked)
				if err := _LockedGold.contract.UnpackLog(event, "GoldUnlocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGoldUnlocked is a log parse operation binding the contract event 0xb1a3aef2a332070da206ad1868a5e327f5aa5144e00e9a7b40717c153158a588.
//
// Solidity: event GoldUnlocked(address indexed account, uint256 value, uint256 available)
func (_LockedGold *LockedGoldFilterer) ParseGoldUnlocked(log types.Log) (*LockedGoldGoldUnlocked, error) {
	event := new(LockedGoldGoldUnlocked)
	if err := _LockedGold.contract.UnpackLog(event, "GoldUnlocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldGoldWithdrawnIterator is returned from FilterGoldWithdrawn and is used to iterate over the raw logs and unpacked data for GoldWithdrawn events raised by the LockedGold contract.
type LockedGoldGoldWithdrawnIterator struct {
	Event *LockedGoldGoldWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldGoldWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldGoldWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldGoldWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldGoldWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldGoldWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldGoldWithdrawn represents a GoldWithdrawn event raised by the LockedGold contract.
type LockedGoldGoldWithdrawn struct {
	Account common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGoldWithdrawn is a free log retrieval operation binding the contract event 0x292d39ba701489b7f640c83806d3eeabe0a32c9f0a61b49e95612ebad42211cd.
//
// Solidity: event GoldWithdrawn(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) FilterGoldWithdrawn(opts *bind.FilterOpts, account []common.Address) (*LockedGoldGoldWithdrawnIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "GoldWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldGoldWithdrawnIterator{contract: _LockedGold.contract, event: "GoldWithdrawn", logs: logs, sub: sub}, nil
}

// WatchGoldWithdrawn is a free log subscription operation binding the contract event 0x292d39ba701489b7f640c83806d3eeabe0a32c9f0a61b49e95612ebad42211cd.
//
// Solidity: event GoldWithdrawn(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) WatchGoldWithdrawn(opts *bind.WatchOpts, sink chan<- *LockedGoldGoldWithdrawn, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "GoldWithdrawn", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldGoldWithdrawn)
				if err := _LockedGold.contract.UnpackLog(event, "GoldWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGoldWithdrawn is a log parse operation binding the contract event 0x292d39ba701489b7f640c83806d3eeabe0a32c9f0a61b49e95612ebad42211cd.
//
// Solidity: event GoldWithdrawn(address indexed account, uint256 value)
func (_LockedGold *LockedGoldFilterer) ParseGoldWithdrawn(log types.Log) (*LockedGoldGoldWithdrawn, error) {
	event := new(LockedGoldGoldWithdrawn)
	if err := _LockedGold.contract.UnpackLog(event, "GoldWithdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LockedGold contract.
type LockedGoldOwnershipTransferredIterator struct {
	Event *LockedGoldOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldOwnershipTransferred represents a OwnershipTransferred event raised by the LockedGold contract.
type LockedGoldOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LockedGold *LockedGoldFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LockedGoldOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldOwnershipTransferredIterator{contract: _LockedGold.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LockedGold *LockedGoldFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LockedGoldOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldOwnershipTransferred)
				if err := _LockedGold.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LockedGold *LockedGoldFilterer) ParseOwnershipTransferred(log types.Log) (*LockedGoldOwnershipTransferred, error) {
	event := new(LockedGoldOwnershipTransferred)
	if err := _LockedGold.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the LockedGold contract.
type LockedGoldRegistrySetIterator struct {
	Event *LockedGoldRegistrySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldRegistrySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldRegistrySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldRegistrySet represents a RegistrySet event raised by the LockedGold contract.
type LockedGoldRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_LockedGold *LockedGoldFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*LockedGoldRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldRegistrySetIterator{contract: _LockedGold.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_LockedGold *LockedGoldFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *LockedGoldRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldRegistrySet)
				if err := _LockedGold.contract.UnpackLog(event, "RegistrySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrySet is a log parse operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_LockedGold *LockedGoldFilterer) ParseRegistrySet(log types.Log) (*LockedGoldRegistrySet, error) {
	event := new(LockedGoldRegistrySet)
	if err := _LockedGold.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldSlasherWhitelistAddedIterator is returned from FilterSlasherWhitelistAdded and is used to iterate over the raw logs and unpacked data for SlasherWhitelistAdded events raised by the LockedGold contract.
type LockedGoldSlasherWhitelistAddedIterator struct {
	Event *LockedGoldSlasherWhitelistAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldSlasherWhitelistAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldSlasherWhitelistAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldSlasherWhitelistAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldSlasherWhitelistAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldSlasherWhitelistAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldSlasherWhitelistAdded represents a SlasherWhitelistAdded event raised by the LockedGold contract.
type LockedGoldSlasherWhitelistAdded struct {
	SlasherIdentifier common.Hash
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSlasherWhitelistAdded is a free log retrieval operation binding the contract event 0x92a16cb9e1846d175c3007fc61953d186452c9ea1aa34183eb4b7f88cd3f07bb.
//
// Solidity: event SlasherWhitelistAdded(string indexed slasherIdentifier)
func (_LockedGold *LockedGoldFilterer) FilterSlasherWhitelistAdded(opts *bind.FilterOpts, slasherIdentifier []string) (*LockedGoldSlasherWhitelistAddedIterator, error) {

	var slasherIdentifierRule []interface{}
	for _, slasherIdentifierItem := range slasherIdentifier {
		slasherIdentifierRule = append(slasherIdentifierRule, slasherIdentifierItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "SlasherWhitelistAdded", slasherIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldSlasherWhitelistAddedIterator{contract: _LockedGold.contract, event: "SlasherWhitelistAdded", logs: logs, sub: sub}, nil
}

// WatchSlasherWhitelistAdded is a free log subscription operation binding the contract event 0x92a16cb9e1846d175c3007fc61953d186452c9ea1aa34183eb4b7f88cd3f07bb.
//
// Solidity: event SlasherWhitelistAdded(string indexed slasherIdentifier)
func (_LockedGold *LockedGoldFilterer) WatchSlasherWhitelistAdded(opts *bind.WatchOpts, sink chan<- *LockedGoldSlasherWhitelistAdded, slasherIdentifier []string) (event.Subscription, error) {

	var slasherIdentifierRule []interface{}
	for _, slasherIdentifierItem := range slasherIdentifier {
		slasherIdentifierRule = append(slasherIdentifierRule, slasherIdentifierItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "SlasherWhitelistAdded", slasherIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldSlasherWhitelistAdded)
				if err := _LockedGold.contract.UnpackLog(event, "SlasherWhitelistAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlasherWhitelistAdded is a log parse operation binding the contract event 0x92a16cb9e1846d175c3007fc61953d186452c9ea1aa34183eb4b7f88cd3f07bb.
//
// Solidity: event SlasherWhitelistAdded(string indexed slasherIdentifier)
func (_LockedGold *LockedGoldFilterer) ParseSlasherWhitelistAdded(log types.Log) (*LockedGoldSlasherWhitelistAdded, error) {
	event := new(LockedGoldSlasherWhitelistAdded)
	if err := _LockedGold.contract.UnpackLog(event, "SlasherWhitelistAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldSlasherWhitelistRemovedIterator is returned from FilterSlasherWhitelistRemoved and is used to iterate over the raw logs and unpacked data for SlasherWhitelistRemoved events raised by the LockedGold contract.
type LockedGoldSlasherWhitelistRemovedIterator struct {
	Event *LockedGoldSlasherWhitelistRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldSlasherWhitelistRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldSlasherWhitelistRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldSlasherWhitelistRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldSlasherWhitelistRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldSlasherWhitelistRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldSlasherWhitelistRemoved represents a SlasherWhitelistRemoved event raised by the LockedGold contract.
type LockedGoldSlasherWhitelistRemoved struct {
	SlasherIdentifier common.Hash
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSlasherWhitelistRemoved is a free log retrieval operation binding the contract event 0xaee8df56d95b5766042c2ff4dcb39a120f0a09dd21bb9c143f86a314eff4b714.
//
// Solidity: event SlasherWhitelistRemoved(string indexed slasherIdentifier)
func (_LockedGold *LockedGoldFilterer) FilterSlasherWhitelistRemoved(opts *bind.FilterOpts, slasherIdentifier []string) (*LockedGoldSlasherWhitelistRemovedIterator, error) {

	var slasherIdentifierRule []interface{}
	for _, slasherIdentifierItem := range slasherIdentifier {
		slasherIdentifierRule = append(slasherIdentifierRule, slasherIdentifierItem)
	}

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "SlasherWhitelistRemoved", slasherIdentifierRule)
	if err != nil {
		return nil, err
	}
	return &LockedGoldSlasherWhitelistRemovedIterator{contract: _LockedGold.contract, event: "SlasherWhitelistRemoved", logs: logs, sub: sub}, nil
}

// WatchSlasherWhitelistRemoved is a free log subscription operation binding the contract event 0xaee8df56d95b5766042c2ff4dcb39a120f0a09dd21bb9c143f86a314eff4b714.
//
// Solidity: event SlasherWhitelistRemoved(string indexed slasherIdentifier)
func (_LockedGold *LockedGoldFilterer) WatchSlasherWhitelistRemoved(opts *bind.WatchOpts, sink chan<- *LockedGoldSlasherWhitelistRemoved, slasherIdentifier []string) (event.Subscription, error) {

	var slasherIdentifierRule []interface{}
	for _, slasherIdentifierItem := range slasherIdentifier {
		slasherIdentifierRule = append(slasherIdentifierRule, slasherIdentifierItem)
	}

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "SlasherWhitelistRemoved", slasherIdentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldSlasherWhitelistRemoved)
				if err := _LockedGold.contract.UnpackLog(event, "SlasherWhitelistRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSlasherWhitelistRemoved is a log parse operation binding the contract event 0xaee8df56d95b5766042c2ff4dcb39a120f0a09dd21bb9c143f86a314eff4b714.
//
// Solidity: event SlasherWhitelistRemoved(string indexed slasherIdentifier)
func (_LockedGold *LockedGoldFilterer) ParseSlasherWhitelistRemoved(log types.Log) (*LockedGoldSlasherWhitelistRemoved, error) {
	event := new(LockedGoldSlasherWhitelistRemoved)
	if err := _LockedGold.contract.UnpackLog(event, "SlasherWhitelistRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LockedGoldUnlockingPeriodSetIterator is returned from FilterUnlockingPeriodSet and is used to iterate over the raw logs and unpacked data for UnlockingPeriodSet events raised by the LockedGold contract.
type LockedGoldUnlockingPeriodSetIterator struct {
	Event *LockedGoldUnlockingPeriodSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *LockedGoldUnlockingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LockedGoldUnlockingPeriodSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(LockedGoldUnlockingPeriodSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *LockedGoldUnlockingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LockedGoldUnlockingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LockedGoldUnlockingPeriodSet represents a UnlockingPeriodSet event raised by the LockedGold contract.
type LockedGoldUnlockingPeriodSet struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnlockingPeriodSet is a free log retrieval operation binding the contract event 0xd9274a7c98edc7c66931fc71872764091e7023fe3867358f8504d4c21b161fc5.
//
// Solidity: event UnlockingPeriodSet(uint256 period)
func (_LockedGold *LockedGoldFilterer) FilterUnlockingPeriodSet(opts *bind.FilterOpts) (*LockedGoldUnlockingPeriodSetIterator, error) {

	logs, sub, err := _LockedGold.contract.FilterLogs(opts, "UnlockingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &LockedGoldUnlockingPeriodSetIterator{contract: _LockedGold.contract, event: "UnlockingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchUnlockingPeriodSet is a free log subscription operation binding the contract event 0xd9274a7c98edc7c66931fc71872764091e7023fe3867358f8504d4c21b161fc5.
//
// Solidity: event UnlockingPeriodSet(uint256 period)
func (_LockedGold *LockedGoldFilterer) WatchUnlockingPeriodSet(opts *bind.WatchOpts, sink chan<- *LockedGoldUnlockingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _LockedGold.contract.WatchLogs(opts, "UnlockingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LockedGoldUnlockingPeriodSet)
				if err := _LockedGold.contract.UnpackLog(event, "UnlockingPeriodSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnlockingPeriodSet is a log parse operation binding the contract event 0xd9274a7c98edc7c66931fc71872764091e7023fe3867358f8504d4c21b161fc5.
//
// Solidity: event UnlockingPeriodSet(uint256 period)
func (_LockedGold *LockedGoldFilterer) ParseUnlockingPeriodSet(log types.Log) (*LockedGoldUnlockingPeriodSet, error) {
	event := new(LockedGoldUnlockingPeriodSet)
	if err := _LockedGold.contract.UnpackLog(event, "UnlockingPeriodSet", log); err != nil {
		return nil, err
	}
	return event, nil
}
