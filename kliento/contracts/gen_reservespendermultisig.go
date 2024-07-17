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

// ReserveSpenderMultiSigMetaData contains all meta data concerning the ReserveSpenderMultiSig contract.
var ReserveSpenderMultiSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"internalRequired\",\"type\":\"uint256\"}],\"name\":\"InternalRequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_internalRequired\",\"type\":\"uint256\"}],\"name\":\"changeInternalRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_internalRequired\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"internalRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReserveSpenderMultiSigABI is the input ABI used to generate the binding from.
// Deprecated: Use ReserveSpenderMultiSigMetaData.ABI instead.
var ReserveSpenderMultiSigABI = ReserveSpenderMultiSigMetaData.ABI

// ReserveSpenderMultiSig is an auto generated Go binding around an Ethereum contract.
type ReserveSpenderMultiSig struct {
	ReserveSpenderMultiSigCaller     // Read-only binding to the contract
	ReserveSpenderMultiSigTransactor // Write-only binding to the contract
	ReserveSpenderMultiSigFilterer   // Log filterer for contract events
}

// ReserveSpenderMultiSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveSpenderMultiSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSpenderMultiSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveSpenderMultiSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSpenderMultiSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveSpenderMultiSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSpenderMultiSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveSpenderMultiSigSession struct {
	Contract     *ReserveSpenderMultiSig // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ReserveSpenderMultiSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveSpenderMultiSigCallerSession struct {
	Contract *ReserveSpenderMultiSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ReserveSpenderMultiSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveSpenderMultiSigTransactorSession struct {
	Contract     *ReserveSpenderMultiSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ReserveSpenderMultiSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveSpenderMultiSigRaw struct {
	Contract *ReserveSpenderMultiSig // Generic contract binding to access the raw methods on
}

// ReserveSpenderMultiSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveSpenderMultiSigCallerRaw struct {
	Contract *ReserveSpenderMultiSigCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveSpenderMultiSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveSpenderMultiSigTransactorRaw struct {
	Contract *ReserveSpenderMultiSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserveSpenderMultiSig creates a new instance of ReserveSpenderMultiSig, bound to a specific deployed contract.
func NewReserveSpenderMultiSig(address common.Address, backend bind.ContractBackend) (*ReserveSpenderMultiSig, error) {
	contract, err := bindReserveSpenderMultiSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSig{ReserveSpenderMultiSigCaller: ReserveSpenderMultiSigCaller{contract: contract}, ReserveSpenderMultiSigTransactor: ReserveSpenderMultiSigTransactor{contract: contract}, ReserveSpenderMultiSigFilterer: ReserveSpenderMultiSigFilterer{contract: contract}}, nil
}

// NewReserveSpenderMultiSigCaller creates a new read-only instance of ReserveSpenderMultiSig, bound to a specific deployed contract.
func NewReserveSpenderMultiSigCaller(address common.Address, caller bind.ContractCaller) (*ReserveSpenderMultiSigCaller, error) {
	contract, err := bindReserveSpenderMultiSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigCaller{contract: contract}, nil
}

// NewReserveSpenderMultiSigTransactor creates a new write-only instance of ReserveSpenderMultiSig, bound to a specific deployed contract.
func NewReserveSpenderMultiSigTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveSpenderMultiSigTransactor, error) {
	contract, err := bindReserveSpenderMultiSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigTransactor{contract: contract}, nil
}

// NewReserveSpenderMultiSigFilterer creates a new log filterer instance of ReserveSpenderMultiSig, bound to a specific deployed contract.
func NewReserveSpenderMultiSigFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveSpenderMultiSigFilterer, error) {
	contract, err := bindReserveSpenderMultiSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigFilterer{contract: contract}, nil
}

