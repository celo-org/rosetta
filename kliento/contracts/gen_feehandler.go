// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/celo-org/celo-blockchain"
	"github.com/celo-org/celo-blockchain/accounts/abi"
	"github.com/celo-org/celo-blockchain/accounts/abi/bind"
	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FeeHandlerMetaData contains all meta data concerning the FeeHandler contract.
var FeeHandlerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"BurnFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burning\",\"type\":\"uint256\"}],\"name\":\"DailyLimitHit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"DailyLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DailySellLimitUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"FeeBeneficiarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxSlippage\",\"type\":\"uint256\"}],\"name\":\"MaxSlippageSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"}],\"name\":\"RouterUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SoldAndBurnedToken\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"FIXED1_UINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BURN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burnFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"feeBeneficiary\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastLimitDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_registryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newFeeBeneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newBurnFraction\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"handlers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newLimits\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newMaxSlippages\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenHandler\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenMaxSlippage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenDailySellLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenCurrentDaySellLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getTokenToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"setFeeBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"setBurnFraction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"sell\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"activateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"deactivateToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"}],\"name\":\"setHandler\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"distribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newMax\",\"type\":\"uint256\"}],\"name\":\"setMaxSplippage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newLimit\",\"type\":\"uint256\"}],\"name\":\"setDailySellLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"burnCelo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"distributeAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"handleAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"handle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPastBurnForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountToBurn\",\"type\":\"uint256\"}],\"name\":\"dailySellLimitHit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FeeHandlerABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeHandlerMetaData.ABI instead.
var FeeHandlerABI = FeeHandlerMetaData.ABI

// FeeHandler is an auto generated Go binding around an Ethereum contract.
type FeeHandler struct {
	FeeHandlerCaller     // Read-only binding to the contract
	FeeHandlerTransactor // Write-only binding to the contract
	FeeHandlerFilterer   // Log filterer for contract events
}

// FeeHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FeeHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeHandlerSession struct {
	Contract     *FeeHandler       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeHandlerCallerSession struct {
	Contract *FeeHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// FeeHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeHandlerTransactorSession struct {
	Contract     *FeeHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// FeeHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FeeHandlerRaw struct {
	Contract *FeeHandler // Generic contract binding to access the raw methods on
}

// FeeHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeHandlerCallerRaw struct {
	Contract *FeeHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// FeeHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeHandlerTransactorRaw struct {
	Contract *FeeHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeHandler creates a new instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandler(address common.Address, backend bind.ContractBackend) (*FeeHandler, error) {
	contract, err := bindFeeHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeHandler{FeeHandlerCaller: FeeHandlerCaller{contract: contract}, FeeHandlerTransactor: FeeHandlerTransactor{contract: contract}, FeeHandlerFilterer: FeeHandlerFilterer{contract: contract}}, nil
}

// NewFeeHandlerCaller creates a new read-only instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerCaller(address common.Address, caller bind.ContractCaller) (*FeeHandlerCaller, error) {
	contract, err := bindFeeHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerCaller{contract: contract}, nil
}

// NewFeeHandlerTransactor creates a new write-only instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*FeeHandlerTransactor, error) {
	contract, err := bindFeeHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerTransactor{contract: contract}, nil
}

// NewFeeHandlerFilterer creates a new log filterer instance of FeeHandler, bound to a specific deployed contract.
func NewFeeHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*FeeHandlerFilterer, error) {
	contract, err := bindFeeHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerFilterer{contract: contract}, nil
}

// bindFeeHandler binds a generic wrapper to an already deployed contract.
func bindFeeHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseFeeHandlerABI parses the ABI
func ParseFeeHandlerABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(FeeHandlerABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandler *FeeHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandler.Contract.FeeHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandler *FeeHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.Contract.FeeHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandler *FeeHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandler.Contract.FeeHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeHandler *FeeHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeHandler *FeeHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeHandler *FeeHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeHandler.Contract.contract.Transact(opts, method, params...)
}

// FIXED1UINT is a free data retrieval call binding the contract method 0x036235a6.
//
// Solidity: function FIXED1_UINT() view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) FIXED1UINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "FIXED1_UINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FIXED1UINT is a free data retrieval call binding the contract method 0x036235a6.
//
// Solidity: function FIXED1_UINT() view returns(uint256)
func (_FeeHandler *FeeHandlerSession) FIXED1UINT() (*big.Int, error) {
	return _FeeHandler.Contract.FIXED1UINT(&_FeeHandler.CallOpts)
}

// FIXED1UINT is a free data retrieval call binding the contract method 0x036235a6.
//
// Solidity: function FIXED1_UINT() view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) FIXED1UINT() (*big.Int, error) {
	return _FeeHandler.Contract.FIXED1UINT(&_FeeHandler.CallOpts)
}

// MINBURN is a free data retrieval call binding the contract method 0x49844b1c.
//
// Solidity: function MIN_BURN() view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) MINBURN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "MIN_BURN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBURN is a free data retrieval call binding the contract method 0x49844b1c.
//
// Solidity: function MIN_BURN() view returns(uint256)
func (_FeeHandler *FeeHandlerSession) MINBURN() (*big.Int, error) {
	return _FeeHandler.Contract.MINBURN(&_FeeHandler.CallOpts)
}

// MINBURN is a free data retrieval call binding the contract method 0x49844b1c.
//
// Solidity: function MIN_BURN() view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) MINBURN() (*big.Int, error) {
	return _FeeHandler.Contract.MINBURN(&_FeeHandler.CallOpts)
}

// BurnFraction is a free data retrieval call binding the contract method 0x8de065b6.
//
// Solidity: function burnFraction() view returns(uint256 value)
func (_FeeHandler *FeeHandlerCaller) BurnFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "burnFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BurnFraction is a free data retrieval call binding the contract method 0x8de065b6.
//
// Solidity: function burnFraction() view returns(uint256 value)
func (_FeeHandler *FeeHandlerSession) BurnFraction() (*big.Int, error) {
	return _FeeHandler.Contract.BurnFraction(&_FeeHandler.CallOpts)
}

// BurnFraction is a free data retrieval call binding the contract method 0x8de065b6.
//
// Solidity: function burnFraction() view returns(uint256 value)
func (_FeeHandler *FeeHandlerCallerSession) BurnFraction() (*big.Int, error) {
	return _FeeHandler.Contract.BurnFraction(&_FeeHandler.CallOpts)
}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_FeeHandler *FeeHandlerCaller) FeeBeneficiary(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "feeBeneficiary")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_FeeHandler *FeeHandlerSession) FeeBeneficiary() (common.Address, error) {
	return _FeeHandler.Contract.FeeBeneficiary(&_FeeHandler.CallOpts)
}

// FeeBeneficiary is a free data retrieval call binding the contract method 0x492fb343.
//
// Solidity: function feeBeneficiary() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) FeeBeneficiary() (common.Address, error) {
	return _FeeHandler.Contract.FeeBeneficiary(&_FeeHandler.CallOpts)
}

// GetActiveTokens is a free data retrieval call binding the contract method 0x5f5817e3.
//
// Solidity: function getActiveTokens() view returns(address[])
func (_FeeHandler *FeeHandlerCaller) GetActiveTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getActiveTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveTokens is a free data retrieval call binding the contract method 0x5f5817e3.
//
// Solidity: function getActiveTokens() view returns(address[])
func (_FeeHandler *FeeHandlerSession) GetActiveTokens() ([]common.Address, error) {
	return _FeeHandler.Contract.GetActiveTokens(&_FeeHandler.CallOpts)
}

