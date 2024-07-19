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

// EpochRewardsMetaData contains all meta data concerning the EpochRewards contract.
var EpochRewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"CarbonOffsettingFundSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"CommunityRewardFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"underspendAdjustmentFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"overspendAdjustmentFactor\",\"type\":\"uint256\"}],\"name\":\"RewardsMultiplierParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"payment\",\"type\":\"uint256\"}],\"name\":\"TargetValidatorEpochPaymentSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"TargetVotingGoldFractionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"adjustmentFactor\",\"type\":\"uint256\"}],\"name\":\"TargetVotingYieldParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"TargetVotingYieldSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fraction\",\"type\":\"uint256\"}],\"name\":\"TargetVotingYieldUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"carbonOffsettingPartner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"targetValidatorEpochPayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"targetVotingYieldInitial\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetVotingYieldMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetVotingYieldAdjustmentFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsMultiplierMax\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsMultiplierUnderspendAdjustmentFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardsMultiplierOverspendAdjustmentFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetVotingGoldFraction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetValidatorEpochPayment\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_communityRewardFraction\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_carbonOffsettingPartner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_carbonOffsettingFraction\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetVotingYieldParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardsMultiplierParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setCommunityRewardFraction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommunityRewardFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"partner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setCarbonOffsettingFund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCarbonOffsettingFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setTargetVotingGoldFraction\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetVotingGoldFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setTargetValidatorEpochPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"underspendAdjustmentFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"overspendAdjustmentFactor\",\"type\":\"uint256\"}],\"name\":\"setRewardsMultiplierParameters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"adjustmentFactor\",\"type\":\"uint256\"}],\"name\":\"setTargetVotingYieldParameters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targetVotingYield\",\"type\":\"uint256\"}],\"name\":\"setTargetVotingYield\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetGoldTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetVoterRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTargetTotalEpochPaymentsInGold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRewardsMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVotingGoldFraction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateTargetVotingYield\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isReserveLow\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"calculateTargetEpochRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// EpochRewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use EpochRewardsMetaData.ABI instead.
var EpochRewardsABI = EpochRewardsMetaData.ABI

// EpochRewards is an auto generated Go binding around an Ethereum contract.
type EpochRewards struct {
	EpochRewardsCaller     // Read-only binding to the contract
	EpochRewardsTransactor // Write-only binding to the contract
	EpochRewardsFilterer   // Log filterer for contract events
}

// EpochRewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EpochRewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochRewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EpochRewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochRewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EpochRewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EpochRewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EpochRewardsSession struct {
	Contract     *EpochRewards     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EpochRewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EpochRewardsCallerSession struct {
	Contract *EpochRewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// EpochRewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EpochRewardsTransactorSession struct {
	Contract     *EpochRewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// EpochRewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EpochRewardsRaw struct {
	Contract *EpochRewards // Generic contract binding to access the raw methods on
}

// EpochRewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EpochRewardsCallerRaw struct {
	Contract *EpochRewardsCaller // Generic read-only contract binding to access the raw methods on
}

// EpochRewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EpochRewardsTransactorRaw struct {
	Contract *EpochRewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEpochRewards creates a new instance of EpochRewards, bound to a specific deployed contract.
func NewEpochRewards(address common.Address, backend bind.ContractBackend) (*EpochRewards, error) {
	contract, err := bindEpochRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EpochRewards{EpochRewardsCaller: EpochRewardsCaller{contract: contract}, EpochRewardsTransactor: EpochRewardsTransactor{contract: contract}, EpochRewardsFilterer: EpochRewardsFilterer{contract: contract}}, nil
}

// NewEpochRewardsCaller creates a new read-only instance of EpochRewards, bound to a specific deployed contract.
func NewEpochRewardsCaller(address common.Address, caller bind.ContractCaller) (*EpochRewardsCaller, error) {
	contract, err := bindEpochRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EpochRewardsCaller{contract: contract}, nil
}

// NewEpochRewardsTransactor creates a new write-only instance of EpochRewards, bound to a specific deployed contract.
func NewEpochRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*EpochRewardsTransactor, error) {
	contract, err := bindEpochRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EpochRewardsTransactor{contract: contract}, nil
}

// NewEpochRewardsFilterer creates a new log filterer instance of EpochRewards, bound to a specific deployed contract.
func NewEpochRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*EpochRewardsFilterer, error) {
	contract, err := bindEpochRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EpochRewardsFilterer{contract: contract}, nil
}

// bindEpochRewards binds a generic wrapper to an already deployed contract.
func bindEpochRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EpochRewardsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseEpochRewardsABI parses the ABI
func ParseEpochRewardsABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(EpochRewardsABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EpochRewards *EpochRewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EpochRewards.Contract.EpochRewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EpochRewards *EpochRewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochRewards.Contract.EpochRewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EpochRewards *EpochRewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EpochRewards.Contract.EpochRewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EpochRewards *EpochRewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EpochRewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EpochRewards *EpochRewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochRewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EpochRewards *EpochRewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EpochRewards.Contract.contract.Transact(opts, method, params...)
}

// CalculateTargetEpochRewards is a free data retrieval call binding the contract method 0x64347043.
//
// Solidity: function calculateTargetEpochRewards() view returns(uint256, uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCaller) CalculateTargetEpochRewards(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "calculateTargetEpochRewards")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// CalculateTargetEpochRewards is a free data retrieval call binding the contract method 0x64347043.
//
// Solidity: function calculateTargetEpochRewards() view returns(uint256, uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsSession) CalculateTargetEpochRewards() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.CalculateTargetEpochRewards(&_EpochRewards.CallOpts)
}

// CalculateTargetEpochRewards is a free data retrieval call binding the contract method 0x64347043.
//
// Solidity: function calculateTargetEpochRewards() view returns(uint256, uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCallerSession) CalculateTargetEpochRewards() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.CalculateTargetEpochRewards(&_EpochRewards.CallOpts)
}

// CarbonOffsettingPartner is a free data retrieval call binding the contract method 0x22dae21f.
//
// Solidity: function carbonOffsettingPartner() view returns(address)
func (_EpochRewards *EpochRewardsCaller) CarbonOffsettingPartner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "carbonOffsettingPartner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CarbonOffsettingPartner is a free data retrieval call binding the contract method 0x22dae21f.
//
// Solidity: function carbonOffsettingPartner() view returns(address)
func (_EpochRewards *EpochRewardsSession) CarbonOffsettingPartner() (common.Address, error) {
	return _EpochRewards.Contract.CarbonOffsettingPartner(&_EpochRewards.CallOpts)
}

