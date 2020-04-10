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

// AccountsABI is the input ABI used to generate the binding from.
const AccountsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedBy\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"AttestationSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"VoteSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"ValidatorSignerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSigner\",\"type\":\"address\"}],\"name\":\"AttestationSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSigner\",\"type\":\"address\"}],\"name\":\"VoteSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldSigner\",\"type\":\"address\"}],\"name\":\"ValidatorSignerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"}],\"name\":\"AccountDataEncryptionKeySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"name\",\"type\":\"string\"}],\"name\":\"AccountNameSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"metadataURL\",\"type\":\"string\"}],\"name\":\"AccountMetadataURLSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"walletAddress\",\"type\":\"address\"}],\"name\":\"AccountWalletAddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AccountCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"},{\"name\":\"walletAddress\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setWalletAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"}],\"name\":\"setAccountDataEncryptionKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"metadataURL\",\"type\":\"string\"}],\"name\":\"setMetadataURL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeVoteSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeValidatorSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"authorizeValidatorSignerWithPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"authorizeValidatorSignerWithKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeAttestationSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeVoteSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeValidatorSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAttestationSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"attestationSignerToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"validatorSignerToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"voteSignerToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"signerToAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getVoteSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidatorSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAttestationSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAuthorizedVoteSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAuthorizedValidatorSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasAuthorizedAttestationSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getName\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMetadataURL\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accountsToQuery\",\"type\":\"address[]\"}],\"name\":\"batchGetMetadataURL\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getDataEncryptionKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getWalletAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isAccount\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"isAuthorizedSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Accounts is an auto generated Go binding around an Ethereum contract.
type Accounts struct {
	AccountsCaller     // Read-only binding to the contract
	AccountsTransactor // Write-only binding to the contract
	AccountsFilterer   // Log filterer for contract events
}

// AccountsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AccountsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AccountsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AccountsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AccountsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AccountsSession struct {
	Contract     *Accounts         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AccountsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AccountsCallerSession struct {
	Contract *AccountsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AccountsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AccountsTransactorSession struct {
	Contract     *AccountsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AccountsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AccountsRaw struct {
	Contract *Accounts // Generic contract binding to access the raw methods on
}

// AccountsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AccountsCallerRaw struct {
	Contract *AccountsCaller // Generic read-only contract binding to access the raw methods on
}

// AccountsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AccountsTransactorRaw struct {
	Contract *AccountsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAccounts creates a new instance of Accounts, bound to a specific deployed contract.
func NewAccounts(address common.Address, backend bind.ContractBackend) (*Accounts, error) {
	contract, err := bindAccounts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Accounts{AccountsCaller: AccountsCaller{contract: contract}, AccountsTransactor: AccountsTransactor{contract: contract}, AccountsFilterer: AccountsFilterer{contract: contract}}, nil
}

// NewAccountsCaller creates a new read-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsCaller(address common.Address, caller bind.ContractCaller) (*AccountsCaller, error) {
	contract, err := bindAccounts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsCaller{contract: contract}, nil
}

// NewAccountsTransactor creates a new write-only instance of Accounts, bound to a specific deployed contract.
func NewAccountsTransactor(address common.Address, transactor bind.ContractTransactor) (*AccountsTransactor, error) {
	contract, err := bindAccounts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AccountsTransactor{contract: contract}, nil
}

// NewAccountsFilterer creates a new log filterer instance of Accounts, bound to a specific deployed contract.
func NewAccountsFilterer(address common.Address, filterer bind.ContractFilterer) (*AccountsFilterer, error) {
	contract, err := bindAccounts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AccountsFilterer{contract: contract}, nil
}

// bindAccounts binds a generic wrapper to an already deployed contract.
func bindAccounts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseAccountsABI parses the ABI
func ParseAccountsABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(AccountsABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.AccountsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.AccountsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Accounts *AccountsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Accounts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Accounts *AccountsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Accounts *AccountsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Accounts.Contract.contract.Transact(opts, method, params...)
}

// AttestationSignerToAccount is a free data retrieval call binding the contract method 0x7b2434cb.
//
// Solidity: function attestationSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCaller) AttestationSignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "attestationSignerToAccount", signer)
	return *ret0, err
}

// AttestationSignerToAccount is a free data retrieval call binding the contract method 0x7b2434cb.
//
// Solidity: function attestationSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsSession) AttestationSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.AttestationSignerToAccount(&_Accounts.CallOpts, signer)
}

// AttestationSignerToAccount is a free data retrieval call binding the contract method 0x7b2434cb.
//
// Solidity: function attestationSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCallerSession) AttestationSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.AttestationSignerToAccount(&_Accounts.CallOpts, signer)
}

// AuthorizedBy is a free data retrieval call binding the contract method 0xb5a664c2.
//
// Solidity: function authorizedBy(address ) constant returns(address)
func (_Accounts *AccountsCaller) AuthorizedBy(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "authorizedBy", arg0)
	return *ret0, err
}

// AuthorizedBy is a free data retrieval call binding the contract method 0xb5a664c2.
//
// Solidity: function authorizedBy(address ) constant returns(address)
func (_Accounts *AccountsSession) AuthorizedBy(arg0 common.Address) (common.Address, error) {
	return _Accounts.Contract.AuthorizedBy(&_Accounts.CallOpts, arg0)
}

// AuthorizedBy is a free data retrieval call binding the contract method 0xb5a664c2.
//
// Solidity: function authorizedBy(address ) constant returns(address)
func (_Accounts *AccountsCallerSession) AuthorizedBy(arg0 common.Address) (common.Address, error) {
	return _Accounts.Contract.AuthorizedBy(&_Accounts.CallOpts, arg0)
}

// BatchGetMetadataURL is a free data retrieval call binding the contract method 0x8adaf96f.
//
// Solidity: function batchGetMetadataURL(address[] accountsToQuery) constant returns(uint256[], bytes)
func (_Accounts *AccountsCaller) BatchGetMetadataURL(opts *bind.CallOpts, accountsToQuery []common.Address) ([]*big.Int, []byte, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Accounts.contract.Call(opts, out, "batchGetMetadataURL", accountsToQuery)
	return *ret0, *ret1, err
}

// BatchGetMetadataURL is a free data retrieval call binding the contract method 0x8adaf96f.
//
// Solidity: function batchGetMetadataURL(address[] accountsToQuery) constant returns(uint256[], bytes)
func (_Accounts *AccountsSession) BatchGetMetadataURL(accountsToQuery []common.Address) ([]*big.Int, []byte, error) {
	return _Accounts.Contract.BatchGetMetadataURL(&_Accounts.CallOpts, accountsToQuery)
}

