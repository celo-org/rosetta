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

// GasPriceMinimumABI is the input ABI used to generate the binding from.
const GasPriceMinimumABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasPriceMinimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetDensity\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"adjustmentSpeed\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gasPriceMinimumFloor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"targetDensity\",\"type\":\"uint256\"}],\"name\":\"TargetDensitySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"gasPriceMinimumFloor\",\"type\":\"uint256\"}],\"name\":\"GasPriceMinimumFloorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"AdjustmentSpeedSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_registryAddress\",\"type\":\"address\"},{\"name\":\"_gasPriceMinimumFloor\",\"type\":\"uint256\"},{\"name\":\"_targetDensity\",\"type\":\"uint256\"},{\"name\":\"_adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"setAdjustmentSpeed\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_targetDensity\",\"type\":\"uint256\"}],\"name\":\"setTargetDensity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_gasPriceMinimumFloor\",\"type\":\"uint256\"}],\"name\":\"setGasPriceMinimumFloor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"getGasPriceMinimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blockGasTotal\",\"type\":\"uint256\"},{\"name\":\"blockGasLimit\",\"type\":\"uint256\"}],\"name\":\"updateGasPriceMinimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockGasTotal\",\"type\":\"uint256\"},{\"name\":\"blockGasLimit\",\"type\":\"uint256\"}],\"name\":\"getUpdatedGasPriceMinimum\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// GasPriceMinimum is an auto generated Go binding around an Ethereum contract.
type GasPriceMinimum struct {
	GasPriceMinimumCaller     // Read-only binding to the contract
	GasPriceMinimumTransactor // Write-only binding to the contract
	GasPriceMinimumFilterer   // Log filterer for contract events
}

// GasPriceMinimumCaller is an auto generated read-only Go binding around an Ethereum contract.
type GasPriceMinimumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceMinimumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GasPriceMinimumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceMinimumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GasPriceMinimumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GasPriceMinimumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GasPriceMinimumSession struct {
	Contract     *GasPriceMinimum  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GasPriceMinimumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GasPriceMinimumCallerSession struct {
	Contract *GasPriceMinimumCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// GasPriceMinimumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GasPriceMinimumTransactorSession struct {
	Contract     *GasPriceMinimumTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// GasPriceMinimumRaw is an auto generated low-level Go binding around an Ethereum contract.
type GasPriceMinimumRaw struct {
	Contract *GasPriceMinimum // Generic contract binding to access the raw methods on
}

// GasPriceMinimumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GasPriceMinimumCallerRaw struct {
	Contract *GasPriceMinimumCaller // Generic read-only contract binding to access the raw methods on
}

// GasPriceMinimumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GasPriceMinimumTransactorRaw struct {
	Contract *GasPriceMinimumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGasPriceMinimum creates a new instance of GasPriceMinimum, bound to a specific deployed contract.
func NewGasPriceMinimum(address common.Address, backend bind.ContractBackend) (*GasPriceMinimum, error) {
	contract, err := bindGasPriceMinimum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimum{GasPriceMinimumCaller: GasPriceMinimumCaller{contract: contract}, GasPriceMinimumTransactor: GasPriceMinimumTransactor{contract: contract}, GasPriceMinimumFilterer: GasPriceMinimumFilterer{contract: contract}}, nil
}

// NewGasPriceMinimumCaller creates a new read-only instance of GasPriceMinimum, bound to a specific deployed contract.
func NewGasPriceMinimumCaller(address common.Address, caller bind.ContractCaller) (*GasPriceMinimumCaller, error) {
	contract, err := bindGasPriceMinimum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumCaller{contract: contract}, nil
}

// NewGasPriceMinimumTransactor creates a new write-only instance of GasPriceMinimum, bound to a specific deployed contract.
func NewGasPriceMinimumTransactor(address common.Address, transactor bind.ContractTransactor) (*GasPriceMinimumTransactor, error) {
	contract, err := bindGasPriceMinimum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumTransactor{contract: contract}, nil
}

// NewGasPriceMinimumFilterer creates a new log filterer instance of GasPriceMinimum, bound to a specific deployed contract.
func NewGasPriceMinimumFilterer(address common.Address, filterer bind.ContractFilterer) (*GasPriceMinimumFilterer, error) {
	contract, err := bindGasPriceMinimum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumFilterer{contract: contract}, nil
}

// bindGasPriceMinimum binds a generic wrapper to an already deployed contract.
func bindGasPriceMinimum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GasPriceMinimumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseGasPriceMinimumABI parses the ABI
func ParseGasPriceMinimumABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(GasPriceMinimumABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPriceMinimum *GasPriceMinimumRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasPriceMinimum.Contract.GasPriceMinimumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPriceMinimum *GasPriceMinimumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.GasPriceMinimumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPriceMinimum *GasPriceMinimumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.GasPriceMinimumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GasPriceMinimum *GasPriceMinimumCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _GasPriceMinimum.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GasPriceMinimum *GasPriceMinimumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GasPriceMinimum *GasPriceMinimumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.contract.Transact(opts, method, params...)
}

// AdjustmentSpeed is a free data retrieval call binding the contract method 0xa68f548e.
//
// Solidity: function adjustmentSpeed() constant returns(uint256 value)
func (_GasPriceMinimum *GasPriceMinimumCaller) AdjustmentSpeed(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "adjustmentSpeed")
	return *ret0, err
}

