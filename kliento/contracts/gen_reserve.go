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

// ReserveMetaData contains all meta data concerning the Reserve contract.
var ReserveMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"symbols\",\"type\":\"bytes32[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"}],\"name\":\"AssetAllocationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"DailySpendingRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exchangeSpender\",\"type\":\"address\"}],\"name\":\"ExchangeSpenderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"exchangeSpender\",\"type\":\"address\"}],\"name\":\"ExchangeSpenderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otherReserveAddress\",\"type\":\"address\"}],\"name\":\"OtherReserveAddressAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"otherReserveAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"OtherReserveAddressRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ReserveGoldTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"SpenderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"SpenderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TobinTaxReserveRatioSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TobinTaxSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TobinTaxStalenessThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"TokenRemoved\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assetAllocationSymbols\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"assetAllocationWeights\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"exchangeSpenderAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"frozenReserveGoldDays\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"frozenReserveGoldStartBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"frozenReserveGoldStartDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isExchangeSpender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOtherReserveAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isSpender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastSpendingDay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"otherReserveAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spendingLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tobinTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tobinTaxCache\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"numerator\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"timestamp\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tobinTaxReserveRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tobinTaxStalenessThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_tobinTaxStalenessThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_spendingRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_frozenGold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_frozenDays\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"_assetAllocationSymbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_assetAllocationWeights\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_tobinTax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tobinTaxReserveRatio\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setTobinTaxStalenessThreshold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setTobinTax\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setTobinTaxReserveRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setDailySpendingRatio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDailySpendingRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"frozenGold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frozenDays\",\"type\":\"uint256\"}],\"name\":\"setFrozenGold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"symbols\",\"type\":\"bytes32[]\"},{\"internalType\":\"uint256[]\",\"name\":\"weights\",\"type\":\"uint256[]\"}],\"name\":\"setAssetAllocations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"}],\"name\":\"addOtherReserveAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserveAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeOtherReserveAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"addSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"removeSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"addExchangeSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"removeExchangeSpender\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExchangeSpenders\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferGold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferExchangeGold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"getOrComputeTobinTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOtherReserveAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetAllocationSymbols\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAssetAllocationWeights\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUnfrozenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserveGoldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOtherReserveAddressesGoldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getUnfrozenReserveGoldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFrozenReserveGoldBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReserveRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ReserveABI is the input ABI used to generate the binding from.
// Deprecated: Use ReserveMetaData.ABI instead.
var ReserveABI = ReserveMetaData.ABI

// Reserve is an auto generated Go binding around an Ethereum contract.
type Reserve struct {
	ReserveCaller     // Read-only binding to the contract
	ReserveTransactor // Write-only binding to the contract
	ReserveFilterer   // Log filterer for contract events
}

// ReserveCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveSession struct {
	Contract     *Reserve          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReserveCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveCallerSession struct {
	Contract *ReserveCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ReserveTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveTransactorSession struct {
	Contract     *ReserveTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ReserveRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveRaw struct {
	Contract *Reserve // Generic contract binding to access the raw methods on
}

// ReserveCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveCallerRaw struct {
	Contract *ReserveCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveTransactorRaw struct {
	Contract *ReserveTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserve creates a new instance of Reserve, bound to a specific deployed contract.
func NewReserve(address common.Address, backend bind.ContractBackend) (*Reserve, error) {
	contract, err := bindReserve(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Reserve{ReserveCaller: ReserveCaller{contract: contract}, ReserveTransactor: ReserveTransactor{contract: contract}, ReserveFilterer: ReserveFilterer{contract: contract}}, nil
}

// NewReserveCaller creates a new read-only instance of Reserve, bound to a specific deployed contract.
func NewReserveCaller(address common.Address, caller bind.ContractCaller) (*ReserveCaller, error) {
	contract, err := bindReserve(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveCaller{contract: contract}, nil
}

// NewReserveTransactor creates a new write-only instance of Reserve, bound to a specific deployed contract.
func NewReserveTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveTransactor, error) {
	contract, err := bindReserve(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveTransactor{contract: contract}, nil
}

// NewReserveFilterer creates a new log filterer instance of Reserve, bound to a specific deployed contract.
func NewReserveFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveFilterer, error) {
	contract, err := bindReserve(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveFilterer{contract: contract}, nil
}

// bindReserve binds a generic wrapper to an already deployed contract.
func bindReserve(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseReserveABI parses the ABI
func ParseReserveABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.ReserveCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.ReserveTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Reserve *ReserveCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Reserve.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Reserve *ReserveTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Reserve *ReserveTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Reserve.Contract.contract.Transact(opts, method, params...)
}

// AssetAllocationSymbols is a free data retrieval call binding the contract method 0x0db279be.
//
// Solidity: function assetAllocationSymbols(uint256 ) view returns(bytes32)
func (_Reserve *ReserveCaller) AssetAllocationSymbols(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "assetAllocationSymbols", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetAllocationSymbols is a free data retrieval call binding the contract method 0x0db279be.
//
// Solidity: function assetAllocationSymbols(uint256 ) view returns(bytes32)
func (_Reserve *ReserveSession) AssetAllocationSymbols(arg0 *big.Int) ([32]byte, error) {
	return _Reserve.Contract.AssetAllocationSymbols(&_Reserve.CallOpts, arg0)
}

// AssetAllocationSymbols is a free data retrieval call binding the contract method 0x0db279be.
//
// Solidity: function assetAllocationSymbols(uint256 ) view returns(bytes32)
func (_Reserve *ReserveCallerSession) AssetAllocationSymbols(arg0 *big.Int) ([32]byte, error) {
	return _Reserve.Contract.AssetAllocationSymbols(&_Reserve.CallOpts, arg0)
}

// AssetAllocationWeights is a free data retrieval call binding the contract method 0xec4f797b.
//
// Solidity: function assetAllocationWeights(bytes32 ) view returns(uint256)
func (_Reserve *ReserveCaller) AssetAllocationWeights(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "assetAllocationWeights", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AssetAllocationWeights is a free data retrieval call binding the contract method 0xec4f797b.
//
// Solidity: function assetAllocationWeights(bytes32 ) view returns(uint256)
func (_Reserve *ReserveSession) AssetAllocationWeights(arg0 [32]byte) (*big.Int, error) {
	return _Reserve.Contract.AssetAllocationWeights(&_Reserve.CallOpts, arg0)
}

// AssetAllocationWeights is a free data retrieval call binding the contract method 0xec4f797b.
//
// Solidity: function assetAllocationWeights(bytes32 ) view returns(uint256)
func (_Reserve *ReserveCallerSession) AssetAllocationWeights(arg0 [32]byte) (*big.Int, error) {
	return _Reserve.Contract.AssetAllocationWeights(&_Reserve.CallOpts, arg0)
}

// ExchangeSpenderAddresses is a free data retrieval call binding the contract method 0x5a18b08b.
//
// Solidity: function exchangeSpenderAddresses(uint256 ) view returns(address)
func (_Reserve *ReserveCaller) ExchangeSpenderAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "exchangeSpenderAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExchangeSpenderAddresses is a free data retrieval call binding the contract method 0x5a18b08b.
//
// Solidity: function exchangeSpenderAddresses(uint256 ) view returns(address)
func (_Reserve *ReserveSession) ExchangeSpenderAddresses(arg0 *big.Int) (common.Address, error) {
	return _Reserve.Contract.ExchangeSpenderAddresses(&_Reserve.CallOpts, arg0)
}

// ExchangeSpenderAddresses is a free data retrieval call binding the contract method 0x5a18b08b.
//
// Solidity: function exchangeSpenderAddresses(uint256 ) view returns(address)
func (_Reserve *ReserveCallerSession) ExchangeSpenderAddresses(arg0 *big.Int) (common.Address, error) {
	return _Reserve.Contract.ExchangeSpenderAddresses(&_Reserve.CallOpts, arg0)
}

// FrozenReserveGoldDays is a free data retrieval call binding the contract method 0x7090db4e.
//
// Solidity: function frozenReserveGoldDays() view returns(uint256)
func (_Reserve *ReserveCaller) FrozenReserveGoldDays(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "frozenReserveGoldDays")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenReserveGoldDays is a free data retrieval call binding the contract method 0x7090db4e.
//
// Solidity: function frozenReserveGoldDays() view returns(uint256)
func (_Reserve *ReserveSession) FrozenReserveGoldDays() (*big.Int, error) {
	return _Reserve.Contract.FrozenReserveGoldDays(&_Reserve.CallOpts)
}

// FrozenReserveGoldDays is a free data retrieval call binding the contract method 0x7090db4e.
//
// Solidity: function frozenReserveGoldDays() view returns(uint256)
func (_Reserve *ReserveCallerSession) FrozenReserveGoldDays() (*big.Int, error) {
	return _Reserve.Contract.FrozenReserveGoldDays(&_Reserve.CallOpts)
}

// FrozenReserveGoldStartBalance is a free data retrieval call binding the contract method 0x03d835f3.
//
// Solidity: function frozenReserveGoldStartBalance() view returns(uint256)
func (_Reserve *ReserveCaller) FrozenReserveGoldStartBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "frozenReserveGoldStartBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenReserveGoldStartBalance is a free data retrieval call binding the contract method 0x03d835f3.
//
// Solidity: function frozenReserveGoldStartBalance() view returns(uint256)
func (_Reserve *ReserveSession) FrozenReserveGoldStartBalance() (*big.Int, error) {
	return _Reserve.Contract.FrozenReserveGoldStartBalance(&_Reserve.CallOpts)
}

// FrozenReserveGoldStartBalance is a free data retrieval call binding the contract method 0x03d835f3.
//
// Solidity: function frozenReserveGoldStartBalance() view returns(uint256)
func (_Reserve *ReserveCallerSession) FrozenReserveGoldStartBalance() (*big.Int, error) {
	return _Reserve.Contract.FrozenReserveGoldStartBalance(&_Reserve.CallOpts)
}

// FrozenReserveGoldStartDay is a free data retrieval call binding the contract method 0x81b861a6.
//
// Solidity: function frozenReserveGoldStartDay() view returns(uint256)
func (_Reserve *ReserveCaller) FrozenReserveGoldStartDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "frozenReserveGoldStartDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FrozenReserveGoldStartDay is a free data retrieval call binding the contract method 0x81b861a6.
//
// Solidity: function frozenReserveGoldStartDay() view returns(uint256)
func (_Reserve *ReserveSession) FrozenReserveGoldStartDay() (*big.Int, error) {
	return _Reserve.Contract.FrozenReserveGoldStartDay(&_Reserve.CallOpts)
}

// FrozenReserveGoldStartDay is a free data retrieval call binding the contract method 0x81b861a6.
//
// Solidity: function frozenReserveGoldStartDay() view returns(uint256)
func (_Reserve *ReserveCallerSession) FrozenReserveGoldStartDay() (*big.Int, error) {
	return _Reserve.Contract.FrozenReserveGoldStartDay(&_Reserve.CallOpts)
}

// GetAssetAllocationSymbols is a free data retrieval call binding the contract method 0x8438796a.
//
// Solidity: function getAssetAllocationSymbols() view returns(bytes32[])
func (_Reserve *ReserveCaller) GetAssetAllocationSymbols(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getAssetAllocationSymbols")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetAssetAllocationSymbols is a free data retrieval call binding the contract method 0x8438796a.
//
// Solidity: function getAssetAllocationSymbols() view returns(bytes32[])
func (_Reserve *ReserveSession) GetAssetAllocationSymbols() ([][32]byte, error) {
	return _Reserve.Contract.GetAssetAllocationSymbols(&_Reserve.CallOpts)
}

// GetAssetAllocationSymbols is a free data retrieval call binding the contract method 0x8438796a.
//
// Solidity: function getAssetAllocationSymbols() view returns(bytes32[])
func (_Reserve *ReserveCallerSession) GetAssetAllocationSymbols() ([][32]byte, error) {
	return _Reserve.Contract.GetAssetAllocationSymbols(&_Reserve.CallOpts)
}

// GetAssetAllocationWeights is a free data retrieval call binding the contract method 0xe50a6c1e.
//
// Solidity: function getAssetAllocationWeights() view returns(uint256[])
func (_Reserve *ReserveCaller) GetAssetAllocationWeights(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getAssetAllocationWeights")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAssetAllocationWeights is a free data retrieval call binding the contract method 0xe50a6c1e.
//
// Solidity: function getAssetAllocationWeights() view returns(uint256[])
func (_Reserve *ReserveSession) GetAssetAllocationWeights() ([]*big.Int, error) {
	return _Reserve.Contract.GetAssetAllocationWeights(&_Reserve.CallOpts)
}

// GetAssetAllocationWeights is a free data retrieval call binding the contract method 0xe50a6c1e.
//
// Solidity: function getAssetAllocationWeights() view returns(uint256[])
func (_Reserve *ReserveCallerSession) GetAssetAllocationWeights() ([]*big.Int, error) {
	return _Reserve.Contract.GetAssetAllocationWeights(&_Reserve.CallOpts)
}

// GetDailySpendingRatio is a free data retrieval call binding the contract method 0x7897a78e.
//
// Solidity: function getDailySpendingRatio() view returns(uint256)
func (_Reserve *ReserveCaller) GetDailySpendingRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getDailySpendingRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDailySpendingRatio is a free data retrieval call binding the contract method 0x7897a78e.
//
// Solidity: function getDailySpendingRatio() view returns(uint256)
func (_Reserve *ReserveSession) GetDailySpendingRatio() (*big.Int, error) {
	return _Reserve.Contract.GetDailySpendingRatio(&_Reserve.CallOpts)
}

// GetDailySpendingRatio is a free data retrieval call binding the contract method 0x7897a78e.
//
// Solidity: function getDailySpendingRatio() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetDailySpendingRatio() (*big.Int, error) {
	return _Reserve.Contract.GetDailySpendingRatio(&_Reserve.CallOpts)
}

// GetExchangeSpenders is a free data retrieval call binding the contract method 0x1218f982.
//
// Solidity: function getExchangeSpenders() view returns(address[])
func (_Reserve *ReserveCaller) GetExchangeSpenders(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getExchangeSpenders")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetExchangeSpenders is a free data retrieval call binding the contract method 0x1218f982.
//
// Solidity: function getExchangeSpenders() view returns(address[])
func (_Reserve *ReserveSession) GetExchangeSpenders() ([]common.Address, error) {
	return _Reserve.Contract.GetExchangeSpenders(&_Reserve.CallOpts)
}

// GetExchangeSpenders is a free data retrieval call binding the contract method 0x1218f982.
//
// Solidity: function getExchangeSpenders() view returns(address[])
func (_Reserve *ReserveCallerSession) GetExchangeSpenders() ([]common.Address, error) {
	return _Reserve.Contract.GetExchangeSpenders(&_Reserve.CallOpts)
}

// GetFrozenReserveGoldBalance is a free data retrieval call binding the contract method 0x2aa1c16d.
//
// Solidity: function getFrozenReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveCaller) GetFrozenReserveGoldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getFrozenReserveGoldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFrozenReserveGoldBalance is a free data retrieval call binding the contract method 0x2aa1c16d.
//
// Solidity: function getFrozenReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveSession) GetFrozenReserveGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetFrozenReserveGoldBalance(&_Reserve.CallOpts)
}

// GetFrozenReserveGoldBalance is a free data retrieval call binding the contract method 0x2aa1c16d.
//
// Solidity: function getFrozenReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetFrozenReserveGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetFrozenReserveGoldBalance(&_Reserve.CallOpts)
}

// GetOtherReserveAddresses is a free data retrieval call binding the contract method 0x9c3e2f0f.
//
// Solidity: function getOtherReserveAddresses() view returns(address[])
func (_Reserve *ReserveCaller) GetOtherReserveAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getOtherReserveAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOtherReserveAddresses is a free data retrieval call binding the contract method 0x9c3e2f0f.
//
// Solidity: function getOtherReserveAddresses() view returns(address[])
func (_Reserve *ReserveSession) GetOtherReserveAddresses() ([]common.Address, error) {
	return _Reserve.Contract.GetOtherReserveAddresses(&_Reserve.CallOpts)
}

// GetOtherReserveAddresses is a free data retrieval call binding the contract method 0x9c3e2f0f.
//
// Solidity: function getOtherReserveAddresses() view returns(address[])
func (_Reserve *ReserveCallerSession) GetOtherReserveAddresses() ([]common.Address, error) {
	return _Reserve.Contract.GetOtherReserveAddresses(&_Reserve.CallOpts)
}

// GetOtherReserveAddressesGoldBalance is a free data retrieval call binding the contract method 0x765c1fe9.
//
// Solidity: function getOtherReserveAddressesGoldBalance() view returns(uint256)
func (_Reserve *ReserveCaller) GetOtherReserveAddressesGoldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getOtherReserveAddressesGoldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOtherReserveAddressesGoldBalance is a free data retrieval call binding the contract method 0x765c1fe9.
//
// Solidity: function getOtherReserveAddressesGoldBalance() view returns(uint256)
func (_Reserve *ReserveSession) GetOtherReserveAddressesGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetOtherReserveAddressesGoldBalance(&_Reserve.CallOpts)
}

// GetOtherReserveAddressesGoldBalance is a free data retrieval call binding the contract method 0x765c1fe9.
//
// Solidity: function getOtherReserveAddressesGoldBalance() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetOtherReserveAddressesGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetOtherReserveAddressesGoldBalance(&_Reserve.CallOpts)
}

// GetReserveGoldBalance is a free data retrieval call binding the contract method 0x8d9a5e6f.
//
// Solidity: function getReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveCaller) GetReserveGoldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getReserveGoldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveGoldBalance is a free data retrieval call binding the contract method 0x8d9a5e6f.
//
// Solidity: function getReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveSession) GetReserveGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetReserveGoldBalance(&_Reserve.CallOpts)
}

// GetReserveGoldBalance is a free data retrieval call binding the contract method 0x8d9a5e6f.
//
// Solidity: function getReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetReserveGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetReserveGoldBalance(&_Reserve.CallOpts)
}

// GetReserveRatio is a free data retrieval call binding the contract method 0x56b6d0d5.
//
// Solidity: function getReserveRatio() view returns(uint256)
func (_Reserve *ReserveCaller) GetReserveRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getReserveRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveRatio is a free data retrieval call binding the contract method 0x56b6d0d5.
//
// Solidity: function getReserveRatio() view returns(uint256)
func (_Reserve *ReserveSession) GetReserveRatio() (*big.Int, error) {
	return _Reserve.Contract.GetReserveRatio(&_Reserve.CallOpts)
}

// GetReserveRatio is a free data retrieval call binding the contract method 0x56b6d0d5.
//
// Solidity: function getReserveRatio() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetReserveRatio() (*big.Int, error) {
	return _Reserve.Contract.GetReserveRatio(&_Reserve.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Reserve *ReserveCaller) GetTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Reserve *ReserveSession) GetTokens() ([]common.Address, error) {
	return _Reserve.Contract.GetTokens(&_Reserve.CallOpts)
}

// GetTokens is a free data retrieval call binding the contract method 0xaa6ca808.
//
// Solidity: function getTokens() view returns(address[])
func (_Reserve *ReserveCallerSession) GetTokens() ([]common.Address, error) {
	return _Reserve.Contract.GetTokens(&_Reserve.CallOpts)
}

// GetUnfrozenBalance is a free data retrieval call binding the contract method 0xe30f579d.
//
// Solidity: function getUnfrozenBalance() view returns(uint256)
func (_Reserve *ReserveCaller) GetUnfrozenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getUnfrozenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnfrozenBalance is a free data retrieval call binding the contract method 0xe30f579d.
//
// Solidity: function getUnfrozenBalance() view returns(uint256)
func (_Reserve *ReserveSession) GetUnfrozenBalance() (*big.Int, error) {
	return _Reserve.Contract.GetUnfrozenBalance(&_Reserve.CallOpts)
}

// GetUnfrozenBalance is a free data retrieval call binding the contract method 0xe30f579d.
//
// Solidity: function getUnfrozenBalance() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetUnfrozenBalance() (*big.Int, error) {
	return _Reserve.Contract.GetUnfrozenBalance(&_Reserve.CallOpts)
}

// GetUnfrozenReserveGoldBalance is a free data retrieval call binding the contract method 0x8b7df8d4.
//
// Solidity: function getUnfrozenReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveCaller) GetUnfrozenReserveGoldBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getUnfrozenReserveGoldBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnfrozenReserveGoldBalance is a free data retrieval call binding the contract method 0x8b7df8d4.
//
// Solidity: function getUnfrozenReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveSession) GetUnfrozenReserveGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetUnfrozenReserveGoldBalance(&_Reserve.CallOpts)
}

