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

// FreezerMetaData contains all meta data concerning the Freezer contract.
var FreezerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isFrozen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"freeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"unfreeze\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// FreezerABI is the input ABI used to generate the binding from.
// Deprecated: Use FreezerMetaData.ABI instead.
var FreezerABI = FreezerMetaData.ABI

// Freezer is an auto generated Go binding around an Ethereum contract.
type Freezer struct {
	FreezerCaller     // Read-only binding to the contract
	FreezerTransactor // Write-only binding to the contract
	FreezerFilterer   // Log filterer for contract events
}

// FreezerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FreezerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreezerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FreezerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreezerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FreezerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FreezerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FreezerSession struct {
	Contract     *Freezer          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FreezerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FreezerCallerSession struct {
	Contract *FreezerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FreezerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FreezerTransactorSession struct {
	Contract     *FreezerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FreezerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FreezerRaw struct {
	Contract *Freezer // Generic contract binding to access the raw methods on
}

// FreezerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FreezerCallerRaw struct {
	Contract *FreezerCaller // Generic read-only contract binding to access the raw methods on
}

// FreezerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FreezerTransactorRaw struct {
	Contract *FreezerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFreezer creates a new instance of Freezer, bound to a specific deployed contract.
func NewFreezer(address common.Address, backend bind.ContractBackend) (*Freezer, error) {
	contract, err := bindFreezer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Freezer{FreezerCaller: FreezerCaller{contract: contract}, FreezerTransactor: FreezerTransactor{contract: contract}, FreezerFilterer: FreezerFilterer{contract: contract}}, nil
}

// NewFreezerCaller creates a new read-only instance of Freezer, bound to a specific deployed contract.
func NewFreezerCaller(address common.Address, caller bind.ContractCaller) (*FreezerCaller, error) {
	contract, err := bindFreezer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FreezerCaller{contract: contract}, nil
}

// NewFreezerTransactor creates a new write-only instance of Freezer, bound to a specific deployed contract.
func NewFreezerTransactor(address common.Address, transactor bind.ContractTransactor) (*FreezerTransactor, error) {
	contract, err := bindFreezer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FreezerTransactor{contract: contract}, nil
}

// NewFreezerFilterer creates a new log filterer instance of Freezer, bound to a specific deployed contract.
func NewFreezerFilterer(address common.Address, filterer bind.ContractFilterer) (*FreezerFilterer, error) {
	contract, err := bindFreezer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FreezerFilterer{contract: contract}, nil
}

// bindFreezer binds a generic wrapper to an already deployed contract.
func bindFreezer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FreezerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseFreezerABI parses the ABI
func ParseFreezerABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(FreezerABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Freezer *FreezerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Freezer.Contract.FreezerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Freezer *FreezerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Freezer.Contract.FreezerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Freezer *FreezerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Freezer.Contract.FreezerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Freezer *FreezerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Freezer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Freezer *FreezerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Freezer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Freezer *FreezerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Freezer.Contract.contract.Transact(opts, method, params...)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Freezer *FreezerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Freezer.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Freezer *FreezerSession) Initialized() (bool, error) {
	return _Freezer.Contract.Initialized(&_Freezer.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Freezer *FreezerCallerSession) Initialized() (bool, error) {
	return _Freezer.Contract.Initialized(&_Freezer.CallOpts)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address ) view returns(bool)
func (_Freezer *FreezerCaller) IsFrozen(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Freezer.contract.Call(opts, &out, "isFrozen", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address ) view returns(bool)
func (_Freezer *FreezerSession) IsFrozen(arg0 common.Address) (bool, error) {
	return _Freezer.Contract.IsFrozen(&_Freezer.CallOpts, arg0)
}

// IsFrozen is a free data retrieval call binding the contract method 0xe5839836.
//
// Solidity: function isFrozen(address ) view returns(bool)
func (_Freezer *FreezerCallerSession) IsFrozen(arg0 common.Address) (bool, error) {
	return _Freezer.Contract.IsFrozen(&_Freezer.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Freezer *FreezerCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Freezer.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Freezer *FreezerSession) IsOwner() (bool, error) {
	return _Freezer.Contract.IsOwner(&_Freezer.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Freezer *FreezerCallerSession) IsOwner() (bool, error) {
	return _Freezer.Contract.IsOwner(&_Freezer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Freezer *FreezerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Freezer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Freezer *FreezerSession) Owner() (common.Address, error) {
	return _Freezer.Contract.Owner(&_Freezer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Freezer *FreezerCallerSession) Owner() (common.Address, error) {
	return _Freezer.Contract.Owner(&_Freezer.CallOpts)
}

// Freeze is a paid mutator transaction binding the contract method 0x8d1fdf2f.
//
// Solidity: function freeze(address target) returns()
func (_Freezer *FreezerTransactor) Freeze(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Freezer.contract.Transact(opts, "freeze", target)
}

// Freeze is a paid mutator transaction binding the contract method 0x8d1fdf2f.
//
// Solidity: function freeze(address target) returns()
func (_Freezer *FreezerSession) Freeze(target common.Address) (*types.Transaction, error) {
	return _Freezer.Contract.Freeze(&_Freezer.TransactOpts, target)
}

// Freeze is a paid mutator transaction binding the contract method 0x8d1fdf2f.
//
// Solidity: function freeze(address target) returns()
func (_Freezer *FreezerTransactorSession) Freeze(target common.Address) (*types.Transaction, error) {
	return _Freezer.Contract.Freeze(&_Freezer.TransactOpts, target)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Freezer *FreezerTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Freezer.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Freezer *FreezerSession) Initialize() (*types.Transaction, error) {
	return _Freezer.Contract.Initialize(&_Freezer.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Freezer *FreezerTransactorSession) Initialize() (*types.Transaction, error) {
	return _Freezer.Contract.Initialize(&_Freezer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Freezer *FreezerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Freezer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Freezer *FreezerSession) RenounceOwnership() (*types.Transaction, error) {
	return _Freezer.Contract.RenounceOwnership(&_Freezer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Freezer *FreezerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Freezer.Contract.RenounceOwnership(&_Freezer.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Freezer *FreezerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Freezer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Freezer *FreezerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Freezer.Contract.TransferOwnership(&_Freezer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Freezer *FreezerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Freezer.Contract.TransferOwnership(&_Freezer.TransactOpts, newOwner)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x45c8b1a6.
//
// Solidity: function unfreeze(address target) returns()
func (_Freezer *FreezerTransactor) Unfreeze(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _Freezer.contract.Transact(opts, "unfreeze", target)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x45c8b1a6.
//
// Solidity: function unfreeze(address target) returns()
func (_Freezer *FreezerSession) Unfreeze(target common.Address) (*types.Transaction, error) {
	return _Freezer.Contract.Unfreeze(&_Freezer.TransactOpts, target)
}

// Unfreeze is a paid mutator transaction binding the contract method 0x45c8b1a6.
//
// Solidity: function unfreeze(address target) returns()
func (_Freezer *FreezerTransactorSession) Unfreeze(target common.Address) (*types.Transaction, error) {
	return _Freezer.Contract.Unfreeze(&_Freezer.TransactOpts, target)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Freezer *FreezerFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Freezer.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "OwnershipTransferred":
		event, err = _Freezer.ParseOwnershipTransferred(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// FreezerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Freezer contract.
type FreezerOwnershipTransferredIterator struct {
	Event *FreezerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FreezerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FreezerOwnershipTransferred)
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
		it.Event = new(FreezerOwnershipTransferred)
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
func (it *FreezerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FreezerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FreezerOwnershipTransferred represents a OwnershipTransferred event raised by the Freezer contract.
type FreezerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Freezer *FreezerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FreezerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Freezer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FreezerOwnershipTransferredIterator{contract: _Freezer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Freezer *FreezerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FreezerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Freezer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FreezerOwnershipTransferred)
				if err := _Freezer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Freezer *FreezerFilterer) ParseOwnershipTransferred(log types.Log) (*FreezerOwnershipTransferred, error) {
	event := new(FreezerOwnershipTransferred)
	if err := _Freezer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