// AdjustmentSpeed is a free data retrieval call binding the contract method 0xa68f548e.
//
// Solidity: function adjustmentSpeed() constant returns(uint256 value)
func (_GasPriceMinimum *GasPriceMinimumSession) AdjustmentSpeed() (*big.Int, error) {
	return _GasPriceMinimum.Contract.AdjustmentSpeed(&_GasPriceMinimum.CallOpts)
}

// AdjustmentSpeed is a free data retrieval call binding the contract method 0xa68f548e.
//
// Solidity: function adjustmentSpeed() constant returns(uint256 value)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) AdjustmentSpeed() (*big.Int, error) {
	return _GasPriceMinimum.Contract.AdjustmentSpeed(&_GasPriceMinimum.CallOpts)
}

// GasPriceMinimum is a free data retrieval call binding the contract method 0x36945c2d.
//
// Solidity: function gasPriceMinimum() constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCaller) GasPriceMinimum(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "gasPriceMinimum")
	return *ret0, err
}

// GasPriceMinimum is a free data retrieval call binding the contract method 0x36945c2d.
//
// Solidity: function gasPriceMinimum() constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumSession) GasPriceMinimum() (*big.Int, error) {
	return _GasPriceMinimum.Contract.GasPriceMinimum(&_GasPriceMinimum.CallOpts)
}

// GasPriceMinimum is a free data retrieval call binding the contract method 0x36945c2d.
//
// Solidity: function gasPriceMinimum() constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) GasPriceMinimum() (*big.Int, error) {
	return _GasPriceMinimum.Contract.GasPriceMinimum(&_GasPriceMinimum.CallOpts)
}

// GasPriceMinimumFloor is a free data retrieval call binding the contract method 0xceff0bd6.
//
// Solidity: function gasPriceMinimumFloor() constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCaller) GasPriceMinimumFloor(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "gasPriceMinimumFloor")
	return *ret0, err
}

// GasPriceMinimumFloor is a free data retrieval call binding the contract method 0xceff0bd6.
//
// Solidity: function gasPriceMinimumFloor() constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumSession) GasPriceMinimumFloor() (*big.Int, error) {
	return _GasPriceMinimum.Contract.GasPriceMinimumFloor(&_GasPriceMinimum.CallOpts)
}

// GasPriceMinimumFloor is a free data retrieval call binding the contract method 0xceff0bd6.
//
// Solidity: function gasPriceMinimumFloor() constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) GasPriceMinimumFloor() (*big.Int, error) {
	return _GasPriceMinimum.Contract.GasPriceMinimumFloor(&_GasPriceMinimum.CallOpts)
}

// GetGasPriceMinimum is a free data retrieval call binding the contract method 0xa54b7fc0.
//
// Solidity: function getGasPriceMinimum(address tokenAddress) constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCaller) GetGasPriceMinimum(opts *bind.CallOpts, tokenAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "getGasPriceMinimum", tokenAddress)
	return *ret0, err
}

