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

// DowntimeSlasherMetaData contains all meta data concerning the DowntimeSlasher contract.
var DowntimeSlasherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bitmap\",\"type\":\"bytes32\"}],\"name\":\"BitmapSetForInterval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"DowntimeSlashPerformed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"SlashableDowntimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"SlashingIncentivesSet\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bitmaps\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"groupMembershipHistoryIndex\",\"type\":\"uint256\"}],\"name\":\"groupMembershipAtBlock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastSlashedBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"setSlashingIncentives\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashableDowntime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashingIncentives\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_penalty\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_slashableDowntime\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"interval\",\"type\":\"uint256\"}],\"name\":\"setSlashableDowntime\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"getBitmapForInterval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"setBitmapForInterval\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"signerIndex\",\"type\":\"uint256\"}],\"name\":\"wasDownForInterval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"isBitmapSetForInterval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"startBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"endBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"signerIndices\",\"type\":\"uint256[]\"}],\"name\":\"wasDownForIntervals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"startBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"endBlocks\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"signerIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"groupMembershipHistoryIndex\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"validatorElectionLessers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"validatorElectionGreaters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"validatorElectionIndices\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"groupElectionLessers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"groupElectionGreaters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"groupElectionIndices\",\"type\":\"uint256[]\"}],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// DowntimeSlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use DowntimeSlasherMetaData.ABI instead.
var DowntimeSlasherABI = DowntimeSlasherMetaData.ABI

// DowntimeSlasher is an auto generated Go binding around an Ethereum contract.
type DowntimeSlasher struct {
	DowntimeSlasherCaller     // Read-only binding to the contract
	DowntimeSlasherTransactor // Write-only binding to the contract
	DowntimeSlasherFilterer   // Log filterer for contract events
}

// DowntimeSlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type DowntimeSlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DowntimeSlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DowntimeSlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DowntimeSlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DowntimeSlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DowntimeSlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DowntimeSlasherSession struct {
	Contract     *DowntimeSlasher  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DowntimeSlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DowntimeSlasherCallerSession struct {
	Contract *DowntimeSlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// DowntimeSlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DowntimeSlasherTransactorSession struct {
	Contract     *DowntimeSlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// DowntimeSlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type DowntimeSlasherRaw struct {
	Contract *DowntimeSlasher // Generic contract binding to access the raw methods on
}

// DowntimeSlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DowntimeSlasherCallerRaw struct {
	Contract *DowntimeSlasherCaller // Generic read-only contract binding to access the raw methods on
}

// DowntimeSlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DowntimeSlasherTransactorRaw struct {
	Contract *DowntimeSlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDowntimeSlasher creates a new instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasher(address common.Address, backend bind.ContractBackend) (*DowntimeSlasher, error) {
	contract, err := bindDowntimeSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasher{DowntimeSlasherCaller: DowntimeSlasherCaller{contract: contract}, DowntimeSlasherTransactor: DowntimeSlasherTransactor{contract: contract}, DowntimeSlasherFilterer: DowntimeSlasherFilterer{contract: contract}}, nil
}

// NewDowntimeSlasherCaller creates a new read-only instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasherCaller(address common.Address, caller bind.ContractCaller) (*DowntimeSlasherCaller, error) {
	contract, err := bindDowntimeSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherCaller{contract: contract}, nil
}

// NewDowntimeSlasherTransactor creates a new write-only instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*DowntimeSlasherTransactor, error) {
	contract, err := bindDowntimeSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherTransactor{contract: contract}, nil
}

// NewDowntimeSlasherFilterer creates a new log filterer instance of DowntimeSlasher, bound to a specific deployed contract.
func NewDowntimeSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*DowntimeSlasherFilterer, error) {
	contract, err := bindDowntimeSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherFilterer{contract: contract}, nil
}

// bindDowntimeSlasher binds a generic wrapper to an already deployed contract.
func bindDowntimeSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DowntimeSlasherABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseDowntimeSlasherABI parses the ABI
func ParseDowntimeSlasherABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(DowntimeSlasherABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DowntimeSlasher *DowntimeSlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DowntimeSlasher.Contract.DowntimeSlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DowntimeSlasher *DowntimeSlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.DowntimeSlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DowntimeSlasher *DowntimeSlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.DowntimeSlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DowntimeSlasher *DowntimeSlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DowntimeSlasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DowntimeSlasher *DowntimeSlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DowntimeSlasher *DowntimeSlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.contract.Transact(opts, method, params...)
}

