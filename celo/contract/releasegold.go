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

// ReleaseGoldABI is the input ABI used to generate the binding from.
const ReleaseGoldABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"maxDistribution\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"refundAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canVote\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"beneficiary\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXPIRATION_TIME\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalWithdrawn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"revocationInfo\",\"outputs\":[{\"name\":\"revocable\",\"type\":\"bool\"},{\"name\":\"canExpire\",\"type\":\"bool\"},{\"name\":\"releasedBalanceAtRevoke\",\"type\":\"uint256\"},{\"name\":\"revokeTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidityProvisionMet\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canValidate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"releaseSchedule\",\"outputs\":[{\"name\":\"releaseStartTime\",\"type\":\"uint256\"},{\"name\":\"releaseCliff\",\"type\":\"uint256\"},{\"name\":\"numReleasePeriods\",\"type\":\"uint256\"},{\"name\":\"releasePeriod\",\"type\":\"uint256\"},{\"name\":\"amountReleasedPerPeriod\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"atAddress\",\"type\":\"address\"}],\"name\":\"ReleaseGoldInstanceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"revokeTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"releasedBalanceAtRevoke\",\"type\":\"uint256\"}],\"name\":\"ReleaseScheduleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"atAddress\",\"type\":\"address\"}],\"name\":\"ReleaseGoldInstanceDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"maxDistribution\",\"type\":\"uint256\"}],\"name\":\"DistributionLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"LiquidityProvisionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"canExpire\",\"type\":\"bool\"}],\"name\":\"CanExpireSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"BeneficiarySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"releaseStartTime\",\"type\":\"uint256\"},{\"name\":\"releaseCliffTime\",\"type\":\"uint256\"},{\"name\":\"numReleasePeriods\",\"type\":\"uint256\"},{\"name\":\"releasePeriod\",\"type\":\"uint256\"},{\"name\":\"amountReleasedPerPeriod\",\"type\":\"uint256\"},{\"name\":\"revocable\",\"type\":\"bool\"},{\"name\":\"_beneficiary\",\"type\":\"address\"},{\"name\":\"_releaseOwner\",\"type\":\"address\"},{\"name\":\"_refundAddress\",\"type\":\"address\"},{\"name\":\"subjectToLiquidityProvision\",\"type\":\"bool\"},{\"name\":\"initialDistributionRatio\",\"type\":\"uint256\"},{\"name\":\"_canValidate\",\"type\":\"bool\"},{\"name\":\"_canVote\",\"type\":\"bool\"},{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRevoked\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"setLiquidityProvision\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_canExpire\",\"type\":\"bool\"}],\"name\":\"setCanExpire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"distributionRatio\",\"type\":\"uint256\"}],\"name\":\"setMaxDistribution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newBeneficiary\",\"type\":\"address\"}],\"name\":\"setBeneficiary\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"refundAndFinalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"expire\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainingTotalBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainingUnlockedBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRemainingLockedBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentReleasedTotalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"lockGold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"unlockGold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"relockGold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"withdrawLockedGold\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeVoteSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeValidatorSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"authorizeValidatorSignerWithPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"authorizeValidatorSignerWithKeys\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"authorizeAttestationSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"},{\"name\":\"walletAddress\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"createAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"}],\"name\":\"setAccountName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"walletAddress\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"setAccountWalletAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dataEncryptionKey\",\"type\":\"bytes\"}],\"name\":\"setAccountDataEncryptionKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"metadataURL\",\"type\":\"string\"}],\"name\":\"setAccountMetadataURL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"lesser\",\"type\":\"address\"},{\"name\":\"greater\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revokeActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"lesser\",\"type\":\"address\"},{\"name\":\"greater\",\"type\":\"address\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revokePending\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ReleaseGold is an auto generated Go binding around an Ethereum contract.
type ReleaseGold struct {
	ReleaseGoldCaller     // Read-only binding to the contract
	ReleaseGoldTransactor // Write-only binding to the contract
	ReleaseGoldFilterer   // Log filterer for contract events
}

// ReleaseGoldCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReleaseGoldCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseGoldTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReleaseGoldTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseGoldFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReleaseGoldFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReleaseGoldSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReleaseGoldSession struct {
	Contract     *ReleaseGold      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReleaseGoldCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReleaseGoldCallerSession struct {
	Contract *ReleaseGoldCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ReleaseGoldTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReleaseGoldTransactorSession struct {
	Contract     *ReleaseGoldTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ReleaseGoldRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReleaseGoldRaw struct {
	Contract *ReleaseGold // Generic contract binding to access the raw methods on
}

// ReleaseGoldCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReleaseGoldCallerRaw struct {
	Contract *ReleaseGoldCaller // Generic read-only contract binding to access the raw methods on
}

// ReleaseGoldTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReleaseGoldTransactorRaw struct {
	Contract *ReleaseGoldTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReleaseGold creates a new instance of ReleaseGold, bound to a specific deployed contract.
func NewReleaseGold(address common.Address, backend bind.ContractBackend) (*ReleaseGold, error) {
	contract, err := bindReleaseGold(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReleaseGold{ReleaseGoldCaller: ReleaseGoldCaller{contract: contract}, ReleaseGoldTransactor: ReleaseGoldTransactor{contract: contract}, ReleaseGoldFilterer: ReleaseGoldFilterer{contract: contract}}, nil
}

// NewReleaseGoldCaller creates a new read-only instance of ReleaseGold, bound to a specific deployed contract.
func NewReleaseGoldCaller(address common.Address, caller bind.ContractCaller) (*ReleaseGoldCaller, error) {
	contract, err := bindReleaseGold(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldCaller{contract: contract}, nil
}

// NewReleaseGoldTransactor creates a new write-only instance of ReleaseGold, bound to a specific deployed contract.
func NewReleaseGoldTransactor(address common.Address, transactor bind.ContractTransactor) (*ReleaseGoldTransactor, error) {
	contract, err := bindReleaseGold(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldTransactor{contract: contract}, nil
}

// NewReleaseGoldFilterer creates a new log filterer instance of ReleaseGold, bound to a specific deployed contract.
func NewReleaseGoldFilterer(address common.Address, filterer bind.ContractFilterer) (*ReleaseGoldFilterer, error) {
	contract, err := bindReleaseGold(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldFilterer{contract: contract}, nil
}

// bindReleaseGold binds a generic wrapper to an already deployed contract.
func bindReleaseGold(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReleaseGoldABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseReleaseGoldABI parses the ABI
func ParseReleaseGoldABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ReleaseGoldABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseGold *ReleaseGoldRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReleaseGold.Contract.ReleaseGoldCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseGold *ReleaseGoldRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.Contract.ReleaseGoldTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseGold *ReleaseGoldRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseGold.Contract.ReleaseGoldTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReleaseGold *ReleaseGoldCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReleaseGold.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReleaseGold *ReleaseGoldTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReleaseGold *ReleaseGoldTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReleaseGold.Contract.contract.Transact(opts, method, params...)
}

// EXPIRATIONTIME is a free data retrieval call binding the contract method 0x4a5c7348.
//
// Solidity: function EXPIRATION_TIME() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) EXPIRATIONTIME(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "EXPIRATION_TIME")
	return *ret0, err
}

// EXPIRATIONTIME is a free data retrieval call binding the contract method 0x4a5c7348.
//
// Solidity: function EXPIRATION_TIME() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) EXPIRATIONTIME() (*big.Int, error) {
	return _ReleaseGold.Contract.EXPIRATIONTIME(&_ReleaseGold.CallOpts)
}

// EXPIRATIONTIME is a free data retrieval call binding the contract method 0x4a5c7348.
//
// Solidity: function EXPIRATION_TIME() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) EXPIRATIONTIME() (*big.Int, error) {
	return _ReleaseGold.Contract.EXPIRATIONTIME(&_ReleaseGold.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_ReleaseGold *ReleaseGoldCaller) Beneficiary(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "beneficiary")
	return *ret0, err
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_ReleaseGold *ReleaseGoldSession) Beneficiary() (common.Address, error) {
	return _ReleaseGold.Contract.Beneficiary(&_ReleaseGold.CallOpts)
}

// Beneficiary is a free data retrieval call binding the contract method 0x38af3eed.
//
// Solidity: function beneficiary() constant returns(address)
func (_ReleaseGold *ReleaseGoldCallerSession) Beneficiary() (common.Address, error) {
	return _ReleaseGold.Contract.Beneficiary(&_ReleaseGold.CallOpts)
}

// CanValidate is a free data retrieval call binding the contract method 0xd69c2463.
//
// Solidity: function canValidate() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCaller) CanValidate(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "canValidate")
	return *ret0, err
}

// CanValidate is a free data retrieval call binding the contract method 0xd69c2463.
//
// Solidity: function canValidate() constant returns(bool)
func (_ReleaseGold *ReleaseGoldSession) CanValidate() (bool, error) {
	return _ReleaseGold.Contract.CanValidate(&_ReleaseGold.CallOpts)
}

// CanValidate is a free data retrieval call binding the contract method 0xd69c2463.
//
// Solidity: function canValidate() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCallerSession) CanValidate() (bool, error) {
	return _ReleaseGold.Contract.CanValidate(&_ReleaseGold.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0x159e7064.
//
// Solidity: function canVote() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCaller) CanVote(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "canVote")
	return *ret0, err
}

// CanVote is a free data retrieval call binding the contract method 0x159e7064.
//
// Solidity: function canVote() constant returns(bool)
func (_ReleaseGold *ReleaseGoldSession) CanVote() (bool, error) {
	return _ReleaseGold.Contract.CanVote(&_ReleaseGold.CallOpts)
}

// CanVote is a free data retrieval call binding the contract method 0x159e7064.
//
// Solidity: function canVote() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCallerSession) CanVote() (bool, error) {
	return _ReleaseGold.Contract.CanVote(&_ReleaseGold.CallOpts)
}