// GetUnfrozenReserveGoldBalance is a free data retrieval call binding the contract method 0x8b7df8d4.
//
// Solidity: function getUnfrozenReserveGoldBalance() view returns(uint256)
func (_Reserve *ReserveCallerSession) GetUnfrozenReserveGoldBalance() (*big.Int, error) {
	return _Reserve.Contract.GetUnfrozenReserveGoldBalance(&_Reserve.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Reserve *ReserveCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "getVersionNumber")

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
func (_Reserve *ReserveSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Reserve.Contract.GetVersionNumber(&_Reserve.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Reserve *ReserveCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Reserve.Contract.GetVersionNumber(&_Reserve.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Reserve *ReserveCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Reserve *ReserveSession) Initialized() (bool, error) {
	return _Reserve.Contract.Initialized(&_Reserve.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Reserve *ReserveCallerSession) Initialized() (bool, error) {
	return _Reserve.Contract.Initialized(&_Reserve.CallOpts)
}

// IsExchangeSpender is a free data retrieval call binding the contract method 0x4cea8ded.
//
// Solidity: function isExchangeSpender(address ) view returns(bool)
func (_Reserve *ReserveCaller) IsExchangeSpender(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "isExchangeSpender", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExchangeSpender is a free data retrieval call binding the contract method 0x4cea8ded.
//
// Solidity: function isExchangeSpender(address ) view returns(bool)
func (_Reserve *ReserveSession) IsExchangeSpender(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsExchangeSpender(&_Reserve.CallOpts, arg0)
}

// IsExchangeSpender is a free data retrieval call binding the contract method 0x4cea8ded.
//
// Solidity: function isExchangeSpender(address ) view returns(bool)
func (_Reserve *ReserveCallerSession) IsExchangeSpender(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsExchangeSpender(&_Reserve.CallOpts, arg0)
}

// IsOtherReserveAddress is a free data retrieval call binding the contract method 0x7b522075.
//
// Solidity: function isOtherReserveAddress(address ) view returns(bool)
func (_Reserve *ReserveCaller) IsOtherReserveAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "isOtherReserveAddress", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOtherReserveAddress is a free data retrieval call binding the contract method 0x7b522075.
//
// Solidity: function isOtherReserveAddress(address ) view returns(bool)
func (_Reserve *ReserveSession) IsOtherReserveAddress(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsOtherReserveAddress(&_Reserve.CallOpts, arg0)
}

// IsOtherReserveAddress is a free data retrieval call binding the contract method 0x7b522075.
//
// Solidity: function isOtherReserveAddress(address ) view returns(bool)
func (_Reserve *ReserveCallerSession) IsOtherReserveAddress(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsOtherReserveAddress(&_Reserve.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Reserve *ReserveCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Reserve *ReserveSession) IsOwner() (bool, error) {
	return _Reserve.Contract.IsOwner(&_Reserve.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Reserve *ReserveCallerSession) IsOwner() (bool, error) {
	return _Reserve.Contract.IsOwner(&_Reserve.CallOpts)
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(address ) view returns(bool)
func (_Reserve *ReserveCaller) IsSpender(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "isSpender", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(address ) view returns(bool)
func (_Reserve *ReserveSession) IsSpender(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsSpender(&_Reserve.CallOpts, arg0)
}

// IsSpender is a free data retrieval call binding the contract method 0x9a206ece.
//
// Solidity: function isSpender(address ) view returns(bool)
func (_Reserve *ReserveCallerSession) IsSpender(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsSpender(&_Reserve.CallOpts, arg0)
}

// IsToken is a free data retrieval call binding the contract method 0x19f37361.
//
// Solidity: function isToken(address ) view returns(bool)
func (_Reserve *ReserveCaller) IsToken(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "isToken", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsToken is a free data retrieval call binding the contract method 0x19f37361.
//
// Solidity: function isToken(address ) view returns(bool)
func (_Reserve *ReserveSession) IsToken(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsToken(&_Reserve.CallOpts, arg0)
}

// IsToken is a free data retrieval call binding the contract method 0x19f37361.
//
// Solidity: function isToken(address ) view returns(bool)
func (_Reserve *ReserveCallerSession) IsToken(arg0 common.Address) (bool, error) {
	return _Reserve.Contract.IsToken(&_Reserve.CallOpts, arg0)
}

// LastSpendingDay is a free data retrieval call binding the contract method 0xfa9ed95a.
//
// Solidity: function lastSpendingDay() view returns(uint256)
func (_Reserve *ReserveCaller) LastSpendingDay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "lastSpendingDay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSpendingDay is a free data retrieval call binding the contract method 0xfa9ed95a.
//
// Solidity: function lastSpendingDay() view returns(uint256)
func (_Reserve *ReserveSession) LastSpendingDay() (*big.Int, error) {
	return _Reserve.Contract.LastSpendingDay(&_Reserve.CallOpts)
}

// LastSpendingDay is a free data retrieval call binding the contract method 0xfa9ed95a.
//
// Solidity: function lastSpendingDay() view returns(uint256)
func (_Reserve *ReserveCallerSession) LastSpendingDay() (*big.Int, error) {
	return _Reserve.Contract.LastSpendingDay(&_Reserve.CallOpts)
}

// OtherReserveAddresses is a free data retrieval call binding the contract method 0x40899365.
//
// Solidity: function otherReserveAddresses(uint256 ) view returns(address)
func (_Reserve *ReserveCaller) OtherReserveAddresses(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "otherReserveAddresses", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OtherReserveAddresses is a free data retrieval call binding the contract method 0x40899365.
//
// Solidity: function otherReserveAddresses(uint256 ) view returns(address)
func (_Reserve *ReserveSession) OtherReserveAddresses(arg0 *big.Int) (common.Address, error) {
	return _Reserve.Contract.OtherReserveAddresses(&_Reserve.CallOpts, arg0)
}

// OtherReserveAddresses is a free data retrieval call binding the contract method 0x40899365.
//
// Solidity: function otherReserveAddresses(uint256 ) view returns(address)
func (_Reserve *ReserveCallerSession) OtherReserveAddresses(arg0 *big.Int) (common.Address, error) {
	return _Reserve.Contract.OtherReserveAddresses(&_Reserve.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reserve *ReserveCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reserve *ReserveSession) Owner() (common.Address, error) {
	return _Reserve.Contract.Owner(&_Reserve.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Reserve *ReserveCallerSession) Owner() (common.Address, error) {
	return _Reserve.Contract.Owner(&_Reserve.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Reserve *ReserveCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Reserve *ReserveSession) Registry() (common.Address, error) {
	return _Reserve.Contract.Registry(&_Reserve.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Reserve *ReserveCallerSession) Registry() (common.Address, error) {
	return _Reserve.Contract.Registry(&_Reserve.CallOpts)
}

// SpendingLimit is a free data retrieval call binding the contract method 0x39d7f76e.
//
// Solidity: function spendingLimit() view returns(uint256)
func (_Reserve *ReserveCaller) SpendingLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "spendingLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SpendingLimit is a free data retrieval call binding the contract method 0x39d7f76e.
//
// Solidity: function spendingLimit() view returns(uint256)
func (_Reserve *ReserveSession) SpendingLimit() (*big.Int, error) {
	return _Reserve.Contract.SpendingLimit(&_Reserve.CallOpts)
}

// SpendingLimit is a free data retrieval call binding the contract method 0x39d7f76e.
//
// Solidity: function spendingLimit() view returns(uint256)
func (_Reserve *ReserveCallerSession) SpendingLimit() (*big.Int, error) {
	return _Reserve.Contract.SpendingLimit(&_Reserve.CallOpts)
}

// TobinTax is a free data retrieval call binding the contract method 0x894098d6.
//
// Solidity: function tobinTax() view returns(uint256)
func (_Reserve *ReserveCaller) TobinTax(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "tobinTax")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TobinTax is a free data retrieval call binding the contract method 0x894098d6.
//
// Solidity: function tobinTax() view returns(uint256)
func (_Reserve *ReserveSession) TobinTax() (*big.Int, error) {
	return _Reserve.Contract.TobinTax(&_Reserve.CallOpts)
}

// TobinTax is a free data retrieval call binding the contract method 0x894098d6.
//
// Solidity: function tobinTax() view returns(uint256)
func (_Reserve *ReserveCallerSession) TobinTax() (*big.Int, error) {
	return _Reserve.Contract.TobinTax(&_Reserve.CallOpts)
}

// TobinTaxCache is a free data retrieval call binding the contract method 0x22796e83.
//
// Solidity: function tobinTaxCache() view returns(uint128 numerator, uint128 timestamp)
func (_Reserve *ReserveCaller) TobinTaxCache(opts *bind.CallOpts) (struct {
	Numerator *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "tobinTaxCache")

	outstruct := new(struct {
		Numerator *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Numerator = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TobinTaxCache is a free data retrieval call binding the contract method 0x22796e83.
//
// Solidity: function tobinTaxCache() view returns(uint128 numerator, uint128 timestamp)
func (_Reserve *ReserveSession) TobinTaxCache() (struct {
	Numerator *big.Int
	Timestamp *big.Int
}, error) {
	return _Reserve.Contract.TobinTaxCache(&_Reserve.CallOpts)
}

// TobinTaxCache is a free data retrieval call binding the contract method 0x22796e83.
//
// Solidity: function tobinTaxCache() view returns(uint128 numerator, uint128 timestamp)
func (_Reserve *ReserveCallerSession) TobinTaxCache() (struct {
	Numerator *big.Int
	Timestamp *big.Int
}, error) {
	return _Reserve.Contract.TobinTaxCache(&_Reserve.CallOpts)
}

// TobinTaxReserveRatio is a free data retrieval call binding the contract method 0x76769a60.
//
// Solidity: function tobinTaxReserveRatio() view returns(uint256)
func (_Reserve *ReserveCaller) TobinTaxReserveRatio(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "tobinTaxReserveRatio")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TobinTaxReserveRatio is a free data retrieval call binding the contract method 0x76769a60.
//
// Solidity: function tobinTaxReserveRatio() view returns(uint256)
func (_Reserve *ReserveSession) TobinTaxReserveRatio() (*big.Int, error) {
	return _Reserve.Contract.TobinTaxReserveRatio(&_Reserve.CallOpts)
}

// TobinTaxReserveRatio is a free data retrieval call binding the contract method 0x76769a60.
//
// Solidity: function tobinTaxReserveRatio() view returns(uint256)
func (_Reserve *ReserveCallerSession) TobinTaxReserveRatio() (*big.Int, error) {
	return _Reserve.Contract.TobinTaxReserveRatio(&_Reserve.CallOpts)
}

// TobinTaxStalenessThreshold is a free data retrieval call binding the contract method 0xe33a88e7.
//
// Solidity: function tobinTaxStalenessThreshold() view returns(uint256)
func (_Reserve *ReserveCaller) TobinTaxStalenessThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Reserve.contract.Call(opts, &out, "tobinTaxStalenessThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TobinTaxStalenessThreshold is a free data retrieval call binding the contract method 0xe33a88e7.
//
// Solidity: function tobinTaxStalenessThreshold() view returns(uint256)
func (_Reserve *ReserveSession) TobinTaxStalenessThreshold() (*big.Int, error) {
	return _Reserve.Contract.TobinTaxStalenessThreshold(&_Reserve.CallOpts)
}

// TobinTaxStalenessThreshold is a free data retrieval call binding the contract method 0xe33a88e7.
//
// Solidity: function tobinTaxStalenessThreshold() view returns(uint256)
func (_Reserve *ReserveCallerSession) TobinTaxStalenessThreshold() (*big.Int, error) {
	return _Reserve.Contract.TobinTaxStalenessThreshold(&_Reserve.CallOpts)
}

// AddExchangeSpender is a paid mutator transaction binding the contract method 0xf0b7182b.
//
// Solidity: function addExchangeSpender(address spender) returns()
func (_Reserve *ReserveTransactor) AddExchangeSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "addExchangeSpender", spender)
}

// AddExchangeSpender is a paid mutator transaction binding the contract method 0xf0b7182b.
//
// Solidity: function addExchangeSpender(address spender) returns()
func (_Reserve *ReserveSession) AddExchangeSpender(spender common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddExchangeSpender(&_Reserve.TransactOpts, spender)
}

// AddExchangeSpender is a paid mutator transaction binding the contract method 0xf0b7182b.
//
// Solidity: function addExchangeSpender(address spender) returns()
func (_Reserve *ReserveTransactorSession) AddExchangeSpender(spender common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddExchangeSpender(&_Reserve.TransactOpts, spender)
}

// AddOtherReserveAddress is a paid mutator transaction binding the contract method 0x22015968.
//
// Solidity: function addOtherReserveAddress(address reserveAddress) returns(bool)
func (_Reserve *ReserveTransactor) AddOtherReserveAddress(opts *bind.TransactOpts, reserveAddress common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "addOtherReserveAddress", reserveAddress)
}

// AddOtherReserveAddress is a paid mutator transaction binding the contract method 0x22015968.
//
// Solidity: function addOtherReserveAddress(address reserveAddress) returns(bool)
func (_Reserve *ReserveSession) AddOtherReserveAddress(reserveAddress common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddOtherReserveAddress(&_Reserve.TransactOpts, reserveAddress)
}

// AddOtherReserveAddress is a paid mutator transaction binding the contract method 0x22015968.
//
// Solidity: function addOtherReserveAddress(address reserveAddress) returns(bool)
func (_Reserve *ReserveTransactorSession) AddOtherReserveAddress(reserveAddress common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddOtherReserveAddress(&_Reserve.TransactOpts, reserveAddress)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(address spender) returns()
func (_Reserve *ReserveTransactor) AddSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "addSpender", spender)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(address spender) returns()
func (_Reserve *ReserveSession) AddSpender(spender common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddSpender(&_Reserve.TransactOpts, spender)
}

// AddSpender is a paid mutator transaction binding the contract method 0xe7e31e7a.
//
// Solidity: function addSpender(address spender) returns()
func (_Reserve *ReserveTransactorSession) AddSpender(spender common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddSpender(&_Reserve.TransactOpts, spender)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns(bool)
func (_Reserve *ReserveTransactor) AddToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "addToken", token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns(bool)
func (_Reserve *ReserveSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddToken(&_Reserve.TransactOpts, token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address token) returns(bool)
func (_Reserve *ReserveTransactorSession) AddToken(token common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.AddToken(&_Reserve.TransactOpts, token)
}

// GetOrComputeTobinTax is a paid mutator transaction binding the contract method 0x17f9a6f7.
//
// Solidity: function getOrComputeTobinTax() returns(uint256, uint256)
func (_Reserve *ReserveTransactor) GetOrComputeTobinTax(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "getOrComputeTobinTax")
}

// GetOrComputeTobinTax is a paid mutator transaction binding the contract method 0x17f9a6f7.
//
// Solidity: function getOrComputeTobinTax() returns(uint256, uint256)
func (_Reserve *ReserveSession) GetOrComputeTobinTax() (*types.Transaction, error) {
	return _Reserve.Contract.GetOrComputeTobinTax(&_Reserve.TransactOpts)
}

// GetOrComputeTobinTax is a paid mutator transaction binding the contract method 0x17f9a6f7.
//
// Solidity: function getOrComputeTobinTax() returns(uint256, uint256)
func (_Reserve *ReserveTransactorSession) GetOrComputeTobinTax() (*types.Transaction, error) {
	return _Reserve.Contract.GetOrComputeTobinTax(&_Reserve.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa3e1f00d.
//
// Solidity: function initialize(address registryAddress, uint256 _tobinTaxStalenessThreshold, uint256 _spendingRatio, uint256 _frozenGold, uint256 _frozenDays, bytes32[] _assetAllocationSymbols, uint256[] _assetAllocationWeights, uint256 _tobinTax, uint256 _tobinTaxReserveRatio) returns()
func (_Reserve *ReserveTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _tobinTaxStalenessThreshold *big.Int, _spendingRatio *big.Int, _frozenGold *big.Int, _frozenDays *big.Int, _assetAllocationSymbols [][32]byte, _assetAllocationWeights []*big.Int, _tobinTax *big.Int, _tobinTaxReserveRatio *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "initialize", registryAddress, _tobinTaxStalenessThreshold, _spendingRatio, _frozenGold, _frozenDays, _assetAllocationSymbols, _assetAllocationWeights, _tobinTax, _tobinTaxReserveRatio)
}

// Initialize is a paid mutator transaction binding the contract method 0xa3e1f00d.
//
// Solidity: function initialize(address registryAddress, uint256 _tobinTaxStalenessThreshold, uint256 _spendingRatio, uint256 _frozenGold, uint256 _frozenDays, bytes32[] _assetAllocationSymbols, uint256[] _assetAllocationWeights, uint256 _tobinTax, uint256 _tobinTaxReserveRatio) returns()
func (_Reserve *ReserveSession) Initialize(registryAddress common.Address, _tobinTaxStalenessThreshold *big.Int, _spendingRatio *big.Int, _frozenGold *big.Int, _frozenDays *big.Int, _assetAllocationSymbols [][32]byte, _assetAllocationWeights []*big.Int, _tobinTax *big.Int, _tobinTaxReserveRatio *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Initialize(&_Reserve.TransactOpts, registryAddress, _tobinTaxStalenessThreshold, _spendingRatio, _frozenGold, _frozenDays, _assetAllocationSymbols, _assetAllocationWeights, _tobinTax, _tobinTaxReserveRatio)
}

// Initialize is a paid mutator transaction binding the contract method 0xa3e1f00d.
//
// Solidity: function initialize(address registryAddress, uint256 _tobinTaxStalenessThreshold, uint256 _spendingRatio, uint256 _frozenGold, uint256 _frozenDays, bytes32[] _assetAllocationSymbols, uint256[] _assetAllocationWeights, uint256 _tobinTax, uint256 _tobinTaxReserveRatio) returns()
func (_Reserve *ReserveTransactorSession) Initialize(registryAddress common.Address, _tobinTaxStalenessThreshold *big.Int, _spendingRatio *big.Int, _frozenGold *big.Int, _frozenDays *big.Int, _assetAllocationSymbols [][32]byte, _assetAllocationWeights []*big.Int, _tobinTax *big.Int, _tobinTaxReserveRatio *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.Initialize(&_Reserve.TransactOpts, registryAddress, _tobinTaxStalenessThreshold, _spendingRatio, _frozenGold, _frozenDays, _assetAllocationSymbols, _assetAllocationWeights, _tobinTax, _tobinTaxReserveRatio)
}

// RemoveExchangeSpender is a paid mutator transaction binding the contract method 0xb003dcf1.
//
// Solidity: function removeExchangeSpender(address spender, uint256 index) returns()
func (_Reserve *ReserveTransactor) RemoveExchangeSpender(opts *bind.TransactOpts, spender common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "removeExchangeSpender", spender, index)
}

// RemoveExchangeSpender is a paid mutator transaction binding the contract method 0xb003dcf1.
//
// Solidity: function removeExchangeSpender(address spender, uint256 index) returns()
func (_Reserve *ReserveSession) RemoveExchangeSpender(spender common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveExchangeSpender(&_Reserve.TransactOpts, spender, index)
}

// RemoveExchangeSpender is a paid mutator transaction binding the contract method 0xb003dcf1.
//
// Solidity: function removeExchangeSpender(address spender, uint256 index) returns()
func (_Reserve *ReserveTransactorSession) RemoveExchangeSpender(spender common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveExchangeSpender(&_Reserve.TransactOpts, spender, index)
}

// RemoveOtherReserveAddress is a paid mutator transaction binding the contract method 0x5c4a3145.
//
// Solidity: function removeOtherReserveAddress(address reserveAddress, uint256 index) returns(bool)
func (_Reserve *ReserveTransactor) RemoveOtherReserveAddress(opts *bind.TransactOpts, reserveAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "removeOtherReserveAddress", reserveAddress, index)
}

// RemoveOtherReserveAddress is a paid mutator transaction binding the contract method 0x5c4a3145.
//
// Solidity: function removeOtherReserveAddress(address reserveAddress, uint256 index) returns(bool)
func (_Reserve *ReserveSession) RemoveOtherReserveAddress(reserveAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveOtherReserveAddress(&_Reserve.TransactOpts, reserveAddress, index)
}

// RemoveOtherReserveAddress is a paid mutator transaction binding the contract method 0x5c4a3145.
//
// Solidity: function removeOtherReserveAddress(address reserveAddress, uint256 index) returns(bool)
func (_Reserve *ReserveTransactorSession) RemoveOtherReserveAddress(reserveAddress common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveOtherReserveAddress(&_Reserve.TransactOpts, reserveAddress, index)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(address spender) returns()
func (_Reserve *ReserveTransactor) RemoveSpender(opts *bind.TransactOpts, spender common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "removeSpender", spender)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(address spender) returns()
func (_Reserve *ReserveSession) RemoveSpender(spender common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveSpender(&_Reserve.TransactOpts, spender)
}

// RemoveSpender is a paid mutator transaction binding the contract method 0x8ce5877c.
//
// Solidity: function removeSpender(address spender) returns()
func (_Reserve *ReserveTransactorSession) RemoveSpender(spender common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveSpender(&_Reserve.TransactOpts, spender)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(address token, uint256 index) returns(bool)
func (_Reserve *ReserveTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "removeToken", token, index)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(address token, uint256 index) returns(bool)
func (_Reserve *ReserveSession) RemoveToken(token common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveToken(&_Reserve.TransactOpts, token, index)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x13baf1e6.
//
// Solidity: function removeToken(address token, uint256 index) returns(bool)
func (_Reserve *ReserveTransactorSession) RemoveToken(token common.Address, index *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.RemoveToken(&_Reserve.TransactOpts, token, index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reserve *ReserveTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reserve *ReserveSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reserve.Contract.RenounceOwnership(&_Reserve.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Reserve *ReserveTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Reserve.Contract.RenounceOwnership(&_Reserve.TransactOpts)
}

// SetAssetAllocations is a paid mutator transaction binding the contract method 0xca56d33b.
//
// Solidity: function setAssetAllocations(bytes32[] symbols, uint256[] weights) returns()
func (_Reserve *ReserveTransactor) SetAssetAllocations(opts *bind.TransactOpts, symbols [][32]byte, weights []*big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setAssetAllocations", symbols, weights)
}

// SetAssetAllocations is a paid mutator transaction binding the contract method 0xca56d33b.
//
// Solidity: function setAssetAllocations(bytes32[] symbols, uint256[] weights) returns()
func (_Reserve *ReserveSession) SetAssetAllocations(symbols [][32]byte, weights []*big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetAssetAllocations(&_Reserve.TransactOpts, symbols, weights)
}

// SetAssetAllocations is a paid mutator transaction binding the contract method 0xca56d33b.
//
// Solidity: function setAssetAllocations(bytes32[] symbols, uint256[] weights) returns()
func (_Reserve *ReserveTransactorSession) SetAssetAllocations(symbols [][32]byte, weights []*big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetAssetAllocations(&_Reserve.TransactOpts, symbols, weights)
}

// SetDailySpendingRatio is a paid mutator transaction binding the contract method 0x01da32bd.
//
// Solidity: function setDailySpendingRatio(uint256 ratio) returns()
func (_Reserve *ReserveTransactor) SetDailySpendingRatio(opts *bind.TransactOpts, ratio *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setDailySpendingRatio", ratio)
}

// SetDailySpendingRatio is a paid mutator transaction binding the contract method 0x01da32bd.
//
// Solidity: function setDailySpendingRatio(uint256 ratio) returns()
func (_Reserve *ReserveSession) SetDailySpendingRatio(ratio *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetDailySpendingRatio(&_Reserve.TransactOpts, ratio)
}

// SetDailySpendingRatio is a paid mutator transaction binding the contract method 0x01da32bd.
//
// Solidity: function setDailySpendingRatio(uint256 ratio) returns()
func (_Reserve *ReserveTransactorSession) SetDailySpendingRatio(ratio *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetDailySpendingRatio(&_Reserve.TransactOpts, ratio)
}

// SetFrozenGold is a paid mutator transaction binding the contract method 0xe83b373b.
//
// Solidity: function setFrozenGold(uint256 frozenGold, uint256 frozenDays) returns()
func (_Reserve *ReserveTransactor) SetFrozenGold(opts *bind.TransactOpts, frozenGold *big.Int, frozenDays *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setFrozenGold", frozenGold, frozenDays)
}

// SetFrozenGold is a paid mutator transaction binding the contract method 0xe83b373b.
//
// Solidity: function setFrozenGold(uint256 frozenGold, uint256 frozenDays) returns()
func (_Reserve *ReserveSession) SetFrozenGold(frozenGold *big.Int, frozenDays *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetFrozenGold(&_Reserve.TransactOpts, frozenGold, frozenDays)
}

// SetFrozenGold is a paid mutator transaction binding the contract method 0xe83b373b.
//
// Solidity: function setFrozenGold(uint256 frozenGold, uint256 frozenDays) returns()
func (_Reserve *ReserveTransactorSession) SetFrozenGold(frozenGold *big.Int, frozenDays *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetFrozenGold(&_Reserve.TransactOpts, frozenGold, frozenDays)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Reserve *ReserveTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Reserve *ReserveSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SetRegistry(&_Reserve.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Reserve *ReserveTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.SetRegistry(&_Reserve.TransactOpts, registryAddress)
}

// SetTobinTax is a paid mutator transaction binding the contract method 0xe6b76e9c.
//
// Solidity: function setTobinTax(uint256 value) returns()
func (_Reserve *ReserveTransactor) SetTobinTax(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setTobinTax", value)
}

// SetTobinTax is a paid mutator transaction binding the contract method 0xe6b76e9c.
//
// Solidity: function setTobinTax(uint256 value) returns()
func (_Reserve *ReserveSession) SetTobinTax(value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetTobinTax(&_Reserve.TransactOpts, value)
}

// SetTobinTax is a paid mutator transaction binding the contract method 0xe6b76e9c.
//
// Solidity: function setTobinTax(uint256 value) returns()
func (_Reserve *ReserveTransactorSession) SetTobinTax(value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetTobinTax(&_Reserve.TransactOpts, value)
}

// SetTobinTaxReserveRatio is a paid mutator transaction binding the contract method 0xad62ad10.
//
// Solidity: function setTobinTaxReserveRatio(uint256 value) returns()
func (_Reserve *ReserveTransactor) SetTobinTaxReserveRatio(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setTobinTaxReserveRatio", value)
}

// SetTobinTaxReserveRatio is a paid mutator transaction binding the contract method 0xad62ad10.
//
// Solidity: function setTobinTaxReserveRatio(uint256 value) returns()
func (_Reserve *ReserveSession) SetTobinTaxReserveRatio(value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetTobinTaxReserveRatio(&_Reserve.TransactOpts, value)
}

// SetTobinTaxReserveRatio is a paid mutator transaction binding the contract method 0xad62ad10.
//
// Solidity: function setTobinTaxReserveRatio(uint256 value) returns()
func (_Reserve *ReserveTransactorSession) SetTobinTaxReserveRatio(value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetTobinTaxReserveRatio(&_Reserve.TransactOpts, value)
}

// SetTobinTaxStalenessThreshold is a paid mutator transaction binding the contract method 0xa1ab55b3.
//
// Solidity: function setTobinTaxStalenessThreshold(uint256 value) returns()
func (_Reserve *ReserveTransactor) SetTobinTaxStalenessThreshold(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "setTobinTaxStalenessThreshold", value)
}

// SetTobinTaxStalenessThreshold is a paid mutator transaction binding the contract method 0xa1ab55b3.
//
// Solidity: function setTobinTaxStalenessThreshold(uint256 value) returns()
func (_Reserve *ReserveSession) SetTobinTaxStalenessThreshold(value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetTobinTaxStalenessThreshold(&_Reserve.TransactOpts, value)
}

// SetTobinTaxStalenessThreshold is a paid mutator transaction binding the contract method 0xa1ab55b3.
//
// Solidity: function setTobinTaxStalenessThreshold(uint256 value) returns()
func (_Reserve *ReserveTransactorSession) SetTobinTaxStalenessThreshold(value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.SetTobinTaxStalenessThreshold(&_Reserve.TransactOpts, value)
}

// TransferExchangeGold is a paid mutator transaction binding the contract method 0x03a0fea3.
//
// Solidity: function transferExchangeGold(address to, uint256 value) returns(bool)
func (_Reserve *ReserveTransactor) TransferExchangeGold(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "transferExchangeGold", to, value)
}

// TransferExchangeGold is a paid mutator transaction binding the contract method 0x03a0fea3.
//
// Solidity: function transferExchangeGold(address to, uint256 value) returns(bool)
func (_Reserve *ReserveSession) TransferExchangeGold(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.TransferExchangeGold(&_Reserve.TransactOpts, to, value)
}

// TransferExchangeGold is a paid mutator transaction binding the contract method 0x03a0fea3.
//
// Solidity: function transferExchangeGold(address to, uint256 value) returns(bool)
func (_Reserve *ReserveTransactorSession) TransferExchangeGold(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.TransferExchangeGold(&_Reserve.TransactOpts, to, value)
}

// TransferGold is a paid mutator transaction binding the contract method 0x1c39c7d5.
//
// Solidity: function transferGold(address to, uint256 value) returns(bool)
func (_Reserve *ReserveTransactor) TransferGold(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "transferGold", to, value)
}

// TransferGold is a paid mutator transaction binding the contract method 0x1c39c7d5.
//
// Solidity: function transferGold(address to, uint256 value) returns(bool)
func (_Reserve *ReserveSession) TransferGold(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.TransferGold(&_Reserve.TransactOpts, to, value)
}

// TransferGold is a paid mutator transaction binding the contract method 0x1c39c7d5.
//
// Solidity: function transferGold(address to, uint256 value) returns(bool)
func (_Reserve *ReserveTransactorSession) TransferGold(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Reserve.Contract.TransferGold(&_Reserve.TransactOpts, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reserve *ReserveTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Reserve.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reserve *ReserveSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.TransferOwnership(&_Reserve.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Reserve *ReserveTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Reserve.Contract.TransferOwnership(&_Reserve.TransactOpts, newOwner)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Reserve *ReserveFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Reserve.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AssetAllocationSet":
		event, err = _Reserve.ParseAssetAllocationSet(log)
	case "DailySpendingRatioSet":
		event, err = _Reserve.ParseDailySpendingRatioSet(log)
	case "ExchangeSpenderAdded":
		event, err = _Reserve.ParseExchangeSpenderAdded(log)
	case "ExchangeSpenderRemoved":
		event, err = _Reserve.ParseExchangeSpenderRemoved(log)
	case "OtherReserveAddressAdded":
		event, err = _Reserve.ParseOtherReserveAddressAdded(log)
	case "OtherReserveAddressRemoved":
		event, err = _Reserve.ParseOtherReserveAddressRemoved(log)
	case "OwnershipTransferred":
		event, err = _Reserve.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Reserve.ParseRegistrySet(log)
	case "ReserveGoldTransferred":
		event, err = _Reserve.ParseReserveGoldTransferred(log)
	case "SpenderAdded":
		event, err = _Reserve.ParseSpenderAdded(log)
	case "SpenderRemoved":
		event, err = _Reserve.ParseSpenderRemoved(log)
	case "TobinTaxReserveRatioSet":
		event, err = _Reserve.ParseTobinTaxReserveRatioSet(log)
	case "TobinTaxSet":
		event, err = _Reserve.ParseTobinTaxSet(log)
	case "TobinTaxStalenessThresholdSet":
		event, err = _Reserve.ParseTobinTaxStalenessThresholdSet(log)
	case "TokenAdded":
		event, err = _Reserve.ParseTokenAdded(log)
	case "TokenRemoved":
		event, err = _Reserve.ParseTokenRemoved(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Reserve *ReserveTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Reserve.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Reserve *ReserveSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Reserve.Contract.Fallback(&_Reserve.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Reserve *ReserveTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Reserve.Contract.Fallback(&_Reserve.TransactOpts, calldata)
}

// ReserveAssetAllocationSetIterator is returned from FilterAssetAllocationSet and is used to iterate over the raw logs and unpacked data for AssetAllocationSet events raised by the Reserve contract.
type ReserveAssetAllocationSetIterator struct {
	Event *ReserveAssetAllocationSet // Event containing the contract specifics and raw log

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
func (it *ReserveAssetAllocationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveAssetAllocationSet)
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
		it.Event = new(ReserveAssetAllocationSet)
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
func (it *ReserveAssetAllocationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveAssetAllocationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveAssetAllocationSet represents a AssetAllocationSet event raised by the Reserve contract.
type ReserveAssetAllocationSet struct {
	Symbols [][32]byte
	Weights []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAssetAllocationSet is a free log retrieval operation binding the contract event 0x55b488abd19ae7621712324d3d42c2ef7a9575f64f5503103286a1161fb40855.
//
// Solidity: event AssetAllocationSet(bytes32[] symbols, uint256[] weights)
func (_Reserve *ReserveFilterer) FilterAssetAllocationSet(opts *bind.FilterOpts) (*ReserveAssetAllocationSetIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "AssetAllocationSet")
	if err != nil {
		return nil, err
	}
	return &ReserveAssetAllocationSetIterator{contract: _Reserve.contract, event: "AssetAllocationSet", logs: logs, sub: sub}, nil
}

// WatchAssetAllocationSet is a free log subscription operation binding the contract event 0x55b488abd19ae7621712324d3d42c2ef7a9575f64f5503103286a1161fb40855.
//
// Solidity: event AssetAllocationSet(bytes32[] symbols, uint256[] weights)
func (_Reserve *ReserveFilterer) WatchAssetAllocationSet(opts *bind.WatchOpts, sink chan<- *ReserveAssetAllocationSet) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "AssetAllocationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveAssetAllocationSet)
				if err := _Reserve.contract.UnpackLog(event, "AssetAllocationSet", log); err != nil {
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

// ParseAssetAllocationSet is a log parse operation binding the contract event 0x55b488abd19ae7621712324d3d42c2ef7a9575f64f5503103286a1161fb40855.
//
// Solidity: event AssetAllocationSet(bytes32[] symbols, uint256[] weights)
func (_Reserve *ReserveFilterer) ParseAssetAllocationSet(log types.Log) (*ReserveAssetAllocationSet, error) {
	event := new(ReserveAssetAllocationSet)
	if err := _Reserve.contract.UnpackLog(event, "AssetAllocationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveDailySpendingRatioSetIterator is returned from FilterDailySpendingRatioSet and is used to iterate over the raw logs and unpacked data for DailySpendingRatioSet events raised by the Reserve contract.
type ReserveDailySpendingRatioSetIterator struct {
	Event *ReserveDailySpendingRatioSet // Event containing the contract specifics and raw log

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
func (it *ReserveDailySpendingRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveDailySpendingRatioSet)
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
		it.Event = new(ReserveDailySpendingRatioSet)
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
func (it *ReserveDailySpendingRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveDailySpendingRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveDailySpendingRatioSet represents a DailySpendingRatioSet event raised by the Reserve contract.
type ReserveDailySpendingRatioSet struct {
	Ratio *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterDailySpendingRatioSet is a free log retrieval operation binding the contract event 0xb08f0607338ad77f5b08ccf831e533cefcc2d373c173e87a8f61144f1d82be1e.
//
// Solidity: event DailySpendingRatioSet(uint256 ratio)
func (_Reserve *ReserveFilterer) FilterDailySpendingRatioSet(opts *bind.FilterOpts) (*ReserveDailySpendingRatioSetIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "DailySpendingRatioSet")
	if err != nil {
		return nil, err
	}
	return &ReserveDailySpendingRatioSetIterator{contract: _Reserve.contract, event: "DailySpendingRatioSet", logs: logs, sub: sub}, nil
}

// WatchDailySpendingRatioSet is a free log subscription operation binding the contract event 0xb08f0607338ad77f5b08ccf831e533cefcc2d373c173e87a8f61144f1d82be1e.
//
// Solidity: event DailySpendingRatioSet(uint256 ratio)
func (_Reserve *ReserveFilterer) WatchDailySpendingRatioSet(opts *bind.WatchOpts, sink chan<- *ReserveDailySpendingRatioSet) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "DailySpendingRatioSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveDailySpendingRatioSet)
				if err := _Reserve.contract.UnpackLog(event, "DailySpendingRatioSet", log); err != nil {
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

// ParseDailySpendingRatioSet is a log parse operation binding the contract event 0xb08f0607338ad77f5b08ccf831e533cefcc2d373c173e87a8f61144f1d82be1e.
//
// Solidity: event DailySpendingRatioSet(uint256 ratio)
func (_Reserve *ReserveFilterer) ParseDailySpendingRatioSet(log types.Log) (*ReserveDailySpendingRatioSet, error) {
	event := new(ReserveDailySpendingRatioSet)
	if err := _Reserve.contract.UnpackLog(event, "DailySpendingRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveExchangeSpenderAddedIterator is returned from FilterExchangeSpenderAdded and is used to iterate over the raw logs and unpacked data for ExchangeSpenderAdded events raised by the Reserve contract.
type ReserveExchangeSpenderAddedIterator struct {
	Event *ReserveExchangeSpenderAdded // Event containing the contract specifics and raw log

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
func (it *ReserveExchangeSpenderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveExchangeSpenderAdded)
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
		it.Event = new(ReserveExchangeSpenderAdded)
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
func (it *ReserveExchangeSpenderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveExchangeSpenderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveExchangeSpenderAdded represents a ExchangeSpenderAdded event raised by the Reserve contract.
type ReserveExchangeSpenderAdded struct {
	ExchangeSpender common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExchangeSpenderAdded is a free log retrieval operation binding the contract event 0x71bccdb89fff4d914e3d2e472b327e3debaf4c4d6f1dfe528f430447e4cbcf5f.
//
// Solidity: event ExchangeSpenderAdded(address indexed exchangeSpender)
func (_Reserve *ReserveFilterer) FilterExchangeSpenderAdded(opts *bind.FilterOpts, exchangeSpender []common.Address) (*ReserveExchangeSpenderAddedIterator, error) {

	var exchangeSpenderRule []interface{}
	for _, exchangeSpenderItem := range exchangeSpender {
		exchangeSpenderRule = append(exchangeSpenderRule, exchangeSpenderItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "ExchangeSpenderAdded", exchangeSpenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveExchangeSpenderAddedIterator{contract: _Reserve.contract, event: "ExchangeSpenderAdded", logs: logs, sub: sub}, nil
}

// WatchExchangeSpenderAdded is a free log subscription operation binding the contract event 0x71bccdb89fff4d914e3d2e472b327e3debaf4c4d6f1dfe528f430447e4cbcf5f.
//
// Solidity: event ExchangeSpenderAdded(address indexed exchangeSpender)
func (_Reserve *ReserveFilterer) WatchExchangeSpenderAdded(opts *bind.WatchOpts, sink chan<- *ReserveExchangeSpenderAdded, exchangeSpender []common.Address) (event.Subscription, error) {

	var exchangeSpenderRule []interface{}
	for _, exchangeSpenderItem := range exchangeSpender {
		exchangeSpenderRule = append(exchangeSpenderRule, exchangeSpenderItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "ExchangeSpenderAdded", exchangeSpenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveExchangeSpenderAdded)
				if err := _Reserve.contract.UnpackLog(event, "ExchangeSpenderAdded", log); err != nil {
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

// ParseExchangeSpenderAdded is a log parse operation binding the contract event 0x71bccdb89fff4d914e3d2e472b327e3debaf4c4d6f1dfe528f430447e4cbcf5f.
//
// Solidity: event ExchangeSpenderAdded(address indexed exchangeSpender)
func (_Reserve *ReserveFilterer) ParseExchangeSpenderAdded(log types.Log) (*ReserveExchangeSpenderAdded, error) {
	event := new(ReserveExchangeSpenderAdded)
	if err := _Reserve.contract.UnpackLog(event, "ExchangeSpenderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveExchangeSpenderRemovedIterator is returned from FilterExchangeSpenderRemoved and is used to iterate over the raw logs and unpacked data for ExchangeSpenderRemoved events raised by the Reserve contract.
type ReserveExchangeSpenderRemovedIterator struct {
	Event *ReserveExchangeSpenderRemoved // Event containing the contract specifics and raw log

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
func (it *ReserveExchangeSpenderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveExchangeSpenderRemoved)
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
		it.Event = new(ReserveExchangeSpenderRemoved)
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
func (it *ReserveExchangeSpenderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveExchangeSpenderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveExchangeSpenderRemoved represents a ExchangeSpenderRemoved event raised by the Reserve contract.
type ReserveExchangeSpenderRemoved struct {
	ExchangeSpender common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExchangeSpenderRemoved is a free log retrieval operation binding the contract event 0x20aaa18caa668680a42b328a15fd50d580bac65d8bd346e104355473c6373ff3.
//
// Solidity: event ExchangeSpenderRemoved(address indexed exchangeSpender)
func (_Reserve *ReserveFilterer) FilterExchangeSpenderRemoved(opts *bind.FilterOpts, exchangeSpender []common.Address) (*ReserveExchangeSpenderRemovedIterator, error) {

	var exchangeSpenderRule []interface{}
	for _, exchangeSpenderItem := range exchangeSpender {
		exchangeSpenderRule = append(exchangeSpenderRule, exchangeSpenderItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "ExchangeSpenderRemoved", exchangeSpenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveExchangeSpenderRemovedIterator{contract: _Reserve.contract, event: "ExchangeSpenderRemoved", logs: logs, sub: sub}, nil
}

// WatchExchangeSpenderRemoved is a free log subscription operation binding the contract event 0x20aaa18caa668680a42b328a15fd50d580bac65d8bd346e104355473c6373ff3.
//
// Solidity: event ExchangeSpenderRemoved(address indexed exchangeSpender)
func (_Reserve *ReserveFilterer) WatchExchangeSpenderRemoved(opts *bind.WatchOpts, sink chan<- *ReserveExchangeSpenderRemoved, exchangeSpender []common.Address) (event.Subscription, error) {

	var exchangeSpenderRule []interface{}
	for _, exchangeSpenderItem := range exchangeSpender {
		exchangeSpenderRule = append(exchangeSpenderRule, exchangeSpenderItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "ExchangeSpenderRemoved", exchangeSpenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveExchangeSpenderRemoved)
				if err := _Reserve.contract.UnpackLog(event, "ExchangeSpenderRemoved", log); err != nil {
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

// ParseExchangeSpenderRemoved is a log parse operation binding the contract event 0x20aaa18caa668680a42b328a15fd50d580bac65d8bd346e104355473c6373ff3.
//
// Solidity: event ExchangeSpenderRemoved(address indexed exchangeSpender)
func (_Reserve *ReserveFilterer) ParseExchangeSpenderRemoved(log types.Log) (*ReserveExchangeSpenderRemoved, error) {
	event := new(ReserveExchangeSpenderRemoved)
	if err := _Reserve.contract.UnpackLog(event, "ExchangeSpenderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveOtherReserveAddressAddedIterator is returned from FilterOtherReserveAddressAdded and is used to iterate over the raw logs and unpacked data for OtherReserveAddressAdded events raised by the Reserve contract.
type ReserveOtherReserveAddressAddedIterator struct {
	Event *ReserveOtherReserveAddressAdded // Event containing the contract specifics and raw log

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
func (it *ReserveOtherReserveAddressAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveOtherReserveAddressAdded)
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
		it.Event = new(ReserveOtherReserveAddressAdded)
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
func (it *ReserveOtherReserveAddressAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveOtherReserveAddressAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveOtherReserveAddressAdded represents a OtherReserveAddressAdded event raised by the Reserve contract.
type ReserveOtherReserveAddressAdded struct {
	OtherReserveAddress common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOtherReserveAddressAdded is a free log retrieval operation binding the contract event 0xd78793225285ecf9cf5f0f84b1cdc335c2cb4d6810ff0b9fd156ad6026c89cea.
//
// Solidity: event OtherReserveAddressAdded(address indexed otherReserveAddress)
func (_Reserve *ReserveFilterer) FilterOtherReserveAddressAdded(opts *bind.FilterOpts, otherReserveAddress []common.Address) (*ReserveOtherReserveAddressAddedIterator, error) {

	var otherReserveAddressRule []interface{}
	for _, otherReserveAddressItem := range otherReserveAddress {
		otherReserveAddressRule = append(otherReserveAddressRule, otherReserveAddressItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "OtherReserveAddressAdded", otherReserveAddressRule)
	if err != nil {
		return nil, err
	}
	return &ReserveOtherReserveAddressAddedIterator{contract: _Reserve.contract, event: "OtherReserveAddressAdded", logs: logs, sub: sub}, nil
}

// WatchOtherReserveAddressAdded is a free log subscription operation binding the contract event 0xd78793225285ecf9cf5f0f84b1cdc335c2cb4d6810ff0b9fd156ad6026c89cea.
//
// Solidity: event OtherReserveAddressAdded(address indexed otherReserveAddress)
func (_Reserve *ReserveFilterer) WatchOtherReserveAddressAdded(opts *bind.WatchOpts, sink chan<- *ReserveOtherReserveAddressAdded, otherReserveAddress []common.Address) (event.Subscription, error) {

	var otherReserveAddressRule []interface{}
	for _, otherReserveAddressItem := range otherReserveAddress {
		otherReserveAddressRule = append(otherReserveAddressRule, otherReserveAddressItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "OtherReserveAddressAdded", otherReserveAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveOtherReserveAddressAdded)
				if err := _Reserve.contract.UnpackLog(event, "OtherReserveAddressAdded", log); err != nil {
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

// ParseOtherReserveAddressAdded is a log parse operation binding the contract event 0xd78793225285ecf9cf5f0f84b1cdc335c2cb4d6810ff0b9fd156ad6026c89cea.
//
// Solidity: event OtherReserveAddressAdded(address indexed otherReserveAddress)
func (_Reserve *ReserveFilterer) ParseOtherReserveAddressAdded(log types.Log) (*ReserveOtherReserveAddressAdded, error) {
	event := new(ReserveOtherReserveAddressAdded)
	if err := _Reserve.contract.UnpackLog(event, "OtherReserveAddressAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveOtherReserveAddressRemovedIterator is returned from FilterOtherReserveAddressRemoved and is used to iterate over the raw logs and unpacked data for OtherReserveAddressRemoved events raised by the Reserve contract.
type ReserveOtherReserveAddressRemovedIterator struct {
	Event *ReserveOtherReserveAddressRemoved // Event containing the contract specifics and raw log

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
func (it *ReserveOtherReserveAddressRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveOtherReserveAddressRemoved)
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
		it.Event = new(ReserveOtherReserveAddressRemoved)
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
func (it *ReserveOtherReserveAddressRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveOtherReserveAddressRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveOtherReserveAddressRemoved represents a OtherReserveAddressRemoved event raised by the Reserve contract.
type ReserveOtherReserveAddressRemoved struct {
	OtherReserveAddress common.Address
	Index               *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterOtherReserveAddressRemoved is a free log retrieval operation binding the contract event 0x89b4ee5cecfdfb246ede373c10283b5038afe56a531fc1d2f3ed8c5507a52fcb.
//
// Solidity: event OtherReserveAddressRemoved(address indexed otherReserveAddress, uint256 index)
func (_Reserve *ReserveFilterer) FilterOtherReserveAddressRemoved(opts *bind.FilterOpts, otherReserveAddress []common.Address) (*ReserveOtherReserveAddressRemovedIterator, error) {

	var otherReserveAddressRule []interface{}
	for _, otherReserveAddressItem := range otherReserveAddress {
		otherReserveAddressRule = append(otherReserveAddressRule, otherReserveAddressItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "OtherReserveAddressRemoved", otherReserveAddressRule)
	if err != nil {
		return nil, err
	}
	return &ReserveOtherReserveAddressRemovedIterator{contract: _Reserve.contract, event: "OtherReserveAddressRemoved", logs: logs, sub: sub}, nil
}

// WatchOtherReserveAddressRemoved is a free log subscription operation binding the contract event 0x89b4ee5cecfdfb246ede373c10283b5038afe56a531fc1d2f3ed8c5507a52fcb.
//
// Solidity: event OtherReserveAddressRemoved(address indexed otherReserveAddress, uint256 index)
func (_Reserve *ReserveFilterer) WatchOtherReserveAddressRemoved(opts *bind.WatchOpts, sink chan<- *ReserveOtherReserveAddressRemoved, otherReserveAddress []common.Address) (event.Subscription, error) {

	var otherReserveAddressRule []interface{}
	for _, otherReserveAddressItem := range otherReserveAddress {
		otherReserveAddressRule = append(otherReserveAddressRule, otherReserveAddressItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "OtherReserveAddressRemoved", otherReserveAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveOtherReserveAddressRemoved)
				if err := _Reserve.contract.UnpackLog(event, "OtherReserveAddressRemoved", log); err != nil {
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

// ParseOtherReserveAddressRemoved is a log parse operation binding the contract event 0x89b4ee5cecfdfb246ede373c10283b5038afe56a531fc1d2f3ed8c5507a52fcb.
//
// Solidity: event OtherReserveAddressRemoved(address indexed otherReserveAddress, uint256 index)
func (_Reserve *ReserveFilterer) ParseOtherReserveAddressRemoved(log types.Log) (*ReserveOtherReserveAddressRemoved, error) {
	event := new(ReserveOtherReserveAddressRemoved)
	if err := _Reserve.contract.UnpackLog(event, "OtherReserveAddressRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Reserve contract.
type ReserveOwnershipTransferredIterator struct {
	Event *ReserveOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReserveOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveOwnershipTransferred)
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
		it.Event = new(ReserveOwnershipTransferred)
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
func (it *ReserveOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveOwnershipTransferred represents a OwnershipTransferred event raised by the Reserve contract.
type ReserveOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reserve *ReserveFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReserveOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReserveOwnershipTransferredIterator{contract: _Reserve.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Reserve *ReserveFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReserveOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveOwnershipTransferred)
				if err := _Reserve.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseOwnershipTransferred(log types.Log) (*ReserveOwnershipTransferred, error) {
	event := new(ReserveOwnershipTransferred)
	if err := _Reserve.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Reserve contract.
type ReserveRegistrySetIterator struct {
	Event *ReserveRegistrySet // Event containing the contract specifics and raw log

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
func (it *ReserveRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveRegistrySet)
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
		it.Event = new(ReserveRegistrySet)
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
func (it *ReserveRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveRegistrySet represents a RegistrySet event raised by the Reserve contract.
type ReserveRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Reserve *ReserveFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*ReserveRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &ReserveRegistrySetIterator{contract: _Reserve.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Reserve *ReserveFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *ReserveRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveRegistrySet)
				if err := _Reserve.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Reserve *ReserveFilterer) ParseRegistrySet(log types.Log) (*ReserveRegistrySet, error) {
	event := new(ReserveRegistrySet)
	if err := _Reserve.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveReserveGoldTransferredIterator is returned from FilterReserveGoldTransferred and is used to iterate over the raw logs and unpacked data for ReserveGoldTransferred events raised by the Reserve contract.
type ReserveReserveGoldTransferredIterator struct {
	Event *ReserveReserveGoldTransferred // Event containing the contract specifics and raw log

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
func (it *ReserveReserveGoldTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveReserveGoldTransferred)
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
		it.Event = new(ReserveReserveGoldTransferred)
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
func (it *ReserveReserveGoldTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveReserveGoldTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveReserveGoldTransferred represents a ReserveGoldTransferred event raised by the Reserve contract.
type ReserveReserveGoldTransferred struct {
	Spender common.Address
	To      common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveGoldTransferred is a free log retrieval operation binding the contract event 0x4dd1abe16ad3d4f829372dc77766ca2cce34e205af9b10f8cc1fab370425864f.
//
// Solidity: event ReserveGoldTransferred(address indexed spender, address indexed to, uint256 value)
func (_Reserve *ReserveFilterer) FilterReserveGoldTransferred(opts *bind.FilterOpts, spender []common.Address, to []common.Address) (*ReserveReserveGoldTransferredIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "ReserveGoldTransferred", spenderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ReserveReserveGoldTransferredIterator{contract: _Reserve.contract, event: "ReserveGoldTransferred", logs: logs, sub: sub}, nil
}

// WatchReserveGoldTransferred is a free log subscription operation binding the contract event 0x4dd1abe16ad3d4f829372dc77766ca2cce34e205af9b10f8cc1fab370425864f.
//
// Solidity: event ReserveGoldTransferred(address indexed spender, address indexed to, uint256 value)
func (_Reserve *ReserveFilterer) WatchReserveGoldTransferred(opts *bind.WatchOpts, sink chan<- *ReserveReserveGoldTransferred, spender []common.Address, to []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "ReserveGoldTransferred", spenderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveReserveGoldTransferred)
				if err := _Reserve.contract.UnpackLog(event, "ReserveGoldTransferred", log); err != nil {
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

// ParseReserveGoldTransferred is a log parse operation binding the contract event 0x4dd1abe16ad3d4f829372dc77766ca2cce34e205af9b10f8cc1fab370425864f.
//
// Solidity: event ReserveGoldTransferred(address indexed spender, address indexed to, uint256 value)
func (_Reserve *ReserveFilterer) ParseReserveGoldTransferred(log types.Log) (*ReserveReserveGoldTransferred, error) {
	event := new(ReserveReserveGoldTransferred)
	if err := _Reserve.contract.UnpackLog(event, "ReserveGoldTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderAddedIterator is returned from FilterSpenderAdded and is used to iterate over the raw logs and unpacked data for SpenderAdded events raised by the Reserve contract.
type ReserveSpenderAddedIterator struct {
	Event *ReserveSpenderAdded // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderAdded)
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
		it.Event = new(ReserveSpenderAdded)
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
func (it *ReserveSpenderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderAdded represents a SpenderAdded event raised by the Reserve contract.
type ReserveSpenderAdded struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderAdded is a free log retrieval operation binding the contract event 0x3139419c41cdd7abca84fa19dd21118cd285d3e2ce1a9444e8161ce9fa62fdcd.
//
// Solidity: event SpenderAdded(address indexed spender)
func (_Reserve *ReserveFilterer) FilterSpenderAdded(opts *bind.FilterOpts, spender []common.Address) (*ReserveSpenderAddedIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "SpenderAdded", spenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderAddedIterator{contract: _Reserve.contract, event: "SpenderAdded", logs: logs, sub: sub}, nil
}

// WatchSpenderAdded is a free log subscription operation binding the contract event 0x3139419c41cdd7abca84fa19dd21118cd285d3e2ce1a9444e8161ce9fa62fdcd.
//
// Solidity: event SpenderAdded(address indexed spender)
func (_Reserve *ReserveFilterer) WatchSpenderAdded(opts *bind.WatchOpts, sink chan<- *ReserveSpenderAdded, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "SpenderAdded", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderAdded)
				if err := _Reserve.contract.UnpackLog(event, "SpenderAdded", log); err != nil {
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

// ParseSpenderAdded is a log parse operation binding the contract event 0x3139419c41cdd7abca84fa19dd21118cd285d3e2ce1a9444e8161ce9fa62fdcd.
//
// Solidity: event SpenderAdded(address indexed spender)
func (_Reserve *ReserveFilterer) ParseSpenderAdded(log types.Log) (*ReserveSpenderAdded, error) {
	event := new(ReserveSpenderAdded)
	if err := _Reserve.contract.UnpackLog(event, "SpenderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveSpenderRemovedIterator is returned from FilterSpenderRemoved and is used to iterate over the raw logs and unpacked data for SpenderRemoved events raised by the Reserve contract.
type ReserveSpenderRemovedIterator struct {
	Event *ReserveSpenderRemoved // Event containing the contract specifics and raw log

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
func (it *ReserveSpenderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveSpenderRemoved)
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
		it.Event = new(ReserveSpenderRemoved)
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
func (it *ReserveSpenderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveSpenderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveSpenderRemoved represents a SpenderRemoved event raised by the Reserve contract.
type ReserveSpenderRemoved struct {
	Spender common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpenderRemoved is a free log retrieval operation binding the contract event 0xab8cff50266d80b9c9d9703af934ca455b9218286bf4fcaa05653a564c499e4b.
//
// Solidity: event SpenderRemoved(address indexed spender)
func (_Reserve *ReserveFilterer) FilterSpenderRemoved(opts *bind.FilterOpts, spender []common.Address) (*ReserveSpenderRemovedIterator, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "SpenderRemoved", spenderRule)
	if err != nil {
		return nil, err
	}
	return &ReserveSpenderRemovedIterator{contract: _Reserve.contract, event: "SpenderRemoved", logs: logs, sub: sub}, nil
}

// WatchSpenderRemoved is a free log subscription operation binding the contract event 0xab8cff50266d80b9c9d9703af934ca455b9218286bf4fcaa05653a564c499e4b.
//
// Solidity: event SpenderRemoved(address indexed spender)
func (_Reserve *ReserveFilterer) WatchSpenderRemoved(opts *bind.WatchOpts, sink chan<- *ReserveSpenderRemoved, spender []common.Address) (event.Subscription, error) {

	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "SpenderRemoved", spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveSpenderRemoved)
				if err := _Reserve.contract.UnpackLog(event, "SpenderRemoved", log); err != nil {
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

// ParseSpenderRemoved is a log parse operation binding the contract event 0xab8cff50266d80b9c9d9703af934ca455b9218286bf4fcaa05653a564c499e4b.
//
// Solidity: event SpenderRemoved(address indexed spender)
func (_Reserve *ReserveFilterer) ParseSpenderRemoved(log types.Log) (*ReserveSpenderRemoved, error) {
	event := new(ReserveSpenderRemoved)
	if err := _Reserve.contract.UnpackLog(event, "SpenderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveTobinTaxReserveRatioSetIterator is returned from FilterTobinTaxReserveRatioSet and is used to iterate over the raw logs and unpacked data for TobinTaxReserveRatioSet events raised by the Reserve contract.
type ReserveTobinTaxReserveRatioSetIterator struct {
	Event *ReserveTobinTaxReserveRatioSet // Event containing the contract specifics and raw log

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
func (it *ReserveTobinTaxReserveRatioSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTobinTaxReserveRatioSet)
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
		it.Event = new(ReserveTobinTaxReserveRatioSet)
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
func (it *ReserveTobinTaxReserveRatioSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTobinTaxReserveRatioSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTobinTaxReserveRatioSet represents a TobinTaxReserveRatioSet event raised by the Reserve contract.
type ReserveTobinTaxReserveRatioSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTobinTaxReserveRatioSet is a free log retrieval operation binding the contract event 0x4da8e8b2223fbbb897200fb9dfb6b986c1b4188621114d407ee8ec363569fc37.
//
// Solidity: event TobinTaxReserveRatioSet(uint256 value)
func (_Reserve *ReserveFilterer) FilterTobinTaxReserveRatioSet(opts *bind.FilterOpts) (*ReserveTobinTaxReserveRatioSetIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "TobinTaxReserveRatioSet")
	if err != nil {
		return nil, err
	}
	return &ReserveTobinTaxReserveRatioSetIterator{contract: _Reserve.contract, event: "TobinTaxReserveRatioSet", logs: logs, sub: sub}, nil
}

// WatchTobinTaxReserveRatioSet is a free log subscription operation binding the contract event 0x4da8e8b2223fbbb897200fb9dfb6b986c1b4188621114d407ee8ec363569fc37.
//
// Solidity: event TobinTaxReserveRatioSet(uint256 value)
func (_Reserve *ReserveFilterer) WatchTobinTaxReserveRatioSet(opts *bind.WatchOpts, sink chan<- *ReserveTobinTaxReserveRatioSet) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "TobinTaxReserveRatioSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTobinTaxReserveRatioSet)
				if err := _Reserve.contract.UnpackLog(event, "TobinTaxReserveRatioSet", log); err != nil {
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

// ParseTobinTaxReserveRatioSet is a log parse operation binding the contract event 0x4da8e8b2223fbbb897200fb9dfb6b986c1b4188621114d407ee8ec363569fc37.
//
// Solidity: event TobinTaxReserveRatioSet(uint256 value)
func (_Reserve *ReserveFilterer) ParseTobinTaxReserveRatioSet(log types.Log) (*ReserveTobinTaxReserveRatioSet, error) {
	event := new(ReserveTobinTaxReserveRatioSet)
	if err := _Reserve.contract.UnpackLog(event, "TobinTaxReserveRatioSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveTobinTaxSetIterator is returned from FilterTobinTaxSet and is used to iterate over the raw logs and unpacked data for TobinTaxSet events raised by the Reserve contract.
type ReserveTobinTaxSetIterator struct {
	Event *ReserveTobinTaxSet // Event containing the contract specifics and raw log

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
func (it *ReserveTobinTaxSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTobinTaxSet)
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
		it.Event = new(ReserveTobinTaxSet)
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
func (it *ReserveTobinTaxSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTobinTaxSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTobinTaxSet represents a TobinTaxSet event raised by the Reserve contract.
type ReserveTobinTaxSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTobinTaxSet is a free log retrieval operation binding the contract event 0xfe69856ffb1b1d6cb00c1d8151726e6e95032b1666282eeb293ecadd58b29a6e.
//
// Solidity: event TobinTaxSet(uint256 value)
func (_Reserve *ReserveFilterer) FilterTobinTaxSet(opts *bind.FilterOpts) (*ReserveTobinTaxSetIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "TobinTaxSet")
	if err != nil {
		return nil, err
	}
	return &ReserveTobinTaxSetIterator{contract: _Reserve.contract, event: "TobinTaxSet", logs: logs, sub: sub}, nil
}

// WatchTobinTaxSet is a free log subscription operation binding the contract event 0xfe69856ffb1b1d6cb00c1d8151726e6e95032b1666282eeb293ecadd58b29a6e.
//
// Solidity: event TobinTaxSet(uint256 value)
func (_Reserve *ReserveFilterer) WatchTobinTaxSet(opts *bind.WatchOpts, sink chan<- *ReserveTobinTaxSet) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "TobinTaxSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTobinTaxSet)
				if err := _Reserve.contract.UnpackLog(event, "TobinTaxSet", log); err != nil {
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

// ParseTobinTaxSet is a log parse operation binding the contract event 0xfe69856ffb1b1d6cb00c1d8151726e6e95032b1666282eeb293ecadd58b29a6e.
//
// Solidity: event TobinTaxSet(uint256 value)
func (_Reserve *ReserveFilterer) ParseTobinTaxSet(log types.Log) (*ReserveTobinTaxSet, error) {
	event := new(ReserveTobinTaxSet)
	if err := _Reserve.contract.UnpackLog(event, "TobinTaxSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveTobinTaxStalenessThresholdSetIterator is returned from FilterTobinTaxStalenessThresholdSet and is used to iterate over the raw logs and unpacked data for TobinTaxStalenessThresholdSet events raised by the Reserve contract.
type ReserveTobinTaxStalenessThresholdSetIterator struct {
	Event *ReserveTobinTaxStalenessThresholdSet // Event containing the contract specifics and raw log

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
func (it *ReserveTobinTaxStalenessThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTobinTaxStalenessThresholdSet)
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
		it.Event = new(ReserveTobinTaxStalenessThresholdSet)
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
func (it *ReserveTobinTaxStalenessThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTobinTaxStalenessThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTobinTaxStalenessThresholdSet represents a TobinTaxStalenessThresholdSet event raised by the Reserve contract.
type ReserveTobinTaxStalenessThresholdSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTobinTaxStalenessThresholdSet is a free log retrieval operation binding the contract event 0x7bfe94ca3147f135fcd6d94ebf61d33fa34fbe904f933ccae66911b9548544f2.
//
// Solidity: event TobinTaxStalenessThresholdSet(uint256 value)
func (_Reserve *ReserveFilterer) FilterTobinTaxStalenessThresholdSet(opts *bind.FilterOpts) (*ReserveTobinTaxStalenessThresholdSetIterator, error) {

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "TobinTaxStalenessThresholdSet")
	if err != nil {
		return nil, err
	}
	return &ReserveTobinTaxStalenessThresholdSetIterator{contract: _Reserve.contract, event: "TobinTaxStalenessThresholdSet", logs: logs, sub: sub}, nil
}

// WatchTobinTaxStalenessThresholdSet is a free log subscription operation binding the contract event 0x7bfe94ca3147f135fcd6d94ebf61d33fa34fbe904f933ccae66911b9548544f2.
//
// Solidity: event TobinTaxStalenessThresholdSet(uint256 value)
func (_Reserve *ReserveFilterer) WatchTobinTaxStalenessThresholdSet(opts *bind.WatchOpts, sink chan<- *ReserveTobinTaxStalenessThresholdSet) (event.Subscription, error) {

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "TobinTaxStalenessThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTobinTaxStalenessThresholdSet)
				if err := _Reserve.contract.UnpackLog(event, "TobinTaxStalenessThresholdSet", log); err != nil {
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

// ParseTobinTaxStalenessThresholdSet is a log parse operation binding the contract event 0x7bfe94ca3147f135fcd6d94ebf61d33fa34fbe904f933ccae66911b9548544f2.
//
// Solidity: event TobinTaxStalenessThresholdSet(uint256 value)
func (_Reserve *ReserveFilterer) ParseTobinTaxStalenessThresholdSet(log types.Log) (*ReserveTobinTaxStalenessThresholdSet, error) {
	event := new(ReserveTobinTaxStalenessThresholdSet)
	if err := _Reserve.contract.UnpackLog(event, "TobinTaxStalenessThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveTokenAddedIterator is returned from FilterTokenAdded and is used to iterate over the raw logs and unpacked data for TokenAdded events raised by the Reserve contract.
type ReserveTokenAddedIterator struct {
	Event *ReserveTokenAdded // Event containing the contract specifics and raw log

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
func (it *ReserveTokenAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTokenAdded)
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
		it.Event = new(ReserveTokenAdded)
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
func (it *ReserveTokenAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTokenAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTokenAdded represents a TokenAdded event raised by the Reserve contract.
type ReserveTokenAdded struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenAdded is a free log retrieval operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_Reserve *ReserveFilterer) FilterTokenAdded(opts *bind.FilterOpts, token []common.Address) (*ReserveTokenAddedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenAddedIterator{contract: _Reserve.contract, event: "TokenAdded", logs: logs, sub: sub}, nil
}

// WatchTokenAdded is a free log subscription operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_Reserve *ReserveFilterer) WatchTokenAdded(opts *bind.WatchOpts, sink chan<- *ReserveTokenAdded, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "TokenAdded", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTokenAdded)
				if err := _Reserve.contract.UnpackLog(event, "TokenAdded", log); err != nil {
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

// ParseTokenAdded is a log parse operation binding the contract event 0x784c8f4dbf0ffedd6e72c76501c545a70f8b203b30a26ce542bf92ba87c248a4.
//
// Solidity: event TokenAdded(address indexed token)
func (_Reserve *ReserveFilterer) ParseTokenAdded(log types.Log) (*ReserveTokenAdded, error) {
	event := new(ReserveTokenAdded)
	if err := _Reserve.contract.UnpackLog(event, "TokenAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveTokenRemovedIterator is returned from FilterTokenRemoved and is used to iterate over the raw logs and unpacked data for TokenRemoved events raised by the Reserve contract.
type ReserveTokenRemovedIterator struct {
	Event *ReserveTokenRemoved // Event containing the contract specifics and raw log

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
func (it *ReserveTokenRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReserveTokenRemoved)
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
		it.Event = new(ReserveTokenRemoved)
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
func (it *ReserveTokenRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReserveTokenRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReserveTokenRemoved represents a TokenRemoved event raised by the Reserve contract.
type ReserveTokenRemoved struct {
	Token common.Address
	Index *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenRemoved is a free log retrieval operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address indexed token, uint256 index)
func (_Reserve *ReserveFilterer) FilterTokenRemoved(opts *bind.FilterOpts, token []common.Address) (*ReserveTokenRemovedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Reserve.contract.FilterLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return &ReserveTokenRemovedIterator{contract: _Reserve.contract, event: "TokenRemoved", logs: logs, sub: sub}, nil
}

// WatchTokenRemoved is a free log subscription operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address indexed token, uint256 index)
func (_Reserve *ReserveFilterer) WatchTokenRemoved(opts *bind.WatchOpts, sink chan<- *ReserveTokenRemoved, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Reserve.contract.WatchLogs(opts, "TokenRemoved", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReserveTokenRemoved)
				if err := _Reserve.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
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

// ParseTokenRemoved is a log parse operation binding the contract event 0xbe9bb4bdca0a094babd75e3a98b1d2e2390633430d0a2f6e2b9970e2ee03fb2e.
//
// Solidity: event TokenRemoved(address indexed token, uint256 index)
func (_Reserve *ReserveFilterer) ParseTokenRemoved(log types.Log) (*ReserveTokenRemoved, error) {
	event := new(ReserveTokenRemoved)
	if err := _Reserve.contract.UnpackLog(event, "TokenRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