// Bitmaps is a free data retrieval call binding the contract method 0x91275b4f.
//
// Solidity: function bitmaps(address , uint256 , uint256 ) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) Bitmaps(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "bitmaps", arg0, arg1, arg2)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Bitmaps is a free data retrieval call binding the contract method 0x91275b4f.
//
// Solidity: function bitmaps(address , uint256 , uint256 ) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) Bitmaps(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.Bitmaps(&_DowntimeSlasher.CallOpts, arg0, arg1, arg2)
}

// Bitmaps is a free data retrieval call binding the contract method 0x91275b4f.
//
// Solidity: function bitmaps(address , uint256 , uint256 ) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Bitmaps(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.Bitmaps(&_DowntimeSlasher.CallOpts, arg0, arg1, arg2)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _DowntimeSlasher.Contract.CheckProofOfPossession(&_DowntimeSlasher.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _DowntimeSlasher.Contract.CheckProofOfPossession(&_DowntimeSlasher.CallOpts, sender, blsKey, blsPop)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

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
func (_DowntimeSlasher *DowntimeSlasherSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _DowntimeSlasher.Contract.FractionMulExp(&_DowntimeSlasher.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _DowntimeSlasher.Contract.FractionMulExp(&_DowntimeSlasher.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBitmapForInterval is a free data retrieval call binding the contract method 0xa654a494.
//
// Solidity: function getBitmapForInterval(uint256 startBlock, uint256 endBlock) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetBitmapForInterval(opts *bind.CallOpts, startBlock *big.Int, endBlock *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getBitmapForInterval", startBlock, endBlock)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetBitmapForInterval is a free data retrieval call binding the contract method 0xa654a494.
//
// Solidity: function getBitmapForInterval(uint256 startBlock, uint256 endBlock) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) GetBitmapForInterval(startBlock *big.Int, endBlock *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetBitmapForInterval(&_DowntimeSlasher.CallOpts, startBlock, endBlock)
}

// GetBitmapForInterval is a free data retrieval call binding the contract method 0xa654a494.
//
// Solidity: function getBitmapForInterval(uint256 startBlock, uint256 endBlock) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetBitmapForInterval(startBlock *big.Int, endBlock *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetBitmapForInterval(&_DowntimeSlasher.CallOpts, startBlock, endBlock)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetBlockNumberFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetBlockNumberFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetEpochNumber() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumber(&_DowntimeSlasher.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetEpochNumber() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumber(&_DowntimeSlasher.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumberOfBlock(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochNumberOfBlock(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) GetEpochSize() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochSize(&_DowntimeSlasher.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetEpochSize() (*big.Int, error) {
	return _DowntimeSlasher.Contract.GetEpochSize(&_DowntimeSlasher.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetParentSealBitmap(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetParentSealBitmap(&_DowntimeSlasher.CallOpts, blockNumber)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetVerifiedSealBitmapFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.GetVerifiedSealBitmapFromHeader(&_DowntimeSlasher.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "getVersionNumber")

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
func (_DowntimeSlasher *DowntimeSlasherSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _DowntimeSlasher.Contract.GetVersionNumber(&_DowntimeSlasher.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _DowntimeSlasher.Contract.GetVersionNumber(&_DowntimeSlasher.CallOpts)
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) GroupMembershipAtBlock(opts *bind.CallOpts, validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "groupMembershipAtBlock", validator, blockNumber, groupMembershipHistoryIndex)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) GroupMembershipAtBlock(validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.GroupMembershipAtBlock(&_DowntimeSlasher.CallOpts, validator, blockNumber, groupMembershipHistoryIndex)
}

// GroupMembershipAtBlock is a free data retrieval call binding the contract method 0x88498aaf.
//
// Solidity: function groupMembershipAtBlock(address validator, uint256 blockNumber, uint256 groupMembershipHistoryIndex) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) GroupMembershipAtBlock(validator common.Address, blockNumber *big.Int, groupMembershipHistoryIndex *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.GroupMembershipAtBlock(&_DowntimeSlasher.CallOpts, validator, blockNumber, groupMembershipHistoryIndex)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) HashHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.HashHeader(&_DowntimeSlasher.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _DowntimeSlasher.Contract.HashHeader(&_DowntimeSlasher.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) Initialized() (bool, error) {
	return _DowntimeSlasher.Contract.Initialized(&_DowntimeSlasher.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Initialized() (bool, error) {
	return _DowntimeSlasher.Contract.Initialized(&_DowntimeSlasher.CallOpts)
}

// IsBitmapSetForInterval is a free data retrieval call binding the contract method 0x1bf0925b.
//
// Solidity: function isBitmapSetForInterval(uint256 startBlock, uint256 endBlock) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) IsBitmapSetForInterval(opts *bind.CallOpts, startBlock *big.Int, endBlock *big.Int) (bool, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "isBitmapSetForInterval", startBlock, endBlock)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBitmapSetForInterval is a free data retrieval call binding the contract method 0x1bf0925b.
//
// Solidity: function isBitmapSetForInterval(uint256 startBlock, uint256 endBlock) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) IsBitmapSetForInterval(startBlock *big.Int, endBlock *big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.IsBitmapSetForInterval(&_DowntimeSlasher.CallOpts, startBlock, endBlock)
}

// IsBitmapSetForInterval is a free data retrieval call binding the contract method 0x1bf0925b.
//
// Solidity: function isBitmapSetForInterval(uint256 startBlock, uint256 endBlock) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) IsBitmapSetForInterval(startBlock *big.Int, endBlock *big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.IsBitmapSetForInterval(&_DowntimeSlasher.CallOpts, startBlock, endBlock)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) IsOwner() (bool, error) {
	return _DowntimeSlasher.Contract.IsOwner(&_DowntimeSlasher.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) IsOwner() (bool, error) {
	return _DowntimeSlasher.Contract.IsOwner(&_DowntimeSlasher.CallOpts)
}

// LastSlashedBlock is a free data retrieval call binding the contract method 0x222d6b9f.
//
// Solidity: function lastSlashedBlock(address ) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) LastSlashedBlock(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "lastSlashedBlock", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSlashedBlock is a free data retrieval call binding the contract method 0x222d6b9f.
//
// Solidity: function lastSlashedBlock(address ) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) LastSlashedBlock(arg0 common.Address) (*big.Int, error) {
	return _DowntimeSlasher.Contract.LastSlashedBlock(&_DowntimeSlasher.CallOpts, arg0)
}

// LastSlashedBlock is a free data retrieval call binding the contract method 0x222d6b9f.
//
// Solidity: function lastSlashedBlock(address ) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) LastSlashedBlock(arg0 common.Address) (*big.Int, error) {
	return _DowntimeSlasher.Contract.LastSlashedBlock(&_DowntimeSlasher.CallOpts, arg0)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSize(&_DowntimeSlasher.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSize(&_DowntimeSlasher.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSizeInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.MinQuorumSizeInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInCurrentSet(&_DowntimeSlasher.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInSet(&_DowntimeSlasher.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _DowntimeSlasher.Contract.NumberValidatorsInSet(&_DowntimeSlasher.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) Owner() (common.Address, error) {
	return _DowntimeSlasher.Contract.Owner(&_DowntimeSlasher.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Owner() (common.Address, error) {
	return _DowntimeSlasher.Contract.Owner(&_DowntimeSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) Registry() (common.Address, error) {
	return _DowntimeSlasher.Contract.Registry(&_DowntimeSlasher.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) Registry() (common.Address, error) {
	return _DowntimeSlasher.Contract.Registry(&_DowntimeSlasher.CallOpts)
}

// SlashableDowntime is a free data retrieval call binding the contract method 0x4227d971.
//
// Solidity: function slashableDowntime() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCaller) SlashableDowntime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "slashableDowntime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashableDowntime is a free data retrieval call binding the contract method 0x4227d971.
//
// Solidity: function slashableDowntime() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherSession) SlashableDowntime() (*big.Int, error) {
	return _DowntimeSlasher.Contract.SlashableDowntime(&_DowntimeSlasher.CallOpts)
}

// SlashableDowntime is a free data retrieval call binding the contract method 0x4227d971.
//
// Solidity: function slashableDowntime() view returns(uint256)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) SlashableDowntime() (*big.Int, error) {
	return _DowntimeSlasher.Contract.SlashableDowntime(&_DowntimeSlasher.CallOpts)
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() view returns(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherCaller) SlashingIncentives(opts *bind.CallOpts) (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "slashingIncentives")

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
func (_DowntimeSlasher *DowntimeSlasherSession) SlashingIncentives() (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	return _DowntimeSlasher.Contract.SlashingIncentives(&_DowntimeSlasher.CallOpts)
}

// SlashingIncentives is a free data retrieval call binding the contract method 0x0a05cd84.
//
// Solidity: function slashingIncentives() view returns(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) SlashingIncentives() (struct {
	Penalty *big.Int
	Reward  *big.Int
}, error) {
	return _DowntimeSlasher.Contract.SlashingIncentives(&_DowntimeSlasher.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromCurrentSet(&_DowntimeSlasher.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromCurrentSet(&_DowntimeSlasher.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromSet(&_DowntimeSlasher.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _DowntimeSlasher.Contract.ValidatorSignerAddressFromSet(&_DowntimeSlasher.CallOpts, index, blockNumber)
}

// WasDownForInterval is a free data retrieval call binding the contract method 0xec611ffc.
//
// Solidity: function wasDownForInterval(uint256 startBlock, uint256 endBlock, uint256 signerIndex) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) WasDownForInterval(opts *bind.CallOpts, startBlock *big.Int, endBlock *big.Int, signerIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "wasDownForInterval", startBlock, endBlock, signerIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasDownForInterval is a free data retrieval call binding the contract method 0xec611ffc.
//
// Solidity: function wasDownForInterval(uint256 startBlock, uint256 endBlock, uint256 signerIndex) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) WasDownForInterval(startBlock *big.Int, endBlock *big.Int, signerIndex *big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.WasDownForInterval(&_DowntimeSlasher.CallOpts, startBlock, endBlock, signerIndex)
}

// WasDownForInterval is a free data retrieval call binding the contract method 0xec611ffc.
//
// Solidity: function wasDownForInterval(uint256 startBlock, uint256 endBlock, uint256 signerIndex) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) WasDownForInterval(startBlock *big.Int, endBlock *big.Int, signerIndex *big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.WasDownForInterval(&_DowntimeSlasher.CallOpts, startBlock, endBlock, signerIndex)
}

// WasDownForIntervals is a free data retrieval call binding the contract method 0xe252e904.
//
// Solidity: function wasDownForIntervals(uint256[] startBlocks, uint256[] endBlocks, uint256[] signerIndices) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCaller) WasDownForIntervals(opts *bind.CallOpts, startBlocks []*big.Int, endBlocks []*big.Int, signerIndices []*big.Int) (bool, error) {
	var out []interface{}
	err := _DowntimeSlasher.contract.Call(opts, &out, "wasDownForIntervals", startBlocks, endBlocks, signerIndices)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WasDownForIntervals is a free data retrieval call binding the contract method 0xe252e904.
//
// Solidity: function wasDownForIntervals(uint256[] startBlocks, uint256[] endBlocks, uint256[] signerIndices) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherSession) WasDownForIntervals(startBlocks []*big.Int, endBlocks []*big.Int, signerIndices []*big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.WasDownForIntervals(&_DowntimeSlasher.CallOpts, startBlocks, endBlocks, signerIndices)
}

// WasDownForIntervals is a free data retrieval call binding the contract method 0xe252e904.
//
// Solidity: function wasDownForIntervals(uint256[] startBlocks, uint256[] endBlocks, uint256[] signerIndices) view returns(bool)
func (_DowntimeSlasher *DowntimeSlasherCallerSession) WasDownForIntervals(startBlocks []*big.Int, endBlocks []*big.Int, signerIndices []*big.Int) (bool, error) {
	return _DowntimeSlasher.Contract.WasDownForIntervals(&_DowntimeSlasher.CallOpts, startBlocks, endBlocks, signerIndices)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward, uint256 _slashableDowntime) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _penalty *big.Int, _reward *big.Int, _slashableDowntime *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "initialize", registryAddress, _penalty, _reward, _slashableDowntime)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward, uint256 _slashableDowntime) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) Initialize(registryAddress common.Address, _penalty *big.Int, _reward *big.Int, _slashableDowntime *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Initialize(&_DowntimeSlasher.TransactOpts, registryAddress, _penalty, _reward, _slashableDowntime)
}

// Initialize is a paid mutator transaction binding the contract method 0x4ec81af1.
//
// Solidity: function initialize(address registryAddress, uint256 _penalty, uint256 _reward, uint256 _slashableDowntime) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) Initialize(registryAddress common.Address, _penalty *big.Int, _reward *big.Int, _slashableDowntime *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Initialize(&_DowntimeSlasher.TransactOpts, registryAddress, _penalty, _reward, _slashableDowntime)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DowntimeSlasher *DowntimeSlasherSession) RenounceOwnership() (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.RenounceOwnership(&_DowntimeSlasher.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.RenounceOwnership(&_DowntimeSlasher.TransactOpts)
}

// SetBitmapForInterval is a paid mutator transaction binding the contract method 0xfafec0f6.
//
// Solidity: function setBitmapForInterval(uint256 startBlock, uint256 endBlock) returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetBitmapForInterval(opts *bind.TransactOpts, startBlock *big.Int, endBlock *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setBitmapForInterval", startBlock, endBlock)
}

// SetBitmapForInterval is a paid mutator transaction binding the contract method 0xfafec0f6.
//
// Solidity: function setBitmapForInterval(uint256 startBlock, uint256 endBlock) returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherSession) SetBitmapForInterval(startBlock *big.Int, endBlock *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetBitmapForInterval(&_DowntimeSlasher.TransactOpts, startBlock, endBlock)
}

// SetBitmapForInterval is a paid mutator transaction binding the contract method 0xfafec0f6.
//
// Solidity: function setBitmapForInterval(uint256 startBlock, uint256 endBlock) returns(bytes32)
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetBitmapForInterval(startBlock *big.Int, endBlock *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetBitmapForInterval(&_DowntimeSlasher.TransactOpts, startBlock, endBlock)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetRegistry(&_DowntimeSlasher.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetRegistry(&_DowntimeSlasher.TransactOpts, registryAddress)
}

// SetSlashableDowntime is a paid mutator transaction binding the contract method 0x4d643e17.
//
// Solidity: function setSlashableDowntime(uint256 interval) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetSlashableDowntime(opts *bind.TransactOpts, interval *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setSlashableDowntime", interval)
}

// SetSlashableDowntime is a paid mutator transaction binding the contract method 0x4d643e17.
//
// Solidity: function setSlashableDowntime(uint256 interval) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) SetSlashableDowntime(interval *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashableDowntime(&_DowntimeSlasher.TransactOpts, interval)
}

// SetSlashableDowntime is a paid mutator transaction binding the contract method 0x4d643e17.
//
// Solidity: function setSlashableDowntime(uint256 interval) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetSlashableDowntime(interval *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashableDowntime(&_DowntimeSlasher.TransactOpts, interval)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) SetSlashingIncentives(opts *bind.TransactOpts, penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "setSlashingIncentives", penalty, reward)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) SetSlashingIncentives(penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashingIncentives(&_DowntimeSlasher.TransactOpts, penalty, reward)
}

// SetSlashingIncentives is a paid mutator transaction binding the contract method 0xbd0d9979.
//
// Solidity: function setSlashingIncentives(uint256 penalty, uint256 reward) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) SetSlashingIncentives(penalty *big.Int, reward *big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.SetSlashingIncentives(&_DowntimeSlasher.TransactOpts, penalty, reward)
}

