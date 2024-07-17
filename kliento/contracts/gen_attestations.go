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

// AttestationsMetaData contains all meta data concerning the Attestations contract.
var AttestationsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"AttestationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AttestationExpiryBlocksSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"attestationRequestFeeToken\",\"type\":\"address\"}],\"name\":\"AttestationIssuerSelected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"AttestationRequestFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attestationsRequested\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"attestationRequestFeeToken\",\"type\":\"address\"}],\"name\":\"AttestationsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toAccount\",\"type\":\"address\"}],\"name\":\"AttestationsTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"MaxAttestationsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SelectIssuersWaitBlocksSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"indentifier\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"TransferApproval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"attestationExpiryBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"attestationRequestFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxAttestations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"selectIssuersWaitBlocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transferApprovals\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_attestationExpiryBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_selectIssuersWaitBlocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxAttestations\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"attestationRequestFeeTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"attestationRequestFeeValues\",\"type\":\"uint256[]\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revoke\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUnselectedRequest\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAttestationIssuers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAttestationStats\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"identifiersToLookup\",\"type\":\"bytes32[]\"}],\"name\":\"batchGetAttestationStats\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"},{\"internalType\":\"uint64[]\",\"name\":\"\",\"type\":\"uint64[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"}],\"name\":\"getAttestationState\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCompletableAttestations\",\"outputs\":[{\"internalType\":\"uint32[]\",\"name\":\"\",\"type\":\"uint32[]\"},{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAttestationRequestFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setAttestationRequestFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_attestationExpiryBlocks\",\"type\":\"uint256\"}],\"name\":\"setAttestationExpiryBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_selectIssuersWaitBlocks\",\"type\":\"uint256\"}],\"name\":\"setSelectIssuersWaitBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxAttestations\",\"type\":\"uint256\"}],\"name\":\"setMaxAttestations\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxAttestations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"validateAttestationCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"}],\"name\":\"lookupAccountsForIdentifier\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"identifier\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"expected\",\"type\":\"uint32\"}],\"name\":\"requireNAttestationsRequested\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AttestationsABI is the input ABI used to generate the binding from.
// Deprecated: Use AttestationsMetaData.ABI instead.
var AttestationsABI = AttestationsMetaData.ABI

// Attestations is an auto generated Go binding around an Ethereum contract.
type Attestations struct {
	AttestationsCaller     // Read-only binding to the contract
	AttestationsTransactor // Write-only binding to the contract
	AttestationsFilterer   // Log filterer for contract events
}

// AttestationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AttestationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AttestationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AttestationsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AttestationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AttestationsSession struct {
	Contract     *Attestations     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AttestationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AttestationsCallerSession struct {
	Contract *AttestationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// AttestationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AttestationsTransactorSession struct {
	Contract     *AttestationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AttestationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AttestationsRaw struct {
	Contract *Attestations // Generic contract binding to access the raw methods on
}

// AttestationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AttestationsCallerRaw struct {
	Contract *AttestationsCaller // Generic read-only contract binding to access the raw methods on
}

// AttestationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AttestationsTransactorRaw struct {
	Contract *AttestationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAttestations creates a new instance of Attestations, bound to a specific deployed contract.
