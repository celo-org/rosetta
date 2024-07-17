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

// DoubleSigningSlasherMetaData contains all meta data concerning the DoubleSigningSlasher contract.
var DoubleSigningSlasherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"DoubleSigningSlashPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"SlashingIncentivesSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"groupMembershipHistoryIndex\",\"type\":\"uint256\"}],\"name\":\"groupMembershipAtBlock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setSlashingIncentives\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashingIncentives\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"headerA\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerB\",\"type\":\"bytes\"}],\"name\":\"checkForDoubleSigning\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"headerA\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"headerB\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"groupMembershipHistoryIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validatorElectionLessers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"validatorElectionGreaters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"validatorElectionIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"groupElectionLessers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"groupElectionGreaters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"groupElectionIndices\",\"type\":\"uint256[]\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DoubleSigningSlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use DoubleSigningSlasherMetaData.ABI instead.
var DoubleSigningSlasherABI = DoubleSigningSlasherMetaData.ABI

// DoubleSigningSlasher is an auto generated Go binding around an Ethereum contract.
type DoubleSigningSlasher struct {
	DoubleSigningSlasherCaller     // Read-only binding to the contract
	DoubleSigningSlasherTransactor // Write-only binding to the contract
	DoubleSigningSlasherFilterer   // Log filterer for contract events
}

// DoubleSigningSlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type DoubleSigningSlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleSigningSlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DoubleSigningSlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleSigningSlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DoubleSigningSlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DoubleSigningSlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DoubleSigningSlasherSession struct {
	Contract     *DoubleSigningSlasher // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DoubleSigningSlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DoubleSigningSlasherCallerSession struct {
	Contract *DoubleSigningSlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// DoubleSigningSlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DoubleSigningSlasherTransactorSession struct {
	Contract     *DoubleSigningSlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// DoubleSigningSlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type DoubleSigningSlasherRaw struct {
	Contract *DoubleSigningSlasher // Generic contract binding to access the raw methods on
}

// DoubleSigningSlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DoubleSigningSlasherCallerRaw struct {
	Contract *DoubleSigningSlasherCaller // Generic read-only contract binding to access the raw methods on
}

// DoubleSigningSlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DoubleSigningSlasherTransactorRaw struct {
	Contract *DoubleSigningSlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDoubleSigningSlasher creates a new instance of DoubleSigningSlasher, bound to a specific deployed contract.
func NewDoubleSigningSlasher(address common.Address, backend bind.ContractBackend) (*DoubleSigningSlasher, error) {
	contract, err := bindDoubleSigningSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasher{DoubleSigningSlasherCaller: DoubleSigningSlasherCaller{contract: contract}, DoubleSigningSlasherTransactor: DoubleSigningSlasherTransactor{contract: contract}, DoubleSigningSlasherFilterer: DoubleSigningSlasherFilterer{contract: contract}}, nil
}

// NewDoubleSigningSlasherCaller creates a new read-only instance of DoubleSigningSlasher, bound to a specific deployed contract.
func NewDoubleSigningSlasherCaller(address common.Address, caller bind.ContractCaller) (*DoubleSigningSlasherCaller, error) {
	contract, err := bindDoubleSigningSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherCaller{contract: contract}, nil
}

// NewDoubleSigningSlasherTransactor creates a new write-only instance of DoubleSigningSlasher, bound to a specific deployed contract.
func NewDoubleSigningSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*DoubleSigningSlasherTransactor, error) {
	contract, err := bindDoubleSigningSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherTransactor{contract: contract}, nil
}

// NewDoubleSigningSlasherFilterer creates a new log filterer instance of DoubleSigningSlasher, bound to a specific deployed contract.
func NewDoubleSigningSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*DoubleSigningSlasherFilterer, error) {
	contract, err := bindDoubleSigningSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherFilterer{contract: contract}, nil
}

// bindDoubleSigningSlasher binds a generic wrapper to an already deployed contract.
func bindDoubleSigningSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DoubleSigningSlasherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseDoubleSigningSlasherABI parses the ABI
func ParseDoubleSigningSlasherABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(DoubleSigningSlasherABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleSigningSlasher *DoubleSigningSlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleSigningSlasher.Contract.DoubleSigningSlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleSigningSlasher *DoubleSigningSlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.DoubleSigningSlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleSigningSlasher *DoubleSigningSlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.DoubleSigningSlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DoubleSigningSlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.contract.Transact(opts, method, params...)
}

