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

// MultiSigMetaData contains all meta data concerning the MultiSig contract.
var MultiSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"internalRequired\",\"type\":\"uint256\"}],\"name\":\"InternalRequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"internalRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_internalRequired\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_internalRequired\",\"type\":\"uint256\"}],\"name\":\"changeInternalRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MultiSigABI is the input ABI used to generate the binding from.
// Deprecated: Use MultiSigMetaData.ABI instead.
var MultiSigABI = MultiSigMetaData.ABI

// MultiSig is an auto generated Go binding around an Ethereum contract.
type MultiSig struct {
	MultiSigCaller     // Read-only binding to the contract
	MultiSigTransactor // Write-only binding to the contract
	MultiSigFilterer   // Log filterer for contract events
}

// MultiSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type MultiSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MultiSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MultiSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MultiSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MultiSigSession struct {
	Contract     *MultiSig         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MultiSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MultiSigCallerSession struct {
	Contract *MultiSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MultiSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MultiSigTransactorSession struct {
	Contract     *MultiSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MultiSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type MultiSigRaw struct {
	Contract *MultiSig // Generic contract binding to access the raw methods on
}

// MultiSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MultiSigCallerRaw struct {
	Contract *MultiSigCaller // Generic read-only contract binding to access the raw methods on
}

// MultiSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MultiSigTransactorRaw struct {
	Contract *MultiSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMultiSig creates a new instance of MultiSig, bound to a specific deployed contract.
func NewMultiSig(address common.Address, backend bind.ContractBackend) (*MultiSig, error) {
	contract, err := bindMultiSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MultiSig{MultiSigCaller: MultiSigCaller{contract: contract}, MultiSigTransactor: MultiSigTransactor{contract: contract}, MultiSigFilterer: MultiSigFilterer{contract: contract}}, nil
}

// NewMultiSigCaller creates a new read-only instance of MultiSig, bound to a specific deployed contract.
func NewMultiSigCaller(address common.Address, caller bind.ContractCaller) (*MultiSigCaller, error) {
	contract, err := bindMultiSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigCaller{contract: contract}, nil
}

// NewMultiSigTransactor creates a new write-only instance of MultiSig, bound to a specific deployed contract.
func NewMultiSigTransactor(address common.Address, transactor bind.ContractTransactor) (*MultiSigTransactor, error) {
	contract, err := bindMultiSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MultiSigTransactor{contract: contract}, nil
}

// NewMultiSigFilterer creates a new log filterer instance of MultiSig, bound to a specific deployed contract.
func NewMultiSigFilterer(address common.Address, filterer bind.ContractFilterer) (*MultiSigFilterer, error) {
	contract, err := bindMultiSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MultiSigFilterer{contract: contract}, nil
}

// bindMultiSig binds a generic wrapper to an already deployed contract.
func bindMultiSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseMultiSigABI parses the ABI
func ParseMultiSigABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(MultiSigABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSig *MultiSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSig.Contract.MultiSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSig *MultiSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSig.Contract.MultiSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSig *MultiSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSig.Contract.MultiSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MultiSig *MultiSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MultiSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MultiSig *MultiSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MultiSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MultiSig *MultiSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MultiSig.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSig *MultiSigCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "MAX_OWNER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSig *MultiSigSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSig.Contract.MAXOWNERCOUNT(&_MultiSig.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_MultiSig *MultiSigCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MultiSig.Contract.MAXOWNERCOUNT(&_MultiSig.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSig *MultiSigCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSig *MultiSigSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSig.Contract.Confirmations(&_MultiSig.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_MultiSig *MultiSigCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MultiSig.Contract.Confirmations(&_MultiSig.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSig *MultiSigCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getConfirmationCount", transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSig *MultiSigSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSig.Contract.GetConfirmationCount(&_MultiSig.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_MultiSig *MultiSigCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MultiSig.Contract.GetConfirmationCount(&_MultiSig.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSig *MultiSigCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getConfirmations", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSig *MultiSigSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSig.Contract.GetConfirmations(&_MultiSig.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_MultiSig *MultiSigCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MultiSig.Contract.GetConfirmations(&_MultiSig.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSig *MultiSigCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSig *MultiSigSession) GetOwners() ([]common.Address, error) {
	return _MultiSig.Contract.GetOwners(&_MultiSig.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_MultiSig *MultiSigCallerSession) GetOwners() ([]common.Address, error) {
	return _MultiSig.Contract.GetOwners(&_MultiSig.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSig *MultiSigCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getTransactionCount", pending, executed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSig *MultiSigSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSig.Contract.GetTransactionCount(&_MultiSig.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_MultiSig *MultiSigCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MultiSig.Contract.GetTransactionCount(&_MultiSig.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSig *MultiSigCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "getTransactionIds", from, to, pending, executed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSig *MultiSigSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSig.Contract.GetTransactionIds(&_MultiSig.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_MultiSig *MultiSigCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MultiSig.Contract.GetTransactionIds(&_MultiSig.CallOpts, from, to, pending, executed)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MultiSig *MultiSigCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MultiSig *MultiSigSession) Initialized() (bool, error) {
	return _MultiSig.Contract.Initialized(&_MultiSig.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_MultiSig *MultiSigCallerSession) Initialized() (bool, error) {
	return _MultiSig.Contract.Initialized(&_MultiSig.CallOpts)
}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_MultiSig *MultiSigCaller) InternalRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "internalRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_MultiSig *MultiSigSession) InternalRequired() (*big.Int, error) {
	return _MultiSig.Contract.InternalRequired(&_MultiSig.CallOpts)
}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_MultiSig *MultiSigCallerSession) InternalRequired() (*big.Int, error) {
	return _MultiSig.Contract.InternalRequired(&_MultiSig.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSig *MultiSigCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSig *MultiSigSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSig.Contract.IsConfirmed(&_MultiSig.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_MultiSig *MultiSigCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MultiSig.Contract.IsConfirmed(&_MultiSig.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSig *MultiSigCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSig *MultiSigSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSig.Contract.IsOwner(&_MultiSig.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_MultiSig *MultiSigCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MultiSig.Contract.IsOwner(&_MultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSig *MultiSigCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSig *MultiSigSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSig.Contract.Owners(&_MultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_MultiSig *MultiSigCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MultiSig.Contract.Owners(&_MultiSig.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSig *MultiSigCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSig *MultiSigSession) Required() (*big.Int, error) {
	return _MultiSig.Contract.Required(&_MultiSig.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_MultiSig *MultiSigCallerSession) Required() (*big.Int, error) {
	return _MultiSig.Contract.Required(&_MultiSig.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSig *MultiSigCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSig *MultiSigSession) TransactionCount() (*big.Int, error) {
	return _MultiSig.Contract.TransactionCount(&_MultiSig.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_MultiSig *MultiSigCallerSession) TransactionCount() (*big.Int, error) {
	return _MultiSig.Contract.TransactionCount(&_MultiSig.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSig *MultiSigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	var out []interface{}
	err := _MultiSig.contract.Call(opts, &out, "transactions", arg0)

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
func (_MultiSig *MultiSigSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSig.Contract.Transactions(&_MultiSig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_MultiSig *MultiSigCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MultiSig.Contract.Transactions(&_MultiSig.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSig *MultiSigTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSig *MultiSigSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSig.Contract.AddOwner(&_MultiSig.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_MultiSig *MultiSigTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSig.Contract.AddOwner(&_MultiSig.TransactOpts, owner)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_MultiSig *MultiSigTransactor) ChangeInternalRequirement(opts *bind.TransactOpts, _internalRequired *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "changeInternalRequirement", _internalRequired)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_MultiSig *MultiSigSession) ChangeInternalRequirement(_internalRequired *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ChangeInternalRequirement(&_MultiSig.TransactOpts, _internalRequired)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_MultiSig *MultiSigTransactorSession) ChangeInternalRequirement(_internalRequired *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ChangeInternalRequirement(&_MultiSig.TransactOpts, _internalRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSig *MultiSigTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSig *MultiSigSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ChangeRequirement(&_MultiSig.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_MultiSig *MultiSigTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ChangeRequirement(&_MultiSig.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSig *MultiSigTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSig *MultiSigSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ConfirmTransaction(&_MultiSig.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_MultiSig *MultiSigTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ConfirmTransaction(&_MultiSig.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSig *MultiSigTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSig *MultiSigSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ExecuteTransaction(&_MultiSig.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_MultiSig *MultiSigTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.ExecuteTransaction(&_MultiSig.TransactOpts, transactionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_MultiSig *MultiSigTransactor) Initialize(opts *bind.TransactOpts, _owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "initialize", _owners, _required, _internalRequired)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_MultiSig *MultiSigSession) Initialize(_owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.Initialize(&_MultiSig.TransactOpts, _owners, _required, _internalRequired)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_MultiSig *MultiSigTransactorSession) Initialize(_owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.Initialize(&_MultiSig.TransactOpts, _owners, _required, _internalRequired)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSig *MultiSigTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSig *MultiSigSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSig.Contract.RemoveOwner(&_MultiSig.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_MultiSig *MultiSigTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MultiSig.Contract.RemoveOwner(&_MultiSig.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSig *MultiSigTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSig *MultiSigSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSig.Contract.ReplaceOwner(&_MultiSig.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_MultiSig *MultiSigTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MultiSig.Contract.ReplaceOwner(&_MultiSig.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSig *MultiSigTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSig *MultiSigSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.RevokeConfirmation(&_MultiSig.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_MultiSig *MultiSigTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MultiSig.Contract.RevokeConfirmation(&_MultiSig.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSig *MultiSigTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSig.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSig *MultiSigSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSig.Contract.SubmitTransaction(&_MultiSig.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_MultiSig *MultiSigTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MultiSig.Contract.SubmitTransaction(&_MultiSig.TransactOpts, destination, value, data)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_MultiSig *MultiSigFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _MultiSig.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "Confirmation":
		event, err = _MultiSig.ParseConfirmation(log)
	case "Deposit":
		event, err = _MultiSig.ParseDeposit(log)
	case "Execution":
		event, err = _MultiSig.ParseExecution(log)
	case "InternalRequirementChange":
		event, err = _MultiSig.ParseInternalRequirementChange(log)
	case "OwnerAddition":
		event, err = _MultiSig.ParseOwnerAddition(log)
	case "OwnerRemoval":
		event, err = _MultiSig.ParseOwnerRemoval(log)
	case "RequirementChange":
		event, err = _MultiSig.ParseRequirementChange(log)
	case "Revocation":
		event, err = _MultiSig.ParseRevocation(log)
	case "Submission":
		event, err = _MultiSig.ParseSubmission(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSig *MultiSigTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MultiSig.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSig *MultiSigSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSig.Contract.Fallback(&_MultiSig.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MultiSig *MultiSigTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MultiSig.Contract.Fallback(&_MultiSig.TransactOpts, calldata)
}

// MultiSigConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MultiSig contract.
type MultiSigConfirmationIterator struct {
	Event *MultiSigConfirmation // Event containing the contract specifics and raw log

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
func (it *MultiSigConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigConfirmation)
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
		it.Event = new(MultiSigConfirmation)
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
func (it *MultiSigConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigConfirmation represents a Confirmation event raised by the MultiSig contract.
type MultiSigConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSig *MultiSigFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigConfirmationIterator{contract: _MultiSig.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_MultiSig *MultiSigFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MultiSigConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigConfirmation)
				if err := _MultiSig.contract.UnpackLog(event, "Confirmation", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseConfirmation(log types.Log) (*MultiSigConfirmation, error) {
	event := new(MultiSigConfirmation)
	if err := _MultiSig.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the MultiSig contract.
type MultiSigDepositIterator struct {
	Event *MultiSigDeposit // Event containing the contract specifics and raw log

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
func (it *MultiSigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigDeposit)
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
		it.Event = new(MultiSigDeposit)
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
func (it *MultiSigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigDeposit represents a Deposit event raised by the MultiSig contract.
type MultiSigDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSig *MultiSigFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*MultiSigDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigDepositIterator{contract: _MultiSig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_MultiSig *MultiSigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MultiSigDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigDeposit)
				if err := _MultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseDeposit(log types.Log) (*MultiSigDeposit, error) {
	event := new(MultiSigDeposit)
	if err := _MultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MultiSig contract.
type MultiSigExecutionIterator struct {
	Event *MultiSigExecution // Event containing the contract specifics and raw log

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
func (it *MultiSigExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigExecution)
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
		it.Event = new(MultiSigExecution)
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
func (it *MultiSigExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigExecution represents a Execution event raised by the MultiSig contract.
type MultiSigExecution struct {
	TransactionId *big.Int
	ReturnData    []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_MultiSig *MultiSigFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigExecutionIterator{contract: _MultiSig.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_MultiSig *MultiSigFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MultiSigExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigExecution)
				if err := _MultiSig.contract.UnpackLog(event, "Execution", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseExecution(log types.Log) (*MultiSigExecution, error) {
	event := new(MultiSigExecution)
	if err := _MultiSig.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigInternalRequirementChangeIterator is returned from FilterInternalRequirementChange and is used to iterate over the raw logs and unpacked data for InternalRequirementChange events raised by the MultiSig contract.
type MultiSigInternalRequirementChangeIterator struct {
	Event *MultiSigInternalRequirementChange // Event containing the contract specifics and raw log

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
func (it *MultiSigInternalRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigInternalRequirementChange)
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
		it.Event = new(MultiSigInternalRequirementChange)
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
func (it *MultiSigInternalRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigInternalRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigInternalRequirementChange represents a InternalRequirementChange event raised by the MultiSig contract.
type MultiSigInternalRequirementChange struct {
	InternalRequired *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInternalRequirementChange is a free log retrieval operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_MultiSig *MultiSigFilterer) FilterInternalRequirementChange(opts *bind.FilterOpts) (*MultiSigInternalRequirementChangeIterator, error) {

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "InternalRequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigInternalRequirementChangeIterator{contract: _MultiSig.contract, event: "InternalRequirementChange", logs: logs, sub: sub}, nil
}

// WatchInternalRequirementChange is a free log subscription operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_MultiSig *MultiSigFilterer) WatchInternalRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigInternalRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "InternalRequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigInternalRequirementChange)
				if err := _MultiSig.contract.UnpackLog(event, "InternalRequirementChange", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseInternalRequirementChange(log types.Log) (*MultiSigInternalRequirementChange, error) {
	event := new(MultiSigInternalRequirementChange)
	if err := _MultiSig.contract.UnpackLog(event, "InternalRequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MultiSig contract.
type MultiSigOwnerAdditionIterator struct {
	Event *MultiSigOwnerAddition // Event containing the contract specifics and raw log

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
func (it *MultiSigOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigOwnerAddition)
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
		it.Event = new(MultiSigOwnerAddition)
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
func (it *MultiSigOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigOwnerAddition represents a OwnerAddition event raised by the MultiSig contract.
type MultiSigOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSig *MultiSigFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MultiSigOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigOwnerAdditionIterator{contract: _MultiSig.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_MultiSig *MultiSigFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MultiSigOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigOwnerAddition)
				if err := _MultiSig.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseOwnerAddition(log types.Log) (*MultiSigOwnerAddition, error) {
	event := new(MultiSigOwnerAddition)
	if err := _MultiSig.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MultiSig contract.
type MultiSigOwnerRemovalIterator struct {
	Event *MultiSigOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *MultiSigOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigOwnerRemoval)
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
		it.Event = new(MultiSigOwnerRemoval)
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
func (it *MultiSigOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigOwnerRemoval represents a OwnerRemoval event raised by the MultiSig contract.
type MultiSigOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSig *MultiSigFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MultiSigOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigOwnerRemovalIterator{contract: _MultiSig.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_MultiSig *MultiSigFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MultiSigOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigOwnerRemoval)
				if err := _MultiSig.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseOwnerRemoval(log types.Log) (*MultiSigOwnerRemoval, error) {
	event := new(MultiSigOwnerRemoval)
	if err := _MultiSig.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MultiSig contract.
type MultiSigRequirementChangeIterator struct {
	Event *MultiSigRequirementChange // Event containing the contract specifics and raw log

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
func (it *MultiSigRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigRequirementChange)
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
		it.Event = new(MultiSigRequirementChange)
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
func (it *MultiSigRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigRequirementChange represents a RequirementChange event raised by the MultiSig contract.
type MultiSigRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSig *MultiSigFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*MultiSigRequirementChangeIterator, error) {

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MultiSigRequirementChangeIterator{contract: _MultiSig.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_MultiSig *MultiSigFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MultiSigRequirementChange) (event.Subscription, error) {

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigRequirementChange)
				if err := _MultiSig.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseRequirementChange(log types.Log) (*MultiSigRequirementChange, error) {
	event := new(MultiSigRequirementChange)
	if err := _MultiSig.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MultiSig contract.
type MultiSigRevocationIterator struct {
	Event *MultiSigRevocation // Event containing the contract specifics and raw log

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
func (it *MultiSigRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigRevocation)
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
		it.Event = new(MultiSigRevocation)
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
func (it *MultiSigRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigRevocation represents a Revocation event raised by the MultiSig contract.
type MultiSigRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSig *MultiSigFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MultiSigRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigRevocationIterator{contract: _MultiSig.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_MultiSig *MultiSigFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MultiSigRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigRevocation)
				if err := _MultiSig.contract.UnpackLog(event, "Revocation", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseRevocation(log types.Log) (*MultiSigRevocation, error) {
	event := new(MultiSigRevocation)
	if err := _MultiSig.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MultiSigSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MultiSig contract.
type MultiSigSubmissionIterator struct {
	Event *MultiSigSubmission // Event containing the contract specifics and raw log

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
func (it *MultiSigSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MultiSigSubmission)
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
		it.Event = new(MultiSigSubmission)
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
func (it *MultiSigSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MultiSigSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MultiSigSubmission represents a Submission event raised by the MultiSig contract.
type MultiSigSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSig *MultiSigFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MultiSigSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MultiSigSubmissionIterator{contract: _MultiSig.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_MultiSig *MultiSigFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MultiSigSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MultiSig.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MultiSigSubmission)
				if err := _MultiSig.contract.UnpackLog(event, "Submission", log); err != nil {
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
func (_MultiSig *MultiSigFilterer) ParseSubmission(log types.Log) (*MultiSigSubmission, error) {
	event := new(MultiSigSubmission)
	if err := _MultiSig.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