func NewAttestations(address common.Address, backend bind.ContractBackend) (*Attestations, error) {
	contract, err := bindAttestations(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Attestations{AttestationsCaller: AttestationsCaller{contract: contract}, AttestationsTransactor: AttestationsTransactor{contract: contract}, AttestationsFilterer: AttestationsFilterer{contract: contract}}, nil
}

// NewAttestationsCaller creates a new read-only instance of Attestations, bound to a specific deployed contract.
func NewAttestationsCaller(address common.Address, caller bind.ContractCaller) (*AttestationsCaller, error) {
	contract, err := bindAttestations(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationsCaller{contract: contract}, nil
}

// NewAttestationsTransactor creates a new write-only instance of Attestations, bound to a specific deployed contract.
func NewAttestationsTransactor(address common.Address, transactor bind.ContractTransactor) (*AttestationsTransactor, error) {
	contract, err := bindAttestations(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AttestationsTransactor{contract: contract}, nil
}

// NewAttestationsFilterer creates a new log filterer instance of Attestations, bound to a specific deployed contract.
func NewAttestationsFilterer(address common.Address, filterer bind.ContractFilterer) (*AttestationsFilterer, error) {
	contract, err := bindAttestations(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AttestationsFilterer{contract: contract}, nil
}

// bindAttestations binds a generic wrapper to an already deployed contract.
func bindAttestations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseAttestationsABI parses the ABI
func ParseAttestationsABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(AttestationsABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestations *AttestationsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestations.Contract.AttestationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestations *AttestationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestations.Contract.AttestationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestations *AttestationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestations.Contract.AttestationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Attestations *AttestationsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Attestations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Attestations *AttestationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Attestations *AttestationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Attestations.Contract.contract.Transact(opts, method, params...)
}

// AttestationExpiryBlocks is a free data retrieval call binding the contract method 0xb45eb7da.
//
// Solidity: function attestationExpiryBlocks() view returns(uint256)
func (_Attestations *AttestationsCaller) AttestationExpiryBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "attestationExpiryBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationExpiryBlocks is a free data retrieval call binding the contract method 0xb45eb7da.
//
// Solidity: function attestationExpiryBlocks() view returns(uint256)
func (_Attestations *AttestationsSession) AttestationExpiryBlocks() (*big.Int, error) {
	return _Attestations.Contract.AttestationExpiryBlocks(&_Attestations.CallOpts)
}

// AttestationExpiryBlocks is a free data retrieval call binding the contract method 0xb45eb7da.
//
// Solidity: function attestationExpiryBlocks() view returns(uint256)
func (_Attestations *AttestationsCallerSession) AttestationExpiryBlocks() (*big.Int, error) {
	return _Attestations.Contract.AttestationExpiryBlocks(&_Attestations.CallOpts)
}

// AttestationRequestFees is a free data retrieval call binding the contract method 0xbd93f998.
//
// Solidity: function attestationRequestFees(address ) view returns(uint256)
func (_Attestations *AttestationsCaller) AttestationRequestFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "attestationRequestFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AttestationRequestFees is a free data retrieval call binding the contract method 0xbd93f998.
//
// Solidity: function attestationRequestFees(address ) view returns(uint256)
func (_Attestations *AttestationsSession) AttestationRequestFees(arg0 common.Address) (*big.Int, error) {
	return _Attestations.Contract.AttestationRequestFees(&_Attestations.CallOpts, arg0)
}

// AttestationRequestFees is a free data retrieval call binding the contract method 0xbd93f998.
//
// Solidity: function attestationRequestFees(address ) view returns(uint256)
func (_Attestations *AttestationsCallerSession) AttestationRequestFees(arg0 common.Address) (*big.Int, error) {
	return _Attestations.Contract.AttestationRequestFees(&_Attestations.CallOpts, arg0)
}

// BatchGetAttestationStats is a free data retrieval call binding the contract method 0x96357c0a.
//
// Solidity: function batchGetAttestationStats(bytes32[] identifiersToLookup) view returns(uint256[], address[], uint64[], uint64[])
func (_Attestations *AttestationsCaller) BatchGetAttestationStats(opts *bind.CallOpts, identifiersToLookup [][32]byte) ([]*big.Int, []common.Address, []uint64, []uint64, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "batchGetAttestationStats", identifiersToLookup)

	if err != nil {
		return *new([]*big.Int), *new([]common.Address), *new([]uint64), *new([]uint64), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]uint64)).(*[]uint64)
	out3 := *abi.ConvertType(out[3], new([]uint64)).(*[]uint64)

	return out0, out1, out2, out3, err

}

// BatchGetAttestationStats is a free data retrieval call binding the contract method 0x96357c0a.
//
// Solidity: function batchGetAttestationStats(bytes32[] identifiersToLookup) view returns(uint256[], address[], uint64[], uint64[])
func (_Attestations *AttestationsSession) BatchGetAttestationStats(identifiersToLookup [][32]byte) ([]*big.Int, []common.Address, []uint64, []uint64, error) {
	return _Attestations.Contract.BatchGetAttestationStats(&_Attestations.CallOpts, identifiersToLookup)
}

// BatchGetAttestationStats is a free data retrieval call binding the contract method 0x96357c0a.
//
// Solidity: function batchGetAttestationStats(bytes32[] identifiersToLookup) view returns(uint256[], address[], uint64[], uint64[])
func (_Attestations *AttestationsCallerSession) BatchGetAttestationStats(identifiersToLookup [][32]byte) ([]*big.Int, []common.Address, []uint64, []uint64, error) {
	return _Attestations.Contract.BatchGetAttestationStats(&_Attestations.CallOpts, identifiersToLookup)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Attestations *AttestationsCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Attestations *AttestationsSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Attestations.Contract.CheckProofOfPossession(&_Attestations.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Attestations *AttestationsCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Attestations.Contract.CheckProofOfPossession(&_Attestations.CallOpts, sender, blsKey, blsPop)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Attestations *AttestationsCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

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
func (_Attestations *AttestationsSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Attestations.Contract.FractionMulExp(&_Attestations.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Attestations *AttestationsCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Attestations.Contract.FractionMulExp(&_Attestations.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetAttestationIssuers is a free data retrieval call binding the contract method 0x5fc5c916.
//
// Solidity: function getAttestationIssuers(bytes32 identifier, address account) view returns(address[])
func (_Attestations *AttestationsCaller) GetAttestationIssuers(opts *bind.CallOpts, identifier [32]byte, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getAttestationIssuers", identifier, account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAttestationIssuers is a free data retrieval call binding the contract method 0x5fc5c916.
//
// Solidity: function getAttestationIssuers(bytes32 identifier, address account) view returns(address[])
func (_Attestations *AttestationsSession) GetAttestationIssuers(identifier [32]byte, account common.Address) ([]common.Address, error) {
	return _Attestations.Contract.GetAttestationIssuers(&_Attestations.CallOpts, identifier, account)
}

// GetAttestationIssuers is a free data retrieval call binding the contract method 0x5fc5c916.
//
// Solidity: function getAttestationIssuers(bytes32 identifier, address account) view returns(address[])
func (_Attestations *AttestationsCallerSession) GetAttestationIssuers(identifier [32]byte, account common.Address) ([]common.Address, error) {
	return _Attestations.Contract.GetAttestationIssuers(&_Attestations.CallOpts, identifier, account)
}

// GetAttestationRequestFee is a free data retrieval call binding the contract method 0x623d5931.
//
// Solidity: function getAttestationRequestFee(address token) view returns(uint256)
func (_Attestations *AttestationsCaller) GetAttestationRequestFee(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getAttestationRequestFee", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAttestationRequestFee is a free data retrieval call binding the contract method 0x623d5931.
//
// Solidity: function getAttestationRequestFee(address token) view returns(uint256)
func (_Attestations *AttestationsSession) GetAttestationRequestFee(token common.Address) (*big.Int, error) {
	return _Attestations.Contract.GetAttestationRequestFee(&_Attestations.CallOpts, token)
}

// GetAttestationRequestFee is a free data retrieval call binding the contract method 0x623d5931.
//
// Solidity: function getAttestationRequestFee(address token) view returns(uint256)
func (_Attestations *AttestationsCallerSession) GetAttestationRequestFee(token common.Address) (*big.Int, error) {
	return _Attestations.Contract.GetAttestationRequestFee(&_Attestations.CallOpts, token)
}

// GetAttestationState is a free data retrieval call binding the contract method 0xb5599cc6.
//
// Solidity: function getAttestationState(bytes32 identifier, address account, address issuer) view returns(uint8, uint32, address)
func (_Attestations *AttestationsCaller) GetAttestationState(opts *bind.CallOpts, identifier [32]byte, account common.Address, issuer common.Address) (uint8, uint32, common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getAttestationState", identifier, account, issuer)

	if err != nil {
		return *new(uint8), *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetAttestationState is a free data retrieval call binding the contract method 0xb5599cc6.
//
// Solidity: function getAttestationState(bytes32 identifier, address account, address issuer) view returns(uint8, uint32, address)
func (_Attestations *AttestationsSession) GetAttestationState(identifier [32]byte, account common.Address, issuer common.Address) (uint8, uint32, common.Address, error) {
	return _Attestations.Contract.GetAttestationState(&_Attestations.CallOpts, identifier, account, issuer)
}

// GetAttestationState is a free data retrieval call binding the contract method 0xb5599cc6.
//
// Solidity: function getAttestationState(bytes32 identifier, address account, address issuer) view returns(uint8, uint32, address)
func (_Attestations *AttestationsCallerSession) GetAttestationState(identifier [32]byte, account common.Address, issuer common.Address) (uint8, uint32, common.Address, error) {
	return _Attestations.Contract.GetAttestationState(&_Attestations.CallOpts, identifier, account, issuer)
}

// GetAttestationStats is a free data retrieval call binding the contract method 0x596abea5.
//
// Solidity: function getAttestationStats(bytes32 identifier, address account) view returns(uint32, uint32)
func (_Attestations *AttestationsCaller) GetAttestationStats(opts *bind.CallOpts, identifier [32]byte, account common.Address) (uint32, uint32, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getAttestationStats", identifier, account)

	if err != nil {
		return *new(uint32), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetAttestationStats is a free data retrieval call binding the contract method 0x596abea5.
//
// Solidity: function getAttestationStats(bytes32 identifier, address account) view returns(uint32, uint32)
func (_Attestations *AttestationsSession) GetAttestationStats(identifier [32]byte, account common.Address) (uint32, uint32, error) {
	return _Attestations.Contract.GetAttestationStats(&_Attestations.CallOpts, identifier, account)
}

// GetAttestationStats is a free data retrieval call binding the contract method 0x596abea5.
//
// Solidity: function getAttestationStats(bytes32 identifier, address account) view returns(uint32, uint32)
func (_Attestations *AttestationsCallerSession) GetAttestationStats(identifier [32]byte, account common.Address) (uint32, uint32, error) {
	return _Attestations.Contract.GetAttestationStats(&_Attestations.CallOpts, identifier, account)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Attestations *AttestationsCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Attestations *AttestationsSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Attestations.Contract.GetBlockNumberFromHeader(&_Attestations.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Attestations *AttestationsCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Attestations.Contract.GetBlockNumberFromHeader(&_Attestations.CallOpts, header)
}

// GetCompletableAttestations is a free data retrieval call binding the contract method 0x4eef7e85.
//
// Solidity: function getCompletableAttestations(bytes32 identifier, address account) view returns(uint32[], address[], uint256[], bytes)
func (_Attestations *AttestationsCaller) GetCompletableAttestations(opts *bind.CallOpts, identifier [32]byte, account common.Address) ([]uint32, []common.Address, []*big.Int, []byte, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getCompletableAttestations", identifier, account)

	if err != nil {
		return *new([]uint32), *new([]common.Address), *new([]*big.Int), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint32)).(*[]uint32)
	out1 := *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	out2 := *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	out3 := *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return out0, out1, out2, out3, err

}

// GetCompletableAttestations is a free data retrieval call binding the contract method 0x4eef7e85.
//
// Solidity: function getCompletableAttestations(bytes32 identifier, address account) view returns(uint32[], address[], uint256[], bytes)
func (_Attestations *AttestationsSession) GetCompletableAttestations(identifier [32]byte, account common.Address) ([]uint32, []common.Address, []*big.Int, []byte, error) {
	return _Attestations.Contract.GetCompletableAttestations(&_Attestations.CallOpts, identifier, account)
}

// GetCompletableAttestations is a free data retrieval call binding the contract method 0x4eef7e85.
//
// Solidity: function getCompletableAttestations(bytes32 identifier, address account) view returns(uint32[], address[], uint256[], bytes)
func (_Attestations *AttestationsCallerSession) GetCompletableAttestations(identifier [32]byte, account common.Address) ([]uint32, []common.Address, []*big.Int, []byte, error) {
	return _Attestations.Contract.GetCompletableAttestations(&_Attestations.CallOpts, identifier, account)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Attestations *AttestationsCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Attestations *AttestationsSession) GetEpochNumber() (*big.Int, error) {
	return _Attestations.Contract.GetEpochNumber(&_Attestations.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Attestations *AttestationsCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Attestations.Contract.GetEpochNumber(&_Attestations.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Attestations.Contract.GetEpochNumberOfBlock(&_Attestations.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Attestations.Contract.GetEpochNumberOfBlock(&_Attestations.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Attestations *AttestationsCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Attestations *AttestationsSession) GetEpochSize() (*big.Int, error) {
	return _Attestations.Contract.GetEpochSize(&_Attestations.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Attestations *AttestationsCallerSession) GetEpochSize() (*big.Int, error) {
	return _Attestations.Contract.GetEpochSize(&_Attestations.CallOpts)
}

// GetMaxAttestations is a free data retrieval call binding the contract method 0x7796a684.
//
// Solidity: function getMaxAttestations() view returns(uint256)
func (_Attestations *AttestationsCaller) GetMaxAttestations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getMaxAttestations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxAttestations is a free data retrieval call binding the contract method 0x7796a684.
//
// Solidity: function getMaxAttestations() view returns(uint256)
func (_Attestations *AttestationsSession) GetMaxAttestations() (*big.Int, error) {
	return _Attestations.Contract.GetMaxAttestations(&_Attestations.CallOpts)
}

// GetMaxAttestations is a free data retrieval call binding the contract method 0x7796a684.
//
// Solidity: function getMaxAttestations() view returns(uint256)
func (_Attestations *AttestationsCallerSession) GetMaxAttestations() (*big.Int, error) {
	return _Attestations.Contract.GetMaxAttestations(&_Attestations.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Attestations *AttestationsCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Attestations *AttestationsSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Attestations.Contract.GetParentSealBitmap(&_Attestations.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Attestations *AttestationsCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Attestations.Contract.GetParentSealBitmap(&_Attestations.CallOpts, blockNumber)
}

// GetUnselectedRequest is a free data retrieval call binding the contract method 0xe3d0f66f.
//
// Solidity: function getUnselectedRequest(bytes32 identifier, address account) view returns(uint32, uint32, address)
func (_Attestations *AttestationsCaller) GetUnselectedRequest(opts *bind.CallOpts, identifier [32]byte, account common.Address) (uint32, uint32, common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getUnselectedRequest", identifier, account)

	if err != nil {
		return *new(uint32), *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return out0, out1, out2, err

}

// GetUnselectedRequest is a free data retrieval call binding the contract method 0xe3d0f66f.
//
// Solidity: function getUnselectedRequest(bytes32 identifier, address account) view returns(uint32, uint32, address)
func (_Attestations *AttestationsSession) GetUnselectedRequest(identifier [32]byte, account common.Address) (uint32, uint32, common.Address, error) {
	return _Attestations.Contract.GetUnselectedRequest(&_Attestations.CallOpts, identifier, account)
}

// GetUnselectedRequest is a free data retrieval call binding the contract method 0xe3d0f66f.
//
// Solidity: function getUnselectedRequest(bytes32 identifier, address account) view returns(uint32, uint32, address)
func (_Attestations *AttestationsCallerSession) GetUnselectedRequest(identifier [32]byte, account common.Address) (uint32, uint32, common.Address, error) {
	return _Attestations.Contract.GetUnselectedRequest(&_Attestations.CallOpts, identifier, account)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Attestations *AttestationsCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Attestations *AttestationsSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Attestations.Contract.GetVerifiedSealBitmapFromHeader(&_Attestations.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Attestations *AttestationsCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Attestations.Contract.GetVerifiedSealBitmapFromHeader(&_Attestations.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Attestations *AttestationsCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "getVersionNumber")

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
func (_Attestations *AttestationsSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Attestations.Contract.GetVersionNumber(&_Attestations.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Attestations *AttestationsCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Attestations.Contract.GetVersionNumber(&_Attestations.CallOpts)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Attestations *AttestationsCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Attestations *AttestationsSession) HashHeader(header []byte) ([32]byte, error) {
	return _Attestations.Contract.HashHeader(&_Attestations.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Attestations *AttestationsCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Attestations.Contract.HashHeader(&_Attestations.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Attestations *AttestationsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Attestations *AttestationsSession) Initialized() (bool, error) {
	return _Attestations.Contract.Initialized(&_Attestations.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Attestations *AttestationsCallerSession) Initialized() (bool, error) {
	return _Attestations.Contract.Initialized(&_Attestations.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Attestations *AttestationsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Attestations *AttestationsSession) IsOwner() (bool, error) {
	return _Attestations.Contract.IsOwner(&_Attestations.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Attestations *AttestationsCallerSession) IsOwner() (bool, error) {
	return _Attestations.Contract.IsOwner(&_Attestations.CallOpts)
}

// LookupAccountsForIdentifier is a free data retrieval call binding the contract method 0x03cc1aff.
//
// Solidity: function lookupAccountsForIdentifier(bytes32 identifier) view returns(address[])
func (_Attestations *AttestationsCaller) LookupAccountsForIdentifier(opts *bind.CallOpts, identifier [32]byte) ([]common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "lookupAccountsForIdentifier", identifier)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// LookupAccountsForIdentifier is a free data retrieval call binding the contract method 0x03cc1aff.
//
// Solidity: function lookupAccountsForIdentifier(bytes32 identifier) view returns(address[])
func (_Attestations *AttestationsSession) LookupAccountsForIdentifier(identifier [32]byte) ([]common.Address, error) {
	return _Attestations.Contract.LookupAccountsForIdentifier(&_Attestations.CallOpts, identifier)
}

// LookupAccountsForIdentifier is a free data retrieval call binding the contract method 0x03cc1aff.
//
// Solidity: function lookupAccountsForIdentifier(bytes32 identifier) view returns(address[])
func (_Attestations *AttestationsCallerSession) LookupAccountsForIdentifier(identifier [32]byte) ([]common.Address, error) {
	return _Attestations.Contract.LookupAccountsForIdentifier(&_Attestations.CallOpts, identifier)
}

// MaxAttestations is a free data retrieval call binding the contract method 0x8218c6fe.
//
// Solidity: function maxAttestations() view returns(uint256)
func (_Attestations *AttestationsCaller) MaxAttestations(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "maxAttestations")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxAttestations is a free data retrieval call binding the contract method 0x8218c6fe.
//
// Solidity: function maxAttestations() view returns(uint256)
func (_Attestations *AttestationsSession) MaxAttestations() (*big.Int, error) {
	return _Attestations.Contract.MaxAttestations(&_Attestations.CallOpts)
}

// MaxAttestations is a free data retrieval call binding the contract method 0x8218c6fe.
//
// Solidity: function maxAttestations() view returns(uint256)
func (_Attestations *AttestationsCallerSession) MaxAttestations() (*big.Int, error) {
	return _Attestations.Contract.MaxAttestations(&_Attestations.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Attestations.Contract.MinQuorumSize(&_Attestations.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Attestations.Contract.MinQuorumSize(&_Attestations.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Attestations *AttestationsCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Attestations *AttestationsSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Attestations.Contract.MinQuorumSizeInCurrentSet(&_Attestations.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Attestations *AttestationsCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Attestations.Contract.MinQuorumSizeInCurrentSet(&_Attestations.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Attestations *AttestationsCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Attestations *AttestationsSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Attestations.Contract.NumberValidatorsInCurrentSet(&_Attestations.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Attestations *AttestationsCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Attestations.Contract.NumberValidatorsInCurrentSet(&_Attestations.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Attestations.Contract.NumberValidatorsInSet(&_Attestations.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Attestations *AttestationsCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Attestations.Contract.NumberValidatorsInSet(&_Attestations.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attestations *AttestationsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attestations *AttestationsSession) Owner() (common.Address, error) {
	return _Attestations.Contract.Owner(&_Attestations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Attestations *AttestationsCallerSession) Owner() (common.Address, error) {
	return _Attestations.Contract.Owner(&_Attestations.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xe831be58.
//
// Solidity: function pendingWithdrawals(address , address ) view returns(uint256)
func (_Attestations *AttestationsCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "pendingWithdrawals", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xe831be58.
//
// Solidity: function pendingWithdrawals(address , address ) view returns(uint256)
func (_Attestations *AttestationsSession) PendingWithdrawals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Attestations.Contract.PendingWithdrawals(&_Attestations.CallOpts, arg0, arg1)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xe831be58.
//
// Solidity: function pendingWithdrawals(address , address ) view returns(uint256)
func (_Attestations *AttestationsCallerSession) PendingWithdrawals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Attestations.Contract.PendingWithdrawals(&_Attestations.CallOpts, arg0, arg1)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Attestations *AttestationsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Attestations *AttestationsSession) Registry() (common.Address, error) {
	return _Attestations.Contract.Registry(&_Attestations.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Attestations *AttestationsCallerSession) Registry() (common.Address, error) {
	return _Attestations.Contract.Registry(&_Attestations.CallOpts)
}

// RequireNAttestationsRequested is a free data retrieval call binding the contract method 0xa762825a.
//
// Solidity: function requireNAttestationsRequested(bytes32 identifier, address account, uint32 expected) view returns()
func (_Attestations *AttestationsCaller) RequireNAttestationsRequested(opts *bind.CallOpts, identifier [32]byte, account common.Address, expected uint32) error {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "requireNAttestationsRequested", identifier, account, expected)

	if err != nil {
		return err
	}

	return err

}

// RequireNAttestationsRequested is a free data retrieval call binding the contract method 0xa762825a.
//
// Solidity: function requireNAttestationsRequested(bytes32 identifier, address account, uint32 expected) view returns()
func (_Attestations *AttestationsSession) RequireNAttestationsRequested(identifier [32]byte, account common.Address, expected uint32) error {
	return _Attestations.Contract.RequireNAttestationsRequested(&_Attestations.CallOpts, identifier, account, expected)
}

// RequireNAttestationsRequested is a free data retrieval call binding the contract method 0xa762825a.
//
// Solidity: function requireNAttestationsRequested(bytes32 identifier, address account, uint32 expected) view returns()
func (_Attestations *AttestationsCallerSession) RequireNAttestationsRequested(identifier [32]byte, account common.Address, expected uint32) error {
	return _Attestations.Contract.RequireNAttestationsRequested(&_Attestations.CallOpts, identifier, account, expected)
}

// SelectIssuersWaitBlocks is a free data retrieval call binding the contract method 0x89d35286.
//
// Solidity: function selectIssuersWaitBlocks() view returns(uint256)
func (_Attestations *AttestationsCaller) SelectIssuersWaitBlocks(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "selectIssuersWaitBlocks")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SelectIssuersWaitBlocks is a free data retrieval call binding the contract method 0x89d35286.
//
// Solidity: function selectIssuersWaitBlocks() view returns(uint256)
func (_Attestations *AttestationsSession) SelectIssuersWaitBlocks() (*big.Int, error) {
	return _Attestations.Contract.SelectIssuersWaitBlocks(&_Attestations.CallOpts)
}

// SelectIssuersWaitBlocks is a free data retrieval call binding the contract method 0x89d35286.
//
// Solidity: function selectIssuersWaitBlocks() view returns(uint256)
func (_Attestations *AttestationsCallerSession) SelectIssuersWaitBlocks() (*big.Int, error) {
	return _Attestations.Contract.SelectIssuersWaitBlocks(&_Attestations.CallOpts)
}

// TransferApprovals is a free data retrieval call binding the contract method 0x84a1a4fc.
//
// Solidity: function transferApprovals(address , bytes32 ) view returns(bool)
func (_Attestations *AttestationsCaller) TransferApprovals(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "transferApprovals", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferApprovals is a free data retrieval call binding the contract method 0x84a1a4fc.
//
// Solidity: function transferApprovals(address , bytes32 ) view returns(bool)
func (_Attestations *AttestationsSession) TransferApprovals(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Attestations.Contract.TransferApprovals(&_Attestations.CallOpts, arg0, arg1)
}

// TransferApprovals is a free data retrieval call binding the contract method 0x84a1a4fc.
//
// Solidity: function transferApprovals(address , bytes32 ) view returns(bool)
func (_Attestations *AttestationsCallerSession) TransferApprovals(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _Attestations.Contract.TransferApprovals(&_Attestations.CallOpts, arg0, arg1)
}

// ValidateAttestationCode is a free data retrieval call binding the contract method 0x5ce9bc07.
//
// Solidity: function validateAttestationCode(bytes32 identifier, address account, uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Attestations *AttestationsCaller) ValidateAttestationCode(opts *bind.CallOpts, identifier [32]byte, account common.Address, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "validateAttestationCode", identifier, account, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidateAttestationCode is a free data retrieval call binding the contract method 0x5ce9bc07.
//
// Solidity: function validateAttestationCode(bytes32 identifier, address account, uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Attestations *AttestationsSession) ValidateAttestationCode(identifier [32]byte, account common.Address, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Attestations.Contract.ValidateAttestationCode(&_Attestations.CallOpts, identifier, account, v, r, s)
}

// ValidateAttestationCode is a free data retrieval call binding the contract method 0x5ce9bc07.
//
// Solidity: function validateAttestationCode(bytes32 identifier, address account, uint8 v, bytes32 r, bytes32 s) view returns(address)
func (_Attestations *AttestationsCallerSession) ValidateAttestationCode(identifier [32]byte, account common.Address, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Attestations.Contract.ValidateAttestationCode(&_Attestations.CallOpts, identifier, account, v, r, s)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Attestations *AttestationsCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Attestations *AttestationsSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Attestations.Contract.ValidatorSignerAddressFromCurrentSet(&_Attestations.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Attestations *AttestationsCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Attestations.Contract.ValidatorSignerAddressFromCurrentSet(&_Attestations.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Attestations *AttestationsCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Attestations.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Attestations *AttestationsSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Attestations.Contract.ValidatorSignerAddressFromSet(&_Attestations.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Attestations *AttestationsCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Attestations.Contract.ValidatorSignerAddressFromSet(&_Attestations.CallOpts, index, blockNumber)
}

// Initialize is a paid mutator transaction binding the contract method 0xfd536f5d.
//
// Solidity: function initialize(address registryAddress, uint256 _attestationExpiryBlocks, uint256 _selectIssuersWaitBlocks, uint256 _maxAttestations, address[] attestationRequestFeeTokens, uint256[] attestationRequestFeeValues) returns()
func (_Attestations *AttestationsTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _attestationExpiryBlocks *big.Int, _selectIssuersWaitBlocks *big.Int, _maxAttestations *big.Int, attestationRequestFeeTokens []common.Address, attestationRequestFeeValues []*big.Int) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "initialize", registryAddress, _attestationExpiryBlocks, _selectIssuersWaitBlocks, _maxAttestations, attestationRequestFeeTokens, attestationRequestFeeValues)
}

// Initialize is a paid mutator transaction binding the contract method 0xfd536f5d.
//
// Solidity: function initialize(address registryAddress, uint256 _attestationExpiryBlocks, uint256 _selectIssuersWaitBlocks, uint256 _maxAttestations, address[] attestationRequestFeeTokens, uint256[] attestationRequestFeeValues) returns()
func (_Attestations *AttestationsSession) Initialize(registryAddress common.Address, _attestationExpiryBlocks *big.Int, _selectIssuersWaitBlocks *big.Int, _maxAttestations *big.Int, attestationRequestFeeTokens []common.Address, attestationRequestFeeValues []*big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.Initialize(&_Attestations.TransactOpts, registryAddress, _attestationExpiryBlocks, _selectIssuersWaitBlocks, _maxAttestations, attestationRequestFeeTokens, attestationRequestFeeValues)
}

// Initialize is a paid mutator transaction binding the contract method 0xfd536f5d.
//
// Solidity: function initialize(address registryAddress, uint256 _attestationExpiryBlocks, uint256 _selectIssuersWaitBlocks, uint256 _maxAttestations, address[] attestationRequestFeeTokens, uint256[] attestationRequestFeeValues) returns()
func (_Attestations *AttestationsTransactorSession) Initialize(registryAddress common.Address, _attestationExpiryBlocks *big.Int, _selectIssuersWaitBlocks *big.Int, _maxAttestations *big.Int, attestationRequestFeeTokens []common.Address, attestationRequestFeeValues []*big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.Initialize(&_Attestations.TransactOpts, registryAddress, _attestationExpiryBlocks, _selectIssuersWaitBlocks, _maxAttestations, attestationRequestFeeTokens, attestationRequestFeeValues)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Attestations *AttestationsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Attestations *AttestationsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Attestations.Contract.RenounceOwnership(&_Attestations.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Attestations *AttestationsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Attestations.Contract.RenounceOwnership(&_Attestations.TransactOpts)
}

// Revoke is a paid mutator transaction binding the contract method 0xbb46942f.
//
// Solidity: function revoke(bytes32 identifier, uint256 index) returns()
func (_Attestations *AttestationsTransactor) Revoke(opts *bind.TransactOpts, identifier [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "revoke", identifier, index)
}

// Revoke is a paid mutator transaction binding the contract method 0xbb46942f.
//
// Solidity: function revoke(bytes32 identifier, uint256 index) returns()
func (_Attestations *AttestationsSession) Revoke(identifier [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.Revoke(&_Attestations.TransactOpts, identifier, index)
}

// Revoke is a paid mutator transaction binding the contract method 0xbb46942f.
//
// Solidity: function revoke(bytes32 identifier, uint256 index) returns()
func (_Attestations *AttestationsTransactorSession) Revoke(identifier [32]byte, index *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.Revoke(&_Attestations.TransactOpts, identifier, index)
}

// SetAttestationExpiryBlocks is a paid mutator transaction binding the contract method 0xa6437e73.
//
// Solidity: function setAttestationExpiryBlocks(uint256 _attestationExpiryBlocks) returns()
func (_Attestations *AttestationsTransactor) SetAttestationExpiryBlocks(opts *bind.TransactOpts, _attestationExpiryBlocks *big.Int) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "setAttestationExpiryBlocks", _attestationExpiryBlocks)
}

// SetAttestationExpiryBlocks is a paid mutator transaction binding the contract method 0xa6437e73.
//
// Solidity: function setAttestationExpiryBlocks(uint256 _attestationExpiryBlocks) returns()
func (_Attestations *AttestationsSession) SetAttestationExpiryBlocks(_attestationExpiryBlocks *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetAttestationExpiryBlocks(&_Attestations.TransactOpts, _attestationExpiryBlocks)
}

// SetAttestationExpiryBlocks is a paid mutator transaction binding the contract method 0xa6437e73.
//
// Solidity: function setAttestationExpiryBlocks(uint256 _attestationExpiryBlocks) returns()
func (_Attestations *AttestationsTransactorSession) SetAttestationExpiryBlocks(_attestationExpiryBlocks *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetAttestationExpiryBlocks(&_Attestations.TransactOpts, _attestationExpiryBlocks)
}

// SetAttestationRequestFee is a paid mutator transaction binding the contract method 0xf3ff26c6.
//
// Solidity: function setAttestationRequestFee(address token, uint256 fee) returns()
func (_Attestations *AttestationsTransactor) SetAttestationRequestFee(opts *bind.TransactOpts, token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "setAttestationRequestFee", token, fee)
}

// SetAttestationRequestFee is a paid mutator transaction binding the contract method 0xf3ff26c6.
//
// Solidity: function setAttestationRequestFee(address token, uint256 fee) returns()
func (_Attestations *AttestationsSession) SetAttestationRequestFee(token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetAttestationRequestFee(&_Attestations.TransactOpts, token, fee)
}

// SetAttestationRequestFee is a paid mutator transaction binding the contract method 0xf3ff26c6.
//
// Solidity: function setAttestationRequestFee(address token, uint256 fee) returns()
func (_Attestations *AttestationsTransactorSession) SetAttestationRequestFee(token common.Address, fee *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetAttestationRequestFee(&_Attestations.TransactOpts, token, fee)
}

// SetMaxAttestations is a paid mutator transaction binding the contract method 0xbe2c47a6.
//
// Solidity: function setMaxAttestations(uint256 _maxAttestations) returns()
func (_Attestations *AttestationsTransactor) SetMaxAttestations(opts *bind.TransactOpts, _maxAttestations *big.Int) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "setMaxAttestations", _maxAttestations)
}

// SetMaxAttestations is a paid mutator transaction binding the contract method 0xbe2c47a6.
//
// Solidity: function setMaxAttestations(uint256 _maxAttestations) returns()
func (_Attestations *AttestationsSession) SetMaxAttestations(_maxAttestations *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetMaxAttestations(&_Attestations.TransactOpts, _maxAttestations)
}

// SetMaxAttestations is a paid mutator transaction binding the contract method 0xbe2c47a6.
//
// Solidity: function setMaxAttestations(uint256 _maxAttestations) returns()
func (_Attestations *AttestationsTransactorSession) SetMaxAttestations(_maxAttestations *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetMaxAttestations(&_Attestations.TransactOpts, _maxAttestations)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Attestations *AttestationsTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Attestations *AttestationsSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Attestations.Contract.SetRegistry(&_Attestations.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Attestations *AttestationsTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Attestations.Contract.SetRegistry(&_Attestations.TransactOpts, registryAddress)
}

// SetSelectIssuersWaitBlocks is a paid mutator transaction binding the contract method 0xe02659ce.
//
// Solidity: function setSelectIssuersWaitBlocks(uint256 _selectIssuersWaitBlocks) returns()
func (_Attestations *AttestationsTransactor) SetSelectIssuersWaitBlocks(opts *bind.TransactOpts, _selectIssuersWaitBlocks *big.Int) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "setSelectIssuersWaitBlocks", _selectIssuersWaitBlocks)
}

// SetSelectIssuersWaitBlocks is a paid mutator transaction binding the contract method 0xe02659ce.
//
// Solidity: function setSelectIssuersWaitBlocks(uint256 _selectIssuersWaitBlocks) returns()
func (_Attestations *AttestationsSession) SetSelectIssuersWaitBlocks(_selectIssuersWaitBlocks *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetSelectIssuersWaitBlocks(&_Attestations.TransactOpts, _selectIssuersWaitBlocks)
}

// SetSelectIssuersWaitBlocks is a paid mutator transaction binding the contract method 0xe02659ce.
//
// Solidity: function setSelectIssuersWaitBlocks(uint256 _selectIssuersWaitBlocks) returns()
func (_Attestations *AttestationsTransactorSession) SetSelectIssuersWaitBlocks(_selectIssuersWaitBlocks *big.Int) (*types.Transaction, error) {
	return _Attestations.Contract.SetSelectIssuersWaitBlocks(&_Attestations.TransactOpts, _selectIssuersWaitBlocks)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Attestations *AttestationsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Attestations *AttestationsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Attestations.Contract.TransferOwnership(&_Attestations.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Attestations *AttestationsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Attestations.Contract.TransferOwnership(&_Attestations.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_Attestations *AttestationsTransactor) Withdraw(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Attestations.contract.Transact(opts, "withdraw", token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_Attestations *AttestationsSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _Attestations.Contract.Withdraw(&_Attestations.TransactOpts, token)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address token) returns()
func (_Attestations *AttestationsTransactorSession) Withdraw(token common.Address) (*types.Transaction, error) {
	return _Attestations.Contract.Withdraw(&_Attestations.TransactOpts, token)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Attestations *AttestationsFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Attestations.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AttestationCompleted":
		event, err = _Attestations.ParseAttestationCompleted(log)
	case "AttestationExpiryBlocksSet":
		event, err = _Attestations.ParseAttestationExpiryBlocksSet(log)
	case "AttestationIssuerSelected":
		event, err = _Attestations.ParseAttestationIssuerSelected(log)
	case "AttestationRequestFeeSet":
		event, err = _Attestations.ParseAttestationRequestFeeSet(log)
	case "AttestationsRequested":
		event, err = _Attestations.ParseAttestationsRequested(log)
	case "AttestationsTransferred":
		event, err = _Attestations.ParseAttestationsTransferred(log)
	case "MaxAttestationsSet":
		event, err = _Attestations.ParseMaxAttestationsSet(log)
	case "OwnershipTransferred":
		event, err = _Attestations.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Attestations.ParseRegistrySet(log)
	case "SelectIssuersWaitBlocksSet":
		event, err = _Attestations.ParseSelectIssuersWaitBlocksSet(log)
	case "TransferApproval":
		event, err = _Attestations.ParseTransferApproval(log)
	case "Withdrawal":
		event, err = _Attestations.ParseWithdrawal(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// AttestationsAttestationCompletedIterator is returned from FilterAttestationCompleted and is used to iterate over the raw logs and unpacked data for AttestationCompleted events raised by the Attestations contract.
type AttestationsAttestationCompletedIterator struct {
	Event *AttestationsAttestationCompleted // Event containing the contract specifics and raw log

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
func (it *AttestationsAttestationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsAttestationCompleted)
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
		it.Event = new(AttestationsAttestationCompleted)
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
func (it *AttestationsAttestationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsAttestationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsAttestationCompleted represents a AttestationCompleted event raised by the Attestations contract.
type AttestationsAttestationCompleted struct {
	Identifier [32]byte
	Account    common.Address
	Issuer     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAttestationCompleted is a free log retrieval operation binding the contract event 0x414ff2c18c092697c4b8de49f515ac44f8bebc19b24553cf58ace913a6ac639d.
//
// Solidity: event AttestationCompleted(bytes32 indexed identifier, address indexed account, address indexed issuer)
func (_Attestations *AttestationsFilterer) FilterAttestationCompleted(opts *bind.FilterOpts, identifier [][32]byte, account []common.Address, issuer []common.Address) (*AttestationsAttestationCompletedIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "AttestationCompleted", identifierRule, accountRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsAttestationCompletedIterator{contract: _Attestations.contract, event: "AttestationCompleted", logs: logs, sub: sub}, nil
}

// WatchAttestationCompleted is a free log subscription operation binding the contract event 0x414ff2c18c092697c4b8de49f515ac44f8bebc19b24553cf58ace913a6ac639d.
//
// Solidity: event AttestationCompleted(bytes32 indexed identifier, address indexed account, address indexed issuer)
func (_Attestations *AttestationsFilterer) WatchAttestationCompleted(opts *bind.WatchOpts, sink chan<- *AttestationsAttestationCompleted, identifier [][32]byte, account []common.Address, issuer []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "AttestationCompleted", identifierRule, accountRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsAttestationCompleted)
				if err := _Attestations.contract.UnpackLog(event, "AttestationCompleted", log); err != nil {
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

// ParseAttestationCompleted is a log parse operation binding the contract event 0x414ff2c18c092697c4b8de49f515ac44f8bebc19b24553cf58ace913a6ac639d.
//
// Solidity: event AttestationCompleted(bytes32 indexed identifier, address indexed account, address indexed issuer)
func (_Attestations *AttestationsFilterer) ParseAttestationCompleted(log types.Log) (*AttestationsAttestationCompleted, error) {
	event := new(AttestationsAttestationCompleted)
	if err := _Attestations.contract.UnpackLog(event, "AttestationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsAttestationExpiryBlocksSetIterator is returned from FilterAttestationExpiryBlocksSet and is used to iterate over the raw logs and unpacked data for AttestationExpiryBlocksSet events raised by the Attestations contract.
type AttestationsAttestationExpiryBlocksSetIterator struct {
	Event *AttestationsAttestationExpiryBlocksSet // Event containing the contract specifics and raw log

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
func (it *AttestationsAttestationExpiryBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsAttestationExpiryBlocksSet)
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
		it.Event = new(AttestationsAttestationExpiryBlocksSet)
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
func (it *AttestationsAttestationExpiryBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsAttestationExpiryBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsAttestationExpiryBlocksSet represents a AttestationExpiryBlocksSet event raised by the Attestations contract.
type AttestationsAttestationExpiryBlocksSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttestationExpiryBlocksSet is a free log retrieval operation binding the contract event 0x4fbe976a07a9260091c2d347f8780c4bc636392e34d5b249b367baf8a5c7ca69.
//
// Solidity: event AttestationExpiryBlocksSet(uint256 value)
func (_Attestations *AttestationsFilterer) FilterAttestationExpiryBlocksSet(opts *bind.FilterOpts) (*AttestationsAttestationExpiryBlocksSetIterator, error) {

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "AttestationExpiryBlocksSet")
	if err != nil {
		return nil, err
	}
	return &AttestationsAttestationExpiryBlocksSetIterator{contract: _Attestations.contract, event: "AttestationExpiryBlocksSet", logs: logs, sub: sub}, nil
}

// WatchAttestationExpiryBlocksSet is a free log subscription operation binding the contract event 0x4fbe976a07a9260091c2d347f8780c4bc636392e34d5b249b367baf8a5c7ca69.
//
// Solidity: event AttestationExpiryBlocksSet(uint256 value)
func (_Attestations *AttestationsFilterer) WatchAttestationExpiryBlocksSet(opts *bind.WatchOpts, sink chan<- *AttestationsAttestationExpiryBlocksSet) (event.Subscription, error) {

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "AttestationExpiryBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsAttestationExpiryBlocksSet)
				if err := _Attestations.contract.UnpackLog(event, "AttestationExpiryBlocksSet", log); err != nil {
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

// ParseAttestationExpiryBlocksSet is a log parse operation binding the contract event 0x4fbe976a07a9260091c2d347f8780c4bc636392e34d5b249b367baf8a5c7ca69.
//
// Solidity: event AttestationExpiryBlocksSet(uint256 value)
func (_Attestations *AttestationsFilterer) ParseAttestationExpiryBlocksSet(log types.Log) (*AttestationsAttestationExpiryBlocksSet, error) {
	event := new(AttestationsAttestationExpiryBlocksSet)
	if err := _Attestations.contract.UnpackLog(event, "AttestationExpiryBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsAttestationIssuerSelectedIterator is returned from FilterAttestationIssuerSelected and is used to iterate over the raw logs and unpacked data for AttestationIssuerSelected events raised by the Attestations contract.
type AttestationsAttestationIssuerSelectedIterator struct {
	Event *AttestationsAttestationIssuerSelected // Event containing the contract specifics and raw log

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
func (it *AttestationsAttestationIssuerSelectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsAttestationIssuerSelected)
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
		it.Event = new(AttestationsAttestationIssuerSelected)
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
func (it *AttestationsAttestationIssuerSelectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsAttestationIssuerSelectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsAttestationIssuerSelected represents a AttestationIssuerSelected event raised by the Attestations contract.
type AttestationsAttestationIssuerSelected struct {
	Identifier                 [32]byte
	Account                    common.Address
	Issuer                     common.Address
	AttestationRequestFeeToken common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAttestationIssuerSelected is a free log retrieval operation binding the contract event 0xaf7f470b643316cf44c1f2898328a075e7602945b4f8584f48ba4ad2d8a2ea9d.
//
// Solidity: event AttestationIssuerSelected(bytes32 indexed identifier, address indexed account, address indexed issuer, address attestationRequestFeeToken)
func (_Attestations *AttestationsFilterer) FilterAttestationIssuerSelected(opts *bind.FilterOpts, identifier [][32]byte, account []common.Address, issuer []common.Address) (*AttestationsAttestationIssuerSelectedIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "AttestationIssuerSelected", identifierRule, accountRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsAttestationIssuerSelectedIterator{contract: _Attestations.contract, event: "AttestationIssuerSelected", logs: logs, sub: sub}, nil
}

// WatchAttestationIssuerSelected is a free log subscription operation binding the contract event 0xaf7f470b643316cf44c1f2898328a075e7602945b4f8584f48ba4ad2d8a2ea9d.
//
// Solidity: event AttestationIssuerSelected(bytes32 indexed identifier, address indexed account, address indexed issuer, address attestationRequestFeeToken)
func (_Attestations *AttestationsFilterer) WatchAttestationIssuerSelected(opts *bind.WatchOpts, sink chan<- *AttestationsAttestationIssuerSelected, identifier [][32]byte, account []common.Address, issuer []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "AttestationIssuerSelected", identifierRule, accountRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsAttestationIssuerSelected)
				if err := _Attestations.contract.UnpackLog(event, "AttestationIssuerSelected", log); err != nil {
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

// ParseAttestationIssuerSelected is a log parse operation binding the contract event 0xaf7f470b643316cf44c1f2898328a075e7602945b4f8584f48ba4ad2d8a2ea9d.
//
// Solidity: event AttestationIssuerSelected(bytes32 indexed identifier, address indexed account, address indexed issuer, address attestationRequestFeeToken)
func (_Attestations *AttestationsFilterer) ParseAttestationIssuerSelected(log types.Log) (*AttestationsAttestationIssuerSelected, error) {
	event := new(AttestationsAttestationIssuerSelected)
	if err := _Attestations.contract.UnpackLog(event, "AttestationIssuerSelected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsAttestationRequestFeeSetIterator is returned from FilterAttestationRequestFeeSet and is used to iterate over the raw logs and unpacked data for AttestationRequestFeeSet events raised by the Attestations contract.
type AttestationsAttestationRequestFeeSetIterator struct {
	Event *AttestationsAttestationRequestFeeSet // Event containing the contract specifics and raw log

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
func (it *AttestationsAttestationRequestFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsAttestationRequestFeeSet)
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
		it.Event = new(AttestationsAttestationRequestFeeSet)
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
func (it *AttestationsAttestationRequestFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsAttestationRequestFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsAttestationRequestFeeSet represents a AttestationRequestFeeSet event raised by the Attestations contract.
type AttestationsAttestationRequestFeeSet struct {
	Token common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterAttestationRequestFeeSet is a free log retrieval operation binding the contract event 0x7cf8b633f218e9f9bc2c06107bcaddcfee6b90580863768acdcfd4f05d7af394.
//
// Solidity: event AttestationRequestFeeSet(address indexed token, uint256 value)
func (_Attestations *AttestationsFilterer) FilterAttestationRequestFeeSet(opts *bind.FilterOpts, token []common.Address) (*AttestationsAttestationRequestFeeSetIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "AttestationRequestFeeSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsAttestationRequestFeeSetIterator{contract: _Attestations.contract, event: "AttestationRequestFeeSet", logs: logs, sub: sub}, nil
}

// WatchAttestationRequestFeeSet is a free log subscription operation binding the contract event 0x7cf8b633f218e9f9bc2c06107bcaddcfee6b90580863768acdcfd4f05d7af394.
//
// Solidity: event AttestationRequestFeeSet(address indexed token, uint256 value)
func (_Attestations *AttestationsFilterer) WatchAttestationRequestFeeSet(opts *bind.WatchOpts, sink chan<- *AttestationsAttestationRequestFeeSet, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "AttestationRequestFeeSet", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsAttestationRequestFeeSet)
				if err := _Attestations.contract.UnpackLog(event, "AttestationRequestFeeSet", log); err != nil {
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

// ParseAttestationRequestFeeSet is a log parse operation binding the contract event 0x7cf8b633f218e9f9bc2c06107bcaddcfee6b90580863768acdcfd4f05d7af394.
//
// Solidity: event AttestationRequestFeeSet(address indexed token, uint256 value)
func (_Attestations *AttestationsFilterer) ParseAttestationRequestFeeSet(log types.Log) (*AttestationsAttestationRequestFeeSet, error) {
	event := new(AttestationsAttestationRequestFeeSet)
	if err := _Attestations.contract.UnpackLog(event, "AttestationRequestFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsAttestationsRequestedIterator is returned from FilterAttestationsRequested and is used to iterate over the raw logs and unpacked data for AttestationsRequested events raised by the Attestations contract.
type AttestationsAttestationsRequestedIterator struct {
	Event *AttestationsAttestationsRequested // Event containing the contract specifics and raw log

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
func (it *AttestationsAttestationsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsAttestationsRequested)
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
		it.Event = new(AttestationsAttestationsRequested)
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
func (it *AttestationsAttestationsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsAttestationsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsAttestationsRequested represents a AttestationsRequested event raised by the Attestations contract.
type AttestationsAttestationsRequested struct {
	Identifier                 [32]byte
	Account                    common.Address
	AttestationsRequested      *big.Int
	AttestationRequestFeeToken common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAttestationsRequested is a free log retrieval operation binding the contract event 0x381545d9b1fffcb94ffbbd0bccfff9f1fb3acd474d34f7d59112a5c9973fee49.
//
// Solidity: event AttestationsRequested(bytes32 indexed identifier, address indexed account, uint256 attestationsRequested, address attestationRequestFeeToken)
func (_Attestations *AttestationsFilterer) FilterAttestationsRequested(opts *bind.FilterOpts, identifier [][32]byte, account []common.Address) (*AttestationsAttestationsRequestedIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "AttestationsRequested", identifierRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsAttestationsRequestedIterator{contract: _Attestations.contract, event: "AttestationsRequested", logs: logs, sub: sub}, nil
}

// WatchAttestationsRequested is a free log subscription operation binding the contract event 0x381545d9b1fffcb94ffbbd0bccfff9f1fb3acd474d34f7d59112a5c9973fee49.
//
// Solidity: event AttestationsRequested(bytes32 indexed identifier, address indexed account, uint256 attestationsRequested, address attestationRequestFeeToken)
func (_Attestations *AttestationsFilterer) WatchAttestationsRequested(opts *bind.WatchOpts, sink chan<- *AttestationsAttestationsRequested, identifier [][32]byte, account []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "AttestationsRequested", identifierRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsAttestationsRequested)
				if err := _Attestations.contract.UnpackLog(event, "AttestationsRequested", log); err != nil {
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

// ParseAttestationsRequested is a log parse operation binding the contract event 0x381545d9b1fffcb94ffbbd0bccfff9f1fb3acd474d34f7d59112a5c9973fee49.
//
// Solidity: event AttestationsRequested(bytes32 indexed identifier, address indexed account, uint256 attestationsRequested, address attestationRequestFeeToken)
func (_Attestations *AttestationsFilterer) ParseAttestationsRequested(log types.Log) (*AttestationsAttestationsRequested, error) {
	event := new(AttestationsAttestationsRequested)
	if err := _Attestations.contract.UnpackLog(event, "AttestationsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsAttestationsTransferredIterator is returned from FilterAttestationsTransferred and is used to iterate over the raw logs and unpacked data for AttestationsTransferred events raised by the Attestations contract.
type AttestationsAttestationsTransferredIterator struct {
	Event *AttestationsAttestationsTransferred // Event containing the contract specifics and raw log

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
func (it *AttestationsAttestationsTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsAttestationsTransferred)
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
		it.Event = new(AttestationsAttestationsTransferred)
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
func (it *AttestationsAttestationsTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsAttestationsTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsAttestationsTransferred represents a AttestationsTransferred event raised by the Attestations contract.
type AttestationsAttestationsTransferred struct {
	Identifier  [32]byte
	FromAccount common.Address
	ToAccount   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAttestationsTransferred is a free log retrieval operation binding the contract event 0x35bc19e2c74829d0a96c765bb41b09ce24a9d0757486ced0d075e79089323638.
//
// Solidity: event AttestationsTransferred(bytes32 indexed identifier, address indexed fromAccount, address indexed toAccount)
func (_Attestations *AttestationsFilterer) FilterAttestationsTransferred(opts *bind.FilterOpts, identifier [][32]byte, fromAccount []common.Address, toAccount []common.Address) (*AttestationsAttestationsTransferredIterator, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var fromAccountRule []interface{}
	for _, fromAccountItem := range fromAccount {
		fromAccountRule = append(fromAccountRule, fromAccountItem)
	}
	var toAccountRule []interface{}
	for _, toAccountItem := range toAccount {
		toAccountRule = append(toAccountRule, toAccountItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "AttestationsTransferred", identifierRule, fromAccountRule, toAccountRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsAttestationsTransferredIterator{contract: _Attestations.contract, event: "AttestationsTransferred", logs: logs, sub: sub}, nil
}

// WatchAttestationsTransferred is a free log subscription operation binding the contract event 0x35bc19e2c74829d0a96c765bb41b09ce24a9d0757486ced0d075e79089323638.
//
// Solidity: event AttestationsTransferred(bytes32 indexed identifier, address indexed fromAccount, address indexed toAccount)
func (_Attestations *AttestationsFilterer) WatchAttestationsTransferred(opts *bind.WatchOpts, sink chan<- *AttestationsAttestationsTransferred, identifier [][32]byte, fromAccount []common.Address, toAccount []common.Address) (event.Subscription, error) {

	var identifierRule []interface{}
	for _, identifierItem := range identifier {
		identifierRule = append(identifierRule, identifierItem)
	}
	var fromAccountRule []interface{}
	for _, fromAccountItem := range fromAccount {
		fromAccountRule = append(fromAccountRule, fromAccountItem)
	}
	var toAccountRule []interface{}
	for _, toAccountItem := range toAccount {
		toAccountRule = append(toAccountRule, toAccountItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "AttestationsTransferred", identifierRule, fromAccountRule, toAccountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsAttestationsTransferred)
				if err := _Attestations.contract.UnpackLog(event, "AttestationsTransferred", log); err != nil {
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

// ParseAttestationsTransferred is a log parse operation binding the contract event 0x35bc19e2c74829d0a96c765bb41b09ce24a9d0757486ced0d075e79089323638.
//
// Solidity: event AttestationsTransferred(bytes32 indexed identifier, address indexed fromAccount, address indexed toAccount)
func (_Attestations *AttestationsFilterer) ParseAttestationsTransferred(log types.Log) (*AttestationsAttestationsTransferred, error) {
	event := new(AttestationsAttestationsTransferred)
	if err := _Attestations.contract.UnpackLog(event, "AttestationsTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsMaxAttestationsSetIterator is returned from FilterMaxAttestationsSet and is used to iterate over the raw logs and unpacked data for MaxAttestationsSet events raised by the Attestations contract.
type AttestationsMaxAttestationsSetIterator struct {
	Event *AttestationsMaxAttestationsSet // Event containing the contract specifics and raw log

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
func (it *AttestationsMaxAttestationsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsMaxAttestationsSet)
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
		it.Event = new(AttestationsMaxAttestationsSet)
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
func (it *AttestationsMaxAttestationsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsMaxAttestationsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsMaxAttestationsSet represents a MaxAttestationsSet event raised by the Attestations contract.
type AttestationsMaxAttestationsSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMaxAttestationsSet is a free log retrieval operation binding the contract event 0xc1f217a1246a98ce04e938768309107630ed86c1e0e9f9995af28e23a9c06178.
//
// Solidity: event MaxAttestationsSet(uint256 value)
func (_Attestations *AttestationsFilterer) FilterMaxAttestationsSet(opts *bind.FilterOpts) (*AttestationsMaxAttestationsSetIterator, error) {

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "MaxAttestationsSet")
	if err != nil {
		return nil, err
	}
	return &AttestationsMaxAttestationsSetIterator{contract: _Attestations.contract, event: "MaxAttestationsSet", logs: logs, sub: sub}, nil
}

// WatchMaxAttestationsSet is a free log subscription operation binding the contract event 0xc1f217a1246a98ce04e938768309107630ed86c1e0e9f9995af28e23a9c06178.
//
// Solidity: event MaxAttestationsSet(uint256 value)
func (_Attestations *AttestationsFilterer) WatchMaxAttestationsSet(opts *bind.WatchOpts, sink chan<- *AttestationsMaxAttestationsSet) (event.Subscription, error) {

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "MaxAttestationsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsMaxAttestationsSet)
				if err := _Attestations.contract.UnpackLog(event, "MaxAttestationsSet", log); err != nil {
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

// ParseMaxAttestationsSet is a log parse operation binding the contract event 0xc1f217a1246a98ce04e938768309107630ed86c1e0e9f9995af28e23a9c06178.
//
// Solidity: event MaxAttestationsSet(uint256 value)
func (_Attestations *AttestationsFilterer) ParseMaxAttestationsSet(log types.Log) (*AttestationsMaxAttestationsSet, error) {
	event := new(AttestationsMaxAttestationsSet)
	if err := _Attestations.contract.UnpackLog(event, "MaxAttestationsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Attestations contract.
type AttestationsOwnershipTransferredIterator struct {
	Event *AttestationsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AttestationsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsOwnershipTransferred)
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
		it.Event = new(AttestationsOwnershipTransferred)
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
func (it *AttestationsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsOwnershipTransferred represents a OwnershipTransferred event raised by the Attestations contract.
type AttestationsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Attestations *AttestationsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AttestationsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsOwnershipTransferredIterator{contract: _Attestations.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Attestations *AttestationsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AttestationsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsOwnershipTransferred)
				if err := _Attestations.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Attestations *AttestationsFilterer) ParseOwnershipTransferred(log types.Log) (*AttestationsOwnershipTransferred, error) {
	event := new(AttestationsOwnershipTransferred)
	if err := _Attestations.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Attestations contract.
type AttestationsRegistrySetIterator struct {
	Event *AttestationsRegistrySet // Event containing the contract specifics and raw log

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
func (it *AttestationsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsRegistrySet)
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
		it.Event = new(AttestationsRegistrySet)
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
func (it *AttestationsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsRegistrySet represents a RegistrySet event raised by the Attestations contract.
type AttestationsRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Attestations *AttestationsFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*AttestationsRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsRegistrySetIterator{contract: _Attestations.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Attestations *AttestationsFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *AttestationsRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsRegistrySet)
				if err := _Attestations.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Attestations *AttestationsFilterer) ParseRegistrySet(log types.Log) (*AttestationsRegistrySet, error) {
	event := new(AttestationsRegistrySet)
	if err := _Attestations.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsSelectIssuersWaitBlocksSetIterator is returned from FilterSelectIssuersWaitBlocksSet and is used to iterate over the raw logs and unpacked data for SelectIssuersWaitBlocksSet events raised by the Attestations contract.
type AttestationsSelectIssuersWaitBlocksSetIterator struct {
	Event *AttestationsSelectIssuersWaitBlocksSet // Event containing the contract specifics and raw log

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
func (it *AttestationsSelectIssuersWaitBlocksSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsSelectIssuersWaitBlocksSet)
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
		it.Event = new(AttestationsSelectIssuersWaitBlocksSet)
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
func (it *AttestationsSelectIssuersWaitBlocksSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsSelectIssuersWaitBlocksSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsSelectIssuersWaitBlocksSet represents a SelectIssuersWaitBlocksSet event raised by the Attestations contract.
type AttestationsSelectIssuersWaitBlocksSet struct {
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSelectIssuersWaitBlocksSet is a free log retrieval operation binding the contract event 0x954fa47fa6f4e8017b99f93c73f4fbe599d786f9f5da73fe9086ab473fb455d8.
//
// Solidity: event SelectIssuersWaitBlocksSet(uint256 value)
func (_Attestations *AttestationsFilterer) FilterSelectIssuersWaitBlocksSet(opts *bind.FilterOpts) (*AttestationsSelectIssuersWaitBlocksSetIterator, error) {

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "SelectIssuersWaitBlocksSet")
	if err != nil {
		return nil, err
	}
	return &AttestationsSelectIssuersWaitBlocksSetIterator{contract: _Attestations.contract, event: "SelectIssuersWaitBlocksSet", logs: logs, sub: sub}, nil
}

// WatchSelectIssuersWaitBlocksSet is a free log subscription operation binding the contract event 0x954fa47fa6f4e8017b99f93c73f4fbe599d786f9f5da73fe9086ab473fb455d8.
//
// Solidity: event SelectIssuersWaitBlocksSet(uint256 value)
func (_Attestations *AttestationsFilterer) WatchSelectIssuersWaitBlocksSet(opts *bind.WatchOpts, sink chan<- *AttestationsSelectIssuersWaitBlocksSet) (event.Subscription, error) {

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "SelectIssuersWaitBlocksSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsSelectIssuersWaitBlocksSet)
				if err := _Attestations.contract.UnpackLog(event, "SelectIssuersWaitBlocksSet", log); err != nil {
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

// ParseSelectIssuersWaitBlocksSet is a log parse operation binding the contract event 0x954fa47fa6f4e8017b99f93c73f4fbe599d786f9f5da73fe9086ab473fb455d8.
//
// Solidity: event SelectIssuersWaitBlocksSet(uint256 value)
func (_Attestations *AttestationsFilterer) ParseSelectIssuersWaitBlocksSet(log types.Log) (*AttestationsSelectIssuersWaitBlocksSet, error) {
	event := new(AttestationsSelectIssuersWaitBlocksSet)
	if err := _Attestations.contract.UnpackLog(event, "SelectIssuersWaitBlocksSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsTransferApprovalIterator is returned from FilterTransferApproval and is used to iterate over the raw logs and unpacked data for TransferApproval events raised by the Attestations contract.
type AttestationsTransferApprovalIterator struct {
	Event *AttestationsTransferApproval // Event containing the contract specifics and raw log

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
func (it *AttestationsTransferApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsTransferApproval)
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
		it.Event = new(AttestationsTransferApproval)
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
func (it *AttestationsTransferApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsTransferApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsTransferApproval represents a TransferApproval event raised by the Attestations contract.
type AttestationsTransferApproval struct {
	Approver    common.Address
	Indentifier [32]byte
	From        common.Address
	To          common.Address
	Approved    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferApproval is a free log retrieval operation binding the contract event 0x14d7ffb83f4265cb6fb62188eb603269555bf46efbc2923909ed7ac313d57af7.
//
// Solidity: event TransferApproval(address indexed approver, bytes32 indexed indentifier, address from, address to, bool approved)
func (_Attestations *AttestationsFilterer) FilterTransferApproval(opts *bind.FilterOpts, approver []common.Address, indentifier [][32]byte) (*AttestationsTransferApprovalIterator, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}
	var indentifierRule []interface{}
	for _, indentifierItem := range indentifier {
		indentifierRule = append(indentifierRule, indentifierItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "TransferApproval", approverRule, indentifierRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsTransferApprovalIterator{contract: _Attestations.contract, event: "TransferApproval", logs: logs, sub: sub}, nil
}

// WatchTransferApproval is a free log subscription operation binding the contract event 0x14d7ffb83f4265cb6fb62188eb603269555bf46efbc2923909ed7ac313d57af7.
//
// Solidity: event TransferApproval(address indexed approver, bytes32 indexed indentifier, address from, address to, bool approved)
func (_Attestations *AttestationsFilterer) WatchTransferApproval(opts *bind.WatchOpts, sink chan<- *AttestationsTransferApproval, approver []common.Address, indentifier [][32]byte) (event.Subscription, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}
	var indentifierRule []interface{}
	for _, indentifierItem := range indentifier {
		indentifierRule = append(indentifierRule, indentifierItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "TransferApproval", approverRule, indentifierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsTransferApproval)
				if err := _Attestations.contract.UnpackLog(event, "TransferApproval", log); err != nil {
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

// ParseTransferApproval is a log parse operation binding the contract event 0x14d7ffb83f4265cb6fb62188eb603269555bf46efbc2923909ed7ac313d57af7.
//
// Solidity: event TransferApproval(address indexed approver, bytes32 indexed indentifier, address from, address to, bool approved)
func (_Attestations *AttestationsFilterer) ParseTransferApproval(log types.Log) (*AttestationsTransferApproval, error) {
	event := new(AttestationsTransferApproval)
	if err := _Attestations.contract.UnpackLog(event, "TransferApproval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AttestationsWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Attestations contract.
type AttestationsWithdrawalIterator struct {
	Event *AttestationsWithdrawal // Event containing the contract specifics and raw log

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
func (it *AttestationsWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AttestationsWithdrawal)
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
		it.Event = new(AttestationsWithdrawal)
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
func (it *AttestationsWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AttestationsWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AttestationsWithdrawal represents a Withdrawal event raised by the Attestations contract.
type AttestationsWithdrawal struct {
	Account common.Address
	Token   common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed account, address indexed token, uint256 amount)
func (_Attestations *AttestationsFilterer) FilterWithdrawal(opts *bind.FilterOpts, account []common.Address, token []common.Address) (*AttestationsWithdrawalIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Attestations.contract.FilterLogs(opts, "Withdrawal", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &AttestationsWithdrawalIterator{contract: _Attestations.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed account, address indexed token, uint256 amount)
func (_Attestations *AttestationsFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *AttestationsWithdrawal, account []common.Address, token []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Attestations.contract.WatchLogs(opts, "Withdrawal", accountRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AttestationsWithdrawal)
				if err := _Attestations.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x2717ead6b9200dd235aad468c9809ea400fe33ac69b5bfaa6d3e90fc922b6398.
//
// Solidity: event Withdrawal(address indexed account, address indexed token, uint256 amount)
func (_Attestations *AttestationsFilterer) ParseWithdrawal(log types.Log) (*AttestationsWithdrawal, error) {
	event := new(AttestationsWithdrawal)
	if err := _Attestations.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