// CheckForDoubleSigning is a free data retrieval call binding the contract method 0x09f99447.
//
// Solidity: function checkForDoubleSigning(address signer, uint256 index, bytes headerA, bytes headerB) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) CheckForDoubleSigning(opts *bind.CallOpts, signer common.Address, index *big.Int, headerA []byte, headerB []byte) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "checkForDoubleSigning", signer, index, headerA, headerB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckForDoubleSigning is a free data retrieval call binding the contract method 0x09f99447.
//
// Solidity: function checkForDoubleSigning(address signer, uint256 index, bytes headerA, bytes headerB) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) CheckForDoubleSigning(signer common.Address, index *big.Int, headerA []byte, headerB []byte) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.CheckForDoubleSigning(&_DoubleSigningSlasher.CallOpts, signer, index, headerA, headerB)
}

// CheckForDoubleSigning is a free data retrieval call binding the contract method 0x09f99447.
//
// Solidity: function checkForDoubleSigning(address signer, uint256 index, bytes headerA, bytes headerB) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) CheckForDoubleSigning(signer common.Address, index *big.Int, headerA []byte, headerB []byte) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.CheckForDoubleSigning(&_DoubleSigningSlasher.CallOpts, signer, index, headerA, headerB)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _DoubleSigningSlasher.Contract.CheckProofOfPossession(&_DoubleSigningSlasher.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _DoubleSigningSlasher.Contract.CheckProofOfPossession(&_DoubleSigningSlasher.CallOpts, sender, blsKey, blsPop)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _DoubleSigningSlasher.Contract.FractionMulExp(&_DoubleSigningSlasher.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _DoubleSigningSlasher.Contract.FractionMulExp(&_DoubleSigningSlasher.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetBlockNumberFromHeader(&_DoubleSigningSlasher.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetBlockNumberFromHeader(&_DoubleSigningSlasher.CallOpts, header)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetEpochNumber() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetEpochNumber(&_DoubleSigningSlasher.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetEpochNumber() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetEpochNumber(&_DoubleSigningSlasher.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetEpochNumberOfBlock(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetEpochNumberOfBlock(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetEpochSize() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetEpochSize(&_DoubleSigningSlasher.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetEpochSize() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetEpochSize(&_DoubleSigningSlasher.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _DoubleSigningSlasher.Contract.GetParentSealBitmap(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _DoubleSigningSlasher.Contract.GetParentSealBitmap(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _DoubleSigningSlasher.Contract.GetVerifiedSealBitmapFromHeader(&_DoubleSigningSlasher.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _DoubleSigningSlasher.Contract.GetVerifiedSealBitmapFromHeader(&_DoubleSigningSlasher.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "getVersionNumber")

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
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetVersionNumber(&_DoubleSigningSlasher.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _DoubleSigningSlasher.Contract.GetVersionNumber(&_DoubleSigningSlasher.CallOpts)
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) GroupMembershipAtBlock(opts *bind.CallOpts, validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "groupMembershipAtBlock", validator, blockNumber, groupMembershipHistoryIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) GroupMembershipAtBlock(validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	return _DoubleSigningSlasher.Contract.GroupMembershipAtBlock(&_DoubleSigningSlasher.CallOpts, validator, blockNumber, groupMembershipHistoryIndex)
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) GroupMembershipAtBlock(validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	return _DoubleSigningSlasher.Contract.GroupMembershipAtBlock(&_DoubleSigningSlasher.CallOpts, validator, blockNumber, groupMembershipHistoryIndex)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) HashHeader(header []byte) ([32]byte, error) {
	return _DoubleSigningSlasher.Contract.HashHeader(&_DoubleSigningSlasher.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _DoubleSigningSlasher.Contract.HashHeader(&_DoubleSigningSlasher.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) Initialized() (bool, error) {
	return _DoubleSigningSlasher.Contract.Initialized(&_DoubleSigningSlasher.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) Initialized() (bool, error) {
	return _DoubleSigningSlasher.Contract.Initialized(&_DoubleSigningSlasher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) IsOwner() (bool, error) {
	return _DoubleSigningSlasher.Contract.IsOwner(&_DoubleSigningSlasher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) IsOwner() (bool, error) {
	return _DoubleSigningSlasher.Contract.IsOwner(&_DoubleSigningSlasher.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.MinQuorumSize(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.MinQuorumSize(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.MinQuorumSizeInCurrentSet(&_DoubleSigningSlasher.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.MinQuorumSizeInCurrentSet(&_DoubleSigningSlasher.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.NumberValidatorsInCurrentSet(&_DoubleSigningSlasher.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.NumberValidatorsInCurrentSet(&_DoubleSigningSlasher.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.NumberValidatorsInSet(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _DoubleSigningSlasher.Contract.NumberValidatorsInSet(&_DoubleSigningSlasher.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) Owner() (common.Address, error) {
	return _DoubleSigningSlasher.Contract.Owner(&_DoubleSigningSlasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) Owner() (common.Address, error) {
	return _DoubleSigningSlasher.Contract.Owner(&_DoubleSigningSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) Registry() (common.Address, error) {
	return _DoubleSigningSlasher.Contract.Registry(&_DoubleSigningSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) Registry() (common.Address, error) {
	return _DoubleSigningSlasher.Contract.Registry(&_DoubleSigningSlasher.CallOpts)
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() view returns(uint256 penalty, uint256 reward)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) SlashingIncentives(opts *bind.CallOpts) (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "slashingIncentives")

	outstruct := new(struct {
		Penalty *big.Int
		Reward  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Penalty = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reward = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() view returns(uint256 penalty, uint256 reward)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) SlashingIncentives() (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	return _DoubleSigningSlasher.Contract.SlashingIncentives(&_DoubleSigningSlasher.CallOpts)
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() view returns(uint256 penalty, uint256 reward)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) SlashingIncentives() (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	return _DoubleSigningSlasher.Contract.SlashingIncentives(&_DoubleSigningSlasher.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _DoubleSigningSlasher.Contract.ValidatorSignerAddressFromCurrentSet(&_DoubleSigningSlasher.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _DoubleSigningSlasher.Contract.ValidatorSignerAddressFromCurrentSet(&_DoubleSigningSlasher.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DoubleSigningSlasher.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _DoubleSigningSlasher.Contract.ValidatorSignerAddressFromSet(&_DoubleSigningSlasher.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_DoubleSigningSlasher *DoubleSigningSlasherCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _DoubleSigningSlasher.Contract.ValidatorSignerAddressFromSet(&_DoubleSigningSlasher.CallOpts, index, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _penalty *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.contract.Transact(opts, "initialize", registryAddress, _penalty, _reward)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) Initialize(registryAddress common.Address, _penalty *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.Initialize(&_DoubleSigningSlasher.TransactOpts, registryAddress, _penalty, _reward)
}

// Initialize is a paid mutator transaction binding the contract method 0x7a1ac61e.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorSession) Initialize(registryAddress common.Address, _penalty *big.Int, _reward *big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.Initialize(&_DoubleSigningSlasher.TransactOpts, registryAddress, _penalty, _reward)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DoubleSigningSlasher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) RenounceOwnership() (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.RenounceOwnership(&_DoubleSigningSlasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.RenounceOwnership(&_DoubleSigningSlasher.TransactOpts)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _DoubleSigningSlasher.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.SetRegistry(&_DoubleSigningSlasher.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.SetRegistry(&_DoubleSigningSlasher.TransactOpts, registryAddress)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactor) SetSlashingIncentives(opts *bind.TransactOpts, penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.contract.Transact(opts, "setSlashingIncentives", penalty, reward)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) SetSlashingIncentives(penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.SetSlashingIncentives(&_DoubleSigningSlasher.TransactOpts, penalty, reward)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorSession) SetSlashingIncentives(penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.SetSlashingIncentives(&_DoubleSigningSlasher.TransactOpts, penalty, reward)
}

// Slash is a paid mutator transaction binding the contract method 0x8cc26910.
//
// Solidity: function slash(address signer, uint256 index, bytes headerA, bytes headerB, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactor) Slash(opts *bind.TransactOpts, signer common.Address, index *big.Int, headerA []byte, headerB []byte, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.contract.Transact(opts, "slash", signer, index, headerA, headerB, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0x8cc26910.
//
// Solidity: function slash(address signer, uint256 index, bytes headerA, bytes headerB, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) Slash(signer common.Address, index *big.Int, headerA []byte, headerB []byte, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.Slash(&_DoubleSigningSlasher.TransactOpts, signer, index, headerA, headerB, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0x8cc26910.
//
// Solidity: function slash(address signer, uint256 index, bytes headerA, bytes headerB, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorSession) Slash(signer common.Address, index *big.Int, headerA []byte, headerB []byte, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.Slash(&_DoubleSigningSlasher.TransactOpts, signer, index, headerA, headerB, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DoubleSigningSlasher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.TransferOwnership(&_DoubleSigningSlasher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DoubleSigningSlasher *DoubleSigningSlasherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DoubleSigningSlasher.Contract.TransferOwnership(&_DoubleSigningSlasher.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _DoubleSigningSlasher.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "DoubleSigningSlashPerformed":
		event, err = _DoubleSigningSlasher.ParseDoubleSigningSlashPerformed(log)
	case "OwnershipTransferred":
		event, err = _DoubleSigningSlasher.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _DoubleSigningSlasher.ParseRegistrySet(log)
	case "SlashingIncentivesSet":
		event, err = _DoubleSigningSlasher.ParseSlashingIncentivesSet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// DoubleSigningSlasherDoubleSigningSlashPerformedIterator is returned from FilterDoubleSigningSlashPerformed and is used to iterate over the raw logs and unpacked data for DoubleSigningSlashPerformed events raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherDoubleSigningSlashPerformedIterator struct {
	Event *DoubleSigningSlasherDoubleSigningSlashPerformed // Event containing the contract specifics and raw log

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
func (it *DoubleSigningSlasherDoubleSigningSlashPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleSigningSlasherDoubleSigningSlashPerformed)
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
		it.Event = new(DoubleSigningSlasherDoubleSigningSlashPerformed)
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
func (it *DoubleSigningSlasherDoubleSigningSlashPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleSigningSlasherDoubleSigningSlashPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleSigningSlasherDoubleSigningSlashPerformed represents a DoubleSigningSlashPerformed event raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherDoubleSigningSlashPerformed struct {
	Validator   common.Address
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDoubleSigningSlashPerformed is a free log retrieval operation binding the contract event 0xca7992de940988854714f90c0236621d5b6b850313f03eeea47f7028aaecea40.
//
// Solidity: event DoubleSigningSlashPerformed(address indexed validator, uint256 indexed blockNumber)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) FilterDoubleSigningSlashPerformed(opts *bind.FilterOpts, validator []common.Address, blockNumber []*big.Int) (*DoubleSigningSlasherDoubleSigningSlashPerformedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _DoubleSigningSlasher.contract.FilterLogs(opts, "DoubleSigningSlashPerformed", validatorRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherDoubleSigningSlashPerformedIterator{contract: _DoubleSigningSlasher.contract, event: "DoubleSigningSlashPerformed", logs: logs, sub: sub}, nil
}

// WatchDoubleSigningSlashPerformed is a free log subscription operation binding the contract event 0xca7992de940988854714f90c0236621d5b6b850313f03eeea47f7028aaecea40.
//
// Solidity: event DoubleSigningSlashPerformed(address indexed validator, uint256 indexed blockNumber)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) WatchDoubleSigningSlashPerformed(opts *bind.WatchOpts, sink chan<- *DoubleSigningSlasherDoubleSigningSlashPerformed, validator []common.Address, blockNumber []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var blockNumberRule []interface{}
	for _, blockNumberItem := range blockNumber {
		blockNumberRule = append(blockNumberRule, blockNumberItem)
	}

	logs, sub, err := _DoubleSigningSlasher.contract.WatchLogs(opts, "DoubleSigningSlashPerformed", validatorRule, blockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleSigningSlasherDoubleSigningSlashPerformed)
				if err := _DoubleSigningSlasher.contract.UnpackLog(event, "DoubleSigningSlashPerformed", log); err != nil {
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

// ParseDoubleSigningSlashPerformed is a log parse operation binding the contract event 0xca7992de940988854714f90c0236621d5b6b850313f03eeea47f7028aaecea40.
//
// Solidity: event DoubleSigningSlashPerformed(address indexed validator, uint256 indexed blockNumber)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) ParseDoubleSigningSlashPerformed(log types.Log) (*DoubleSigningSlasherDoubleSigningSlashPerformed, error) {
	event := new(DoubleSigningSlasherDoubleSigningSlashPerformed)
	if err := _DoubleSigningSlasher.contract.UnpackLog(event, "DoubleSigningSlashPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleSigningSlasherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherOwnershipTransferredIterator struct {
	Event *DoubleSigningSlasherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DoubleSigningSlasherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleSigningSlasherOwnershipTransferred)
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
		it.Event = new(DoubleSigningSlasherOwnershipTransferred)
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
func (it *DoubleSigningSlasherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleSigningSlasherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleSigningSlasherOwnershipTransferred represents a OwnershipTransferred event raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DoubleSigningSlasherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DoubleSigningSlasher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherOwnershipTransferredIterator{contract: _DoubleSigningSlasher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DoubleSigningSlasherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DoubleSigningSlasher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleSigningSlasherOwnershipTransferred)
				if err := _DoubleSigningSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) ParseOwnershipTransferred(log types.Log) (*DoubleSigningSlasherOwnershipTransferred, error) {
	event := new(DoubleSigningSlasherOwnershipTransferred)
	if err := _DoubleSigningSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleSigningSlasherRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherRegistrySetIterator struct {
	Event *DoubleSigningSlasherRegistrySet // Event containing the contract specifics and raw log

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
func (it *DoubleSigningSlasherRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleSigningSlasherRegistrySet)
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
		it.Event = new(DoubleSigningSlasherRegistrySet)
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
func (it *DoubleSigningSlasherRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleSigningSlasherRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleSigningSlasherRegistrySet represents a RegistrySet event raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*DoubleSigningSlasherRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _DoubleSigningSlasher.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherRegistrySetIterator{contract: _DoubleSigningSlasher.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *DoubleSigningSlasherRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _DoubleSigningSlasher.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleSigningSlasherRegistrySet)
				if err := _DoubleSigningSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) ParseRegistrySet(log types.Log) (*DoubleSigningSlasherRegistrySet, error) {
	event := new(DoubleSigningSlasherRegistrySet)
	if err := _DoubleSigningSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DoubleSigningSlasherSlashingIncentivesSetIterator is returned from FilterSlashingIncentivesSet and is used to iterate over the raw logs and unpacked data for SlashingIncentivesSet events raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherSlashingIncentivesSetIterator struct {
	Event *DoubleSigningSlasherSlashingIncentivesSet // Event containing the contract specifics and raw log

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
func (it *DoubleSigningSlasherSlashingIncentivesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DoubleSigningSlasherSlashingIncentivesSet)
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
		it.Event = new(DoubleSigningSlasherSlashingIncentivesSet)
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
func (it *DoubleSigningSlasherSlashingIncentivesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DoubleSigningSlasherSlashingIncentivesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DoubleSigningSlasherSlashingIncentivesSet represents a SlashingIncentivesSet event raised by the DoubleSigningSlasher contract.
type DoubleSigningSlasherSlashingIncentivesSet struct {
	Penalty *big.Int
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashingIncentivesSet is a free log retrieval operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) FilterSlashingIncentivesSet(opts *bind.FilterOpts) (*DoubleSigningSlasherSlashingIncentivesSetIterator, error) {

	logs, sub, err := _DoubleSigningSlasher.contract.FilterLogs(opts, "SlashingIncentivesSet")
	if err != nil {
		return nil, err
	}
	return &DoubleSigningSlasherSlashingIncentivesSetIterator{contract: _DoubleSigningSlasher.contract, event: "SlashingIncentivesSet", logs: logs, sub: sub}, nil
}

// WatchSlashingIncentivesSet is a free log subscription operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) WatchSlashingIncentivesSet(opts *bind.WatchOpts, sink chan<- *DoubleSigningSlasherSlashingIncentivesSet) (event.Subscription, error) {

	logs, sub, err := _DoubleSigningSlasher.contract.WatchLogs(opts, "SlashingIncentivesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DoubleSigningSlasherSlashingIncentivesSet)
				if err := _DoubleSigningSlasher.contract.UnpackLog(event, "SlashingIncentivesSet", log); err != nil {
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

// ParseSlashingIncentivesSet is a log parse operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DoubleSigningSlasher *DoubleSigningSlasherFilterer) ParseSlashingIncentivesSet(log types.Log) (*DoubleSigningSlasherSlashingIncentivesSet, error) {
	event := new(DoubleSigningSlasherSlashingIncentivesSet)
	if err := _DoubleSigningSlasher.contract.UnpackLog(event, "SlashingIncentivesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