// GetActiveTokens is a free data retrieval call binding the contract method 0x5f5817e3.
//
// Solidity: function getActiveTokens() view returns(address[])
func (_FeeHandler *FeeHandlerCallerSession) GetActiveTokens() ([]common.Address, error) {
	return _FeeHandler.Contract.GetActiveTokens(&_FeeHandler.CallOpts)
}

// GetPastBurnForToken is a free data retrieval call binding the contract method 0xec4bd8ae.
//
// Solidity: function getPastBurnForToken(address token) view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) GetPastBurnForToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getPastBurnForToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPastBurnForToken is a free data retrieval call binding the contract method 0xec4bd8ae.
//
// Solidity: function getPastBurnForToken(address token) view returns(uint256)
func (_FeeHandler *FeeHandlerSession) GetPastBurnForToken(token common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetPastBurnForToken(&_FeeHandler.CallOpts, token)
}

// GetPastBurnForToken is a free data retrieval call binding the contract method 0xec4bd8ae.
//
// Solidity: function getPastBurnForToken(address token) view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) GetPastBurnForToken(token common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetPastBurnForToken(&_FeeHandler.CallOpts, token)
}

// GetTokenActive is a free data retrieval call binding the contract method 0x6c6c65ad.
//
// Solidity: function getTokenActive(address tokenAddress) view returns(bool)
func (_FeeHandler *FeeHandlerCaller) GetTokenActive(opts *bind.CallOpts, tokenAddress common.Address) (bool, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getTokenActive", tokenAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetTokenActive is a free data retrieval call binding the contract method 0x6c6c65ad.
//
// Solidity: function getTokenActive(address tokenAddress) view returns(bool)
func (_FeeHandler *FeeHandlerSession) GetTokenActive(tokenAddress common.Address) (bool, error) {
	return _FeeHandler.Contract.GetTokenActive(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenActive is a free data retrieval call binding the contract method 0x6c6c65ad.
//
// Solidity: function getTokenActive(address tokenAddress) view returns(bool)
func (_FeeHandler *FeeHandlerCallerSession) GetTokenActive(tokenAddress common.Address) (bool, error) {
	return _FeeHandler.Contract.GetTokenActive(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenCurrentDaySellLimit is a free data retrieval call binding the contract method 0x13e33cea.
//
// Solidity: function getTokenCurrentDaySellLimit(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) GetTokenCurrentDaySellLimit(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getTokenCurrentDaySellLimit", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenCurrentDaySellLimit is a free data retrieval call binding the contract method 0x13e33cea.
//
// Solidity: function getTokenCurrentDaySellLimit(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerSession) GetTokenCurrentDaySellLimit(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenCurrentDaySellLimit(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenCurrentDaySellLimit is a free data retrieval call binding the contract method 0x13e33cea.
//
// Solidity: function getTokenCurrentDaySellLimit(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) GetTokenCurrentDaySellLimit(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenCurrentDaySellLimit(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenDailySellLimit is a free data retrieval call binding the contract method 0xce4773d3.
//
// Solidity: function getTokenDailySellLimit(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) GetTokenDailySellLimit(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getTokenDailySellLimit", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenDailySellLimit is a free data retrieval call binding the contract method 0xce4773d3.
//
// Solidity: function getTokenDailySellLimit(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerSession) GetTokenDailySellLimit(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenDailySellLimit(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenDailySellLimit is a free data retrieval call binding the contract method 0xce4773d3.
//
// Solidity: function getTokenDailySellLimit(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) GetTokenDailySellLimit(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenDailySellLimit(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenHandler is a free data retrieval call binding the contract method 0x3b9e3ad6.
//
// Solidity: function getTokenHandler(address tokenAddress) view returns(address)
func (_FeeHandler *FeeHandlerCaller) GetTokenHandler(opts *bind.CallOpts, tokenAddress common.Address) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getTokenHandler", tokenAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenHandler is a free data retrieval call binding the contract method 0x3b9e3ad6.
//
// Solidity: function getTokenHandler(address tokenAddress) view returns(address)
func (_FeeHandler *FeeHandlerSession) GetTokenHandler(tokenAddress common.Address) (common.Address, error) {
	return _FeeHandler.Contract.GetTokenHandler(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenHandler is a free data retrieval call binding the contract method 0x3b9e3ad6.
//
// Solidity: function getTokenHandler(address tokenAddress) view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) GetTokenHandler(tokenAddress common.Address) (common.Address, error) {
	return _FeeHandler.Contract.GetTokenHandler(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenMaxSlippage is a free data retrieval call binding the contract method 0x92f8bce3.
//
// Solidity: function getTokenMaxSlippage(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) GetTokenMaxSlippage(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getTokenMaxSlippage", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenMaxSlippage is a free data retrieval call binding the contract method 0x92f8bce3.
//
// Solidity: function getTokenMaxSlippage(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerSession) GetTokenMaxSlippage(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenMaxSlippage(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenMaxSlippage is a free data retrieval call binding the contract method 0x92f8bce3.
//
// Solidity: function getTokenMaxSlippage(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) GetTokenMaxSlippage(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenMaxSlippage(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenToDistribute is a free data retrieval call binding the contract method 0x6654f435.
//
// Solidity: function getTokenToDistribute(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) GetTokenToDistribute(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getTokenToDistribute", tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenToDistribute is a free data retrieval call binding the contract method 0x6654f435.
//
// Solidity: function getTokenToDistribute(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerSession) GetTokenToDistribute(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenToDistribute(&_FeeHandler.CallOpts, tokenAddress)
}

// GetTokenToDistribute is a free data retrieval call binding the contract method 0x6654f435.
//
// Solidity: function getTokenToDistribute(address tokenAddress) view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) GetTokenToDistribute(tokenAddress common.Address) (*big.Int, error) {
	return _FeeHandler.Contract.GetTokenToDistribute(&_FeeHandler.CallOpts, tokenAddress)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_FeeHandler *FeeHandlerCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "getVersionNumber")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_FeeHandler *FeeHandlerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FeeHandler.Contract.GetVersionNumber(&_FeeHandler.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_FeeHandler *FeeHandlerCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _FeeHandler.Contract.GetVersionNumber(&_FeeHandler.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeHandler *FeeHandlerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeHandler *FeeHandlerSession) Initialized() (bool, error) {
	return _FeeHandler.Contract.Initialized(&_FeeHandler.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_FeeHandler *FeeHandlerCallerSession) Initialized() (bool, error) {
	return _FeeHandler.Contract.Initialized(&_FeeHandler.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FeeHandler *FeeHandlerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FeeHandler *FeeHandlerSession) IsOwner() (bool, error) {
	return _FeeHandler.Contract.IsOwner(&_FeeHandler.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_FeeHandler *FeeHandlerCallerSession) IsOwner() (bool, error) {
	return _FeeHandler.Contract.IsOwner(&_FeeHandler.CallOpts)
}

// LastLimitDay is a free data retrieval call binding the contract method 0xdf3712e4.
//
// Solidity: function lastLimitDay() view returns(uint256)
func (_FeeHandler *FeeHandlerCaller) LastLimitDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "lastLimitDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastLimitDay is a free data retrieval call binding the contract method 0xdf3712e4.
//
// Solidity: function lastLimitDay() view returns(uint256)
func (_FeeHandler *FeeHandlerSession) LastLimitDay() (*big.Int, error) {
	return _FeeHandler.Contract.LastLimitDay(&_FeeHandler.CallOpts)
}

// LastLimitDay is a free data retrieval call binding the contract method 0xdf3712e4.
//
// Solidity: function lastLimitDay() view returns(uint256)
func (_FeeHandler *FeeHandlerCallerSession) LastLimitDay() (*big.Int, error) {
	return _FeeHandler.Contract.LastLimitDay(&_FeeHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeHandler *FeeHandlerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeHandler *FeeHandlerSession) Owner() (common.Address, error) {
	return _FeeHandler.Contract.Owner(&_FeeHandler.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) Owner() (common.Address, error) {
	return _FeeHandler.Contract.Owner(&_FeeHandler.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_FeeHandler *FeeHandlerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeHandler.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_FeeHandler *FeeHandlerSession) Registry() (common.Address, error) {
	return _FeeHandler.Contract.Registry(&_FeeHandler.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_FeeHandler *FeeHandlerCallerSession) Registry() (common.Address, error) {
	return _FeeHandler.Contract.Registry(&_FeeHandler.CallOpts)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) ActivateToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "activateToken", tokenAddress)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerSession) ActivateToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.ActivateToken(&_FeeHandler.TransactOpts, tokenAddress)
}

// ActivateToken is a paid mutator transaction binding the contract method 0x0d1ce2d2.
//
// Solidity: function activateToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) ActivateToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.ActivateToken(&_FeeHandler.TransactOpts, tokenAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address tokenAddress, address handlerAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) AddToken(opts *bind.TransactOpts, tokenAddress common.Address, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "addToken", tokenAddress, handlerAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address tokenAddress, address handlerAddress) returns()
func (_FeeHandler *FeeHandlerSession) AddToken(tokenAddress common.Address, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.AddToken(&_FeeHandler.TransactOpts, tokenAddress, handlerAddress)
}

// AddToken is a paid mutator transaction binding the contract method 0x5476bd72.
//
// Solidity: function addToken(address tokenAddress, address handlerAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) AddToken(tokenAddress common.Address, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.AddToken(&_FeeHandler.TransactOpts, tokenAddress, handlerAddress)
}

// BurnCelo is a paid mutator transaction binding the contract method 0xb8b99e4d.
//
// Solidity: function burnCelo() returns()
func (_FeeHandler *FeeHandlerTransactor) BurnCelo(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "burnCelo")
}

// BurnCelo is a paid mutator transaction binding the contract method 0xb8b99e4d.
//
// Solidity: function burnCelo() returns()
func (_FeeHandler *FeeHandlerSession) BurnCelo() (*types.Transaction, error) {
	return _FeeHandler.Contract.BurnCelo(&_FeeHandler.TransactOpts)
}

// BurnCelo is a paid mutator transaction binding the contract method 0xb8b99e4d.
//
// Solidity: function burnCelo() returns()
func (_FeeHandler *FeeHandlerTransactorSession) BurnCelo() (*types.Transaction, error) {
	return _FeeHandler.Contract.BurnCelo(&_FeeHandler.TransactOpts)
}

// DailySellLimitHit is a paid mutator transaction binding the contract method 0x31828a5b.
//
// Solidity: function dailySellLimitHit(address token, uint256 amountToBurn) returns(bool)
func (_FeeHandler *FeeHandlerTransactor) DailySellLimitHit(opts *bind.TransactOpts, token common.Address, amountToBurn *big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "dailySellLimitHit", token, amountToBurn)
}

// DailySellLimitHit is a paid mutator transaction binding the contract method 0x31828a5b.
//
// Solidity: function dailySellLimitHit(address token, uint256 amountToBurn) returns(bool)
func (_FeeHandler *FeeHandlerSession) DailySellLimitHit(token common.Address, amountToBurn *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.DailySellLimitHit(&_FeeHandler.TransactOpts, token, amountToBurn)
}

// DailySellLimitHit is a paid mutator transaction binding the contract method 0x31828a5b.
//
// Solidity: function dailySellLimitHit(address token, uint256 amountToBurn) returns(bool)
func (_FeeHandler *FeeHandlerTransactorSession) DailySellLimitHit(token common.Address, amountToBurn *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.DailySellLimitHit(&_FeeHandler.TransactOpts, token, amountToBurn)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) DeactivateToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "deactivateToken", tokenAddress)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerSession) DeactivateToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.DeactivateToken(&_FeeHandler.TransactOpts, tokenAddress)
}

// DeactivateToken is a paid mutator transaction binding the contract method 0x68173bcf.
//
// Solidity: function deactivateToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) DeactivateToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.DeactivateToken(&_FeeHandler.TransactOpts, tokenAddress)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) Distribute(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "distribute", tokenAddress)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerSession) Distribute(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.Distribute(&_FeeHandler.TransactOpts, tokenAddress)
}

// Distribute is a paid mutator transaction binding the contract method 0x63453ae1.
//
// Solidity: function distribute(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) Distribute(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.Distribute(&_FeeHandler.TransactOpts, tokenAddress)
}

// DistributeAll is a paid mutator transaction binding the contract method 0x436596c4.
//
// Solidity: function distributeAll() returns()
func (_FeeHandler *FeeHandlerTransactor) DistributeAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "distributeAll")
}

// DistributeAll is a paid mutator transaction binding the contract method 0x436596c4.
//
// Solidity: function distributeAll() returns()
func (_FeeHandler *FeeHandlerSession) DistributeAll() (*types.Transaction, error) {
	return _FeeHandler.Contract.DistributeAll(&_FeeHandler.TransactOpts)
}

// DistributeAll is a paid mutator transaction binding the contract method 0x436596c4.
//
// Solidity: function distributeAll() returns()
func (_FeeHandler *FeeHandlerTransactorSession) DistributeAll() (*types.Transaction, error) {
	return _FeeHandler.Contract.DistributeAll(&_FeeHandler.TransactOpts)
}

// Handle is a paid mutator transaction binding the contract method 0x7b763140.
//
// Solidity: function handle(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) Handle(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "handle", tokenAddress)
}

