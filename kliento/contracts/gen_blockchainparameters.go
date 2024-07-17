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

// BlockchainParametersMetaData contains all meta data concerning the BlockchainParameters contract.
var BlockchainParametersMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"BlockGasLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"IntrinsicGasForAlternativeFeeCurrencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"window\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"activationEpoch\",\"type\":\"uint256\"}],\"name\":\"UptimeLookbackWindowSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"intrinsicGasForAlternativeFeeCurrency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uptimeLookbackWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"oldValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nextValueActivationEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_gasForNonGoldCurrencies\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lookbackWindow\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasLimit\",\"type\":\"uint256\"}],\"name\":\"setBlockGasLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gas\",\"type\":\"uint256\"}],\"name\":\"setIntrinsicGasForAlternativeFeeCurrency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"window\",\"type\":\"uint256\"}],\"name\":\"setUptimeLookbackWindow\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUptimeLookbackWindow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lookbackWindow\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BlockchainParametersABI is the input ABI used to generate the binding from.
// Deprecated: Use BlockchainParametersMetaData.ABI instead.
var BlockchainParametersABI = BlockchainParametersMetaData.ABI

// BlockchainParameters is an auto generated Go binding around an Ethereum contract.
type BlockchainParameters struct {
	BlockchainParametersCaller     // Read-only binding to the contract
	BlockchainParametersTransactor // Write-only binding to the contract
	BlockchainParametersFilterer   // Log filterer for contract events
}