// GetCurrentReleasedTotalAmount is a free data retrieval call binding the contract method 0x1689eac7.
//
// Solidity: function getCurrentReleasedTotalAmount() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) GetCurrentReleasedTotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "getCurrentReleasedTotalAmount")
	return *ret0, err
}

// GetCurrentReleasedTotalAmount is a free data retrieval call binding the contract method 0x1689eac7.
//
// Solidity: function getCurrentReleasedTotalAmount() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) GetCurrentReleasedTotalAmount() (*big.Int, error) {
	return _ReleaseGold.Contract.GetCurrentReleasedTotalAmount(&_ReleaseGold.CallOpts)
}

// GetCurrentReleasedTotalAmount is a free data retrieval call binding the contract method 0x1689eac7.
//
// Solidity: function getCurrentReleasedTotalAmount() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) GetCurrentReleasedTotalAmount() (*big.Int, error) {
	return _ReleaseGold.Contract.GetCurrentReleasedTotalAmount(&_ReleaseGold.CallOpts)
}

// GetRemainingLockedBalance is a free data retrieval call binding the contract method 0xcf23c33a.
//
// Solidity: function getRemainingLockedBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) GetRemainingLockedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "getRemainingLockedBalance")
	return *ret0, err
}

// GetRemainingLockedBalance is a free data retrieval call binding the contract method 0xcf23c33a.
//
// Solidity: function getRemainingLockedBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) GetRemainingLockedBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetRemainingLockedBalance(&_ReleaseGold.CallOpts)
}

// GetRemainingLockedBalance is a free data retrieval call binding the contract method 0xcf23c33a.
//
// Solidity: function getRemainingLockedBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) GetRemainingLockedBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetRemainingLockedBalance(&_ReleaseGold.CallOpts)
}

// GetRemainingTotalBalance is a free data retrieval call binding the contract method 0x178f42e9.
//
// Solidity: function getRemainingTotalBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) GetRemainingTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "getRemainingTotalBalance")
	return *ret0, err
}

// GetRemainingTotalBalance is a free data retrieval call binding the contract method 0x178f42e9.
//
// Solidity: function getRemainingTotalBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) GetRemainingTotalBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetRemainingTotalBalance(&_ReleaseGold.CallOpts)
}

// GetRemainingTotalBalance is a free data retrieval call binding the contract method 0x178f42e9.
//
// Solidity: function getRemainingTotalBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) GetRemainingTotalBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetRemainingTotalBalance(&_ReleaseGold.CallOpts)
}

// GetRemainingUnlockedBalance is a free data retrieval call binding the contract method 0xd0484677.
//
// Solidity: function getRemainingUnlockedBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) GetRemainingUnlockedBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "getRemainingUnlockedBalance")
	return *ret0, err
}

// GetRemainingUnlockedBalance is a free data retrieval call binding the contract method 0xd0484677.
//
// Solidity: function getRemainingUnlockedBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) GetRemainingUnlockedBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetRemainingUnlockedBalance(&_ReleaseGold.CallOpts)
}

// GetRemainingUnlockedBalance is a free data retrieval call binding the contract method 0xd0484677.
//
// Solidity: function getRemainingUnlockedBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) GetRemainingUnlockedBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetRemainingUnlockedBalance(&_ReleaseGold.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) GetTotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "getTotalBalance")
	return *ret0, err
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) GetTotalBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetTotalBalance(&_ReleaseGold.CallOpts)
}

// GetTotalBalance is a free data retrieval call binding the contract method 0x12b58349.
//
// Solidity: function getTotalBalance() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) GetTotalBalance() (*big.Int, error) {
	return _ReleaseGold.Contract.GetTotalBalance(&_ReleaseGold.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_ReleaseGold *ReleaseGoldSession) Initialized() (bool, error) {
	return _ReleaseGold.Contract.Initialized(&_ReleaseGold.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCallerSession) Initialized() (bool, error) {
	return _ReleaseGold.Contract.Initialized(&_ReleaseGold.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ReleaseGold *ReleaseGoldSession) IsOwner() (bool, error) {
	return _ReleaseGold.Contract.IsOwner(&_ReleaseGold.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCallerSession) IsOwner() (bool, error) {
	return _ReleaseGold.Contract.IsOwner(&_ReleaseGold.CallOpts)
}

// IsRevoked is a free data retrieval call binding the contract method 0x2bc9ed02.
//
// Solidity: function isRevoked() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCaller) IsRevoked(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "isRevoked")
	return *ret0, err
}

// IsRevoked is a free data retrieval call binding the contract method 0x2bc9ed02.
//
// Solidity: function isRevoked() constant returns(bool)
func (_ReleaseGold *ReleaseGoldSession) IsRevoked() (bool, error) {
	return _ReleaseGold.Contract.IsRevoked(&_ReleaseGold.CallOpts)
}

// IsRevoked is a free data retrieval call binding the contract method 0x2bc9ed02.
//
// Solidity: function isRevoked() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCallerSession) IsRevoked() (bool, error) {
	return _ReleaseGold.Contract.IsRevoked(&_ReleaseGold.CallOpts)
}

// LiquidityProvisionMet is a free data retrieval call binding the contract method 0xbeb28d7c.
//
// Solidity: function liquidityProvisionMet() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCaller) LiquidityProvisionMet(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "liquidityProvisionMet")
	return *ret0, err
}

// LiquidityProvisionMet is a free data retrieval call binding the contract method 0xbeb28d7c.
//
// Solidity: function liquidityProvisionMet() constant returns(bool)
func (_ReleaseGold *ReleaseGoldSession) LiquidityProvisionMet() (bool, error) {
	return _ReleaseGold.Contract.LiquidityProvisionMet(&_ReleaseGold.CallOpts)
}

// LiquidityProvisionMet is a free data retrieval call binding the contract method 0xbeb28d7c.
//
// Solidity: function liquidityProvisionMet() constant returns(bool)
func (_ReleaseGold *ReleaseGoldCallerSession) LiquidityProvisionMet() (bool, error) {
	return _ReleaseGold.Contract.LiquidityProvisionMet(&_ReleaseGold.CallOpts)
}

// MaxDistribution is a free data retrieval call binding the contract method 0x044e0ea2.
//
// Solidity: function maxDistribution() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) MaxDistribution(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "maxDistribution")
	return *ret0, err
}

// MaxDistribution is a free data retrieval call binding the contract method 0x044e0ea2.
//
// Solidity: function maxDistribution() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) MaxDistribution() (*big.Int, error) {
	return _ReleaseGold.Contract.MaxDistribution(&_ReleaseGold.CallOpts)
}