// Handle is a paid mutator transaction binding the contract method 0x7b763140.
//
// Solidity: function handle(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerSession) Handle(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.Handle(&_FeeHandler.TransactOpts, tokenAddress)
}

// Handle is a paid mutator transaction binding the contract method 0x7b763140.
//
// Solidity: function handle(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) Handle(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.Handle(&_FeeHandler.TransactOpts, tokenAddress)
}

// HandleAll is a paid mutator transaction binding the contract method 0xee8df72d.
//
// Solidity: function handleAll() returns()
func (_FeeHandler *FeeHandlerTransactor) HandleAll(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "handleAll")
}

// HandleAll is a paid mutator transaction binding the contract method 0xee8df72d.
//
// Solidity: function handleAll() returns()
func (_FeeHandler *FeeHandlerSession) HandleAll() (*types.Transaction, error) {
	return _FeeHandler.Contract.HandleAll(&_FeeHandler.TransactOpts)
}

// HandleAll is a paid mutator transaction binding the contract method 0xee8df72d.
//
// Solidity: function handleAll() returns()
func (_FeeHandler *FeeHandlerTransactorSession) HandleAll() (*types.Transaction, error) {
	return _FeeHandler.Contract.HandleAll(&_FeeHandler.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x650a1605.
//
// Solidity: function initialize(address _registryAddress, address newFeeBeneficiary, uint256 newBurnFraction, address[] tokens, address[] handlers, uint256[] newLimits, uint256[] newMaxSlippages) returns()
func (_FeeHandler *FeeHandlerTransactor) Initialize(opts *bind.TransactOpts, _registryAddress common.Address, newFeeBeneficiary common.Address, newBurnFraction *big.Int, tokens []common.Address, handlers []common.Address, newLimits []*big.Int, newMaxSlippages []*big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "initialize", _registryAddress, newFeeBeneficiary, newBurnFraction, tokens, handlers, newLimits, newMaxSlippages)
}

// Initialize is a paid mutator transaction binding the contract method 0x650a1605.
//
// Solidity: function initialize(address _registryAddress, address newFeeBeneficiary, uint256 newBurnFraction, address[] tokens, address[] handlers, uint256[] newLimits, uint256[] newMaxSlippages) returns()
func (_FeeHandler *FeeHandlerSession) Initialize(_registryAddress common.Address, newFeeBeneficiary common.Address, newBurnFraction *big.Int, tokens []common.Address, handlers []common.Address, newLimits []*big.Int, newMaxSlippages []*big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.Initialize(&_FeeHandler.TransactOpts, _registryAddress, newFeeBeneficiary, newBurnFraction, tokens, handlers, newLimits, newMaxSlippages)
}

// Initialize is a paid mutator transaction binding the contract method 0x650a1605.
//
// Solidity: function initialize(address _registryAddress, address newFeeBeneficiary, uint256 newBurnFraction, address[] tokens, address[] handlers, uint256[] newLimits, uint256[] newMaxSlippages) returns()
func (_FeeHandler *FeeHandlerTransactorSession) Initialize(_registryAddress common.Address, newFeeBeneficiary common.Address, newBurnFraction *big.Int, tokens []common.Address, handlers []common.Address, newLimits []*big.Int, newMaxSlippages []*big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.Initialize(&_FeeHandler.TransactOpts, _registryAddress, newFeeBeneficiary, newBurnFraction, tokens, handlers, newLimits, newMaxSlippages)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) RemoveToken(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "removeToken", tokenAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerSession) RemoveToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RemoveToken(&_FeeHandler.TransactOpts, tokenAddress)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) RemoveToken(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.RemoveToken(&_FeeHandler.TransactOpts, tokenAddress)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeHandler *FeeHandlerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeHandler *FeeHandlerSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeHandler.Contract.RenounceOwnership(&_FeeHandler.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeHandler *FeeHandlerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeHandler.Contract.RenounceOwnership(&_FeeHandler.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x94b6f9d4.
//
// Solidity: function sell(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) Sell(opts *bind.TransactOpts, tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "sell", tokenAddress)
}

// Sell is a paid mutator transaction binding the contract method 0x94b6f9d4.
//
// Solidity: function sell(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerSession) Sell(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.Sell(&_FeeHandler.TransactOpts, tokenAddress)
}

// Sell is a paid mutator transaction binding the contract method 0x94b6f9d4.
//
// Solidity: function sell(address tokenAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) Sell(tokenAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.Sell(&_FeeHandler.TransactOpts, tokenAddress)
}