// CarbonOffsettingPartner is a free data retrieval call binding the contract method 0x22dae21f.
//
// Solidity: function carbonOffsettingPartner() view returns(address)
func (_EpochRewards *EpochRewardsCallerSession) CarbonOffsettingPartner() (common.Address, error) {
	return _EpochRewards.Contract.CarbonOffsettingPartner(&_EpochRewards.CallOpts)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_EpochRewards *EpochRewardsCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_EpochRewards *EpochRewardsSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _EpochRewards.Contract.CheckProofOfPossession(&_EpochRewards.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_EpochRewards *EpochRewardsCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _EpochRewards.Contract.CheckProofOfPossession(&_EpochRewards.CallOpts, sender, blsKey, blsPop)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_EpochRewards *EpochRewardsCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

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
func (_EpochRewards *EpochRewardsSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _EpochRewards.Contract.FractionMulExp(&_EpochRewards.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_EpochRewards *EpochRewardsCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _EpochRewards.Contract.FractionMulExp(&_EpochRewards.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _EpochRewards.Contract.GetBlockNumberFromHeader(&_EpochRewards.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _EpochRewards.Contract.GetBlockNumberFromHeader(&_EpochRewards.CallOpts, header)
}

// GetCarbonOffsettingFraction is a free data retrieval call binding the contract method 0x7d164125.
//
// Solidity: function getCarbonOffsettingFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetCarbonOffsettingFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getCarbonOffsettingFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCarbonOffsettingFraction is a free data retrieval call binding the contract method 0x7d164125.
//
// Solidity: function getCarbonOffsettingFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetCarbonOffsettingFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetCarbonOffsettingFraction(&_EpochRewards.CallOpts)
}

// GetCarbonOffsettingFraction is a free data retrieval call binding the contract method 0x7d164125.
//
// Solidity: function getCarbonOffsettingFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetCarbonOffsettingFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetCarbonOffsettingFraction(&_EpochRewards.CallOpts)
}

// GetCommunityRewardFraction is a free data retrieval call binding the contract method 0x9917907f.
//
// Solidity: function getCommunityRewardFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetCommunityRewardFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getCommunityRewardFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommunityRewardFraction is a free data retrieval call binding the contract method 0x9917907f.
//
// Solidity: function getCommunityRewardFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetCommunityRewardFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetCommunityRewardFraction(&_EpochRewards.CallOpts)
}

// GetCommunityRewardFraction is a free data retrieval call binding the contract method 0x9917907f.
//
// Solidity: function getCommunityRewardFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetCommunityRewardFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetCommunityRewardFraction(&_EpochRewards.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetEpochNumber() (*big.Int, error) {
	return _EpochRewards.Contract.GetEpochNumber(&_EpochRewards.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetEpochNumber() (*big.Int, error) {
	return _EpochRewards.Contract.GetEpochNumber(&_EpochRewards.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _EpochRewards.Contract.GetEpochNumberOfBlock(&_EpochRewards.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _EpochRewards.Contract.GetEpochNumberOfBlock(&_EpochRewards.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetEpochSize() (*big.Int, error) {
	return _EpochRewards.Contract.GetEpochSize(&_EpochRewards.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetEpochSize() (*big.Int, error) {
	return _EpochRewards.Contract.GetEpochSize(&_EpochRewards.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_EpochRewards *EpochRewardsCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_EpochRewards *EpochRewardsSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _EpochRewards.Contract.GetParentSealBitmap(&_EpochRewards.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_EpochRewards *EpochRewardsCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _EpochRewards.Contract.GetParentSealBitmap(&_EpochRewards.CallOpts, blockNumber)
}

// GetRewardsMultiplier is a free data retrieval call binding the contract method 0x0203ab24.
//
// Solidity: function getRewardsMultiplier() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetRewardsMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getRewardsMultiplier")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardsMultiplier is a free data retrieval call binding the contract method 0x0203ab24.
//
// Solidity: function getRewardsMultiplier() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetRewardsMultiplier() (*big.Int, error) {
	return _EpochRewards.Contract.GetRewardsMultiplier(&_EpochRewards.CallOpts)
}

// GetRewardsMultiplier is a free data retrieval call binding the contract method 0x0203ab24.
//
// Solidity: function getRewardsMultiplier() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetRewardsMultiplier() (*big.Int, error) {
	return _EpochRewards.Contract.GetRewardsMultiplier(&_EpochRewards.CallOpts)
}

// GetRewardsMultiplierParameters is a free data retrieval call binding the contract method 0x5f396e48.
//
// Solidity: function getRewardsMultiplierParameters() view returns(uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCaller) GetRewardsMultiplierParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getRewardsMultiplierParameters")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetRewardsMultiplierParameters is a free data retrieval call binding the contract method 0x5f396e48.
//
// Solidity: function getRewardsMultiplierParameters() view returns(uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsSession) GetRewardsMultiplierParameters() (*big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.GetRewardsMultiplierParameters(&_EpochRewards.CallOpts)
}

// GetRewardsMultiplierParameters is a free data retrieval call binding the contract method 0x5f396e48.
//
// Solidity: function getRewardsMultiplierParameters() view returns(uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetRewardsMultiplierParameters() (*big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.GetRewardsMultiplierParameters(&_EpochRewards.CallOpts)
}

// GetTargetGoldTotalSupply is a free data retrieval call binding the contract method 0x5049890f.
//
// Solidity: function getTargetGoldTotalSupply() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetTargetGoldTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getTargetGoldTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetGoldTotalSupply is a free data retrieval call binding the contract method 0x5049890f.
//
// Solidity: function getTargetGoldTotalSupply() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetTargetGoldTotalSupply() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetGoldTotalSupply(&_EpochRewards.CallOpts)
}

// GetTargetGoldTotalSupply is a free data retrieval call binding the contract method 0x5049890f.
//
// Solidity: function getTargetGoldTotalSupply() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetTargetGoldTotalSupply() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetGoldTotalSupply(&_EpochRewards.CallOpts)
}

// GetTargetTotalEpochPaymentsInGold is a free data retrieval call binding the contract method 0x4901c725.
//
// Solidity: function getTargetTotalEpochPaymentsInGold() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetTargetTotalEpochPaymentsInGold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getTargetTotalEpochPaymentsInGold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetTotalEpochPaymentsInGold is a free data retrieval call binding the contract method 0x4901c725.
//
// Solidity: function getTargetTotalEpochPaymentsInGold() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetTargetTotalEpochPaymentsInGold() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetTotalEpochPaymentsInGold(&_EpochRewards.CallOpts)
}

// GetTargetTotalEpochPaymentsInGold is a free data retrieval call binding the contract method 0x4901c725.
//
// Solidity: function getTargetTotalEpochPaymentsInGold() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetTargetTotalEpochPaymentsInGold() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetTotalEpochPaymentsInGold(&_EpochRewards.CallOpts)
}

// GetTargetVoterRewards is a free data retrieval call binding the contract method 0x2848f9e3.
//
// Solidity: function getTargetVoterRewards() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetTargetVoterRewards(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getTargetVoterRewards")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetVoterRewards is a free data retrieval call binding the contract method 0x2848f9e3.
//
// Solidity: function getTargetVoterRewards() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetTargetVoterRewards() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetVoterRewards(&_EpochRewards.CallOpts)
}

// GetTargetVoterRewards is a free data retrieval call binding the contract method 0x2848f9e3.
//
// Solidity: function getTargetVoterRewards() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetTargetVoterRewards() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetVoterRewards(&_EpochRewards.CallOpts)
}

// GetTargetVotingGoldFraction is a free data retrieval call binding the contract method 0xae098de2.
//
// Solidity: function getTargetVotingGoldFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetTargetVotingGoldFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getTargetVotingGoldFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTargetVotingGoldFraction is a free data retrieval call binding the contract method 0xae098de2.
//
// Solidity: function getTargetVotingGoldFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetTargetVotingGoldFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetVotingGoldFraction(&_EpochRewards.CallOpts)
}

// GetTargetVotingGoldFraction is a free data retrieval call binding the contract method 0xae098de2.
//
// Solidity: function getTargetVotingGoldFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetTargetVotingGoldFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetTargetVotingGoldFraction(&_EpochRewards.CallOpts)
}

// GetTargetVotingYieldParameters is a free data retrieval call binding the contract method 0x171af90f.
//
// Solidity: function getTargetVotingYieldParameters() view returns(uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCaller) GetTargetVotingYieldParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getTargetVotingYieldParameters")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetTargetVotingYieldParameters is a free data retrieval call binding the contract method 0x171af90f.
//
// Solidity: function getTargetVotingYieldParameters() view returns(uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsSession) GetTargetVotingYieldParameters() (*big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.GetTargetVotingYieldParameters(&_EpochRewards.CallOpts)
}