// Slash is a paid mutator transaction binding the contract method 0x190ad68b.
//
// Solidity: function slash(uint256[] startBlocks, uint256[] endBlocks, uint256[] signerIndices, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) Slash(opts *bind.TransactOpts, startBlocks []*big.Int, endBlocks []*big.Int, signerIndices []*big.Int, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "slash", startBlocks, endBlocks, signerIndices, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0x190ad68b.
//
// Solidity: function slash(uint256[] startBlocks, uint256[] endBlocks, uint256[] signerIndices, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) Slash(startBlocks []*big.Int, endBlocks []*big.Int, signerIndices []*big.Int, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Slash(&_DowntimeSlasher.TransactOpts, startBlocks, endBlocks, signerIndices, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// Slash is a paid mutator transaction binding the contract method 0x190ad68b.
//
// Solidity: function slash(uint256[] startBlocks, uint256[] endBlocks, uint256[] signerIndices, uint256 groupMembershipHistoryIndex, address[] validatorElectionLessers, address[] validatorElectionGreaters, uint256[] validatorElectionIndices, address[] groupElectionLessers, address[] groupElectionGreaters, uint256[] groupElectionIndices) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) Slash(startBlocks []*big.Int, endBlocks []*big.Int, signerIndices []*big.Int, groupMembershipHistoryIndex *big.Int, validatorElectionLessers []common.Address, validatorElectionGreaters []common.Address, validatorElectionIndices []*big.Int, groupElectionLessers []common.Address, groupElectionGreaters []common.Address, groupElectionIndices []*big.Int) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.Slash(&_DowntimeSlasher.TransactOpts, startBlocks, endBlocks, signerIndices, groupMembershipHistoryIndex, validatorElectionLessers, validatorElectionGreaters, validatorElectionIndices, groupElectionLessers, groupElectionGreaters, groupElectionIndices)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DowntimeSlasher *DowntimeSlasherSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.TransferOwnership(&_DowntimeSlasher.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_DowntimeSlasher *DowntimeSlasherTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _DowntimeSlasher.Contract.TransferOwnership(&_DowntimeSlasher.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_DowntimeSlasher *DowntimeSlasherFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _DowntimeSlasher.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "BitmapSetForInterval":
		event, err = _DowntimeSlasher.ParseBitmapSetForInterval(log)
	case "DowntimeSlashPerformed":
		event, err = _DowntimeSlasher.ParseDowntimeSlashPerformed(log)
	case "OwnershipTransferred":
		event, err = _DowntimeSlasher.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _DowntimeSlasher.ParseRegistrySet(log)
	case "SlashableDowntimeSet":
		event, err = _DowntimeSlasher.ParseSlashableDowntimeSet(log)
	case "SlashingIncentivesSet":
		event, err = _DowntimeSlasher.ParseSlashingIncentivesSet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// DowntimeSlasherBitmapSetForIntervalIterator is returned from FilterBitmapSetForInterval and is used to iterate over the raw logs and unpacked data for BitmapSetForInterval events raised by the DowntimeSlasher contract.