// MaxDistribution is a free data retrieval call binding the contract method 0x044e0ea2.
//
// Solidity: function maxDistribution() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) MaxDistribution() (*big.Int, error) {
	return _ReleaseGold.Contract.MaxDistribution(&_ReleaseGold.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReleaseGold *ReleaseGoldCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReleaseGold *ReleaseGoldSession) Owner() (common.Address, error) {
	return _ReleaseGold.Contract.Owner(&_ReleaseGold.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ReleaseGold *ReleaseGoldCallerSession) Owner() (common.Address, error) {
	return _ReleaseGold.Contract.Owner(&_ReleaseGold.CallOpts)
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() constant returns(address)
func (_ReleaseGold *ReleaseGoldCaller) RefundAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "refundAddress")
	return *ret0, err
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() constant returns(address)
func (_ReleaseGold *ReleaseGoldSession) RefundAddress() (common.Address, error) {
	return _ReleaseGold.Contract.RefundAddress(&_ReleaseGold.CallOpts)
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() constant returns(address)
func (_ReleaseGold *ReleaseGoldCallerSession) RefundAddress() (common.Address, error) {
	return _ReleaseGold.Contract.RefundAddress(&_ReleaseGold.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_ReleaseGold *ReleaseGoldCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_ReleaseGold *ReleaseGoldSession) Registry() (common.Address, error) {
	return _ReleaseGold.Contract.Registry(&_ReleaseGold.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_ReleaseGold *ReleaseGoldCallerSession) Registry() (common.Address, error) {
	return _ReleaseGold.Contract.Registry(&_ReleaseGold.CallOpts)
}

// ReleaseOwner is a free data retrieval call binding the contract method 0x83d0aae9.
//
// Solidity: function releaseOwner() constant returns(address)
func (_ReleaseGold *ReleaseGoldCaller) ReleaseOwner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "releaseOwner")
	return *ret0, err
}

// ReleaseOwner is a free data retrieval call binding the contract method 0x83d0aae9.
//
// Solidity: function releaseOwner() constant returns(address)
func (_ReleaseGold *ReleaseGoldSession) ReleaseOwner() (common.Address, error) {
	return _ReleaseGold.Contract.ReleaseOwner(&_ReleaseGold.CallOpts)
}

// ReleaseOwner is a free data retrieval call binding the contract method 0x83d0aae9.
//
// Solidity: function releaseOwner() constant returns(address)
func (_ReleaseGold *ReleaseGoldCallerSession) ReleaseOwner() (common.Address, error) {
	return _ReleaseGold.Contract.ReleaseOwner(&_ReleaseGold.CallOpts)
}

// ReleaseSchedule is a free data retrieval call binding the contract method 0xfc9f3ef0.
//
// Solidity: function releaseSchedule() constant returns(uint256 releaseStartTime, uint256 releaseCliff, uint256 numReleasePeriods, uint256 releasePeriod, uint256 amountReleasedPerPeriod)
func (_ReleaseGold *ReleaseGoldCaller) ReleaseSchedule(opts *bind.CallOpts) (struct {
	ReleaseStartTime        *big.Int
	ReleaseCliff            *big.Int
	NumReleasePeriods       *big.Int
	ReleasePeriod           *big.Int
	AmountReleasedPerPeriod *big.Int
}, error) {
	ret := new(struct {
		ReleaseStartTime        *big.Int
		ReleaseCliff            *big.Int
		NumReleasePeriods       *big.Int
		ReleasePeriod           *big.Int
		AmountReleasedPerPeriod *big.Int
	})
	out := ret
	err := _ReleaseGold.contract.Call(opts, out, "releaseSchedule")
	return *ret, err
}

// ReleaseSchedule is a free data retrieval call binding the contract method 0xfc9f3ef0.
//
// Solidity: function releaseSchedule() constant returns(uint256 releaseStartTime, uint256 releaseCliff, uint256 numReleasePeriods, uint256 releasePeriod, uint256 amountReleasedPerPeriod)
func (_ReleaseGold *ReleaseGoldSession) ReleaseSchedule() (struct {
	ReleaseStartTime        *big.Int
	ReleaseCliff            *big.Int
	NumReleasePeriods       *big.Int
	ReleasePeriod           *big.Int
	AmountReleasedPerPeriod *big.Int
}, error) {
	return _ReleaseGold.Contract.ReleaseSchedule(&_ReleaseGold.CallOpts)
}

// ReleaseSchedule is a free data retrieval call binding the contract method 0xfc9f3ef0.
//
// Solidity: function releaseSchedule() constant returns(uint256 releaseStartTime, uint256 releaseCliff, uint256 numReleasePeriods, uint256 releasePeriod, uint256 amountReleasedPerPeriod)
func (_ReleaseGold *ReleaseGoldCallerSession) ReleaseSchedule() (struct {
	ReleaseStartTime        *big.Int
	ReleaseCliff            *big.Int
	NumReleasePeriods       *big.Int
	ReleasePeriod           *big.Int
	AmountReleasedPerPeriod *big.Int
}, error) {
	return _ReleaseGold.Contract.ReleaseSchedule(&_ReleaseGold.CallOpts)
}

// RevocationInfo is a free data retrieval call binding the contract method 0x969e83af.
//
// Solidity: function revocationInfo() constant returns(bool revocable, bool canExpire, uint256 releasedBalanceAtRevoke, uint256 revokeTime)
func (_ReleaseGold *ReleaseGoldCaller) RevocationInfo(opts *bind.CallOpts) (struct {
	Revocable               bool
	CanExpire               bool
	ReleasedBalanceAtRevoke *big.Int
	RevokeTime              *big.Int
}, error) {
	ret := new(struct {
		Revocable               bool
		CanExpire               bool
		ReleasedBalanceAtRevoke *big.Int
		RevokeTime              *big.Int
	})
	out := ret
	err := _ReleaseGold.contract.Call(opts, out, "revocationInfo")
	return *ret, err
}

// RevocationInfo is a free data retrieval call binding the contract method 0x969e83af.
//
// Solidity: function revocationInfo() constant returns(bool revocable, bool canExpire, uint256 releasedBalanceAtRevoke, uint256 revokeTime)
func (_ReleaseGold *ReleaseGoldSession) RevocationInfo() (struct {
	Revocable               bool
	CanExpire               bool
	ReleasedBalanceAtRevoke *big.Int
	RevokeTime              *big.Int
}, error) {
	return _ReleaseGold.Contract.RevocationInfo(&_ReleaseGold.CallOpts)
}

// RevocationInfo is a free data retrieval call binding the contract method 0x969e83af.
//
// Solidity: function revocationInfo() constant returns(bool revocable, bool canExpire, uint256 releasedBalanceAtRevoke, uint256 revokeTime)
func (_ReleaseGold *ReleaseGoldCallerSession) RevocationInfo() (struct {
	Revocable               bool
	CanExpire               bool
	ReleasedBalanceAtRevoke *big.Int
	RevokeTime              *big.Int
}, error) {
	return _ReleaseGold.Contract.RevocationInfo(&_ReleaseGold.CallOpts)
}

// TotalWithdrawn is a free data retrieval call binding the contract method 0x4b319713.
//
// Solidity: function totalWithdrawn() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCaller) TotalWithdrawn(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ReleaseGold.contract.Call(opts, out, "totalWithdrawn")
	return *ret0, err
}

// TotalWithdrawn is a free data retrieval call binding the contract method 0x4b319713.
//
// Solidity: function totalWithdrawn() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldSession) TotalWithdrawn() (*big.Int, error) {
	return _ReleaseGold.Contract.TotalWithdrawn(&_ReleaseGold.CallOpts)
}

// TotalWithdrawn is a free data retrieval call binding the contract method 0x4b319713.
//
// Solidity: function totalWithdrawn() constant returns(uint256)
func (_ReleaseGold *ReleaseGoldCallerSession) TotalWithdrawn() (*big.Int, error) {
	return _ReleaseGold.Contract.TotalWithdrawn(&_ReleaseGold.CallOpts)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactor) AuthorizeAttestationSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "authorizeAttestationSigner", signer, v, r, s)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldSession) AuthorizeAttestationSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeAttestationSigner(&_ReleaseGold.TransactOpts, signer, v, r, s)
}

// AuthorizeAttestationSigner is a paid mutator transaction binding the contract method 0x76afa04c.
//
// Solidity: function authorizeAttestationSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) AuthorizeAttestationSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeAttestationSigner(&_ReleaseGold.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactor) AuthorizeValidatorSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "authorizeValidatorSigner", signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldSession) AuthorizeValidatorSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeValidatorSigner(&_ReleaseGold.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSigner is a paid mutator transaction binding the contract method 0xbaf7ef0f.
//
// Solidity: function authorizeValidatorSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) AuthorizeValidatorSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeValidatorSigner(&_ReleaseGold.TransactOpts, signer, v, r, s)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_ReleaseGold *ReleaseGoldTransactor) AuthorizeValidatorSignerWithKeys(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "authorizeValidatorSignerWithKeys", signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_ReleaseGold *ReleaseGoldSession) AuthorizeValidatorSignerWithKeys(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeValidatorSignerWithKeys(&_ReleaseGold.TransactOpts, signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithKeys is a paid mutator transaction binding the contract method 0x1465b923.
//
// Solidity: function authorizeValidatorSignerWithKeys(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) AuthorizeValidatorSignerWithKeys(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeValidatorSignerWithKeys(&_ReleaseGold.TransactOpts, signer, v, r, s, ecdsaPublicKey, blsPublicKey, blsPop)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_ReleaseGold *ReleaseGoldTransactor) AuthorizeValidatorSignerWithPublicKey(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "authorizeValidatorSignerWithPublicKey", signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_ReleaseGold *ReleaseGoldSession) AuthorizeValidatorSignerWithPublicKey(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeValidatorSignerWithPublicKey(&_ReleaseGold.TransactOpts, signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeValidatorSignerWithPublicKey is a paid mutator transaction binding the contract method 0x0fa750d2.
//
// Solidity: function authorizeValidatorSignerWithPublicKey(address signer, uint8 v, bytes32 r, bytes32 s, bytes ecdsaPublicKey) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) AuthorizeValidatorSignerWithPublicKey(signer common.Address, v uint8, r [32]byte, s [32]byte, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeValidatorSignerWithPublicKey(&_ReleaseGold.TransactOpts, signer, v, r, s, ecdsaPublicKey)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactor) AuthorizeVoteSigner(opts *bind.TransactOpts, signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "authorizeVoteSigner", signer, v, r, s)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldSession) AuthorizeVoteSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeVoteSigner(&_ReleaseGold.TransactOpts, signer, v, r, s)
}

