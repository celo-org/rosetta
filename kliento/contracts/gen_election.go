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

// ElectionMetaData contains all meta data concerning the Election contract.
var ElectionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"AllowedToVoteOverMaxNumberOfGroups\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"electabilityThreshold\",\"type\":\"uint256\"}],\"name\":\"ElectabilityThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"ElectableValidatorsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"EpochRewardsDistributedToVoters\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxNumGroupsVotedFor\",\"type\":\"uint256\"}],\"name\":\"MaxNumGroupsVotedForSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"units\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupActiveVoteRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMarkedEligible\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMarkedIneligible\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupPendingVoteRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"units\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupVoteActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupVoteCast\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowedToVoteOverMaxNumberOfGroups\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cachedVotesByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalVotes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electabilityThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electableValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxNumGroupsVotedFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minElectableValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxElectableValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxNumGroupsVotedFor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_electabilityThreshold\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"min\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"max\",\"type\":\"uint256\"}],\"name\":\"setElectableValidators\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getElectableValidators\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxNumGroupsVotedFor\",\"type\":\"uint256\"}],\"name\":\"setMaxNumGroupsVotedFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setElectabilityThreshold\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getElectabilityThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lesser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"greater\",\"type\":\"address\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"activate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"activateForAccount\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"hasActivatablePendingVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lesser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"greater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revokePending\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lesser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"greater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revokeAllActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lesser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"greater\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"revokeActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalVotesByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"updateTotalVotesByAccountForGroup\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getPendingVotesForGroupByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getActiveVotesForGroupByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalVotesForGroupByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getActiveVoteUnitsForGroupByAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getActiveVoteUnitsForGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getTotalVotesForGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getActiveVotesForGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getPendingVotesForGroup\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getGroupEligibility\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"totalEpochRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"uptimes\",\"type\":\"uint256[]\"}],\"name\":\"getGroupEpochRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lesser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"greater\",\"type\":\"address\"}],\"name\":\"distributeEpochRewards\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"markGroupIneligible\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lesser\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"greater\",\"type\":\"address\"}],\"name\":\"markGroupEligible\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGroupsVotedForByAccount\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"canReceiveVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"group\",\"type\":\"address\"}],\"name\":\"getNumVotesReceivable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getActiveVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEligibleValidatorGroups\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalVotesForEligibleValidatorGroups\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"groups\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"electValidatorSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minElectableValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxElectableValidators\",\"type\":\"uint256\"}],\"name\":\"electNValidatorSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentValidatorSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"setAllowedToVoteOverMaxNumberOfGroups\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"lessers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"greaters\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"forceDecrementVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ElectionABI is the input ABI used to generate the binding from.
// Deprecated: Use ElectionMetaData.ABI instead.
var ElectionABI = ElectionMetaData.ABI

// Election is an auto generated Go binding around an Ethereum contract.
type Election struct {
	ElectionCaller     // Read-only binding to the contract
	ElectionTransactor // Write-only binding to the contract
	ElectionFilterer   // Log filterer for contract events
}

// ElectionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ElectionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ElectionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ElectionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ElectionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ElectionSession struct {
	Contract     *Election         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ElectionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ElectionCallerSession struct {
	Contract *ElectionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ElectionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ElectionTransactorSession struct {
	Contract     *ElectionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ElectionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ElectionRaw struct {
	Contract *Election // Generic contract binding to access the raw methods on
}

// ElectionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ElectionCallerRaw struct {
	Contract *ElectionCaller // Generic read-only contract binding to access the raw methods on
}

// ElectionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ElectionTransactorRaw struct {
	Contract *ElectionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewElection creates a new instance of Election, bound to a specific deployed contract.
func NewElection(address common.Address, backend bind.ContractBackend) (*Election, error) {
	contract, err := bindElection(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Election{ElectionCaller: ElectionCaller{contract: contract}, ElectionTransactor: ElectionTransactor{contract: contract}, ElectionFilterer: ElectionFilterer{contract: contract}}, nil
}

// NewElectionCaller creates a new read-only instance of Election, bound to a specific deployed contract.
func NewElectionCaller(address common.Address, caller bind.ContractCaller) (*ElectionCaller, error) {
	contract, err := bindElection(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionCaller{contract: contract}, nil
}

// NewElectionTransactor creates a new write-only instance of Election, bound to a specific deployed contract.
func NewElectionTransactor(address common.Address, transactor bind.ContractTransactor) (*ElectionTransactor, error) {
	contract, err := bindElection(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ElectionTransactor{contract: contract}, nil
}

// NewElectionFilterer creates a new log filterer instance of Election, bound to a specific deployed contract.
func NewElectionFilterer(address common.Address, filterer bind.ContractFilterer) (*ElectionFilterer, error) {
	contract, err := bindElection(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ElectionFilterer{contract: contract}, nil
}

// bindElection binds a generic wrapper to an already deployed contract.
func bindElection(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseElectionABI parses the ABI
func ParseElectionABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ElectionABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Election *ElectionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Election.Contract.ElectionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Election *ElectionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.Contract.ElectionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Election *ElectionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Election.Contract.ElectionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Election *ElectionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Election.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Election *ElectionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Election *ElectionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Election.Contract.contract.Transact(opts, method, params...)
}

// AllowedToVoteOverMaxNumberOfGroups is a free data retrieval call binding the contract method 0xcb08c993.
//
// Solidity: function allowedToVoteOverMaxNumberOfGroups(address ) view returns(bool)
func (_Election *ElectionCaller) AllowedToVoteOverMaxNumberOfGroups(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "allowedToVoteOverMaxNumberOfGroups", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedToVoteOverMaxNumberOfGroups is a free data retrieval call binding the contract method 0xcb08c993.
//
// Solidity: function allowedToVoteOverMaxNumberOfGroups(address ) view returns(bool)
func (_Election *ElectionSession) AllowedToVoteOverMaxNumberOfGroups(arg0 common.Address) (bool, error) {
	return _Election.Contract.AllowedToVoteOverMaxNumberOfGroups(&_Election.CallOpts, arg0)
}

// AllowedToVoteOverMaxNumberOfGroups is a free data retrieval call binding the contract method 0xcb08c993.
//
// Solidity: function allowedToVoteOverMaxNumberOfGroups(address ) view returns(bool)
func (_Election *ElectionCallerSession) AllowedToVoteOverMaxNumberOfGroups(arg0 common.Address) (bool, error) {
	return _Election.Contract.AllowedToVoteOverMaxNumberOfGroups(&_Election.CallOpts, arg0)
}

// CachedVotesByAccount is a free data retrieval call binding the contract method 0xd190d580.
//
// Solidity: function cachedVotesByAccount(address ) view returns(uint256 totalVotes)
func (_Election *ElectionCaller) CachedVotesByAccount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "cachedVotesByAccount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CachedVotesByAccount is a free data retrieval call binding the contract method 0xd190d580.
//
// Solidity: function cachedVotesByAccount(address ) view returns(uint256 totalVotes)
func (_Election *ElectionSession) CachedVotesByAccount(arg0 common.Address) (*big.Int, error) {
	return _Election.Contract.CachedVotesByAccount(&_Election.CallOpts, arg0)
}

// CachedVotesByAccount is a free data retrieval call binding the contract method 0xd190d580.
//
// Solidity: function cachedVotesByAccount(address ) view returns(uint256 totalVotes)
func (_Election *ElectionCallerSession) CachedVotesByAccount(arg0 common.Address) (*big.Int, error) {
	return _Election.Contract.CachedVotesByAccount(&_Election.CallOpts, arg0)
}

// CanReceiveVotes is a free data retrieval call binding the contract method 0xe59ea3e8.
//
// Solidity: function canReceiveVotes(address group, uint256 value) view returns(bool)
func (_Election *ElectionCaller) CanReceiveVotes(opts *bind.CallOpts, group common.Address, value *big.Int) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "canReceiveVotes", group, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanReceiveVotes is a free data retrieval call binding the contract method 0xe59ea3e8.
//
// Solidity: function canReceiveVotes(address group, uint256 value) view returns(bool)
func (_Election *ElectionSession) CanReceiveVotes(group common.Address, value *big.Int) (bool, error) {
	return _Election.Contract.CanReceiveVotes(&_Election.CallOpts, group, value)
}

// CanReceiveVotes is a free data retrieval call binding the contract method 0xe59ea3e8.
//
// Solidity: function canReceiveVotes(address group, uint256 value) view returns(bool)
func (_Election *ElectionCallerSession) CanReceiveVotes(group common.Address, value *big.Int) (bool, error) {
	return _Election.Contract.CanReceiveVotes(&_Election.CallOpts, group, value)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Election *ElectionCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Election *ElectionSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Election.Contract.CheckProofOfPossession(&_Election.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Election *ElectionCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Election.Contract.CheckProofOfPossession(&_Election.CallOpts, sender, blsKey, blsPop)
}

// ElectNValidatorSigners is a free data retrieval call binding the contract method 0x90a4dd5c.
//
// Solidity: function electNValidatorSigners(uint256 minElectableValidators, uint256 maxElectableValidators) view returns(address[])
func (_Election *ElectionCaller) ElectNValidatorSigners(opts *bind.CallOpts, minElectableValidators *big.Int, maxElectableValidators *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "electNValidatorSigners", minElectableValidators, maxElectableValidators)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ElectNValidatorSigners is a free data retrieval call binding the contract method 0x90a4dd5c.
//
// Solidity: function electNValidatorSigners(uint256 minElectableValidators, uint256 maxElectableValidators) view returns(address[])
func (_Election *ElectionSession) ElectNValidatorSigners(minElectableValidators *big.Int, maxElectableValidators *big.Int) ([]common.Address, error) {
	return _Election.Contract.ElectNValidatorSigners(&_Election.CallOpts, minElectableValidators, maxElectableValidators)
}

// ElectNValidatorSigners is a free data retrieval call binding the contract method 0x90a4dd5c.
//
// Solidity: function electNValidatorSigners(uint256 minElectableValidators, uint256 maxElectableValidators) view returns(address[])
func (_Election *ElectionCallerSession) ElectNValidatorSigners(minElectableValidators *big.Int, maxElectableValidators *big.Int) ([]common.Address, error) {
	return _Election.Contract.ElectNValidatorSigners(&_Election.CallOpts, minElectableValidators, maxElectableValidators)
}

// ElectValidatorSigners is a free data retrieval call binding the contract method 0x2ba38e69.
//
// Solidity: function electValidatorSigners() view returns(address[])
func (_Election *ElectionCaller) ElectValidatorSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "electValidatorSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// ElectValidatorSigners is a free data retrieval call binding the contract method 0x2ba38e69.
//
// Solidity: function electValidatorSigners() view returns(address[])
func (_Election *ElectionSession) ElectValidatorSigners() ([]common.Address, error) {
	return _Election.Contract.ElectValidatorSigners(&_Election.CallOpts)
}

// ElectValidatorSigners is a free data retrieval call binding the contract method 0x2ba38e69.
//
// Solidity: function electValidatorSigners() view returns(address[])
func (_Election *ElectionCallerSession) ElectValidatorSigners() ([]common.Address, error) {
	return _Election.Contract.ElectValidatorSigners(&_Election.CallOpts)
}

// ElectabilityThreshold is a free data retrieval call binding the contract method 0x4be8843b.
//
// Solidity: function electabilityThreshold() view returns(uint256 value)
func (_Election *ElectionCaller) ElectabilityThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "electabilityThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ElectabilityThreshold is a free data retrieval call binding the contract method 0x4be8843b.
//
// Solidity: function electabilityThreshold() view returns(uint256 value)
func (_Election *ElectionSession) ElectabilityThreshold() (*big.Int, error) {
	return _Election.Contract.ElectabilityThreshold(&_Election.CallOpts)
}

// ElectabilityThreshold is a free data retrieval call binding the contract method 0x4be8843b.
//
// Solidity: function electabilityThreshold() view returns(uint256 value)
func (_Election *ElectionCallerSession) ElectabilityThreshold() (*big.Int, error) {
	return _Election.Contract.ElectabilityThreshold(&_Election.CallOpts)
}

// ElectableValidators is a free data retrieval call binding the contract method 0xf9d7daae.
//
// Solidity: function electableValidators() view returns(uint256 min, uint256 max)
func (_Election *ElectionCaller) ElectableValidators(opts *bind.CallOpts) (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "electableValidators")

	outstruct := new(struct {
		Min *big.Int
		Max *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Min = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Max = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ElectableValidators is a free data retrieval call binding the contract method 0xf9d7daae.
//
// Solidity: function electableValidators() view returns(uint256 min, uint256 max)
func (_Election *ElectionSession) ElectableValidators() (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _Election.Contract.ElectableValidators(&_Election.CallOpts)
}

// ElectableValidators is a free data retrieval call binding the contract method 0xf9d7daae.
//
// Solidity: function electableValidators() view returns(uint256 min, uint256 max)
func (_Election *ElectionCallerSession) ElectableValidators() (struct {
	Min *big.Int
	Max *big.Int
}, error) {
	return _Election.Contract.ElectableValidators(&_Election.CallOpts)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Election *ElectionCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

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
func (_Election *ElectionSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Election.Contract.FractionMulExp(&_Election.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Election *ElectionCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Election.Contract.FractionMulExp(&_Election.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetActiveVoteUnitsForGroup is a free data retrieval call binding the contract method 0x5bb5acfb.
//
// Solidity: function getActiveVoteUnitsForGroup(address group) view returns(uint256)
func (_Election *ElectionCaller) GetActiveVoteUnitsForGroup(opts *bind.CallOpts, group common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getActiveVoteUnitsForGroup", group)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveVoteUnitsForGroup is a free data retrieval call binding the contract method 0x5bb5acfb.
//
// Solidity: function getActiveVoteUnitsForGroup(address group) view returns(uint256)
func (_Election *ElectionSession) GetActiveVoteUnitsForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVoteUnitsForGroup(&_Election.CallOpts, group)
}

// GetActiveVoteUnitsForGroup is a free data retrieval call binding the contract method 0x5bb5acfb.
//
// Solidity: function getActiveVoteUnitsForGroup(address group) view returns(uint256)
func (_Election *ElectionCallerSession) GetActiveVoteUnitsForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVoteUnitsForGroup(&_Election.CallOpts, group)
}

// GetActiveVoteUnitsForGroupByAccount is a free data retrieval call binding the contract method 0xa2fb4ddf.
//
// Solidity: function getActiveVoteUnitsForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCaller) GetActiveVoteUnitsForGroupByAccount(opts *bind.CallOpts, group common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getActiveVoteUnitsForGroupByAccount", group, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveVoteUnitsForGroupByAccount is a free data retrieval call binding the contract method 0xa2fb4ddf.
//
// Solidity: function getActiveVoteUnitsForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionSession) GetActiveVoteUnitsForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVoteUnitsForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetActiveVoteUnitsForGroupByAccount is a free data retrieval call binding the contract method 0xa2fb4ddf.
//
// Solidity: function getActiveVoteUnitsForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCallerSession) GetActiveVoteUnitsForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVoteUnitsForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetActiveVotes is a free data retrieval call binding the contract method 0x1f604243.
//
// Solidity: function getActiveVotes() view returns(uint256)
func (_Election *ElectionCaller) GetActiveVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getActiveVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveVotes is a free data retrieval call binding the contract method 0x1f604243.
//
// Solidity: function getActiveVotes() view returns(uint256)
func (_Election *ElectionSession) GetActiveVotes() (*big.Int, error) {
	return _Election.Contract.GetActiveVotes(&_Election.CallOpts)
}

// GetActiveVotes is a free data retrieval call binding the contract method 0x1f604243.
//
// Solidity: function getActiveVotes() view returns(uint256)
func (_Election *ElectionCallerSession) GetActiveVotes() (*big.Int, error) {
	return _Election.Contract.GetActiveVotes(&_Election.CallOpts)
}

// GetActiveVotesForGroup is a free data retrieval call binding the contract method 0x926d00ca.
//
// Solidity: function getActiveVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionCaller) GetActiveVotesForGroup(opts *bind.CallOpts, group common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getActiveVotesForGroup", group)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveVotesForGroup is a free data retrieval call binding the contract method 0x926d00ca.
//
// Solidity: function getActiveVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionSession) GetActiveVotesForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVotesForGroup(&_Election.CallOpts, group)
}

// GetActiveVotesForGroup is a free data retrieval call binding the contract method 0x926d00ca.
//
// Solidity: function getActiveVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionCallerSession) GetActiveVotesForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVotesForGroup(&_Election.CallOpts, group)
}

// GetActiveVotesForGroupByAccount is a free data retrieval call binding the contract method 0xd3e242a4.
//
// Solidity: function getActiveVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCaller) GetActiveVotesForGroupByAccount(opts *bind.CallOpts, group common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getActiveVotesForGroupByAccount", group, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveVotesForGroupByAccount is a free data retrieval call binding the contract method 0xd3e242a4.
//
// Solidity: function getActiveVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionSession) GetActiveVotesForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVotesForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetActiveVotesForGroupByAccount is a free data retrieval call binding the contract method 0xd3e242a4.
//
// Solidity: function getActiveVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCallerSession) GetActiveVotesForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetActiveVotesForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Election *ElectionCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Election *ElectionSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Election.Contract.GetBlockNumberFromHeader(&_Election.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Election *ElectionCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Election.Contract.GetBlockNumberFromHeader(&_Election.CallOpts, header)
}

// GetCurrentValidatorSigners is a free data retrieval call binding the contract method 0x448144c8.
//
// Solidity: function getCurrentValidatorSigners() view returns(address[])
func (_Election *ElectionCaller) GetCurrentValidatorSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getCurrentValidatorSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentValidatorSigners is a free data retrieval call binding the contract method 0x448144c8.
//
// Solidity: function getCurrentValidatorSigners() view returns(address[])
func (_Election *ElectionSession) GetCurrentValidatorSigners() ([]common.Address, error) {
	return _Election.Contract.GetCurrentValidatorSigners(&_Election.CallOpts)
}

// GetCurrentValidatorSigners is a free data retrieval call binding the contract method 0x448144c8.
//
// Solidity: function getCurrentValidatorSigners() view returns(address[])
func (_Election *ElectionCallerSession) GetCurrentValidatorSigners() ([]common.Address, error) {
	return _Election.Contract.GetCurrentValidatorSigners(&_Election.CallOpts)
}

// GetElectabilityThreshold is a free data retrieval call binding the contract method 0xbdd14318.
//
// Solidity: function getElectabilityThreshold() view returns(uint256)
func (_Election *ElectionCaller) GetElectabilityThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getElectabilityThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetElectabilityThreshold is a free data retrieval call binding the contract method 0xbdd14318.
//
// Solidity: function getElectabilityThreshold() view returns(uint256)
func (_Election *ElectionSession) GetElectabilityThreshold() (*big.Int, error) {
	return _Election.Contract.GetElectabilityThreshold(&_Election.CallOpts)
}

// GetElectabilityThreshold is a free data retrieval call binding the contract method 0xbdd14318.
//
// Solidity: function getElectabilityThreshold() view returns(uint256)
func (_Election *ElectionCallerSession) GetElectabilityThreshold() (*big.Int, error) {
	return _Election.Contract.GetElectabilityThreshold(&_Election.CallOpts)
}

// GetElectableValidators is a free data retrieval call binding the contract method 0xf9f41a7a.
//
// Solidity: function getElectableValidators() view returns(uint256, uint256)
func (_Election *ElectionCaller) GetElectableValidators(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getElectableValidators")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetElectableValidators is a free data retrieval call binding the contract method 0xf9f41a7a.
//
// Solidity: function getElectableValidators() view returns(uint256, uint256)
func (_Election *ElectionSession) GetElectableValidators() (*big.Int, *big.Int, error) {
	return _Election.Contract.GetElectableValidators(&_Election.CallOpts)
}

// GetElectableValidators is a free data retrieval call binding the contract method 0xf9f41a7a.
//
// Solidity: function getElectableValidators() view returns(uint256, uint256)
func (_Election *ElectionCallerSession) GetElectableValidators() (*big.Int, *big.Int, error) {
	return _Election.Contract.GetElectableValidators(&_Election.CallOpts)
}

// GetEligibleValidatorGroups is a free data retrieval call binding the contract method 0xa5826ab2.
//
// Solidity: function getEligibleValidatorGroups() view returns(address[])
func (_Election *ElectionCaller) GetEligibleValidatorGroups(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getEligibleValidatorGroups")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetEligibleValidatorGroups is a free data retrieval call binding the contract method 0xa5826ab2.
//
// Solidity: function getEligibleValidatorGroups() view returns(address[])
func (_Election *ElectionSession) GetEligibleValidatorGroups() ([]common.Address, error) {
	return _Election.Contract.GetEligibleValidatorGroups(&_Election.CallOpts)
}

// GetEligibleValidatorGroups is a free data retrieval call binding the contract method 0xa5826ab2.
//
// Solidity: function getEligibleValidatorGroups() view returns(address[])
func (_Election *ElectionCallerSession) GetEligibleValidatorGroups() ([]common.Address, error) {
	return _Election.Contract.GetEligibleValidatorGroups(&_Election.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Election *ElectionCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Election *ElectionSession) GetEpochNumber() (*big.Int, error) {
	return _Election.Contract.GetEpochNumber(&_Election.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Election *ElectionCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Election.Contract.GetEpochNumber(&_Election.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Election.Contract.GetEpochNumberOfBlock(&_Election.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Election.Contract.GetEpochNumberOfBlock(&_Election.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Election *ElectionCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Election *ElectionSession) GetEpochSize() (*big.Int, error) {
	return _Election.Contract.GetEpochSize(&_Election.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Election *ElectionCallerSession) GetEpochSize() (*big.Int, error) {
	return _Election.Contract.GetEpochSize(&_Election.CallOpts)
}

// GetGroupEligibility is a free data retrieval call binding the contract method 0x8c666775.
//
// Solidity: function getGroupEligibility(address group) view returns(bool)
func (_Election *ElectionCaller) GetGroupEligibility(opts *bind.CallOpts, group common.Address) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getGroupEligibility", group)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetGroupEligibility is a free data retrieval call binding the contract method 0x8c666775.
//
// Solidity: function getGroupEligibility(address group) view returns(bool)
func (_Election *ElectionSession) GetGroupEligibility(group common.Address) (bool, error) {
	return _Election.Contract.GetGroupEligibility(&_Election.CallOpts, group)
}

// GetGroupEligibility is a free data retrieval call binding the contract method 0x8c666775.
//
// Solidity: function getGroupEligibility(address group) view returns(bool)
func (_Election *ElectionCallerSession) GetGroupEligibility(group common.Address) (bool, error) {
	return _Election.Contract.GetGroupEligibility(&_Election.CallOpts, group)
}

// GetGroupEpochRewards is a free data retrieval call binding the contract method 0xf23263f9.
//
// Solidity: function getGroupEpochRewards(address group, uint256 totalEpochRewards, uint256[] uptimes) view returns(uint256)
func (_Election *ElectionCaller) GetGroupEpochRewards(opts *bind.CallOpts, group common.Address, totalEpochRewards *big.Int, uptimes []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getGroupEpochRewards", group, totalEpochRewards, uptimes)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGroupEpochRewards is a free data retrieval call binding the contract method 0xf23263f9.
//
// Solidity: function getGroupEpochRewards(address group, uint256 totalEpochRewards, uint256[] uptimes) view returns(uint256)
func (_Election *ElectionSession) GetGroupEpochRewards(group common.Address, totalEpochRewards *big.Int, uptimes []*big.Int) (*big.Int, error) {
	return _Election.Contract.GetGroupEpochRewards(&_Election.CallOpts, group, totalEpochRewards, uptimes)
}

// GetGroupEpochRewards is a free data retrieval call binding the contract method 0xf23263f9.
//
// Solidity: function getGroupEpochRewards(address group, uint256 totalEpochRewards, uint256[] uptimes) view returns(uint256)
func (_Election *ElectionCallerSession) GetGroupEpochRewards(group common.Address, totalEpochRewards *big.Int, uptimes []*big.Int) (*big.Int, error) {
	return _Election.Contract.GetGroupEpochRewards(&_Election.CallOpts, group, totalEpochRewards, uptimes)
}

// GetGroupsVotedForByAccount is a free data retrieval call binding the contract method 0x457578a3.
//
// Solidity: function getGroupsVotedForByAccount(address account) view returns(address[])
func (_Election *ElectionCaller) GetGroupsVotedForByAccount(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getGroupsVotedForByAccount", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetGroupsVotedForByAccount is a free data retrieval call binding the contract method 0x457578a3.
//
// Solidity: function getGroupsVotedForByAccount(address account) view returns(address[])
func (_Election *ElectionSession) GetGroupsVotedForByAccount(account common.Address) ([]common.Address, error) {
	return _Election.Contract.GetGroupsVotedForByAccount(&_Election.CallOpts, account)
}

// GetGroupsVotedForByAccount is a free data retrieval call binding the contract method 0x457578a3.
//
// Solidity: function getGroupsVotedForByAccount(address account) view returns(address[])
func (_Election *ElectionCallerSession) GetGroupsVotedForByAccount(account common.Address) ([]common.Address, error) {
	return _Election.Contract.GetGroupsVotedForByAccount(&_Election.CallOpts, account)
}

// GetNumVotesReceivable is a free data retrieval call binding the contract method 0x2c3b7916.
//
// Solidity: function getNumVotesReceivable(address group) view returns(uint256)
func (_Election *ElectionCaller) GetNumVotesReceivable(opts *bind.CallOpts, group common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getNumVotesReceivable", group)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumVotesReceivable is a free data retrieval call binding the contract method 0x2c3b7916.
//
// Solidity: function getNumVotesReceivable(address group) view returns(uint256)
func (_Election *ElectionSession) GetNumVotesReceivable(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetNumVotesReceivable(&_Election.CallOpts, group)
}

// GetNumVotesReceivable is a free data retrieval call binding the contract method 0x2c3b7916.
//
// Solidity: function getNumVotesReceivable(address group) view returns(uint256)
func (_Election *ElectionCallerSession) GetNumVotesReceivable(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetNumVotesReceivable(&_Election.CallOpts, group)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Election *ElectionCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Election *ElectionSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Election.Contract.GetParentSealBitmap(&_Election.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Election *ElectionCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Election.Contract.GetParentSealBitmap(&_Election.CallOpts, blockNumber)
}

// GetPendingVotesForGroup is a free data retrieval call binding the contract method 0x95128ce3.
//
// Solidity: function getPendingVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionCaller) GetPendingVotesForGroup(opts *bind.CallOpts, group common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getPendingVotesForGroup", group)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingVotesForGroup is a free data retrieval call binding the contract method 0x95128ce3.
//
// Solidity: function getPendingVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionSession) GetPendingVotesForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetPendingVotesForGroup(&_Election.CallOpts, group)
}

// GetPendingVotesForGroup is a free data retrieval call binding the contract method 0x95128ce3.
//
// Solidity: function getPendingVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionCallerSession) GetPendingVotesForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetPendingVotesForGroup(&_Election.CallOpts, group)
}

// GetPendingVotesForGroupByAccount is a free data retrieval call binding the contract method 0x9b95975f.
//
// Solidity: function getPendingVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCaller) GetPendingVotesForGroupByAccount(opts *bind.CallOpts, group common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getPendingVotesForGroupByAccount", group, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPendingVotesForGroupByAccount is a free data retrieval call binding the contract method 0x9b95975f.
//
// Solidity: function getPendingVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionSession) GetPendingVotesForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetPendingVotesForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetPendingVotesForGroupByAccount is a free data retrieval call binding the contract method 0x9b95975f.
//
// Solidity: function getPendingVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCallerSession) GetPendingVotesForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetPendingVotesForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetTotalVotes is a free data retrieval call binding the contract method 0x9a0e7d66.
//
// Solidity: function getTotalVotes() view returns(uint256)
func (_Election *ElectionCaller) GetTotalVotes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getTotalVotes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotes is a free data retrieval call binding the contract method 0x9a0e7d66.
//
// Solidity: function getTotalVotes() view returns(uint256)
func (_Election *ElectionSession) GetTotalVotes() (*big.Int, error) {
	return _Election.Contract.GetTotalVotes(&_Election.CallOpts)
}

// GetTotalVotes is a free data retrieval call binding the contract method 0x9a0e7d66.
//
// Solidity: function getTotalVotes() view returns(uint256)
func (_Election *ElectionCallerSession) GetTotalVotes() (*big.Int, error) {
	return _Election.Contract.GetTotalVotes(&_Election.CallOpts)
}

// GetTotalVotesByAccount is a free data retrieval call binding the contract method 0x6c781a2c.
//
// Solidity: function getTotalVotesByAccount(address account) view returns(uint256)
func (_Election *ElectionCaller) GetTotalVotesByAccount(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getTotalVotesByAccount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotesByAccount is a free data retrieval call binding the contract method 0x6c781a2c.
//
// Solidity: function getTotalVotesByAccount(address account) view returns(uint256)
func (_Election *ElectionSession) GetTotalVotesByAccount(account common.Address) (*big.Int, error) {
	return _Election.Contract.GetTotalVotesByAccount(&_Election.CallOpts, account)
}

// GetTotalVotesByAccount is a free data retrieval call binding the contract method 0x6c781a2c.
//
// Solidity: function getTotalVotesByAccount(address account) view returns(uint256)
func (_Election *ElectionCallerSession) GetTotalVotesByAccount(account common.Address) (*big.Int, error) {
	return _Election.Contract.GetTotalVotesByAccount(&_Election.CallOpts, account)
}

// GetTotalVotesForEligibleValidatorGroups is a free data retrieval call binding the contract method 0x7046c96b.
//
// Solidity: function getTotalVotesForEligibleValidatorGroups() view returns(address[] groups, uint256[] values)
func (_Election *ElectionCaller) GetTotalVotesForEligibleValidatorGroups(opts *bind.CallOpts) (struct {
	Groups []common.Address
	Values []*big.Int
}, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getTotalVotesForEligibleValidatorGroups")

	outstruct := new(struct {
		Groups []common.Address
		Values []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Groups = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Values = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetTotalVotesForEligibleValidatorGroups is a free data retrieval call binding the contract method 0x7046c96b.
//
// Solidity: function getTotalVotesForEligibleValidatorGroups() view returns(address[] groups, uint256[] values)
func (_Election *ElectionSession) GetTotalVotesForEligibleValidatorGroups() (struct {
	Groups []common.Address
	Values []*big.Int
}, error) {
	return _Election.Contract.GetTotalVotesForEligibleValidatorGroups(&_Election.CallOpts)
}

// GetTotalVotesForEligibleValidatorGroups is a free data retrieval call binding the contract method 0x7046c96b.
//
// Solidity: function getTotalVotesForEligibleValidatorGroups() view returns(address[] groups, uint256[] values)
func (_Election *ElectionCallerSession) GetTotalVotesForEligibleValidatorGroups() (struct {
	Groups []common.Address
	Values []*big.Int
}, error) {
	return _Election.Contract.GetTotalVotesForEligibleValidatorGroups(&_Election.CallOpts)
}

// GetTotalVotesForGroup is a free data retrieval call binding the contract method 0xdedafeae.
//
// Solidity: function getTotalVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionCaller) GetTotalVotesForGroup(opts *bind.CallOpts, group common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getTotalVotesForGroup", group)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotesForGroup is a free data retrieval call binding the contract method 0xdedafeae.
//
// Solidity: function getTotalVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionSession) GetTotalVotesForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetTotalVotesForGroup(&_Election.CallOpts, group)
}

// GetTotalVotesForGroup is a free data retrieval call binding the contract method 0xdedafeae.
//
// Solidity: function getTotalVotesForGroup(address group) view returns(uint256)
func (_Election *ElectionCallerSession) GetTotalVotesForGroup(group common.Address) (*big.Int, error) {
	return _Election.Contract.GetTotalVotesForGroup(&_Election.CallOpts, group)
}

// GetTotalVotesForGroupByAccount is a free data retrieval call binding the contract method 0x38617272.
//
// Solidity: function getTotalVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCaller) GetTotalVotesForGroupByAccount(opts *bind.CallOpts, group common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getTotalVotesForGroupByAccount", group, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalVotesForGroupByAccount is a free data retrieval call binding the contract method 0x38617272.
//
// Solidity: function getTotalVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionSession) GetTotalVotesForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetTotalVotesForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetTotalVotesForGroupByAccount is a free data retrieval call binding the contract method 0x38617272.
//
// Solidity: function getTotalVotesForGroupByAccount(address group, address account) view returns(uint256)
func (_Election *ElectionCallerSession) GetTotalVotesForGroupByAccount(group common.Address, account common.Address) (*big.Int, error) {
	return _Election.Contract.GetTotalVotesForGroupByAccount(&_Election.CallOpts, group, account)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Election *ElectionCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Election *ElectionSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Election.Contract.GetVerifiedSealBitmapFromHeader(&_Election.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Election *ElectionCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Election.Contract.GetVerifiedSealBitmapFromHeader(&_Election.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Election *ElectionCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "getVersionNumber")

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
func (_Election *ElectionSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Election.Contract.GetVersionNumber(&_Election.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Election *ElectionCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Election.Contract.GetVersionNumber(&_Election.CallOpts)
}

// HasActivatablePendingVotes is a free data retrieval call binding the contract method 0x263ecf74.
//
// Solidity: function hasActivatablePendingVotes(address account, address group) view returns(bool)
func (_Election *ElectionCaller) HasActivatablePendingVotes(opts *bind.CallOpts, account common.Address, group common.Address) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "hasActivatablePendingVotes", account, group)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasActivatablePendingVotes is a free data retrieval call binding the contract method 0x263ecf74.
//
// Solidity: function hasActivatablePendingVotes(address account, address group) view returns(bool)
func (_Election *ElectionSession) HasActivatablePendingVotes(account common.Address, group common.Address) (bool, error) {
	return _Election.Contract.HasActivatablePendingVotes(&_Election.CallOpts, account, group)
}

// HasActivatablePendingVotes is a free data retrieval call binding the contract method 0x263ecf74.
//
// Solidity: function hasActivatablePendingVotes(address account, address group) view returns(bool)
func (_Election *ElectionCallerSession) HasActivatablePendingVotes(account common.Address, group common.Address) (bool, error) {
	return _Election.Contract.HasActivatablePendingVotes(&_Election.CallOpts, account, group)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Election *ElectionCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Election *ElectionSession) HashHeader(header []byte) ([32]byte, error) {
	return _Election.Contract.HashHeader(&_Election.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Election *ElectionCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Election.Contract.HashHeader(&_Election.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Election *ElectionCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Election *ElectionSession) Initialized() (bool, error) {
	return _Election.Contract.Initialized(&_Election.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Election *ElectionCallerSession) Initialized() (bool, error) {
	return _Election.Contract.Initialized(&_Election.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Election *ElectionCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Election *ElectionSession) IsOwner() (bool, error) {
	return _Election.Contract.IsOwner(&_Election.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Election *ElectionCallerSession) IsOwner() (bool, error) {
	return _Election.Contract.IsOwner(&_Election.CallOpts)
}

// MaxNumGroupsVotedFor is a free data retrieval call binding the contract method 0xac839d69.
//
// Solidity: function maxNumGroupsVotedFor() view returns(uint256)
func (_Election *ElectionCaller) MaxNumGroupsVotedFor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "maxNumGroupsVotedFor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNumGroupsVotedFor is a free data retrieval call binding the contract method 0xac839d69.
//
// Solidity: function maxNumGroupsVotedFor() view returns(uint256)
func (_Election *ElectionSession) MaxNumGroupsVotedFor() (*big.Int, error) {
	return _Election.Contract.MaxNumGroupsVotedFor(&_Election.CallOpts)
}

// MaxNumGroupsVotedFor is a free data retrieval call binding the contract method 0xac839d69.
//
// Solidity: function maxNumGroupsVotedFor() view returns(uint256)
func (_Election *ElectionCallerSession) MaxNumGroupsVotedFor() (*big.Int, error) {
	return _Election.Contract.MaxNumGroupsVotedFor(&_Election.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Election.Contract.MinQuorumSize(&_Election.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Election.Contract.MinQuorumSize(&_Election.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Election *ElectionCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Election *ElectionSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Election.Contract.MinQuorumSizeInCurrentSet(&_Election.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Election *ElectionCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Election.Contract.MinQuorumSizeInCurrentSet(&_Election.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Election *ElectionCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Election *ElectionSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Election.Contract.NumberValidatorsInCurrentSet(&_Election.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Election *ElectionCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Election.Contract.NumberValidatorsInCurrentSet(&_Election.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Election.Contract.NumberValidatorsInSet(&_Election.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Election *ElectionCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Election.Contract.NumberValidatorsInSet(&_Election.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Election *ElectionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Election *ElectionSession) Owner() (common.Address, error) {
	return _Election.Contract.Owner(&_Election.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Election *ElectionCallerSession) Owner() (common.Address, error) {
	return _Election.Contract.Owner(&_Election.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Election *ElectionCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Election *ElectionSession) Registry() (common.Address, error) {
	return _Election.Contract.Registry(&_Election.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Election *ElectionCallerSession) Registry() (common.Address, error) {
	return _Election.Contract.Registry(&_Election.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Election *ElectionCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Election *ElectionSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Election.Contract.ValidatorSignerAddressFromCurrentSet(&_Election.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Election *ElectionCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Election.Contract.ValidatorSignerAddressFromCurrentSet(&_Election.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Election *ElectionCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Election.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Election *ElectionSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Election.Contract.ValidatorSignerAddressFromSet(&_Election.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Election *ElectionCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Election.Contract.ValidatorSignerAddressFromSet(&_Election.CallOpts, index, blockNumber)
}

// Activate is a paid mutator transaction binding the contract method 0x1c5a9d9c.
//
// Solidity: function activate(address group) returns(bool)
func (_Election *ElectionTransactor) Activate(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "activate", group)
}

// Activate is a paid mutator transaction binding the contract method 0x1c5a9d9c.
//
// Solidity: function activate(address group) returns(bool)
func (_Election *ElectionSession) Activate(group common.Address) (*types.Transaction, error) {
	return _Election.Contract.Activate(&_Election.TransactOpts, group)
}

// Activate is a paid mutator transaction binding the contract method 0x1c5a9d9c.
//
// Solidity: function activate(address group) returns(bool)
func (_Election *ElectionTransactorSession) Activate(group common.Address) (*types.Transaction, error) {
	return _Election.Contract.Activate(&_Election.TransactOpts, group)
}

// ActivateForAccount is a paid mutator transaction binding the contract method 0xc14470c4.
//
// Solidity: function activateForAccount(address group, address account) returns(bool)
func (_Election *ElectionTransactor) ActivateForAccount(opts *bind.TransactOpts, group common.Address, account common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "activateForAccount", group, account)
}

// ActivateForAccount is a paid mutator transaction binding the contract method 0xc14470c4.
//
// Solidity: function activateForAccount(address group, address account) returns(bool)
func (_Election *ElectionSession) ActivateForAccount(group common.Address, account common.Address) (*types.Transaction, error) {
	return _Election.Contract.ActivateForAccount(&_Election.TransactOpts, group, account)
}

// ActivateForAccount is a paid mutator transaction binding the contract method 0xc14470c4.
//
// Solidity: function activateForAccount(address group, address account) returns(bool)
func (_Election *ElectionTransactorSession) ActivateForAccount(group common.Address, account common.Address) (*types.Transaction, error) {
	return _Election.Contract.ActivateForAccount(&_Election.TransactOpts, group, account)
}

// DistributeEpochRewards is a paid mutator transaction binding the contract method 0x12541a6b.
//
// Solidity: function distributeEpochRewards(address group, uint256 value, address lesser, address greater) returns()
func (_Election *ElectionTransactor) DistributeEpochRewards(opts *bind.TransactOpts, group common.Address, value *big.Int, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "distributeEpochRewards", group, value, lesser, greater)
}

// DistributeEpochRewards is a paid mutator transaction binding the contract method 0x12541a6b.
//
// Solidity: function distributeEpochRewards(address group, uint256 value, address lesser, address greater) returns()
func (_Election *ElectionSession) DistributeEpochRewards(group common.Address, value *big.Int, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.Contract.DistributeEpochRewards(&_Election.TransactOpts, group, value, lesser, greater)
}

// DistributeEpochRewards is a paid mutator transaction binding the contract method 0x12541a6b.
//
// Solidity: function distributeEpochRewards(address group, uint256 value, address lesser, address greater) returns()
func (_Election *ElectionTransactorSession) DistributeEpochRewards(group common.Address, value *big.Int, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.Contract.DistributeEpochRewards(&_Election.TransactOpts, group, value, lesser, greater)
}

// ForceDecrementVotes is a paid mutator transaction binding the contract method 0x8ef01def.
//
// Solidity: function forceDecrementVotes(address account, uint256 value, address[] lessers, address[] greaters, uint256[] indices) returns(uint256)
func (_Election *ElectionTransactor) ForceDecrementVotes(opts *bind.TransactOpts, account common.Address, value *big.Int, lessers []common.Address, greaters []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "forceDecrementVotes", account, value, lessers, greaters, indices)
}

// ForceDecrementVotes is a paid mutator transaction binding the contract method 0x8ef01def.
//
// Solidity: function forceDecrementVotes(address account, uint256 value, address[] lessers, address[] greaters, uint256[] indices) returns(uint256)
func (_Election *ElectionSession) ForceDecrementVotes(account common.Address, value *big.Int, lessers []common.Address, greaters []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _Election.Contract.ForceDecrementVotes(&_Election.TransactOpts, account, value, lessers, greaters, indices)
}

// ForceDecrementVotes is a paid mutator transaction binding the contract method 0x8ef01def.
//
// Solidity: function forceDecrementVotes(address account, uint256 value, address[] lessers, address[] greaters, uint256[] indices) returns(uint256)
func (_Election *ElectionTransactorSession) ForceDecrementVotes(account common.Address, value *big.Int, lessers []common.Address, greaters []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _Election.Contract.ForceDecrementVotes(&_Election.TransactOpts, account, value, lessers, greaters, indices)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address registryAddress, uint256 minElectableValidators, uint256 maxElectableValidators, uint256 _maxNumGroupsVotedFor, uint256 _electabilityThreshold) returns()
func (_Election *ElectionTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, minElectableValidators *big.Int, maxElectableValidators *big.Int, _maxNumGroupsVotedFor *big.Int, _electabilityThreshold *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "initialize", registryAddress, minElectableValidators, maxElectableValidators, _maxNumGroupsVotedFor, _electabilityThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address registryAddress, uint256 minElectableValidators, uint256 maxElectableValidators, uint256 _maxNumGroupsVotedFor, uint256 _electabilityThreshold) returns()
func (_Election *ElectionSession) Initialize(registryAddress common.Address, minElectableValidators *big.Int, maxElectableValidators *big.Int, _maxNumGroupsVotedFor *big.Int, _electabilityThreshold *big.Int) (*types.Transaction, error) {
	return _Election.Contract.Initialize(&_Election.TransactOpts, registryAddress, minElectableValidators, maxElectableValidators, _maxNumGroupsVotedFor, _electabilityThreshold)
}

// Initialize is a paid mutator transaction binding the contract method 0xf92ad219.
//
// Solidity: function initialize(address registryAddress, uint256 minElectableValidators, uint256 maxElectableValidators, uint256 _maxNumGroupsVotedFor, uint256 _electabilityThreshold) returns()
func (_Election *ElectionTransactorSession) Initialize(registryAddress common.Address, minElectableValidators *big.Int, maxElectableValidators *big.Int, _maxNumGroupsVotedFor *big.Int, _electabilityThreshold *big.Int) (*types.Transaction, error) {
	return _Election.Contract.Initialize(&_Election.TransactOpts, registryAddress, minElectableValidators, maxElectableValidators, _maxNumGroupsVotedFor, _electabilityThreshold)
}

// MarkGroupEligible is a paid mutator transaction binding the contract method 0xa18fb2db.
//
// Solidity: function markGroupEligible(address group, address lesser, address greater) returns()
func (_Election *ElectionTransactor) MarkGroupEligible(opts *bind.TransactOpts, group common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "markGroupEligible", group, lesser, greater)
}

// MarkGroupEligible is a paid mutator transaction binding the contract method 0xa18fb2db.
//
// Solidity: function markGroupEligible(address group, address lesser, address greater) returns()
func (_Election *ElectionSession) MarkGroupEligible(group common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.Contract.MarkGroupEligible(&_Election.TransactOpts, group, lesser, greater)
}

// MarkGroupEligible is a paid mutator transaction binding the contract method 0xa18fb2db.
//
// Solidity: function markGroupEligible(address group, address lesser, address greater) returns()
func (_Election *ElectionTransactorSession) MarkGroupEligible(group common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.Contract.MarkGroupEligible(&_Election.TransactOpts, group, lesser, greater)
}

// MarkGroupIneligible is a paid mutator transaction binding the contract method 0xa8e45871.
//
// Solidity: function markGroupIneligible(address group) returns()
func (_Election *ElectionTransactor) MarkGroupIneligible(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "markGroupIneligible", group)
}

// MarkGroupIneligible is a paid mutator transaction binding the contract method 0xa8e45871.
//
// Solidity: function markGroupIneligible(address group) returns()
func (_Election *ElectionSession) MarkGroupIneligible(group common.Address) (*types.Transaction, error) {
	return _Election.Contract.MarkGroupIneligible(&_Election.TransactOpts, group)
}

// MarkGroupIneligible is a paid mutator transaction binding the contract method 0xa8e45871.
//
// Solidity: function markGroupIneligible(address group) returns()
func (_Election *ElectionTransactorSession) MarkGroupIneligible(group common.Address) (*types.Transaction, error) {
	return _Election.Contract.MarkGroupIneligible(&_Election.TransactOpts, group)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Election *ElectionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Election *ElectionSession) RenounceOwnership() (*types.Transaction, error) {
	return _Election.Contract.RenounceOwnership(&_Election.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Election *ElectionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Election.Contract.RenounceOwnership(&_Election.TransactOpts)
}

// RevokeActive is a paid mutator transaction binding the contract method 0x6e198475.
//
// Solidity: function revokeActive(address group, uint256 value, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionTransactor) RevokeActive(opts *bind.TransactOpts, group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "revokeActive", group, value, lesser, greater, index)
}

// RevokeActive is a paid mutator transaction binding the contract method 0x6e198475.
//
// Solidity: function revokeActive(address group, uint256 value, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionSession) RevokeActive(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.Contract.RevokeActive(&_Election.TransactOpts, group, value, lesser, greater, index)
}

// RevokeActive is a paid mutator transaction binding the contract method 0x6e198475.
//
// Solidity: function revokeActive(address group, uint256 value, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionTransactorSession) RevokeActive(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.Contract.RevokeActive(&_Election.TransactOpts, group, value, lesser, greater, index)
}

// RevokeAllActive is a paid mutator transaction binding the contract method 0xe0a2ab52.
//
// Solidity: function revokeAllActive(address group, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionTransactor) RevokeAllActive(opts *bind.TransactOpts, group common.Address, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "revokeAllActive", group, lesser, greater, index)
}

// RevokeAllActive is a paid mutator transaction binding the contract method 0xe0a2ab52.
//
// Solidity: function revokeAllActive(address group, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionSession) RevokeAllActive(group common.Address, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.Contract.RevokeAllActive(&_Election.TransactOpts, group, lesser, greater, index)
}

// RevokeAllActive is a paid mutator transaction binding the contract method 0xe0a2ab52.
//
// Solidity: function revokeAllActive(address group, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionTransactorSession) RevokeAllActive(group common.Address, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.Contract.RevokeAllActive(&_Election.TransactOpts, group, lesser, greater, index)
}

// RevokePending is a paid mutator transaction binding the contract method 0x9dfb6081.
//
// Solidity: function revokePending(address group, uint256 value, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionTransactor) RevokePending(opts *bind.TransactOpts, group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "revokePending", group, value, lesser, greater, index)
}

// RevokePending is a paid mutator transaction binding the contract method 0x9dfb6081.
//
// Solidity: function revokePending(address group, uint256 value, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionSession) RevokePending(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.Contract.RevokePending(&_Election.TransactOpts, group, value, lesser, greater, index)
}

// RevokePending is a paid mutator transaction binding the contract method 0x9dfb6081.
//
// Solidity: function revokePending(address group, uint256 value, address lesser, address greater, uint256 index) returns(bool)
func (_Election *ElectionTransactorSession) RevokePending(group common.Address, value *big.Int, lesser common.Address, greater common.Address, index *big.Int) (*types.Transaction, error) {
	return _Election.Contract.RevokePending(&_Election.TransactOpts, group, value, lesser, greater, index)
}

// SetAllowedToVoteOverMaxNumberOfGroups is a paid mutator transaction binding the contract method 0xd52758aa.
//
// Solidity: function setAllowedToVoteOverMaxNumberOfGroups(bool flag) returns()
func (_Election *ElectionTransactor) SetAllowedToVoteOverMaxNumberOfGroups(opts *bind.TransactOpts, flag bool) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "setAllowedToVoteOverMaxNumberOfGroups", flag)
}

// SetAllowedToVoteOverMaxNumberOfGroups is a paid mutator transaction binding the contract method 0xd52758aa.
//
// Solidity: function setAllowedToVoteOverMaxNumberOfGroups(bool flag) returns()
func (_Election *ElectionSession) SetAllowedToVoteOverMaxNumberOfGroups(flag bool) (*types.Transaction, error) {
	return _Election.Contract.SetAllowedToVoteOverMaxNumberOfGroups(&_Election.TransactOpts, flag)
}

// SetAllowedToVoteOverMaxNumberOfGroups is a paid mutator transaction binding the contract method 0xd52758aa.
//
// Solidity: function setAllowedToVoteOverMaxNumberOfGroups(bool flag) returns()
func (_Election *ElectionTransactorSession) SetAllowedToVoteOverMaxNumberOfGroups(flag bool) (*types.Transaction, error) {
	return _Election.Contract.SetAllowedToVoteOverMaxNumberOfGroups(&_Election.TransactOpts, flag)
}

// SetElectabilityThreshold is a paid mutator transaction binding the contract method 0x631db7e7.
//
// Solidity: function setElectabilityThreshold(uint256 threshold) returns(bool)
func (_Election *ElectionTransactor) SetElectabilityThreshold(opts *bind.TransactOpts, threshold *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "setElectabilityThreshold", threshold)
}

// SetElectabilityThreshold is a paid mutator transaction binding the contract method 0x631db7e7.
//
// Solidity: function setElectabilityThreshold(uint256 threshold) returns(bool)
func (_Election *ElectionSession) SetElectabilityThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Election.Contract.SetElectabilityThreshold(&_Election.TransactOpts, threshold)
}

// SetElectabilityThreshold is a paid mutator transaction binding the contract method 0x631db7e7.
//
// Solidity: function setElectabilityThreshold(uint256 threshold) returns(bool)
func (_Election *ElectionTransactorSession) SetElectabilityThreshold(threshold *big.Int) (*types.Transaction, error) {
	return _Election.Contract.SetElectabilityThreshold(&_Election.TransactOpts, threshold)
}

// SetElectableValidators is a paid mutator transaction binding the contract method 0xf911f0b7.
//
// Solidity: function setElectableValidators(uint256 min, uint256 max) returns(bool)
func (_Election *ElectionTransactor) SetElectableValidators(opts *bind.TransactOpts, min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "setElectableValidators", min, max)
}

// SetElectableValidators is a paid mutator transaction binding the contract method 0xf911f0b7.
//
// Solidity: function setElectableValidators(uint256 min, uint256 max) returns(bool)
func (_Election *ElectionSession) SetElectableValidators(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Election.Contract.SetElectableValidators(&_Election.TransactOpts, min, max)
}

// SetElectableValidators is a paid mutator transaction binding the contract method 0xf911f0b7.
//
// Solidity: function setElectableValidators(uint256 min, uint256 max) returns(bool)
func (_Election *ElectionTransactorSession) SetElectableValidators(min *big.Int, max *big.Int) (*types.Transaction, error) {
	return _Election.Contract.SetElectableValidators(&_Election.TransactOpts, min, max)
}

// SetMaxNumGroupsVotedFor is a paid mutator transaction binding the contract method 0x3c55a73c.
//
// Solidity: function setMaxNumGroupsVotedFor(uint256 _maxNumGroupsVotedFor) returns(bool)
func (_Election *ElectionTransactor) SetMaxNumGroupsVotedFor(opts *bind.TransactOpts, _maxNumGroupsVotedFor *big.Int) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "setMaxNumGroupsVotedFor", _maxNumGroupsVotedFor)
}

// SetMaxNumGroupsVotedFor is a paid mutator transaction binding the contract method 0x3c55a73c.
//
// Solidity: function setMaxNumGroupsVotedFor(uint256 _maxNumGroupsVotedFor) returns(bool)
func (_Election *ElectionSession) SetMaxNumGroupsVotedFor(_maxNumGroupsVotedFor *big.Int) (*types.Transaction, error) {
	return _Election.Contract.SetMaxNumGroupsVotedFor(&_Election.TransactOpts, _maxNumGroupsVotedFor)
}

// SetMaxNumGroupsVotedFor is a paid mutator transaction binding the contract method 0x3c55a73c.
//
// Solidity: function setMaxNumGroupsVotedFor(uint256 _maxNumGroupsVotedFor) returns(bool)
func (_Election *ElectionTransactorSession) SetMaxNumGroupsVotedFor(_maxNumGroupsVotedFor *big.Int) (*types.Transaction, error) {
	return _Election.Contract.SetMaxNumGroupsVotedFor(&_Election.TransactOpts, _maxNumGroupsVotedFor)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Election *ElectionTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Election *ElectionSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Election.Contract.SetRegistry(&_Election.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Election *ElectionTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Election.Contract.SetRegistry(&_Election.TransactOpts, registryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Election *ElectionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Election *ElectionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Election.Contract.TransferOwnership(&_Election.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Election *ElectionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Election.Contract.TransferOwnership(&_Election.TransactOpts, newOwner)
}

// UpdateTotalVotesByAccountForGroup is a paid mutator transaction binding the contract method 0x35a8376b.
//
// Solidity: function updateTotalVotesByAccountForGroup(address account, address group) returns()
func (_Election *ElectionTransactor) UpdateTotalVotesByAccountForGroup(opts *bind.TransactOpts, account common.Address, group common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "updateTotalVotesByAccountForGroup", account, group)
}

// UpdateTotalVotesByAccountForGroup is a paid mutator transaction binding the contract method 0x35a8376b.
//
// Solidity: function updateTotalVotesByAccountForGroup(address account, address group) returns()
func (_Election *ElectionSession) UpdateTotalVotesByAccountForGroup(account common.Address, group common.Address) (*types.Transaction, error) {
	return _Election.Contract.UpdateTotalVotesByAccountForGroup(&_Election.TransactOpts, account, group)
}

// UpdateTotalVotesByAccountForGroup is a paid mutator transaction binding the contract method 0x35a8376b.
//
// Solidity: function updateTotalVotesByAccountForGroup(address account, address group) returns()
func (_Election *ElectionTransactorSession) UpdateTotalVotesByAccountForGroup(account common.Address, group common.Address) (*types.Transaction, error) {
	return _Election.Contract.UpdateTotalVotesByAccountForGroup(&_Election.TransactOpts, account, group)
}

// Vote is a paid mutator transaction binding the contract method 0x580d747a.
//
// Solidity: function vote(address group, uint256 value, address lesser, address greater) returns(bool)
func (_Election *ElectionTransactor) Vote(opts *bind.TransactOpts, group common.Address, value *big.Int, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.contract.Transact(opts, "vote", group, value, lesser, greater)
}

// Vote is a paid mutator transaction binding the contract method 0x580d747a.
//
// Solidity: function vote(address group, uint256 value, address lesser, address greater) returns(bool)
func (_Election *ElectionSession) Vote(group common.Address, value *big.Int, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.Contract.Vote(&_Election.TransactOpts, group, value, lesser, greater)
}

// Vote is a paid mutator transaction binding the contract method 0x580d747a.
//
// Solidity: function vote(address group, uint256 value, address lesser, address greater) returns(bool)
func (_Election *ElectionTransactorSession) Vote(group common.Address, value *big.Int, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Election.Contract.Vote(&_Election.TransactOpts, group, value, lesser, greater)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Election *ElectionFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Election.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "AllowedToVoteOverMaxNumberOfGroups":
		event, err = _Election.ParseAllowedToVoteOverMaxNumberOfGroups(log)
	case "ElectabilityThresholdSet":
		event, err = _Election.ParseElectabilityThresholdSet(log)
	case "ElectableValidatorsSet":
		event, err = _Election.ParseElectableValidatorsSet(log)
	case "EpochRewardsDistributedToVoters":
		event, err = _Election.ParseEpochRewardsDistributedToVoters(log)
	case "MaxNumGroupsVotedForSet":
		event, err = _Election.ParseMaxNumGroupsVotedForSet(log)
	case "OwnershipTransferred":
		event, err = _Election.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Election.ParseRegistrySet(log)
	case "ValidatorGroupActiveVoteRevoked":
		event, err = _Election.ParseValidatorGroupActiveVoteRevoked(log)
	case "ValidatorGroupMarkedEligible":
		event, err = _Election.ParseValidatorGroupMarkedEligible(log)
	case "ValidatorGroupMarkedIneligible":
		event, err = _Election.ParseValidatorGroupMarkedIneligible(log)
	case "ValidatorGroupPendingVoteRevoked":
		event, err = _Election.ParseValidatorGroupPendingVoteRevoked(log)
	case "ValidatorGroupVoteActivated":
		event, err = _Election.ParseValidatorGroupVoteActivated(log)
	case "ValidatorGroupVoteCast":
		event, err = _Election.ParseValidatorGroupVoteCast(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// ElectionAllowedToVoteOverMaxNumberOfGroupsIterator is returned from FilterAllowedToVoteOverMaxNumberOfGroups and is used to iterate over the raw logs and unpacked data for AllowedToVoteOverMaxNumberOfGroups events raised by the Election contract.
type ElectionAllowedToVoteOverMaxNumberOfGroupsIterator struct {
	Event *ElectionAllowedToVoteOverMaxNumberOfGroups // Event containing the contract specifics and raw log

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
func (it *ElectionAllowedToVoteOverMaxNumberOfGroupsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionAllowedToVoteOverMaxNumberOfGroups)
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
		it.Event = new(ElectionAllowedToVoteOverMaxNumberOfGroups)
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
func (it *ElectionAllowedToVoteOverMaxNumberOfGroupsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionAllowedToVoteOverMaxNumberOfGroupsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionAllowedToVoteOverMaxNumberOfGroups represents a AllowedToVoteOverMaxNumberOfGroups event raised by the Election contract.
type ElectionAllowedToVoteOverMaxNumberOfGroups struct {
	Account common.Address
	Flag    bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAllowedToVoteOverMaxNumberOfGroups is a free log retrieval operation binding the contract event 0xd9ff9bf9c0aa22c57c14972fe77841756843243a74e331dabcb04a4d8dbf11ff.
//
// Solidity: event AllowedToVoteOverMaxNumberOfGroups(address indexed account, bool flag)
func (_Election *ElectionFilterer) FilterAllowedToVoteOverMaxNumberOfGroups(opts *bind.FilterOpts, account []common.Address) (*ElectionAllowedToVoteOverMaxNumberOfGroupsIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "AllowedToVoteOverMaxNumberOfGroups", accountRule)
	if err != nil {
		return nil, err
	}
	return &ElectionAllowedToVoteOverMaxNumberOfGroupsIterator{contract: _Election.contract, event: "AllowedToVoteOverMaxNumberOfGroups", logs: logs, sub: sub}, nil
}

// WatchAllowedToVoteOverMaxNumberOfGroups is a free log subscription operation binding the contract event 0xd9ff9bf9c0aa22c57c14972fe77841756843243a74e331dabcb04a4d8dbf11ff.
//
// Solidity: event AllowedToVoteOverMaxNumberOfGroups(address indexed account, bool flag)
func (_Election *ElectionFilterer) WatchAllowedToVoteOverMaxNumberOfGroups(opts *bind.WatchOpts, sink chan<- *ElectionAllowedToVoteOverMaxNumberOfGroups, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "AllowedToVoteOverMaxNumberOfGroups", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionAllowedToVoteOverMaxNumberOfGroups)
				if err := _Election.contract.UnpackLog(event, "AllowedToVoteOverMaxNumberOfGroups", log); err != nil {
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

// ParseAllowedToVoteOverMaxNumberOfGroups is a log parse operation binding the contract event 0xd9ff9bf9c0aa22c57c14972fe77841756843243a74e331dabcb04a4d8dbf11ff.
//
// Solidity: event AllowedToVoteOverMaxNumberOfGroups(address indexed account, bool flag)
func (_Election *ElectionFilterer) ParseAllowedToVoteOverMaxNumberOfGroups(log types.Log) (*ElectionAllowedToVoteOverMaxNumberOfGroups, error) {
	event := new(ElectionAllowedToVoteOverMaxNumberOfGroups)
	if err := _Election.contract.UnpackLog(event, "AllowedToVoteOverMaxNumberOfGroups", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionElectabilityThresholdSetIterator is returned from FilterElectabilityThresholdSet and is used to iterate over the raw logs and unpacked data for ElectabilityThresholdSet events raised by the Election contract.
type ElectionElectabilityThresholdSetIterator struct {
	Event *ElectionElectabilityThresholdSet // Event containing the contract specifics and raw log

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
func (it *ElectionElectabilityThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionElectabilityThresholdSet)
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
		it.Event = new(ElectionElectabilityThresholdSet)
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
func (it *ElectionElectabilityThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionElectabilityThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionElectabilityThresholdSet represents a ElectabilityThresholdSet event raised by the Election contract.
type ElectionElectabilityThresholdSet struct {
	ElectabilityThreshold *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterElectabilityThresholdSet is a free log retrieval operation binding the contract event 0x9854be03126e38f9c318d8aabe1b150d09cb3a57059b21855b1e11d44e082c1a.
//
// Solidity: event ElectabilityThresholdSet(uint256 electabilityThreshold)
func (_Election *ElectionFilterer) FilterElectabilityThresholdSet(opts *bind.FilterOpts) (*ElectionElectabilityThresholdSetIterator, error) {

	logs, sub, err := _Election.contract.FilterLogs(opts, "ElectabilityThresholdSet")
	if err != nil {
		return nil, err
	}
	return &ElectionElectabilityThresholdSetIterator{contract: _Election.contract, event: "ElectabilityThresholdSet", logs: logs, sub: sub}, nil
}

// WatchElectabilityThresholdSet is a free log subscription operation binding the contract event 0x9854be03126e38f9c318d8aabe1b150d09cb3a57059b21855b1e11d44e082c1a.
//
// Solidity: event ElectabilityThresholdSet(uint256 electabilityThreshold)
func (_Election *ElectionFilterer) WatchElectabilityThresholdSet(opts *bind.WatchOpts, sink chan<- *ElectionElectabilityThresholdSet) (event.Subscription, error) {

	logs, sub, err := _Election.contract.WatchLogs(opts, "ElectabilityThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionElectabilityThresholdSet)
				if err := _Election.contract.UnpackLog(event, "ElectabilityThresholdSet", log); err != nil {
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

// ParseElectabilityThresholdSet is a log parse operation binding the contract event 0x9854be03126e38f9c318d8aabe1b150d09cb3a57059b21855b1e11d44e082c1a.
//
// Solidity: event ElectabilityThresholdSet(uint256 electabilityThreshold)
func (_Election *ElectionFilterer) ParseElectabilityThresholdSet(log types.Log) (*ElectionElectabilityThresholdSet, error) {
	event := new(ElectionElectabilityThresholdSet)
	if err := _Election.contract.UnpackLog(event, "ElectabilityThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionElectableValidatorsSetIterator is returned from FilterElectableValidatorsSet and is used to iterate over the raw logs and unpacked data for ElectableValidatorsSet events raised by the Election contract.
type ElectionElectableValidatorsSetIterator struct {
	Event *ElectionElectableValidatorsSet // Event containing the contract specifics and raw log

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
func (it *ElectionElectableValidatorsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionElectableValidatorsSet)
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
		it.Event = new(ElectionElectableValidatorsSet)
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
func (it *ElectionElectableValidatorsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionElectableValidatorsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionElectableValidatorsSet represents a ElectableValidatorsSet event raised by the Election contract.
type ElectionElectableValidatorsSet struct {
	Min *big.Int
	Max *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterElectableValidatorsSet is a free log retrieval operation binding the contract event 0xb3ae64819ff89f6136eb58b8563cb32c6550f17eaf97f9ecc32f23783229f6de.
//
// Solidity: event ElectableValidatorsSet(uint256 min, uint256 max)
func (_Election *ElectionFilterer) FilterElectableValidatorsSet(opts *bind.FilterOpts) (*ElectionElectableValidatorsSetIterator, error) {

	logs, sub, err := _Election.contract.FilterLogs(opts, "ElectableValidatorsSet")
	if err != nil {
		return nil, err
	}
	return &ElectionElectableValidatorsSetIterator{contract: _Election.contract, event: "ElectableValidatorsSet", logs: logs, sub: sub}, nil
}

// WatchElectableValidatorsSet is a free log subscription operation binding the contract event 0xb3ae64819ff89f6136eb58b8563cb32c6550f17eaf97f9ecc32f23783229f6de.
//
// Solidity: event ElectableValidatorsSet(uint256 min, uint256 max)
func (_Election *ElectionFilterer) WatchElectableValidatorsSet(opts *bind.WatchOpts, sink chan<- *ElectionElectableValidatorsSet) (event.Subscription, error) {

	logs, sub, err := _Election.contract.WatchLogs(opts, "ElectableValidatorsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionElectableValidatorsSet)
				if err := _Election.contract.UnpackLog(event, "ElectableValidatorsSet", log); err != nil {
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

// ParseElectableValidatorsSet is a log parse operation binding the contract event 0xb3ae64819ff89f6136eb58b8563cb32c6550f17eaf97f9ecc32f23783229f6de.
//
// Solidity: event ElectableValidatorsSet(uint256 min, uint256 max)
func (_Election *ElectionFilterer) ParseElectableValidatorsSet(log types.Log) (*ElectionElectableValidatorsSet, error) {
	event := new(ElectionElectableValidatorsSet)
	if err := _Election.contract.UnpackLog(event, "ElectableValidatorsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionEpochRewardsDistributedToVotersIterator is returned from FilterEpochRewardsDistributedToVoters and is used to iterate over the raw logs and unpacked data for EpochRewardsDistributedToVoters events raised by the Election contract.
type ElectionEpochRewardsDistributedToVotersIterator struct {
	Event *ElectionEpochRewardsDistributedToVoters // Event containing the contract specifics and raw log

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
func (it *ElectionEpochRewardsDistributedToVotersIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionEpochRewardsDistributedToVoters)
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
		it.Event = new(ElectionEpochRewardsDistributedToVoters)
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
func (it *ElectionEpochRewardsDistributedToVotersIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionEpochRewardsDistributedToVotersIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionEpochRewardsDistributedToVoters represents a EpochRewardsDistributedToVoters event raised by the Election contract.
type ElectionEpochRewardsDistributedToVoters struct {
	Group common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochRewardsDistributedToVoters is a free log retrieval operation binding the contract event 0x91ba34d62474c14d6c623cd322f4256666c7a45b7fdaa3378e009d39dfcec2a7.
//
// Solidity: event EpochRewardsDistributedToVoters(address indexed group, uint256 value)
func (_Election *ElectionFilterer) FilterEpochRewardsDistributedToVoters(opts *bind.FilterOpts, group []common.Address) (*ElectionEpochRewardsDistributedToVotersIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "EpochRewardsDistributedToVoters", groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionEpochRewardsDistributedToVotersIterator{contract: _Election.contract, event: "EpochRewardsDistributedToVoters", logs: logs, sub: sub}, nil
}

// WatchEpochRewardsDistributedToVoters is a free log subscription operation binding the contract event 0x91ba34d62474c14d6c623cd322f4256666c7a45b7fdaa3378e009d39dfcec2a7.
//
// Solidity: event EpochRewardsDistributedToVoters(address indexed group, uint256 value)
func (_Election *ElectionFilterer) WatchEpochRewardsDistributedToVoters(opts *bind.WatchOpts, sink chan<- *ElectionEpochRewardsDistributedToVoters, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "EpochRewardsDistributedToVoters", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionEpochRewardsDistributedToVoters)
				if err := _Election.contract.UnpackLog(event, "EpochRewardsDistributedToVoters", log); err != nil {
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

// ParseEpochRewardsDistributedToVoters is a log parse operation binding the contract event 0x91ba34d62474c14d6c623cd322f4256666c7a45b7fdaa3378e009d39dfcec2a7.
//
// Solidity: event EpochRewardsDistributedToVoters(address indexed group, uint256 value)
func (_Election *ElectionFilterer) ParseEpochRewardsDistributedToVoters(log types.Log) (*ElectionEpochRewardsDistributedToVoters, error) {
	event := new(ElectionEpochRewardsDistributedToVoters)
	if err := _Election.contract.UnpackLog(event, "EpochRewardsDistributedToVoters", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionMaxNumGroupsVotedForSetIterator is returned from FilterMaxNumGroupsVotedForSet and is used to iterate over the raw logs and unpacked data for MaxNumGroupsVotedForSet events raised by the Election contract.
type ElectionMaxNumGroupsVotedForSetIterator struct {
	Event *ElectionMaxNumGroupsVotedForSet // Event containing the contract specifics and raw log

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
func (it *ElectionMaxNumGroupsVotedForSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionMaxNumGroupsVotedForSet)
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
		it.Event = new(ElectionMaxNumGroupsVotedForSet)
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
func (it *ElectionMaxNumGroupsVotedForSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionMaxNumGroupsVotedForSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionMaxNumGroupsVotedForSet represents a MaxNumGroupsVotedForSet event raised by the Election contract.
type ElectionMaxNumGroupsVotedForSet struct {
	MaxNumGroupsVotedFor *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterMaxNumGroupsVotedForSet is a free log retrieval operation binding the contract event 0x1993a3864c31265ef86eec51d147eff697dee0466c92ac9abddcc4c4c6829348.
//
// Solidity: event MaxNumGroupsVotedForSet(uint256 maxNumGroupsVotedFor)
func (_Election *ElectionFilterer) FilterMaxNumGroupsVotedForSet(opts *bind.FilterOpts) (*ElectionMaxNumGroupsVotedForSetIterator, error) {

	logs, sub, err := _Election.contract.FilterLogs(opts, "MaxNumGroupsVotedForSet")
	if err != nil {
		return nil, err
	}
	return &ElectionMaxNumGroupsVotedForSetIterator{contract: _Election.contract, event: "MaxNumGroupsVotedForSet", logs: logs, sub: sub}, nil
}

// WatchMaxNumGroupsVotedForSet is a free log subscription operation binding the contract event 0x1993a3864c31265ef86eec51d147eff697dee0466c92ac9abddcc4c4c6829348.
//
// Solidity: event MaxNumGroupsVotedForSet(uint256 maxNumGroupsVotedFor)
func (_Election *ElectionFilterer) WatchMaxNumGroupsVotedForSet(opts *bind.WatchOpts, sink chan<- *ElectionMaxNumGroupsVotedForSet) (event.Subscription, error) {

	logs, sub, err := _Election.contract.WatchLogs(opts, "MaxNumGroupsVotedForSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionMaxNumGroupsVotedForSet)
				if err := _Election.contract.UnpackLog(event, "MaxNumGroupsVotedForSet", log); err != nil {
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

// ParseMaxNumGroupsVotedForSet is a log parse operation binding the contract event 0x1993a3864c31265ef86eec51d147eff697dee0466c92ac9abddcc4c4c6829348.
//
// Solidity: event MaxNumGroupsVotedForSet(uint256 maxNumGroupsVotedFor)
func (_Election *ElectionFilterer) ParseMaxNumGroupsVotedForSet(log types.Log) (*ElectionMaxNumGroupsVotedForSet, error) {
	event := new(ElectionMaxNumGroupsVotedForSet)
	if err := _Election.contract.UnpackLog(event, "MaxNumGroupsVotedForSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Election contract.
type ElectionOwnershipTransferredIterator struct {
	Event *ElectionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ElectionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionOwnershipTransferred)
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
		it.Event = new(ElectionOwnershipTransferred)
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
func (it *ElectionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionOwnershipTransferred represents a OwnershipTransferred event raised by the Election contract.
type ElectionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Election *ElectionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ElectionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ElectionOwnershipTransferredIterator{contract: _Election.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Election *ElectionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ElectionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionOwnershipTransferred)
				if err := _Election.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Election *ElectionFilterer) ParseOwnershipTransferred(log types.Log) (*ElectionOwnershipTransferred, error) {
	event := new(ElectionOwnershipTransferred)
	if err := _Election.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Election contract.
type ElectionRegistrySetIterator struct {
	Event *ElectionRegistrySet // Event containing the contract specifics and raw log

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
func (it *ElectionRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionRegistrySet)
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
		it.Event = new(ElectionRegistrySet)
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
func (it *ElectionRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionRegistrySet represents a RegistrySet event raised by the Election contract.
type ElectionRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Election *ElectionFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*ElectionRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &ElectionRegistrySetIterator{contract: _Election.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Election *ElectionFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *ElectionRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionRegistrySet)
				if err := _Election.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Election *ElectionFilterer) ParseRegistrySet(log types.Log) (*ElectionRegistrySet, error) {
	event := new(ElectionRegistrySet)
	if err := _Election.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionValidatorGroupActiveVoteRevokedIterator is returned from FilterValidatorGroupActiveVoteRevoked and is used to iterate over the raw logs and unpacked data for ValidatorGroupActiveVoteRevoked events raised by the Election contract.
type ElectionValidatorGroupActiveVoteRevokedIterator struct {
	Event *ElectionValidatorGroupActiveVoteRevoked // Event containing the contract specifics and raw log

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
func (it *ElectionValidatorGroupActiveVoteRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionValidatorGroupActiveVoteRevoked)
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
		it.Event = new(ElectionValidatorGroupActiveVoteRevoked)
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
func (it *ElectionValidatorGroupActiveVoteRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionValidatorGroupActiveVoteRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionValidatorGroupActiveVoteRevoked represents a ValidatorGroupActiveVoteRevoked event raised by the Election contract.
type ElectionValidatorGroupActiveVoteRevoked struct {
	Account common.Address
	Group   common.Address
	Value   *big.Int
	Units   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupActiveVoteRevoked is a free log retrieval operation binding the contract event 0xae7458f8697a680da6be36406ea0b8f40164915ac9cc40c0dad05a2ff6e8c6a8.
//
// Solidity: event ValidatorGroupActiveVoteRevoked(address indexed account, address indexed group, uint256 value, uint256 units)
func (_Election *ElectionFilterer) FilterValidatorGroupActiveVoteRevoked(opts *bind.FilterOpts, account []common.Address, group []common.Address) (*ElectionValidatorGroupActiveVoteRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "ValidatorGroupActiveVoteRevoked", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionValidatorGroupActiveVoteRevokedIterator{contract: _Election.contract, event: "ValidatorGroupActiveVoteRevoked", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupActiveVoteRevoked is a free log subscription operation binding the contract event 0xae7458f8697a680da6be36406ea0b8f40164915ac9cc40c0dad05a2ff6e8c6a8.
//
// Solidity: event ValidatorGroupActiveVoteRevoked(address indexed account, address indexed group, uint256 value, uint256 units)
func (_Election *ElectionFilterer) WatchValidatorGroupActiveVoteRevoked(opts *bind.WatchOpts, sink chan<- *ElectionValidatorGroupActiveVoteRevoked, account []common.Address, group []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "ValidatorGroupActiveVoteRevoked", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionValidatorGroupActiveVoteRevoked)
				if err := _Election.contract.UnpackLog(event, "ValidatorGroupActiveVoteRevoked", log); err != nil {
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

// ParseValidatorGroupActiveVoteRevoked is a log parse operation binding the contract event 0xae7458f8697a680da6be36406ea0b8f40164915ac9cc40c0dad05a2ff6e8c6a8.
//
// Solidity: event ValidatorGroupActiveVoteRevoked(address indexed account, address indexed group, uint256 value, uint256 units)
func (_Election *ElectionFilterer) ParseValidatorGroupActiveVoteRevoked(log types.Log) (*ElectionValidatorGroupActiveVoteRevoked, error) {
	event := new(ElectionValidatorGroupActiveVoteRevoked)
	if err := _Election.contract.UnpackLog(event, "ValidatorGroupActiveVoteRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionValidatorGroupMarkedEligibleIterator is returned from FilterValidatorGroupMarkedEligible and is used to iterate over the raw logs and unpacked data for ValidatorGroupMarkedEligible events raised by the Election contract.
type ElectionValidatorGroupMarkedEligibleIterator struct {
	Event *ElectionValidatorGroupMarkedEligible // Event containing the contract specifics and raw log

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
func (it *ElectionValidatorGroupMarkedEligibleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionValidatorGroupMarkedEligible)
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
		it.Event = new(ElectionValidatorGroupMarkedEligible)
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
func (it *ElectionValidatorGroupMarkedEligibleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionValidatorGroupMarkedEligibleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionValidatorGroupMarkedEligible represents a ValidatorGroupMarkedEligible event raised by the Election contract.
type ElectionValidatorGroupMarkedEligible struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMarkedEligible is a free log retrieval operation binding the contract event 0x8f21dc7ff6f55d73e4fca52a4ef4fcc14fbda43ac338d24922519d51455d39c1.
//
// Solidity: event ValidatorGroupMarkedEligible(address indexed group)
func (_Election *ElectionFilterer) FilterValidatorGroupMarkedEligible(opts *bind.FilterOpts, group []common.Address) (*ElectionValidatorGroupMarkedEligibleIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "ValidatorGroupMarkedEligible", groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionValidatorGroupMarkedEligibleIterator{contract: _Election.contract, event: "ValidatorGroupMarkedEligible", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMarkedEligible is a free log subscription operation binding the contract event 0x8f21dc7ff6f55d73e4fca52a4ef4fcc14fbda43ac338d24922519d51455d39c1.
//
// Solidity: event ValidatorGroupMarkedEligible(address indexed group)
func (_Election *ElectionFilterer) WatchValidatorGroupMarkedEligible(opts *bind.WatchOpts, sink chan<- *ElectionValidatorGroupMarkedEligible, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "ValidatorGroupMarkedEligible", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionValidatorGroupMarkedEligible)
				if err := _Election.contract.UnpackLog(event, "ValidatorGroupMarkedEligible", log); err != nil {
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

// ParseValidatorGroupMarkedEligible is a log parse operation binding the contract event 0x8f21dc7ff6f55d73e4fca52a4ef4fcc14fbda43ac338d24922519d51455d39c1.
//
// Solidity: event ValidatorGroupMarkedEligible(address indexed group)
func (_Election *ElectionFilterer) ParseValidatorGroupMarkedEligible(log types.Log) (*ElectionValidatorGroupMarkedEligible, error) {
	event := new(ElectionValidatorGroupMarkedEligible)
	if err := _Election.contract.UnpackLog(event, "ValidatorGroupMarkedEligible", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionValidatorGroupMarkedIneligibleIterator is returned from FilterValidatorGroupMarkedIneligible and is used to iterate over the raw logs and unpacked data for ValidatorGroupMarkedIneligible events raised by the Election contract.
type ElectionValidatorGroupMarkedIneligibleIterator struct {
	Event *ElectionValidatorGroupMarkedIneligible // Event containing the contract specifics and raw log

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
func (it *ElectionValidatorGroupMarkedIneligibleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionValidatorGroupMarkedIneligible)
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
		it.Event = new(ElectionValidatorGroupMarkedIneligible)
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
func (it *ElectionValidatorGroupMarkedIneligibleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionValidatorGroupMarkedIneligibleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionValidatorGroupMarkedIneligible represents a ValidatorGroupMarkedIneligible event raised by the Election contract.
type ElectionValidatorGroupMarkedIneligible struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMarkedIneligible is a free log retrieval operation binding the contract event 0x5c8cd4e832f3a7d79f9208c2acf25a412143aa3f751cfd3728c42a0fea4921a8.
//
// Solidity: event ValidatorGroupMarkedIneligible(address indexed group)
func (_Election *ElectionFilterer) FilterValidatorGroupMarkedIneligible(opts *bind.FilterOpts, group []common.Address) (*ElectionValidatorGroupMarkedIneligibleIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "ValidatorGroupMarkedIneligible", groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionValidatorGroupMarkedIneligibleIterator{contract: _Election.contract, event: "ValidatorGroupMarkedIneligible", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMarkedIneligible is a free log subscription operation binding the contract event 0x5c8cd4e832f3a7d79f9208c2acf25a412143aa3f751cfd3728c42a0fea4921a8.
//
// Solidity: event ValidatorGroupMarkedIneligible(address indexed group)
func (_Election *ElectionFilterer) WatchValidatorGroupMarkedIneligible(opts *bind.WatchOpts, sink chan<- *ElectionValidatorGroupMarkedIneligible, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "ValidatorGroupMarkedIneligible", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionValidatorGroupMarkedIneligible)
				if err := _Election.contract.UnpackLog(event, "ValidatorGroupMarkedIneligible", log); err != nil {
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

// ParseValidatorGroupMarkedIneligible is a log parse operation binding the contract event 0x5c8cd4e832f3a7d79f9208c2acf25a412143aa3f751cfd3728c42a0fea4921a8.
//
// Solidity: event ValidatorGroupMarkedIneligible(address indexed group)
func (_Election *ElectionFilterer) ParseValidatorGroupMarkedIneligible(log types.Log) (*ElectionValidatorGroupMarkedIneligible, error) {
	event := new(ElectionValidatorGroupMarkedIneligible)
	if err := _Election.contract.UnpackLog(event, "ValidatorGroupMarkedIneligible", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionValidatorGroupPendingVoteRevokedIterator is returned from FilterValidatorGroupPendingVoteRevoked and is used to iterate over the raw logs and unpacked data for ValidatorGroupPendingVoteRevoked events raised by the Election contract.
type ElectionValidatorGroupPendingVoteRevokedIterator struct {
	Event *ElectionValidatorGroupPendingVoteRevoked // Event containing the contract specifics and raw log

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
func (it *ElectionValidatorGroupPendingVoteRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionValidatorGroupPendingVoteRevoked)
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
		it.Event = new(ElectionValidatorGroupPendingVoteRevoked)
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
func (it *ElectionValidatorGroupPendingVoteRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionValidatorGroupPendingVoteRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionValidatorGroupPendingVoteRevoked represents a ValidatorGroupPendingVoteRevoked event raised by the Election contract.
type ElectionValidatorGroupPendingVoteRevoked struct {
	Account common.Address
	Group   common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupPendingVoteRevoked is a free log retrieval operation binding the contract event 0x148075455e24d5cf538793db3e917a157cbadac69dd6a304186daf11b23f76fe.
//
// Solidity: event ValidatorGroupPendingVoteRevoked(address indexed account, address indexed group, uint256 value)
func (_Election *ElectionFilterer) FilterValidatorGroupPendingVoteRevoked(opts *bind.FilterOpts, account []common.Address, group []common.Address) (*ElectionValidatorGroupPendingVoteRevokedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "ValidatorGroupPendingVoteRevoked", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionValidatorGroupPendingVoteRevokedIterator{contract: _Election.contract, event: "ValidatorGroupPendingVoteRevoked", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupPendingVoteRevoked is a free log subscription operation binding the contract event 0x148075455e24d5cf538793db3e917a157cbadac69dd6a304186daf11b23f76fe.
//
// Solidity: event ValidatorGroupPendingVoteRevoked(address indexed account, address indexed group, uint256 value)
func (_Election *ElectionFilterer) WatchValidatorGroupPendingVoteRevoked(opts *bind.WatchOpts, sink chan<- *ElectionValidatorGroupPendingVoteRevoked, account []common.Address, group []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "ValidatorGroupPendingVoteRevoked", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionValidatorGroupPendingVoteRevoked)
				if err := _Election.contract.UnpackLog(event, "ValidatorGroupPendingVoteRevoked", log); err != nil {
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

// ParseValidatorGroupPendingVoteRevoked is a log parse operation binding the contract event 0x148075455e24d5cf538793db3e917a157cbadac69dd6a304186daf11b23f76fe.
//
// Solidity: event ValidatorGroupPendingVoteRevoked(address indexed account, address indexed group, uint256 value)
func (_Election *ElectionFilterer) ParseValidatorGroupPendingVoteRevoked(log types.Log) (*ElectionValidatorGroupPendingVoteRevoked, error) {
	event := new(ElectionValidatorGroupPendingVoteRevoked)
	if err := _Election.contract.UnpackLog(event, "ValidatorGroupPendingVoteRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionValidatorGroupVoteActivatedIterator is returned from FilterValidatorGroupVoteActivated and is used to iterate over the raw logs and unpacked data for ValidatorGroupVoteActivated events raised by the Election contract.
type ElectionValidatorGroupVoteActivatedIterator struct {
	Event *ElectionValidatorGroupVoteActivated // Event containing the contract specifics and raw log

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
func (it *ElectionValidatorGroupVoteActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionValidatorGroupVoteActivated)
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
		it.Event = new(ElectionValidatorGroupVoteActivated)
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
func (it *ElectionValidatorGroupVoteActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionValidatorGroupVoteActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionValidatorGroupVoteActivated represents a ValidatorGroupVoteActivated event raised by the Election contract.
type ElectionValidatorGroupVoteActivated struct {
	Account common.Address
	Group   common.Address
	Value   *big.Int
	Units   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupVoteActivated is a free log retrieval operation binding the contract event 0x45aac85f38083b18efe2d441a65b9c1ae177c78307cb5a5d4aec8f7dbcaeabfe.
//
// Solidity: event ValidatorGroupVoteActivated(address indexed account, address indexed group, uint256 value, uint256 units)
func (_Election *ElectionFilterer) FilterValidatorGroupVoteActivated(opts *bind.FilterOpts, account []common.Address, group []common.Address) (*ElectionValidatorGroupVoteActivatedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "ValidatorGroupVoteActivated", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionValidatorGroupVoteActivatedIterator{contract: _Election.contract, event: "ValidatorGroupVoteActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupVoteActivated is a free log subscription operation binding the contract event 0x45aac85f38083b18efe2d441a65b9c1ae177c78307cb5a5d4aec8f7dbcaeabfe.
//
// Solidity: event ValidatorGroupVoteActivated(address indexed account, address indexed group, uint256 value, uint256 units)
func (_Election *ElectionFilterer) WatchValidatorGroupVoteActivated(opts *bind.WatchOpts, sink chan<- *ElectionValidatorGroupVoteActivated, account []common.Address, group []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "ValidatorGroupVoteActivated", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionValidatorGroupVoteActivated)
				if err := _Election.contract.UnpackLog(event, "ValidatorGroupVoteActivated", log); err != nil {
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

// ParseValidatorGroupVoteActivated is a log parse operation binding the contract event 0x45aac85f38083b18efe2d441a65b9c1ae177c78307cb5a5d4aec8f7dbcaeabfe.
//
// Solidity: event ValidatorGroupVoteActivated(address indexed account, address indexed group, uint256 value, uint256 units)
func (_Election *ElectionFilterer) ParseValidatorGroupVoteActivated(log types.Log) (*ElectionValidatorGroupVoteActivated, error) {
	event := new(ElectionValidatorGroupVoteActivated)
	if err := _Election.contract.UnpackLog(event, "ValidatorGroupVoteActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ElectionValidatorGroupVoteCastIterator is returned from FilterValidatorGroupVoteCast and is used to iterate over the raw logs and unpacked data for ValidatorGroupVoteCast events raised by the Election contract.
type ElectionValidatorGroupVoteCastIterator struct {
	Event *ElectionValidatorGroupVoteCast // Event containing the contract specifics and raw log

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
func (it *ElectionValidatorGroupVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ElectionValidatorGroupVoteCast)
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
		it.Event = new(ElectionValidatorGroupVoteCast)
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
func (it *ElectionValidatorGroupVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ElectionValidatorGroupVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ElectionValidatorGroupVoteCast represents a ValidatorGroupVoteCast event raised by the Election contract.
type ElectionValidatorGroupVoteCast struct {
	Account common.Address
	Group   common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupVoteCast is a free log retrieval operation binding the contract event 0xd3532f70444893db82221041edb4dc26c94593aeb364b0b14dfc77d5ee905152.
//
// Solidity: event ValidatorGroupVoteCast(address indexed account, address indexed group, uint256 value)
func (_Election *ElectionFilterer) FilterValidatorGroupVoteCast(opts *bind.FilterOpts, account []common.Address, group []common.Address) (*ElectionValidatorGroupVoteCastIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.FilterLogs(opts, "ValidatorGroupVoteCast", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ElectionValidatorGroupVoteCastIterator{contract: _Election.contract, event: "ValidatorGroupVoteCast", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupVoteCast is a free log subscription operation binding the contract event 0xd3532f70444893db82221041edb4dc26c94593aeb364b0b14dfc77d5ee905152.
//
// Solidity: event ValidatorGroupVoteCast(address indexed account, address indexed group, uint256 value)
func (_Election *ElectionFilterer) WatchValidatorGroupVoteCast(opts *bind.WatchOpts, sink chan<- *ElectionValidatorGroupVoteCast, account []common.Address, group []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Election.contract.WatchLogs(opts, "ValidatorGroupVoteCast", accountRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ElectionValidatorGroupVoteCast)
				if err := _Election.contract.UnpackLog(event, "ValidatorGroupVoteCast", log); err != nil {
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

// ParseValidatorGroupVoteCast is a log parse operation binding the contract event 0xd3532f70444893db82221041edb4dc26c94593aeb364b0b14dfc77d5ee905152.
//
// Solidity: event ValidatorGroupVoteCast(address indexed account, address indexed group, uint256 value)
func (_Election *ElectionFilterer) ParseValidatorGroupVoteCast(log types.Log) (*ElectionValidatorGroupVoteCast, error) {
	event := new(ElectionValidatorGroupVoteCast)
	if err := _Election.contract.UnpackLog(event, "ValidatorGroupVoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
