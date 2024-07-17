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

// GovernanceApproverMultiSigMetaData contains all meta data concerning the GovernanceApproverMultiSig contract.
var GovernanceApproverMultiSigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"internalRequired\",\"type\":\"uint256\"}],\"name\":\"InternalRequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_internalRequired\",\"type\":\"uint256\"}],\"name\":\"changeInternalRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"pending\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_owners\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"_required\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_internalRequired\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"internalRequired\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GovernanceApproverMultiSigABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceApproverMultiSigMetaData.ABI instead.
var GovernanceApproverMultiSigABI = GovernanceApproverMultiSigMetaData.ABI

// GovernanceApproverMultiSig is an auto generated Go binding around an Ethereum contract.
type GovernanceApproverMultiSig struct {
	GovernanceApproverMultiSigCaller     // Read-only binding to the contract
	GovernanceApproverMultiSigTransactor // Write-only binding to the contract
	GovernanceApproverMultiSigFilterer   // Log filterer for contract events
}

// GovernanceApproverMultiSigCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceApproverMultiSigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceApproverMultiSigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceApproverMultiSigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceApproverMultiSigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceApproverMultiSigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceApproverMultiSigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceApproverMultiSigSession struct {
	Contract     *GovernanceApproverMultiSig // Generic contract binding to set the session for
	CallOpts     bind.CallOpts               // Call options to use throughout this session
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// GovernanceApproverMultiSigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceApproverMultiSigCallerSession struct {
	Contract *GovernanceApproverMultiSigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                     // Call options to use throughout this session
}

// GovernanceApproverMultiSigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceApproverMultiSigTransactorSession struct {
	Contract     *GovernanceApproverMultiSigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                     // Transaction auth options to use throughout this session
}

// GovernanceApproverMultiSigRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceApproverMultiSigRaw struct {
	Contract *GovernanceApproverMultiSig // Generic contract binding to access the raw methods on
}

// GovernanceApproverMultiSigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceApproverMultiSigCallerRaw struct {
	Contract *GovernanceApproverMultiSigCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceApproverMultiSigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceApproverMultiSigTransactorRaw struct {
	Contract *GovernanceApproverMultiSigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernanceApproverMultiSig creates a new instance of GovernanceApproverMultiSig, bound to a specific deployed contract.
func NewGovernanceApproverMultiSig(address common.Address, backend bind.ContractBackend) (*GovernanceApproverMultiSig, error) {
	contract, err := bindGovernanceApproverMultiSig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSig{GovernanceApproverMultiSigCaller: GovernanceApproverMultiSigCaller{contract: contract}, GovernanceApproverMultiSigTransactor: GovernanceApproverMultiSigTransactor{contract: contract}, GovernanceApproverMultiSigFilterer: GovernanceApproverMultiSigFilterer{contract: contract}}, nil
}

// NewGovernanceApproverMultiSigCaller creates a new read-only instance of GovernanceApproverMultiSig, bound to a specific deployed contract.
func NewGovernanceApproverMultiSigCaller(address common.Address, caller bind.ContractCaller) (*GovernanceApproverMultiSigCaller, error) {
	contract, err := bindGovernanceApproverMultiSig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigCaller{contract: contract}, nil
}

// NewGovernanceApproverMultiSigTransactor creates a new write-only instance of GovernanceApproverMultiSig, bound to a specific deployed contract.
func NewGovernanceApproverMultiSigTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceApproverMultiSigTransactor, error) {
	contract, err := bindGovernanceApproverMultiSig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigTransactor{contract: contract}, nil
}

// NewGovernanceApproverMultiSigFilterer creates a new log filterer instance of GovernanceApproverMultiSig, bound to a specific deployed contract.
func NewGovernanceApproverMultiSigFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceApproverMultiSigFilterer, error) {
	contract, err := bindGovernanceApproverMultiSig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigFilterer{contract: contract}, nil
}