// BatchGetMetadataURL is a free data retrieval call binding the contract method 0x8adaf96f.
//
// Solidity: function batchGetMetadataURL(address[] accountsToQuery) constant returns(uint256[], bytes)
func (_Accounts *AccountsCallerSession) BatchGetMetadataURL(accountsToQuery []common.Address) ([]*big.Int, []byte, error) {
	return _Accounts.Contract.BatchGetMetadataURL(&_Accounts.CallOpts, accountsToQuery)
}

// GetAttestationSigner is a free data retrieval call binding the contract method 0x61bab1ae.
//
// Solidity: function getAttestationSigner(address account) constant returns(address)
func (_Accounts *AccountsCaller) GetAttestationSigner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getAttestationSigner", account)
	return *ret0, err
}

// GetAttestationSigner is a free data retrieval call binding the contract method 0x61bab1ae.
//
// Solidity: function getAttestationSigner(address account) constant returns(address)
func (_Accounts *AccountsSession) GetAttestationSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetAttestationSigner(&_Accounts.CallOpts, account)
}

// GetAttestationSigner is a free data retrieval call binding the contract method 0x61bab1ae.
//
// Solidity: function getAttestationSigner(address account) constant returns(address)
func (_Accounts *AccountsCallerSession) GetAttestationSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetAttestationSigner(&_Accounts.CallOpts, account)
}

// GetDataEncryptionKey is a free data retrieval call binding the contract method 0xae32fa0e.
//
// Solidity: function getDataEncryptionKey(address account) constant returns(bytes)
func (_Accounts *AccountsCaller) GetDataEncryptionKey(opts *bind.CallOpts, account common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getDataEncryptionKey", account)
	return *ret0, err
}

// GetDataEncryptionKey is a free data retrieval call binding the contract method 0xae32fa0e.
//
// Solidity: function getDataEncryptionKey(address account) constant returns(bytes)
func (_Accounts *AccountsSession) GetDataEncryptionKey(account common.Address) ([]byte, error) {
	return _Accounts.Contract.GetDataEncryptionKey(&_Accounts.CallOpts, account)
}

// GetDataEncryptionKey is a free data retrieval call binding the contract method 0xae32fa0e.
//
// Solidity: function getDataEncryptionKey(address account) constant returns(bytes)
func (_Accounts *AccountsCallerSession) GetDataEncryptionKey(account common.Address) ([]byte, error) {
	return _Accounts.Contract.GetDataEncryptionKey(&_Accounts.CallOpts, account)
}

// GetMetadataURL is a free data retrieval call binding the contract method 0xa8ae1a3d.
//
// Solidity: function getMetadataURL(address account) constant returns(string)
func (_Accounts *AccountsCaller) GetMetadataURL(opts *bind.CallOpts, account common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getMetadataURL", account)
	return *ret0, err
}

// GetMetadataURL is a free data retrieval call binding the contract method 0xa8ae1a3d.
//
// Solidity: function getMetadataURL(address account) constant returns(string)
func (_Accounts *AccountsSession) GetMetadataURL(account common.Address) (string, error) {
	return _Accounts.Contract.GetMetadataURL(&_Accounts.CallOpts, account)
}

// GetMetadataURL is a free data retrieval call binding the contract method 0xa8ae1a3d.
//
// Solidity: function getMetadataURL(address account) constant returns(string)
func (_Accounts *AccountsCallerSession) GetMetadataURL(account common.Address) (string, error) {
	return _Accounts.Contract.GetMetadataURL(&_Accounts.CallOpts, account)
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address account) constant returns(string)
func (_Accounts *AccountsCaller) GetName(opts *bind.CallOpts, account common.Address) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getName", account)
	return *ret0, err
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address account) constant returns(string)
func (_Accounts *AccountsSession) GetName(account common.Address) (string, error) {
	return _Accounts.Contract.GetName(&_Accounts.CallOpts, account)
}

// GetName is a free data retrieval call binding the contract method 0x5fd4b08a.
//
// Solidity: function getName(address account) constant returns(string)
func (_Accounts *AccountsCallerSession) GetName(account common.Address) (string, error) {
	return _Accounts.Contract.GetName(&_Accounts.CallOpts, account)
}

// GetValidatorSigner is a free data retrieval call binding the contract method 0x4ce38b5f.
//
// Solidity: function getValidatorSigner(address account) constant returns(address)
func (_Accounts *AccountsCaller) GetValidatorSigner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getValidatorSigner", account)
	return *ret0, err
}

// GetValidatorSigner is a free data retrieval call binding the contract method 0x4ce38b5f.
//
// Solidity: function getValidatorSigner(address account) constant returns(address)
func (_Accounts *AccountsSession) GetValidatorSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetValidatorSigner(&_Accounts.CallOpts, account)
}

// GetValidatorSigner is a free data retrieval call binding the contract method 0x4ce38b5f.
//
// Solidity: function getValidatorSigner(address account) constant returns(address)
func (_Accounts *AccountsCallerSession) GetValidatorSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetValidatorSigner(&_Accounts.CallOpts, account)
}

// GetVoteSigner is a free data retrieval call binding the contract method 0x41ddd880.
//
// Solidity: function getVoteSigner(address account) constant returns(address)
func (_Accounts *AccountsCaller) GetVoteSigner(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getVoteSigner", account)
	return *ret0, err
}

// GetVoteSigner is a free data retrieval call binding the contract method 0x41ddd880.
//
// Solidity: function getVoteSigner(address account) constant returns(address)
func (_Accounts *AccountsSession) GetVoteSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetVoteSigner(&_Accounts.CallOpts, account)
}

// GetVoteSigner is a free data retrieval call binding the contract method 0x41ddd880.
//
// Solidity: function getVoteSigner(address account) constant returns(address)
func (_Accounts *AccountsCallerSession) GetVoteSigner(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetVoteSigner(&_Accounts.CallOpts, account)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address account) constant returns(address)
func (_Accounts *AccountsCaller) GetWalletAddress(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "getWalletAddress", account)
	return *ret0, err
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address account) constant returns(address)
func (_Accounts *AccountsSession) GetWalletAddress(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetWalletAddress(&_Accounts.CallOpts, account)
}

// GetWalletAddress is a free data retrieval call binding the contract method 0x1fd9afa5.
//
// Solidity: function getWalletAddress(address account) constant returns(address)
func (_Accounts *AccountsCallerSession) GetWalletAddress(account common.Address) (common.Address, error) {
	return _Accounts.Contract.GetWalletAddress(&_Accounts.CallOpts, account)
}

// HasAuthorizedAttestationSigner is a free data retrieval call binding the contract method 0xc2e0ee20.
//
// Solidity: function hasAuthorizedAttestationSigner(address account) constant returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedAttestationSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "hasAuthorizedAttestationSigner", account)
	return *ret0, err
}