type DowntimeSlasherBitmapSetForIntervalIterator struct {
	Event *DowntimeSlasherBitmapSetForInterval // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherBitmapSetForIntervalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherBitmapSetForInterval)
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
		it.Event = new(DowntimeSlasherBitmapSetForInterval)
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
func (it *DowntimeSlasherBitmapSetForIntervalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherBitmapSetForIntervalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherBitmapSetForInterval represents a BitmapSetForInterval event raised by the DowntimeSlasher contract.
type DowntimeSlasherBitmapSetForInterval struct {
	Sender     common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	Bitmap     [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBitmapSetForInterval is a free log retrieval operation binding the contract event 0x0aa96aa275a5f936eed2a6a01f082594744dcc2510f575101366f8f479f03235.
//
// Solidity: event BitmapSetForInterval(address indexed sender, uint256 indexed startBlock, uint256 indexed endBlock, bytes32 bitmap)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterBitmapSetForInterval(opts *bind.FilterOpts, sender []common.Address, startBlock []*big.Int, endBlock []*big.Int) (*DowntimeSlasherBitmapSetForIntervalIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "BitmapSetForInterval", senderRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherBitmapSetForIntervalIterator{contract: _DowntimeSlasher.contract, event: "BitmapSetForInterval", logs: logs, sub: sub}, nil
}

// WatchBitmapSetForInterval is a free log subscription operation binding the contract event 0x0aa96aa275a5f936eed2a6a01f082594744dcc2510f575101366f8f479f03235.
//
// Solidity: event BitmapSetForInterval(address indexed sender, uint256 indexed startBlock, uint256 indexed endBlock, bytes32 bitmap)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchBitmapSetForInterval(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherBitmapSetForInterval, sender []common.Address, startBlock []*big.Int, endBlock []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "BitmapSetForInterval", senderRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherBitmapSetForInterval)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "BitmapSetForInterval", log); err != nil {
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

// ParseBitmapSetForInterval is a log parse operation binding the contract event 0x0aa96aa275a5f936eed2a6a01f082594744dcc2510f575101366f8f479f03235.
//
// Solidity: event BitmapSetForInterval(address indexed sender, uint256 indexed startBlock, uint256 indexed endBlock, bytes32 bitmap)
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseBitmapSetForInterval(log types.Log) (*DowntimeSlasherBitmapSetForInterval, error) {
	event := new(DowntimeSlasherBitmapSetForInterval)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "BitmapSetForInterval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DowntimeSlasherDowntimeSlashPerformedIterator is returned from FilterDowntimeSlashPerformed and is used to iterate over the raw logs and unpacked data for DowntimeSlashPerformed events raised by the DowntimeSlasher contract.
type DowntimeSlasherDowntimeSlashPerformedIterator struct {
	Event *DowntimeSlasherDowntimeSlashPerformed // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherDowntimeSlashPerformedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherDowntimeSlashPerformed)
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
		it.Event = new(DowntimeSlasherDowntimeSlashPerformed)
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
func (it *DowntimeSlasherDowntimeSlashPerformedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherDowntimeSlashPerformedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherDowntimeSlashPerformed represents a DowntimeSlashPerformed event raised by the DowntimeSlasher contract.
type DowntimeSlasherDowntimeSlashPerformed struct {
	Validator  common.Address
	StartBlock *big.Int
	EndBlock   *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDowntimeSlashPerformed is a free log retrieval operation binding the contract event 0x229d63d990a0f1068a86ee5bdce0b23fe156ff5d5174cc634d5da8ed3618e0c9.
//
// Solidity: event DowntimeSlashPerformed(address indexed validator, uint256 indexed startBlock, uint256 indexed endBlock)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterDowntimeSlashPerformed(opts *bind.FilterOpts, validator []common.Address, startBlock []*big.Int, endBlock []*big.Int) (*DowntimeSlasherDowntimeSlashPerformedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "DowntimeSlashPerformed", validatorRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherDowntimeSlashPerformedIterator{contract: _DowntimeSlasher.contract, event: "DowntimeSlashPerformed", logs: logs, sub: sub}, nil
}

// WatchDowntimeSlashPerformed is a free log subscription operation binding the contract event 0x229d63d990a0f1068a86ee5bdce0b23fe156ff5d5174cc634d5da8ed3618e0c9.
//
// Solidity: event DowntimeSlashPerformed(address indexed validator, uint256 indexed startBlock, uint256 indexed endBlock)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchDowntimeSlashPerformed(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherDowntimeSlashPerformed, validator []common.Address, startBlock []*big.Int, endBlock []*big.Int) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var startBlockRule []interface{}
	for _, startBlockItem := range startBlock {
		startBlockRule = append(startBlockRule, startBlockItem)
	}
	var endBlockRule []interface{}
	for _, endBlockItem := range endBlock {
		endBlockRule = append(endBlockRule, endBlockItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "DowntimeSlashPerformed", validatorRule, startBlockRule, endBlockRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherDowntimeSlashPerformed)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "DowntimeSlashPerformed", log); err != nil {
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

// ParseDowntimeSlashPerformed is a log parse operation binding the contract event 0x229d63d990a0f1068a86ee5bdce0b23fe156ff5d5174cc634d5da8ed3618e0c9.
//
// Solidity: event DowntimeSlashPerformed(address indexed validator, uint256 indexed startBlock, uint256 indexed endBlock)
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseDowntimeSlashPerformed(log types.Log) (*DowntimeSlasherDowntimeSlashPerformed, error) {
	event := new(DowntimeSlasherDowntimeSlashPerformed)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "DowntimeSlashPerformed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DowntimeSlasherOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the DowntimeSlasher contract.
type DowntimeSlasherOwnershipTransferredIterator struct {
	Event *DowntimeSlasherOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherOwnershipTransferred)
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
		it.Event = new(DowntimeSlasherOwnershipTransferred)
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
func (it *DowntimeSlasherOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherOwnershipTransferred represents a OwnershipTransferred event raised by the DowntimeSlasher contract.
type DowntimeSlasherOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DowntimeSlasherOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherOwnershipTransferredIterator{contract: _DowntimeSlasher.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherOwnershipTransferred)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseOwnershipTransferred(log types.Log) (*DowntimeSlasherOwnershipTransferred, error) {
	event := new(DowntimeSlasherOwnershipTransferred)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DowntimeSlasherRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the DowntimeSlasher contract.
type DowntimeSlasherRegistrySetIterator struct {
	Event *DowntimeSlasherRegistrySet // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherRegistrySet)
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
		it.Event = new(DowntimeSlasherRegistrySet)
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
func (it *DowntimeSlasherRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherRegistrySet represents a RegistrySet event raised by the DowntimeSlasher contract.
type DowntimeSlasherRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*DowntimeSlasherRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherRegistrySetIterator{contract: _DowntimeSlasher.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherRegistrySet)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseRegistrySet(log types.Log) (*DowntimeSlasherRegistrySet, error) {
	event := new(DowntimeSlasherRegistrySet)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DowntimeSlasherSlashableDowntimeSetIterator is returned from FilterSlashableDowntimeSet and is used to iterate over the raw logs and unpacked data for SlashableDowntimeSet events raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashableDowntimeSetIterator struct {
	Event *DowntimeSlasherSlashableDowntimeSet // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherSlashableDowntimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherSlashableDowntimeSet)
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
		it.Event = new(DowntimeSlasherSlashableDowntimeSet)
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
func (it *DowntimeSlasherSlashableDowntimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherSlashableDowntimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherSlashableDowntimeSet represents a SlashableDowntimeSet event raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashableDowntimeSet struct {
	Interval *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSlashableDowntimeSet is a free log retrieval operation binding the contract event 0xc3293b70d45615822039f6f13747ece88efbbb4e645c42070413a6c3fd21d771.
//
// Solidity: event SlashableDowntimeSet(uint256 interval)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterSlashableDowntimeSet(opts *bind.FilterOpts) (*DowntimeSlasherSlashableDowntimeSetIterator, error) {

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "SlashableDowntimeSet")
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherSlashableDowntimeSetIterator{contract: _DowntimeSlasher.contract, event: "SlashableDowntimeSet", logs: logs, sub: sub}, nil
}

// WatchSlashableDowntimeSet is a free log subscription operation binding the contract event 0xc3293b70d45615822039f6f13747ece88efbbb4e645c42070413a6c3fd21d771.
//
// Solidity: event SlashableDowntimeSet(uint256 interval)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchSlashableDowntimeSet(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherSlashableDowntimeSet) (event.Subscription, error) {

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "SlashableDowntimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherSlashableDowntimeSet)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashableDowntimeSet", log); err != nil {
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

// ParseSlashableDowntimeSet is a log parse operation binding the contract event 0xc3293b70d45615822039f6f13747ece88efbbb4e645c42070413a6c3fd21d771.
//
// Solidity: event SlashableDowntimeSet(uint256 interval)
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseSlashableDowntimeSet(log types.Log) (*DowntimeSlasherSlashableDowntimeSet, error) {
	event := new(DowntimeSlasherSlashableDowntimeSet)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashableDowntimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DowntimeSlasherSlashingIncentivesSetIterator is returned from FilterSlashingIncentivesSet and is used to iterate over the raw logs and unpacked data for SlashingIncentivesSet events raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashingIncentivesSetIterator struct {
	Event *DowntimeSlasherSlashingIncentivesSet // Event containing the contract specifics and raw log

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
func (it *DowntimeSlasherSlashingIncentivesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DowntimeSlasherSlashingIncentivesSet)
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
		it.Event = new(DowntimeSlasherSlashingIncentivesSet)
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
func (it *DowntimeSlasherSlashingIncentivesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DowntimeSlasherSlashingIncentivesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DowntimeSlasherSlashingIncentivesSet represents a SlashingIncentivesSet event raised by the DowntimeSlasher contract.
type DowntimeSlasherSlashingIncentivesSet struct {
	Penalty *big.Int
	Reward  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSlashingIncentivesSet is a free log retrieval operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherFilterer) FilterSlashingIncentivesSet(opts *bind.FilterOpts) (*DowntimeSlasherSlashingIncentivesSetIterator, error) {

	logs, sub, err := _DowntimeSlasher.contract.FilterLogs(opts, "SlashingIncentivesSet")
	if err != nil {
		return nil, err
	}
	return &DowntimeSlasherSlashingIncentivesSetIterator{contract: _DowntimeSlasher.contract, event: "SlashingIncentivesSet", logs: logs, sub: sub}, nil
}

// WatchSlashingIncentivesSet is a free log subscription operation binding the contract event 0x716dc7c34384df36c6ccc5a2949f2ce9b019f5d4075ef39139a80038a4fdd1c3.
//
// Solidity: event SlashingIncentivesSet(uint256 penalty, uint256 reward)
func (_DowntimeSlasher *DowntimeSlasherFilterer) WatchSlashingIncentivesSet(opts *bind.WatchOpts, sink chan<- *DowntimeSlasherSlashingIncentivesSet) (event.Subscription, error) {

	logs, sub, err := _DowntimeSlasher.contract.WatchLogs(opts, "SlashingIncentivesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DowntimeSlasherSlashingIncentivesSet)
				if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashingIncentivesSet", log); err != nil {
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
func (_DowntimeSlasher *DowntimeSlasherFilterer) ParseSlashingIncentivesSet(log types.Log) (*DowntimeSlasherSlashingIncentivesSet, error) {
	event := new(DowntimeSlasherSlashingIncentivesSet)
	if err := _DowntimeSlasher.contract.UnpackLog(event, "SlashingIncentivesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