// AuthorizeVoteSigner is a paid mutator transaction binding the contract method 0x4282ee6d.
//
// Solidity: function authorizeVoteSigner(address signer, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) AuthorizeVoteSigner(signer common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.AuthorizeVoteSigner(&_ReleaseGold.TransactOpts, signer, v, r, s)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns()
func (_ReleaseGold *ReleaseGoldTransactor) CreateAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "createAccount")
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns()
func (_ReleaseGold *ReleaseGoldSession) CreateAccount() (*types.Transaction, error) {
	return _ReleaseGold.Contract.CreateAccount(&_ReleaseGold.TransactOpts)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x9dca362f.
//
// Solidity: function createAccount() returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) CreateAccount() (*types.Transaction, error) {
	return _ReleaseGold.Contract.CreateAccount(&_ReleaseGold.TransactOpts)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_ReleaseGold *ReleaseGoldTransactor) Expire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "expire")
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_ReleaseGold *ReleaseGoldSession) Expire() (*types.Transaction, error) {
	return _ReleaseGold.Contract.Expire(&_ReleaseGold.TransactOpts)
}

// Expire is a paid mutator transaction binding the contract method 0x79599f96.
//
// Solidity: function expire() returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) Expire() (*types.Transaction, error) {
	return _ReleaseGold.Contract.Expire(&_ReleaseGold.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x064a2e68.
//
// Solidity: function initialize(uint256 releaseStartTime, uint256 releaseCliffTime, uint256 numReleasePeriods, uint256 releasePeriod, uint256 amountReleasedPerPeriod, bool revocable, address _beneficiary, address _releaseOwner, address _refundAddress, bool subjectToLiquidityProvision, uint256 initialDistributionRatio, bool _canValidate, bool _canVote, address registryAddress) returns()
func (_ReleaseGold *ReleaseGoldTransactor) Initialize(opts *bind.TransactOpts, releaseStartTime *big.Int, releaseCliffTime *big.Int, numReleasePeriods *big.Int, releasePeriod *big.Int, amountReleasedPerPeriod *big.Int, revocable bool, _beneficiary common.Address, _releaseOwner common.Address, _refundAddress common.Address, subjectToLiquidityProvision bool, initialDistributionRatio *big.Int, _canValidate bool, _canVote bool, registryAddress common.Address) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "initialize", releaseStartTime, releaseCliffTime, numReleasePeriods, releasePeriod, amountReleasedPerPeriod, revocable, _beneficiary, _releaseOwner, _refundAddress, subjectToLiquidityProvision, initialDistributionRatio, _canValidate, _canVote, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x064a2e68.
//
// Solidity: function initialize(uint256 releaseStartTime, uint256 releaseCliffTime, uint256 numReleasePeriods, uint256 releasePeriod, uint256 amountReleasedPerPeriod, bool revocable, address _beneficiary, address _releaseOwner, address _refundAddress, bool subjectToLiquidityProvision, uint256 initialDistributionRatio, bool _canValidate, bool _canVote, address registryAddress) returns()
func (_ReleaseGold *ReleaseGoldSession) Initialize(releaseStartTime *big.Int, releaseCliffTime *big.Int, numReleasePeriods *big.Int, releasePeriod *big.Int, amountReleasedPerPeriod *big.Int, revocable bool, _beneficiary common.Address, _releaseOwner common.Address, _refundAddress common.Address, subjectToLiquidityProvision bool, initialDistributionRatio *big.Int, _canValidate bool, _canVote bool, registryAddress common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.Initialize(&_ReleaseGold.TransactOpts, releaseStartTime, releaseCliffTime, numReleasePeriods, releasePeriod, amountReleasedPerPeriod, revocable, _beneficiary, _releaseOwner, _refundAddress, subjectToLiquidityProvision, initialDistributionRatio, _canValidate, _canVote, registryAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x064a2e68.
//
// Solidity: function initialize(uint256 releaseStartTime, uint256 releaseCliffTime, uint256 numReleasePeriods, uint256 releasePeriod, uint256 amountReleasedPerPeriod, bool revocable, address _beneficiary, address _releaseOwner, address _refundAddress, bool subjectToLiquidityProvision, uint256 initialDistributionRatio, bool _canValidate, bool _canVote, address registryAddress) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) Initialize(releaseStartTime *big.Int, releaseCliffTime *big.Int, numReleasePeriods *big.Int, releasePeriod *big.Int, amountReleasedPerPeriod *big.Int, revocable bool, _beneficiary common.Address, _releaseOwner common.Address, _refundAddress common.Address, subjectToLiquidityProvision bool, initialDistributionRatio *big.Int, _canValidate bool, _canVote bool, registryAddress common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.Initialize(&_ReleaseGold.TransactOpts, releaseStartTime, releaseCliffTime, numReleasePeriods, releasePeriod, amountReleasedPerPeriod, revocable, _beneficiary, _releaseOwner, _refundAddress, subjectToLiquidityProvision, initialDistributionRatio, _canValidate, _canVote, registryAddress)
}

// LockGold is a paid mutator transaction binding the contract method 0x58aa72b0.
//
// Solidity: function lockGold(uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactor) LockGold(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "lockGold", value)
}

// LockGold is a paid mutator transaction binding the contract method 0x58aa72b0.
//
// Solidity: function lockGold(uint256 value) returns()
func (_ReleaseGold *ReleaseGoldSession) LockGold(value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.LockGold(&_ReleaseGold.TransactOpts, value)
}

// LockGold is a paid mutator transaction binding the contract method 0x58aa72b0.
//
// Solidity: function lockGold(uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) LockGold(value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.LockGold(&_ReleaseGold.TransactOpts, value)
}

// RefundAndFinalize is a paid mutator transaction binding the contract method 0xde82d860.
//
// Solidity: function refundAndFinalize() returns()
func (_ReleaseGold *ReleaseGoldTransactor) RefundAndFinalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "refundAndFinalize")
}

// RefundAndFinalize is a paid mutator transaction binding the contract method 0xde82d860.
//
// Solidity: function refundAndFinalize() returns()
func (_ReleaseGold *ReleaseGoldSession) RefundAndFinalize() (*types.Transaction, error) {
	return _ReleaseGold.Contract.RefundAndFinalize(&_ReleaseGold.TransactOpts)
}

// RefundAndFinalize is a paid mutator transaction binding the contract method 0xde82d860.
//
// Solidity: function refundAndFinalize() returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) RefundAndFinalize() (*types.Transaction, error) {
	return _ReleaseGold.Contract.RefundAndFinalize(&_ReleaseGold.TransactOpts)
}

// RelockGold is a paid mutator transaction binding the contract method 0x0d7c6e6e.
//
// Solidity: function relockGold(uint256 index, uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactor) RelockGold(opts *bind.TransactOpts, index *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "relockGold", index, value)
}

// RelockGold is a paid mutator transaction binding the contract method 0x0d7c6e6e.
//
// Solidity: function relockGold(uint256 index, uint256 value) returns()
func (_ReleaseGold *ReleaseGoldSession) RelockGold(index *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.RelockGold(&_ReleaseGold.TransactOpts, index, value)
}

// RelockGold is a paid mutator transaction binding the contract method 0x0d7c6e6e.
//
// Solidity: function relockGold(uint256 index, uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) RelockGold(index *big.Int, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.RelockGold(&_ReleaseGold.TransactOpts, index, value)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReleaseGold *ReleaseGoldTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReleaseGold *ReleaseGoldSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReleaseGold.Contract.RenounceOwnership(&_ReleaseGold.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ReleaseGold.Contract.RenounceOwnership(&_ReleaseGold.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ReleaseGold *ReleaseGoldTransactor) Revoke(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "revoke")
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ReleaseGold *ReleaseGoldSession) Revoke() (*types.Transaction, error) {
	return _ReleaseGold.Contract.Revoke(&_ReleaseGold.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xb6549f75.
//
// Solidity: function revoke() returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) Revoke() (*types.Transaction, error) {
	return _ReleaseGold.Contract.Revoke(&_ReleaseGold.TransactOpts)
}

// RevokeActive is a paid mutator transaction binding the contract method 0x6e198475.
//
// Solidity: function revokeActive(address group, uint256 value, address lesser, address greater, uint256 index) returns()
func (_ReleaseGold *ReleaseGoldTransactor) RevokeActive(opts *bind.TransactOpts, group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "revokeActive", group, value, lesser, greater, index)
}