// bindGovernanceApproverMultiSig binds a generic wrapper to an already deployed contract.
func bindGovernanceApproverMultiSig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceApproverMultiSigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseGovernanceApproverMultiSigABI parses the ABI
func ParseGovernanceApproverMultiSigABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceApproverMultiSigABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceApproverMultiSig.Contract.GovernanceApproverMultiSigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.GovernanceApproverMultiSigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.GovernanceApproverMultiSigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GovernanceApproverMultiSig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.contract.Transact(opts, method, params...)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "MAX_OWNER_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.MAXOWNERCOUNT(&_GovernanceApproverMultiSig.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.MAXOWNERCOUNT(&_GovernanceApproverMultiSig.CallOpts)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "confirmations", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GovernanceApproverMultiSig.Contract.Confirmations(&_GovernanceApproverMultiSig.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations(uint256 , address ) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _GovernanceApproverMultiSig.Contract.Confirmations(&_GovernanceApproverMultiSig.CallOpts, arg0, arg1)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "getConfirmationCount", transactionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.GetConfirmationCount(&_GovernanceApproverMultiSig.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(uint256 transactionId) view returns(uint256 count)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.GetConfirmationCount(&_GovernanceApproverMultiSig.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "getConfirmations", transactionId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _GovernanceApproverMultiSig.Contract.GetConfirmations(&_GovernanceApproverMultiSig.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(uint256 transactionId) view returns(address[] _confirmations)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _GovernanceApproverMultiSig.Contract.GetConfirmations(&_GovernanceApproverMultiSig.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "getOwners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) GetOwners() ([]common.Address, error) {
	return _GovernanceApproverMultiSig.Contract.GetOwners(&_GovernanceApproverMultiSig.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() view returns(address[])
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) GetOwners() ([]common.Address, error) {
	return _GovernanceApproverMultiSig.Contract.GetOwners(&_GovernanceApproverMultiSig.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "getTransactionCount", pending, executed)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.GetTransactionCount(&_GovernanceApproverMultiSig.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(bool pending, bool executed) view returns(uint256 count)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.GetTransactionCount(&_GovernanceApproverMultiSig.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "getTransactionIds", from, to, pending, executed)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.GetTransactionIds(&_GovernanceApproverMultiSig.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(uint256 from, uint256 to, bool pending, bool executed) view returns(uint256[] _transactionIds)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.GetTransactionIds(&_GovernanceApproverMultiSig.CallOpts, from, to, pending, executed)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Initialized() (bool, error) {
	return _GovernanceApproverMultiSig.Contract.Initialized(&_GovernanceApproverMultiSig.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) Initialized() (bool, error) {
	return _GovernanceApproverMultiSig.Contract.Initialized(&_GovernanceApproverMultiSig.CallOpts)
}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) InternalRequired(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "internalRequired")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) InternalRequired() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.InternalRequired(&_GovernanceApproverMultiSig.CallOpts)
}

// InternalRequired is a free data retrieval call binding the contract method 0xa24efcdf.
//
// Solidity: function internalRequired() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) InternalRequired() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.InternalRequired(&_GovernanceApproverMultiSig.CallOpts)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "isConfirmed", transactionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _GovernanceApproverMultiSig.Contract.IsConfirmed(&_GovernanceApproverMultiSig.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(uint256 transactionId) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _GovernanceApproverMultiSig.Contract.IsConfirmed(&_GovernanceApproverMultiSig.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "isOwner", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) IsOwner(arg0 common.Address) (bool, error) {
	return _GovernanceApproverMultiSig.Contract.IsOwner(&_GovernanceApproverMultiSig.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address ) view returns(bool)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _GovernanceApproverMultiSig.Contract.IsOwner(&_GovernanceApproverMultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _GovernanceApproverMultiSig.Contract.Owners(&_GovernanceApproverMultiSig.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners(uint256 ) view returns(address)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _GovernanceApproverMultiSig.Contract.Owners(&_GovernanceApproverMultiSig.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "required")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Required() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.Required(&_GovernanceApproverMultiSig.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) Required() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.Required(&_GovernanceApproverMultiSig.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "transactionCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) TransactionCount() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.TransactionCount(&_GovernanceApproverMultiSig.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() view returns(uint256)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) TransactionCount() (*big.Int, error) {
	return _GovernanceApproverMultiSig.Contract.TransactionCount(&_GovernanceApproverMultiSig.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCaller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	var out []interface{}
	err := _GovernanceApproverMultiSig.contract.Call(opts, &out, "transactions", arg0)

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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _GovernanceApproverMultiSig.Contract.Transactions(&_GovernanceApproverMultiSig.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions(uint256 ) view returns(address destination, uint256 value, bytes data, bool executed)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigCallerSession) Transactions(arg0 *big.Int) (struct {
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _GovernanceApproverMultiSig.Contract.Transactions(&_GovernanceApproverMultiSig.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.AddOwner(&_GovernanceApproverMultiSig.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(address owner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.AddOwner(&_GovernanceApproverMultiSig.TransactOpts, owner)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) ChangeInternalRequirement(opts *bind.TransactOpts, _internalRequired *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "changeInternalRequirement", _internalRequired)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) ChangeInternalRequirement(_internalRequired *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ChangeInternalRequirement(&_GovernanceApproverMultiSig.TransactOpts, _internalRequired)
}

// ChangeInternalRequirement is a paid mutator transaction binding the contract method 0x2e6c3721.
//
// Solidity: function changeInternalRequirement(uint256 _internalRequired) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) ChangeInternalRequirement(_internalRequired *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ChangeInternalRequirement(&_GovernanceApproverMultiSig.TransactOpts, _internalRequired)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ChangeRequirement(&_GovernanceApproverMultiSig.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(uint256 _required) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ChangeRequirement(&_GovernanceApproverMultiSig.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ConfirmTransaction(&_GovernanceApproverMultiSig.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ConfirmTransaction(&_GovernanceApproverMultiSig.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ExecuteTransaction(&_GovernanceApproverMultiSig.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ExecuteTransaction(&_GovernanceApproverMultiSig.TransactOpts, transactionId)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) Initialize(opts *bind.TransactOpts, _owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "initialize", _owners, _required, _internalRequired)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Initialize(_owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.Initialize(&_GovernanceApproverMultiSig.TransactOpts, _owners, _required, _internalRequired)
}

// Initialize is a paid mutator transaction binding the contract method 0x5eae7959.
//
// Solidity: function initialize(address[] _owners, uint256 _required, uint256 _internalRequired) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) Initialize(_owners []common.Address, _required *big.Int, _internalRequired *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.Initialize(&_GovernanceApproverMultiSig.TransactOpts, _owners, _required, _internalRequired)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.RemoveOwner(&_GovernanceApproverMultiSig.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(address owner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.RemoveOwner(&_GovernanceApproverMultiSig.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ReplaceOwner(&_GovernanceApproverMultiSig.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(address owner, address newOwner) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.ReplaceOwner(&_GovernanceApproverMultiSig.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.RevokeConfirmation(&_GovernanceApproverMultiSig.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(uint256 transactionId) returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.RevokeConfirmation(&_GovernanceApproverMultiSig.TransactOpts, transactionId)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.SubmitTransaction(&_GovernanceApproverMultiSig.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(address destination, uint256 value, bytes data) returns(uint256 transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.SubmitTransaction(&_GovernanceApproverMultiSig.TransactOpts, destination, value, data)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _GovernanceApproverMultiSig.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "Confirmation":
		event, err = _GovernanceApproverMultiSig.ParseConfirmation(log)
	case "Deposit":
		event, err = _GovernanceApproverMultiSig.ParseDeposit(log)
	case "Execution":
		event, err = _GovernanceApproverMultiSig.ParseExecution(log)
	case "InternalRequirementChange":
		event, err = _GovernanceApproverMultiSig.ParseInternalRequirementChange(log)
	case "OwnerAddition":
		event, err = _GovernanceApproverMultiSig.ParseOwnerAddition(log)
	case "OwnerRemoval":
		event, err = _GovernanceApproverMultiSig.ParseOwnerRemoval(log)
	case "RequirementChange":
		event, err = _GovernanceApproverMultiSig.ParseRequirementChange(log)
	case "Revocation":
		event, err = _GovernanceApproverMultiSig.ParseRevocation(log)
	case "Submission":
		event, err = _GovernanceApproverMultiSig.ParseSubmission(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.Fallback(&_GovernanceApproverMultiSig.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _GovernanceApproverMultiSig.Contract.Fallback(&_GovernanceApproverMultiSig.TransactOpts, calldata)
}

// GovernanceApproverMultiSigConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigConfirmationIterator struct {
	Event *GovernanceApproverMultiSigConfirmation // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigConfirmation)
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
		it.Event = new(GovernanceApproverMultiSigConfirmation)
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
func (it *GovernanceApproverMultiSigConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigConfirmation represents a Confirmation event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigConfirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*GovernanceApproverMultiSigConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigConfirmationIterator{contract: _GovernanceApproverMultiSig.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(address indexed sender, uint256 indexed transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigConfirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigConfirmation)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Confirmation", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseConfirmation(log types.Log) (*GovernanceApproverMultiSigConfirmation, error) {
	event := new(GovernanceApproverMultiSigConfirmation)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Confirmation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigDepositIterator struct {
	Event *GovernanceApproverMultiSigDeposit // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigDeposit)
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
		it.Event = new(GovernanceApproverMultiSigDeposit)
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
func (it *GovernanceApproverMultiSigDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigDeposit represents a Deposit event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigDeposit struct {
	Sender common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterDeposit(opts *bind.FilterOpts, sender []common.Address) (*GovernanceApproverMultiSigDepositIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigDepositIterator{contract: _GovernanceApproverMultiSig.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed sender, uint256 value)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigDeposit, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "Deposit", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigDeposit)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseDeposit(log types.Log) (*GovernanceApproverMultiSigDeposit, error) {
	event := new(GovernanceApproverMultiSigDeposit)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigExecutionIterator struct {
	Event *GovernanceApproverMultiSigExecution // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigExecution)
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
		it.Event = new(GovernanceApproverMultiSigExecution)
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
func (it *GovernanceApproverMultiSigExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigExecution represents a Execution event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigExecution struct {
	TransactionId *big.Int
	ReturnData    []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int) (*GovernanceApproverMultiSigExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigExecutionIterator{contract: _GovernanceApproverMultiSig.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x0c18aae526accb31b01cf9a15bdf435e70632ee31efc4c5c0752c4262ea45d2f.
//
// Solidity: event Execution(uint256 indexed transactionId, bytes returnData)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigExecution, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "Execution", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigExecution)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Execution", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseExecution(log types.Log) (*GovernanceApproverMultiSigExecution, error) {
	event := new(GovernanceApproverMultiSigExecution)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Execution", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigInternalRequirementChangeIterator is returned from FilterInternalRequirementChange and is used to iterate over the raw logs and unpacked data for InternalRequirementChange events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigInternalRequirementChangeIterator struct {
	Event *GovernanceApproverMultiSigInternalRequirementChange // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigInternalRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigInternalRequirementChange)
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
		it.Event = new(GovernanceApproverMultiSigInternalRequirementChange)
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
func (it *GovernanceApproverMultiSigInternalRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigInternalRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigInternalRequirementChange represents a InternalRequirementChange event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigInternalRequirementChange struct {
	InternalRequired *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterInternalRequirementChange is a free log retrieval operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterInternalRequirementChange(opts *bind.FilterOpts) (*GovernanceApproverMultiSigInternalRequirementChangeIterator, error) {

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "InternalRequirementChange")
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigInternalRequirementChangeIterator{contract: _GovernanceApproverMultiSig.contract, event: "InternalRequirementChange", logs: logs, sub: sub}, nil
}

// WatchInternalRequirementChange is a free log subscription operation binding the contract event 0xa07eff79ea50418b0e96ff7c01d65eb6c3a5a240ee91cd81c70c89503dd41239.
//
// Solidity: event InternalRequirementChange(uint256 internalRequired)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchInternalRequirementChange(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigInternalRequirementChange) (event.Subscription, error) {

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "InternalRequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigInternalRequirementChange)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "InternalRequirementChange", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseInternalRequirementChange(log types.Log) (*GovernanceApproverMultiSigInternalRequirementChange, error) {
	event := new(GovernanceApproverMultiSigInternalRequirementChange)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "InternalRequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigOwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigOwnerAdditionIterator struct {
	Event *GovernanceApproverMultiSigOwnerAddition // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigOwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigOwnerAddition)
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
		it.Event = new(GovernanceApproverMultiSigOwnerAddition)
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
func (it *GovernanceApproverMultiSigOwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigOwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigOwnerAddition represents a OwnerAddition event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigOwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*GovernanceApproverMultiSigOwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigOwnerAdditionIterator{contract: _GovernanceApproverMultiSig.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(address indexed owner)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigOwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigOwnerAddition)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseOwnerAddition(log types.Log) (*GovernanceApproverMultiSigOwnerAddition, error) {
	event := new(GovernanceApproverMultiSigOwnerAddition)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigOwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigOwnerRemovalIterator struct {
	Event *GovernanceApproverMultiSigOwnerRemoval // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigOwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigOwnerRemoval)
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
		it.Event = new(GovernanceApproverMultiSigOwnerRemoval)
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
func (it *GovernanceApproverMultiSigOwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigOwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigOwnerRemoval represents a OwnerRemoval event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigOwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*GovernanceApproverMultiSigOwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigOwnerRemovalIterator{contract: _GovernanceApproverMultiSig.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(address indexed owner)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigOwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigOwnerRemoval)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseOwnerRemoval(log types.Log) (*GovernanceApproverMultiSigOwnerRemoval, error) {
	event := new(GovernanceApproverMultiSigOwnerRemoval)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigRequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigRequirementChangeIterator struct {
	Event *GovernanceApproverMultiSigRequirementChange // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigRequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigRequirementChange)
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
		it.Event = new(GovernanceApproverMultiSigRequirementChange)
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
func (it *GovernanceApproverMultiSigRequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigRequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigRequirementChange represents a RequirementChange event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigRequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterRequirementChange(opts *bind.FilterOpts) (*GovernanceApproverMultiSigRequirementChangeIterator, error) {

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigRequirementChangeIterator{contract: _GovernanceApproverMultiSig.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(uint256 required)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigRequirementChange) (event.Subscription, error) {

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigRequirementChange)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "RequirementChange", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseRequirementChange(log types.Log) (*GovernanceApproverMultiSigRequirementChange, error) {
	event := new(GovernanceApproverMultiSigRequirementChange)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "RequirementChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigRevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigRevocationIterator struct {
	Event *GovernanceApproverMultiSigRevocation // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigRevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigRevocation)
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
		it.Event = new(GovernanceApproverMultiSigRevocation)
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
func (it *GovernanceApproverMultiSigRevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigRevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigRevocation represents a Revocation event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigRevocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*GovernanceApproverMultiSigRevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigRevocationIterator{contract: _GovernanceApproverMultiSig.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(address indexed sender, uint256 indexed transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigRevocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigRevocation)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Revocation", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseRevocation(log types.Log) (*GovernanceApproverMultiSigRevocation, error) {
	event := new(GovernanceApproverMultiSigRevocation)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Revocation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceApproverMultiSigSubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigSubmissionIterator struct {
	Event *GovernanceApproverMultiSigSubmission // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverMultiSigSubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverMultiSigSubmission)
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
		it.Event = new(GovernanceApproverMultiSigSubmission)
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
func (it *GovernanceApproverMultiSigSubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverMultiSigSubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverMultiSigSubmission represents a Submission event raised by the GovernanceApproverMultiSig contract.
type GovernanceApproverMultiSigSubmission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*GovernanceApproverMultiSigSubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverMultiSigSubmissionIterator{contract: _GovernanceApproverMultiSig.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(uint256 indexed transactionId)
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *GovernanceApproverMultiSigSubmission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _GovernanceApproverMultiSig.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverMultiSigSubmission)
				if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Submission", log); err != nil {
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
func (_GovernanceApproverMultiSig *GovernanceApproverMultiSigFilterer) ParseSubmission(log types.Log) (*GovernanceApproverMultiSigSubmission, error) {
	event := new(GovernanceApproverMultiSigSubmission)
	if err := _GovernanceApproverMultiSig.contract.UnpackLog(event, "Submission", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
