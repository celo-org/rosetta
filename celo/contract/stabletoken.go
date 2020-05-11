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

// StableTokenABI is the input ABI used to generate the binding from.
const StableTokenABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"blsKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"factor\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"name\":\"InflationFactorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"rate\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"updatePeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastUpdated\",\"type\":\"uint256\"}],\"name\":\"InflationParametersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"TransferComment\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"inflationRate\",\"type\":\"uint256\"},{\"name\":\"inflationFactorUpdatePeriod\",\"type\":\"uint256\"},{\"name\":\"initialBalanceAddresses\",\"type\":\"address[]\"},{\"name\":\"initialBalanceValues\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"rate\",\"type\":\"uint256\"},{\"name\":\"updatePeriod\",\"type\":\"uint256\"}],\"name\":\"setInflationParameters\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"comment\",\"type\":\"string\"}],\"name\":\"transferWithComment\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountOwner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountOwner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getInflationParameters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"valueToUnits\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"units\",\"type\":\"uint256\"}],\"name\":\"unitsToValue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"debitGasFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"feeRecipient\",\"type\":\"address\"},{\"name\":\"gatewayFeeRecipient\",\"type\":\"address\"},{\"name\":\"communityFund\",\"type\":\"address\"},{\"name\":\"refund\",\"type\":\"uint256\"},{\"name\":\"tipTxFee\",\"type\":\"uint256\"},{\"name\":\"gatewayFee\",\"type\":\"uint256\"},{\"name\":\"baseTxFee\",\"type\":\"uint256\"}],\"name\":\"creditGasFees\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StableToken is an auto generated Go binding around an Ethereum contract.
type StableToken struct {
	StableTokenCaller     // Read-only binding to the contract
	StableTokenTransactor // Write-only binding to the contract
	StableTokenFilterer   // Log filterer for contract events
}

// StableTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StableTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StableTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StableTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StableTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StableTokenSession struct {
	Contract     *StableToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StableTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StableTokenCallerSession struct {
	Contract *StableTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// StableTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StableTokenTransactorSession struct {
	Contract     *StableTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StableTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StableTokenRaw struct {
	Contract *StableToken // Generic contract binding to access the raw methods on
}

// StableTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StableTokenCallerRaw struct {
	Contract *StableTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StableTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StableTokenTransactorRaw struct {
	Contract *StableTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStableToken creates a new instance of StableToken, bound to a specific deployed contract.
func NewStableToken(address common.Address, backend bind.ContractBackend) (*StableToken, error) {
	contract, err := bindStableToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StableToken{StableTokenCaller: StableTokenCaller{contract: contract}, StableTokenTransactor: StableTokenTransactor{contract: contract}, StableTokenFilterer: StableTokenFilterer{contract: contract}}, nil
}

// NewStableTokenCaller creates a new read-only instance of StableToken, bound to a specific deployed contract.
func NewStableTokenCaller(address common.Address, caller bind.ContractCaller) (*StableTokenCaller, error) {
	contract, err := bindStableToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StableTokenCaller{contract: contract}, nil
}

// NewStableTokenTransactor creates a new write-only instance of StableToken, bound to a specific deployed contract.
func NewStableTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StableTokenTransactor, error) {
	contract, err := bindStableToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StableTokenTransactor{contract: contract}, nil
}

// NewStableTokenFilterer creates a new log filterer instance of StableToken, bound to a specific deployed contract.
func NewStableTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StableTokenFilterer, error) {
	contract, err := bindStableToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StableTokenFilterer{contract: contract}, nil
}