// HasAuthorizedAttestationSigner is a free data retrieval call binding the contract method 0xc2e0ee20.
//
// Solidity: function hasAuthorizedAttestationSigner(address account) constant returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedAttestationSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedAttestationSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedAttestationSigner is a free data retrieval call binding the contract method 0xc2e0ee20.
//
// Solidity: function hasAuthorizedAttestationSigner(address account) constant returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedAttestationSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedAttestationSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedValidatorSigner is a free data retrieval call binding the contract method 0x0127dbed.
//
// Solidity: function hasAuthorizedValidatorSigner(address account) constant returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedValidatorSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "hasAuthorizedValidatorSigner", account)
	return *ret0, err
}

// HasAuthorizedValidatorSigner is a free data retrieval call binding the contract method 0x0127dbed.
//
// Solidity: function hasAuthorizedValidatorSigner(address account) constant returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedValidatorSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedValidatorSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedValidatorSigner is a free data retrieval call binding the contract method 0x0127dbed.
//
// Solidity: function hasAuthorizedValidatorSigner(address account) constant returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedValidatorSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedValidatorSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedVoteSigner is a free data retrieval call binding the contract method 0x614ed493.
//
// Solidity: function hasAuthorizedVoteSigner(address account) constant returns(bool)
func (_Accounts *AccountsCaller) HasAuthorizedVoteSigner(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "hasAuthorizedVoteSigner", account)
	return *ret0, err
}

// HasAuthorizedVoteSigner is a free data retrieval call binding the contract method 0x614ed493.
//
// Solidity: function hasAuthorizedVoteSigner(address account) constant returns(bool)
func (_Accounts *AccountsSession) HasAuthorizedVoteSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedVoteSigner(&_Accounts.CallOpts, account)
}

// HasAuthorizedVoteSigner is a free data retrieval call binding the contract method 0x614ed493.
//
// Solidity: function hasAuthorizedVoteSigner(address account) constant returns(bool)
func (_Accounts *AccountsCallerSession) HasAuthorizedVoteSigner(account common.Address) (bool, error) {
	return _Accounts.Contract.HasAuthorizedVoteSigner(&_Accounts.CallOpts, account)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Accounts *AccountsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Accounts *AccountsSession) Initialized() (bool, error) {
	return _Accounts.Contract.Initialized(&_Accounts.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Accounts *AccountsCallerSession) Initialized() (bool, error) {
	return _Accounts.Contract.Initialized(&_Accounts.CallOpts)
}

// IsAccount is a free data retrieval call binding the contract method 0x25ca4c9c.
//
// Solidity: function isAccount(address account) constant returns(bool)
func (_Accounts *AccountsCaller) IsAccount(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isAccount", account)
	return *ret0, err
}

// IsAccount is a free data retrieval call binding the contract method 0x25ca4c9c.
//
// Solidity: function isAccount(address account) constant returns(bool)
func (_Accounts *AccountsSession) IsAccount(account common.Address) (bool, error) {
	return _Accounts.Contract.IsAccount(&_Accounts.CallOpts, account)
}

// IsAccount is a free data retrieval call binding the contract method 0x25ca4c9c.
//
// Solidity: function isAccount(address account) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsAccount(account common.Address) (bool, error) {
	return _Accounts.Contract.IsAccount(&_Accounts.CallOpts, account)
}

// IsAuthorizedSigner is a free data retrieval call binding the contract method 0x49045e16.
//
// Solidity: function isAuthorizedSigner(address signer) constant returns(bool)
func (_Accounts *AccountsCaller) IsAuthorizedSigner(opts *bind.CallOpts, signer common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isAuthorizedSigner", signer)
	return *ret0, err
}

// IsAuthorizedSigner is a free data retrieval call binding the contract method 0x49045e16.
//
// Solidity: function isAuthorizedSigner(address signer) constant returns(bool)
func (_Accounts *AccountsSession) IsAuthorizedSigner(signer common.Address) (bool, error) {
	return _Accounts.Contract.IsAuthorizedSigner(&_Accounts.CallOpts, signer)
}

// IsAuthorizedSigner is a free data retrieval call binding the contract method 0x49045e16.
//
// Solidity: function isAuthorizedSigner(address signer) constant returns(bool)
func (_Accounts *AccountsCallerSession) IsAuthorizedSigner(signer common.Address) (bool, error) {
	return _Accounts.Contract.IsAuthorizedSigner(&_Accounts.CallOpts, signer)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Accounts *AccountsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Accounts *AccountsSession) IsOwner() (bool, error) {
	return _Accounts.Contract.IsOwner(&_Accounts.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Accounts *AccountsCallerSession) IsOwner() (bool, error) {
	return _Accounts.Contract.IsOwner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Accounts *AccountsCallerSession) Owner() (common.Address, error) {
	return _Accounts.Contract.Owner(&_Accounts.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Accounts *AccountsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Accounts *AccountsSession) Registry() (common.Address, error) {
	return _Accounts.Contract.Registry(&_Accounts.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Accounts *AccountsCallerSession) Registry() (common.Address, error) {
	return _Accounts.Contract.Registry(&_Accounts.CallOpts)
}

// SignerToAccount is a free data retrieval call binding the contract method 0x93c5c487.
//
// Solidity: function signerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCaller) SignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "signerToAccount", signer)
	return *ret0, err
}

// SignerToAccount is a free data retrieval call binding the contract method 0x93c5c487.
//
// Solidity: function signerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsSession) SignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.SignerToAccount(&_Accounts.CallOpts, signer)
}

// SignerToAccount is a free data retrieval call binding the contract method 0x93c5c487.
//
// Solidity: function signerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCallerSession) SignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.SignerToAccount(&_Accounts.CallOpts, signer)
}

// ValidatorSignerToAccount is a free data retrieval call binding the contract method 0x64439b43.
//
// Solidity: function validatorSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCaller) ValidatorSignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "validatorSignerToAccount", signer)
	return *ret0, err
}

// ValidatorSignerToAccount is a free data retrieval call binding the contract method 0x64439b43.
//
// Solidity: function validatorSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsSession) ValidatorSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.ValidatorSignerToAccount(&_Accounts.CallOpts, signer)
}

// ValidatorSignerToAccount is a free data retrieval call binding the contract method 0x64439b43.
//
// Solidity: function validatorSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCallerSession) ValidatorSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.ValidatorSignerToAccount(&_Accounts.CallOpts, signer)
}

// VoteSignerToAccount is a free data retrieval call binding the contract method 0x6642d594.
//
// Solidity: function voteSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCaller) VoteSignerToAccount(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Accounts.contract.Call(opts, out, "voteSignerToAccount", signer)
	return *ret0, err
}

// VoteSignerToAccount is a free data retrieval call binding the contract method 0x6642d594.
//
// Solidity: function voteSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsSession) VoteSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.VoteSignerToAccount(&_Accounts.CallOpts, signer)
}