// GetGasPriceMinimum is a free data retrieval call binding the contract method 0xa54b7fc0.
//
// Solidity: function getGasPriceMinimum(address tokenAddress) constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumSession) GetGasPriceMinimum(tokenAddress common.Address) (*big.Int, error) {
	return _GasPriceMinimum.Contract.GetGasPriceMinimum(&_GasPriceMinimum.CallOpts, tokenAddress)
}

// GetGasPriceMinimum is a free data retrieval call binding the contract method 0xa54b7fc0.
//
// Solidity: function getGasPriceMinimum(address tokenAddress) constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) GetGasPriceMinimum(tokenAddress common.Address) (*big.Int, error) {
	return _GasPriceMinimum.Contract.GetGasPriceMinimum(&_GasPriceMinimum.CallOpts, tokenAddress)
}

// GetUpdatedGasPriceMinimum is a free data retrieval call binding the contract method 0xef712c5b.
//
// Solidity: function getUpdatedGasPriceMinimum(uint256 blockGasTotal, uint256 blockGasLimit) constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCaller) GetUpdatedGasPriceMinimum(opts *bind.CallOpts, blockGasTotal *big.Int, blockGasLimit *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "getUpdatedGasPriceMinimum", blockGasTotal, blockGasLimit)
	return *ret0, err
}

// GetUpdatedGasPriceMinimum is a free data retrieval call binding the contract method 0xef712c5b.
//
// Solidity: function getUpdatedGasPriceMinimum(uint256 blockGasTotal, uint256 blockGasLimit) constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumSession) GetUpdatedGasPriceMinimum(blockGasTotal *big.Int, blockGasLimit *big.Int) (*big.Int, error) {
	return _GasPriceMinimum.Contract.GetUpdatedGasPriceMinimum(&_GasPriceMinimum.CallOpts, blockGasTotal, blockGasLimit)
}