// bindStableToken binds a generic wrapper to an already deployed contract.
func bindStableToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StableTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseStableTokenABI parses the ABI
func ParseStableTokenABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(StableTokenABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableToken *StableTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StableToken.Contract.StableTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableToken *StableTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableToken.Contract.StableTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableToken *StableTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableToken.Contract.StableTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StableToken *StableTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _StableToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StableToken *StableTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StableToken *StableTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StableToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address accountOwner, address spender) constant returns(uint256)
func (_StableToken *StableTokenCaller) Allowance(opts *bind.CallOpts, accountOwner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "allowance", accountOwner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address accountOwner, address spender) constant returns(uint256)
func (_StableToken *StableTokenSession) Allowance(accountOwner common.Address, spender common.Address) (*big.Int, error) {
	return _StableToken.Contract.Allowance(&_StableToken.CallOpts, accountOwner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address accountOwner, address spender) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) Allowance(accountOwner common.Address, spender common.Address) (*big.Int, error) {
	return _StableToken.Contract.Allowance(&_StableToken.CallOpts, accountOwner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accountOwner) constant returns(uint256)
func (_StableToken *StableTokenCaller) BalanceOf(opts *bind.CallOpts, accountOwner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "balanceOf", accountOwner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accountOwner) constant returns(uint256)
func (_StableToken *StableTokenSession) BalanceOf(accountOwner common.Address) (*big.Int, error) {
	return _StableToken.Contract.BalanceOf(&_StableToken.CallOpts, accountOwner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address accountOwner) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) BalanceOf(accountOwner common.Address) (*big.Int, error) {
	return _StableToken.Contract.BalanceOf(&_StableToken.CallOpts, accountOwner)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_StableToken *StableTokenCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "checkProofOfPossession", sender, blsKey, blsPop)
	return *ret0, err
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_StableToken *StableTokenSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _StableToken.Contract.CheckProofOfPossession(&_StableToken.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_StableToken *StableTokenCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _StableToken.Contract.CheckProofOfPossession(&_StableToken.CallOpts, sender, blsKey, blsPop)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StableToken *StableTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StableToken *StableTokenSession) Decimals() (uint8, error) {
	return _StableToken.Contract.Decimals(&_StableToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_StableToken *StableTokenCallerSession) Decimals() (uint8, error) {
	return _StableToken.Contract.Decimals(&_StableToken.CallOpts)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_StableToken *StableTokenCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _StableToken.contract.Call(opts, out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
	return *ret0, *ret1, err
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_StableToken *StableTokenSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _StableToken.Contract.FractionMulExp(&_StableToken.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_StableToken *StableTokenCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _StableToken.Contract.FractionMulExp(&_StableToken.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_StableToken *StableTokenCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "getBlockNumberFromHeader", header)
	return *ret0, err
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_StableToken *StableTokenSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _StableToken.Contract.GetBlockNumberFromHeader(&_StableToken.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _StableToken.Contract.GetBlockNumberFromHeader(&_StableToken.CallOpts, header)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_StableToken *StableTokenCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "getEpochNumber")
	return *ret0, err
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_StableToken *StableTokenSession) GetEpochNumber() (*big.Int, error) {
	return _StableToken.Contract.GetEpochNumber(&_StableToken.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_StableToken *StableTokenCallerSession) GetEpochNumber() (*big.Int, error) {
	return _StableToken.Contract.GetEpochNumber(&_StableToken.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "getEpochNumberOfBlock", blockNumber)
	return *ret0, err
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _StableToken.Contract.GetEpochNumberOfBlock(&_StableToken.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _StableToken.Contract.GetEpochNumberOfBlock(&_StableToken.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_StableToken *StableTokenCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "getEpochSize")
	return *ret0, err
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_StableToken *StableTokenSession) GetEpochSize() (*big.Int, error) {
	return _StableToken.Contract.GetEpochSize(&_StableToken.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_StableToken *StableTokenCallerSession) GetEpochSize() (*big.Int, error) {
	return _StableToken.Contract.GetEpochSize(&_StableToken.CallOpts)
}

// GetInflationParameters is a free data retrieval call binding the contract method 0xa67f8747.
//
// Solidity: function getInflationParameters() constant returns(uint256, uint256, uint256, uint256)
func (_StableToken *StableTokenCaller) GetInflationParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _StableToken.contract.Call(opts, out, "getInflationParameters")
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetInflationParameters is a free data retrieval call binding the contract method 0xa67f8747.
//
// Solidity: function getInflationParameters() constant returns(uint256, uint256, uint256, uint256)
func (_StableToken *StableTokenSession) GetInflationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _StableToken.Contract.GetInflationParameters(&_StableToken.CallOpts)
}

// GetInflationParameters is a free data retrieval call binding the contract method 0xa67f8747.
//
// Solidity: function getInflationParameters() constant returns(uint256, uint256, uint256, uint256)
func (_StableToken *StableTokenCallerSession) GetInflationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _StableToken.Contract.GetInflationParameters(&_StableToken.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_StableToken *StableTokenCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "getParentSealBitmap", blockNumber)
	return *ret0, err
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_StableToken *StableTokenSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _StableToken.Contract.GetParentSealBitmap(&_StableToken.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_StableToken *StableTokenCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _StableToken.Contract.GetParentSealBitmap(&_StableToken.CallOpts, blockNumber)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_StableToken *StableTokenCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "getVerifiedSealBitmapFromHeader", header)
	return *ret0, err
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_StableToken *StableTokenSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _StableToken.Contract.GetVerifiedSealBitmapFromHeader(&_StableToken.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_StableToken *StableTokenCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _StableToken.Contract.GetVerifiedSealBitmapFromHeader(&_StableToken.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_StableToken *StableTokenCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "hashHeader", header)
	return *ret0, err
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_StableToken *StableTokenSession) HashHeader(header []byte) ([32]byte, error) {
	return _StableToken.Contract.HashHeader(&_StableToken.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_StableToken *StableTokenCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _StableToken.Contract.HashHeader(&_StableToken.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_StableToken *StableTokenCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_StableToken *StableTokenSession) Initialized() (bool, error) {
	return _StableToken.Contract.Initialized(&_StableToken.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_StableToken *StableTokenCallerSession) Initialized() (bool, error) {
	return _StableToken.Contract.Initialized(&_StableToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StableToken *StableTokenCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StableToken *StableTokenSession) IsOwner() (bool, error) {
	return _StableToken.Contract.IsOwner(&_StableToken.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_StableToken *StableTokenCallerSession) IsOwner() (bool, error) {
	return _StableToken.Contract.IsOwner(&_StableToken.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "minQuorumSize", blockNumber)
	return *ret0, err
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _StableToken.Contract.MinQuorumSize(&_StableToken.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _StableToken.Contract.MinQuorumSize(&_StableToken.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_StableToken *StableTokenCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "minQuorumSizeInCurrentSet")
	return *ret0, err
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_StableToken *StableTokenSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _StableToken.Contract.MinQuorumSizeInCurrentSet(&_StableToken.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_StableToken *StableTokenCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _StableToken.Contract.MinQuorumSizeInCurrentSet(&_StableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StableToken *StableTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StableToken *StableTokenSession) Name() (string, error) {
	return _StableToken.Contract.Name(&_StableToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_StableToken *StableTokenCallerSession) Name() (string, error) {
	return _StableToken.Contract.Name(&_StableToken.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_StableToken *StableTokenCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "numberValidatorsInCurrentSet")
	return *ret0, err
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_StableToken *StableTokenSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _StableToken.Contract.NumberValidatorsInCurrentSet(&_StableToken.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_StableToken *StableTokenCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _StableToken.Contract.NumberValidatorsInCurrentSet(&_StableToken.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "numberValidatorsInSet", blockNumber)
	return *ret0, err
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _StableToken.Contract.NumberValidatorsInSet(&_StableToken.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _StableToken.Contract.NumberValidatorsInSet(&_StableToken.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StableToken *StableTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StableToken *StableTokenSession) Owner() (common.Address, error) {
	return _StableToken.Contract.Owner(&_StableToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_StableToken *StableTokenCallerSession) Owner() (common.Address, error) {
	return _StableToken.Contract.Owner(&_StableToken.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_StableToken *StableTokenCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_StableToken *StableTokenSession) Registry() (common.Address, error) {
	return _StableToken.Contract.Registry(&_StableToken.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_StableToken *StableTokenCallerSession) Registry() (common.Address, error) {
	return _StableToken.Contract.Registry(&_StableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StableToken *StableTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StableToken *StableTokenSession) Symbol() (string, error) {
	return _StableToken.Contract.Symbol(&_StableToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_StableToken *StableTokenCallerSession) Symbol() (string, error) {
	return _StableToken.Contract.Symbol(&_StableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StableToken *StableTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StableToken *StableTokenSession) TotalSupply() (*big.Int, error) {
	return _StableToken.Contract.TotalSupply(&_StableToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_StableToken *StableTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StableToken.Contract.TotalSupply(&_StableToken.CallOpts)
}

// UnitsToValue is a free data retrieval call binding the contract method 0xaf31f587.
//
// Solidity: function unitsToValue(uint256 units) constant returns(uint256)
func (_StableToken *StableTokenCaller) UnitsToValue(opts *bind.CallOpts, units *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "unitsToValue", units)
	return *ret0, err
}

// UnitsToValue is a free data retrieval call binding the contract method 0xaf31f587.
//
// Solidity: function unitsToValue(uint256 units) constant returns(uint256)
func (_StableToken *StableTokenSession) UnitsToValue(units *big.Int) (*big.Int, error) {
	return _StableToken.Contract.UnitsToValue(&_StableToken.CallOpts, units)
}

// UnitsToValue is a free data retrieval call binding the contract method 0xaf31f587.
//
// Solidity: function unitsToValue(uint256 units) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) UnitsToValue(units *big.Int) (*big.Int, error) {
	return _StableToken.Contract.UnitsToValue(&_StableToken.CallOpts, units)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_StableToken *StableTokenCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "validatorSignerAddressFromCurrentSet", index)
	return *ret0, err
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_StableToken *StableTokenSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _StableToken.Contract.ValidatorSignerAddressFromCurrentSet(&_StableToken.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_StableToken *StableTokenCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _StableToken.Contract.ValidatorSignerAddressFromCurrentSet(&_StableToken.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_StableToken *StableTokenCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "validatorSignerAddressFromSet", index, blockNumber)
	return *ret0, err
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_StableToken *StableTokenSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _StableToken.Contract.ValidatorSignerAddressFromSet(&_StableToken.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_StableToken *StableTokenCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _StableToken.Contract.ValidatorSignerAddressFromSet(&_StableToken.CallOpts, index, blockNumber)
}

// ValueToUnits is a free data retrieval call binding the contract method 0x12c6c099.
//
// Solidity: function valueToUnits(uint256 value) constant returns(uint256)
func (_StableToken *StableTokenCaller) ValueToUnits(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _StableToken.contract.Call(opts, out, "valueToUnits", value)
	return *ret0, err
}

// ValueToUnits is a free data retrieval call binding the contract method 0x12c6c099.
//
// Solidity: function valueToUnits(uint256 value) constant returns(uint256)
func (_StableToken *StableTokenSession) ValueToUnits(value *big.Int) (*big.Int, error) {
	return _StableToken.Contract.ValueToUnits(&_StableToken.CallOpts, value)
}

// ValueToUnits is a free data retrieval call binding the contract method 0x12c6c099.
//
// Solidity: function valueToUnits(uint256 value) constant returns(uint256)
func (_StableToken *StableTokenCallerSession) ValueToUnits(value *big.Int) (*big.Int, error) {
	return _StableToken.Contract.ValueToUnits(&_StableToken.CallOpts, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Approve(&_StableToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Approve(&_StableToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_StableToken *StableTokenSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Burn(&_StableToken.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Burn(&_StableToken.TransactOpts, value)
}

// CreditGasFees is a paid mutator transaction binding the contract method 0x6a30b253.
//
// Solidity: function creditGasFees(address from, address feeRecipient, address gatewayFeeRecipient, address communityFund, uint256 refund, uint256 tipTxFee, uint256 gatewayFee, uint256 baseTxFee) returns()
func (_StableToken *StableTokenTransactor) CreditGasFees(opts *bind.TransactOpts, from common.Address, feeRecipient common.Address, gatewayFeeRecipient common.Address, communityFund common.Address, refund *big.Int, tipTxFee *big.Int, gatewayFee *big.Int, baseTxFee *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "creditGasFees", from, feeRecipient, gatewayFeeRecipient, communityFund, refund, tipTxFee, gatewayFee, baseTxFee)
}

// CreditGasFees is a paid mutator transaction binding the contract method 0x6a30b253.
//
// Solidity: function creditGasFees(address from, address feeRecipient, address gatewayFeeRecipient, address communityFund, uint256 refund, uint256 tipTxFee, uint256 gatewayFee, uint256 baseTxFee) returns()
func (_StableToken *StableTokenSession) CreditGasFees(from common.Address, feeRecipient common.Address, gatewayFeeRecipient common.Address, communityFund common.Address, refund *big.Int, tipTxFee *big.Int, gatewayFee *big.Int, baseTxFee *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.CreditGasFees(&_StableToken.TransactOpts, from, feeRecipient, gatewayFeeRecipient, communityFund, refund, tipTxFee, gatewayFee, baseTxFee)
}

// CreditGasFees is a paid mutator transaction binding the contract method 0x6a30b253.
//
// Solidity: function creditGasFees(address from, address feeRecipient, address gatewayFeeRecipient, address communityFund, uint256 refund, uint256 tipTxFee, uint256 gatewayFee, uint256 baseTxFee) returns()
func (_StableToken *StableTokenTransactorSession) CreditGasFees(from common.Address, feeRecipient common.Address, gatewayFeeRecipient common.Address, communityFund common.Address, refund *big.Int, tipTxFee *big.Int, gatewayFee *big.Int, baseTxFee *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.CreditGasFees(&_StableToken.TransactOpts, from, feeRecipient, gatewayFeeRecipient, communityFund, refund, tipTxFee, gatewayFee, baseTxFee)
}

// DebitGasFees is a paid mutator transaction binding the contract method 0x58cf9672.
//
// Solidity: function debitGasFees(address from, uint256 value) returns()
func (_StableToken *StableTokenTransactor) DebitGasFees(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "debitGasFees", from, value)
}

// DebitGasFees is a paid mutator transaction binding the contract method 0x58cf9672.
//
// Solidity: function debitGasFees(address from, uint256 value) returns()
func (_StableToken *StableTokenSession) DebitGasFees(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.DebitGasFees(&_StableToken.TransactOpts, from, value)
}

// DebitGasFees is a paid mutator transaction binding the contract method 0x58cf9672.
//
// Solidity: function debitGasFees(address from, uint256 value) returns()
func (_StableToken *StableTokenTransactorSession) DebitGasFees(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.DebitGasFees(&_StableToken.TransactOpts, from, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "decreaseAllowance", spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenSession) DecreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.DecreaseAllowance(&_StableToken.TransactOpts, spender, value)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) DecreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.DecreaseAllowance(&_StableToken.TransactOpts, spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "increaseAllowance", spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenSession) IncreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.IncreaseAllowance(&_StableToken.TransactOpts, spender, value)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) IncreaseAllowance(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.IncreaseAllowance(&_StableToken.TransactOpts, spender, value)
}

// Initialize is a paid mutator transaction binding the contract method 0x82b93516.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address registryAddress, uint256 inflationRate, uint256 inflationFactorUpdatePeriod, address[] initialBalanceAddresses, uint256[] initialBalanceValues) returns()
func (_StableToken *StableTokenTransactor) Initialize(opts *bind.TransactOpts, _name string, _symbol string, _decimals uint8, registryAddress common.Address, inflationRate *big.Int, inflationFactorUpdatePeriod *big.Int, initialBalanceAddresses []common.Address, initialBalanceValues []*big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "initialize", _name, _symbol, _decimals, registryAddress, inflationRate, inflationFactorUpdatePeriod, initialBalanceAddresses, initialBalanceValues)
}

// Initialize is a paid mutator transaction binding the contract method 0x82b93516.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address registryAddress, uint256 inflationRate, uint256 inflationFactorUpdatePeriod, address[] initialBalanceAddresses, uint256[] initialBalanceValues) returns()
func (_StableToken *StableTokenSession) Initialize(_name string, _symbol string, _decimals uint8, registryAddress common.Address, inflationRate *big.Int, inflationFactorUpdatePeriod *big.Int, initialBalanceAddresses []common.Address, initialBalanceValues []*big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Initialize(&_StableToken.TransactOpts, _name, _symbol, _decimals, registryAddress, inflationRate, inflationFactorUpdatePeriod, initialBalanceAddresses, initialBalanceValues)
}

// Initialize is a paid mutator transaction binding the contract method 0x82b93516.
//
// Solidity: function initialize(string _name, string _symbol, uint8 _decimals, address registryAddress, uint256 inflationRate, uint256 inflationFactorUpdatePeriod, address[] initialBalanceAddresses, uint256[] initialBalanceValues) returns()
func (_StableToken *StableTokenTransactorSession) Initialize(_name string, _symbol string, _decimals uint8, registryAddress common.Address, inflationRate *big.Int, inflationFactorUpdatePeriod *big.Int, initialBalanceAddresses []common.Address, initialBalanceValues []*big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Initialize(&_StableToken.TransactOpts, _name, _symbol, _decimals, registryAddress, inflationRate, inflationFactorUpdatePeriod, initialBalanceAddresses, initialBalanceValues)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_StableToken *StableTokenSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Mint(&_StableToken.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Mint(&_StableToken.TransactOpts, to, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StableToken *StableTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StableToken *StableTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _StableToken.Contract.RenounceOwnership(&_StableToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StableToken *StableTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StableToken.Contract.RenounceOwnership(&_StableToken.TransactOpts)
}

// SetInflationParameters is a paid mutator transaction binding the contract method 0x222836ad.
//
// Solidity: function setInflationParameters(uint256 rate, uint256 updatePeriod) returns()
func (_StableToken *StableTokenTransactor) SetInflationParameters(opts *bind.TransactOpts, rate *big.Int, updatePeriod *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "setInflationParameters", rate, updatePeriod)
}

// SetInflationParameters is a paid mutator transaction binding the contract method 0x222836ad.
//
// Solidity: function setInflationParameters(uint256 rate, uint256 updatePeriod) returns()
func (_StableToken *StableTokenSession) SetInflationParameters(rate *big.Int, updatePeriod *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.SetInflationParameters(&_StableToken.TransactOpts, rate, updatePeriod)
}

// SetInflationParameters is a paid mutator transaction binding the contract method 0x222836ad.
//
// Solidity: function setInflationParameters(uint256 rate, uint256 updatePeriod) returns()
func (_StableToken *StableTokenTransactorSession) SetInflationParameters(rate *big.Int, updatePeriod *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.SetInflationParameters(&_StableToken.TransactOpts, rate, updatePeriod)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_StableToken *StableTokenTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_StableToken *StableTokenSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _StableToken.Contract.SetRegistry(&_StableToken.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_StableToken *StableTokenTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _StableToken.Contract.SetRegistry(&_StableToken.TransactOpts, registryAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StableToken *StableTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Transfer(&_StableToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.Transfer(&_StableToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StableToken *StableTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.TransferFrom(&_StableToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StableToken *StableTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StableToken.Contract.TransferFrom(&_StableToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StableToken *StableTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StableToken *StableTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StableToken.Contract.TransferOwnership(&_StableToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StableToken *StableTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StableToken.Contract.TransferOwnership(&_StableToken.TransactOpts, newOwner)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_StableToken *StableTokenTransactor) TransferWithComment(opts *bind.TransactOpts, to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _StableToken.contract.Transact(opts, "transferWithComment", to, value, comment)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_StableToken *StableTokenSession) TransferWithComment(to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _StableToken.Contract.TransferWithComment(&_StableToken.TransactOpts, to, value, comment)
}

// TransferWithComment is a paid mutator transaction binding the contract method 0xe1d6aceb.
//
// Solidity: function transferWithComment(address to, uint256 value, string comment) returns(bool)
func (_StableToken *StableTokenTransactorSession) TransferWithComment(to common.Address, value *big.Int, comment string) (*types.Transaction, error) {
	return _StableToken.Contract.TransferWithComment(&_StableToken.TransactOpts, to, value, comment)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_StableToken *StableTokenFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _StableToken.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "Approval":
		event, err = _StableToken.ParseApproval(log)
	case "InflationFactorUpdated":
		event, err = _StableToken.ParseInflationFactorUpdated(log)
	case "InflationParametersUpdated":
		event, err = _StableToken.ParseInflationParametersUpdated(log)
	case "OwnershipTransferred":
		event, err = _StableToken.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _StableToken.ParseRegistrySet(log)
	case "Transfer":
		event, err = _StableToken.ParseTransfer(log)
	case "TransferComment":
		event, err = _StableToken.ParseTransferComment(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// StableTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StableToken contract.
type StableTokenApprovalIterator struct {
	Event *StableTokenApproval // Event containing the contract specifics and raw log

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
func (it *StableTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenApproval)
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
		it.Event = new(StableTokenApproval)
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
func (it *StableTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenApproval represents a Approval event raised by the StableToken contract.
type StableTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableToken *StableTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StableTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StableTokenApprovalIterator{contract: _StableToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableToken *StableTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StableTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenApproval)
				if err := _StableToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StableToken *StableTokenFilterer) ParseApproval(log types.Log) (*StableTokenApproval, error) {
	event := new(StableTokenApproval)
	if err := _StableToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableTokenInflationFactorUpdatedIterator is returned from FilterInflationFactorUpdated and is used to iterate over the raw logs and unpacked data for InflationFactorUpdated events raised by the StableToken contract.
type StableTokenInflationFactorUpdatedIterator struct {
	Event *StableTokenInflationFactorUpdated // Event containing the contract specifics and raw log

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
func (it *StableTokenInflationFactorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenInflationFactorUpdated)
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
		it.Event = new(StableTokenInflationFactorUpdated)
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
func (it *StableTokenInflationFactorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenInflationFactorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenInflationFactorUpdated represents a InflationFactorUpdated event raised by the StableToken contract.
type StableTokenInflationFactorUpdated struct {
	Factor      *big.Int
	LastUpdated *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterInflationFactorUpdated is a free log retrieval operation binding the contract event 0x08f3ed03ec9e579d1f6ab2f9e0d3dc661704696deabe37a6b6df7014f1b30a97.
//
// Solidity: event InflationFactorUpdated(uint256 factor, uint256 lastUpdated)
func (_StableToken *StableTokenFilterer) FilterInflationFactorUpdated(opts *bind.FilterOpts) (*StableTokenInflationFactorUpdatedIterator, error) {

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "InflationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return &StableTokenInflationFactorUpdatedIterator{contract: _StableToken.contract, event: "InflationFactorUpdated", logs: logs, sub: sub}, nil
}

// WatchInflationFactorUpdated is a free log subscription operation binding the contract event 0x08f3ed03ec9e579d1f6ab2f9e0d3dc661704696deabe37a6b6df7014f1b30a97.
//
// Solidity: event InflationFactorUpdated(uint256 factor, uint256 lastUpdated)
func (_StableToken *StableTokenFilterer) WatchInflationFactorUpdated(opts *bind.WatchOpts, sink chan<- *StableTokenInflationFactorUpdated) (event.Subscription, error) {

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "InflationFactorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenInflationFactorUpdated)
				if err := _StableToken.contract.UnpackLog(event, "InflationFactorUpdated", log); err != nil {
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

// ParseInflationFactorUpdated is a log parse operation binding the contract event 0x08f3ed03ec9e579d1f6ab2f9e0d3dc661704696deabe37a6b6df7014f1b30a97.
//
// Solidity: event InflationFactorUpdated(uint256 factor, uint256 lastUpdated)
func (_StableToken *StableTokenFilterer) ParseInflationFactorUpdated(log types.Log) (*StableTokenInflationFactorUpdated, error) {
	event := new(StableTokenInflationFactorUpdated)
	if err := _StableToken.contract.UnpackLog(event, "InflationFactorUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableTokenInflationParametersUpdatedIterator is returned from FilterInflationParametersUpdated and is used to iterate over the raw logs and unpacked data for InflationParametersUpdated events raised by the StableToken contract.
type StableTokenInflationParametersUpdatedIterator struct {
	Event *StableTokenInflationParametersUpdated // Event containing the contract specifics and raw log

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
func (it *StableTokenInflationParametersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenInflationParametersUpdated)
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
		it.Event = new(StableTokenInflationParametersUpdated)
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
func (it *StableTokenInflationParametersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenInflationParametersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenInflationParametersUpdated represents a InflationParametersUpdated event raised by the StableToken contract.
type StableTokenInflationParametersUpdated struct {
	Rate         *big.Int
	UpdatePeriod *big.Int
	LastUpdated  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInflationParametersUpdated is a free log retrieval operation binding the contract event 0xa0035d6667ffb7d387c86c7228141c4a877e8ed831b267ac928a2f5b651c155d.
//
// Solidity: event InflationParametersUpdated(uint256 rate, uint256 updatePeriod, uint256 lastUpdated)
func (_StableToken *StableTokenFilterer) FilterInflationParametersUpdated(opts *bind.FilterOpts) (*StableTokenInflationParametersUpdatedIterator, error) {

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "InflationParametersUpdated")
	if err != nil {
		return nil, err
	}
	return &StableTokenInflationParametersUpdatedIterator{contract: _StableToken.contract, event: "InflationParametersUpdated", logs: logs, sub: sub}, nil
}

// WatchInflationParametersUpdated is a free log subscription operation binding the contract event 0xa0035d6667ffb7d387c86c7228141c4a877e8ed831b267ac928a2f5b651c155d.
//
// Solidity: event InflationParametersUpdated(uint256 rate, uint256 updatePeriod, uint256 lastUpdated)
func (_StableToken *StableTokenFilterer) WatchInflationParametersUpdated(opts *bind.WatchOpts, sink chan<- *StableTokenInflationParametersUpdated) (event.Subscription, error) {

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "InflationParametersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenInflationParametersUpdated)
				if err := _StableToken.contract.UnpackLog(event, "InflationParametersUpdated", log); err != nil {
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

// ParseInflationParametersUpdated is a log parse operation binding the contract event 0xa0035d6667ffb7d387c86c7228141c4a877e8ed831b267ac928a2f5b651c155d.
//
// Solidity: event InflationParametersUpdated(uint256 rate, uint256 updatePeriod, uint256 lastUpdated)
func (_StableToken *StableTokenFilterer) ParseInflationParametersUpdated(log types.Log) (*StableTokenInflationParametersUpdated, error) {
	event := new(StableTokenInflationParametersUpdated)
	if err := _StableToken.contract.UnpackLog(event, "InflationParametersUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StableToken contract.
type StableTokenOwnershipTransferredIterator struct {
	Event *StableTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StableTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenOwnershipTransferred)
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
		it.Event = new(StableTokenOwnershipTransferred)
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
func (it *StableTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenOwnershipTransferred represents a OwnershipTransferred event raised by the StableToken contract.
type StableTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StableToken *StableTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StableTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StableTokenOwnershipTransferredIterator{contract: _StableToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StableToken *StableTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StableTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenOwnershipTransferred)
				if err := _StableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StableToken *StableTokenFilterer) ParseOwnershipTransferred(log types.Log) (*StableTokenOwnershipTransferred, error) {
	event := new(StableTokenOwnershipTransferred)
	if err := _StableToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableTokenRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the StableToken contract.
type StableTokenRegistrySetIterator struct {
	Event *StableTokenRegistrySet // Event containing the contract specifics and raw log

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
func (it *StableTokenRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenRegistrySet)
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
		it.Event = new(StableTokenRegistrySet)
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
func (it *StableTokenRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenRegistrySet represents a RegistrySet event raised by the StableToken contract.
type StableTokenRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_StableToken *StableTokenFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*StableTokenRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &StableTokenRegistrySetIterator{contract: _StableToken.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_StableToken *StableTokenFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *StableTokenRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenRegistrySet)
				if err := _StableToken.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_StableToken *StableTokenFilterer) ParseRegistrySet(log types.Log) (*StableTokenRegistrySet, error) {
	event := new(StableTokenRegistrySet)
	if err := _StableToken.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StableToken contract.
type StableTokenTransferIterator struct {
	Event *StableTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StableTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenTransfer)
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
		it.Event = new(StableTokenTransfer)
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
func (it *StableTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenTransfer represents a Transfer event raised by the StableToken contract.
type StableTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableToken *StableTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StableTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StableTokenTransferIterator{contract: _StableToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableToken *StableTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StableTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenTransfer)
				if err := _StableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StableToken *StableTokenFilterer) ParseTransfer(log types.Log) (*StableTokenTransfer, error) {
	event := new(StableTokenTransfer)
	if err := _StableToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StableTokenTransferCommentIterator is returned from FilterTransferComment and is used to iterate over the raw logs and unpacked data for TransferComment events raised by the StableToken contract.
type StableTokenTransferCommentIterator struct {
	Event *StableTokenTransferComment // Event containing the contract specifics and raw log

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
func (it *StableTokenTransferCommentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StableTokenTransferComment)
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
		it.Event = new(StableTokenTransferComment)
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
func (it *StableTokenTransferCommentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StableTokenTransferCommentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StableTokenTransferComment represents a TransferComment event raised by the StableToken contract.
type StableTokenTransferComment struct {
	Comment string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferComment is a free log retrieval operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_StableToken *StableTokenFilterer) FilterTransferComment(opts *bind.FilterOpts) (*StableTokenTransferCommentIterator, error) {

	logs, sub, err := _StableToken.contract.FilterLogs(opts, "TransferComment")
	if err != nil {
		return nil, err
	}
	return &StableTokenTransferCommentIterator{contract: _StableToken.contract, event: "TransferComment", logs: logs, sub: sub}, nil
}

// WatchTransferComment is a free log subscription operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_StableToken *StableTokenFilterer) WatchTransferComment(opts *bind.WatchOpts, sink chan<- *StableTokenTransferComment) (event.Subscription, error) {

	logs, sub, err := _StableToken.contract.WatchLogs(opts, "TransferComment")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StableTokenTransferComment)
				if err := _StableToken.contract.UnpackLog(event, "TransferComment", log); err != nil {
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

// ParseTransferComment is a log parse operation binding the contract event 0xe5d4e30fb8364e57bc4d662a07d0cf36f4c34552004c4c3624620a2c1d1c03dc.
//
// Solidity: event TransferComment(string comment)
func (_StableToken *StableTokenFilterer) ParseTransferComment(log types.Log) (*StableTokenTransferComment, error) {
	event := new(StableTokenTransferComment)
	if err := _StableToken.contract.UnpackLog(event, "TransferComment", log); err != nil {
		return nil, err
	}
	return event, nil
}