// BlockchainParametersCaller is an auto generated read-only Go binding around an Ethereum contract.
type BlockchainParametersCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainParametersTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BlockchainParametersTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainParametersFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BlockchainParametersFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BlockchainParametersSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BlockchainParametersSession struct {
	Contract     *BlockchainParameters // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BlockchainParametersCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BlockchainParametersCallerSession struct {
	Contract *BlockchainParametersCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BlockchainParametersTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BlockchainParametersTransactorSession struct {
	Contract     *BlockchainParametersTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BlockchainParametersRaw is an auto generated low-level Go binding around an Ethereum contract.
type BlockchainParametersRaw struct {
	Contract *BlockchainParameters // Generic contract binding to access the raw methods on
}

// BlockchainParametersCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BlockchainParametersCallerRaw struct {
	Contract *BlockchainParametersCaller // Generic read-only contract binding to access the raw methods on
}

// BlockchainParametersTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BlockchainParametersTransactorRaw struct {
	Contract *BlockchainParametersTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBlockchainParameters creates a new instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParameters(address common.Address, backend bind.ContractBackend) (*BlockchainParameters, error) {
	contract, err := bindBlockchainParameters(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BlockchainParameters{BlockchainParametersCaller: BlockchainParametersCaller{contract: contract}, BlockchainParametersTransactor: BlockchainParametersTransactor{contract: contract}, BlockchainParametersFilterer: BlockchainParametersFilterer{contract: contract}}, nil
}

// NewBlockchainParametersCaller creates a new read-only instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParametersCaller(address common.Address, caller bind.ContractCaller) (*BlockchainParametersCaller, error) {
	contract, err := bindBlockchainParameters(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersCaller{contract: contract}, nil
}

// NewBlockchainParametersTransactor creates a new write-only instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParametersTransactor(address common.Address, transactor bind.ContractTransactor) (*BlockchainParametersTransactor, error) {
	contract, err := bindBlockchainParameters(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersTransactor{contract: contract}, nil
}

// NewBlockchainParametersFilterer creates a new log filterer instance of BlockchainParameters, bound to a specific deployed contract.
func NewBlockchainParametersFilterer(address common.Address, filterer bind.ContractFilterer) (*BlockchainParametersFilterer, error) {
	contract, err := bindBlockchainParameters(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersFilterer{contract: contract}, nil
}

// bindBlockchainParameters binds a generic wrapper to an already deployed contract.
func bindBlockchainParameters(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainParametersABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseBlockchainParametersABI parses the ABI
func ParseBlockchainParametersABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(BlockchainParametersABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockchainParameters *BlockchainParametersRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockchainParameters.Contract.BlockchainParametersCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockchainParameters *BlockchainParametersRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.BlockchainParametersTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockchainParameters *BlockchainParametersRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.BlockchainParametersTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BlockchainParameters *BlockchainParametersCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BlockchainParameters.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BlockchainParameters *BlockchainParametersTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BlockchainParameters *BlockchainParametersTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.contract.Transact(opts, method, params...)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) BlockGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "blockGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) BlockGasLimit() (*big.Int, error) {
	return _BlockchainParameters.Contract.BlockGasLimit(&_BlockchainParameters.CallOpts)
}

// BlockGasLimit is a free data retrieval call binding the contract method 0x7877a797.
//
// Solidity: function blockGasLimit() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) BlockGasLimit() (*big.Int, error) {
	return _BlockchainParameters.Contract.BlockGasLimit(&_BlockchainParameters.CallOpts)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_BlockchainParameters *BlockchainParametersCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_BlockchainParameters *BlockchainParametersSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _BlockchainParameters.Contract.CheckProofOfPossession(&_BlockchainParameters.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_BlockchainParameters *BlockchainParametersCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _BlockchainParameters.Contract.CheckProofOfPossession(&_BlockchainParameters.CallOpts, sender, blsKey, blsPop)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_BlockchainParameters *BlockchainParametersCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

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
func (_BlockchainParameters *BlockchainParametersSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _BlockchainParameters.Contract.FractionMulExp(&_BlockchainParameters.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _BlockchainParameters.Contract.FractionMulExp(&_BlockchainParameters.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _BlockchainParameters.Contract.GetBlockNumberFromHeader(&_BlockchainParameters.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _BlockchainParameters.Contract.GetBlockNumberFromHeader(&_BlockchainParameters.CallOpts, header)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) GetEpochNumber() (*big.Int, error) {
	return _BlockchainParameters.Contract.GetEpochNumber(&_BlockchainParameters.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetEpochNumber() (*big.Int, error) {
	return _BlockchainParameters.Contract.GetEpochNumber(&_BlockchainParameters.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _BlockchainParameters.Contract.GetEpochNumberOfBlock(&_BlockchainParameters.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _BlockchainParameters.Contract.GetEpochNumberOfBlock(&_BlockchainParameters.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) GetEpochSize() (*big.Int, error) {
	return _BlockchainParameters.Contract.GetEpochSize(&_BlockchainParameters.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetEpochSize() (*big.Int, error) {
	return _BlockchainParameters.Contract.GetEpochSize(&_BlockchainParameters.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _BlockchainParameters.Contract.GetParentSealBitmap(&_BlockchainParameters.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _BlockchainParameters.Contract.GetParentSealBitmap(&_BlockchainParameters.CallOpts, blockNumber)
}

// GetUptimeLookbackWindow is a free data retrieval call binding the contract method 0x52bed4d7.
//
// Solidity: function getUptimeLookbackWindow() view returns(uint256 lookbackWindow)
func (_BlockchainParameters *BlockchainParametersCaller) GetUptimeLookbackWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getUptimeLookbackWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUptimeLookbackWindow is a free data retrieval call binding the contract method 0x52bed4d7.
//
// Solidity: function getUptimeLookbackWindow() view returns(uint256 lookbackWindow)
func (_BlockchainParameters *BlockchainParametersSession) GetUptimeLookbackWindow() (*big.Int, error) {
	return _BlockchainParameters.Contract.GetUptimeLookbackWindow(&_BlockchainParameters.CallOpts)
}

// GetUptimeLookbackWindow is a free data retrieval call binding the contract method 0x52bed4d7.
//
// Solidity: function getUptimeLookbackWindow() view returns(uint256 lookbackWindow)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetUptimeLookbackWindow() (*big.Int, error) {
	return _BlockchainParameters.Contract.GetUptimeLookbackWindow(&_BlockchainParameters.CallOpts)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _BlockchainParameters.Contract.GetVerifiedSealBitmapFromHeader(&_BlockchainParameters.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _BlockchainParameters.Contract.GetVerifiedSealBitmapFromHeader(&_BlockchainParameters.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_BlockchainParameters *BlockchainParametersCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "getVersionNumber")

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
func (_BlockchainParameters *BlockchainParametersSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BlockchainParameters.Contract.GetVersionNumber(&_BlockchainParameters.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _BlockchainParameters.Contract.GetVersionNumber(&_BlockchainParameters.CallOpts)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersSession) HashHeader(header []byte) ([32]byte, error) {
	return _BlockchainParameters.Contract.HashHeader(&_BlockchainParameters.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_BlockchainParameters *BlockchainParametersCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _BlockchainParameters.Contract.HashHeader(&_BlockchainParameters.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_BlockchainParameters *BlockchainParametersCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_BlockchainParameters *BlockchainParametersSession) Initialized() (bool, error) {
	return _BlockchainParameters.Contract.Initialized(&_BlockchainParameters.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_BlockchainParameters *BlockchainParametersCallerSession) Initialized() (bool, error) {
	return _BlockchainParameters.Contract.Initialized(&_BlockchainParameters.CallOpts)
}

// IntrinsicGasForAlternativeFeeCurrency is a free data retrieval call binding the contract method 0x808474f1.
//
// Solidity: function intrinsicGasForAlternativeFeeCurrency() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) IntrinsicGasForAlternativeFeeCurrency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "intrinsicGasForAlternativeFeeCurrency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IntrinsicGasForAlternativeFeeCurrency is a free data retrieval call binding the contract method 0x808474f1.
//
// Solidity: function intrinsicGasForAlternativeFeeCurrency() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) IntrinsicGasForAlternativeFeeCurrency() (*big.Int, error) {
	return _BlockchainParameters.Contract.IntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.CallOpts)
}

// IntrinsicGasForAlternativeFeeCurrency is a free data retrieval call binding the contract method 0x808474f1.
//
// Solidity: function intrinsicGasForAlternativeFeeCurrency() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) IntrinsicGasForAlternativeFeeCurrency() (*big.Int, error) {
	return _BlockchainParameters.Contract.IntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BlockchainParameters *BlockchainParametersCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BlockchainParameters *BlockchainParametersSession) IsOwner() (bool, error) {
	return _BlockchainParameters.Contract.IsOwner(&_BlockchainParameters.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_BlockchainParameters *BlockchainParametersCallerSession) IsOwner() (bool, error) {
	return _BlockchainParameters.Contract.IsOwner(&_BlockchainParameters.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _BlockchainParameters.Contract.MinQuorumSize(&_BlockchainParameters.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _BlockchainParameters.Contract.MinQuorumSize(&_BlockchainParameters.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _BlockchainParameters.Contract.MinQuorumSizeInCurrentSet(&_BlockchainParameters.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _BlockchainParameters.Contract.MinQuorumSizeInCurrentSet(&_BlockchainParameters.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _BlockchainParameters.Contract.NumberValidatorsInCurrentSet(&_BlockchainParameters.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _BlockchainParameters.Contract.NumberValidatorsInCurrentSet(&_BlockchainParameters.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _BlockchainParameters.Contract.NumberValidatorsInSet(&_BlockchainParameters.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_BlockchainParameters *BlockchainParametersCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _BlockchainParameters.Contract.NumberValidatorsInSet(&_BlockchainParameters.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockchainParameters *BlockchainParametersCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockchainParameters *BlockchainParametersSession) Owner() (common.Address, error) {
	return _BlockchainParameters.Contract.Owner(&_BlockchainParameters.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BlockchainParameters *BlockchainParametersCallerSession) Owner() (common.Address, error) {
	return _BlockchainParameters.Contract.Owner(&_BlockchainParameters.CallOpts)
}

// UptimeLookbackWindow is a free data retrieval call binding the contract method 0x61568828.
//
// Solidity: function uptimeLookbackWindow() view returns(uint256 oldValue, uint256 nextValue, uint256 nextValueActivationEpoch)
func (_BlockchainParameters *BlockchainParametersCaller) UptimeLookbackWindow(opts *bind.CallOpts) (struct {
	OldValue                 *big.Int
	NextValue                *big.Int
	NextValueActivationEpoch *big.Int
}, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "uptimeLookbackWindow")

	outstruct := new(struct {
		OldValue                 *big.Int
		NextValue                *big.Int
		NextValueActivationEpoch *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OldValue = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NextValue = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.NextValueActivationEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UptimeLookbackWindow is a free data retrieval call binding the contract method 0x61568828.
//
// Solidity: function uptimeLookbackWindow() view returns(uint256 oldValue, uint256 nextValue, uint256 nextValueActivationEpoch)
func (_BlockchainParameters *BlockchainParametersSession) UptimeLookbackWindow() (struct {
	OldValue                 *big.Int
	NextValue                *big.Int
	NextValueActivationEpoch *big.Int
}, error) {
	return _BlockchainParameters.Contract.UptimeLookbackWindow(&_BlockchainParameters.CallOpts)
}

// UptimeLookbackWindow is a free data retrieval call binding the contract method 0x61568828.
//
// Solidity: function uptimeLookbackWindow() view returns(uint256 oldValue, uint256 nextValue, uint256 nextValueActivationEpoch)
func (_BlockchainParameters *BlockchainParametersCallerSession) UptimeLookbackWindow() (struct {
	OldValue                 *big.Int
	NextValue                *big.Int
	NextValueActivationEpoch *big.Int
}, error) {
	return _BlockchainParameters.Contract.UptimeLookbackWindow(&_BlockchainParameters.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_BlockchainParameters *BlockchainParametersCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_BlockchainParameters *BlockchainParametersSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _BlockchainParameters.Contract.ValidatorSignerAddressFromCurrentSet(&_BlockchainParameters.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_BlockchainParameters *BlockchainParametersCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _BlockchainParameters.Contract.ValidatorSignerAddressFromCurrentSet(&_BlockchainParameters.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_BlockchainParameters *BlockchainParametersCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BlockchainParameters.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_BlockchainParameters *BlockchainParametersSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _BlockchainParameters.Contract.ValidatorSignerAddressFromSet(&_BlockchainParameters.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_BlockchainParameters *BlockchainParametersCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _BlockchainParameters.Contract.ValidatorSignerAddressFromSet(&_BlockchainParameters.CallOpts, index, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _gasForNonGoldCurrencies, uint256 gasLimit, uint256 lookbackWindow) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) Initialize(opts *bind.TransactOpts, _gasForNonGoldCurrencies *big.Int, gasLimit *big.Int, lookbackWindow *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "initialize", _gasForNonGoldCurrencies, gasLimit, lookbackWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _gasForNonGoldCurrencies, uint256 gasLimit, uint256 lookbackWindow) returns()
func (_BlockchainParameters *BlockchainParametersSession) Initialize(_gasForNonGoldCurrencies *big.Int, gasLimit *big.Int, lookbackWindow *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.Initialize(&_BlockchainParameters.TransactOpts, _gasForNonGoldCurrencies, gasLimit, lookbackWindow)
}

// Initialize is a paid mutator transaction binding the contract method 0x80d85911.
//
// Solidity: function initialize(uint256 _gasForNonGoldCurrencies, uint256 gasLimit, uint256 lookbackWindow) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) Initialize(_gasForNonGoldCurrencies *big.Int, gasLimit *big.Int, lookbackWindow *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.Initialize(&_BlockchainParameters.TransactOpts, _gasForNonGoldCurrencies, gasLimit, lookbackWindow)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockchainParameters *BlockchainParametersTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockchainParameters *BlockchainParametersSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockchainParameters.Contract.RenounceOwnership(&_BlockchainParameters.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BlockchainParameters.Contract.RenounceOwnership(&_BlockchainParameters.TransactOpts)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xa69257f3.
//
// Solidity: function setBlockGasLimit(uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) SetBlockGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "setBlockGasLimit", gasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xa69257f3.
//
// Solidity: function setBlockGasLimit(uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersSession) SetBlockGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetBlockGasLimit(&_BlockchainParameters.TransactOpts, gasLimit)
}

// SetBlockGasLimit is a paid mutator transaction binding the contract method 0xa69257f3.
//
// Solidity: function setBlockGasLimit(uint256 gasLimit) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) SetBlockGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetBlockGasLimit(&_BlockchainParameters.TransactOpts, gasLimit)
}

// SetIntrinsicGasForAlternativeFeeCurrency is a paid mutator transaction binding the contract method 0xcb0ec628.
//
// Solidity: function setIntrinsicGasForAlternativeFeeCurrency(uint256 gas) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) SetIntrinsicGasForAlternativeFeeCurrency(opts *bind.TransactOpts, gas *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "setIntrinsicGasForAlternativeFeeCurrency", gas)
}

// SetIntrinsicGasForAlternativeFeeCurrency is a paid mutator transaction binding the contract method 0xcb0ec628.
//
// Solidity: function setIntrinsicGasForAlternativeFeeCurrency(uint256 gas) returns()
func (_BlockchainParameters *BlockchainParametersSession) SetIntrinsicGasForAlternativeFeeCurrency(gas *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetIntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.TransactOpts, gas)
}

// SetIntrinsicGasForAlternativeFeeCurrency is a paid mutator transaction binding the contract method 0xcb0ec628.
//
// Solidity: function setIntrinsicGasForAlternativeFeeCurrency(uint256 gas) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) SetIntrinsicGasForAlternativeFeeCurrency(gas *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetIntrinsicGasForAlternativeFeeCurrency(&_BlockchainParameters.TransactOpts, gas)
}

// SetUptimeLookbackWindow is a paid mutator transaction binding the contract method 0xe94fd109.
//
// Solidity: function setUptimeLookbackWindow(uint256 window) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) SetUptimeLookbackWindow(opts *bind.TransactOpts, window *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "setUptimeLookbackWindow", window)
}

// SetUptimeLookbackWindow is a paid mutator transaction binding the contract method 0xe94fd109.
//
// Solidity: function setUptimeLookbackWindow(uint256 window) returns()
func (_BlockchainParameters *BlockchainParametersSession) SetUptimeLookbackWindow(window *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetUptimeLookbackWindow(&_BlockchainParameters.TransactOpts, window)
}

// SetUptimeLookbackWindow is a paid mutator transaction binding the contract method 0xe94fd109.
//
// Solidity: function setUptimeLookbackWindow(uint256 window) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) SetUptimeLookbackWindow(window *big.Int) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.SetUptimeLookbackWindow(&_BlockchainParameters.TransactOpts, window)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainParameters *BlockchainParametersTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainParameters.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainParameters *BlockchainParametersSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.TransferOwnership(&_BlockchainParameters.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BlockchainParameters *BlockchainParametersTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BlockchainParameters.Contract.TransferOwnership(&_BlockchainParameters.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_BlockchainParameters *BlockchainParametersFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _BlockchainParameters.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "BlockGasLimitSet":
		event, err = _BlockchainParameters.ParseBlockGasLimitSet(log)
	case "IntrinsicGasForAlternativeFeeCurrencySet":
		event, err = _BlockchainParameters.ParseIntrinsicGasForAlternativeFeeCurrencySet(log)
	case "OwnershipTransferred":
		event, err = _BlockchainParameters.ParseOwnershipTransferred(log)
	case "UptimeLookbackWindowSet":
		event, err = _BlockchainParameters.ParseUptimeLookbackWindowSet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// BlockchainParametersBlockGasLimitSetIterator is returned from FilterBlockGasLimitSet and is used to iterate over the raw logs and unpacked data for BlockGasLimitSet events raised by the BlockchainParameters contract.
type BlockchainParametersBlockGasLimitSetIterator struct {
	Event *BlockchainParametersBlockGasLimitSet // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersBlockGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersBlockGasLimitSet)
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
		it.Event = new(BlockchainParametersBlockGasLimitSet)
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
func (it *BlockchainParametersBlockGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersBlockGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersBlockGasLimitSet represents a BlockGasLimitSet event raised by the BlockchainParameters contract.
type BlockchainParametersBlockGasLimitSet struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterBlockGasLimitSet is a free log retrieval operation binding the contract event 0x55311ae9c14427b0863f38ed97a2a5944c50d824bbf692836246512e6822c3cf.
//
// Solidity: event BlockGasLimitSet(uint256 limit)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterBlockGasLimitSet(opts *bind.FilterOpts) (*BlockchainParametersBlockGasLimitSetIterator, error) {

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "BlockGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersBlockGasLimitSetIterator{contract: _BlockchainParameters.contract, event: "BlockGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchBlockGasLimitSet is a free log subscription operation binding the contract event 0x55311ae9c14427b0863f38ed97a2a5944c50d824bbf692836246512e6822c3cf.
//
// Solidity: event BlockGasLimitSet(uint256 limit)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchBlockGasLimitSet(opts *bind.WatchOpts, sink chan<- *BlockchainParametersBlockGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "BlockGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersBlockGasLimitSet)
				if err := _BlockchainParameters.contract.UnpackLog(event, "BlockGasLimitSet", log); err != nil {
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

// ParseBlockGasLimitSet is a log parse operation binding the contract event 0x55311ae9c14427b0863f38ed97a2a5944c50d824bbf692836246512e6822c3cf.
//
// Solidity: event BlockGasLimitSet(uint256 limit)
func (_BlockchainParameters *BlockchainParametersFilterer) ParseBlockGasLimitSet(log types.Log) (*BlockchainParametersBlockGasLimitSet, error) {
	event := new(BlockchainParametersBlockGasLimitSet)
	if err := _BlockchainParameters.contract.UnpackLog(event, "BlockGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator is returned from FilterIntrinsicGasForAlternativeFeeCurrencySet and is used to iterate over the raw logs and unpacked data for IntrinsicGasForAlternativeFeeCurrencySet events raised by the BlockchainParameters contract.
type BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator struct {
	Event *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
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
		it.Event = new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
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
func (it *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet represents a IntrinsicGasForAlternativeFeeCurrencySet event raised by the BlockchainParameters contract.
type BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet struct {
	Gas *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterIntrinsicGasForAlternativeFeeCurrencySet is a free log retrieval operation binding the contract event 0xba9c6f28c7d9990745a5b5282dbee04706c28cae24a44736c3ba99b57c021f3e.
//
// Solidity: event IntrinsicGasForAlternativeFeeCurrencySet(uint256 gas)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterIntrinsicGasForAlternativeFeeCurrencySet(opts *bind.FilterOpts) (*BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator, error) {

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "IntrinsicGasForAlternativeFeeCurrencySet")
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySetIterator{contract: _BlockchainParameters.contract, event: "IntrinsicGasForAlternativeFeeCurrencySet", logs: logs, sub: sub}, nil
}

// WatchIntrinsicGasForAlternativeFeeCurrencySet is a free log subscription operation binding the contract event 0xba9c6f28c7d9990745a5b5282dbee04706c28cae24a44736c3ba99b57c021f3e.
//
// Solidity: event IntrinsicGasForAlternativeFeeCurrencySet(uint256 gas)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchIntrinsicGasForAlternativeFeeCurrencySet(opts *bind.WatchOpts, sink chan<- *BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet) (event.Subscription, error) {

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "IntrinsicGasForAlternativeFeeCurrencySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
				if err := _BlockchainParameters.contract.UnpackLog(event, "IntrinsicGasForAlternativeFeeCurrencySet", log); err != nil {
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

// ParseIntrinsicGasForAlternativeFeeCurrencySet is a log parse operation binding the contract event 0xba9c6f28c7d9990745a5b5282dbee04706c28cae24a44736c3ba99b57c021f3e.
//
// Solidity: event IntrinsicGasForAlternativeFeeCurrencySet(uint256 gas)
func (_BlockchainParameters *BlockchainParametersFilterer) ParseIntrinsicGasForAlternativeFeeCurrencySet(log types.Log) (*BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet, error) {
	event := new(BlockchainParametersIntrinsicGasForAlternativeFeeCurrencySet)
	if err := _BlockchainParameters.contract.UnpackLog(event, "IntrinsicGasForAlternativeFeeCurrencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainParametersOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BlockchainParameters contract.
type BlockchainParametersOwnershipTransferredIterator struct {
	Event *BlockchainParametersOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersOwnershipTransferred)
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
		it.Event = new(BlockchainParametersOwnershipTransferred)
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
func (it *BlockchainParametersOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersOwnershipTransferred represents a OwnershipTransferred event raised by the BlockchainParameters contract.
type BlockchainParametersOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BlockchainParametersOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersOwnershipTransferredIterator{contract: _BlockchainParameters.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BlockchainParametersOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersOwnershipTransferred)
				if err := _BlockchainParameters.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_BlockchainParameters *BlockchainParametersFilterer) ParseOwnershipTransferred(log types.Log) (*BlockchainParametersOwnershipTransferred, error) {
	event := new(BlockchainParametersOwnershipTransferred)
	if err := _BlockchainParameters.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BlockchainParametersUptimeLookbackWindowSetIterator is returned from FilterUptimeLookbackWindowSet and is used to iterate over the raw logs and unpacked data for UptimeLookbackWindowSet events raised by the BlockchainParameters contract.
type BlockchainParametersUptimeLookbackWindowSetIterator struct {
	Event *BlockchainParametersUptimeLookbackWindowSet // Event containing the contract specifics and raw log

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
func (it *BlockchainParametersUptimeLookbackWindowSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BlockchainParametersUptimeLookbackWindowSet)
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
		it.Event = new(BlockchainParametersUptimeLookbackWindowSet)
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
func (it *BlockchainParametersUptimeLookbackWindowSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BlockchainParametersUptimeLookbackWindowSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BlockchainParametersUptimeLookbackWindowSet represents a UptimeLookbackWindowSet event raised by the BlockchainParameters contract.
type BlockchainParametersUptimeLookbackWindowSet struct {
	Window          *big.Int
	ActivationEpoch *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUptimeLookbackWindowSet is a free log retrieval operation binding the contract event 0x484a24d7faca8c4330aaf9ba5f131e6bd474ed6877a555511f39d16a1d71d15a.
//
// Solidity: event UptimeLookbackWindowSet(uint256 window, uint256 activationEpoch)
func (_BlockchainParameters *BlockchainParametersFilterer) FilterUptimeLookbackWindowSet(opts *bind.FilterOpts) (*BlockchainParametersUptimeLookbackWindowSetIterator, error) {

	logs, sub, err := _BlockchainParameters.contract.FilterLogs(opts, "UptimeLookbackWindowSet")
	if err != nil {
		return nil, err
	}
	return &BlockchainParametersUptimeLookbackWindowSetIterator{contract: _BlockchainParameters.contract, event: "UptimeLookbackWindowSet", logs: logs, sub: sub}, nil
}

// WatchUptimeLookbackWindowSet is a free log subscription operation binding the contract event 0x484a24d7faca8c4330aaf9ba5f131e6bd474ed6877a555511f39d16a1d71d15a.
//
// Solidity: event UptimeLookbackWindowSet(uint256 window, uint256 activationEpoch)
func (_BlockchainParameters *BlockchainParametersFilterer) WatchUptimeLookbackWindowSet(opts *bind.WatchOpts, sink chan<- *BlockchainParametersUptimeLookbackWindowSet) (event.Subscription, error) {

	logs, sub, err := _BlockchainParameters.contract.WatchLogs(opts, "UptimeLookbackWindowSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BlockchainParametersUptimeLookbackWindowSet)
				if err := _BlockchainParameters.contract.UnpackLog(event, "UptimeLookbackWindowSet", log); err != nil {
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

// ParseUptimeLookbackWindowSet is a log parse operation binding the contract event 0x484a24d7faca8c4330aaf9ba5f131e6bd474ed6877a555511f39d16a1d71d15a.
//
// Solidity: event UptimeLookbackWindowSet(uint256 window, uint256 activationEpoch)
func (_BlockchainParameters *BlockchainParametersFilterer) ParseUptimeLookbackWindowSet(log types.Log) (*BlockchainParametersUptimeLookbackWindowSet, error) {
	event := new(BlockchainParametersUptimeLookbackWindowSet)
	if err := _BlockchainParameters.contract.UnpackLog(event, "UptimeLookbackWindowSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