// GetUpdatedGasPriceMinimum is a free data retrieval call binding the contract method 0xef712c5b.
//
// Solidity: function getUpdatedGasPriceMinimum(uint256 blockGasTotal, uint256 blockGasLimit) constant returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) GetUpdatedGasPriceMinimum(blockGasTotal *big.Int, blockGasLimit *big.Int) (*big.Int, error) {
	return _GasPriceMinimum.Contract.GetUpdatedGasPriceMinimum(&_GasPriceMinimum.CallOpts, blockGasTotal, blockGasLimit)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GasPriceMinimum *GasPriceMinimumCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GasPriceMinimum *GasPriceMinimumSession) Initialized() (bool, error) {
	return _GasPriceMinimum.Contract.Initialized(&_GasPriceMinimum.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) Initialized() (bool, error) {
	return _GasPriceMinimum.Contract.Initialized(&_GasPriceMinimum.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GasPriceMinimum *GasPriceMinimumCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GasPriceMinimum *GasPriceMinimumSession) IsOwner() (bool, error) {
	return _GasPriceMinimum.Contract.IsOwner(&_GasPriceMinimum.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) IsOwner() (bool, error) {
	return _GasPriceMinimum.Contract.IsOwner(&_GasPriceMinimum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GasPriceMinimum *GasPriceMinimumCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GasPriceMinimum *GasPriceMinimumSession) Owner() (common.Address, error) {
	return _GasPriceMinimum.Contract.Owner(&_GasPriceMinimum.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) Owner() (common.Address, error) {
	return _GasPriceMinimum.Contract.Owner(&_GasPriceMinimum.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GasPriceMinimum *GasPriceMinimumCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GasPriceMinimum *GasPriceMinimumSession) Registry() (common.Address, error) {
	return _GasPriceMinimum.Contract.Registry(&_GasPriceMinimum.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) Registry() (common.Address, error) {
	return _GasPriceMinimum.Contract.Registry(&_GasPriceMinimum.CallOpts)
}

// TargetDensity is a free data retrieval call binding the contract method 0x4a3d5fe2.
//
// Solidity: function targetDensity() constant returns(uint256 value)
func (_GasPriceMinimum *GasPriceMinimumCaller) TargetDensity(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _GasPriceMinimum.contract.Call(opts, out, "targetDensity")
	return *ret0, err
}

// TargetDensity is a free data retrieval call binding the contract method 0x4a3d5fe2.
//
// Solidity: function targetDensity() constant returns(uint256 value)
func (_GasPriceMinimum *GasPriceMinimumSession) TargetDensity() (*big.Int, error) {
	return _GasPriceMinimum.Contract.TargetDensity(&_GasPriceMinimum.CallOpts)
}

// TargetDensity is a free data retrieval call binding the contract method 0x4a3d5fe2.
//
// Solidity: function targetDensity() constant returns(uint256 value)
func (_GasPriceMinimum *GasPriceMinimumCallerSession) TargetDensity() (*big.Int, error) {
	return _GasPriceMinimum.Contract.TargetDensity(&_GasPriceMinimum.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _registryAddress, uint256 _gasPriceMinimumFloor, uint256 _targetDensity, uint256 _adjustmentSpeed) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) Initialize(opts *bind.TransactOpts, _registryAddress common.Address, _gasPriceMinimumFloor *big.Int, _targetDensity *big.Int, _adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "initialize", _registryAddress, _gasPriceMinimumFloor, _targetDensity, _adjustmentSpeed)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _registryAddress, uint256 _gasPriceMinimumFloor, uint256 _targetDensity, uint256 _adjustmentSpeed) returns()
func (_GasPriceMinimum *GasPriceMinimumSession) Initialize(_registryAddress common.Address, _gasPriceMinimumFloor *big.Int, _targetDensity *big.Int, _adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.Initialize(&_GasPriceMinimum.TransactOpts, _registryAddress, _gasPriceMinimumFloor, _targetDensity, _adjustmentSpeed)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address _registryAddress, uint256 _gasPriceMinimumFloor, uint256 _targetDensity, uint256 _adjustmentSpeed) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) Initialize(_registryAddress common.Address, _gasPriceMinimumFloor *big.Int, _targetDensity *big.Int, _adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.Initialize(&_GasPriceMinimum.TransactOpts, _registryAddress, _gasPriceMinimumFloor, _targetDensity, _adjustmentSpeed)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceMinimum *GasPriceMinimumSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.RenounceOwnership(&_GasPriceMinimum.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.RenounceOwnership(&_GasPriceMinimum.TransactOpts)
}

// SetAdjustmentSpeed is a paid mutator transaction binding the contract method 0x30f726b9.
//
// Solidity: function setAdjustmentSpeed(uint256 _adjustmentSpeed) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) SetAdjustmentSpeed(opts *bind.TransactOpts, _adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "setAdjustmentSpeed", _adjustmentSpeed)
}

// SetAdjustmentSpeed is a paid mutator transaction binding the contract method 0x30f726b9.
//
// Solidity: function setAdjustmentSpeed(uint256 _adjustmentSpeed) returns()
func (_GasPriceMinimum *GasPriceMinimumSession) SetAdjustmentSpeed(_adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetAdjustmentSpeed(&_GasPriceMinimum.TransactOpts, _adjustmentSpeed)
}

// SetAdjustmentSpeed is a paid mutator transaction binding the contract method 0x30f726b9.
//
// Solidity: function setAdjustmentSpeed(uint256 _adjustmentSpeed) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) SetAdjustmentSpeed(_adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetAdjustmentSpeed(&_GasPriceMinimum.TransactOpts, _adjustmentSpeed)
}

// SetGasPriceMinimumFloor is a paid mutator transaction binding the contract method 0xb830f4a4.
//
// Solidity: function setGasPriceMinimumFloor(uint256 _gasPriceMinimumFloor) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) SetGasPriceMinimumFloor(opts *bind.TransactOpts, _gasPriceMinimumFloor *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "setGasPriceMinimumFloor", _gasPriceMinimumFloor)
}

// SetGasPriceMinimumFloor is a paid mutator transaction binding the contract method 0xb830f4a4.
//
// Solidity: function setGasPriceMinimumFloor(uint256 _gasPriceMinimumFloor) returns()
func (_GasPriceMinimum *GasPriceMinimumSession) SetGasPriceMinimumFloor(_gasPriceMinimumFloor *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetGasPriceMinimumFloor(&_GasPriceMinimum.TransactOpts, _gasPriceMinimumFloor)
}

// SetGasPriceMinimumFloor is a paid mutator transaction binding the contract method 0xb830f4a4.
//
// Solidity: function setGasPriceMinimumFloor(uint256 _gasPriceMinimumFloor) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) SetGasPriceMinimumFloor(_gasPriceMinimumFloor *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetGasPriceMinimumFloor(&_GasPriceMinimum.TransactOpts, _gasPriceMinimumFloor)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GasPriceMinimum *GasPriceMinimumSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetRegistry(&_GasPriceMinimum.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetRegistry(&_GasPriceMinimum.TransactOpts, registryAddress)
}

// SetTargetDensity is a paid mutator transaction binding the contract method 0x93ca6fc4.
//
// Solidity: function setTargetDensity(uint256 _targetDensity) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) SetTargetDensity(opts *bind.TransactOpts, _targetDensity *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "setTargetDensity", _targetDensity)
}

// SetTargetDensity is a paid mutator transaction binding the contract method 0x93ca6fc4.
//
// Solidity: function setTargetDensity(uint256 _targetDensity) returns()
func (_GasPriceMinimum *GasPriceMinimumSession) SetTargetDensity(_targetDensity *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetTargetDensity(&_GasPriceMinimum.TransactOpts, _targetDensity)
}

// SetTargetDensity is a paid mutator transaction binding the contract method 0x93ca6fc4.
//
// Solidity: function setTargetDensity(uint256 _targetDensity) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) SetTargetDensity(_targetDensity *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.SetTargetDensity(&_GasPriceMinimum.TransactOpts, _targetDensity)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceMinimum *GasPriceMinimumSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.TransferOwnership(&_GasPriceMinimum.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.TransferOwnership(&_GasPriceMinimum.TransactOpts, newOwner)
}