// bindReserveSpenderMultiSig binds a generic wrapper to an already deployed contract.
func bindReserveSpenderMultiSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveSpenderMultiSigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseReserveSpenderMultiSigABI parses the ABI
func ParseReserveSpenderMultiSigABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveSpenderMultiSigABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReserveSpenderMultiSig.Contract.ReserveSpenderMultiSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ReserveSpenderMultiSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ReserveSpenderMultiSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReserveSpenderMultiSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "MAX_OWNER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.MAXOWNERCOUNT(&_ReserveSpenderMultiSig.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.MAXOWNERCOUNT(&_ReserveSpenderMultiSig.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ReserveSpenderMultiSig.Contract.Confirmations(&_ReserveSpenderMultiSig.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _ReserveSpenderMultiSig.Contract.Confirmations(&_ReserveSpenderMultiSig.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "getConfirmationCount", transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.GetConfirmationCount(&_ReserveSpenderMultiSig.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.GetConfirmationCount(&_ReserveSpenderMultiSig.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "getConfirmations", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _ReserveSpenderMultiSig.Contract.GetConfirmations(&_ReserveSpenderMultiSig.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _ReserveSpenderMultiSig.Contract.GetConfirmations(&_ReserveSpenderMultiSig.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) GetOwners() ([]common.Address, error) {
	return _ReserveSpenderMultiSig.Contract.GetOwners(&_ReserveSpenderMultiSig.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) GetOwners() ([]common.Address, error) {
	return _ReserveSpenderMultiSig.Contract.GetOwners(&_ReserveSpenderMultiSig.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "getTransactionCount", pending, executed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.GetTransactionCount(&_ReserveSpenderMultiSig.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.GetTransactionCount(&_ReserveSpenderMultiSig.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "getTransactionIds", from, to, pending, executed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.GetTransactionIds(&_ReserveSpenderMultiSig.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.GetTransactionIds(&_ReserveSpenderMultiSig.CallOpts, from, to, pending, executed)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Initialized() (bool, error) {
	return _ReserveSpenderMultiSig.Contract.Initialized(&_ReserveSpenderMultiSig.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) Initialized() (bool, error) {
	return _ReserveSpenderMultiSig.Contract.Initialized(&_ReserveSpenderMultiSig.CallOpts)
}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) InternalRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "internalRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) InternalRequired() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.InternalRequired(&_ReserveSpenderMultiSig.CallOpts)
}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) InternalRequired() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.InternalRequired(&_ReserveSpenderMultiSig.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _ReserveSpenderMultiSig.Contract.IsConfirmed(&_ReserveSpenderMultiSig.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _ReserveSpenderMultiSig.Contract.IsConfirmed(&_ReserveSpenderMultiSig.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) IsOwner(arg0 common.Address) (bool, error) {
	return _ReserveSpenderMultiSig.Contract.IsOwner(&_ReserveSpenderMultiSig.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _ReserveSpenderMultiSig.Contract.IsOwner(&_ReserveSpenderMultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _ReserveSpenderMultiSig.Contract.Owners(&_ReserveSpenderMultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _ReserveSpenderMultiSig.Contract.Owners(&_ReserveSpenderMultiSig.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Required() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.Required(&_ReserveSpenderMultiSig.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) Required() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.Required(&_ReserveSpenderMultiSig.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) TransactionCount() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.TransactionCount(&_ReserveSpenderMultiSig.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) TransactionCount() (*big.Int, error) {
	return _ReserveSpenderMultiSig.Contract.TransactionCount(&_ReserveSpenderMultiSig.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	var out []interface{}
	err := _ReserveSpenderMultiSig.contract.Call(opts, &out, "transactions", arg0)

	outstruct := new(struct {
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Destination = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Data = *abi.ConvertType(out[2], new([]byte)).(*[]byte)
	outstruct.Executed = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _ReserveSpenderMultiSig.Contract.Transactions(&_ReserveSpenderMultiSig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _ReserveSpenderMultiSig.Contract.Transactions(&_ReserveSpenderMultiSig.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.AddOwner(&_ReserveSpenderMultiSig.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.AddOwner(&_ReserveSpenderMultiSig.TransactOpts, owner)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) ChangeInternalRequirement(opts *bind.TransactOpts, _internalRequired *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "changeInternalRequirement", _internalRequired)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) ChangeInternalRequirement(_internalRequired *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ChangeInternalRequirement(&_ReserveSpenderMultiSig.TransactOpts, _internalRequired)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) ChangeInternalRequirement(_internalRequired *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ChangeInternalRequirement(&_ReserveSpenderMultiSig.TransactOpts, _internalRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ChangeRequirement(&_ReserveSpenderMultiSig.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ChangeRequirement(&_ReserveSpenderMultiSig.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ConfirmTransaction(&_ReserveSpenderMultiSig.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ConfirmTransaction(&_ReserveSpenderMultiSig.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ExecuteTransaction(&_ReserveSpenderMultiSig.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ExecuteTransaction(&_ReserveSpenderMultiSig.TransactOpts, transactionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) Initialize(opts *bind.TransactOpts, _owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "initialize", _owners, _required, _internalRequired)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Initialize(_owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.Initialize(&_ReserveSpenderMultiSig.TransactOpts, _owners, _required, _internalRequired)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) Initialize(_owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.Initialize(&_ReserveSpenderMultiSig.TransactOpts, _owners, _required, _internalRequired)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.RemoveOwner(&_ReserveSpenderMultiSig.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.RemoveOwner(&_ReserveSpenderMultiSig.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ReplaceOwner(&_ReserveSpenderMultiSig.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.ReplaceOwner(&_ReserveSpenderMultiSig.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.RevokeConfirmation(&_ReserveSpenderMultiSig.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.RevokeConfirmation(&_ReserveSpenderMultiSig.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.SubmitTransaction(&_ReserveSpenderMultiSig.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.SubmitTransaction(&_ReserveSpenderMultiSig.TransactOpts, destination, value, data)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _ReserveSpenderMultiSig.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "Confirmation":
		event, err = _ReserveSpenderMultiSig.ParseConfirmation(log)
	case "Deposit":
		event, err = _ReserveSpenderMultiSig.ParseDeposit(log)
	case "Execution":
		event, err = _ReserveSpenderMultiSig.ParseExecution(log)
	case "InternalRequirementChange":
		event, err = _ReserveSpenderMultiSig.ParseInternalRequirementChange(log)
	case "OwnerAddition":
		event, err = _ReserveSpenderMultiSig.ParseOwnerAddition(log)
	case "OwnerRemoval":
		event, err = _ReserveSpenderMultiSig.ParseOwnerRemoval(log)
	case "RequirementChange":
		event, err = _ReserveSpenderMultiSig.ParseRequirementChange(log)
	case "Revocation":
		event, err = _ReserveSpenderMultiSig.ParseRevocation(log)
	case "Submission":
		event, err = _ReserveSpenderMultiSig.ParseSubmission(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.Fallback(&_ReserveSpenderMultiSig.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _ReserveSpenderMultiSig.Contract.Fallback(&_ReserveSpenderMultiSig.TransactOpts, calldata)
}

// ReserveSpenderMultiSigConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigConfirmationIterator struct {
	Event *ReserveSpenderMultiSigConfirmation // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigConfirmation)
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
		it.Event = new(ReserveSpenderMultiSigConfirmation)
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
func (it *ReserveSpenderMultiSigConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigConfirmation represents a Confirmation event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*ReserveSpenderMultiSigConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigConfirmationIterator{contract: _ReserveSpenderMultiSig.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigConfirmation)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Confirmation", log); err != nil {
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

// ParseConfirmation is a log parse operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseConfirmation(log types.Log) (*ReserveSpenderMultiSigConfirmation, error) {
	event := new(ReserveSpenderMultiSigConfirmation)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigDepositIterator struct {
	Event *ReserveSpenderMultiSigDeposit // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigDeposit)
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
		it.Event = new(ReserveSpenderMultiSigDeposit)
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
func (it *ReserveSpenderMultiSigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigDeposit represents a Deposit event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*ReserveSpenderMultiSigDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigDepositIterator{contract: _ReserveSpenderMultiSig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigDeposit)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseDeposit(log types.Log) (*ReserveSpenderMultiSigDeposit, error) {
	event := new(ReserveSpenderMultiSigDeposit)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigExecutionIterator struct {
	Event *ReserveSpenderMultiSigExecution // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigExecution)
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
		it.Event = new(ReserveSpenderMultiSigExecution)
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
func (it *ReserveSpenderMultiSigExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigExecution represents a Execution event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigExecution struct {
	TransactionId *big.Int
	ReturnData    []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*ReserveSpenderMultiSigExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigExecutionIterator{contract: _ReserveSpenderMultiSig.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigExecution)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Execution", log); err != nil {
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

// ParseExecution is a log parse operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseExecution(log types.Log) (*ReserveSpenderMultiSigExecution, error) {
	event := new(ReserveSpenderMultiSigExecution)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigInternalRequirementChangeIterator is returned from FilterInternalRequirementChange and is used to iterate over the raw logs and unpacked data for InternalRequirementChange events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigInternalRequirementChangeIterator struct {
	Event *ReserveSpenderMultiSigInternalRequirementChange // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigInternalRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigInternalRequirementChange)
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
		it.Event = new(ReserveSpenderMultiSigInternalRequirementChange)
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
func (it *ReserveSpenderMultiSigInternalRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigInternalRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigInternalRequirementChange represents a InternalRequirementChange event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigInternalRequirementChange struct {
	InternalRequired *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInternalRequirementChange is a free log retrieval operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterInternalRequirementChange(opts *bind.FilterOpts) (*ReserveSpenderMultiSigInternalRequirementChangeIterator, error) {

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "InternalRequirementChange")
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigInternalRequirementChangeIterator{contract: _ReserveSpenderMultiSig.contract, event: "InternalRequirementChange", logs: logs, sub: sub}, nil
}

// WatchInternalRequirementChange is a free log subscription operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchInternalRequirementChange(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigInternalRequirementChange) (event.Subscription, error) {

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "InternalRequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigInternalRequirementChange)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "InternalRequirementChange", log); err != nil {
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

// ParseInternalRequirementChange is a log parse operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseInternalRequirementChange(log types.Log) (*ReserveSpenderMultiSigInternalRequirementChange, error) {
	event := new(ReserveSpenderMultiSigInternalRequirementChange)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "InternalRequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigOwnerAdditionIterator struct {
	Event *ReserveSpenderMultiSigOwnerAddition // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigOwnerAddition)
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
		it.Event = new(ReserveSpenderMultiSigOwnerAddition)
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
func (it *ReserveSpenderMultiSigOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigOwnerAddition represents a OwnerAddition event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*ReserveSpenderMultiSigOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigOwnerAdditionIterator{contract: _ReserveSpenderMultiSig.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigOwnerAddition)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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

// ParseOwnerAddition is a log parse operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseOwnerAddition(log types.Log) (*ReserveSpenderMultiSigOwnerAddition, error) {
	event := new(ReserveSpenderMultiSigOwnerAddition)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigOwnerRemovalIterator struct {
	Event *ReserveSpenderMultiSigOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigOwnerRemoval)
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
		it.Event = new(ReserveSpenderMultiSigOwnerRemoval)
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
func (it *ReserveSpenderMultiSigOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigOwnerRemoval represents a OwnerRemoval event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*ReserveSpenderMultiSigOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigOwnerRemovalIterator{contract: _ReserveSpenderMultiSig.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigOwnerRemoval)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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

// ParseOwnerRemoval is a log parse operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseOwnerRemoval(log types.Log) (*ReserveSpenderMultiSigOwnerRemoval, error) {
	event := new(ReserveSpenderMultiSigOwnerRemoval)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigRequirementChangeIterator struct {
	Event *ReserveSpenderMultiSigRequirementChange // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigRequirementChange)
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
		it.Event = new(ReserveSpenderMultiSigRequirementChange)
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
func (it *ReserveSpenderMultiSigRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigRequirementChange represents a RequirementChange event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*ReserveSpenderMultiSigRequirementChangeIterator, error) {

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigRequirementChangeIterator{contract: _ReserveSpenderMultiSig.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigRequirementChange) (event.Subscription, error) {

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigRequirementChange)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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

// ParseRequirementChange is a log parse operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseRequirementChange(log types.Log) (*ReserveSpenderMultiSigRequirementChange, error) {
	event := new(ReserveSpenderMultiSigRequirementChange)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigRevocationIterator struct {
	Event *ReserveSpenderMultiSigRevocation // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigRevocation)
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
		it.Event = new(ReserveSpenderMultiSigRevocation)
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
func (it *ReserveSpenderMultiSigRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigRevocation represents a Revocation event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*ReserveSpenderMultiSigRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigRevocationIterator{contract: _ReserveSpenderMultiSig.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigRevocation)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Revocation", log); err != nil {
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

// ParseRevocation is a log parse operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseRevocation(log types.Log) (*ReserveSpenderMultiSigRevocation, error) {
	event := new(ReserveSpenderMultiSigRevocation)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderMultiSigSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigSubmissionIterator struct {
	Event *ReserveSpenderMultiSigSubmission // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderMultiSigSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderMultiSigSubmission)
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
		it.Event = new(ReserveSpenderMultiSigSubmission)
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
func (it *ReserveSpenderMultiSigSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderMultiSigSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderMultiSigSubmission represents a Submission event raised by the ReserveSpenderMultiSig contract.
type ReserveSpenderMultiSigSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*ReserveSpenderMultiSigSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderMultiSigSubmissionIterator{contract: _ReserveSpenderMultiSig.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *ReserveSpenderMultiSigSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _ReserveSpenderMultiSig.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderMultiSigSubmission)
				if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Submission", log); err != nil {
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

// ParseSubmission is a log parse operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_ReserveSpenderMultiSig *ReserveSpenderMultiSigFilterer) ParseSubmission(log types.Log) (*ReserveSpenderMultiSigSubmission, error) {
	event := new(ReserveSpenderMultiSigSubmission)
	if err := _ReserveSpenderMultiSig.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