// GetTargetVotingYieldParameters is a free data retrieval call binding the contract method 0x171af90f.
//
// Solidity: function getTargetVotingYieldParameters() view returns(uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetTargetVotingYieldParameters() (*big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.GetTargetVotingYieldParameters(&_EpochRewards.CallOpts)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_EpochRewards *EpochRewardsCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_EpochRewards *EpochRewardsSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _EpochRewards.Contract.GetVerifiedSealBitmapFromHeader(&_EpochRewards.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_EpochRewards *EpochRewardsCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _EpochRewards.Contract.GetVerifiedSealBitmapFromHeader(&_EpochRewards.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getVersionNumber")

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
func (_EpochRewards *EpochRewardsSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.GetVersionNumber(&_EpochRewards.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _EpochRewards.Contract.GetVersionNumber(&_EpochRewards.CallOpts)
}

// GetVotingGoldFraction is a free data retrieval call binding the contract method 0xa1b95962.
//
// Solidity: function getVotingGoldFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) GetVotingGoldFraction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "getVotingGoldFraction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotingGoldFraction is a free data retrieval call binding the contract method 0xa1b95962.
//
// Solidity: function getVotingGoldFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) GetVotingGoldFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetVotingGoldFraction(&_EpochRewards.CallOpts)
}

// GetVotingGoldFraction is a free data retrieval call binding the contract method 0xa1b95962.
//
// Solidity: function getVotingGoldFraction() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) GetVotingGoldFraction() (*big.Int, error) {
	return _EpochRewards.Contract.GetVotingGoldFraction(&_EpochRewards.CallOpts)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_EpochRewards *EpochRewardsCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_EpochRewards *EpochRewardsSession) HashHeader(header []byte) ([32]byte, error) {
	return _EpochRewards.Contract.HashHeader(&_EpochRewards.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_EpochRewards *EpochRewardsCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _EpochRewards.Contract.HashHeader(&_EpochRewards.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_EpochRewards *EpochRewardsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_EpochRewards *EpochRewardsSession) Initialized() (bool, error) {
	return _EpochRewards.Contract.Initialized(&_EpochRewards.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_EpochRewards *EpochRewardsCallerSession) Initialized() (bool, error) {
	return _EpochRewards.Contract.Initialized(&_EpochRewards.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EpochRewards *EpochRewardsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EpochRewards *EpochRewardsSession) IsOwner() (bool, error) {
	return _EpochRewards.Contract.IsOwner(&_EpochRewards.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_EpochRewards *EpochRewardsCallerSession) IsOwner() (bool, error) {
	return _EpochRewards.Contract.IsOwner(&_EpochRewards.CallOpts)
}

// IsReserveLow is a free data retrieval call binding the contract method 0x9ad0cce7.
//
// Solidity: function isReserveLow() view returns(bool)
func (_EpochRewards *EpochRewardsCaller) IsReserveLow(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "isReserveLow")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReserveLow is a free data retrieval call binding the contract method 0x9ad0cce7.
//
// Solidity: function isReserveLow() view returns(bool)
func (_EpochRewards *EpochRewardsSession) IsReserveLow() (bool, error) {
	return _EpochRewards.Contract.IsReserveLow(&_EpochRewards.CallOpts)
}

// IsReserveLow is a free data retrieval call binding the contract method 0x9ad0cce7.
//
// Solidity: function isReserveLow() view returns(bool)
func (_EpochRewards *EpochRewardsCallerSession) IsReserveLow() (bool, error) {
	return _EpochRewards.Contract.IsReserveLow(&_EpochRewards.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _EpochRewards.Contract.MinQuorumSize(&_EpochRewards.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _EpochRewards.Contract.MinQuorumSize(&_EpochRewards.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _EpochRewards.Contract.MinQuorumSizeInCurrentSet(&_EpochRewards.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _EpochRewards.Contract.MinQuorumSizeInCurrentSet(&_EpochRewards.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _EpochRewards.Contract.NumberValidatorsInCurrentSet(&_EpochRewards.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _EpochRewards.Contract.NumberValidatorsInCurrentSet(&_EpochRewards.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _EpochRewards.Contract.NumberValidatorsInSet(&_EpochRewards.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _EpochRewards.Contract.NumberValidatorsInSet(&_EpochRewards.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EpochRewards *EpochRewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EpochRewards *EpochRewardsSession) Owner() (common.Address, error) {
	return _EpochRewards.Contract.Owner(&_EpochRewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_EpochRewards *EpochRewardsCallerSession) Owner() (common.Address, error) {
	return _EpochRewards.Contract.Owner(&_EpochRewards.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EpochRewards *EpochRewardsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EpochRewards *EpochRewardsSession) Registry() (common.Address, error) {
	return _EpochRewards.Contract.Registry(&_EpochRewards.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_EpochRewards *EpochRewardsCallerSession) Registry() (common.Address, error) {
	return _EpochRewards.Contract.Registry(&_EpochRewards.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) StartTime() (*big.Int, error) {
	return _EpochRewards.Contract.StartTime(&_EpochRewards.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) StartTime() (*big.Int, error) {
	return _EpochRewards.Contract.StartTime(&_EpochRewards.CallOpts)
}

// TargetValidatorEpochPayment is a free data retrieval call binding the contract method 0xe185aaa8.
//
// Solidity: function targetValidatorEpochPayment() view returns(uint256)
func (_EpochRewards *EpochRewardsCaller) TargetValidatorEpochPayment(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "targetValidatorEpochPayment")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetValidatorEpochPayment is a free data retrieval call binding the contract method 0xe185aaa8.
//
// Solidity: function targetValidatorEpochPayment() view returns(uint256)
func (_EpochRewards *EpochRewardsSession) TargetValidatorEpochPayment() (*big.Int, error) {
	return _EpochRewards.Contract.TargetValidatorEpochPayment(&_EpochRewards.CallOpts)
}

// TargetValidatorEpochPayment is a free data retrieval call binding the contract method 0xe185aaa8.
//
// Solidity: function targetValidatorEpochPayment() view returns(uint256)
func (_EpochRewards *EpochRewardsCallerSession) TargetValidatorEpochPayment() (*big.Int, error) {
	return _EpochRewards.Contract.TargetValidatorEpochPayment(&_EpochRewards.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_EpochRewards *EpochRewardsCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_EpochRewards *EpochRewardsSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _EpochRewards.Contract.ValidatorSignerAddressFromCurrentSet(&_EpochRewards.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_EpochRewards *EpochRewardsCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _EpochRewards.Contract.ValidatorSignerAddressFromCurrentSet(&_EpochRewards.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_EpochRewards *EpochRewardsCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _EpochRewards.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_EpochRewards *EpochRewardsSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _EpochRewards.Contract.ValidatorSignerAddressFromSet(&_EpochRewards.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_EpochRewards *EpochRewardsCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _EpochRewards.Contract.ValidatorSignerAddressFromSet(&_EpochRewards.CallOpts, index, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cca2a3c.
//
// Solidity: function initialize(address registryAddress, uint256 targetVotingYieldInitial, uint256 targetVotingYieldMax, uint256 targetVotingYieldAdjustmentFactor, uint256 rewardsMultiplierMax, uint256 rewardsMultiplierUnderspendAdjustmentFactor, uint256 rewardsMultiplierOverspendAdjustmentFactor, uint256 _targetVotingGoldFraction, uint256 _targetValidatorEpochPayment, uint256 _communityRewardFraction, address _carbonOffsettingPartner, uint256 _carbonOffsettingFraction) returns()
func (_EpochRewards *EpochRewardsTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, targetVotingYieldInitial *big.Int, targetVotingYieldMax *big.Int, targetVotingYieldAdjustmentFactor *big.Int, rewardsMultiplierMax *big.Int, rewardsMultiplierUnderspendAdjustmentFactor *big.Int, rewardsMultiplierOverspendAdjustmentFactor *big.Int, _targetVotingGoldFraction *big.Int, _targetValidatorEpochPayment *big.Int, _communityRewardFraction *big.Int, _carbonOffsettingPartner common.Address, _carbonOffsettingFraction *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "initialize", registryAddress, targetVotingYieldInitial, targetVotingYieldMax, targetVotingYieldAdjustmentFactor, rewardsMultiplierMax, rewardsMultiplierUnderspendAdjustmentFactor, rewardsMultiplierOverspendAdjustmentFactor, _targetVotingGoldFraction, _targetValidatorEpochPayment, _communityRewardFraction, _carbonOffsettingPartner, _carbonOffsettingFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cca2a3c.
//
// Solidity: function initialize(address registryAddress, uint256 targetVotingYieldInitial, uint256 targetVotingYieldMax, uint256 targetVotingYieldAdjustmentFactor, uint256 rewardsMultiplierMax, uint256 rewardsMultiplierUnderspendAdjustmentFactor, uint256 rewardsMultiplierOverspendAdjustmentFactor, uint256 _targetVotingGoldFraction, uint256 _targetValidatorEpochPayment, uint256 _communityRewardFraction, address _carbonOffsettingPartner, uint256 _carbonOffsettingFraction) returns()
func (_EpochRewards *EpochRewardsSession) Initialize(registryAddress common.Address, targetVotingYieldInitial *big.Int, targetVotingYieldMax *big.Int, targetVotingYieldAdjustmentFactor *big.Int, rewardsMultiplierMax *big.Int, rewardsMultiplierUnderspendAdjustmentFactor *big.Int, rewardsMultiplierOverspendAdjustmentFactor *big.Int, _targetVotingGoldFraction *big.Int, _targetValidatorEpochPayment *big.Int, _communityRewardFraction *big.Int, _carbonOffsettingPartner common.Address, _carbonOffsettingFraction *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.Initialize(&_EpochRewards.TransactOpts, registryAddress, targetVotingYieldInitial, targetVotingYieldMax, targetVotingYieldAdjustmentFactor, rewardsMultiplierMax, rewardsMultiplierUnderspendAdjustmentFactor, rewardsMultiplierOverspendAdjustmentFactor, _targetVotingGoldFraction, _targetValidatorEpochPayment, _communityRewardFraction, _carbonOffsettingPartner, _carbonOffsettingFraction)
}

// Initialize is a paid mutator transaction binding the contract method 0x7cca2a3c.
//
// Solidity: function initialize(address registryAddress, uint256 targetVotingYieldInitial, uint256 targetVotingYieldMax, uint256 targetVotingYieldAdjustmentFactor, uint256 rewardsMultiplierMax, uint256 rewardsMultiplierUnderspendAdjustmentFactor, uint256 rewardsMultiplierOverspendAdjustmentFactor, uint256 _targetVotingGoldFraction, uint256 _targetValidatorEpochPayment, uint256 _communityRewardFraction, address _carbonOffsettingPartner, uint256 _carbonOffsettingFraction) returns()
func (_EpochRewards *EpochRewardsTransactorSession) Initialize(registryAddress common.Address, targetVotingYieldInitial *big.Int, targetVotingYieldMax *big.Int, targetVotingYieldAdjustmentFactor *big.Int, rewardsMultiplierMax *big.Int, rewardsMultiplierUnderspendAdjustmentFactor *big.Int, rewardsMultiplierOverspendAdjustmentFactor *big.Int, _targetVotingGoldFraction *big.Int, _targetValidatorEpochPayment *big.Int, _communityRewardFraction *big.Int, _carbonOffsettingPartner common.Address, _carbonOffsettingFraction *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.Initialize(&_EpochRewards.TransactOpts, registryAddress, targetVotingYieldInitial, targetVotingYieldMax, targetVotingYieldAdjustmentFactor, rewardsMultiplierMax, rewardsMultiplierUnderspendAdjustmentFactor, rewardsMultiplierOverspendAdjustmentFactor, _targetVotingGoldFraction, _targetValidatorEpochPayment, _communityRewardFraction, _carbonOffsettingPartner, _carbonOffsettingFraction)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EpochRewards *EpochRewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EpochRewards *EpochRewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _EpochRewards.Contract.RenounceOwnership(&_EpochRewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_EpochRewards *EpochRewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _EpochRewards.Contract.RenounceOwnership(&_EpochRewards.TransactOpts)
}

// SetCarbonOffsettingFund is a paid mutator transaction binding the contract method 0x434c99c4.
//
// Solidity: function setCarbonOffsettingFund(address partner, uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetCarbonOffsettingFund(opts *bind.TransactOpts, partner common.Address, value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setCarbonOffsettingFund", partner, value)
}

// SetCarbonOffsettingFund is a paid mutator transaction binding the contract method 0x434c99c4.
//
// Solidity: function setCarbonOffsettingFund(address partner, uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetCarbonOffsettingFund(partner common.Address, value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetCarbonOffsettingFund(&_EpochRewards.TransactOpts, partner, value)
}

// SetCarbonOffsettingFund is a paid mutator transaction binding the contract method 0x434c99c4.
//
// Solidity: function setCarbonOffsettingFund(address partner, uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetCarbonOffsettingFund(partner common.Address, value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetCarbonOffsettingFund(&_EpochRewards.TransactOpts, partner, value)
}

// SetCommunityRewardFraction is a paid mutator transaction binding the contract method 0xb63b4a23.
//
// Solidity: function setCommunityRewardFraction(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetCommunityRewardFraction(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setCommunityRewardFraction", value)
}

// SetCommunityRewardFraction is a paid mutator transaction binding the contract method 0xb63b4a23.
//
// Solidity: function setCommunityRewardFraction(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetCommunityRewardFraction(value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetCommunityRewardFraction(&_EpochRewards.TransactOpts, value)
}

// SetCommunityRewardFraction is a paid mutator transaction binding the contract method 0xb63b4a23.
//
// Solidity: function setCommunityRewardFraction(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetCommunityRewardFraction(value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetCommunityRewardFraction(&_EpochRewards.TransactOpts, value)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_EpochRewards *EpochRewardsTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_EpochRewards *EpochRewardsSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetRegistry(&_EpochRewards.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_EpochRewards *EpochRewardsTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetRegistry(&_EpochRewards.TransactOpts, registryAddress)
}

// SetRewardsMultiplierParameters is a paid mutator transaction binding the contract method 0x8331c1d7.
//
// Solidity: function setRewardsMultiplierParameters(uint256 max, uint256 underspendAdjustmentFactor, uint256 overspendAdjustmentFactor) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetRewardsMultiplierParameters(opts *bind.TransactOpts, max *big.Int, underspendAdjustmentFactor *big.Int, overspendAdjustmentFactor *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setRewardsMultiplierParameters", max, underspendAdjustmentFactor, overspendAdjustmentFactor)
}

// SetRewardsMultiplierParameters is a paid mutator transaction binding the contract method 0x8331c1d7.
//
// Solidity: function setRewardsMultiplierParameters(uint256 max, uint256 underspendAdjustmentFactor, uint256 overspendAdjustmentFactor) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetRewardsMultiplierParameters(max *big.Int, underspendAdjustmentFactor *big.Int, overspendAdjustmentFactor *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetRewardsMultiplierParameters(&_EpochRewards.TransactOpts, max, underspendAdjustmentFactor, overspendAdjustmentFactor)
}

// SetRewardsMultiplierParameters is a paid mutator transaction binding the contract method 0x8331c1d7.
//
// Solidity: function setRewardsMultiplierParameters(uint256 max, uint256 underspendAdjustmentFactor, uint256 overspendAdjustmentFactor) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetRewardsMultiplierParameters(max *big.Int, underspendAdjustmentFactor *big.Int, overspendAdjustmentFactor *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetRewardsMultiplierParameters(&_EpochRewards.TransactOpts, max, underspendAdjustmentFactor, overspendAdjustmentFactor)
}

// SetTargetValidatorEpochPayment is a paid mutator transaction binding the contract method 0x96c3d2fd.
//
// Solidity: function setTargetValidatorEpochPayment(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetTargetValidatorEpochPayment(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setTargetValidatorEpochPayment", value)
}

// SetTargetValidatorEpochPayment is a paid mutator transaction binding the contract method 0x96c3d2fd.
//
// Solidity: function setTargetValidatorEpochPayment(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetTargetValidatorEpochPayment(value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetValidatorEpochPayment(&_EpochRewards.TransactOpts, value)
}

// SetTargetValidatorEpochPayment is a paid mutator transaction binding the contract method 0x96c3d2fd.
//
// Solidity: function setTargetValidatorEpochPayment(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetTargetValidatorEpochPayment(value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetValidatorEpochPayment(&_EpochRewards.TransactOpts, value)
}

// SetTargetVotingGoldFraction is a paid mutator transaction binding the contract method 0x94028384.
//
// Solidity: function setTargetVotingGoldFraction(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetTargetVotingGoldFraction(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setTargetVotingGoldFraction", value)
}

// SetTargetVotingGoldFraction is a paid mutator transaction binding the contract method 0x94028384.
//
// Solidity: function setTargetVotingGoldFraction(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetTargetVotingGoldFraction(value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetVotingGoldFraction(&_EpochRewards.TransactOpts, value)
}

// SetTargetVotingGoldFraction is a paid mutator transaction binding the contract method 0x94028384.
//
// Solidity: function setTargetVotingGoldFraction(uint256 value) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetTargetVotingGoldFraction(value *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetVotingGoldFraction(&_EpochRewards.TransactOpts, value)
}

// SetTargetVotingYield is a paid mutator transaction binding the contract method 0xcd52782e.
//
// Solidity: function setTargetVotingYield(uint256 targetVotingYield) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetTargetVotingYield(opts *bind.TransactOpts, targetVotingYield *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setTargetVotingYield", targetVotingYield)
}

// SetTargetVotingYield is a paid mutator transaction binding the contract method 0xcd52782e.
//
// Solidity: function setTargetVotingYield(uint256 targetVotingYield) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetTargetVotingYield(targetVotingYield *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetVotingYield(&_EpochRewards.TransactOpts, targetVotingYield)
}

// SetTargetVotingYield is a paid mutator transaction binding the contract method 0xcd52782e.
//
// Solidity: function setTargetVotingYield(uint256 targetVotingYield) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetTargetVotingYield(targetVotingYield *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetVotingYield(&_EpochRewards.TransactOpts, targetVotingYield)
}

// SetTargetVotingYieldParameters is a paid mutator transaction binding the contract method 0x5918bb58.
//
// Solidity: function setTargetVotingYieldParameters(uint256 max, uint256 adjustmentFactor) returns(bool)
func (_EpochRewards *EpochRewardsTransactor) SetTargetVotingYieldParameters(opts *bind.TransactOpts, max *big.Int, adjustmentFactor *big.Int) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "setTargetVotingYieldParameters", max, adjustmentFactor)
}

// SetTargetVotingYieldParameters is a paid mutator transaction binding the contract method 0x5918bb58.
//
// Solidity: function setTargetVotingYieldParameters(uint256 max, uint256 adjustmentFactor) returns(bool)
func (_EpochRewards *EpochRewardsSession) SetTargetVotingYieldParameters(max *big.Int, adjustmentFactor *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetVotingYieldParameters(&_EpochRewards.TransactOpts, max, adjustmentFactor)
}

// SetTargetVotingYieldParameters is a paid mutator transaction binding the contract method 0x5918bb58.
//
// Solidity: function setTargetVotingYieldParameters(uint256 max, uint256 adjustmentFactor) returns(bool)
func (_EpochRewards *EpochRewardsTransactorSession) SetTargetVotingYieldParameters(max *big.Int, adjustmentFactor *big.Int) (*types.Transaction, error) {
	return _EpochRewards.Contract.SetTargetVotingYieldParameters(&_EpochRewards.TransactOpts, max, adjustmentFactor)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EpochRewards *EpochRewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EpochRewards *EpochRewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EpochRewards.Contract.TransferOwnership(&_EpochRewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_EpochRewards *EpochRewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _EpochRewards.Contract.TransferOwnership(&_EpochRewards.TransactOpts, newOwner)
}

// UpdateTargetVotingYield is a paid mutator transaction binding the contract method 0x92ecd745.
//
// Solidity: function updateTargetVotingYield() returns()
func (_EpochRewards *EpochRewardsTransactor) UpdateTargetVotingYield(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EpochRewards.contract.Transact(opts, "updateTargetVotingYield")
}

// UpdateTargetVotingYield is a paid mutator transaction binding the contract method 0x92ecd745.
//
// Solidity: function updateTargetVotingYield() returns()
func (_EpochRewards *EpochRewardsSession) UpdateTargetVotingYield() (*types.Transaction, error) {
	return _EpochRewards.Contract.UpdateTargetVotingYield(&_EpochRewards.TransactOpts)
}

// UpdateTargetVotingYield is a paid mutator transaction binding the contract method 0x92ecd745.
//
// Solidity: function updateTargetVotingYield() returns()
func (_EpochRewards *EpochRewardsTransactorSession) UpdateTargetVotingYield() (*types.Transaction, error) {
	return _EpochRewards.Contract.UpdateTargetVotingYield(&_EpochRewards.TransactOpts)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_EpochRewards *EpochRewardsFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _EpochRewards.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "CarbonOffsettingFundSet":
		event, err = _EpochRewards.ParseCarbonOffsettingFundSet(log)
	case "CommunityRewardFractionSet":
		event, err = _EpochRewards.ParseCommunityRewardFractionSet(log)
	case "OwnershipTransferred":
		event, err = _EpochRewards.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _EpochRewards.ParseRegistrySet(log)
	case "RewardsMultiplierParametersSet":
		event, err = _EpochRewards.ParseRewardsMultiplierParametersSet(log)
	case "TargetValidatorEpochPaymentSet":
		event, err = _EpochRewards.ParseTargetValidatorEpochPaymentSet(log)
	case "TargetVotingGoldFractionSet":
		event, err = _EpochRewards.ParseTargetVotingGoldFractionSet(log)
	case "TargetVotingYieldParametersSet":
		event, err = _EpochRewards.ParseTargetVotingYieldParametersSet(log)
	case "TargetVotingYieldSet":
		event, err = _EpochRewards.ParseTargetVotingYieldSet(log)
	case "TargetVotingYieldUpdated":
		event, err = _EpochRewards.ParseTargetVotingYieldUpdated(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// EpochRewardsCarbonOffsettingFundSetIterator is returned from FilterCarbonOffsettingFundSet and is used to iterate over the raw logs and unpacked data for CarbonOffsettingFundSet events raised by the EpochRewards contract.
type EpochRewardsCarbonOffsettingFundSetIterator struct {
	Event *EpochRewardsCarbonOffsettingFundSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsCarbonOffsettingFundSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsCarbonOffsettingFundSet)
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
		it.Event = new(EpochRewardsCarbonOffsettingFundSet)
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
func (it *EpochRewardsCarbonOffsettingFundSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsCarbonOffsettingFundSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsCarbonOffsettingFundSet represents a CarbonOffsettingFundSet event raised by the EpochRewards contract.
type EpochRewardsCarbonOffsettingFundSet struct {
	Partner  common.Address
	Fraction *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCarbonOffsettingFundSet is a free log retrieval operation binding the contract event 0xe296227209b47bb8f4a76768ebd564dcde1c44be325a5d262f27c1fd4fd4538b.
//
// Solidity: event CarbonOffsettingFundSet(address indexed partner, uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) FilterCarbonOffsettingFundSet(opts *bind.FilterOpts, partner []common.Address) (*EpochRewardsCarbonOffsettingFundSetIterator, error) {

	var partnerRule []interface{}
	for _, partnerItem := range partner {
		partnerRule = append(partnerRule, partnerItem)
	}

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "CarbonOffsettingFundSet", partnerRule)
	if err != nil {
		return nil, err
	}
	return &EpochRewardsCarbonOffsettingFundSetIterator{contract: _EpochRewards.contract, event: "CarbonOffsettingFundSet", logs: logs, sub: sub}, nil
}

// WatchCarbonOffsettingFundSet is a free log subscription operation binding the contract event 0xe296227209b47bb8f4a76768ebd564dcde1c44be325a5d262f27c1fd4fd4538b.
//
// Solidity: event CarbonOffsettingFundSet(address indexed partner, uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) WatchCarbonOffsettingFundSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsCarbonOffsettingFundSet, partner []common.Address) (event.Subscription, error) {

	var partnerRule []interface{}
	for _, partnerItem := range partner {
		partnerRule = append(partnerRule, partnerItem)
	}

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "CarbonOffsettingFundSet", partnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsCarbonOffsettingFundSet)
				if err := _EpochRewards.contract.UnpackLog(event, "CarbonOffsettingFundSet", log); err != nil {
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

// ParseCarbonOffsettingFundSet is a log parse operation binding the contract event 0xe296227209b47bb8f4a76768ebd564dcde1c44be325a5d262f27c1fd4fd4538b.
//
// Solidity: event CarbonOffsettingFundSet(address indexed partner, uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) ParseCarbonOffsettingFundSet(log types.Log) (*EpochRewardsCarbonOffsettingFundSet, error) {
	event := new(EpochRewardsCarbonOffsettingFundSet)
	if err := _EpochRewards.contract.UnpackLog(event, "CarbonOffsettingFundSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsCommunityRewardFractionSetIterator is returned from FilterCommunityRewardFractionSet and is used to iterate over the raw logs and unpacked data for CommunityRewardFractionSet events raised by the EpochRewards contract.
type EpochRewardsCommunityRewardFractionSetIterator struct {
	Event *EpochRewardsCommunityRewardFractionSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsCommunityRewardFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsCommunityRewardFractionSet)
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
		it.Event = new(EpochRewardsCommunityRewardFractionSet)
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
func (it *EpochRewardsCommunityRewardFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsCommunityRewardFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsCommunityRewardFractionSet represents a CommunityRewardFractionSet event raised by the EpochRewards contract.
type EpochRewardsCommunityRewardFractionSet struct {
	Fraction *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCommunityRewardFractionSet is a free log retrieval operation binding the contract event 0xe6c1b64ad7e601924731051286b7b408b1a6f3ae96dcd6d2d9cd82578372ef9e.
//
// Solidity: event CommunityRewardFractionSet(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) FilterCommunityRewardFractionSet(opts *bind.FilterOpts) (*EpochRewardsCommunityRewardFractionSetIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "CommunityRewardFractionSet")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsCommunityRewardFractionSetIterator{contract: _EpochRewards.contract, event: "CommunityRewardFractionSet", logs: logs, sub: sub}, nil
}

// WatchCommunityRewardFractionSet is a free log subscription operation binding the contract event 0xe6c1b64ad7e601924731051286b7b408b1a6f3ae96dcd6d2d9cd82578372ef9e.
//
// Solidity: event CommunityRewardFractionSet(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) WatchCommunityRewardFractionSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsCommunityRewardFractionSet) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "CommunityRewardFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsCommunityRewardFractionSet)
				if err := _EpochRewards.contract.UnpackLog(event, "CommunityRewardFractionSet", log); err != nil {
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

// ParseCommunityRewardFractionSet is a log parse operation binding the contract event 0xe6c1b64ad7e601924731051286b7b408b1a6f3ae96dcd6d2d9cd82578372ef9e.
//
// Solidity: event CommunityRewardFractionSet(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) ParseCommunityRewardFractionSet(log types.Log) (*EpochRewardsCommunityRewardFractionSet, error) {
	event := new(EpochRewardsCommunityRewardFractionSet)
	if err := _EpochRewards.contract.UnpackLog(event, "CommunityRewardFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the EpochRewards contract.
type EpochRewardsOwnershipTransferredIterator struct {
	Event *EpochRewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *EpochRewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsOwnershipTransferred)
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
		it.Event = new(EpochRewardsOwnershipTransferred)
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
func (it *EpochRewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsOwnershipTransferred represents a OwnershipTransferred event raised by the EpochRewards contract.
type EpochRewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EpochRewards *EpochRewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*EpochRewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &EpochRewardsOwnershipTransferredIterator{contract: _EpochRewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_EpochRewards *EpochRewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *EpochRewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsOwnershipTransferred)
				if err := _EpochRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_EpochRewards *EpochRewardsFilterer) ParseOwnershipTransferred(log types.Log) (*EpochRewardsOwnershipTransferred, error) {
	event := new(EpochRewardsOwnershipTransferred)
	if err := _EpochRewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the EpochRewards contract.
type EpochRewardsRegistrySetIterator struct {
	Event *EpochRewardsRegistrySet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsRegistrySet)
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
		it.Event = new(EpochRewardsRegistrySet)
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
func (it *EpochRewardsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsRegistrySet represents a RegistrySet event raised by the EpochRewards contract.
type EpochRewardsRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_EpochRewards *EpochRewardsFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*EpochRewardsRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &EpochRewardsRegistrySetIterator{contract: _EpochRewards.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_EpochRewards *EpochRewardsFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *EpochRewardsRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsRegistrySet)
				if err := _EpochRewards.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_EpochRewards *EpochRewardsFilterer) ParseRegistrySet(log types.Log) (*EpochRewardsRegistrySet, error) {
	event := new(EpochRewardsRegistrySet)
	if err := _EpochRewards.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsRewardsMultiplierParametersSetIterator is returned from FilterRewardsMultiplierParametersSet and is used to iterate over the raw logs and unpacked data for RewardsMultiplierParametersSet events raised by the EpochRewards contract.
type EpochRewardsRewardsMultiplierParametersSetIterator struct {
	Event *EpochRewardsRewardsMultiplierParametersSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsRewardsMultiplierParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsRewardsMultiplierParametersSet)
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
		it.Event = new(EpochRewardsRewardsMultiplierParametersSet)
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
func (it *EpochRewardsRewardsMultiplierParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsRewardsMultiplierParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsRewardsMultiplierParametersSet represents a RewardsMultiplierParametersSet event raised by the EpochRewards contract.
type EpochRewardsRewardsMultiplierParametersSet struct {
	Max                        *big.Int
	UnderspendAdjustmentFactor *big.Int
	OverspendAdjustmentFactor  *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterRewardsMultiplierParametersSet is a free log retrieval operation binding the contract event 0x191445ee0115396c9725b9c642b985d63820ca57d54e42e5eb38faec4022f05d.
//
// Solidity: event RewardsMultiplierParametersSet(uint256 max, uint256 underspendAdjustmentFactor, uint256 overspendAdjustmentFactor)
func (_EpochRewards *EpochRewardsFilterer) FilterRewardsMultiplierParametersSet(opts *bind.FilterOpts) (*EpochRewardsRewardsMultiplierParametersSetIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "RewardsMultiplierParametersSet")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsRewardsMultiplierParametersSetIterator{contract: _EpochRewards.contract, event: "RewardsMultiplierParametersSet", logs: logs, sub: sub}, nil
}

// WatchRewardsMultiplierParametersSet is a free log subscription operation binding the contract event 0x191445ee0115396c9725b9c642b985d63820ca57d54e42e5eb38faec4022f05d.
//
// Solidity: event RewardsMultiplierParametersSet(uint256 max, uint256 underspendAdjustmentFactor, uint256 overspendAdjustmentFactor)
func (_EpochRewards *EpochRewardsFilterer) WatchRewardsMultiplierParametersSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsRewardsMultiplierParametersSet) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "RewardsMultiplierParametersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsRewardsMultiplierParametersSet)
				if err := _EpochRewards.contract.UnpackLog(event, "RewardsMultiplierParametersSet", log); err != nil {
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

// ParseRewardsMultiplierParametersSet is a log parse operation binding the contract event 0x191445ee0115396c9725b9c642b985d63820ca57d54e42e5eb38faec4022f05d.
//
// Solidity: event RewardsMultiplierParametersSet(uint256 max, uint256 underspendAdjustmentFactor, uint256 overspendAdjustmentFactor)
func (_EpochRewards *EpochRewardsFilterer) ParseRewardsMultiplierParametersSet(log types.Log) (*EpochRewardsRewardsMultiplierParametersSet, error) {
	event := new(EpochRewardsRewardsMultiplierParametersSet)
	if err := _EpochRewards.contract.UnpackLog(event, "RewardsMultiplierParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsTargetValidatorEpochPaymentSetIterator is returned from FilterTargetValidatorEpochPaymentSet and is used to iterate over the raw logs and unpacked data for TargetValidatorEpochPaymentSet events raised by the EpochRewards contract.
type EpochRewardsTargetValidatorEpochPaymentSetIterator struct {
	Event *EpochRewardsTargetValidatorEpochPaymentSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsTargetValidatorEpochPaymentSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsTargetValidatorEpochPaymentSet)
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
		it.Event = new(EpochRewardsTargetValidatorEpochPaymentSet)
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
func (it *EpochRewardsTargetValidatorEpochPaymentSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsTargetValidatorEpochPaymentSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsTargetValidatorEpochPaymentSet represents a TargetValidatorEpochPaymentSet event raised by the EpochRewards contract.
type EpochRewardsTargetValidatorEpochPaymentSet struct {
	Payment *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTargetValidatorEpochPaymentSet is a free log retrieval operation binding the contract event 0xa21d141963bb2c1064b5376f762d22d3e9c4c51617edcf105bcec0f14e36800c.
//
// Solidity: event TargetValidatorEpochPaymentSet(uint256 payment)
func (_EpochRewards *EpochRewardsFilterer) FilterTargetValidatorEpochPaymentSet(opts *bind.FilterOpts) (*EpochRewardsTargetValidatorEpochPaymentSetIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "TargetValidatorEpochPaymentSet")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsTargetValidatorEpochPaymentSetIterator{contract: _EpochRewards.contract, event: "TargetValidatorEpochPaymentSet", logs: logs, sub: sub}, nil
}

// WatchTargetValidatorEpochPaymentSet is a free log subscription operation binding the contract event 0xa21d141963bb2c1064b5376f762d22d3e9c4c51617edcf105bcec0f14e36800c.
//
// Solidity: event TargetValidatorEpochPaymentSet(uint256 payment)
func (_EpochRewards *EpochRewardsFilterer) WatchTargetValidatorEpochPaymentSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsTargetValidatorEpochPaymentSet) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "TargetValidatorEpochPaymentSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsTargetValidatorEpochPaymentSet)
				if err := _EpochRewards.contract.UnpackLog(event, "TargetValidatorEpochPaymentSet", log); err != nil {
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

// ParseTargetValidatorEpochPaymentSet is a log parse operation binding the contract event 0xa21d141963bb2c1064b5376f762d22d3e9c4c51617edcf105bcec0f14e36800c.
//
// Solidity: event TargetValidatorEpochPaymentSet(uint256 payment)
func (_EpochRewards *EpochRewardsFilterer) ParseTargetValidatorEpochPaymentSet(log types.Log) (*EpochRewardsTargetValidatorEpochPaymentSet, error) {
	event := new(EpochRewardsTargetValidatorEpochPaymentSet)
	if err := _EpochRewards.contract.UnpackLog(event, "TargetValidatorEpochPaymentSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsTargetVotingGoldFractionSetIterator is returned from FilterTargetVotingGoldFractionSet and is used to iterate over the raw logs and unpacked data for TargetVotingGoldFractionSet events raised by the EpochRewards contract.
type EpochRewardsTargetVotingGoldFractionSetIterator struct {
	Event *EpochRewardsTargetVotingGoldFractionSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsTargetVotingGoldFractionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsTargetVotingGoldFractionSet)
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
		it.Event = new(EpochRewardsTargetVotingGoldFractionSet)
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
func (it *EpochRewardsTargetVotingGoldFractionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsTargetVotingGoldFractionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsTargetVotingGoldFractionSet represents a TargetVotingGoldFractionSet event raised by the EpochRewards contract.
type EpochRewardsTargetVotingGoldFractionSet struct {
	Fraction *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTargetVotingGoldFractionSet is a free log retrieval operation binding the contract event 0xbae2f33c70949fbc7325c98655f3039e5e1c7f774874c99fd4f31ec5f432b159.
//
// Solidity: event TargetVotingGoldFractionSet(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) FilterTargetVotingGoldFractionSet(opts *bind.FilterOpts) (*EpochRewardsTargetVotingGoldFractionSetIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "TargetVotingGoldFractionSet")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsTargetVotingGoldFractionSetIterator{contract: _EpochRewards.contract, event: "TargetVotingGoldFractionSet", logs: logs, sub: sub}, nil
}

// WatchTargetVotingGoldFractionSet is a free log subscription operation binding the contract event 0xbae2f33c70949fbc7325c98655f3039e5e1c7f774874c99fd4f31ec5f432b159.
//
// Solidity: event TargetVotingGoldFractionSet(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) WatchTargetVotingGoldFractionSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsTargetVotingGoldFractionSet) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "TargetVotingGoldFractionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsTargetVotingGoldFractionSet)
				if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingGoldFractionSet", log); err != nil {
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

// ParseTargetVotingGoldFractionSet is a log parse operation binding the contract event 0xbae2f33c70949fbc7325c98655f3039e5e1c7f774874c99fd4f31ec5f432b159.
//
// Solidity: event TargetVotingGoldFractionSet(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) ParseTargetVotingGoldFractionSet(log types.Log) (*EpochRewardsTargetVotingGoldFractionSet, error) {
	event := new(EpochRewardsTargetVotingGoldFractionSet)
	if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingGoldFractionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsTargetVotingYieldParametersSetIterator is returned from FilterTargetVotingYieldParametersSet and is used to iterate over the raw logs and unpacked data for TargetVotingYieldParametersSet events raised by the EpochRewards contract.
type EpochRewardsTargetVotingYieldParametersSetIterator struct {
	Event *EpochRewardsTargetVotingYieldParametersSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsTargetVotingYieldParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsTargetVotingYieldParametersSet)
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
		it.Event = new(EpochRewardsTargetVotingYieldParametersSet)
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
func (it *EpochRewardsTargetVotingYieldParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsTargetVotingYieldParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsTargetVotingYieldParametersSet represents a TargetVotingYieldParametersSet event raised by the EpochRewards contract.
type EpochRewardsTargetVotingYieldParametersSet struct {
	Max              *big.Int
	AdjustmentFactor *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterTargetVotingYieldParametersSet is a free log retrieval operation binding the contract event 0x1b76e38f3fdd1f284ed4d47c9d50ff407748c516ff9761616ff638c233107625.
//
// Solidity: event TargetVotingYieldParametersSet(uint256 max, uint256 adjustmentFactor)
func (_EpochRewards *EpochRewardsFilterer) FilterTargetVotingYieldParametersSet(opts *bind.FilterOpts) (*EpochRewardsTargetVotingYieldParametersSetIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "TargetVotingYieldParametersSet")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsTargetVotingYieldParametersSetIterator{contract: _EpochRewards.contract, event: "TargetVotingYieldParametersSet", logs: logs, sub: sub}, nil
}

// WatchTargetVotingYieldParametersSet is a free log subscription operation binding the contract event 0x1b76e38f3fdd1f284ed4d47c9d50ff407748c516ff9761616ff638c233107625.
//
// Solidity: event TargetVotingYieldParametersSet(uint256 max, uint256 adjustmentFactor)
func (_EpochRewards *EpochRewardsFilterer) WatchTargetVotingYieldParametersSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsTargetVotingYieldParametersSet) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "TargetVotingYieldParametersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsTargetVotingYieldParametersSet)
				if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingYieldParametersSet", log); err != nil {
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

// ParseTargetVotingYieldParametersSet is a log parse operation binding the contract event 0x1b76e38f3fdd1f284ed4d47c9d50ff407748c516ff9761616ff638c233107625.
//
// Solidity: event TargetVotingYieldParametersSet(uint256 max, uint256 adjustmentFactor)
func (_EpochRewards *EpochRewardsFilterer) ParseTargetVotingYieldParametersSet(log types.Log) (*EpochRewardsTargetVotingYieldParametersSet, error) {
	event := new(EpochRewardsTargetVotingYieldParametersSet)
	if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingYieldParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsTargetVotingYieldSetIterator is returned from FilterTargetVotingYieldSet and is used to iterate over the raw logs and unpacked data for TargetVotingYieldSet events raised by the EpochRewards contract.
type EpochRewardsTargetVotingYieldSetIterator struct {
	Event *EpochRewardsTargetVotingYieldSet // Event containing the contract specifics and raw log

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
func (it *EpochRewardsTargetVotingYieldSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsTargetVotingYieldSet)
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
		it.Event = new(EpochRewardsTargetVotingYieldSet)
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
func (it *EpochRewardsTargetVotingYieldSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsTargetVotingYieldSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsTargetVotingYieldSet represents a TargetVotingYieldSet event raised by the EpochRewards contract.
type EpochRewardsTargetVotingYieldSet struct {
	Target *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTargetVotingYieldSet is a free log retrieval operation binding the contract event 0x152c3fc1e1cd415804bc9ae15876b37e62d8909358b940e6f4847ca927f46637.
//
// Solidity: event TargetVotingYieldSet(uint256 target)
func (_EpochRewards *EpochRewardsFilterer) FilterTargetVotingYieldSet(opts *bind.FilterOpts) (*EpochRewardsTargetVotingYieldSetIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "TargetVotingYieldSet")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsTargetVotingYieldSetIterator{contract: _EpochRewards.contract, event: "TargetVotingYieldSet", logs: logs, sub: sub}, nil
}

// WatchTargetVotingYieldSet is a free log subscription operation binding the contract event 0x152c3fc1e1cd415804bc9ae15876b37e62d8909358b940e6f4847ca927f46637.
//
// Solidity: event TargetVotingYieldSet(uint256 target)
func (_EpochRewards *EpochRewardsFilterer) WatchTargetVotingYieldSet(opts *bind.WatchOpts, sink chan<- *EpochRewardsTargetVotingYieldSet) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "TargetVotingYieldSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsTargetVotingYieldSet)
				if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingYieldSet", log); err != nil {
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

// ParseTargetVotingYieldSet is a log parse operation binding the contract event 0x152c3fc1e1cd415804bc9ae15876b37e62d8909358b940e6f4847ca927f46637.
//
// Solidity: event TargetVotingYieldSet(uint256 target)
func (_EpochRewards *EpochRewardsFilterer) ParseTargetVotingYieldSet(log types.Log) (*EpochRewardsTargetVotingYieldSet, error) {
	event := new(EpochRewardsTargetVotingYieldSet)
	if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingYieldSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EpochRewardsTargetVotingYieldUpdatedIterator is returned from FilterTargetVotingYieldUpdated and is used to iterate over the raw logs and unpacked data for TargetVotingYieldUpdated events raised by the EpochRewards contract.
type EpochRewardsTargetVotingYieldUpdatedIterator struct {
	Event *EpochRewardsTargetVotingYieldUpdated // Event containing the contract specifics and raw log

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
func (it *EpochRewardsTargetVotingYieldUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EpochRewardsTargetVotingYieldUpdated)
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
		it.Event = new(EpochRewardsTargetVotingYieldUpdated)
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
func (it *EpochRewardsTargetVotingYieldUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EpochRewardsTargetVotingYieldUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EpochRewardsTargetVotingYieldUpdated represents a TargetVotingYieldUpdated event raised by the EpochRewards contract.
type EpochRewardsTargetVotingYieldUpdated struct {
	Fraction *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTargetVotingYieldUpdated is a free log retrieval operation binding the contract event 0x49d8cdfe05bae61517c234f65f4088454013bafe561115126a8fe0074dc7700e.
//
// Solidity: event TargetVotingYieldUpdated(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) FilterTargetVotingYieldUpdated(opts *bind.FilterOpts) (*EpochRewardsTargetVotingYieldUpdatedIterator, error) {

	logs, sub, err := _EpochRewards.contract.FilterLogs(opts, "TargetVotingYieldUpdated")
	if err != nil {
		return nil, err
	}
	return &EpochRewardsTargetVotingYieldUpdatedIterator{contract: _EpochRewards.contract, event: "TargetVotingYieldUpdated", logs: logs, sub: sub}, nil
}

// WatchTargetVotingYieldUpdated is a free log subscription operation binding the contract event 0x49d8cdfe05bae61517c234f65f4088454013bafe561115126a8fe0074dc7700e.
//
// Solidity: event TargetVotingYieldUpdated(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) WatchTargetVotingYieldUpdated(opts *bind.WatchOpts, sink chan<- *EpochRewardsTargetVotingYieldUpdated) (event.Subscription, error) {

	logs, sub, err := _EpochRewards.contract.WatchLogs(opts, "TargetVotingYieldUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EpochRewardsTargetVotingYieldUpdated)
				if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingYieldUpdated", log); err != nil {
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

// ParseTargetVotingYieldUpdated is a log parse operation binding the contract event 0x49d8cdfe05bae61517c234f65f4088454013bafe561115126a8fe0074dc7700e.
//
// Solidity: event TargetVotingYieldUpdated(uint256 fraction)
func (_EpochRewards *EpochRewardsFilterer) ParseTargetVotingYieldUpdated(log types.Log) (*EpochRewardsTargetVotingYieldUpdated, error) {
	event := new(EpochRewardsTargetVotingYieldUpdated)
	if err := _EpochRewards.contract.UnpackLog(event, "TargetVotingYieldUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