// UpdateGasPriceMinimum is a paid mutator transaction binding the contract method 0xc12398b4.
//
// Solidity: function updateGasPriceMinimum(uint256 blockGasTotal, uint256 blockGasLimit) returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumTransactor) UpdateGasPriceMinimum(opts *bind.TransactOpts, blockGasTotal *big.Int, blockGasLimit *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.contract.Transact(opts, "updateGasPriceMinimum", blockGasTotal, blockGasLimit)
}

// UpdateGasPriceMinimum is a paid mutator transaction binding the contract method 0xc12398b4.
//
// Solidity: function updateGasPriceMinimum(uint256 blockGasTotal, uint256 blockGasLimit) returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumSession) UpdateGasPriceMinimum(blockGasTotal *big.Int, blockGasLimit *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.UpdateGasPriceMinimum(&_GasPriceMinimum.TransactOpts, blockGasTotal, blockGasLimit)
}

// UpdateGasPriceMinimum is a paid mutator transaction binding the contract method 0xc12398b4.
//
// Solidity: function updateGasPriceMinimum(uint256 blockGasTotal, uint256 blockGasLimit) returns(uint256)
func (_GasPriceMinimum *GasPriceMinimumTransactorSession) UpdateGasPriceMinimum(blockGasTotal *big.Int, blockGasLimit *big.Int) (*types.Transaction, error) {
	return _GasPriceMinimum.Contract.UpdateGasPriceMinimum(&_GasPriceMinimum.TransactOpts, blockGasTotal, blockGasLimit)
}

