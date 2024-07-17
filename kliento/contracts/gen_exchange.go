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

// ExchangeMetaData contains all meta data concerning the Exchange contract.
var ExchangeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"goldBucket\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBucket\",\"type\":\"uint256\"}],\"name\":\"BucketsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exchanger\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"soldGold\",\"type\":\"bool\"}],\"name\":\"Exchanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minimumReports\",\"type\":\"uint256\"}],\"name\":\"MinimumReportsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserveFraction\",\"type\":\"uint256\"}],\"name\":\"ReserveFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"spread\",\"type\":\"uint256\"}],\"name\":\"SpreadSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"stable\",\"type\":\"address\"}],\"name\":\"StableTokenSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updateFrequency\",\"type\":\"uint256\"}],\"name\":\"UpdateFrequencySet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"goldBucket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastBucketUpdate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minimumReports\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spread\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stable\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableBucket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stableTokenRegistryId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"updateFrequency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"stableTokenIdentifier\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_spread\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserveFraction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_updateFrequency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumReports\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"activateStable\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBuyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sellGold\",\"type\":\"bool\"}],\"name\":\"sell\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBuyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sellGold\",\"type\":\"bool\"}],\"name\":\"exchange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxSellAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"buyGold\",\"type\":\"bool\"}],\"name\":\"buy\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sellAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sellGold\",\"type\":\"bool\"}],\"name\":\"getBuyTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sellGold\",\"type\":\"bool\"}],\"name\":\"getSellTokenAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"sellGold\",\"type\":\"bool\"}],\"name\":\"getBuyAndSellBuckets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newUpdateFrequency\",\"type\":\"uint256\"}],\"name\":\"setUpdateFrequency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newMininumReports\",\"type\":\"uint256\"}],\"name\":\"setMinimumReports\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newStableToken\",\"type\":\"address\"}],\"name\":\"setStableToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newSpread\",\"type\":\"uint256\"}],\"name\":\"setSpread\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFraction\",\"type\":\"uint256\"}],\"name\":\"setReserveFraction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ExchangeABI is the input ABI used to generate the binding from.
// Deprecated: Use ExchangeMetaData.ABI instead.
var ExchangeABI = ExchangeMetaData.ABI

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseExchangeABI parses the ABI
func ParseExchangeABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// GetBuyAndSellBuckets is a free data retrieval call binding the contract method 0x78ba9cfd.
//
// Solidity: function getBuyAndSellBuckets(bool sellGold) view returns(uint256, uint256)
func (_Exchange *ExchangeCaller) GetBuyAndSellBuckets(opts *bind.CallOpts, sellGold bool) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getBuyAndSellBuckets", sellGold)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBuyAndSellBuckets is a free data retrieval call binding the contract method 0x78ba9cfd.
//
// Solidity: function getBuyAndSellBuckets(bool sellGold) view returns(uint256, uint256)
func (_Exchange *ExchangeSession) GetBuyAndSellBuckets(sellGold bool) (*big.Int, *big.Int, error) {
	return _Exchange.Contract.GetBuyAndSellBuckets(&_Exchange.CallOpts, sellGold)
}

// GetBuyAndSellBuckets is a free data retrieval call binding the contract method 0x78ba9cfd.
//
// Solidity: function getBuyAndSellBuckets(bool sellGold) view returns(uint256, uint256)
func (_Exchange *ExchangeCallerSession) GetBuyAndSellBuckets(sellGold bool) (*big.Int, *big.Int, error) {
	return _Exchange.Contract.GetBuyAndSellBuckets(&_Exchange.CallOpts, sellGold)
}

// GetBuyTokenAmount is a free data retrieval call binding the contract method 0x9ed02b58.
//
// Solidity: function getBuyTokenAmount(uint256 sellAmount, bool sellGold) view returns(uint256)
func (_Exchange *ExchangeCaller) GetBuyTokenAmount(opts *bind.CallOpts, sellAmount *big.Int, sellGold bool) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getBuyTokenAmount", sellAmount, sellGold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBuyTokenAmount is a free data retrieval call binding the contract method 0x9ed02b58.
//
// Solidity: function getBuyTokenAmount(uint256 sellAmount, bool sellGold) view returns(uint256)
func (_Exchange *ExchangeSession) GetBuyTokenAmount(sellAmount *big.Int, sellGold bool) (*big.Int, error) {
	return _Exchange.Contract.GetBuyTokenAmount(&_Exchange.CallOpts, sellAmount, sellGold)
}

// GetBuyTokenAmount is a free data retrieval call binding the contract method 0x9ed02b58.
//
// Solidity: function getBuyTokenAmount(uint256 sellAmount, bool sellGold) view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetBuyTokenAmount(sellAmount *big.Int, sellGold bool) (*big.Int, error) {
	return _Exchange.Contract.GetBuyTokenAmount(&_Exchange.CallOpts, sellAmount, sellGold)
}

// GetSellTokenAmount is a free data retrieval call binding the contract method 0x4c0226a2.
//
// Solidity: function getSellTokenAmount(uint256 buyAmount, bool sellGold) view returns(uint256)
func (_Exchange *ExchangeCaller) GetSellTokenAmount(opts *bind.CallOpts, buyAmount *big.Int, sellGold bool) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getSellTokenAmount", buyAmount, sellGold)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSellTokenAmount is a free data retrieval call binding the contract method 0x4c0226a2.
//
// Solidity: function getSellTokenAmount(uint256 buyAmount, bool sellGold) view returns(uint256)
func (_Exchange *ExchangeSession) GetSellTokenAmount(buyAmount *big.Int, sellGold bool) (*big.Int, error) {
	return _Exchange.Contract.GetSellTokenAmount(&_Exchange.CallOpts, buyAmount, sellGold)
}