// RevokeActive is a paid mutator transaction binding the contract method 0x6e198475.
//
// Solidity: function revokeActive(address group, uint256 value, address lesser, address greater, uint256 index) returns()
func (_ReleaseGold *ReleaseGoldSession) RevokeActive(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.RevokeActive(&_ReleaseGold.TransactOpts, group, value, lesser, greater, index)
}

// RevokeActive is a paid mutator transaction binding the contract method 0x6e198475.
//
// Solidity: function revokeActive(address group, uint256 value, address lesser, address greater, uint256 index) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) RevokeActive(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.RevokeActive(&_ReleaseGold.TransactOpts, group, value, lesser, greater, index)
}

// RevokePending is a paid mutator transaction binding the contract method 0x9dfb6081.
//
// Solidity: function revokePending(address group, uint256 value, address lesser, address greater, uint256 index) returns()
func (_ReleaseGold *ReleaseGoldTransactor) RevokePending(opts *bind.TransactOpts, group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "revokePending", group, value, lesser, greater, index)
}

// RevokePending is a paid mutator transaction binding the contract method 0x9dfb6081.
//
// Solidity: function revokePending(address group, uint256 value, address lesser, address greater, uint256 index) returns()
func (_ReleaseGold *ReleaseGoldSession) RevokePending(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.RevokePending(&_ReleaseGold.TransactOpts, group, value, lesser, greater, index)
}

// RevokePending is a paid mutator transaction binding the contract method 0x9dfb6081.
//
// Solidity: function revokePending(address group, uint256 value, address lesser, address greater, uint256 index) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) RevokePending(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.RevokePending(&_ReleaseGold.TransactOpts, group, value, lesser, greater, index)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetAccount(opts *bind.TransactOpts, name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setAccount", name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldSession) SetAccount(name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccount(&_ReleaseGold.TransactOpts, name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccount is a paid mutator transaction binding the contract method 0x90b12b47.
//
// Solidity: function setAccount(string name, bytes dataEncryptionKey, address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetAccount(name string, dataEncryptionKey []byte, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccount(&_ReleaseGold.TransactOpts, name, dataEncryptionKey, walletAddress, v, r, s)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetAccountDataEncryptionKey(opts *bind.TransactOpts, dataEncryptionKey []byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setAccountDataEncryptionKey", dataEncryptionKey)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_ReleaseGold *ReleaseGoldSession) SetAccountDataEncryptionKey(dataEncryptionKey []byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountDataEncryptionKey(&_ReleaseGold.TransactOpts, dataEncryptionKey)
}

// SetAccountDataEncryptionKey is a paid mutator transaction binding the contract method 0x0fe7abab.
//
// Solidity: function setAccountDataEncryptionKey(bytes dataEncryptionKey) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetAccountDataEncryptionKey(dataEncryptionKey []byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountDataEncryptionKey(&_ReleaseGold.TransactOpts, dataEncryptionKey)
}

// SetAccountMetadataURL is a paid mutator transaction binding the contract method 0x7981b664.
//
// Solidity: function setAccountMetadataURL(string metadataURL) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetAccountMetadataURL(opts *bind.TransactOpts, metadataURL string) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setAccountMetadataURL", metadataURL)
}

// SetAccountMetadataURL is a paid mutator transaction binding the contract method 0x7981b664.
//
// Solidity: function setAccountMetadataURL(string metadataURL) returns()
func (_ReleaseGold *ReleaseGoldSession) SetAccountMetadataURL(metadataURL string) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountMetadataURL(&_ReleaseGold.TransactOpts, metadataURL)
}

// SetAccountMetadataURL is a paid mutator transaction binding the contract method 0x7981b664.
//
// Solidity: function setAccountMetadataURL(string metadataURL) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetAccountMetadataURL(metadataURL string) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountMetadataURL(&_ReleaseGold.TransactOpts, metadataURL)
}

// SetAccountName is a paid mutator transaction binding the contract method 0x2e95adad.
//
// Solidity: function setAccountName(string name) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetAccountName(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setAccountName", name)
}

// SetAccountName is a paid mutator transaction binding the contract method 0x2e95adad.
//
// Solidity: function setAccountName(string name) returns()
func (_ReleaseGold *ReleaseGoldSession) SetAccountName(name string) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountName(&_ReleaseGold.TransactOpts, name)
}

// SetAccountName is a paid mutator transaction binding the contract method 0x2e95adad.
//
// Solidity: function setAccountName(string name) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetAccountName(name string) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountName(&_ReleaseGold.TransactOpts, name)
}

// SetAccountWalletAddress is a paid mutator transaction binding the contract method 0xd15ef2e6.
//
// Solidity: function setAccountWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetAccountWalletAddress(opts *bind.TransactOpts, walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setAccountWalletAddress", walletAddress, v, r, s)
}

// SetAccountWalletAddress is a paid mutator transaction binding the contract method 0xd15ef2e6.
//
// Solidity: function setAccountWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldSession) SetAccountWalletAddress(walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountWalletAddress(&_ReleaseGold.TransactOpts, walletAddress, v, r, s)
}

// SetAccountWalletAddress is a paid mutator transaction binding the contract method 0xd15ef2e6.
//
// Solidity: function setAccountWalletAddress(address walletAddress, uint8 v, bytes32 r, bytes32 s) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetAccountWalletAddress(walletAddress common.Address, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetAccountWalletAddress(&_ReleaseGold.TransactOpts, walletAddress, v, r, s)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetBeneficiary(opts *bind.TransactOpts, newBeneficiary common.Address) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setBeneficiary", newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ReleaseGold *ReleaseGoldSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetBeneficiary(&_ReleaseGold.TransactOpts, newBeneficiary)
}

// SetBeneficiary is a paid mutator transaction binding the contract method 0x1c31f710.
//
// Solidity: function setBeneficiary(address newBeneficiary) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetBeneficiary(newBeneficiary common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetBeneficiary(&_ReleaseGold.TransactOpts, newBeneficiary)
}

// SetCanExpire is a paid mutator transaction binding the contract method 0x1b3f3ac6.
//
// Solidity: function setCanExpire(bool _canExpire) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetCanExpire(opts *bind.TransactOpts, _canExpire bool) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setCanExpire", _canExpire)
}

// SetCanExpire is a paid mutator transaction binding the contract method 0x1b3f3ac6.
//
// Solidity: function setCanExpire(bool _canExpire) returns()
func (_ReleaseGold *ReleaseGoldSession) SetCanExpire(_canExpire bool) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetCanExpire(&_ReleaseGold.TransactOpts, _canExpire)
}

// SetCanExpire is a paid mutator transaction binding the contract method 0x1b3f3ac6.
//
// Solidity: function setCanExpire(bool _canExpire) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetCanExpire(_canExpire bool) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetCanExpire(&_ReleaseGold.TransactOpts, _canExpire)
}

// SetLiquidityProvision is a paid mutator transaction binding the contract method 0x6db9b090.
//
// Solidity: function setLiquidityProvision() returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetLiquidityProvision(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setLiquidityProvision")
}

// SetLiquidityProvision is a paid mutator transaction binding the contract method 0x6db9b090.
//
// Solidity: function setLiquidityProvision() returns()
func (_ReleaseGold *ReleaseGoldSession) SetLiquidityProvision() (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetLiquidityProvision(&_ReleaseGold.TransactOpts)
}

// SetLiquidityProvision is a paid mutator transaction binding the contract method 0x6db9b090.
//
// Solidity: function setLiquidityProvision() returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetLiquidityProvision() (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetLiquidityProvision(&_ReleaseGold.TransactOpts)
}

// SetMaxDistribution is a paid mutator transaction binding the contract method 0xdf6f6630.
//
// Solidity: function setMaxDistribution(uint256 distributionRatio) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetMaxDistribution(opts *bind.TransactOpts, distributionRatio *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setMaxDistribution", distributionRatio)
}

// SetMaxDistribution is a paid mutator transaction binding the contract method 0xdf6f6630.
//
// Solidity: function setMaxDistribution(uint256 distributionRatio) returns()
func (_ReleaseGold *ReleaseGoldSession) SetMaxDistribution(distributionRatio *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetMaxDistribution(&_ReleaseGold.TransactOpts, distributionRatio)
}