// VoteSignerToAccount is a free data retrieval call binding the contract method 0x6642d594.
//
// Solidity: function voteSignerToAccount(address signer) constant returns(address)
func (_Accounts *AccountsCallerSession) VoteSignerToAccount(signer common.Address) (common.Address, error) {
	return _Accounts.Contract.VoteSignerToAccount(&_Accounts.CallOpts, signer)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeAttestationSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeAttestationSigner", signer, v, r, s)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeAttestationSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeAttestationSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeAttestationSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeAttestationSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeValidatorSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeValidatorSigner", signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeValidatorSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeValidatorSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_Accounts *AccountsTransactor) AuthorizeValidatorSignerWithKeys(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeValidatorSignerWithKeys", signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_Accounts *AccountsSession) AuthorizeValidatorSignerWithKeys(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithKeys(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeValidatorSignerWithKeys(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithKeys(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_Accounts *AccountsTransactor) AuthorizeValidatorSignerWithPublicKey(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeValidatorSignerWithPublicKey", signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_Accounts *AccountsSession) AuthorizeValidatorSignerWithPublicKey(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithPublicKey(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeValidatorSignerWithPublicKey(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeValidatorSignerWithPublicKey(&_Accounts.TransactOpts, signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) AuthorizeVoteSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "authorizeVoteSigner", signer, v, r, s)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) AuthorizeVoteSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeVoteSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) AuthorizeVoteSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.AuthorizeVoteSigner(&_Accounts.TransactOpts, signer, v, r, s)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Accounts *AccountsTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Accounts *AccountsSession) CreateAccount() (*types.Transaction, error) {
	return _Accounts.Contract.CreateAccount(&_Accounts.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns(bool)
func (_Accounts *AccountsTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _Accounts.Contract.CreateAccount(&_Accounts.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Accounts *AccountsTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "initialize", registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Accounts *AccountsSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address registryAddress) returns()
func (_Accounts *AccountsTransactorSession) Initialize(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.Initialize(&_Accounts.TransactOpts, registryAddress)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x760fbbb2.
//
// Solidity: function removeAttestationSigner() returns()
func (_Accounts *AccountsTransactor) RemoveAttestationSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeAttestationSigner")
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x760fbbb2.
//
// Solidity: function removeAttestationSigner() returns()
func (_Accounts *AccountsSession) RemoveAttestationSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveAttestationSigner(&_Accounts.TransactOpts)
}

// RemoveAttestationSigner is a paid mutator transaction binding the contract method 0x760fbbb2.
//
// Solidity: function removeAttestationSigner() returns()
func (_Accounts *AccountsTransactorSession) RemoveAttestationSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveAttestationSigner(&_Accounts.TransactOpts)
}

// RemoveValidatorSigner is a paid mutator transaction binding the contract method 0xa5ec94f9.
//
// Solidity: function removeValidatorSigner() returns()
func (_Accounts *AccountsTransactor) RemoveValidatorSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeValidatorSigner")
}

// RemoveValidatorSigner is a paid mutator transaction binding the contract method 0xa5ec94f9.
//
// Solidity: function removeValidatorSigner() returns()
func (_Accounts *AccountsSession) RemoveValidatorSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveValidatorSigner(&_Accounts.TransactOpts)
}

// RemoveValidatorSigner is a paid mutator transaction binding the contract method 0xa5ec94f9.
//
// Solidity: function removeValidatorSigner() returns()
func (_Accounts *AccountsTransactorSession) RemoveValidatorSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveValidatorSigner(&_Accounts.TransactOpts)
}

// RemoveVoteSigner is a paid mutator transaction binding the contract method 0x10c504b5.
//
// Solidity: function removeVoteSigner() returns()
func (_Accounts *AccountsTransactor) RemoveVoteSigner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "removeVoteSigner")
}

// RemoveVoteSigner is a paid mutator transaction binding the contract method 0x10c504b5.
//
// Solidity: function removeVoteSigner() returns()
func (_Accounts *AccountsSession) RemoveVoteSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveVoteSigner(&_Accounts.TransactOpts)
}

// RemoveVoteSigner is a paid mutator transaction binding the contract method 0x10c504b5.
//
// Solidity: function removeVoteSigner() returns()
func (_Accounts *AccountsTransactorSession) RemoveVoteSigner() (*types.Transaction, error) {
	return _Accounts.Contract.RemoveVoteSigner(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Accounts *AccountsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Accounts.Contract.RenounceOwnership(&_Accounts.TransactOpts)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) SetAccount(opts *bind.TransactOpts, name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setAccount", name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) SetAccount(name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccount(&_Accounts.TransactOpts, name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) SetAccount(name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccount(&_Accounts.TransactOpts, name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_Accounts *AccountsTransactor) SetAccountDataEncryptionKey(opts *bind.TransactOpts, dataEncryptionKey []byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setAccountDataEncryptionKey", dataEncryptionKey)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_Accounts *AccountsSession) SetAccountDataEncryptionKey(dataEncryptionKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccountDataEncryptionKey(&_Accounts.TransactOpts, dataEncryptionKey)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_Accounts *AccountsTransactorSession) SetAccountDataEncryptionKey(dataEncryptionKey []byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetAccountDataEncryptionKey(&_Accounts.TransactOpts, dataEncryptionKey)
}

// SetMetadataURL is a paid mutator transaction binding the contract method 0x747daec5.
//
// Solidity: function setMetadataURL(string metadataURL) returns()
func (_Accounts *AccountsTransactor) SetMetadataURL(opts *bind.TransactOpts, metadataURL string) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setMetadataURL", metadataURL)
}

// SetMetadataURL is a paid mutator transaction binding the contract method 0x747daec5.
//
// Solidity: function setMetadataURL(string metadataURL) returns()
func (_Accounts *AccountsSession) SetMetadataURL(metadataURL string) (*types.Transaction, error) {
	return _Accounts.Contract.SetMetadataURL(&_Accounts.TransactOpts, metadataURL)
}

// SetMetadataURL is a paid mutator transaction binding the contract method 0x747daec5.
//
// Solidity: function setMetadataURL(string metadataURL) returns()
func (_Accounts *AccountsTransactorSession) SetMetadataURL(metadataURL string) (*types.Transaction, error) {
	return _Accounts.Contract.SetMetadataURL(&_Accounts.TransactOpts, metadataURL)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Accounts *AccountsTransactor) SetName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setName", name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Accounts *AccountsSession) SetName(name string) (*types.Transaction, error) {
	return _Accounts.Contract.SetName(&_Accounts.TransactOpts, name)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string name) returns()
