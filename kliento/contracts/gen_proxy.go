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

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"ImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerSet\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"_transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callbackData\",\"type\":\"bytes\"}],\"name\":\"_setAndInitializeImplementation\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_getImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseProxyABI parses the ABI
func ParseProxyABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// GetImplementation is a free data retrieval call binding the contract method 0x42404e07.
//
// Solidity: function _getImplementation() view returns(address implementation)
func (_Proxy *ProxyCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "_getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0x42404e07.
//
// Solidity: function _getImplementation() view returns(address implementation)
func (_Proxy *ProxySession) GetImplementation() (common.Address, error) {
	return _Proxy.Contract.GetImplementation(&_Proxy.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0x42404e07.
//
// Solidity: function _getImplementation() view returns(address implementation)
func (_Proxy *ProxyCallerSession) GetImplementation() (common.Address, error) {
	return _Proxy.Contract.GetImplementation(&_Proxy.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0xf7e6af80.
//
// Solidity: function _getOwner() view returns(address owner)
func (_Proxy *ProxyCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "_getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0xf7e6af80.
//
// Solidity: function _getOwner() view returns(address owner)
func (_Proxy *ProxySession) GetOwner() (common.Address, error) {
	return _Proxy.Contract.GetOwner(&_Proxy.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0xf7e6af80.
//
// Solidity: function _getOwner() view returns(address owner)
func (_Proxy *ProxyCallerSession) GetOwner() (common.Address, error) {
	return _Proxy.Contract.GetOwner(&_Proxy.CallOpts)
}

// SetAndInitializeImplementation is a paid mutator transaction binding the contract method 0x03386ba3.
//
// Solidity: function _setAndInitializeImplementation(address implementation, bytes callbackData) payable returns()
func (_Proxy *ProxyTransactor) SetAndInitializeImplementation(opts *bind.TransactOpts, implementation common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "_setAndInitializeImplementation", implementation, callbackData)
}

// SetAndInitializeImplementation is a paid mutator transaction binding the contract method 0x03386ba3.
//
// Solidity: function _setAndInitializeImplementation(address implementation, bytes callbackData) payable returns()
func (_Proxy *ProxySession) SetAndInitializeImplementation(implementation common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetAndInitializeImplementation(&_Proxy.TransactOpts, implementation, callbackData)
}

// SetAndInitializeImplementation is a paid mutator transaction binding the contract method 0x03386ba3.
//
// Solidity: function _setAndInitializeImplementation(address implementation, bytes callbackData) payable returns()
func (_Proxy *ProxyTransactorSession) SetAndInitializeImplementation(implementation common.Address, callbackData []byte) (*types.Transaction, error) {
	return _Proxy.Contract.SetAndInitializeImplementation(&_Proxy.TransactOpts, implementation, callbackData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xbb913f41.
//
// Solidity: function _setImplementation(address implementation) returns()
func (_Proxy *ProxyTransactor) SetImplementation(opts *bind.TransactOpts, implementation common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "_setImplementation", implementation)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xbb913f41.
//
// Solidity: function _setImplementation(address implementation) returns()
func (_Proxy *ProxySession) SetImplementation(implementation common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetImplementation(&_Proxy.TransactOpts, implementation)
}

// SetImplementation is a paid mutator transaction binding the contract method 0xbb913f41.
//
// Solidity: function _setImplementation(address implementation) returns()
func (_Proxy *ProxyTransactorSession) SetImplementation(implementation common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.SetImplementation(&_Proxy.TransactOpts, implementation)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xd29d44ee.
//
// Solidity: function _transferOwnership(address newOwner) returns()
func (_Proxy *ProxyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "_transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xd29d44ee.
//
// Solidity: function _transferOwnership(address newOwner) returns()
func (_Proxy *ProxySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xd29d44ee.
//
// Solidity: function _transferOwnership(address newOwner) returns()
func (_Proxy *ProxyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Proxy.Contract.TransferOwnership(&_Proxy.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Proxy *ProxyFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Proxy.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "ImplementationSet":
		event, err = _Proxy.ParseImplementationSet(log)
	case "OwnerSet":
		event, err = _Proxy.ParseOwnerSet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxy *ProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxy.Contract.Fallback(&_Proxy.TransactOpts, calldata)
}

// ProxyImplementationSetIterator is returned from FilterImplementationSet and is used to iterate over the raw logs and unpacked data for ImplementationSet events raised by the Proxy contract.
type ProxyImplementationSetIterator struct {
	Event *ProxyImplementationSet // Event containing the contract specifics and raw log

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
func (it *ProxyImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyImplementationSet)
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
		it.Event = new(ProxyImplementationSet)
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
func (it *ProxyImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyImplementationSet represents a ImplementationSet event raised by the Proxy contract.
type ProxyImplementationSet struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterImplementationSet is a free log retrieval operation binding the contract event 0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13.
//
// Solidity: event ImplementationSet(address indexed implementation)
func (_Proxy *ProxyFilterer) FilterImplementationSet(opts *bind.FilterOpts, implementation []common.Address) (*ProxyImplementationSetIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "ImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return &ProxyImplementationSetIterator{contract: _Proxy.contract, event: "ImplementationSet", logs: logs, sub: sub}, nil
}

// WatchImplementationSet is a free log subscription operation binding the contract event 0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13.
//
// Solidity: event ImplementationSet(address indexed implementation)
func (_Proxy *ProxyFilterer) WatchImplementationSet(opts *bind.WatchOpts, sink chan<- *ProxyImplementationSet, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "ImplementationSet", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyImplementationSet)
				if err := _Proxy.contract.UnpackLog(event, "ImplementationSet", log); err != nil {
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

// ParseImplementationSet is a log parse operation binding the contract event 0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13.
//
// Solidity: event ImplementationSet(address indexed implementation)
func (_Proxy *ProxyFilterer) ParseImplementationSet(log types.Log) (*ProxyImplementationSet, error) {
	event := new(ProxyImplementationSet)
	if err := _Proxy.contract.UnpackLog(event, "ImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyOwnerSetIterator is returned from FilterOwnerSet and is used to iterate over the raw logs and unpacked data for OwnerSet events raised by the Proxy contract.
type ProxyOwnerSetIterator struct {
	Event *ProxyOwnerSet // Event containing the contract specifics and raw log

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
func (it *ProxyOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProxyOwnerSet)
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
		it.Event = new(ProxyOwnerSet)
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
func (it *ProxyOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProxyOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProxyOwnerSet represents a OwnerSet event raised by the Proxy contract.
type ProxyOwnerSet struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerSet is a free log retrieval operation binding the contract event 0x50146d0e3c60aa1d17a70635b05494f864e86144a2201275021014fbf08bafe2.
//
// Solidity: event OwnerSet(address indexed owner)
func (_Proxy *ProxyFilterer) FilterOwnerSet(opts *bind.FilterOpts, owner []common.Address) (*ProxyOwnerSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Proxy.contract.FilterLogs(opts, "OwnerSet", ownerRule)
	if err != nil {
		return nil, err
	}
	return &ProxyOwnerSetIterator{contract: _Proxy.contract, event: "OwnerSet", logs: logs, sub: sub}, nil
}

// WatchOwnerSet is a free log subscription operation binding the contract event 0x50146d0e3c60aa1d17a70635b05494f864e86144a2201275021014fbf08bafe2.
//
// Solidity: event OwnerSet(address indexed owner)
func (_Proxy *ProxyFilterer) WatchOwnerSet(opts *bind.WatchOpts, sink chan<- *ProxyOwnerSet, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Proxy.contract.WatchLogs(opts, "OwnerSet", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProxyOwnerSet)
				if err := _Proxy.contract.UnpackLog(event, "OwnerSet", log); err != nil {
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

// ParseOwnerSet is a log parse operation binding the contract event 0x50146d0e3c60aa1d17a70635b05494f864e86144a2201275021014fbf08bafe2.
//
// Solidity: event OwnerSet(address indexed owner)
func (_Proxy *ProxyFilterer) ParseOwnerSet(log types.Log) (*ProxyOwnerSet, error) {
	event := new(ProxyOwnerSet)
	if err := _Proxy.contract.UnpackLog(event, "OwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