// SetBurnFraction is a paid mutator transaction binding the contract method 0x384995cd.
//
// Solidity: function setBurnFraction(uint256 fraction) returns()
func (_FeeHandler *FeeHandlerTransactor) SetBurnFraction(opts *bind.TransactOpts, fraction *big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setBurnFraction", fraction)
}

// SetBurnFraction is a paid mutator transaction binding the contract method 0x384995cd.
//
// Solidity: function setBurnFraction(uint256 fraction) returns()
func (_FeeHandler *FeeHandlerSession) SetBurnFraction(fraction *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetBurnFraction(&_FeeHandler.TransactOpts, fraction)
}

// SetBurnFraction is a paid mutator transaction binding the contract method 0x384995cd.
//
// Solidity: function setBurnFraction(uint256 fraction) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetBurnFraction(fraction *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetBurnFraction(&_FeeHandler.TransactOpts, fraction)
}

// SetDailySellLimit is a paid mutator transaction binding the contract method 0xc558df38.
//
// Solidity: function setDailySellLimit(address token, uint256 newLimit) returns()
func (_FeeHandler *FeeHandlerTransactor) SetDailySellLimit(opts *bind.TransactOpts, token common.Address, newLimit *big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setDailySellLimit", token, newLimit)
}

// SetDailySellLimit is a paid mutator transaction binding the contract method 0xc558df38.
//
// Solidity: function setDailySellLimit(address token, uint256 newLimit) returns()
func (_FeeHandler *FeeHandlerSession) SetDailySellLimit(token common.Address, newLimit *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetDailySellLimit(&_FeeHandler.TransactOpts, token, newLimit)
}

// SetDailySellLimit is a paid mutator transaction binding the contract method 0xc558df38.
//
// Solidity: function setDailySellLimit(address token, uint256 newLimit) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetDailySellLimit(token common.Address, newLimit *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetDailySellLimit(&_FeeHandler.TransactOpts, token, newLimit)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address beneficiary) returns()
func (_FeeHandler *FeeHandlerTransactor) SetFeeBeneficiary(opts *bind.TransactOpts, beneficiary common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setFeeBeneficiary", beneficiary)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address beneficiary) returns()
func (_FeeHandler *FeeHandlerSession) SetFeeBeneficiary(beneficiary common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetFeeBeneficiary(&_FeeHandler.TransactOpts, beneficiary)
}

// SetFeeBeneficiary is a paid mutator transaction binding the contract method 0x5a0a3d82.
//
// Solidity: function setFeeBeneficiary(address beneficiary) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetFeeBeneficiary(beneficiary common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetFeeBeneficiary(&_FeeHandler.TransactOpts, beneficiary)
}

// SetHandler is a paid mutator transaction binding the contract method 0xe9f1bbde.
//
// Solidity: function setHandler(address tokenAddress, address handlerAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) SetHandler(opts *bind.TransactOpts, tokenAddress common.Address, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setHandler", tokenAddress, handlerAddress)
}

// SetHandler is a paid mutator transaction binding the contract method 0xe9f1bbde.
//
// Solidity: function setHandler(address tokenAddress, address handlerAddress) returns()
func (_FeeHandler *FeeHandlerSession) SetHandler(tokenAddress common.Address, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetHandler(&_FeeHandler.TransactOpts, tokenAddress, handlerAddress)
}

// SetHandler is a paid mutator transaction binding the contract method 0xe9f1bbde.
//
// Solidity: function setHandler(address tokenAddress, address handlerAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetHandler(tokenAddress common.Address, handlerAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetHandler(&_FeeHandler.TransactOpts, tokenAddress, handlerAddress)
}

// SetMaxSplippage is a paid mutator transaction binding the contract method 0x4e73db99.
//
// Solidity: function setMaxSplippage(address token, uint256 newMax) returns()
func (_FeeHandler *FeeHandlerTransactor) SetMaxSplippage(opts *bind.TransactOpts, token common.Address, newMax *big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setMaxSplippage", token, newMax)
}