func (_Accounts *AccountsTransactorSession) SetName(name string) (*types.Transaction, error) {
	return _Accounts.Contract.SetName(&_Accounts.TransactOpts, name)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Accounts *AccountsTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Accounts *AccountsSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetRegistry(&_Accounts.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Accounts *AccountsTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.SetRegistry(&_Accounts.TransactOpts, registryAddress)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0x9cafb2a1.
//
// Solidity: function setWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactor) SetWalletAddress(opts *bind.TransactOpts, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "setWalletAddress", walletAddress, v, r, s)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0x9cafb2a1.
//
// Solidity: function setWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsSession) SetWalletAddress(walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetWalletAddress(&_Accounts.TransactOpts, walletAddress, v, r, s)
}

// SetWalletAddress is a paid mutator transaction binding the contract method 0x9cafb2a1.
//
// Solidity: function setWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_Accounts *AccountsTransactorSession) SetWalletAddress(walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Accounts.Contract.SetWalletAddress(&_Accounts.TransactOpts, walletAddress, v, r, s)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Accounts *AccountsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Accounts.Contract.TransferOwnership(&_Accounts.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Accounts *AccountsFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Accounts.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AccountCreated":
		event, err = _Accounts.ParseAccountCreated(log)
	case "AccountDataEncryptionKeySet":
		event, err = _Accounts.ParseAccountDataEncryptionKeySet(log)
	case "AccountMetadataURLSet":
		event, err = _Accounts.ParseAccountMetadataURLSet(log)
	case "AccountNameSet":
		event, err = _Accounts.ParseAccountNameSet(log)
	case "AccountWalletAddressSet":
		event, err = _Accounts.ParseAccountWalletAddressSet(log)
	case "AttestationSignerAuthorized":
		event, err = _Accounts.ParseAttestationSignerAuthorized(log)
	case "AttestationSignerRemoved":
		event, err = _Accounts.ParseAttestationSignerRemoved(log)
	case "OwnershipTransferred":
		event, err = _Accounts.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Accounts.ParseRegistrySet(log)
	case "ValidatorSignerAuthorized":
		event, err = _Accounts.ParseValidatorSignerAuthorized(log)
	case "ValidatorSignerRemoved":
		event, err = _Accounts.ParseValidatorSignerRemoved(log)
	case "VoteSignerAuthorized":
		event, err = _Accounts.ParseVoteSignerAuthorized(log)
	case "VoteSignerRemoved":
		event, err = _Accounts.ParseVoteSignerRemoved(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// AccountsAccountCreatedIterator is returned from FilterAccountCreated and is used to iterate over the raw logs and unpacked data for AccountCreated events raised by the Accounts contract.
type AccountsAccountCreatedIterator struct {
	Event *AccountsAccountCreated // Event containing the contract specifics and raw log

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
func (it *AccountsAccountCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountCreated)
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
		it.Event = new(AccountsAccountCreated)
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
func (it *AccountsAccountCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountCreated represents a AccountCreated event raised by the Accounts contract.
type AccountsAccountCreated struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountCreated is a free log retrieval operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed account)
func (_Accounts *AccountsFilterer) FilterAccountCreated(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountCreatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountCreatedIterator{contract: _Accounts.contract, event: "AccountCreated", logs: logs, sub: sub}, nil
}

// WatchAccountCreated is a free log subscription operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed account)
func (_Accounts *AccountsFilterer) WatchAccountCreated(opts *bind.WatchOpts, sink chan<- *AccountsAccountCreated, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountCreated", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountCreated)
				if err := _Accounts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
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

// ParseAccountCreated is a log parse operation binding the contract event 0x805996f252884581e2f74cf3d2b03564d5ec26ccc90850ae12653dc1b72d1fa2.
//
// Solidity: event AccountCreated(address indexed account)
func (_Accounts *AccountsFilterer) ParseAccountCreated(log types.Log) (*AccountsAccountCreated, error) {
	event := new(AccountsAccountCreated)
	if err := _Accounts.contract.UnpackLog(event, "AccountCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsAccountDataEncryptionKeySetIterator is returned from FilterAccountDataEncryptionKeySet and is used to iterate over the raw logs and unpacked data for AccountDataEncryptionKeySet events raised by the Accounts contract.
type AccountsAccountDataEncryptionKeySetIterator struct {
	Event *AccountsAccountDataEncryptionKeySet // Event containing the contract specifics and raw log

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
func (it *AccountsAccountDataEncryptionKeySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountDataEncryptionKeySet)
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
		it.Event = new(AccountsAccountDataEncryptionKeySet)
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
func (it *AccountsAccountDataEncryptionKeySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountDataEncryptionKeySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountDataEncryptionKeySet represents a AccountDataEncryptionKeySet event raised by the Accounts contract.
type AccountsAccountDataEncryptionKeySet struct {
	Account           common.Address
	DataEncryptionKey []byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterAccountDataEncryptionKeySet is a free log retrieval operation binding the contract event 0x43fdefe0a824cb0e3bbaf9c4bc97669187996136fe9282382baf10787f0d808d.
//
// Solidity: event AccountDataEncryptionKeySet(address indexed account, bytes dataEncryptionKey)
func (_Accounts *AccountsFilterer) FilterAccountDataEncryptionKeySet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountDataEncryptionKeySetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountDataEncryptionKeySet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountDataEncryptionKeySetIterator{contract: _Accounts.contract, event: "AccountDataEncryptionKeySet", logs: logs, sub: sub}, nil
}

// WatchAccountDataEncryptionKeySet is a free log subscription operation binding the contract event 0x43fdefe0a824cb0e3bbaf9c4bc97669187996136fe9282382baf10787f0d808d.
//
// Solidity: event AccountDataEncryptionKeySet(address indexed account, bytes dataEncryptionKey)
func (_Accounts *AccountsFilterer) WatchAccountDataEncryptionKeySet(opts *bind.WatchOpts, sink chan<- *AccountsAccountDataEncryptionKeySet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountDataEncryptionKeySet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountDataEncryptionKeySet)
				if err := _Accounts.contract.UnpackLog(event, "AccountDataEncryptionKeySet", log); err != nil {
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

// ParseAccountDataEncryptionKeySet is a log parse operation binding the contract event 0x43fdefe0a824cb0e3bbaf9c4bc97669187996136fe9282382baf10787f0d808d.
//
// Solidity: event AccountDataEncryptionKeySet(address indexed account, bytes dataEncryptionKey)
func (_Accounts *AccountsFilterer) ParseAccountDataEncryptionKeySet(log types.Log) (*AccountsAccountDataEncryptionKeySet, error) {
	event := new(AccountsAccountDataEncryptionKeySet)
	if err := _Accounts.contract.UnpackLog(event, "AccountDataEncryptionKeySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsAccountMetadataURLSetIterator is returned from FilterAccountMetadataURLSet and is used to iterate over the raw logs and unpacked data for AccountMetadataURLSet events raised by the Accounts contract.
type AccountsAccountMetadataURLSetIterator struct {
	Event *AccountsAccountMetadataURLSet // Event containing the contract specifics and raw log

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
func (it *AccountsAccountMetadataURLSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountMetadataURLSet)
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
		it.Event = new(AccountsAccountMetadataURLSet)
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
func (it *AccountsAccountMetadataURLSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountMetadataURLSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountMetadataURLSet represents a AccountMetadataURLSet event raised by the Accounts contract.
type AccountsAccountMetadataURLSet struct {
	Account     common.Address
	MetadataURL string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAccountMetadataURLSet is a free log retrieval operation binding the contract event 0x0b5629fec5b6b5a1c2cfe0de7495111627a8cf297dced72e0669527425d3f01b.
//
// Solidity: event AccountMetadataURLSet(address indexed account, string metadataURL)
func (_Accounts *AccountsFilterer) FilterAccountMetadataURLSet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountMetadataURLSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountMetadataURLSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountMetadataURLSetIterator{contract: _Accounts.contract, event: "AccountMetadataURLSet", logs: logs, sub: sub}, nil
}

// WatchAccountMetadataURLSet is a free log subscription operation binding the contract event 0x0b5629fec5b6b5a1c2cfe0de7495111627a8cf297dced72e0669527425d3f01b.
//
// Solidity: event AccountMetadataURLSet(address indexed account, string metadataURL)
func (_Accounts *AccountsFilterer) WatchAccountMetadataURLSet(opts *bind.WatchOpts, sink chan<- *AccountsAccountMetadataURLSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountMetadataURLSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountMetadataURLSet)
				if err := _Accounts.contract.UnpackLog(event, "AccountMetadataURLSet", log); err != nil {
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

// ParseAccountMetadataURLSet is a log parse operation binding the contract event 0x0b5629fec5b6b5a1c2cfe0de7495111627a8cf297dced72e0669527425d3f01b.
//
// Solidity: event AccountMetadataURLSet(address indexed account, string metadataURL)
func (_Accounts *AccountsFilterer) ParseAccountMetadataURLSet(log types.Log) (*AccountsAccountMetadataURLSet, error) {
	event := new(AccountsAccountMetadataURLSet)
	if err := _Accounts.contract.UnpackLog(event, "AccountMetadataURLSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsAccountNameSetIterator is returned from FilterAccountNameSet and is used to iterate over the raw logs and unpacked data for AccountNameSet events raised by the Accounts contract.
type AccountsAccountNameSetIterator struct {
	Event *AccountsAccountNameSet // Event containing the contract specifics and raw log

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
func (it *AccountsAccountNameSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountNameSet)
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
		it.Event = new(AccountsAccountNameSet)
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
func (it *AccountsAccountNameSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountNameSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountNameSet represents a AccountNameSet event raised by the Accounts contract.
type AccountsAccountNameSet struct {
	Account common.Address
	Name    string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountNameSet is a free log retrieval operation binding the contract event 0xa6e2c5a23bb917ba0a584c4b250257ddad698685829b66a8813c004b39934fe4.
//
// Solidity: event AccountNameSet(address indexed account, string name)
func (_Accounts *AccountsFilterer) FilterAccountNameSet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountNameSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountNameSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountNameSetIterator{contract: _Accounts.contract, event: "AccountNameSet", logs: logs, sub: sub}, nil
}

// WatchAccountNameSet is a free log subscription operation binding the contract event 0xa6e2c5a23bb917ba0a584c4b250257ddad698685829b66a8813c004b39934fe4.
//
// Solidity: event AccountNameSet(address indexed account, string name)
func (_Accounts *AccountsFilterer) WatchAccountNameSet(opts *bind.WatchOpts, sink chan<- *AccountsAccountNameSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountNameSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountNameSet)
				if err := _Accounts.contract.UnpackLog(event, "AccountNameSet", log); err != nil {
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

// ParseAccountNameSet is a log parse operation binding the contract event 0xa6e2c5a23bb917ba0a584c4b250257ddad698685829b66a8813c004b39934fe4.
//
// Solidity: event AccountNameSet(address indexed account, string name)
func (_Accounts *AccountsFilterer) ParseAccountNameSet(log types.Log) (*AccountsAccountNameSet, error) {
	event := new(AccountsAccountNameSet)
	if err := _Accounts.contract.UnpackLog(event, "AccountNameSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsAccountWalletAddressSetIterator is returned from FilterAccountWalletAddressSet and is used to iterate over the raw logs and unpacked data for AccountWalletAddressSet events raised by the Accounts contract.
type AccountsAccountWalletAddressSetIterator struct {
	Event *AccountsAccountWalletAddressSet // Event containing the contract specifics and raw log

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
func (it *AccountsAccountWalletAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAccountWalletAddressSet)
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
		it.Event = new(AccountsAccountWalletAddressSet)
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
func (it *AccountsAccountWalletAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAccountWalletAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAccountWalletAddressSet represents a AccountWalletAddressSet event raised by the Accounts contract.
type AccountsAccountWalletAddressSet struct {
	Account       common.Address
	WalletAddress common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAccountWalletAddressSet is a free log retrieval operation binding the contract event 0xf81d74398fd47e35c36b714019df15f200f623dde569b5b531d6a0b4da5c5f26.
//
// Solidity: event AccountWalletAddressSet(address indexed account, address walletAddress)
func (_Accounts *AccountsFilterer) FilterAccountWalletAddressSet(opts *bind.FilterOpts, account []common.Address) (*AccountsAccountWalletAddressSetIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AccountWalletAddressSet", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAccountWalletAddressSetIterator{contract: _Accounts.contract, event: "AccountWalletAddressSet", logs: logs, sub: sub}, nil
}

// WatchAccountWalletAddressSet is a free log subscription operation binding the contract event 0xf81d74398fd47e35c36b714019df15f200f623dde569b5b531d6a0b4da5c5f26.
//
// Solidity: event AccountWalletAddressSet(address indexed account, address walletAddress)
func (_Accounts *AccountsFilterer) WatchAccountWalletAddressSet(opts *bind.WatchOpts, sink chan<- *AccountsAccountWalletAddressSet, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AccountWalletAddressSet", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAccountWalletAddressSet)
				if err := _Accounts.contract.UnpackLog(event, "AccountWalletAddressSet", log); err != nil {
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

// ParseAccountWalletAddressSet is a log parse operation binding the contract event 0xf81d74398fd47e35c36b714019df15f200f623dde569b5b531d6a0b4da5c5f26.
//
// Solidity: event AccountWalletAddressSet(address indexed account, address walletAddress)
func (_Accounts *AccountsFilterer) ParseAccountWalletAddressSet(log types.Log) (*AccountsAccountWalletAddressSet, error) {
	event := new(AccountsAccountWalletAddressSet)
	if err := _Accounts.contract.UnpackLog(event, "AccountWalletAddressSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsAttestationSignerAuthorizedIterator is returned from FilterAttestationSignerAuthorized and is used to iterate over the raw logs and unpacked data for AttestationSignerAuthorized events raised by the Accounts contract.
type AccountsAttestationSignerAuthorizedIterator struct {
	Event *AccountsAttestationSignerAuthorized // Event containing the contract specifics and raw log

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
func (it *AccountsAttestationSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAttestationSignerAuthorized)
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
		it.Event = new(AccountsAttestationSignerAuthorized)
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
func (it *AccountsAttestationSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAttestationSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAttestationSignerAuthorized represents a AttestationSignerAuthorized event raised by the Accounts contract.
type AccountsAttestationSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAttestationSignerAuthorized is a free log retrieval operation binding the contract event 0x9dfbc5a621c3e2d0d83beee687a17dfc796bbce2118793e5e254409bb265ca0b.
//
// Solidity: event AttestationSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) FilterAttestationSignerAuthorized(opts *bind.FilterOpts, account []common.Address) (*AccountsAttestationSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AttestationSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAttestationSignerAuthorizedIterator{contract: _Accounts.contract, event: "AttestationSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchAttestationSignerAuthorized is a free log subscription operation binding the contract event 0x9dfbc5a621c3e2d0d83beee687a17dfc796bbce2118793e5e254409bb265ca0b.
//
// Solidity: event AttestationSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) WatchAttestationSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsAttestationSignerAuthorized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AttestationSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAttestationSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "AttestationSignerAuthorized", log); err != nil {
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

// ParseAttestationSignerAuthorized is a log parse operation binding the contract event 0x9dfbc5a621c3e2d0d83beee687a17dfc796bbce2118793e5e254409bb265ca0b.
//
// Solidity: event AttestationSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) ParseAttestationSignerAuthorized(log types.Log) (*AccountsAttestationSignerAuthorized, error) {
	event := new(AccountsAttestationSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "AttestationSignerAuthorized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsAttestationSignerRemovedIterator is returned from FilterAttestationSignerRemoved and is used to iterate over the raw logs and unpacked data for AttestationSignerRemoved events raised by the Accounts contract.
type AccountsAttestationSignerRemovedIterator struct {
	Event *AccountsAttestationSignerRemoved // Event containing the contract specifics and raw log

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
func (it *AccountsAttestationSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsAttestationSignerRemoved)
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
		it.Event = new(AccountsAttestationSignerRemoved)
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
func (it *AccountsAttestationSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsAttestationSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsAttestationSignerRemoved represents a AttestationSignerRemoved event raised by the Accounts contract.
type AccountsAttestationSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAttestationSignerRemoved is a free log retrieval operation binding the contract event 0x14670729407debb6ed03d885f8ba57155de89ce39bf17127ae4900ec7c2ad103.
//
// Solidity: event AttestationSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) FilterAttestationSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsAttestationSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "AttestationSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsAttestationSignerRemovedIterator{contract: _Accounts.contract, event: "AttestationSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchAttestationSignerRemoved is a free log subscription operation binding the contract event 0x14670729407debb6ed03d885f8ba57155de89ce39bf17127ae4900ec7c2ad103.
//
// Solidity: event AttestationSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) WatchAttestationSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsAttestationSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "AttestationSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsAttestationSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "AttestationSignerRemoved", log); err != nil {
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

// ParseAttestationSignerRemoved is a log parse operation binding the contract event 0x14670729407debb6ed03d885f8ba57155de89ce39bf17127ae4900ec7c2ad103.
//
// Solidity: event AttestationSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) ParseAttestationSignerRemoved(log types.Log) (*AccountsAttestationSignerRemoved, error) {
	event := new(AccountsAttestationSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "AttestationSignerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Accounts contract.
type AccountsOwnershipTransferredIterator struct {
	Event *AccountsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AccountsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsOwnershipTransferred)
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
		it.Event = new(AccountsOwnershipTransferred)
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
func (it *AccountsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsOwnershipTransferred represents a OwnershipTransferred event raised by the Accounts contract.
type AccountsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AccountsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AccountsOwnershipTransferredIterator{contract: _Accounts.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Accounts *AccountsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AccountsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsOwnershipTransferred)
				if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseOwnershipTransferred(log types.Log) (*AccountsOwnershipTransferred, error) {
	event := new(AccountsOwnershipTransferred)
	if err := _Accounts.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Accounts contract.
type AccountsRegistrySetIterator struct {
	Event *AccountsRegistrySet // Event containing the contract specifics and raw log

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
func (it *AccountsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsRegistrySet)
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
		it.Event = new(AccountsRegistrySet)
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
func (it *AccountsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsRegistrySet represents a RegistrySet event raised by the Accounts contract.
type AccountsRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Accounts *AccountsFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*AccountsRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &AccountsRegistrySetIterator{contract: _Accounts.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Accounts *AccountsFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *AccountsRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsRegistrySet)
				if err := _Accounts.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Accounts *AccountsFilterer) ParseRegistrySet(log types.Log) (*AccountsRegistrySet, error) {
	event := new(AccountsRegistrySet)
	if err := _Accounts.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsValidatorSignerAuthorizedIterator is returned from FilterValidatorSignerAuthorized and is used to iterate over the raw logs and unpacked data for ValidatorSignerAuthorized events raised by the Accounts contract.
type AccountsValidatorSignerAuthorizedIterator struct {
	Event *AccountsValidatorSignerAuthorized // Event containing the contract specifics and raw log

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
func (it *AccountsValidatorSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsValidatorSignerAuthorized)
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
		it.Event = new(AccountsValidatorSignerAuthorized)
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
func (it *AccountsValidatorSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsValidatorSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsValidatorSignerAuthorized represents a ValidatorSignerAuthorized event raised by the Accounts contract.
type AccountsValidatorSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorSignerAuthorized is a free log retrieval operation binding the contract event 0x16e382723fb40543364faf68863212ba253a099607bf6d3a5b47e50a8bf94943.
//
// Solidity: event ValidatorSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) FilterValidatorSignerAuthorized(opts *bind.FilterOpts, account []common.Address) (*AccountsValidatorSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "ValidatorSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsValidatorSignerAuthorizedIterator{contract: _Accounts.contract, event: "ValidatorSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchValidatorSignerAuthorized is a free log subscription operation binding the contract event 0x16e382723fb40543364faf68863212ba253a099607bf6d3a5b47e50a8bf94943.
//
// Solidity: event ValidatorSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) WatchValidatorSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsValidatorSignerAuthorized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "ValidatorSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsValidatorSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerAuthorized", log); err != nil {
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

// ParseValidatorSignerAuthorized is a log parse operation binding the contract event 0x16e382723fb40543364faf68863212ba253a099607bf6d3a5b47e50a8bf94943.
//
// Solidity: event ValidatorSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) ParseValidatorSignerAuthorized(log types.Log) (*AccountsValidatorSignerAuthorized, error) {
	event := new(AccountsValidatorSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerAuthorized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsValidatorSignerRemovedIterator is returned from FilterValidatorSignerRemoved and is used to iterate over the raw logs and unpacked data for ValidatorSignerRemoved events raised by the Accounts contract.
type AccountsValidatorSignerRemovedIterator struct {
	Event *AccountsValidatorSignerRemoved // Event containing the contract specifics and raw log

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
func (it *AccountsValidatorSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsValidatorSignerRemoved)
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
		it.Event = new(AccountsValidatorSignerRemoved)
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
func (it *AccountsValidatorSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsValidatorSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsValidatorSignerRemoved represents a ValidatorSignerRemoved event raised by the Accounts contract.
type AccountsValidatorSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSignerRemoved is a free log retrieval operation binding the contract event 0xa54764c62865ff0cd3f271fb1d4635662bff10f0878694f1654fb7fbdecb830d.
//
// Solidity: event ValidatorSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) FilterValidatorSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsValidatorSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "ValidatorSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsValidatorSignerRemovedIterator{contract: _Accounts.contract, event: "ValidatorSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorSignerRemoved is a free log subscription operation binding the contract event 0xa54764c62865ff0cd3f271fb1d4635662bff10f0878694f1654fb7fbdecb830d.
//
// Solidity: event ValidatorSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) WatchValidatorSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsValidatorSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "ValidatorSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsValidatorSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerRemoved", log); err != nil {
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

// ParseValidatorSignerRemoved is a log parse operation binding the contract event 0xa54764c62865ff0cd3f271fb1d4635662bff10f0878694f1654fb7fbdecb830d.
//
// Solidity: event ValidatorSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) ParseValidatorSignerRemoved(log types.Log) (*AccountsValidatorSignerRemoved, error) {
	event := new(AccountsValidatorSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "ValidatorSignerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsVoteSignerAuthorizedIterator is returned from FilterVoteSignerAuthorized and is used to iterate over the raw logs and unpacked data for VoteSignerAuthorized events raised by the Accounts contract.
type AccountsVoteSignerAuthorizedIterator struct {
	Event *AccountsVoteSignerAuthorized // Event containing the contract specifics and raw log

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
func (it *AccountsVoteSignerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsVoteSignerAuthorized)
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
		it.Event = new(AccountsVoteSignerAuthorized)
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
func (it *AccountsVoteSignerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsVoteSignerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsVoteSignerAuthorized represents a VoteSignerAuthorized event raised by the Accounts contract.
type AccountsVoteSignerAuthorized struct {
	Account common.Address
	Signer  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVoteSignerAuthorized is a free log retrieval operation binding the contract event 0xaab5f8a189373aaa290f42ae65ea5d7971b732366ca5bf66556e76263944af28.
//
// Solidity: event VoteSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) FilterVoteSignerAuthorized(opts *bind.FilterOpts, account []common.Address) (*AccountsVoteSignerAuthorizedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "VoteSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsVoteSignerAuthorizedIterator{contract: _Accounts.contract, event: "VoteSignerAuthorized", logs: logs, sub: sub}, nil
}

// WatchVoteSignerAuthorized is a free log subscription operation binding the contract event 0xaab5f8a189373aaa290f42ae65ea5d7971b732366ca5bf66556e76263944af28.
//
// Solidity: event VoteSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) WatchVoteSignerAuthorized(opts *bind.WatchOpts, sink chan<- *AccountsVoteSignerAuthorized, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "VoteSignerAuthorized", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsVoteSignerAuthorized)
				if err := _Accounts.contract.UnpackLog(event, "VoteSignerAuthorized", log); err != nil {
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

// ParseVoteSignerAuthorized is a log parse operation binding the contract event 0xaab5f8a189373aaa290f42ae65ea5d7971b732366ca5bf66556e76263944af28.
//
// Solidity: event VoteSignerAuthorized(address indexed account, address signer)
func (_Accounts *AccountsFilterer) ParseVoteSignerAuthorized(log types.Log) (*AccountsVoteSignerAuthorized, error) {
	event := new(AccountsVoteSignerAuthorized)
	if err := _Accounts.contract.UnpackLog(event, "VoteSignerAuthorized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AccountsVoteSignerRemovedIterator is returned from FilterVoteSignerRemoved and is used to iterate over the raw logs and unpacked data for VoteSignerRemoved events raised by the Accounts contract.
type AccountsVoteSignerRemovedIterator struct {
	Event *AccountsVoteSignerRemoved // Event containing the contract specifics and raw log

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
func (it *AccountsVoteSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AccountsVoteSignerRemoved)
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
		it.Event = new(AccountsVoteSignerRemoved)
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
func (it *AccountsVoteSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AccountsVoteSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AccountsVoteSignerRemoved represents a VoteSignerRemoved event raised by the Accounts contract.
type AccountsVoteSignerRemoved struct {
	Account   common.Address
	OldSigner common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVoteSignerRemoved is a free log retrieval operation binding the contract event 0xa197481f404d8a8082368ad7445380f01e75f27dea6b7aef234a4ce071127fae.
//
// Solidity: event VoteSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) FilterVoteSignerRemoved(opts *bind.FilterOpts, account []common.Address) (*AccountsVoteSignerRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.FilterLogs(opts, "VoteSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &AccountsVoteSignerRemovedIterator{contract: _Accounts.contract, event: "VoteSignerRemoved", logs: logs, sub: sub}, nil
}

// WatchVoteSignerRemoved is a free log subscription operation binding the contract event 0xa197481f404d8a8082368ad7445380f01e75f27dea6b7aef234a4ce071127fae.
//
// Solidity: event VoteSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) WatchVoteSignerRemoved(opts *bind.WatchOpts, sink chan<- *AccountsVoteSignerRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Accounts.contract.WatchLogs(opts, "VoteSignerRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AccountsVoteSignerRemoved)
				if err := _Accounts.contract.UnpackLog(event, "VoteSignerRemoved", log); err != nil {
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

// ParseVoteSignerRemoved is a log parse operation binding the contract event 0xa197481f404d8a8082368ad7445380f01e75f27dea6b7aef234a4ce071127fae.
//
// Solidity: event VoteSignerRemoved(address indexed account, address oldSigner)
func (_Accounts *AccountsFilterer) ParseVoteSignerRemoved(log types.Log) (*AccountsVoteSignerRemoved, error) {
	event := new(AccountsVoteSignerRemoved)
	if err := _Accounts.contract.UnpackLog(event, "VoteSignerRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}