// SetMaxDistribution is a paid mutator transaction binding the contract method 0xdf6f6630.
//
// Solidity: function setMaxDistribution(uint256 distributionRatio) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetMaxDistribution(distributionRatio *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetMaxDistribution(&_ReleaseGold.TransactOpts, distributionRatio)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_ReleaseGold *ReleaseGoldTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_ReleaseGold *ReleaseGoldSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetRegistry(&_ReleaseGold.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.SetRegistry(&_ReleaseGold.TransactOpts, registryAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_ReleaseGold *ReleaseGoldSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.Transfer(&_ReleaseGold.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.Transfer(&_ReleaseGold.TransactOpts, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReleaseGold *ReleaseGoldTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReleaseGold *ReleaseGoldSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.TransferOwnership(&_ReleaseGold.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ReleaseGold.Contract.TransferOwnership(&_ReleaseGold.TransactOpts, newOwner)
}

// UnlockGold is a paid mutator transaction binding the contract method 0x68c7d70f.
//
// Solidity: function unlockGold(uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactor) UnlockGold(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "unlockGold", value)
}

// UnlockGold is a paid mutator transaction binding the contract method 0x68c7d70f.
//
// Solidity: function unlockGold(uint256 value) returns()
func (_ReleaseGold *ReleaseGoldSession) UnlockGold(value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.UnlockGold(&_ReleaseGold.TransactOpts, value)
}

// UnlockGold is a paid mutator transaction binding the contract method 0x68c7d70f.
//
// Solidity: function unlockGold(uint256 value) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) UnlockGold(value *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.UnlockGold(&_ReleaseGold.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ReleaseGold *ReleaseGoldTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ReleaseGold *ReleaseGoldSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.Withdraw(&_ReleaseGold.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.Withdraw(&_ReleaseGold.TransactOpts, amount)
}

// WithdrawLockedGold is a paid mutator transaction binding the contract method 0x52d2f864.
//
// Solidity: function withdrawLockedGold(uint256 index) returns()
func (_ReleaseGold *ReleaseGoldTransactor) WithdrawLockedGold(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.contract.Transact(opts, "withdrawLockedGold", index)
}

// WithdrawLockedGold is a paid mutator transaction binding the contract method 0x52d2f864.
//
// Solidity: function withdrawLockedGold(uint256 index) returns()
func (_ReleaseGold *ReleaseGoldSession) WithdrawLockedGold(index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.WithdrawLockedGold(&_ReleaseGold.TransactOpts, index)
}

// WithdrawLockedGold is a paid mutator transaction binding the contract method 0x52d2f864.
//
// Solidity: function withdrawLockedGold(uint256 index) returns()
func (_ReleaseGold *ReleaseGoldTransactorSession) WithdrawLockedGold(index *big.Int) (*types.Transaction, error) {
	return _ReleaseGold.Contract.WithdrawLockedGold(&_ReleaseGold.TransactOpts, index)
}