// SetMaxSplippage is a paid mutator transaction binding the contract method 0x4e73db99.
//
// Solidity: function setMaxSplippage(address token, uint256 newMax) returns()
func (_FeeHandler *FeeHandlerSession) SetMaxSplippage(token common.Address, newMax *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetMaxSplippage(&_FeeHandler.TransactOpts, token, newMax)
}

// SetMaxSplippage is a paid mutator transaction binding the contract method 0x4e73db99.
//
// Solidity: function setMaxSplippage(address token, uint256 newMax) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetMaxSplippage(token common.Address, newMax *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetMaxSplippage(&_FeeHandler.TransactOpts, token, newMax)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_FeeHandler *FeeHandlerTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_FeeHandler *FeeHandlerSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetRegistry(&_FeeHandler.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_FeeHandler *FeeHandlerTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.SetRegistry(&_FeeHandler.TransactOpts, registryAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address recipient, uint256 value) returns(bool)
func (_FeeHandler *FeeHandlerTransactor) Transfer(opts *bind.TransactOpts, token common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "transfer", token, recipient, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address recipient, uint256 value) returns(bool)
func (_FeeHandler *FeeHandlerSession) Transfer(token common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.Transfer(&_FeeHandler.TransactOpts, token, recipient, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address token, address recipient, uint256 value) returns(bool)
func (_FeeHandler *FeeHandlerTransactorSession) Transfer(token common.Address, recipient common.Address, value *big.Int) (*types.Transaction, error) {
	return _FeeHandler.Contract.Transfer(&_FeeHandler.TransactOpts, token, recipient, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeHandler *FeeHandlerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeHandler.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeHandler *FeeHandlerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.TransferOwnership(&_FeeHandler.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeHandler *FeeHandlerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeHandler.Contract.TransferOwnership(&_FeeHandler.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_FeeHandler *FeeHandlerFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _FeeHandler.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "BurnFractionSet":
		event, err = _FeeHandler.ParseBurnFractionSet(log)
	case "DailyLimitHit":
		event, err = _FeeHandler.ParseDailyLimitHit(log)
	case "DailyLimitSet":
		event, err = _FeeHandler.ParseDailyLimitSet(log)
	case "DailySellLimitUpdated":
		event, err = _FeeHandler.ParseDailySellLimitUpdated(log)
	case "FeeBeneficiarySet":
		event, err = _FeeHandler.ParseFeeBeneficiarySet(log)
	case "MaxSlippageSet":
		event, err = _FeeHandler.ParseMaxSlippageSet(log)
	case "OwnershipTransferred":
		event, err = _FeeHandler.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _FeeHandler.ParseRegistrySet(log)
	case "RouterAddressRemoved":
		event, err = _FeeHandler.ParseRouterAddressRemoved(log)
	case "RouterAddressSet":
		event, err = _FeeHandler.ParseRouterAddressSet(log)
	case "RouterUsed":
		event, err = _FeeHandler.ParseRouterUsed(log)
	case "SoldAndBurnedToken":
		event, err = _FeeHandler.ParseSoldAndBurnedToken(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeHandler *FeeHandlerTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _FeeHandler.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeHandler *FeeHandlerSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FeeHandler.Contract.Fallback(&_FeeHandler.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_FeeHandler *FeeHandlerTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _FeeHandler.Contract.Fallback(&_FeeHandler.TransactOpts, calldata)
}

// FeeHandlerBurnFractionSetIterator is returned from FilterBurnFractionSet and is used to iterate over the raw logs and unpacked data for BurnFractionSet events raised by the FeeHandler contract.
type FeeHandlerBurnFractionSetIterator struct {
	Event *FeeHandlerBurnFractionSet // Event containing the contract specifics and raw log

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
func (it *FeeHandlerBurnFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerBurnFractionSet)
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
		it.Event = new(FeeHandlerBurnFractionSet)
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
func (it *FeeHandlerBurnFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerBurnFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerBurnFractionSet represents a BurnFractionSet event raised by the FeeHandler contract.
type FeeHandlerBurnFractionSet struct {
	Fraction *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBurnFractionSet is a free log retrieval operation binding the contract event 0x41c679f4bcdc2c95f79a3647e2237162d9763d86685ef6c667781230c8ba9157.
//
// Solidity: event BurnFractionSet(uint256 fraction)
func (_FeeHandler *FeeHandlerFilterer) FilterBurnFractionSet(opts *bind.FilterOpts) (*FeeHandlerBurnFractionSetIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "BurnFractionSet")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerBurnFractionSetIterator{contract: _FeeHandler.contract, event: "BurnFractionSet", logs: logs, sub: sub}, nil
}

// WatchBurnFractionSet is a free log subscription operation binding the contract event 0x41c679f4bcdc2c95f79a3647e2237162d9763d86685ef6c667781230c8ba9157.
//
// Solidity: event BurnFractionSet(uint256 fraction)
func (_FeeHandler *FeeHandlerFilterer) WatchBurnFractionSet(opts *bind.WatchOpts, sink chan<- *FeeHandlerBurnFractionSet) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "BurnFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerBurnFractionSet)
				if err := _FeeHandler.contract.UnpackLog(event, "BurnFractionSet", log); err != nil {
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

// ParseBurnFractionSet is a log parse operation binding the contract event 0x41c679f4bcdc2c95f79a3647e2237162d9763d86685ef6c667781230c8ba9157.
//
// Solidity: event BurnFractionSet(uint256 fraction)
func (_FeeHandler *FeeHandlerFilterer) ParseBurnFractionSet(log types.Log) (*FeeHandlerBurnFractionSet, error) {
	event := new(FeeHandlerBurnFractionSet)
	if err := _FeeHandler.contract.UnpackLog(event, "BurnFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerDailyLimitHitIterator is returned from FilterDailyLimitHit and is used to iterate over the raw logs and unpacked data for DailyLimitHit events raised by the FeeHandler contract.
type FeeHandlerDailyLimitHitIterator struct {
	Event *FeeHandlerDailyLimitHit // Event containing the contract specifics and raw log

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
func (it *FeeHandlerDailyLimitHitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerDailyLimitHit)
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
		it.Event = new(FeeHandlerDailyLimitHit)
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
func (it *FeeHandlerDailyLimitHitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerDailyLimitHitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerDailyLimitHit represents a DailyLimitHit event raised by the FeeHandler contract.
type FeeHandlerDailyLimitHit struct {
	Token   common.Address
	Burning *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitHit is a free log retrieval operation binding the contract event 0xb1a68b0b66260ca392f760fd4dda4a94818d69c89a4eeb6610eb41db7bab8c37.
//
// Solidity: event DailyLimitHit(address token, uint256 burning)
func (_FeeHandler *FeeHandlerFilterer) FilterDailyLimitHit(opts *bind.FilterOpts) (*FeeHandlerDailyLimitHitIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "DailyLimitHit")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerDailyLimitHitIterator{contract: _FeeHandler.contract, event: "DailyLimitHit", logs: logs, sub: sub}, nil
}

// WatchDailyLimitHit is a free log subscription operation binding the contract event 0xb1a68b0b66260ca392f760fd4dda4a94818d69c89a4eeb6610eb41db7bab8c37.
//
// Solidity: event DailyLimitHit(address token, uint256 burning)
func (_FeeHandler *FeeHandlerFilterer) WatchDailyLimitHit(opts *bind.WatchOpts, sink chan<- *FeeHandlerDailyLimitHit) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "DailyLimitHit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerDailyLimitHit)
				if err := _FeeHandler.contract.UnpackLog(event, "DailyLimitHit", log); err != nil {
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

// ParseDailyLimitHit is a log parse operation binding the contract event 0xb1a68b0b66260ca392f760fd4dda4a94818d69c89a4eeb6610eb41db7bab8c37.
//
// Solidity: event DailyLimitHit(address token, uint256 burning)
func (_FeeHandler *FeeHandlerFilterer) ParseDailyLimitHit(log types.Log) (*FeeHandlerDailyLimitHit, error) {
	event := new(FeeHandlerDailyLimitHit)
	if err := _FeeHandler.contract.UnpackLog(event, "DailyLimitHit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerDailyLimitSetIterator is returned from FilterDailyLimitSet and is used to iterate over the raw logs and unpacked data for DailyLimitSet events raised by the FeeHandler contract.
type FeeHandlerDailyLimitSetIterator struct {
	Event *FeeHandlerDailyLimitSet // Event containing the contract specifics and raw log

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
func (it *FeeHandlerDailyLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerDailyLimitSet)
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
		it.Event = new(FeeHandlerDailyLimitSet)
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
func (it *FeeHandlerDailyLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerDailyLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerDailyLimitSet represents a DailyLimitSet event raised by the FeeHandler contract.
type FeeHandlerDailyLimitSet struct {
	TokenAddress common.Address
	NewLimit     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDailyLimitSet is a free log retrieval operation binding the contract event 0xd3d22ffd28b02735cf411bd7f925bd8da01212c7028153e0d632e2953ac3088e.
//
// Solidity: event DailyLimitSet(address tokenAddress, uint256 newLimit)
func (_FeeHandler *FeeHandlerFilterer) FilterDailyLimitSet(opts *bind.FilterOpts) (*FeeHandlerDailyLimitSetIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "DailyLimitSet")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerDailyLimitSetIterator{contract: _FeeHandler.contract, event: "DailyLimitSet", logs: logs, sub: sub}, nil
}

// WatchDailyLimitSet is a free log subscription operation binding the contract event 0xd3d22ffd28b02735cf411bd7f925bd8da01212c7028153e0d632e2953ac3088e.
//
// Solidity: event DailyLimitSet(address tokenAddress, uint256 newLimit)
func (_FeeHandler *FeeHandlerFilterer) WatchDailyLimitSet(opts *bind.WatchOpts, sink chan<- *FeeHandlerDailyLimitSet) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "DailyLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerDailyLimitSet)
				if err := _FeeHandler.contract.UnpackLog(event, "DailyLimitSet", log); err != nil {
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

// ParseDailyLimitSet is a log parse operation binding the contract event 0xd3d22ffd28b02735cf411bd7f925bd8da01212c7028153e0d632e2953ac3088e.
//
// Solidity: event DailyLimitSet(address tokenAddress, uint256 newLimit)
func (_FeeHandler *FeeHandlerFilterer) ParseDailyLimitSet(log types.Log) (*FeeHandlerDailyLimitSet, error) {
	event := new(FeeHandlerDailyLimitSet)
	if err := _FeeHandler.contract.UnpackLog(event, "DailyLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerDailySellLimitUpdatedIterator is returned from FilterDailySellLimitUpdated and is used to iterate over the raw logs and unpacked data for DailySellLimitUpdated events raised by the FeeHandler contract.
type FeeHandlerDailySellLimitUpdatedIterator struct {
	Event *FeeHandlerDailySellLimitUpdated // Event containing the contract specifics and raw log

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
func (it *FeeHandlerDailySellLimitUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerDailySellLimitUpdated)
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
		it.Event = new(FeeHandlerDailySellLimitUpdated)
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
func (it *FeeHandlerDailySellLimitUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerDailySellLimitUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerDailySellLimitUpdated represents a DailySellLimitUpdated event raised by the FeeHandler contract.
type FeeHandlerDailySellLimitUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDailySellLimitUpdated is a free log retrieval operation binding the contract event 0xcdcea7139bd245b1c7468bc1cfb59ad732b3b0909bafa9f9436ad74c81d0aafb.
//
// Solidity: event DailySellLimitUpdated(uint256 amount)
func (_FeeHandler *FeeHandlerFilterer) FilterDailySellLimitUpdated(opts *bind.FilterOpts) (*FeeHandlerDailySellLimitUpdatedIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "DailySellLimitUpdated")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerDailySellLimitUpdatedIterator{contract: _FeeHandler.contract, event: "DailySellLimitUpdated", logs: logs, sub: sub}, nil
}

// WatchDailySellLimitUpdated is a free log subscription operation binding the contract event 0xcdcea7139bd245b1c7468bc1cfb59ad732b3b0909bafa9f9436ad74c81d0aafb.
//
// Solidity: event DailySellLimitUpdated(uint256 amount)
func (_FeeHandler *FeeHandlerFilterer) WatchDailySellLimitUpdated(opts *bind.WatchOpts, sink chan<- *FeeHandlerDailySellLimitUpdated) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "DailySellLimitUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerDailySellLimitUpdated)
				if err := _FeeHandler.contract.UnpackLog(event, "DailySellLimitUpdated", log); err != nil {
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

// ParseDailySellLimitUpdated is a log parse operation binding the contract event 0xcdcea7139bd245b1c7468bc1cfb59ad732b3b0909bafa9f9436ad74c81d0aafb.
//
// Solidity: event DailySellLimitUpdated(uint256 amount)
func (_FeeHandler *FeeHandlerFilterer) ParseDailySellLimitUpdated(log types.Log) (*FeeHandlerDailySellLimitUpdated, error) {
	event := new(FeeHandlerDailySellLimitUpdated)
	if err := _FeeHandler.contract.UnpackLog(event, "DailySellLimitUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerFeeBeneficiarySetIterator is returned from FilterFeeBeneficiarySet and is used to iterate over the raw logs and unpacked data for FeeBeneficiarySet events raised by the FeeHandler contract.
type FeeHandlerFeeBeneficiarySetIterator struct {
	Event *FeeHandlerFeeBeneficiarySet // Event containing the contract specifics and raw log

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
func (it *FeeHandlerFeeBeneficiarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerFeeBeneficiarySet)
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
		it.Event = new(FeeHandlerFeeBeneficiarySet)
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
func (it *FeeHandlerFeeBeneficiarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerFeeBeneficiarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerFeeBeneficiarySet represents a FeeBeneficiarySet event raised by the FeeHandler contract.
type FeeHandlerFeeBeneficiarySet struct {
	NewBeneficiary common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterFeeBeneficiarySet is a free log retrieval operation binding the contract event 0xf7015098f8d6fa48f0560725ffd5f81bf687ee5ac45153b590bdcb04648bbdd3.
//
// Solidity: event FeeBeneficiarySet(address newBeneficiary)
func (_FeeHandler *FeeHandlerFilterer) FilterFeeBeneficiarySet(opts *bind.FilterOpts) (*FeeHandlerFeeBeneficiarySetIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "FeeBeneficiarySet")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerFeeBeneficiarySetIterator{contract: _FeeHandler.contract, event: "FeeBeneficiarySet", logs: logs, sub: sub}, nil
}

// WatchFeeBeneficiarySet is a free log subscription operation binding the contract event 0xf7015098f8d6fa48f0560725ffd5f81bf687ee5ac45153b590bdcb04648bbdd3.
//
// Solidity: event FeeBeneficiarySet(address newBeneficiary)
func (_FeeHandler *FeeHandlerFilterer) WatchFeeBeneficiarySet(opts *bind.WatchOpts, sink chan<- *FeeHandlerFeeBeneficiarySet) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "FeeBeneficiarySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerFeeBeneficiarySet)
				if err := _FeeHandler.contract.UnpackLog(event, "FeeBeneficiarySet", log); err != nil {
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

// ParseFeeBeneficiarySet is a log parse operation binding the contract event 0xf7015098f8d6fa48f0560725ffd5f81bf687ee5ac45153b590bdcb04648bbdd3.
//
// Solidity: event FeeBeneficiarySet(address newBeneficiary)
func (_FeeHandler *FeeHandlerFilterer) ParseFeeBeneficiarySet(log types.Log) (*FeeHandlerFeeBeneficiarySet, error) {
	event := new(FeeHandlerFeeBeneficiarySet)
	if err := _FeeHandler.contract.UnpackLog(event, "FeeBeneficiarySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerMaxSlippageSetIterator is returned from FilterMaxSlippageSet and is used to iterate over the raw logs and unpacked data for MaxSlippageSet events raised by the FeeHandler contract.
type FeeHandlerMaxSlippageSetIterator struct {
	Event *FeeHandlerMaxSlippageSet // Event containing the contract specifics and raw log

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
func (it *FeeHandlerMaxSlippageSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerMaxSlippageSet)
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
		it.Event = new(FeeHandlerMaxSlippageSet)
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
func (it *FeeHandlerMaxSlippageSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerMaxSlippageSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerMaxSlippageSet represents a MaxSlippageSet event raised by the FeeHandler contract.
type FeeHandlerMaxSlippageSet struct {
	Token       common.Address
	MaxSlippage *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMaxSlippageSet is a free log retrieval operation binding the contract event 0xd8df93d785f3d0d4294fd7b61e5d749c20eec95a2fed5b6b502a4cad09199ca6.
//
// Solidity: event MaxSlippageSet(address token, uint256 maxSlippage)
func (_FeeHandler *FeeHandlerFilterer) FilterMaxSlippageSet(opts *bind.FilterOpts) (*FeeHandlerMaxSlippageSetIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "MaxSlippageSet")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerMaxSlippageSetIterator{contract: _FeeHandler.contract, event: "MaxSlippageSet", logs: logs, sub: sub}, nil
}

// WatchMaxSlippageSet is a free log subscription operation binding the contract event 0xd8df93d785f3d0d4294fd7b61e5d749c20eec95a2fed5b6b502a4cad09199ca6.
//
// Solidity: event MaxSlippageSet(address token, uint256 maxSlippage)
func (_FeeHandler *FeeHandlerFilterer) WatchMaxSlippageSet(opts *bind.WatchOpts, sink chan<- *FeeHandlerMaxSlippageSet) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "MaxSlippageSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerMaxSlippageSet)
				if err := _FeeHandler.contract.UnpackLog(event, "MaxSlippageSet", log); err != nil {
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

// ParseMaxSlippageSet is a log parse operation binding the contract event 0xd8df93d785f3d0d4294fd7b61e5d749c20eec95a2fed5b6b502a4cad09199ca6.
//
// Solidity: event MaxSlippageSet(address token, uint256 maxSlippage)
func (_FeeHandler *FeeHandlerFilterer) ParseMaxSlippageSet(log types.Log) (*FeeHandlerMaxSlippageSet, error) {
	event := new(FeeHandlerMaxSlippageSet)
	if err := _FeeHandler.contract.UnpackLog(event, "MaxSlippageSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeHandler contract.
type FeeHandlerOwnershipTransferredIterator struct {
	Event *FeeHandlerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeHandlerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerOwnershipTransferred)
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
		it.Event = new(FeeHandlerOwnershipTransferred)
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
func (it *FeeHandlerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerOwnershipTransferred represents a OwnershipTransferred event raised by the FeeHandler contract.
type FeeHandlerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeHandler *FeeHandlerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeHandlerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerOwnershipTransferredIterator{contract: _FeeHandler.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeHandler *FeeHandlerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeHandlerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerOwnershipTransferred)
				if err := _FeeHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeeHandler *FeeHandlerFilterer) ParseOwnershipTransferred(log types.Log) (*FeeHandlerOwnershipTransferred, error) {
	event := new(FeeHandlerOwnershipTransferred)
	if err := _FeeHandler.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the FeeHandler contract.
type FeeHandlerRegistrySetIterator struct {
	Event *FeeHandlerRegistrySet // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRegistrySet)
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
		it.Event = new(FeeHandlerRegistrySet)
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
func (it *FeeHandlerRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRegistrySet represents a RegistrySet event raised by the FeeHandler contract.
type FeeHandlerRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_FeeHandler *FeeHandlerFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*FeeHandlerRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRegistrySetIterator{contract: _FeeHandler.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_FeeHandler *FeeHandlerFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *FeeHandlerRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRegistrySet)
				if err := _FeeHandler.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_FeeHandler *FeeHandlerFilterer) ParseRegistrySet(log types.Log) (*FeeHandlerRegistrySet, error) {
	event := new(FeeHandlerRegistrySet)
	if err := _FeeHandler.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterAddressRemovedIterator is returned from FilterRouterAddressRemoved and is used to iterate over the raw logs and unpacked data for RouterAddressRemoved events raised by the FeeHandler contract.
type FeeHandlerRouterAddressRemovedIterator struct {
	Event *FeeHandlerRouterAddressRemoved // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterAddressRemoved)
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
		it.Event = new(FeeHandlerRouterAddressRemoved)
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
func (it *FeeHandlerRouterAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterAddressRemoved represents a RouterAddressRemoved event raised by the FeeHandler contract.
type FeeHandlerRouterAddressRemoved struct {
	Token  common.Address
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterAddressRemoved is a free log retrieval operation binding the contract event 0x044c4b00bcc14b6c00430f73b8bc07f33aecb2387c7b188142d6d497342de89a.
//
// Solidity: event RouterAddressRemoved(address token, address router)
func (_FeeHandler *FeeHandlerFilterer) FilterRouterAddressRemoved(opts *bind.FilterOpts) (*FeeHandlerRouterAddressRemovedIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RouterAddressRemoved")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterAddressRemovedIterator{contract: _FeeHandler.contract, event: "RouterAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchRouterAddressRemoved is a free log subscription operation binding the contract event 0x044c4b00bcc14b6c00430f73b8bc07f33aecb2387c7b188142d6d497342de89a.
//
// Solidity: event RouterAddressRemoved(address token, address router)
func (_FeeHandler *FeeHandlerFilterer) WatchRouterAddressRemoved(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterAddressRemoved) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RouterAddressRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterAddressRemoved)
				if err := _FeeHandler.contract.UnpackLog(event, "RouterAddressRemoved", log); err != nil {
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

// ParseRouterAddressRemoved is a log parse operation binding the contract event 0x044c4b00bcc14b6c00430f73b8bc07f33aecb2387c7b188142d6d497342de89a.
//
// Solidity: event RouterAddressRemoved(address token, address router)
func (_FeeHandler *FeeHandlerFilterer) ParseRouterAddressRemoved(log types.Log) (*FeeHandlerRouterAddressRemoved, error) {
	event := new(FeeHandlerRouterAddressRemoved)
	if err := _FeeHandler.contract.UnpackLog(event, "RouterAddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterAddressSetIterator is returned from FilterRouterAddressSet and is used to iterate over the raw logs and unpacked data for RouterAddressSet events raised by the FeeHandler contract.
type FeeHandlerRouterAddressSetIterator struct {
	Event *FeeHandlerRouterAddressSet // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterAddressSet)
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
		it.Event = new(FeeHandlerRouterAddressSet)
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
func (it *FeeHandlerRouterAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterAddressSet represents a RouterAddressSet event raised by the FeeHandler contract.
type FeeHandlerRouterAddressSet struct {
	Token  common.Address
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterAddressSet is a free log retrieval operation binding the contract event 0xb3cbb74e835466bdbf8838b1acb70fa4a8b73e1a00cd5bacb9f68cf4dfc79cf3.
//
// Solidity: event RouterAddressSet(address token, address router)
func (_FeeHandler *FeeHandlerFilterer) FilterRouterAddressSet(opts *bind.FilterOpts) (*FeeHandlerRouterAddressSetIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RouterAddressSet")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterAddressSetIterator{contract: _FeeHandler.contract, event: "RouterAddressSet", logs: logs, sub: sub}, nil
}

// WatchRouterAddressSet is a free log subscription operation binding the contract event 0xb3cbb74e835466bdbf8838b1acb70fa4a8b73e1a00cd5bacb9f68cf4dfc79cf3.
//
// Solidity: event RouterAddressSet(address token, address router)
func (_FeeHandler *FeeHandlerFilterer) WatchRouterAddressSet(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterAddressSet) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RouterAddressSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterAddressSet)
				if err := _FeeHandler.contract.UnpackLog(event, "RouterAddressSet", log); err != nil {
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

// ParseRouterAddressSet is a log parse operation binding the contract event 0xb3cbb74e835466bdbf8838b1acb70fa4a8b73e1a00cd5bacb9f68cf4dfc79cf3.
//
// Solidity: event RouterAddressSet(address token, address router)
func (_FeeHandler *FeeHandlerFilterer) ParseRouterAddressSet(log types.Log) (*FeeHandlerRouterAddressSet, error) {
	event := new(FeeHandlerRouterAddressSet)
	if err := _FeeHandler.contract.UnpackLog(event, "RouterAddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerRouterUsedIterator is returned from FilterRouterUsed and is used to iterate over the raw logs and unpacked data for RouterUsed events raised by the FeeHandler contract.
type FeeHandlerRouterUsedIterator struct {
	Event *FeeHandlerRouterUsed // Event containing the contract specifics and raw log

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
func (it *FeeHandlerRouterUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerRouterUsed)
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
		it.Event = new(FeeHandlerRouterUsed)
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
func (it *FeeHandlerRouterUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerRouterUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerRouterUsed represents a RouterUsed event raised by the FeeHandler contract.
type FeeHandlerRouterUsed struct {
	Router common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRouterUsed is a free log retrieval operation binding the contract event 0x59afd9b3bf745a06d8739721397432b9f161243cee13868b9d6d6fca05b6e551.
//
// Solidity: event RouterUsed(address router)
func (_FeeHandler *FeeHandlerFilterer) FilterRouterUsed(opts *bind.FilterOpts) (*FeeHandlerRouterUsedIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "RouterUsed")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerRouterUsedIterator{contract: _FeeHandler.contract, event: "RouterUsed", logs: logs, sub: sub}, nil
}

// WatchRouterUsed is a free log subscription operation binding the contract event 0x59afd9b3bf745a06d8739721397432b9f161243cee13868b9d6d6fca05b6e551.
//
// Solidity: event RouterUsed(address router)
func (_FeeHandler *FeeHandlerFilterer) WatchRouterUsed(opts *bind.WatchOpts, sink chan<- *FeeHandlerRouterUsed) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "RouterUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerRouterUsed)
				if err := _FeeHandler.contract.UnpackLog(event, "RouterUsed", log); err != nil {
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

// ParseRouterUsed is a log parse operation binding the contract event 0x59afd9b3bf745a06d8739721397432b9f161243cee13868b9d6d6fca05b6e551.
//
// Solidity: event RouterUsed(address router)
func (_FeeHandler *FeeHandlerFilterer) ParseRouterUsed(log types.Log) (*FeeHandlerRouterUsed, error) {
	event := new(FeeHandlerRouterUsed)
	if err := _FeeHandler.contract.UnpackLog(event, "RouterUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeHandlerSoldAndBurnedTokenIterator is returned from FilterSoldAndBurnedToken and is used to iterate over the raw logs and unpacked data for SoldAndBurnedToken events raised by the FeeHandler contract.
type FeeHandlerSoldAndBurnedTokenIterator struct {
	Event *FeeHandlerSoldAndBurnedToken // Event containing the contract specifics and raw log

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
func (it *FeeHandlerSoldAndBurnedTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeHandlerSoldAndBurnedToken)
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
		it.Event = new(FeeHandlerSoldAndBurnedToken)
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
func (it *FeeHandlerSoldAndBurnedTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeHandlerSoldAndBurnedTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeHandlerSoldAndBurnedToken represents a SoldAndBurnedToken event raised by the FeeHandler contract.
type FeeHandlerSoldAndBurnedToken struct {
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSoldAndBurnedToken is a free log retrieval operation binding the contract event 0xac094032b4e9dccb3a000eedb94cf30146ca0d7c39be85229f478413fa21d1d8.
//
// Solidity: event SoldAndBurnedToken(address token, uint256 value)
func (_FeeHandler *FeeHandlerFilterer) FilterSoldAndBurnedToken(opts *bind.FilterOpts) (*FeeHandlerSoldAndBurnedTokenIterator, error) {

	logs, sub, err := _FeeHandler.contract.FilterLogs(opts, "SoldAndBurnedToken")
	if err != nil {
		return nil, err
	}
	return &FeeHandlerSoldAndBurnedTokenIterator{contract: _FeeHandler.contract, event: "SoldAndBurnedToken", logs: logs, sub: sub}, nil
}

// WatchSoldAndBurnedToken is a free log subscription operation binding the contract event 0xac094032b4e9dccb3a000eedb94cf30146ca0d7c39be85229f478413fa21d1d8.
//
// Solidity: event SoldAndBurnedToken(address token, uint256 value)
func (_FeeHandler *FeeHandlerFilterer) WatchSoldAndBurnedToken(opts *bind.WatchOpts, sink chan<- *FeeHandlerSoldAndBurnedToken) (event.Subscription, error) {

	logs, sub, err := _FeeHandler.contract.WatchLogs(opts, "SoldAndBurnedToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeHandlerSoldAndBurnedToken)
				if err := _FeeHandler.contract.UnpackLog(event, "SoldAndBurnedToken", log); err != nil {
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

// ParseSoldAndBurnedToken is a log parse operation binding the contract event 0xac094032b4e9dccb3a000eedb94cf30146ca0d7c39be85229f478413fa21d1d8.
//
// Solidity: event SoldAndBurnedToken(address token, uint256 value)
func (_FeeHandler *FeeHandlerFilterer) ParseSoldAndBurnedToken(log types.Log) (*FeeHandlerSoldAndBurnedToken, error) {
	event := new(FeeHandlerSoldAndBurnedToken)
	if err := _FeeHandler.contract.UnpackLog(event, "SoldAndBurnedToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