// TryParseLog attemps to parse a log. Returns the parsed log, evenName and wether it was succesfull
func (_GasPriceMinimum *GasPriceMinimumFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _GasPriceMinimum.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AdjustmentSpeedSet":
		event, err = _GasPriceMinimum.ParseAdjustmentSpeedSet(log)
	case "GasPriceMinimumFloorSet":
		event, err = _GasPriceMinimum.ParseGasPriceMinimumFloorSet(log)
	case "OwnershipTransferred":
		event, err = _GasPriceMinimum.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _GasPriceMinimum.ParseRegistrySet(log)
	case "TargetDensitySet":
		event, err = _GasPriceMinimum.ParseTargetDensitySet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// GasPriceMinimumAdjustmentSpeedSetIterator is returned from FilterAdjustmentSpeedSet and is used to iterate over the raw logs and unpacked data for AdjustmentSpeedSet events raised by the GasPriceMinimum contract.
type GasPriceMinimumAdjustmentSpeedSetIterator struct {
	Event *GasPriceMinimumAdjustmentSpeedSet // Event containing the contract specifics and raw log

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
func (it *GasPriceMinimumAdjustmentSpeedSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceMinimumAdjustmentSpeedSet)
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
		it.Event = new(GasPriceMinimumAdjustmentSpeedSet)
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
func (it *GasPriceMinimumAdjustmentSpeedSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceMinimumAdjustmentSpeedSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceMinimumAdjustmentSpeedSet represents a AdjustmentSpeedSet event raised by the GasPriceMinimum contract.
type GasPriceMinimumAdjustmentSpeedSet struct {
	AdjustmentSpeed *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterAdjustmentSpeedSet is a free log retrieval operation binding the contract event 0xd2e71cd7012df1df07d4908ff75ae4b2bfbb6c49d39144404661f1fd47253283.
//
// Solidity: event AdjustmentSpeedSet(uint256 adjustmentSpeed)
func (_GasPriceMinimum *GasPriceMinimumFilterer) FilterAdjustmentSpeedSet(opts *bind.FilterOpts) (*GasPriceMinimumAdjustmentSpeedSetIterator, error) {

	logs, sub, err := _GasPriceMinimum.contract.FilterLogs(opts, "AdjustmentSpeedSet")
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumAdjustmentSpeedSetIterator{contract: _GasPriceMinimum.contract, event: "AdjustmentSpeedSet", logs: logs, sub: sub}, nil
}

// WatchAdjustmentSpeedSet is a free log subscription operation binding the contract event 0xd2e71cd7012df1df07d4908ff75ae4b2bfbb6c49d39144404661f1fd47253283.
//
// Solidity: event AdjustmentSpeedSet(uint256 adjustmentSpeed)
func (_GasPriceMinimum *GasPriceMinimumFilterer) WatchAdjustmentSpeedSet(opts *bind.WatchOpts, sink chan<- *GasPriceMinimumAdjustmentSpeedSet) (event.Subscription, error) {

	logs, sub, err := _GasPriceMinimum.contract.WatchLogs(opts, "AdjustmentSpeedSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceMinimumAdjustmentSpeedSet)
				if err := _GasPriceMinimum.contract.UnpackLog(event, "AdjustmentSpeedSet", log); err != nil {
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

// ParseAdjustmentSpeedSet is a log parse operation binding the contract event 0xd2e71cd7012df1df07d4908ff75ae4b2bfbb6c49d39144404661f1fd47253283.
//
// Solidity: event AdjustmentSpeedSet(uint256 adjustmentSpeed)
func (_GasPriceMinimum *GasPriceMinimumFilterer) ParseAdjustmentSpeedSet(log types.Log) (*GasPriceMinimumAdjustmentSpeedSet, error) {
	event := new(GasPriceMinimumAdjustmentSpeedSet)
	if err := _GasPriceMinimum.contract.UnpackLog(event, "AdjustmentSpeedSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasPriceMinimumGasPriceMinimumFloorSetIterator is returned from FilterGasPriceMinimumFloorSet and is used to iterate over the raw logs and unpacked data for GasPriceMinimumFloorSet events raised by the GasPriceMinimum contract.
type GasPriceMinimumGasPriceMinimumFloorSetIterator struct {
	Event *GasPriceMinimumGasPriceMinimumFloorSet // Event containing the contract specifics and raw log

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
func (it *GasPriceMinimumGasPriceMinimumFloorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceMinimumGasPriceMinimumFloorSet)
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
		it.Event = new(GasPriceMinimumGasPriceMinimumFloorSet)
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
func (it *GasPriceMinimumGasPriceMinimumFloorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceMinimumGasPriceMinimumFloorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceMinimumGasPriceMinimumFloorSet represents a GasPriceMinimumFloorSet event raised by the GasPriceMinimum contract.
type GasPriceMinimumGasPriceMinimumFloorSet struct {
	GasPriceMinimumFloor *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterGasPriceMinimumFloorSet is a free log retrieval operation binding the contract event 0x5548a13ccc1d9e4e2860461edda5ad49ba8a4fda485f67d954f9d7da8d2aff27.
//
// Solidity: event GasPriceMinimumFloorSet(uint256 gasPriceMinimumFloor)
func (_GasPriceMinimum *GasPriceMinimumFilterer) FilterGasPriceMinimumFloorSet(opts *bind.FilterOpts) (*GasPriceMinimumGasPriceMinimumFloorSetIterator, error) {

	logs, sub, err := _GasPriceMinimum.contract.FilterLogs(opts, "GasPriceMinimumFloorSet")
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumGasPriceMinimumFloorSetIterator{contract: _GasPriceMinimum.contract, event: "GasPriceMinimumFloorSet", logs: logs, sub: sub}, nil
}

// WatchGasPriceMinimumFloorSet is a free log subscription operation binding the contract event 0x5548a13ccc1d9e4e2860461edda5ad49ba8a4fda485f67d954f9d7da8d2aff27.
//
// Solidity: event GasPriceMinimumFloorSet(uint256 gasPriceMinimumFloor)
func (_GasPriceMinimum *GasPriceMinimumFilterer) WatchGasPriceMinimumFloorSet(opts *bind.WatchOpts, sink chan<- *GasPriceMinimumGasPriceMinimumFloorSet) (event.Subscription, error) {

	logs, sub, err := _GasPriceMinimum.contract.WatchLogs(opts, "GasPriceMinimumFloorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceMinimumGasPriceMinimumFloorSet)
				if err := _GasPriceMinimum.contract.UnpackLog(event, "GasPriceMinimumFloorSet", log); err != nil {
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

// ParseGasPriceMinimumFloorSet is a log parse operation binding the contract event 0x5548a13ccc1d9e4e2860461edda5ad49ba8a4fda485f67d954f9d7da8d2aff27.
//
// Solidity: event GasPriceMinimumFloorSet(uint256 gasPriceMinimumFloor)
func (_GasPriceMinimum *GasPriceMinimumFilterer) ParseGasPriceMinimumFloorSet(log types.Log) (*GasPriceMinimumGasPriceMinimumFloorSet, error) {
	event := new(GasPriceMinimumGasPriceMinimumFloorSet)
	if err := _GasPriceMinimum.contract.UnpackLog(event, "GasPriceMinimumFloorSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasPriceMinimumOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the GasPriceMinimum contract.
type GasPriceMinimumOwnershipTransferredIterator struct {
	Event *GasPriceMinimumOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GasPriceMinimumOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceMinimumOwnershipTransferred)
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
		it.Event = new(GasPriceMinimumOwnershipTransferred)
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
func (it *GasPriceMinimumOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceMinimumOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceMinimumOwnershipTransferred represents a OwnershipTransferred event raised by the GasPriceMinimum contract.
type GasPriceMinimumOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPriceMinimum *GasPriceMinimumFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GasPriceMinimumOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPriceMinimum.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumOwnershipTransferredIterator{contract: _GasPriceMinimum.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_GasPriceMinimum *GasPriceMinimumFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GasPriceMinimumOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _GasPriceMinimum.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceMinimumOwnershipTransferred)
				if err := _GasPriceMinimum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_GasPriceMinimum *GasPriceMinimumFilterer) ParseOwnershipTransferred(log types.Log) (*GasPriceMinimumOwnershipTransferred, error) {
	event := new(GasPriceMinimumOwnershipTransferred)
	if err := _GasPriceMinimum.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasPriceMinimumRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the GasPriceMinimum contract.
type GasPriceMinimumRegistrySetIterator struct {
	Event *GasPriceMinimumRegistrySet // Event containing the contract specifics and raw log

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
func (it *GasPriceMinimumRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceMinimumRegistrySet)
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
		it.Event = new(GasPriceMinimumRegistrySet)
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
func (it *GasPriceMinimumRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceMinimumRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceMinimumRegistrySet represents a RegistrySet event raised by the GasPriceMinimum contract.
type GasPriceMinimumRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GasPriceMinimum *GasPriceMinimumFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*GasPriceMinimumRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _GasPriceMinimum.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumRegistrySetIterator{contract: _GasPriceMinimum.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_GasPriceMinimum *GasPriceMinimumFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *GasPriceMinimumRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _GasPriceMinimum.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceMinimumRegistrySet)
				if err := _GasPriceMinimum.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_GasPriceMinimum *GasPriceMinimumFilterer) ParseRegistrySet(log types.Log) (*GasPriceMinimumRegistrySet, error) {
	event := new(GasPriceMinimumRegistrySet)
	if err := _GasPriceMinimum.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GasPriceMinimumTargetDensitySetIterator is returned from FilterTargetDensitySet and is used to iterate over the raw logs and unpacked data for TargetDensitySet events raised by the GasPriceMinimum contract.
type GasPriceMinimumTargetDensitySetIterator struct {
	Event *GasPriceMinimumTargetDensitySet // Event containing the contract specifics and raw log

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
func (it *GasPriceMinimumTargetDensitySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GasPriceMinimumTargetDensitySet)
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
		it.Event = new(GasPriceMinimumTargetDensitySet)
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
func (it *GasPriceMinimumTargetDensitySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GasPriceMinimumTargetDensitySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GasPriceMinimumTargetDensitySet represents a TargetDensitySet event raised by the GasPriceMinimum contract.
type GasPriceMinimumTargetDensitySet struct {
	TargetDensity *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTargetDensitySet is a free log retrieval operation binding the contract event 0x2a109bad06121312708ed2a3e9b3556ea85ef8eadd4d10d8181f50d114eb4fab.
//
// Solidity: event TargetDensitySet(uint256 targetDensity)
func (_GasPriceMinimum *GasPriceMinimumFilterer) FilterTargetDensitySet(opts *bind.FilterOpts) (*GasPriceMinimumTargetDensitySetIterator, error) {

	logs, sub, err := _GasPriceMinimum.contract.FilterLogs(opts, "TargetDensitySet")
	if err != nil {
		return nil, err
	}
	return &GasPriceMinimumTargetDensitySetIterator{contract: _GasPriceMinimum.contract, event: "TargetDensitySet", logs: logs, sub: sub}, nil
}

// WatchTargetDensitySet is a free log subscription operation binding the contract event 0x2a109bad06121312708ed2a3e9b3556ea85ef8eadd4d10d8181f50d114eb4fab.
//
// Solidity: event TargetDensitySet(uint256 targetDensity)
func (_GasPriceMinimum *GasPriceMinimumFilterer) WatchTargetDensitySet(opts *bind.WatchOpts, sink chan<- *GasPriceMinimumTargetDensitySet) (event.Subscription, error) {

	logs, sub, err := _GasPriceMinimum.contract.WatchLogs(opts, "TargetDensitySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GasPriceMinimumTargetDensitySet)
				if err := _GasPriceMinimum.contract.UnpackLog(event, "TargetDensitySet", log); err != nil {
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

// ParseTargetDensitySet is a log parse operation binding the contract event 0x2a109bad06121312708ed2a3e9b3556ea85ef8eadd4d10d8181f50d114eb4fab.
//
// Solidity: event TargetDensitySet(uint256 targetDensity)
func (_GasPriceMinimum *GasPriceMinimumFilterer) ParseTargetDensitySet(log types.Log) (*GasPriceMinimumTargetDensitySet, error) {
	event := new(GasPriceMinimumTargetDensitySet)
	if err := _GasPriceMinimum.contract.UnpackLog(event, "TargetDensitySet", log); err != nil {
		return nil, err
	}
	return event, nil
}