// TryParseLog attempts to parse a log. Returns the parsed log, eventName and whether it was successful
func (_ReleaseGold *ReleaseGoldFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _ReleaseGold.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "BeneficiarySet":
		event, err = _ReleaseGold.ParseBeneficiarySet(log)
	case "CanExpireSet":
		event, err = _ReleaseGold.ParseCanExpireSet(log)
	case "DistributionLimitSet":
		event, err = _ReleaseGold.ParseDistributionLimitSet(log)
	case "LiquidityProvisionSet":
		event, err = _ReleaseGold.ParseLiquidityProvisionSet(log)
	case "OwnershipTransferred":
		event, err = _ReleaseGold.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _ReleaseGold.ParseRegistrySet(log)
	case "ReleaseGoldInstanceCreated":
		event, err = _ReleaseGold.ParseReleaseGoldInstanceCreated(log)
	case "ReleaseGoldInstanceDestroyed":
		event, err = _ReleaseGold.ParseReleaseGoldInstanceDestroyed(log)
	case "ReleaseScheduleRevoked":
		event, err = _ReleaseGold.ParseReleaseScheduleRevoked(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// ReleaseGoldBeneficiarySetIterator is returned from FilterBeneficiarySet and is used to iterate over the raw logs and unpacked data for BeneficiarySet events raised by the ReleaseGold contract.
type ReleaseGoldBeneficiarySetIterator struct {
	Event *ReleaseGoldBeneficiarySet // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldBeneficiarySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldBeneficiarySet)
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
		it.Event = new(ReleaseGoldBeneficiarySet)
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
func (it *ReleaseGoldBeneficiarySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldBeneficiarySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldBeneficiarySet represents a BeneficiarySet event raised by the ReleaseGold contract.
type ReleaseGoldBeneficiarySet struct {
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBeneficiarySet is a free log retrieval operation binding the contract event 0x04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a24.
//
// Solidity: event BeneficiarySet(address indexed beneficiary)
func (_ReleaseGold *ReleaseGoldFilterer) FilterBeneficiarySet(opts *bind.FilterOpts, beneficiary []common.Address) (*ReleaseGoldBeneficiarySetIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "BeneficiarySet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldBeneficiarySetIterator{contract: _ReleaseGold.contract, event: "BeneficiarySet", logs: logs, sub: sub}, nil
}

// WatchBeneficiarySet is a free log subscription operation binding the contract event 0x04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a24.
//
// Solidity: event BeneficiarySet(address indexed beneficiary)
func (_ReleaseGold *ReleaseGoldFilterer) WatchBeneficiarySet(opts *bind.WatchOpts, sink chan<- *ReleaseGoldBeneficiarySet, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "BeneficiarySet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldBeneficiarySet)
				if err := _ReleaseGold.contract.UnpackLog(event, "BeneficiarySet", log); err != nil {
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

// ParseBeneficiarySet is a log parse operation binding the contract event 0x04d55a8be181fb8d75b76f2d48aa0b2ee40f47e53d6e61763eeeec46feea8a24.
//
// Solidity: event BeneficiarySet(address indexed beneficiary)
func (_ReleaseGold *ReleaseGoldFilterer) ParseBeneficiarySet(log types.Log) (*ReleaseGoldBeneficiarySet, error) {
	event := new(ReleaseGoldBeneficiarySet)
	if err := _ReleaseGold.contract.UnpackLog(event, "BeneficiarySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldCanExpireSetIterator is returned from FilterCanExpireSet and is used to iterate over the raw logs and unpacked data for CanExpireSet events raised by the ReleaseGold contract.
type ReleaseGoldCanExpireSetIterator struct {
	Event *ReleaseGoldCanExpireSet // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldCanExpireSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldCanExpireSet)
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
		it.Event = new(ReleaseGoldCanExpireSet)
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
func (it *ReleaseGoldCanExpireSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldCanExpireSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldCanExpireSet represents a CanExpireSet event raised by the ReleaseGold contract.
type ReleaseGoldCanExpireSet struct {
	CanExpire bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanExpireSet is a free log retrieval operation binding the contract event 0xcc86d6c701b54780681be776baa19da4895cf2feaef7cb81b7b4b4b512949ce3.
//
// Solidity: event CanExpireSet(bool canExpire)
func (_ReleaseGold *ReleaseGoldFilterer) FilterCanExpireSet(opts *bind.FilterOpts) (*ReleaseGoldCanExpireSetIterator, error) {

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "CanExpireSet")
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldCanExpireSetIterator{contract: _ReleaseGold.contract, event: "CanExpireSet", logs: logs, sub: sub}, nil
}

// WatchCanExpireSet is a free log subscription operation binding the contract event 0xcc86d6c701b54780681be776baa19da4895cf2feaef7cb81b7b4b4b512949ce3.
//
// Solidity: event CanExpireSet(bool canExpire)
func (_ReleaseGold *ReleaseGoldFilterer) WatchCanExpireSet(opts *bind.WatchOpts, sink chan<- *ReleaseGoldCanExpireSet) (event.Subscription, error) {

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "CanExpireSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldCanExpireSet)
				if err := _ReleaseGold.contract.UnpackLog(event, "CanExpireSet", log); err != nil {
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

// ParseCanExpireSet is a log parse operation binding the contract event 0xcc86d6c701b54780681be776baa19da4895cf2feaef7cb81b7b4b4b512949ce3.
//
// Solidity: event CanExpireSet(bool canExpire)
func (_ReleaseGold *ReleaseGoldFilterer) ParseCanExpireSet(log types.Log) (*ReleaseGoldCanExpireSet, error) {
	event := new(ReleaseGoldCanExpireSet)
	if err := _ReleaseGold.contract.UnpackLog(event, "CanExpireSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldDistributionLimitSetIterator is returned from FilterDistributionLimitSet and is used to iterate over the raw logs and unpacked data for DistributionLimitSet events raised by the ReleaseGold contract.
type ReleaseGoldDistributionLimitSetIterator struct {
	Event *ReleaseGoldDistributionLimitSet // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldDistributionLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldDistributionLimitSet)
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
		it.Event = new(ReleaseGoldDistributionLimitSet)
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
func (it *ReleaseGoldDistributionLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldDistributionLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldDistributionLimitSet represents a DistributionLimitSet event raised by the ReleaseGold contract.
type ReleaseGoldDistributionLimitSet struct {
	Beneficiary     common.Address
	MaxDistribution *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributionLimitSet is a free log retrieval operation binding the contract event 0x4fe054e42e0978de058ed4f176a4af2412f9ea91f6953494733d32f4284432b4.
//
// Solidity: event DistributionLimitSet(address indexed beneficiary, uint256 maxDistribution)
func (_ReleaseGold *ReleaseGoldFilterer) FilterDistributionLimitSet(opts *bind.FilterOpts, beneficiary []common.Address) (*ReleaseGoldDistributionLimitSetIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "DistributionLimitSet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldDistributionLimitSetIterator{contract: _ReleaseGold.contract, event: "DistributionLimitSet", logs: logs, sub: sub}, nil
}

// WatchDistributionLimitSet is a free log subscription operation binding the contract event 0x4fe054e42e0978de058ed4f176a4af2412f9ea91f6953494733d32f4284432b4.
//
// Solidity: event DistributionLimitSet(address indexed beneficiary, uint256 maxDistribution)
func (_ReleaseGold *ReleaseGoldFilterer) WatchDistributionLimitSet(opts *bind.WatchOpts, sink chan<- *ReleaseGoldDistributionLimitSet, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "DistributionLimitSet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldDistributionLimitSet)
				if err := _ReleaseGold.contract.UnpackLog(event, "DistributionLimitSet", log); err != nil {
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

// ParseDistributionLimitSet is a log parse operation binding the contract event 0x4fe054e42e0978de058ed4f176a4af2412f9ea91f6953494733d32f4284432b4.
//
// Solidity: event DistributionLimitSet(address indexed beneficiary, uint256 maxDistribution)
func (_ReleaseGold *ReleaseGoldFilterer) ParseDistributionLimitSet(log types.Log) (*ReleaseGoldDistributionLimitSet, error) {
	event := new(ReleaseGoldDistributionLimitSet)
	if err := _ReleaseGold.contract.UnpackLog(event, "DistributionLimitSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldLiquidityProvisionSetIterator is returned from FilterLiquidityProvisionSet and is used to iterate over the raw logs and unpacked data for LiquidityProvisionSet events raised by the ReleaseGold contract.
type ReleaseGoldLiquidityProvisionSetIterator struct {
	Event *ReleaseGoldLiquidityProvisionSet // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldLiquidityProvisionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldLiquidityProvisionSet)
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
		it.Event = new(ReleaseGoldLiquidityProvisionSet)
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
func (it *ReleaseGoldLiquidityProvisionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldLiquidityProvisionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldLiquidityProvisionSet represents a LiquidityProvisionSet event raised by the ReleaseGold contract.
type ReleaseGoldLiquidityProvisionSet struct {
	Beneficiary common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLiquidityProvisionSet is a free log retrieval operation binding the contract event 0x261df30ebbe3cd6d1f12a020dddf136c254764125eb7cfd51ca12a692f95c343.
//
// Solidity: event LiquidityProvisionSet(address indexed beneficiary)
func (_ReleaseGold *ReleaseGoldFilterer) FilterLiquidityProvisionSet(opts *bind.FilterOpts, beneficiary []common.Address) (*ReleaseGoldLiquidityProvisionSetIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "LiquidityProvisionSet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldLiquidityProvisionSetIterator{contract: _ReleaseGold.contract, event: "LiquidityProvisionSet", logs: logs, sub: sub}, nil
}

// WatchLiquidityProvisionSet is a free log subscription operation binding the contract event 0x261df30ebbe3cd6d1f12a020dddf136c254764125eb7cfd51ca12a692f95c343.
//
// Solidity: event LiquidityProvisionSet(address indexed beneficiary)
func (_ReleaseGold *ReleaseGoldFilterer) WatchLiquidityProvisionSet(opts *bind.WatchOpts, sink chan<- *ReleaseGoldLiquidityProvisionSet, beneficiary []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "LiquidityProvisionSet", beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldLiquidityProvisionSet)
				if err := _ReleaseGold.contract.UnpackLog(event, "LiquidityProvisionSet", log); err != nil {
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

// ParseLiquidityProvisionSet is a log parse operation binding the contract event 0x261df30ebbe3cd6d1f12a020dddf136c254764125eb7cfd51ca12a692f95c343.
//
// Solidity: event LiquidityProvisionSet(address indexed beneficiary)
func (_ReleaseGold *ReleaseGoldFilterer) ParseLiquidityProvisionSet(log types.Log) (*ReleaseGoldLiquidityProvisionSet, error) {
	event := new(ReleaseGoldLiquidityProvisionSet)
	if err := _ReleaseGold.contract.UnpackLog(event, "LiquidityProvisionSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ReleaseGold contract.
type ReleaseGoldOwnershipTransferredIterator struct {
	Event *ReleaseGoldOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldOwnershipTransferred)
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
		it.Event = new(ReleaseGoldOwnershipTransferred)
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
func (it *ReleaseGoldOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldOwnershipTransferred represents a OwnershipTransferred event raised by the ReleaseGold contract.
type ReleaseGoldOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReleaseGold *ReleaseGoldFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ReleaseGoldOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldOwnershipTransferredIterator{contract: _ReleaseGold.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ReleaseGold *ReleaseGoldFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ReleaseGoldOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldOwnershipTransferred)
				if err := _ReleaseGold.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ReleaseGold *ReleaseGoldFilterer) ParseOwnershipTransferred(log types.Log) (*ReleaseGoldOwnershipTransferred, error) {
	event := new(ReleaseGoldOwnershipTransferred)
	if err := _ReleaseGold.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the ReleaseGold contract.
type ReleaseGoldRegistrySetIterator struct {
	Event *ReleaseGoldRegistrySet // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldRegistrySet)
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
		it.Event = new(ReleaseGoldRegistrySet)
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
func (it *ReleaseGoldRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldRegistrySet represents a RegistrySet event raised by the ReleaseGold contract.
type ReleaseGoldRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_ReleaseGold *ReleaseGoldFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*ReleaseGoldRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldRegistrySetIterator{contract: _ReleaseGold.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_ReleaseGold *ReleaseGoldFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *ReleaseGoldRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldRegistrySet)
				if err := _ReleaseGold.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_ReleaseGold *ReleaseGoldFilterer) ParseRegistrySet(log types.Log) (*ReleaseGoldRegistrySet, error) {
	event := new(ReleaseGoldRegistrySet)
	if err := _ReleaseGold.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldReleaseGoldInstanceCreatedIterator is returned from FilterReleaseGoldInstanceCreated and is used to iterate over the raw logs and unpacked data for ReleaseGoldInstanceCreated events raised by the ReleaseGold contract.
type ReleaseGoldReleaseGoldInstanceCreatedIterator struct {
	Event *ReleaseGoldReleaseGoldInstanceCreated // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldReleaseGoldInstanceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldReleaseGoldInstanceCreated)
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
		it.Event = new(ReleaseGoldReleaseGoldInstanceCreated)
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
func (it *ReleaseGoldReleaseGoldInstanceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldReleaseGoldInstanceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldReleaseGoldInstanceCreated represents a ReleaseGoldInstanceCreated event raised by the ReleaseGold contract.
type ReleaseGoldReleaseGoldInstanceCreated struct {
	Beneficiary common.Address
	AtAddress   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleaseGoldInstanceCreated is a free log retrieval operation binding the contract event 0x0dcf2df6e6ef000384771a1b2a8f9c2872b94e86df3afb0703d68d79a820eef5.
//
// Solidity: event ReleaseGoldInstanceCreated(address indexed beneficiary, address indexed atAddress)
func (_ReleaseGold *ReleaseGoldFilterer) FilterReleaseGoldInstanceCreated(opts *bind.FilterOpts, beneficiary []common.Address, atAddress []common.Address) (*ReleaseGoldReleaseGoldInstanceCreatedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var atAddressRule []interface{}
	for _, atAddressItem := range atAddress {
		atAddressRule = append(atAddressRule, atAddressItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "ReleaseGoldInstanceCreated", beneficiaryRule, atAddressRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldReleaseGoldInstanceCreatedIterator{contract: _ReleaseGold.contract, event: "ReleaseGoldInstanceCreated", logs: logs, sub: sub}, nil
}

// WatchReleaseGoldInstanceCreated is a free log subscription operation binding the contract event 0x0dcf2df6e6ef000384771a1b2a8f9c2872b94e86df3afb0703d68d79a820eef5.
//
// Solidity: event ReleaseGoldInstanceCreated(address indexed beneficiary, address indexed atAddress)
func (_ReleaseGold *ReleaseGoldFilterer) WatchReleaseGoldInstanceCreated(opts *bind.WatchOpts, sink chan<- *ReleaseGoldReleaseGoldInstanceCreated, beneficiary []common.Address, atAddress []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var atAddressRule []interface{}
	for _, atAddressItem := range atAddress {
		atAddressRule = append(atAddressRule, atAddressItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "ReleaseGoldInstanceCreated", beneficiaryRule, atAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldReleaseGoldInstanceCreated)
				if err := _ReleaseGold.contract.UnpackLog(event, "ReleaseGoldInstanceCreated", log); err != nil {
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

// ParseReleaseGoldInstanceCreated is a log parse operation binding the contract event 0x0dcf2df6e6ef000384771a1b2a8f9c2872b94e86df3afb0703d68d79a820eef5.
//
// Solidity: event ReleaseGoldInstanceCreated(address indexed beneficiary, address indexed atAddress)
func (_ReleaseGold *ReleaseGoldFilterer) ParseReleaseGoldInstanceCreated(log types.Log) (*ReleaseGoldReleaseGoldInstanceCreated, error) {
	event := new(ReleaseGoldReleaseGoldInstanceCreated)
	if err := _ReleaseGold.contract.UnpackLog(event, "ReleaseGoldInstanceCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldReleaseGoldInstanceDestroyedIterator is returned from FilterReleaseGoldInstanceDestroyed and is used to iterate over the raw logs and unpacked data for ReleaseGoldInstanceDestroyed events raised by the ReleaseGold contract.
type ReleaseGoldReleaseGoldInstanceDestroyedIterator struct {
	Event *ReleaseGoldReleaseGoldInstanceDestroyed // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldReleaseGoldInstanceDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldReleaseGoldInstanceDestroyed)
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
		it.Event = new(ReleaseGoldReleaseGoldInstanceDestroyed)
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
func (it *ReleaseGoldReleaseGoldInstanceDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldReleaseGoldInstanceDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldReleaseGoldInstanceDestroyed represents a ReleaseGoldInstanceDestroyed event raised by the ReleaseGold contract.
type ReleaseGoldReleaseGoldInstanceDestroyed struct {
	Beneficiary common.Address
	AtAddress   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterReleaseGoldInstanceDestroyed is a free log retrieval operation binding the contract event 0x6a394f457aa81997bcda400fd5c309a1349f3a3233527e5e6c495be7b63c5bb0.
//
// Solidity: event ReleaseGoldInstanceDestroyed(address indexed beneficiary, address indexed atAddress)
func (_ReleaseGold *ReleaseGoldFilterer) FilterReleaseGoldInstanceDestroyed(opts *bind.FilterOpts, beneficiary []common.Address, atAddress []common.Address) (*ReleaseGoldReleaseGoldInstanceDestroyedIterator, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var atAddressRule []interface{}
	for _, atAddressItem := range atAddress {
		atAddressRule = append(atAddressRule, atAddressItem)
	}

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "ReleaseGoldInstanceDestroyed", beneficiaryRule, atAddressRule)
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldReleaseGoldInstanceDestroyedIterator{contract: _ReleaseGold.contract, event: "ReleaseGoldInstanceDestroyed", logs: logs, sub: sub}, nil
}

// WatchReleaseGoldInstanceDestroyed is a free log subscription operation binding the contract event 0x6a394f457aa81997bcda400fd5c309a1349f3a3233527e5e6c495be7b63c5bb0.
//
// Solidity: event ReleaseGoldInstanceDestroyed(address indexed beneficiary, address indexed atAddress)
func (_ReleaseGold *ReleaseGoldFilterer) WatchReleaseGoldInstanceDestroyed(opts *bind.WatchOpts, sink chan<- *ReleaseGoldReleaseGoldInstanceDestroyed, beneficiary []common.Address, atAddress []common.Address) (event.Subscription, error) {

	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}
	var atAddressRule []interface{}
	for _, atAddressItem := range atAddress {
		atAddressRule = append(atAddressRule, atAddressItem)
	}

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "ReleaseGoldInstanceDestroyed", beneficiaryRule, atAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldReleaseGoldInstanceDestroyed)
				if err := _ReleaseGold.contract.UnpackLog(event, "ReleaseGoldInstanceDestroyed", log); err != nil {
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

// ParseReleaseGoldInstanceDestroyed is a log parse operation binding the contract event 0x6a394f457aa81997bcda400fd5c309a1349f3a3233527e5e6c495be7b63c5bb0.
//
// Solidity: event ReleaseGoldInstanceDestroyed(address indexed beneficiary, address indexed atAddress)
func (_ReleaseGold *ReleaseGoldFilterer) ParseReleaseGoldInstanceDestroyed(log types.Log) (*ReleaseGoldReleaseGoldInstanceDestroyed, error) {
	event := new(ReleaseGoldReleaseGoldInstanceDestroyed)
	if err := _ReleaseGold.contract.UnpackLog(event, "ReleaseGoldInstanceDestroyed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReleaseGoldReleaseScheduleRevokedIterator is returned from FilterReleaseScheduleRevoked and is used to iterate over the raw logs and unpacked data for ReleaseScheduleRevoked events raised by the ReleaseGold contract.
type ReleaseGoldReleaseScheduleRevokedIterator struct {
	Event *ReleaseGoldReleaseScheduleRevoked // Event containing the contract specifics and raw log

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
func (it *ReleaseGoldReleaseScheduleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ReleaseGoldReleaseScheduleRevoked)
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
		it.Event = new(ReleaseGoldReleaseScheduleRevoked)
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
func (it *ReleaseGoldReleaseScheduleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ReleaseGoldReleaseScheduleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ReleaseGoldReleaseScheduleRevoked represents a ReleaseScheduleRevoked event raised by the ReleaseGold contract.
type ReleaseGoldReleaseScheduleRevoked struct {
	RevokeTimestamp         *big.Int
	ReleasedBalanceAtRevoke *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReleaseScheduleRevoked is a free log retrieval operation binding the contract event 0xbe8cff0b1980fe3ca968bf08dddc227b238e8398c0fa79445471303c981bc4c2.
//
// Solidity: event ReleaseScheduleRevoked(uint256 revokeTimestamp, uint256 releasedBalanceAtRevoke)
func (_ReleaseGold *ReleaseGoldFilterer) FilterReleaseScheduleRevoked(opts *bind.FilterOpts) (*ReleaseGoldReleaseScheduleRevokedIterator, error) {

	logs, sub, err := _ReleaseGold.contract.FilterLogs(opts, "ReleaseScheduleRevoked")
	if err != nil {
		return nil, err
	}
	return &ReleaseGoldReleaseScheduleRevokedIterator{contract: _ReleaseGold.contract, event: "ReleaseScheduleRevoked", logs: logs, sub: sub}, nil
}

// WatchReleaseScheduleRevoked is a free log subscription operation binding the contract event 0xbe8cff0b1980fe3ca968bf08dddc227b238e8398c0fa79445471303c981bc4c2.
//
// Solidity: event ReleaseScheduleRevoked(uint256 revokeTimestamp, uint256 releasedBalanceAtRevoke)
func (_ReleaseGold *ReleaseGoldFilterer) WatchReleaseScheduleRevoked(opts *bind.WatchOpts, sink chan<- *ReleaseGoldReleaseScheduleRevoked) (event.Subscription, error) {

	logs, sub, err := _ReleaseGold.contract.WatchLogs(opts, "ReleaseScheduleRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ReleaseGoldReleaseScheduleRevoked)
				if err := _ReleaseGold.contract.UnpackLog(event, "ReleaseScheduleRevoked", log); err != nil {
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

// ParseReleaseScheduleRevoked is a log parse operation binding the contract event 0xbe8cff0b1980fe3ca968bf08dddc227b238e8398c0fa79445471303c981bc4c2.
//
// Solidity: event ReleaseScheduleRevoked(uint256 revokeTimestamp, uint256 releasedBalanceAtRevoke)
func (_ReleaseGold *ReleaseGoldFilterer) ParseReleaseScheduleRevoked(log types.Log) (*ReleaseGoldReleaseScheduleRevoked, error) {
	event := new(ReleaseGoldReleaseScheduleRevoked)
	if err := _ReleaseGold.contract.UnpackLog(event, "ReleaseScheduleRevoked", log); err != nil {
		return nil, err
	}
	return event, nil
}