// GetSellTokenAmount is a free data retrieval call binding the contract method 0x4c0226a2.
//
// Solidity: function getSellTokenAmount(uint256 buyAmount, bool sellGold) view returns(uint256)
func (_Exchange *ExchangeCallerSession) GetSellTokenAmount(buyAmount *big.Int, sellGold bool) (*big.Int, error) {
	return _Exchange.Contract.GetSellTokenAmount(&_Exchange.CallOpts, buyAmount, sellGold)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Exchange *ExchangeCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "getVersionNumber")

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
func (_Exchange *ExchangeSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Exchange.Contract.GetVersionNumber(&_Exchange.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Exchange *ExchangeCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Exchange.Contract.GetVersionNumber(&_Exchange.CallOpts)
}

// GoldBucket is a free data retrieval call binding the contract method 0x62f09084.
//
// Solidity: function goldBucket() view returns(uint256)
func (_Exchange *ExchangeCaller) GoldBucket(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "goldBucket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GoldBucket is a free data retrieval call binding the contract method 0x62f09084.
//
// Solidity: function goldBucket() view returns(uint256)
func (_Exchange *ExchangeSession) GoldBucket() (*big.Int, error) {
	return _Exchange.Contract.GoldBucket(&_Exchange.CallOpts)
}

// GoldBucket is a free data retrieval call binding the contract method 0x62f09084.
//
// Solidity: function goldBucket() view returns(uint256)
func (_Exchange *ExchangeCallerSession) GoldBucket() (*big.Int, error) {
	return _Exchange.Contract.GoldBucket(&_Exchange.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Exchange *ExchangeCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Exchange *ExchangeSession) Initialized() (bool, error) {
	return _Exchange.Contract.Initialized(&_Exchange.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Exchange *ExchangeCallerSession) Initialized() (bool, error) {
	return _Exchange.Contract.Initialized(&_Exchange.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Exchange *ExchangeCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Exchange *ExchangeSession) IsOwner() (bool, error) {
	return _Exchange.Contract.IsOwner(&_Exchange.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Exchange *ExchangeCallerSession) IsOwner() (bool, error) {
	return _Exchange.Contract.IsOwner(&_Exchange.CallOpts)
}

// LastBucketUpdate is a free data retrieval call binding the contract method 0xe0c8b50a.
//
// Solidity: function lastBucketUpdate() view returns(uint256)
func (_Exchange *ExchangeCaller) LastBucketUpdate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "lastBucketUpdate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastBucketUpdate is a free data retrieval call binding the contract method 0xe0c8b50a.
//
// Solidity: function lastBucketUpdate() view returns(uint256)
func (_Exchange *ExchangeSession) LastBucketUpdate() (*big.Int, error) {
	return _Exchange.Contract.LastBucketUpdate(&_Exchange.CallOpts)
}

// LastBucketUpdate is a free data retrieval call binding the contract method 0xe0c8b50a.
//
// Solidity: function lastBucketUpdate() view returns(uint256)
func (_Exchange *ExchangeCallerSession) LastBucketUpdate() (*big.Int, error) {
	return _Exchange.Contract.LastBucketUpdate(&_Exchange.CallOpts)
}

// MinimumReports is a free data retrieval call binding the contract method 0x22503ce5.
//
// Solidity: function minimumReports() view returns(uint256)
func (_Exchange *ExchangeCaller) MinimumReports(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "minimumReports")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumReports is a free data retrieval call binding the contract method 0x22503ce5.
//
// Solidity: function minimumReports() view returns(uint256)
func (_Exchange *ExchangeSession) MinimumReports() (*big.Int, error) {
	return _Exchange.Contract.MinimumReports(&_Exchange.CallOpts)
}

// MinimumReports is a free data retrieval call binding the contract method 0x22503ce5.
//
// Solidity: function minimumReports() view returns(uint256)
func (_Exchange *ExchangeCallerSession) MinimumReports() (*big.Int, error) {
	return _Exchange.Contract.MinimumReports(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Exchange *ExchangeCallerSession) Owner() (common.Address, error) {
	return _Exchange.Contract.Owner(&_Exchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Exchange *ExchangeCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Exchange *ExchangeSession) Registry() (common.Address, error) {
	return _Exchange.Contract.Registry(&_Exchange.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Exchange *ExchangeCallerSession) Registry() (common.Address, error) {
	return _Exchange.Contract.Registry(&_Exchange.CallOpts)
}

// ReserveFraction is a free data retrieval call binding the contract method 0xdda57b93.
//
// Solidity: function reserveFraction() view returns(uint256 value)
func (_Exchange *ExchangeCaller) ReserveFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "reserveFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFraction is a free data retrieval call binding the contract method 0xdda57b93.
//
// Solidity: function reserveFraction() view returns(uint256 value)
func (_Exchange *ExchangeSession) ReserveFraction() (*big.Int, error) {
	return _Exchange.Contract.ReserveFraction(&_Exchange.CallOpts)
}

// ReserveFraction is a free data retrieval call binding the contract method 0xdda57b93.
//
// Solidity: function reserveFraction() view returns(uint256 value)
func (_Exchange *ExchangeCallerSession) ReserveFraction() (*big.Int, error) {
	return _Exchange.Contract.ReserveFraction(&_Exchange.CallOpts)
}

// Spread is a free data retrieval call binding the contract method 0x5c25c76c.
//
// Solidity: function spread() view returns(uint256 value)
func (_Exchange *ExchangeCaller) Spread(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "spread")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Spread is a free data retrieval call binding the contract method 0x5c25c76c.
//
// Solidity: function spread() view returns(uint256 value)
func (_Exchange *ExchangeSession) Spread() (*big.Int, error) {
	return _Exchange.Contract.Spread(&_Exchange.CallOpts)
}

// Spread is a free data retrieval call binding the contract method 0x5c25c76c.
//
// Solidity: function spread() view returns(uint256 value)
func (_Exchange *ExchangeCallerSession) Spread() (*big.Int, error) {
	return _Exchange.Contract.Spread(&_Exchange.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_Exchange *ExchangeCaller) Stable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "stable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_Exchange *ExchangeSession) Stable() (common.Address, error) {
	return _Exchange.Contract.Stable(&_Exchange.CallOpts)
}

// Stable is a free data retrieval call binding the contract method 0x22be3de1.
//
// Solidity: function stable() view returns(address)
func (_Exchange *ExchangeCallerSession) Stable() (common.Address, error) {
	return _Exchange.Contract.Stable(&_Exchange.CallOpts)
}

// StableBucket is a free data retrieval call binding the contract method 0x25ac2de6.
//
// Solidity: function stableBucket() view returns(uint256)
func (_Exchange *ExchangeCaller) StableBucket(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "stableBucket")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StableBucket is a free data retrieval call binding the contract method 0x25ac2de6.
//
// Solidity: function stableBucket() view returns(uint256)
func (_Exchange *ExchangeSession) StableBucket() (*big.Int, error) {
	return _Exchange.Contract.StableBucket(&_Exchange.CallOpts)
}

// StableBucket is a free data retrieval call binding the contract method 0x25ac2de6.
//
// Solidity: function stableBucket() view returns(uint256)
func (_Exchange *ExchangeCallerSession) StableBucket() (*big.Int, error) {
	return _Exchange.Contract.StableBucket(&_Exchange.CallOpts)
}

// StableTokenRegistryId is a free data retrieval call binding the contract method 0x81bb18ec.
//
// Solidity: function stableTokenRegistryId() view returns(bytes32)
func (_Exchange *ExchangeCaller) StableTokenRegistryId(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "stableTokenRegistryId")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// StableTokenRegistryId is a free data retrieval call binding the contract method 0x81bb18ec.
//
// Solidity: function stableTokenRegistryId() view returns(bytes32)
func (_Exchange *ExchangeSession) StableTokenRegistryId() ([32]byte, error) {
	return _Exchange.Contract.StableTokenRegistryId(&_Exchange.CallOpts)
}

// StableTokenRegistryId is a free data retrieval call binding the contract method 0x81bb18ec.
//
// Solidity: function stableTokenRegistryId() view returns(bytes32)
func (_Exchange *ExchangeCallerSession) StableTokenRegistryId() ([32]byte, error) {
	return _Exchange.Contract.StableTokenRegistryId(&_Exchange.CallOpts)
}

// UpdateFrequency is a free data retrieval call binding the contract method 0x673ea086.
//
// Solidity: function updateFrequency() view returns(uint256)
func (_Exchange *ExchangeCaller) UpdateFrequency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Exchange.contract.Call(opts, &out, "updateFrequency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UpdateFrequency is a free data retrieval call binding the contract method 0x673ea086.
//
// Solidity: function updateFrequency() view returns(uint256)
func (_Exchange *ExchangeSession) UpdateFrequency() (*big.Int, error) {
	return _Exchange.Contract.UpdateFrequency(&_Exchange.CallOpts)
}

// UpdateFrequency is a free data retrieval call binding the contract method 0x673ea086.
//
// Solidity: function updateFrequency() view returns(uint256)
func (_Exchange *ExchangeCallerSession) UpdateFrequency() (*big.Int, error) {
	return _Exchange.Contract.UpdateFrequency(&_Exchange.CallOpts)
}

// ActivateStable is a paid mutator transaction binding the contract method 0xd1a2bc11.
//
// Solidity: function activateStable() returns()
func (_Exchange *ExchangeTransactor) ActivateStable(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "activateStable")
}

// ActivateStable is a paid mutator transaction binding the contract method 0xd1a2bc11.
//
// Solidity: function activateStable() returns()
func (_Exchange *ExchangeSession) ActivateStable() (*types.Transaction, error) {
	return _Exchange.Contract.ActivateStable(&_Exchange.TransactOpts)
}

// ActivateStable is a paid mutator transaction binding the contract method 0xd1a2bc11.
//
// Solidity: function activateStable() returns()
func (_Exchange *ExchangeTransactorSession) ActivateStable() (*types.Transaction, error) {
	return _Exchange.Contract.ActivateStable(&_Exchange.TransactOpts)
}

// Buy is a paid mutator transaction binding the contract method 0xc3434883.
//
// Solidity: function buy(uint256 buyAmount, uint256 maxSellAmount, bool buyGold) returns(uint256)
func (_Exchange *ExchangeTransactor) Buy(opts *bind.TransactOpts, buyAmount *big.Int, maxSellAmount *big.Int, buyGold bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "buy", buyAmount, maxSellAmount, buyGold)
}

// Buy is a paid mutator transaction binding the contract method 0xc3434883.
//
// Solidity: function buy(uint256 buyAmount, uint256 maxSellAmount, bool buyGold) returns(uint256)
func (_Exchange *ExchangeSession) Buy(buyAmount *big.Int, maxSellAmount *big.Int, buyGold bool) (*types.Transaction, error) {
	return _Exchange.Contract.Buy(&_Exchange.TransactOpts, buyAmount, maxSellAmount, buyGold)
}

// Buy is a paid mutator transaction binding the contract method 0xc3434883.
//
// Solidity: function buy(uint256 buyAmount, uint256 maxSellAmount, bool buyGold) returns(uint256)
func (_Exchange *ExchangeTransactorSession) Buy(buyAmount *big.Int, maxSellAmount *big.Int, buyGold bool) (*types.Transaction, error) {
	return _Exchange.Contract.Buy(&_Exchange.TransactOpts, buyAmount, maxSellAmount, buyGold)
}

// Exchange is a paid mutator transaction binding the contract method 0x2bc7d67a.
//
// Solidity: function exchange(uint256 sellAmount, uint256 minBuyAmount, bool sellGold) returns(uint256)
func (_Exchange *ExchangeTransactor) Exchange(opts *bind.TransactOpts, sellAmount *big.Int, minBuyAmount *big.Int, sellGold bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "exchange", sellAmount, minBuyAmount, sellGold)
}

// Exchange is a paid mutator transaction binding the contract method 0x2bc7d67a.
//
// Solidity: function exchange(uint256 sellAmount, uint256 minBuyAmount, bool sellGold) returns(uint256)
func (_Exchange *ExchangeSession) Exchange(sellAmount *big.Int, minBuyAmount *big.Int, sellGold bool) (*types.Transaction, error) {
	return _Exchange.Contract.Exchange(&_Exchange.TransactOpts, sellAmount, minBuyAmount, sellGold)
}

// Exchange is a paid mutator transaction binding the contract method 0x2bc7d67a.
//
// Solidity: function exchange(uint256 sellAmount, uint256 minBuyAmount, bool sellGold) returns(uint256)
func (_Exchange *ExchangeTransactorSession) Exchange(sellAmount *big.Int, minBuyAmount *big.Int, sellGold bool) (*types.Transaction, error) {
	return _Exchange.Contract.Exchange(&_Exchange.TransactOpts, sellAmount, minBuyAmount, sellGold)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf317778.
//
// Solidity: function initialize(address registryAddress, string stableTokenIdentifier, uint256 _spread, uint256 _reserveFraction, uint256 _updateFrequency, uint256 _minimumReports) returns()
func (_Exchange *ExchangeTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, stableTokenIdentifier string, _spread *big.Int, _reserveFraction *big.Int, _updateFrequency *big.Int, _minimumReports *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "initialize", registryAddress, stableTokenIdentifier, _spread, _reserveFraction, _updateFrequency, _minimumReports)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf317778.
//
// Solidity: function initialize(address registryAddress, string stableTokenIdentifier, uint256 _spread, uint256 _reserveFraction, uint256 _updateFrequency, uint256 _minimumReports) returns()
func (_Exchange *ExchangeSession) Initialize(registryAddress common.Address, stableTokenIdentifier string, _spread *big.Int, _reserveFraction *big.Int, _updateFrequency *big.Int, _minimumReports *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, registryAddress, stableTokenIdentifier, _spread, _reserveFraction, _updateFrequency, _minimumReports)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf317778.
//
// Solidity: function initialize(address registryAddress, string stableTokenIdentifier, uint256 _spread, uint256 _reserveFraction, uint256 _updateFrequency, uint256 _minimumReports) returns()
func (_Exchange *ExchangeTransactorSession) Initialize(registryAddress common.Address, stableTokenIdentifier string, _spread *big.Int, _reserveFraction *big.Int, _updateFrequency *big.Int, _minimumReports *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.Initialize(&_Exchange.TransactOpts, registryAddress, stableTokenIdentifier, _spread, _reserveFraction, _updateFrequency, _minimumReports)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Exchange *ExchangeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Exchange.Contract.RenounceOwnership(&_Exchange.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x8ab1a5d4.
//
// Solidity: function sell(uint256 sellAmount, uint256 minBuyAmount, bool sellGold) returns(uint256)
func (_Exchange *ExchangeTransactor) Sell(opts *bind.TransactOpts, sellAmount *big.Int, minBuyAmount *big.Int, sellGold bool) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "sell", sellAmount, minBuyAmount, sellGold)
}

// Sell is a paid mutator transaction binding the contract method 0x8ab1a5d4.
//
// Solidity: function sell(uint256 sellAmount, uint256 minBuyAmount, bool sellGold) returns(uint256)
func (_Exchange *ExchangeSession) Sell(sellAmount *big.Int, minBuyAmount *big.Int, sellGold bool) (*types.Transaction, error) {
	return _Exchange.Contract.Sell(&_Exchange.TransactOpts, sellAmount, minBuyAmount, sellGold)
}

// Sell is a paid mutator transaction binding the contract method 0x8ab1a5d4.
//
// Solidity: function sell(uint256 sellAmount, uint256 minBuyAmount, bool sellGold) returns(uint256)
func (_Exchange *ExchangeTransactorSession) Sell(sellAmount *big.Int, minBuyAmount *big.Int, sellGold bool) (*types.Transaction, error) {
	return _Exchange.Contract.Sell(&_Exchange.TransactOpts, sellAmount, minBuyAmount, sellGold)
}

// SetMinimumReports is a paid mutator transaction binding the contract method 0xd404f7f8.
//
// Solidity: function setMinimumReports(uint256 newMininumReports) returns()
func (_Exchange *ExchangeTransactor) SetMinimumReports(opts *bind.TransactOpts, newMininumReports *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setMinimumReports", newMininumReports)
}

// SetMinimumReports is a paid mutator transaction binding the contract method 0xd404f7f8.
//
// Solidity: function setMinimumReports(uint256 newMininumReports) returns()
func (_Exchange *ExchangeSession) SetMinimumReports(newMininumReports *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetMinimumReports(&_Exchange.TransactOpts, newMininumReports)
}

// SetMinimumReports is a paid mutator transaction binding the contract method 0xd404f7f8.
//
// Solidity: function setMinimumReports(uint256 newMininumReports) returns()
func (_Exchange *ExchangeTransactorSession) SetMinimumReports(newMininumReports *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetMinimumReports(&_Exchange.TransactOpts, newMininumReports)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Exchange *ExchangeTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Exchange *ExchangeSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetRegistry(&_Exchange.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Exchange *ExchangeTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetRegistry(&_Exchange.TransactOpts, registryAddress)
}

// SetReserveFraction is a paid mutator transaction binding the contract method 0x6a5eaf47.
//
// Solidity: function setReserveFraction(uint256 newReserveFraction) returns()
func (_Exchange *ExchangeTransactor) SetReserveFraction(opts *bind.TransactOpts, newReserveFraction *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setReserveFraction", newReserveFraction)
}

// SetReserveFraction is a paid mutator transaction binding the contract method 0x6a5eaf47.
//
// Solidity: function setReserveFraction(uint256 newReserveFraction) returns()
func (_Exchange *ExchangeSession) SetReserveFraction(newReserveFraction *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetReserveFraction(&_Exchange.TransactOpts, newReserveFraction)
}

// SetReserveFraction is a paid mutator transaction binding the contract method 0x6a5eaf47.
//
// Solidity: function setReserveFraction(uint256 newReserveFraction) returns()
func (_Exchange *ExchangeTransactorSession) SetReserveFraction(newReserveFraction *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetReserveFraction(&_Exchange.TransactOpts, newReserveFraction)
}

// SetSpread is a paid mutator transaction binding the contract method 0xb66a261c.
//
// Solidity: function setSpread(uint256 newSpread) returns()
func (_Exchange *ExchangeTransactor) SetSpread(opts *bind.TransactOpts, newSpread *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setSpread", newSpread)
}

// SetSpread is a paid mutator transaction binding the contract method 0xb66a261c.
//
// Solidity: function setSpread(uint256 newSpread) returns()
func (_Exchange *ExchangeSession) SetSpread(newSpread *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetSpread(&_Exchange.TransactOpts, newSpread)
}

// SetSpread is a paid mutator transaction binding the contract method 0xb66a261c.
//
// Solidity: function setSpread(uint256 newSpread) returns()
func (_Exchange *ExchangeTransactorSession) SetSpread(newSpread *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetSpread(&_Exchange.TransactOpts, newSpread)
}

// SetStableToken is a paid mutator transaction binding the contract method 0xdb1bc87b.
//
// Solidity: function setStableToken(address newStableToken) returns()
func (_Exchange *ExchangeTransactor) SetStableToken(opts *bind.TransactOpts, newStableToken common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setStableToken", newStableToken)
}

// SetStableToken is a paid mutator transaction binding the contract method 0xdb1bc87b.
//
// Solidity: function setStableToken(address newStableToken) returns()
func (_Exchange *ExchangeSession) SetStableToken(newStableToken common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetStableToken(&_Exchange.TransactOpts, newStableToken)
}

// SetStableToken is a paid mutator transaction binding the contract method 0xdb1bc87b.
//
// Solidity: function setStableToken(address newStableToken) returns()
func (_Exchange *ExchangeTransactorSession) SetStableToken(newStableToken common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.SetStableToken(&_Exchange.TransactOpts, newStableToken)
}

// SetUpdateFrequency is a paid mutator transaction binding the contract method 0x4a1be6cb.
//
// Solidity: function setUpdateFrequency(uint256 newUpdateFrequency) returns()
func (_Exchange *ExchangeTransactor) SetUpdateFrequency(opts *bind.TransactOpts, newUpdateFrequency *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "setUpdateFrequency", newUpdateFrequency)
}

// SetUpdateFrequency is a paid mutator transaction binding the contract method 0x4a1be6cb.
//
// Solidity: function setUpdateFrequency(uint256 newUpdateFrequency) returns()
func (_Exchange *ExchangeSession) SetUpdateFrequency(newUpdateFrequency *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetUpdateFrequency(&_Exchange.TransactOpts, newUpdateFrequency)
}

// SetUpdateFrequency is a paid mutator transaction binding the contract method 0x4a1be6cb.
//
// Solidity: function setUpdateFrequency(uint256 newUpdateFrequency) returns()
func (_Exchange *ExchangeTransactorSession) SetUpdateFrequency(newUpdateFrequency *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.SetUpdateFrequency(&_Exchange.TransactOpts, newUpdateFrequency)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Exchange *ExchangeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Exchange.Contract.TransferOwnership(&_Exchange.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Exchange *ExchangeFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Exchange.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "BucketsUpdated":
		event, err = _Exchange.ParseBucketsUpdated(log)
	case "Exchanged":
		event, err = _Exchange.ParseExchanged(log)
	case "MinimumReportsSet":
		event, err = _Exchange.ParseMinimumReportsSet(log)
	case "OwnershipTransferred":
		event, err = _Exchange.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Exchange.ParseRegistrySet(log)
	case "ReserveFractionSet":
		event, err = _Exchange.ParseReserveFractionSet(log)
	case "SpreadSet":
		event, err = _Exchange.ParseSpreadSet(log)
	case "StableTokenSet":
		event, err = _Exchange.ParseStableTokenSet(log)
	case "UpdateFrequencySet":
		event, err = _Exchange.ParseUpdateFrequencySet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// ExchangeBucketsUpdatedIterator is returned from FilterBucketsUpdated and is used to iterate over the raw logs and unpacked data for BucketsUpdated events raised by the Exchange contract.
type ExchangeBucketsUpdatedIterator struct {
	Event *ExchangeBucketsUpdated // Event containing the contract specifics and raw log

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
func (it *ExchangeBucketsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeBucketsUpdated)
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
		it.Event = new(ExchangeBucketsUpdated)
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
func (it *ExchangeBucketsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeBucketsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeBucketsUpdated represents a BucketsUpdated event raised by the Exchange contract.
type ExchangeBucketsUpdated struct {
	GoldBucket   *big.Int
	StableBucket *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBucketsUpdated is a free log retrieval operation binding the contract event 0xa18ec663cb684011386aa866c4dacb32d2d2ad859a35d3440b6ce7200a76bad8.
//
// Solidity: event BucketsUpdated(uint256 goldBucket, uint256 stableBucket)
func (_Exchange *ExchangeFilterer) FilterBucketsUpdated(opts *bind.FilterOpts) (*ExchangeBucketsUpdatedIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "BucketsUpdated")
	if err != nil {
		return nil, err
	}
	return &ExchangeBucketsUpdatedIterator{contract: _Exchange.contract, event: "BucketsUpdated", logs: logs, sub: sub}, nil
}

// WatchBucketsUpdated is a free log subscription operation binding the contract event 0xa18ec663cb684011386aa866c4dacb32d2d2ad859a35d3440b6ce7200a76bad8.
//
// Solidity: event BucketsUpdated(uint256 goldBucket, uint256 stableBucket)
func (_Exchange *ExchangeFilterer) WatchBucketsUpdated(opts *bind.WatchOpts, sink chan<- *ExchangeBucketsUpdated) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "BucketsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeBucketsUpdated)
				if err := _Exchange.contract.UnpackLog(event, "BucketsUpdated", log); err != nil {
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

// ParseBucketsUpdated is a log parse operation binding the contract event 0xa18ec663cb684011386aa866c4dacb32d2d2ad859a35d3440b6ce7200a76bad8.
//
// Solidity: event BucketsUpdated(uint256 goldBucket, uint256 stableBucket)
func (_Exchange *ExchangeFilterer) ParseBucketsUpdated(log types.Log) (*ExchangeBucketsUpdated, error) {
	event := new(ExchangeBucketsUpdated)
	if err := _Exchange.contract.UnpackLog(event, "BucketsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeExchangedIterator is returned from FilterExchanged and is used to iterate over the raw logs and unpacked data for Exchanged events raised by the Exchange contract.
type ExchangeExchangedIterator struct {
	Event *ExchangeExchanged // Event containing the contract specifics and raw log

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
func (it *ExchangeExchangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeExchanged)
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
		it.Event = new(ExchangeExchanged)
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
func (it *ExchangeExchangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeExchangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeExchanged represents a Exchanged event raised by the Exchange contract.
type ExchangeExchanged struct {
	Exchanger  common.Address
	SellAmount *big.Int
	BuyAmount  *big.Int
	SoldGold   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterExchanged is a free log retrieval operation binding the contract event 0x402ac9185b4616422c2794bf5b118bfcc68ed496d52c0d9841dfa114fdeb05ba.
//
// Solidity: event Exchanged(address indexed exchanger, uint256 sellAmount, uint256 buyAmount, bool soldGold)
func (_Exchange *ExchangeFilterer) FilterExchanged(opts *bind.FilterOpts, exchanger []common.Address) (*ExchangeExchangedIterator, error) {

	var exchangerRule []interface{}
	for _, exchangerItem := range exchanger {
		exchangerRule = append(exchangerRule, exchangerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "Exchanged", exchangerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeExchangedIterator{contract: _Exchange.contract, event: "Exchanged", logs: logs, sub: sub}, nil
}

// WatchExchanged is a free log subscription operation binding the contract event 0x402ac9185b4616422c2794bf5b118bfcc68ed496d52c0d9841dfa114fdeb05ba.
//
// Solidity: event Exchanged(address indexed exchanger, uint256 sellAmount, uint256 buyAmount, bool soldGold)
func (_Exchange *ExchangeFilterer) WatchExchanged(opts *bind.WatchOpts, sink chan<- *ExchangeExchanged, exchanger []common.Address) (event.Subscription, error) {

	var exchangerRule []interface{}
	for _, exchangerItem := range exchanger {
		exchangerRule = append(exchangerRule, exchangerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "Exchanged", exchangerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeExchanged)
				if err := _Exchange.contract.UnpackLog(event, "Exchanged", log); err != nil {
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

// ParseExchanged is a log parse operation binding the contract event 0x402ac9185b4616422c2794bf5b118bfcc68ed496d52c0d9841dfa114fdeb05ba.
//
// Solidity: event Exchanged(address indexed exchanger, uint256 sellAmount, uint256 buyAmount, bool soldGold)
func (_Exchange *ExchangeFilterer) ParseExchanged(log types.Log) (*ExchangeExchanged, error) {
	event := new(ExchangeExchanged)
	if err := _Exchange.contract.UnpackLog(event, "Exchanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeMinimumReportsSetIterator is returned from FilterMinimumReportsSet and is used to iterate over the raw logs and unpacked data for MinimumReportsSet events raised by the Exchange contract.
type ExchangeMinimumReportsSetIterator struct {
	Event *ExchangeMinimumReportsSet // Event containing the contract specifics and raw log

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
func (it *ExchangeMinimumReportsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeMinimumReportsSet)
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
		it.Event = new(ExchangeMinimumReportsSet)
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
func (it *ExchangeMinimumReportsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeMinimumReportsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeMinimumReportsSet represents a MinimumReportsSet event raised by the Exchange contract.
type ExchangeMinimumReportsSet struct {
	MinimumReports *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMinimumReportsSet is a free log retrieval operation binding the contract event 0x08523596abc266fb46d9c40ddf78fdfd3c08142252833ddce1a2b46f76521035.
//
// Solidity: event MinimumReportsSet(uint256 minimumReports)
func (_Exchange *ExchangeFilterer) FilterMinimumReportsSet(opts *bind.FilterOpts) (*ExchangeMinimumReportsSetIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "MinimumReportsSet")
	if err != nil {
		return nil, err
	}
	return &ExchangeMinimumReportsSetIterator{contract: _Exchange.contract, event: "MinimumReportsSet", logs: logs, sub: sub}, nil
}

// WatchMinimumReportsSet is a free log subscription operation binding the contract event 0x08523596abc266fb46d9c40ddf78fdfd3c08142252833ddce1a2b46f76521035.
//
// Solidity: event MinimumReportsSet(uint256 minimumReports)
func (_Exchange *ExchangeFilterer) WatchMinimumReportsSet(opts *bind.WatchOpts, sink chan<- *ExchangeMinimumReportsSet) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "MinimumReportsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeMinimumReportsSet)
				if err := _Exchange.contract.UnpackLog(event, "MinimumReportsSet", log); err != nil {
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

// ParseMinimumReportsSet is a log parse operation binding the contract event 0x08523596abc266fb46d9c40ddf78fdfd3c08142252833ddce1a2b46f76521035.
//
// Solidity: event MinimumReportsSet(uint256 minimumReports)
func (_Exchange *ExchangeFilterer) ParseMinimumReportsSet(log types.Log) (*ExchangeMinimumReportsSet, error) {
	event := new(ExchangeMinimumReportsSet)
	if err := _Exchange.contract.UnpackLog(event, "MinimumReportsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Exchange contract.
type ExchangeOwnershipTransferredIterator struct {
	Event *ExchangeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ExchangeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeOwnershipTransferred)
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
		it.Event = new(ExchangeOwnershipTransferred)
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
func (it *ExchangeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeOwnershipTransferred represents a OwnershipTransferred event raised by the Exchange contract.
type ExchangeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ExchangeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeOwnershipTransferredIterator{contract: _Exchange.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Exchange *ExchangeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ExchangeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeOwnershipTransferred)
				if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseOwnershipTransferred(log types.Log) (*ExchangeOwnershipTransferred, error) {
	event := new(ExchangeOwnershipTransferred)
	if err := _Exchange.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Exchange contract.
type ExchangeRegistrySetIterator struct {
	Event *ExchangeRegistrySet // Event containing the contract specifics and raw log

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
func (it *ExchangeRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRegistrySet)
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
		it.Event = new(ExchangeRegistrySet)
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
func (it *ExchangeRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRegistrySet represents a RegistrySet event raised by the Exchange contract.
type ExchangeRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Exchange *ExchangeFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*ExchangeRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeRegistrySetIterator{contract: _Exchange.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Exchange *ExchangeFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *ExchangeRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRegistrySet)
				if err := _Exchange.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Exchange *ExchangeFilterer) ParseRegistrySet(log types.Log) (*ExchangeRegistrySet, error) {
	event := new(ExchangeRegistrySet)
	if err := _Exchange.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeReserveFractionSetIterator is returned from FilterReserveFractionSet and is used to iterate over the raw logs and unpacked data for ReserveFractionSet events raised by the Exchange contract.
type ExchangeReserveFractionSetIterator struct {
	Event *ExchangeReserveFractionSet // Event containing the contract specifics and raw log

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
func (it *ExchangeReserveFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeReserveFractionSet)
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
		it.Event = new(ExchangeReserveFractionSet)
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
func (it *ExchangeReserveFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeReserveFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeReserveFractionSet represents a ReserveFractionSet event raised by the Exchange contract.
type ExchangeReserveFractionSet struct {
	ReserveFraction *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterReserveFractionSet is a free log retrieval operation binding the contract event 0xb690f84efb1d9039c2834effb7bebc792a85bfec7ef84f4b269528454f363ccf.
//
// Solidity: event ReserveFractionSet(uint256 reserveFraction)
func (_Exchange *ExchangeFilterer) FilterReserveFractionSet(opts *bind.FilterOpts) (*ExchangeReserveFractionSetIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "ReserveFractionSet")
	if err != nil {
		return nil, err
	}
	return &ExchangeReserveFractionSetIterator{contract: _Exchange.contract, event: "ReserveFractionSet", logs: logs, sub: sub}, nil
}

// WatchReserveFractionSet is a free log subscription operation binding the contract event 0xb690f84efb1d9039c2834effb7bebc792a85bfec7ef84f4b269528454f363ccf.
//
// Solidity: event ReserveFractionSet(uint256 reserveFraction)
func (_Exchange *ExchangeFilterer) WatchReserveFractionSet(opts *bind.WatchOpts, sink chan<- *ExchangeReserveFractionSet) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "ReserveFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeReserveFractionSet)
				if err := _Exchange.contract.UnpackLog(event, "ReserveFractionSet", log); err != nil {
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

// ParseReserveFractionSet is a log parse operation binding the contract event 0xb690f84efb1d9039c2834effb7bebc792a85bfec7ef84f4b269528454f363ccf.
//
// Solidity: event ReserveFractionSet(uint256 reserveFraction)
func (_Exchange *ExchangeFilterer) ParseReserveFractionSet(log types.Log) (*ExchangeReserveFractionSet, error) {
	event := new(ExchangeReserveFractionSet)
	if err := _Exchange.contract.UnpackLog(event, "ReserveFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeSpreadSetIterator is returned from FilterSpreadSet and is used to iterate over the raw logs and unpacked data for SpreadSet events raised by the Exchange contract.
type ExchangeSpreadSetIterator struct {
	Event *ExchangeSpreadSet // Event containing the contract specifics and raw log

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
func (it *ExchangeSpreadSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeSpreadSet)
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
		it.Event = new(ExchangeSpreadSet)
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
func (it *ExchangeSpreadSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeSpreadSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeSpreadSet represents a SpreadSet event raised by the Exchange contract.
type ExchangeSpreadSet struct {
	Spread *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSpreadSet is a free log retrieval operation binding the contract event 0x8946f328efcc515b5cc3282f6cd95e87a6c0d3508421af0b52d4d3620b3e2db3.
//
// Solidity: event SpreadSet(uint256 spread)
func (_Exchange *ExchangeFilterer) FilterSpreadSet(opts *bind.FilterOpts) (*ExchangeSpreadSetIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "SpreadSet")
	if err != nil {
		return nil, err
	}
	return &ExchangeSpreadSetIterator{contract: _Exchange.contract, event: "SpreadSet", logs: logs, sub: sub}, nil
}

// WatchSpreadSet is a free log subscription operation binding the contract event 0x8946f328efcc515b5cc3282f6cd95e87a6c0d3508421af0b52d4d3620b3e2db3.
//
// Solidity: event SpreadSet(uint256 spread)
func (_Exchange *ExchangeFilterer) WatchSpreadSet(opts *bind.WatchOpts, sink chan<- *ExchangeSpreadSet) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "SpreadSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeSpreadSet)
				if err := _Exchange.contract.UnpackLog(event, "SpreadSet", log); err != nil {
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

// ParseSpreadSet is a log parse operation binding the contract event 0x8946f328efcc515b5cc3282f6cd95e87a6c0d3508421af0b52d4d3620b3e2db3.
//
// Solidity: event SpreadSet(uint256 spread)
func (_Exchange *ExchangeFilterer) ParseSpreadSet(log types.Log) (*ExchangeSpreadSet, error) {
	event := new(ExchangeSpreadSet)
	if err := _Exchange.contract.UnpackLog(event, "SpreadSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeStableTokenSetIterator is returned from FilterStableTokenSet and is used to iterate over the raw logs and unpacked data for StableTokenSet events raised by the Exchange contract.
type ExchangeStableTokenSetIterator struct {
	Event *ExchangeStableTokenSet // Event containing the contract specifics and raw log

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
func (it *ExchangeStableTokenSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeStableTokenSet)
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
		it.Event = new(ExchangeStableTokenSet)
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
func (it *ExchangeStableTokenSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeStableTokenSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeStableTokenSet represents a StableTokenSet event raised by the Exchange contract.
type ExchangeStableTokenSet struct {
	Stable common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStableTokenSet is a free log retrieval operation binding the contract event 0x119a23392e161a0bc5f9d5f3e2a6040c45b40d43a36973e10ea1de916f3d8a8a.
//
// Solidity: event StableTokenSet(address indexed stable)
func (_Exchange *ExchangeFilterer) FilterStableTokenSet(opts *bind.FilterOpts, stable []common.Address) (*ExchangeStableTokenSetIterator, error) {

	var stableRule []interface{}
	for _, stableItem := range stable {
		stableRule = append(stableRule, stableItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "StableTokenSet", stableRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeStableTokenSetIterator{contract: _Exchange.contract, event: "StableTokenSet", logs: logs, sub: sub}, nil
}

// WatchStableTokenSet is a free log subscription operation binding the contract event 0x119a23392e161a0bc5f9d5f3e2a6040c45b40d43a36973e10ea1de916f3d8a8a.
//
// Solidity: event StableTokenSet(address indexed stable)
func (_Exchange *ExchangeFilterer) WatchStableTokenSet(opts *bind.WatchOpts, sink chan<- *ExchangeStableTokenSet, stable []common.Address) (event.Subscription, error) {

	var stableRule []interface{}
	for _, stableItem := range stable {
		stableRule = append(stableRule, stableItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "StableTokenSet", stableRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeStableTokenSet)
				if err := _Exchange.contract.UnpackLog(event, "StableTokenSet", log); err != nil {
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

// ParseStableTokenSet is a log parse operation binding the contract event 0x119a23392e161a0bc5f9d5f3e2a6040c45b40d43a36973e10ea1de916f3d8a8a.
//
// Solidity: event StableTokenSet(address indexed stable)
func (_Exchange *ExchangeFilterer) ParseStableTokenSet(log types.Log) (*ExchangeStableTokenSet, error) {
	event := new(ExchangeStableTokenSet)
	if err := _Exchange.contract.UnpackLog(event, "StableTokenSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExchangeUpdateFrequencySetIterator is returned from FilterUpdateFrequencySet and is used to iterate over the raw logs and unpacked data for UpdateFrequencySet events raised by the Exchange contract.
type ExchangeUpdateFrequencySetIterator struct {
	Event *ExchangeUpdateFrequencySet // Event containing the contract specifics and raw log

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
func (it *ExchangeUpdateFrequencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeUpdateFrequencySet)
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
		it.Event = new(ExchangeUpdateFrequencySet)
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
func (it *ExchangeUpdateFrequencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeUpdateFrequencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeUpdateFrequencySet represents a UpdateFrequencySet event raised by the Exchange contract.
type ExchangeUpdateFrequencySet struct {
	UpdateFrequency *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUpdateFrequencySet is a free log retrieval operation binding the contract event 0x90c0a4a142fbfbc2ae8c21f50729a2f4bc56e85a66c1a1b6654f1e85092a54a6.
//
// Solidity: event UpdateFrequencySet(uint256 updateFrequency)
func (_Exchange *ExchangeFilterer) FilterUpdateFrequencySet(opts *bind.FilterOpts) (*ExchangeUpdateFrequencySetIterator, error) {

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "UpdateFrequencySet")
	if err != nil {
		return nil, err
	}
	return &ExchangeUpdateFrequencySetIterator{contract: _Exchange.contract, event: "UpdateFrequencySet", logs: logs, sub: sub}, nil
}

// WatchUpdateFrequencySet is a free log subscription operation binding the contract event 0x90c0a4a142fbfbc2ae8c21f50729a2f4bc56e85a66c1a1b6654f1e85092a54a6.
//
// Solidity: event UpdateFrequencySet(uint256 updateFrequency)
func (_Exchange *ExchangeFilterer) WatchUpdateFrequencySet(opts *bind.WatchOpts, sink chan<- *ExchangeUpdateFrequencySet) (event.Subscription, error) {

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "UpdateFrequencySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeUpdateFrequencySet)
				if err := _Exchange.contract.UnpackLog(event, "UpdateFrequencySet", log); err != nil {
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

// ParseUpdateFrequencySet is a log parse operation binding the contract event 0x90c0a4a142fbfbc2ae8c21f50729a2f4bc56e85a66c1a1b6654f1e85092a54a6.
//
// Solidity: event UpdateFrequencySet(uint256 updateFrequency)
func (_Exchange *ExchangeFilterer) ParseUpdateFrequencySet(log types.Log) (*ExchangeUpdateFrequencySet, error) {
	event := new(ExchangeUpdateFrequencySet)
	if err := _Exchange.contract.UnpackLog(event, "UpdateFrequencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
