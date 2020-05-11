// Copyright 2020 Celo Org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// ValidatorsABI is the input ABI used to generate the binding from.
const ValidatorsABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"sender\",\"type\":\"address\"},{\"name\":\"blsKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slashingMultiplierResetPeriod\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"membershipHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxGroupSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"},{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"validatorLockedGoldRequirements\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"groupLockedGoldRequirements\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commissionUpdateDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"MaxGroupSizeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"CommissionUpdateDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"exponent\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"ValidatorScoreParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"GroupLockedGoldRequirementsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"ValidatorLockedGoldRequirementsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"MembershipHistoryLengthSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorAffiliated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorDeaffiliated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorEcdsaPublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorBlsPublicKeyUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"score\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"epochScore\",\"type\":\"uint256\"}],\"name\":\"ValidatorScoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"}],\"name\":\"ValidatorGroupDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMemberAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMemberRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorGroupMemberReordered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"activationBlock\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupCommissionUpdateQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"ValidatorGroupCommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"validatorPayment\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"group\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"groupPayment\",\"type\":\"uint256\"}],\"name\":\"ValidatorEpochPaymentDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"registryAddress\",\"type\":\"address\"},{\"name\":\"groupRequirementValue\",\"type\":\"uint256\"},{\"name\":\"groupRequirementDuration\",\"type\":\"uint256\"},{\"name\":\"validatorRequirementValue\",\"type\":\"uint256\"},{\"name\":\"validatorRequirementDuration\",\"type\":\"uint256\"},{\"name\":\"validatorScoreExponent\",\"type\":\"uint256\"},{\"name\":\"validatorScoreAdjustmentSpeed\",\"type\":\"uint256\"},{\"name\":\"_membershipHistoryLength\",\"type\":\"uint256\"},{\"name\":\"_slashingMultiplierResetPeriod\",\"type\":\"uint256\"},{\"name\":\"_maxGroupSize\",\"type\":\"uint256\"},{\"name\":\"_commissionUpdateDelay\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"delay\",\"type\":\"uint256\"}],\"name\":\"setCommissionUpdateDelay\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"setMaxGroupSize\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"setMembershipHistoryLength\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"exponent\",\"type\":\"uint256\"},{\"name\":\"adjustmentSpeed\",\"type\":\"uint256\"}],\"name\":\"setValidatorScoreParameters\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getMaxGroupSize\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCommissionUpdateDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setGroupLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"setValidatorLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"registerValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorScoreParameters\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMembershipHistory\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uptime\",\"type\":\"uint256\"}],\"name\":\"calculateEpochScore\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"uptimes\",\"type\":\"uint256[]\"}],\"name\":\"calculateGroupEpochScore\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"uptime\",\"type\":\"uint256\"}],\"name\":\"updateValidatorScoreFromSigner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"maxPayment\",\"type\":\"uint256\"}],\"name\":\"distributeEpochPaymentsFromSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deregisterValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"group\",\"type\":\"address\"}],\"name\":\"affiliate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deaffiliate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"updateBlsPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"}],\"name\":\"updateEcdsaPublicKey\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"updatePublicKeys\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"registerValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deregisterValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"addMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"lesser\",\"type\":\"address\"},{\"name\":\"greater\",\"type\":\"address\"}],\"name\":\"addFirstMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"removeMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validator\",\"type\":\"address\"},{\"name\":\"lesserMember\",\"type\":\"address\"},{\"name\":\"greaterMember\",\"type\":\"address\"}],\"name\":\"reorderMember\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"commission\",\"type\":\"uint256\"}],\"name\":\"setNextCommissionUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateCommission\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLockedGoldRequirement\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"meetsAccountLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getValidatorBlsPublicKeyFromSigner\",\"outputs\":[{\"name\":\"blsPublicKey\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidator\",\"outputs\":[{\"name\":\"ecdsaPublicKey\",\"type\":\"bytes\"},{\"name\":\"blsPublicKey\",\"type\":\"bytes\"},{\"name\":\"affiliation\",\"type\":\"address\"},{\"name\":\"score\",\"type\":\"uint256\"},{\"name\":\"signer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256[]\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getGroupNumMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"n\",\"type\":\"uint256\"}],\"name\":\"getTopGroupValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\"}],\"name\":\"getGroupsNumMembers\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumRegisteredValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getValidatorLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getGroupLockedGoldRequirements\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredValidators\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredValidatorSigners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getRegisteredValidatorGroups\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidatorGroup\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isValidator\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"getMembershipInLastEpochFromSigner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMembershipInLastEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"validatorAccount\",\"type\":\"address\"}],\"name\":\"forceDeaffiliateIfValidator\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setSlashingMultiplierResetPeriod\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetSlashingMultiplier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"halveSlashingMultiplier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getValidatorGroupSlashingMultiplier\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"},{\"name\":\"epochNumber\",\"type\":\"uint256\"},{\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"groupMembershipInEpoch\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseValidatorsABI parses the ABI
func ParseValidatorsABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(ValidatorsABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// CalculateEpochScore is a free data retrieval call binding the contract method 0x94903a97.
//
// Solidity: function calculateEpochScore(uint256 uptime) constant returns(uint256)
func (_Validators *ValidatorsCaller) CalculateEpochScore(opts *bind.CallOpts, uptime *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "calculateEpochScore", uptime)
	return *ret0, err
}

// CalculateEpochScore is a free data retrieval call binding the contract method 0x94903a97.
//
// Solidity: function calculateEpochScore(uint256 uptime) constant returns(uint256)
func (_Validators *ValidatorsSession) CalculateEpochScore(uptime *big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateEpochScore(&_Validators.CallOpts, uptime)
}

// CalculateEpochScore is a free data retrieval call binding the contract method 0x94903a97.
//
// Solidity: function calculateEpochScore(uint256 uptime) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CalculateEpochScore(uptime *big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateEpochScore(&_Validators.CallOpts, uptime)
}

// CalculateGroupEpochScore is a free data retrieval call binding the contract method 0x76f7425d.
//
// Solidity: function calculateGroupEpochScore(uint256[] uptimes) constant returns(uint256)
func (_Validators *ValidatorsCaller) CalculateGroupEpochScore(opts *bind.CallOpts, uptimes []*big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "calculateGroupEpochScore", uptimes)
	return *ret0, err
}

// CalculateGroupEpochScore is a free data retrieval call binding the contract method 0x76f7425d.
//
// Solidity: function calculateGroupEpochScore(uint256[] uptimes) constant returns(uint256)
func (_Validators *ValidatorsSession) CalculateGroupEpochScore(uptimes []*big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateGroupEpochScore(&_Validators.CallOpts, uptimes)
}

// CalculateGroupEpochScore is a free data retrieval call binding the contract method 0x76f7425d.
//
// Solidity: function calculateGroupEpochScore(uint256[] uptimes) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CalculateGroupEpochScore(uptimes []*big.Int) (*big.Int, error) {
	return _Validators.Contract.CalculateGroupEpochScore(&_Validators.CallOpts, uptimes)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Validators *ValidatorsCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "checkProofOfPossession", sender, blsKey, blsPop)
	return *ret0, err
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Validators *ValidatorsSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Validators.Contract.CheckProofOfPossession(&_Validators.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) constant returns(bool)
func (_Validators *ValidatorsCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Validators.Contract.CheckProofOfPossession(&_Validators.CallOpts, sender, blsKey, blsPop)
}

// CommissionUpdateDelay is a free data retrieval call binding the contract method 0xe0e3ffe6.
//
// Solidity: function commissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCaller) CommissionUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "commissionUpdateDelay")
	return *ret0, err
}

// CommissionUpdateDelay is a free data retrieval call binding the contract method 0xe0e3ffe6.
//
// Solidity: function commissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsSession) CommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.CommissionUpdateDelay(&_Validators.CallOpts)
}

// CommissionUpdateDelay is a free data retrieval call binding the contract method 0xe0e3ffe6.
//
// Solidity: function commissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) CommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.CommissionUpdateDelay(&_Validators.CallOpts)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
	return *ret0, *ret1, err
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Validators.Contract.FractionMulExp(&_Validators.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Validators.Contract.FractionMulExp(&_Validators.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetAccountLockedGoldRequirement is a free data retrieval call binding the contract method 0xdcff4cf6.
//
// Solidity: function getAccountLockedGoldRequirement(address account) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetAccountLockedGoldRequirement(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getAccountLockedGoldRequirement", account)
	return *ret0, err
}

// GetAccountLockedGoldRequirement is a free data retrieval call binding the contract method 0xdcff4cf6.
//
// Solidity: function getAccountLockedGoldRequirement(address account) constant returns(uint256)
func (_Validators *ValidatorsSession) GetAccountLockedGoldRequirement(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetAccountLockedGoldRequirement(&_Validators.CallOpts, account)
}

// GetAccountLockedGoldRequirement is a free data retrieval call binding the contract method 0xdcff4cf6.
//
// Solidity: function getAccountLockedGoldRequirement(address account) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetAccountLockedGoldRequirement(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetAccountLockedGoldRequirement(&_Validators.CallOpts, account)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getBlockNumberFromHeader", header)
	return *ret0, err
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Validators *ValidatorsSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Validators.Contract.GetBlockNumberFromHeader(&_Validators.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Validators.Contract.GetBlockNumberFromHeader(&_Validators.CallOpts, header)
}

// GetCommissionUpdateDelay is a free data retrieval call binding the contract method 0xb915f530.
//
// Solidity: function getCommissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetCommissionUpdateDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getCommissionUpdateDelay")
	return *ret0, err
}

// GetCommissionUpdateDelay is a free data retrieval call binding the contract method 0xb915f530.
//
// Solidity: function getCommissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsSession) GetCommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.GetCommissionUpdateDelay(&_Validators.CallOpts)
}

// GetCommissionUpdateDelay is a free data retrieval call binding the contract method 0xb915f530.
//
// Solidity: function getCommissionUpdateDelay() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetCommissionUpdateDelay() (*big.Int, error) {
	return _Validators.Contract.GetCommissionUpdateDelay(&_Validators.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getEpochNumber")
	return *ret0, err
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Validators *ValidatorsSession) GetEpochNumber() (*big.Int, error) {
	return _Validators.Contract.GetEpochNumber(&_Validators.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Validators.Contract.GetEpochNumber(&_Validators.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getEpochNumberOfBlock", blockNumber)
	return *ret0, err
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.GetEpochNumberOfBlock(&_Validators.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.GetEpochNumberOfBlock(&_Validators.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getEpochSize")
	return *ret0, err
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Validators *ValidatorsSession) GetEpochSize() (*big.Int, error) {
	return _Validators.Contract.GetEpochSize(&_Validators.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetEpochSize() (*big.Int, error) {
	return _Validators.Contract.GetEpochSize(&_Validators.CallOpts)
}

// GetGroupLockedGoldRequirements is a free data retrieval call binding the contract method 0x6fa47647.
//
// Solidity: function getGroupLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) GetGroupLockedGoldRequirements(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "getGroupLockedGoldRequirements")
	return *ret0, *ret1, err
}

// GetGroupLockedGoldRequirements is a free data retrieval call binding the contract method 0x6fa47647.
//
// Solidity: function getGroupLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) GetGroupLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetGroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GetGroupLockedGoldRequirements is a free data retrieval call binding the contract method 0x6fa47647.
//
// Solidity: function getGroupLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetGroupLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetGroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GetGroupNumMembers is a free data retrieval call binding the contract method 0x39e618e8.
//
// Solidity: function getGroupNumMembers(address account) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetGroupNumMembers(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getGroupNumMembers", account)
	return *ret0, err
}

// GetGroupNumMembers is a free data retrieval call binding the contract method 0x39e618e8.
//
// Solidity: function getGroupNumMembers(address account) constant returns(uint256)
func (_Validators *ValidatorsSession) GetGroupNumMembers(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetGroupNumMembers(&_Validators.CallOpts, account)
}

// GetGroupNumMembers is a free data retrieval call binding the contract method 0x39e618e8.
//
// Solidity: function getGroupNumMembers(address account) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetGroupNumMembers(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetGroupNumMembers(&_Validators.CallOpts, account)
}

// GetGroupsNumMembers is a free data retrieval call binding the contract method 0x70447754.
//
// Solidity: function getGroupsNumMembers(address[] accounts) constant returns(uint256[])
func (_Validators *ValidatorsCaller) GetGroupsNumMembers(opts *bind.CallOpts, accounts []common.Address) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getGroupsNumMembers", accounts)
	return *ret0, err
}

// GetGroupsNumMembers is a free data retrieval call binding the contract method 0x70447754.
//
// Solidity: function getGroupsNumMembers(address[] accounts) constant returns(uint256[])
func (_Validators *ValidatorsSession) GetGroupsNumMembers(accounts []common.Address) ([]*big.Int, error) {
	return _Validators.Contract.GetGroupsNumMembers(&_Validators.CallOpts, accounts)
}

// GetGroupsNumMembers is a free data retrieval call binding the contract method 0x70447754.
//
// Solidity: function getGroupsNumMembers(address[] accounts) constant returns(uint256[])
func (_Validators *ValidatorsCallerSession) GetGroupsNumMembers(accounts []common.Address) ([]*big.Int, error) {
	return _Validators.Contract.GetGroupsNumMembers(&_Validators.CallOpts, accounts)
}

// GetMaxGroupSize is a free data retrieval call binding the contract method 0x43d96699.
//
// Solidity: function getMaxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetMaxGroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getMaxGroupSize")
	return *ret0, err
}

// GetMaxGroupSize is a free data retrieval call binding the contract method 0x43d96699.
//
// Solidity: function getMaxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsSession) GetMaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.GetMaxGroupSize(&_Validators.CallOpts)
}

// GetMaxGroupSize is a free data retrieval call binding the contract method 0x43d96699.
//
// Solidity: function getMaxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetMaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.GetMaxGroupSize(&_Validators.CallOpts)
}

// GetMembershipHistory is a free data retrieval call binding the contract method 0x35244f51.
//
// Solidity: function getMembershipHistory(address account) constant returns(uint256[], address[], uint256, uint256)
func (_Validators *ValidatorsCaller) GetMembershipHistory(opts *bind.CallOpts, account common.Address) ([]*big.Int, []common.Address, *big.Int, *big.Int, error) {
	var (
		ret0 = new([]*big.Int)
		ret1 = new([]common.Address)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
	}
	err := _Validators.contract.Call(opts, out, "getMembershipHistory", account)
	return *ret0, *ret1, *ret2, *ret3, err
}

// GetMembershipHistory is a free data retrieval call binding the contract method 0x35244f51.
//
// Solidity: function getMembershipHistory(address account) constant returns(uint256[], address[], uint256, uint256)
func (_Validators *ValidatorsSession) GetMembershipHistory(account common.Address) ([]*big.Int, []common.Address, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetMembershipHistory(&_Validators.CallOpts, account)
}

// GetMembershipHistory is a free data retrieval call binding the contract method 0x35244f51.
//
// Solidity: function getMembershipHistory(address account) constant returns(uint256[], address[], uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetMembershipHistory(account common.Address) ([]*big.Int, []common.Address, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetMembershipHistory(&_Validators.CallOpts, account)
}

// GetMembershipInLastEpoch is a free data retrieval call binding the contract method 0x0d1312b8.
//
// Solidity: function getMembershipInLastEpoch(address account) constant returns(address)
func (_Validators *ValidatorsCaller) GetMembershipInLastEpoch(opts *bind.CallOpts, account common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getMembershipInLastEpoch", account)
	return *ret0, err
}

// GetMembershipInLastEpoch is a free data retrieval call binding the contract method 0x0d1312b8.
//
// Solidity: function getMembershipInLastEpoch(address account) constant returns(address)
func (_Validators *ValidatorsSession) GetMembershipInLastEpoch(account common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpoch(&_Validators.CallOpts, account)
}

// GetMembershipInLastEpoch is a free data retrieval call binding the contract method 0x0d1312b8.
//
// Solidity: function getMembershipInLastEpoch(address account) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetMembershipInLastEpoch(account common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpoch(&_Validators.CallOpts, account)
}

// GetMembershipInLastEpochFromSigner is a free data retrieval call binding the contract method 0x51b52225.
//
// Solidity: function getMembershipInLastEpochFromSigner(address signer) constant returns(address)
func (_Validators *ValidatorsCaller) GetMembershipInLastEpochFromSigner(opts *bind.CallOpts, signer common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getMembershipInLastEpochFromSigner", signer)
	return *ret0, err
}

// GetMembershipInLastEpochFromSigner is a free data retrieval call binding the contract method 0x51b52225.
//
// Solidity: function getMembershipInLastEpochFromSigner(address signer) constant returns(address)
func (_Validators *ValidatorsSession) GetMembershipInLastEpochFromSigner(signer common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpochFromSigner(&_Validators.CallOpts, signer)
}

// GetMembershipInLastEpochFromSigner is a free data retrieval call binding the contract method 0x51b52225.
//
// Solidity: function getMembershipInLastEpochFromSigner(address signer) constant returns(address)
func (_Validators *ValidatorsCallerSession) GetMembershipInLastEpochFromSigner(signer common.Address) (common.Address, error) {
	return _Validators.Contract.GetMembershipInLastEpochFromSigner(&_Validators.CallOpts, signer)
}

// GetNumRegisteredValidators is a free data retrieval call binding the contract method 0x517f6d33.
//
// Solidity: function getNumRegisteredValidators() constant returns(uint256)
func (_Validators *ValidatorsCaller) GetNumRegisteredValidators(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getNumRegisteredValidators")
	return *ret0, err
}

// GetNumRegisteredValidators is a free data retrieval call binding the contract method 0x517f6d33.
//
// Solidity: function getNumRegisteredValidators() constant returns(uint256)
func (_Validators *ValidatorsSession) GetNumRegisteredValidators() (*big.Int, error) {
	return _Validators.Contract.GetNumRegisteredValidators(&_Validators.CallOpts)
}

// GetNumRegisteredValidators is a free data retrieval call binding the contract method 0x517f6d33.
//
// Solidity: function getNumRegisteredValidators() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetNumRegisteredValidators() (*big.Int, error) {
	return _Validators.Contract.GetNumRegisteredValidators(&_Validators.CallOpts)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Validators *ValidatorsCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getParentSealBitmap", blockNumber)
	return *ret0, err
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Validators *ValidatorsSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Validators.Contract.GetParentSealBitmap(&_Validators.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) constant returns(bytes32)
func (_Validators *ValidatorsCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Validators.Contract.GetParentSealBitmap(&_Validators.CallOpts, blockNumber)
}

// GetRegisteredValidatorGroups is a free data retrieval call binding the contract method 0x3f270898.
//
// Solidity: function getRegisteredValidatorGroups() constant returns(address[])
func (_Validators *ValidatorsCaller) GetRegisteredValidatorGroups(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getRegisteredValidatorGroups")
	return *ret0, err
}

// GetRegisteredValidatorGroups is a free data retrieval call binding the contract method 0x3f270898.
//
// Solidity: function getRegisteredValidatorGroups() constant returns(address[])
func (_Validators *ValidatorsSession) GetRegisteredValidatorGroups() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorGroups(&_Validators.CallOpts)
}

// GetRegisteredValidatorGroups is a free data retrieval call binding the contract method 0x3f270898.
//
// Solidity: function getRegisteredValidatorGroups() constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetRegisteredValidatorGroups() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorGroups(&_Validators.CallOpts)
}

// GetRegisteredValidatorSigners is a free data retrieval call binding the contract method 0xd55dcbcf.
//
// Solidity: function getRegisteredValidatorSigners() constant returns(address[])
func (_Validators *ValidatorsCaller) GetRegisteredValidatorSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getRegisteredValidatorSigners")
	return *ret0, err
}

// GetRegisteredValidatorSigners is a free data retrieval call binding the contract method 0xd55dcbcf.
//
// Solidity: function getRegisteredValidatorSigners() constant returns(address[])
func (_Validators *ValidatorsSession) GetRegisteredValidatorSigners() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorSigners(&_Validators.CallOpts)
}

// GetRegisteredValidatorSigners is a free data retrieval call binding the contract method 0xd55dcbcf.
//
// Solidity: function getRegisteredValidatorSigners() constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetRegisteredValidatorSigners() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidatorSigners(&_Validators.CallOpts)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xd93ab5ad.
//
// Solidity: function getRegisteredValidators() constant returns(address[])
func (_Validators *ValidatorsCaller) GetRegisteredValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getRegisteredValidators")
	return *ret0, err
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xd93ab5ad.
//
// Solidity: function getRegisteredValidators() constant returns(address[])
func (_Validators *ValidatorsSession) GetRegisteredValidators() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidators(&_Validators.CallOpts)
}

// GetRegisteredValidators is a free data retrieval call binding the contract method 0xd93ab5ad.
//
// Solidity: function getRegisteredValidators() constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetRegisteredValidators() ([]common.Address, error) {
	return _Validators.Contract.GetRegisteredValidators(&_Validators.CallOpts)
}

// GetTopGroupValidators is a free data retrieval call binding the contract method 0x8dd31e39.
//
// Solidity: function getTopGroupValidators(address account, uint256 n) constant returns(address[])
func (_Validators *ValidatorsCaller) GetTopGroupValidators(opts *bind.CallOpts, account common.Address, n *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getTopGroupValidators", account, n)
	return *ret0, err
}

// GetTopGroupValidators is a free data retrieval call binding the contract method 0x8dd31e39.
//
// Solidity: function getTopGroupValidators(address account, uint256 n) constant returns(address[])
func (_Validators *ValidatorsSession) GetTopGroupValidators(account common.Address, n *big.Int) ([]common.Address, error) {
	return _Validators.Contract.GetTopGroupValidators(&_Validators.CallOpts, account, n)
}

// GetTopGroupValidators is a free data retrieval call binding the contract method 0x8dd31e39.
//
// Solidity: function getTopGroupValidators(address account, uint256 n) constant returns(address[])
func (_Validators *ValidatorsCallerSession) GetTopGroupValidators(account common.Address, n *big.Int) ([]common.Address, error) {
	return _Validators.Contract.GetTopGroupValidators(&_Validators.CallOpts, account, n)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address account) constant returns(bytes ecdsaPublicKey, bytes blsPublicKey, address affiliation, uint256 score, address signer)
func (_Validators *ValidatorsCaller) GetValidator(opts *bind.CallOpts, account common.Address) (struct {
	EcdsaPublicKey []byte
	BlsPublicKey   []byte
	Affiliation    common.Address
	Score          *big.Int
	Signer         common.Address
}, error) {
	ret := new(struct {
		EcdsaPublicKey []byte
		BlsPublicKey   []byte
		Affiliation    common.Address
		Score          *big.Int
		Signer         common.Address
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "getValidator", account)
	return *ret, err
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address account) constant returns(bytes ecdsaPublicKey, bytes blsPublicKey, address affiliation, uint256 score, address signer)
func (_Validators *ValidatorsSession) GetValidator(account common.Address) (struct {
	EcdsaPublicKey []byte
	BlsPublicKey   []byte
	Affiliation    common.Address
	Score          *big.Int
	Signer         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, account)
}

// GetValidator is a free data retrieval call binding the contract method 0x1904bb2e.
//
// Solidity: function getValidator(address account) constant returns(bytes ecdsaPublicKey, bytes blsPublicKey, address affiliation, uint256 score, address signer)
func (_Validators *ValidatorsCallerSession) GetValidator(account common.Address) (struct {
	EcdsaPublicKey []byte
	BlsPublicKey   []byte
	Affiliation    common.Address
	Score          *big.Int
	Signer         common.Address
}, error) {
	return _Validators.Contract.GetValidator(&_Validators.CallOpts, account)
}

// GetValidatorBlsPublicKeyFromSigner is a free data retrieval call binding the contract method 0xb730a299.
//
// Solidity: function getValidatorBlsPublicKeyFromSigner(address signer) constant returns(bytes blsPublicKey)
func (_Validators *ValidatorsCaller) GetValidatorBlsPublicKeyFromSigner(opts *bind.CallOpts, signer common.Address) ([]byte, error) {
	var (
		ret0 = new([]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getValidatorBlsPublicKeyFromSigner", signer)
	return *ret0, err
}

// GetValidatorBlsPublicKeyFromSigner is a free data retrieval call binding the contract method 0xb730a299.
//
// Solidity: function getValidatorBlsPublicKeyFromSigner(address signer) constant returns(bytes blsPublicKey)
func (_Validators *ValidatorsSession) GetValidatorBlsPublicKeyFromSigner(signer common.Address) ([]byte, error) {
	return _Validators.Contract.GetValidatorBlsPublicKeyFromSigner(&_Validators.CallOpts, signer)
}

// GetValidatorBlsPublicKeyFromSigner is a free data retrieval call binding the contract method 0xb730a299.
//
// Solidity: function getValidatorBlsPublicKeyFromSigner(address signer) constant returns(bytes blsPublicKey)
func (_Validators *ValidatorsCallerSession) GetValidatorBlsPublicKeyFromSigner(signer common.Address) ([]byte, error) {
	return _Validators.Contract.GetValidatorBlsPublicKeyFromSigner(&_Validators.CallOpts, signer)
}

// GetValidatorGroup is a free data retrieval call binding the contract method 0x9b9d5161.
//
// Solidity: function getValidatorGroup(address account) constant returns(address[], uint256, uint256, uint256, uint256[], uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorGroup(opts *bind.CallOpts, account common.Address) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new(*big.Int)
		ret2 = new(*big.Int)
		ret3 = new(*big.Int)
		ret4 = new([]*big.Int)
		ret5 = new(*big.Int)
		ret6 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
		ret5,
		ret6,
	}
	err := _Validators.contract.Call(opts, out, "getValidatorGroup", account)
	return *ret0, *ret1, *ret2, *ret3, *ret4, *ret5, *ret6, err
}

// GetValidatorGroup is a free data retrieval call binding the contract method 0x9b9d5161.
//
// Solidity: function getValidatorGroup(address account) constant returns(address[], uint256, uint256, uint256, uint256[], uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorGroup(account common.Address) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorGroup(&_Validators.CallOpts, account)
}

// GetValidatorGroup is a free data retrieval call binding the contract method 0x9b9d5161.
//
// Solidity: function getValidatorGroup(address account) constant returns(address[], uint256, uint256, uint256, uint256[], uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorGroup(account common.Address) ([]common.Address, *big.Int, *big.Int, *big.Int, []*big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorGroup(&_Validators.CallOpts, account)
}

// GetValidatorGroupSlashingMultiplier is a free data retrieval call binding the contract method 0xdba94fcd.
//
// Solidity: function getValidatorGroupSlashingMultiplier(address account) constant returns(uint256)
func (_Validators *ValidatorsCaller) GetValidatorGroupSlashingMultiplier(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getValidatorGroupSlashingMultiplier", account)
	return *ret0, err
}

// GetValidatorGroupSlashingMultiplier is a free data retrieval call binding the contract method 0xdba94fcd.
//
// Solidity: function getValidatorGroupSlashingMultiplier(address account) constant returns(uint256)
func (_Validators *ValidatorsSession) GetValidatorGroupSlashingMultiplier(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetValidatorGroupSlashingMultiplier(&_Validators.CallOpts, account)
}

// GetValidatorGroupSlashingMultiplier is a free data retrieval call binding the contract method 0xdba94fcd.
//
// Solidity: function getValidatorGroupSlashingMultiplier(address account) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorGroupSlashingMultiplier(account common.Address) (*big.Int, error) {
	return _Validators.Contract.GetValidatorGroupSlashingMultiplier(&_Validators.CallOpts, account)
}

// GetValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xc10c96ef.
//
// Solidity: function getValidatorLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorLockedGoldRequirements(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "getValidatorLockedGoldRequirements")
	return *ret0, *ret1, err
}

// GetValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xc10c96ef.
//
// Solidity: function getValidatorLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// GetValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xc10c96ef.
//
// Solidity: function getValidatorLockedGoldRequirements() constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorLockedGoldRequirements() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// GetValidatorScoreParameters is a free data retrieval call binding the contract method 0x19113e3b.
//
// Solidity: function getValidatorScoreParameters() constant returns(uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorScoreParameters(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Validators.contract.Call(opts, out, "getValidatorScoreParameters")
	return *ret0, *ret1, err
}

// GetValidatorScoreParameters is a free data retrieval call binding the contract method 0x19113e3b.
//
// Solidity: function getValidatorScoreParameters() constant returns(uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorScoreParameters() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorScoreParameters(&_Validators.CallOpts)
}

// GetValidatorScoreParameters is a free data retrieval call binding the contract method 0x19113e3b.
//
// Solidity: function getValidatorScoreParameters() constant returns(uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorScoreParameters() (*big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorScoreParameters(&_Validators.CallOpts)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "getVerifiedSealBitmapFromHeader", header)
	return *ret0, err
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.GetVerifiedSealBitmapFromHeader(&_Validators.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.GetVerifiedSealBitmapFromHeader(&_Validators.CallOpts, header)
}

// GroupLockedGoldRequirements is a free data retrieval call binding the contract method 0xc5805140.
//
// Solidity: function groupLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCaller) GroupLockedGoldRequirements(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	ret := new(struct {
		Value    *big.Int
		Duration *big.Int
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "groupLockedGoldRequirements")
	return *ret, err
}

// GroupLockedGoldRequirements is a free data retrieval call binding the contract method 0xc5805140.
//
// Solidity: function groupLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsSession) GroupLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.GroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GroupLockedGoldRequirements is a free data retrieval call binding the contract method 0xc5805140.
//
// Solidity: function groupLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCallerSession) GroupLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.GroupLockedGoldRequirements(&_Validators.CallOpts)
}

// GroupMembershipInEpoch is a free data retrieval call binding the contract method 0xeb1d0b42.
//
// Solidity: function groupMembershipInEpoch(address account, uint256 epochNumber, uint256 index) constant returns(address)
func (_Validators *ValidatorsCaller) GroupMembershipInEpoch(opts *bind.CallOpts, account common.Address, epochNumber *big.Int, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "groupMembershipInEpoch", account, epochNumber, index)
	return *ret0, err
}

// GroupMembershipInEpoch is a free data retrieval call binding the contract method 0xeb1d0b42.
//
// Solidity: function groupMembershipInEpoch(address account, uint256 epochNumber, uint256 index) constant returns(address)
func (_Validators *ValidatorsSession) GroupMembershipInEpoch(account common.Address, epochNumber *big.Int, index *big.Int) (common.Address, error) {
	return _Validators.Contract.GroupMembershipInEpoch(&_Validators.CallOpts, account, epochNumber, index)
}

// GroupMembershipInEpoch is a free data retrieval call binding the contract method 0xeb1d0b42.
//
// Solidity: function groupMembershipInEpoch(address account, uint256 epochNumber, uint256 index) constant returns(address)
func (_Validators *ValidatorsCallerSession) GroupMembershipInEpoch(account common.Address, epochNumber *big.Int, index *big.Int) (common.Address, error) {
	return _Validators.Contract.GroupMembershipInEpoch(&_Validators.CallOpts, account, epochNumber, index)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "hashHeader", header)
	return *ret0, err
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsSession) HashHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.HashHeader(&_Validators.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) constant returns(bytes32)
func (_Validators *ValidatorsCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Validators.Contract.HashHeader(&_Validators.CallOpts, header)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Validators *ValidatorsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "initialized")
	return *ret0, err
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Validators *ValidatorsSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() constant returns(bool)
func (_Validators *ValidatorsCallerSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Validators *ValidatorsCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Validators *ValidatorsSession) IsOwner() (bool, error) {
	return _Validators.Contract.IsOwner(&_Validators.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsOwner() (bool, error) {
	return _Validators.Contract.IsOwner(&_Validators.CallOpts)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) constant returns(bool)
func (_Validators *ValidatorsCaller) IsValidator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isValidator", account)
	return *ret0, err
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) constant returns(bool)
func (_Validators *ValidatorsSession) IsValidator(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidator(&_Validators.CallOpts, account)
}

// IsValidator is a free data retrieval call binding the contract method 0xfacd743b.
//
// Solidity: function isValidator(address account) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsValidator(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidator(&_Validators.CallOpts, account)
}

// IsValidatorGroup is a free data retrieval call binding the contract method 0x52f13a4e.
//
// Solidity: function isValidatorGroup(address account) constant returns(bool)
func (_Validators *ValidatorsCaller) IsValidatorGroup(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "isValidatorGroup", account)
	return *ret0, err
}

// IsValidatorGroup is a free data retrieval call binding the contract method 0x52f13a4e.
//
// Solidity: function isValidatorGroup(address account) constant returns(bool)
func (_Validators *ValidatorsSession) IsValidatorGroup(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidatorGroup(&_Validators.CallOpts, account)
}

// IsValidatorGroup is a free data retrieval call binding the contract method 0x52f13a4e.
//
// Solidity: function isValidatorGroup(address account) constant returns(bool)
func (_Validators *ValidatorsCallerSession) IsValidatorGroup(account common.Address) (bool, error) {
	return _Validators.Contract.IsValidatorGroup(&_Validators.CallOpts, account)
}

// MaxGroupSize is a free data retrieval call binding the contract method 0x5779e93d.
//
// Solidity: function maxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCaller) MaxGroupSize(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "maxGroupSize")
	return *ret0, err
}

// MaxGroupSize is a free data retrieval call binding the contract method 0x5779e93d.
//
// Solidity: function maxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsSession) MaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.MaxGroupSize(&_Validators.CallOpts)
}

// MaxGroupSize is a free data retrieval call binding the contract method 0x5779e93d.
//
// Solidity: function maxGroupSize() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MaxGroupSize() (*big.Int, error) {
	return _Validators.Contract.MaxGroupSize(&_Validators.CallOpts)
}

// MeetsAccountLockedGoldRequirements is a free data retrieval call binding the contract method 0xc54c1cd4.
//
// Solidity: function meetsAccountLockedGoldRequirements(address account) constant returns(bool)
func (_Validators *ValidatorsCaller) MeetsAccountLockedGoldRequirements(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "meetsAccountLockedGoldRequirements", account)
	return *ret0, err
}

// MeetsAccountLockedGoldRequirements is a free data retrieval call binding the contract method 0xc54c1cd4.
//
// Solidity: function meetsAccountLockedGoldRequirements(address account) constant returns(bool)
func (_Validators *ValidatorsSession) MeetsAccountLockedGoldRequirements(account common.Address) (bool, error) {
	return _Validators.Contract.MeetsAccountLockedGoldRequirements(&_Validators.CallOpts, account)
}

// MeetsAccountLockedGoldRequirements is a free data retrieval call binding the contract method 0xc54c1cd4.
//
// Solidity: function meetsAccountLockedGoldRequirements(address account) constant returns(bool)
func (_Validators *ValidatorsCallerSession) MeetsAccountLockedGoldRequirements(account common.Address) (bool, error) {
	return _Validators.Contract.MeetsAccountLockedGoldRequirements(&_Validators.CallOpts, account)
}

// MembershipHistoryLength is a free data retrieval call binding the contract method 0x4cd76db4.
//
// Solidity: function membershipHistoryLength() constant returns(uint256)
func (_Validators *ValidatorsCaller) MembershipHistoryLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "membershipHistoryLength")
	return *ret0, err
}

// MembershipHistoryLength is a free data retrieval call binding the contract method 0x4cd76db4.
//
// Solidity: function membershipHistoryLength() constant returns(uint256)
func (_Validators *ValidatorsSession) MembershipHistoryLength() (*big.Int, error) {
	return _Validators.Contract.MembershipHistoryLength(&_Validators.CallOpts)
}

// MembershipHistoryLength is a free data retrieval call binding the contract method 0x4cd76db4.
//
// Solidity: function membershipHistoryLength() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MembershipHistoryLength() (*big.Int, error) {
	return _Validators.Contract.MembershipHistoryLength(&_Validators.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "minQuorumSize", blockNumber)
	return *ret0, err
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.MinQuorumSize(&_Validators.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.MinQuorumSize(&_Validators.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "minQuorumSizeInCurrentSet")
	return *ret0, err
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.MinQuorumSizeInCurrentSet(&_Validators.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.MinQuorumSizeInCurrentSet(&_Validators.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "numberValidatorsInCurrentSet")
	return *ret0, err
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInCurrentSet(&_Validators.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInCurrentSet(&_Validators.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "numberValidatorsInSet", blockNumber)
	return *ret0, err
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInSet(&_Validators.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) constant returns(uint256)
func (_Validators *ValidatorsCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Validators.Contract.NumberValidatorsInSet(&_Validators.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsSession) Owner() (common.Address, error) {
	return _Validators.Contract.Owner(&_Validators.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Validators *ValidatorsCallerSession) Owner() (common.Address, error) {
	return _Validators.Contract.Owner(&_Validators.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Validators *ValidatorsCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "registry")
	return *ret0, err
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Validators *ValidatorsSession) Registry() (common.Address, error) {
	return _Validators.Contract.Registry(&_Validators.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() constant returns(address)
func (_Validators *ValidatorsCallerSession) Registry() (common.Address, error) {
	return _Validators.Contract.Registry(&_Validators.CallOpts)
}

// SlashingMultiplierResetPeriod is a free data retrieval call binding the contract method 0x36407b70.
//
// Solidity: function slashingMultiplierResetPeriod() constant returns(uint256)
func (_Validators *ValidatorsCaller) SlashingMultiplierResetPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "slashingMultiplierResetPeriod")
	return *ret0, err
}

// SlashingMultiplierResetPeriod is a free data retrieval call binding the contract method 0x36407b70.
//
// Solidity: function slashingMultiplierResetPeriod() constant returns(uint256)
func (_Validators *ValidatorsSession) SlashingMultiplierResetPeriod() (*big.Int, error) {
	return _Validators.Contract.SlashingMultiplierResetPeriod(&_Validators.CallOpts)
}

// SlashingMultiplierResetPeriod is a free data retrieval call binding the contract method 0x36407b70.
//
// Solidity: function slashingMultiplierResetPeriod() constant returns(uint256)
func (_Validators *ValidatorsCallerSession) SlashingMultiplierResetPeriod() (*big.Int, error) {
	return _Validators.Contract.SlashingMultiplierResetPeriod(&_Validators.CallOpts)
}

// ValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xbd9e9d94.
//
// Solidity: function validatorLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCaller) ValidatorLockedGoldRequirements(opts *bind.CallOpts) (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	ret := new(struct {
		Value    *big.Int
		Duration *big.Int
	})
	out := ret
	err := _Validators.contract.Call(opts, out, "validatorLockedGoldRequirements")
	return *ret, err
}

// ValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xbd9e9d94.
//
// Solidity: function validatorLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsSession) ValidatorLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.ValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// ValidatorLockedGoldRequirements is a free data retrieval call binding the contract method 0xbd9e9d94.
//
// Solidity: function validatorLockedGoldRequirements() constant returns(uint256 value, uint256 duration)
func (_Validators *ValidatorsCallerSession) ValidatorLockedGoldRequirements() (struct {
	Value    *big.Int
	Duration *big.Int
}, error) {
	return _Validators.Contract.ValidatorLockedGoldRequirements(&_Validators.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Validators *ValidatorsCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "validatorSignerAddressFromCurrentSet", index)
	return *ret0, err
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Validators *ValidatorsSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromCurrentSet(&_Validators.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) constant returns(address)
func (_Validators *ValidatorsCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromCurrentSet(&_Validators.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Validators *ValidatorsCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Validators.contract.Call(opts, out, "validatorSignerAddressFromSet", index, blockNumber)
	return *ret0, err
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Validators *ValidatorsSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromSet(&_Validators.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) constant returns(address)
func (_Validators *ValidatorsCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Validators.Contract.ValidatorSignerAddressFromSet(&_Validators.CallOpts, index, blockNumber)
}

// AddFirstMember is a paid mutator transaction binding the contract method 0x3173b8db.
//
// Solidity: function addFirstMember(address validator, address lesser, address greater) returns(bool)
func (_Validators *ValidatorsTransactor) AddFirstMember(opts *bind.TransactOpts, validator common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addFirstMember", validator, lesser, greater)
}

// AddFirstMember is a paid mutator transaction binding the contract method 0x3173b8db.
//
// Solidity: function addFirstMember(address validator, address lesser, address greater) returns(bool)
func (_Validators *ValidatorsSession) AddFirstMember(validator common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddFirstMember(&_Validators.TransactOpts, validator, lesser, greater)
}

// AddFirstMember is a paid mutator transaction binding the contract method 0x3173b8db.
//
// Solidity: function addFirstMember(address validator, address lesser, address greater) returns(bool)
func (_Validators *ValidatorsTransactorSession) AddFirstMember(validator common.Address, lesser common.Address, greater common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddFirstMember(&_Validators.TransactOpts, validator, lesser, greater)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) AddMember(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "addMember", validator)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address validator) returns(bool)
func (_Validators *ValidatorsSession) AddMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddMember(&_Validators.TransactOpts, validator)
}

// AddMember is a paid mutator transaction binding the contract method 0xca6d56dc.
//
// Solidity: function addMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) AddMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.AddMember(&_Validators.TransactOpts, validator)
}

// Affiliate is a paid mutator transaction binding the contract method 0xb591d3a5.
//
// Solidity: function affiliate(address group) returns(bool)
func (_Validators *ValidatorsTransactor) Affiliate(opts *bind.TransactOpts, group common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "affiliate", group)
}

// Affiliate is a paid mutator transaction binding the contract method 0xb591d3a5.
//
// Solidity: function affiliate(address group) returns(bool)
func (_Validators *ValidatorsSession) Affiliate(group common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Affiliate(&_Validators.TransactOpts, group)
}

// Affiliate is a paid mutator transaction binding the contract method 0xb591d3a5.
//
// Solidity: function affiliate(address group) returns(bool)
func (_Validators *ValidatorsTransactorSession) Affiliate(group common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Affiliate(&_Validators.TransactOpts, group)
}

// Deaffiliate is a paid mutator transaction binding the contract method 0xfffdfccb.
//
// Solidity: function deaffiliate() returns(bool)
func (_Validators *ValidatorsTransactor) Deaffiliate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "deaffiliate")
}

// Deaffiliate is a paid mutator transaction binding the contract method 0xfffdfccb.
//
// Solidity: function deaffiliate() returns(bool)
func (_Validators *ValidatorsSession) Deaffiliate() (*types.Transaction, error) {
	return _Validators.Contract.Deaffiliate(&_Validators.TransactOpts)
}

// Deaffiliate is a paid mutator transaction binding the contract method 0xfffdfccb.
//
// Solidity: function deaffiliate() returns(bool)
func (_Validators *ValidatorsTransactorSession) Deaffiliate() (*types.Transaction, error) {
	return _Validators.Contract.Deaffiliate(&_Validators.TransactOpts)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x8b16b1c6.
//
// Solidity: function deregisterValidator(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactor) DeregisterValidator(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "deregisterValidator", index)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x8b16b1c6.
//
// Solidity: function deregisterValidator(uint256 index) returns(bool)
func (_Validators *ValidatorsSession) DeregisterValidator(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidator(&_Validators.TransactOpts, index)
}

// DeregisterValidator is a paid mutator transaction binding the contract method 0x8b16b1c6.
//
// Solidity: function deregisterValidator(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactorSession) DeregisterValidator(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidator(&_Validators.TransactOpts, index)
}

// DeregisterValidatorGroup is a paid mutator transaction binding the contract method 0x60fb822c.
//
// Solidity: function deregisterValidatorGroup(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactor) DeregisterValidatorGroup(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "deregisterValidatorGroup", index)
}

// DeregisterValidatorGroup is a paid mutator transaction binding the contract method 0x60fb822c.
//
// Solidity: function deregisterValidatorGroup(uint256 index) returns(bool)
func (_Validators *ValidatorsSession) DeregisterValidatorGroup(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidatorGroup(&_Validators.TransactOpts, index)
}

// DeregisterValidatorGroup is a paid mutator transaction binding the contract method 0x60fb822c.
//
// Solidity: function deregisterValidatorGroup(uint256 index) returns(bool)
func (_Validators *ValidatorsTransactorSession) DeregisterValidatorGroup(index *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DeregisterValidatorGroup(&_Validators.TransactOpts, index)
}

// DistributeEpochPaymentsFromSigner is a paid mutator transaction binding the contract method 0xd69ef6cf.
//
// Solidity: function distributeEpochPaymentsFromSigner(address signer, uint256 maxPayment) returns(uint256)
func (_Validators *ValidatorsTransactor) DistributeEpochPaymentsFromSigner(opts *bind.TransactOpts, signer common.Address, maxPayment *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "distributeEpochPaymentsFromSigner", signer, maxPayment)
}

// DistributeEpochPaymentsFromSigner is a paid mutator transaction binding the contract method 0xd69ef6cf.
//
// Solidity: function distributeEpochPaymentsFromSigner(address signer, uint256 maxPayment) returns(uint256)
func (_Validators *ValidatorsSession) DistributeEpochPaymentsFromSigner(signer common.Address, maxPayment *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DistributeEpochPaymentsFromSigner(&_Validators.TransactOpts, signer, maxPayment)
}

// DistributeEpochPaymentsFromSigner is a paid mutator transaction binding the contract method 0xd69ef6cf.
//
// Solidity: function distributeEpochPaymentsFromSigner(address signer, uint256 maxPayment) returns(uint256)
func (_Validators *ValidatorsTransactorSession) DistributeEpochPaymentsFromSigner(signer common.Address, maxPayment *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.DistributeEpochPaymentsFromSigner(&_Validators.TransactOpts, signer, maxPayment)
}

// ForceDeaffiliateIfValidator is a paid mutator transaction binding the contract method 0xe33301aa.
//
// Solidity: function forceDeaffiliateIfValidator(address validatorAccount) returns()
func (_Validators *ValidatorsTransactor) ForceDeaffiliateIfValidator(opts *bind.TransactOpts, validatorAccount common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "forceDeaffiliateIfValidator", validatorAccount)
}

// ForceDeaffiliateIfValidator is a paid mutator transaction binding the contract method 0xe33301aa.
//
// Solidity: function forceDeaffiliateIfValidator(address validatorAccount) returns()
func (_Validators *ValidatorsSession) ForceDeaffiliateIfValidator(validatorAccount common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ForceDeaffiliateIfValidator(&_Validators.TransactOpts, validatorAccount)
}

// ForceDeaffiliateIfValidator is a paid mutator transaction binding the contract method 0xe33301aa.
//
// Solidity: function forceDeaffiliateIfValidator(address validatorAccount) returns()
func (_Validators *ValidatorsTransactorSession) ForceDeaffiliateIfValidator(validatorAccount common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ForceDeaffiliateIfValidator(&_Validators.TransactOpts, validatorAccount)
}

// HalveSlashingMultiplier is a paid mutator transaction binding the contract method 0xc22d3bba.
//
// Solidity: function halveSlashingMultiplier(address account) returns()
func (_Validators *ValidatorsTransactor) HalveSlashingMultiplier(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "halveSlashingMultiplier", account)
}

// HalveSlashingMultiplier is a paid mutator transaction binding the contract method 0xc22d3bba.
//
// Solidity: function halveSlashingMultiplier(address account) returns()
func (_Validators *ValidatorsSession) HalveSlashingMultiplier(account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.HalveSlashingMultiplier(&_Validators.TransactOpts, account)
}

// HalveSlashingMultiplier is a paid mutator transaction binding the contract method 0xc22d3bba.
//
// Solidity: function halveSlashingMultiplier(address account) returns()
func (_Validators *ValidatorsTransactorSession) HalveSlashingMultiplier(account common.Address) (*types.Transaction, error) {
	return _Validators.Contract.HalveSlashingMultiplier(&_Validators.TransactOpts, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x78d25456.
//
// Solidity: function initialize(address registryAddress, uint256 groupRequirementValue, uint256 groupRequirementDuration, uint256 validatorRequirementValue, uint256 validatorRequirementDuration, uint256 validatorScoreExponent, uint256 validatorScoreAdjustmentSpeed, uint256 _membershipHistoryLength, uint256 _slashingMultiplierResetPeriod, uint256 _maxGroupSize, uint256 _commissionUpdateDelay) returns()
func (_Validators *ValidatorsTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, groupRequirementValue *big.Int, groupRequirementDuration *big.Int, validatorRequirementValue *big.Int, validatorRequirementDuration *big.Int, validatorScoreExponent *big.Int, validatorScoreAdjustmentSpeed *big.Int, _membershipHistoryLength *big.Int, _slashingMultiplierResetPeriod *big.Int, _maxGroupSize *big.Int, _commissionUpdateDelay *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "initialize", registryAddress, groupRequirementValue, groupRequirementDuration, validatorRequirementValue, validatorRequirementDuration, validatorScoreExponent, validatorScoreAdjustmentSpeed, _membershipHistoryLength, _slashingMultiplierResetPeriod, _maxGroupSize, _commissionUpdateDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x78d25456.
//
// Solidity: function initialize(address registryAddress, uint256 groupRequirementValue, uint256 groupRequirementDuration, uint256 validatorRequirementValue, uint256 validatorRequirementDuration, uint256 validatorScoreExponent, uint256 validatorScoreAdjustmentSpeed, uint256 _membershipHistoryLength, uint256 _slashingMultiplierResetPeriod, uint256 _maxGroupSize, uint256 _commissionUpdateDelay) returns()
func (_Validators *ValidatorsSession) Initialize(registryAddress common.Address, groupRequirementValue *big.Int, groupRequirementDuration *big.Int, validatorRequirementValue *big.Int, validatorRequirementDuration *big.Int, validatorScoreExponent *big.Int, validatorScoreAdjustmentSpeed *big.Int, _membershipHistoryLength *big.Int, _slashingMultiplierResetPeriod *big.Int, _maxGroupSize *big.Int, _commissionUpdateDelay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, registryAddress, groupRequirementValue, groupRequirementDuration, validatorRequirementValue, validatorRequirementDuration, validatorScoreExponent, validatorScoreAdjustmentSpeed, _membershipHistoryLength, _slashingMultiplierResetPeriod, _maxGroupSize, _commissionUpdateDelay)
}

// Initialize is a paid mutator transaction binding the contract method 0x78d25456.
//
// Solidity: function initialize(address registryAddress, uint256 groupRequirementValue, uint256 groupRequirementDuration, uint256 validatorRequirementValue, uint256 validatorRequirementDuration, uint256 validatorScoreExponent, uint256 validatorScoreAdjustmentSpeed, uint256 _membershipHistoryLength, uint256 _slashingMultiplierResetPeriod, uint256 _maxGroupSize, uint256 _commissionUpdateDelay) returns()
func (_Validators *ValidatorsTransactorSession) Initialize(registryAddress common.Address, groupRequirementValue *big.Int, groupRequirementDuration *big.Int, validatorRequirementValue *big.Int, validatorRequirementDuration *big.Int, validatorScoreExponent *big.Int, validatorScoreAdjustmentSpeed *big.Int, _membershipHistoryLength *big.Int, _slashingMultiplierResetPeriod *big.Int, _maxGroupSize *big.Int, _commissionUpdateDelay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, registryAddress, groupRequirementValue, groupRequirementDuration, validatorRequirementValue, validatorRequirementDuration, validatorScoreExponent, validatorScoreAdjustmentSpeed, _membershipHistoryLength, _slashingMultiplierResetPeriod, _maxGroupSize, _commissionUpdateDelay)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xea684f77.
//
// Solidity: function registerValidator(bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactor) RegisterValidator(opts *bind.TransactOpts, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "registerValidator", ecdsaPublicKey, blsPublicKey, blsPop)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xea684f77.
//
// Solidity: function registerValidator(bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsSession) RegisterValidator(ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidator(&_Validators.TransactOpts, ecdsaPublicKey, blsPublicKey, blsPop)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xea684f77.
//
// Solidity: function registerValidator(bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactorSession) RegisterValidator(ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidator(&_Validators.TransactOpts, ecdsaPublicKey, blsPublicKey, blsPop)
}

// RegisterValidatorGroup is a paid mutator transaction binding the contract method 0xee098310.
//
// Solidity: function registerValidatorGroup(uint256 commission) returns(bool)
func (_Validators *ValidatorsTransactor) RegisterValidatorGroup(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "registerValidatorGroup", commission)
}

// RegisterValidatorGroup is a paid mutator transaction binding the contract method 0xee098310.
//
// Solidity: function registerValidatorGroup(uint256 commission) returns(bool)
func (_Validators *ValidatorsSession) RegisterValidatorGroup(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidatorGroup(&_Validators.TransactOpts, commission)
}

// RegisterValidatorGroup is a paid mutator transaction binding the contract method 0xee098310.
//
// Solidity: function registerValidatorGroup(uint256 commission) returns(bool)
func (_Validators *ValidatorsTransactorSession) RegisterValidatorGroup(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.RegisterValidatorGroup(&_Validators.TransactOpts, commission)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) RemoveMember(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "removeMember", validator)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address validator) returns(bool)
func (_Validators *ValidatorsSession) RemoveMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveMember(&_Validators.TransactOpts, validator)
}

// RemoveMember is a paid mutator transaction binding the contract method 0x0b1ca49a.
//
// Solidity: function removeMember(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) RemoveMember(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveMember(&_Validators.TransactOpts, validator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validators.Contract.RenounceOwnership(&_Validators.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Validators *ValidatorsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Validators.Contract.RenounceOwnership(&_Validators.TransactOpts)
}

// ReorderMember is a paid mutator transaction binding the contract method 0x988dcd1f.
//
// Solidity: function reorderMember(address validator, address lesserMember, address greaterMember) returns(bool)
func (_Validators *ValidatorsTransactor) ReorderMember(opts *bind.TransactOpts, validator common.Address, lesserMember common.Address, greaterMember common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "reorderMember", validator, lesserMember, greaterMember)
}

// ReorderMember is a paid mutator transaction binding the contract method 0x988dcd1f.
//
// Solidity: function reorderMember(address validator, address lesserMember, address greaterMember) returns(bool)
func (_Validators *ValidatorsSession) ReorderMember(validator common.Address, lesserMember common.Address, greaterMember common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ReorderMember(&_Validators.TransactOpts, validator, lesserMember, greaterMember)
}

// ReorderMember is a paid mutator transaction binding the contract method 0x988dcd1f.
//
// Solidity: function reorderMember(address validator, address lesserMember, address greaterMember) returns(bool)
func (_Validators *ValidatorsTransactorSession) ReorderMember(validator common.Address, lesserMember common.Address, greaterMember common.Address) (*types.Transaction, error) {
	return _Validators.Contract.ReorderMember(&_Validators.TransactOpts, validator, lesserMember, greaterMember)
}

// ResetSlashingMultiplier is a paid mutator transaction binding the contract method 0xb8f93943.
//
// Solidity: function resetSlashingMultiplier() returns()
func (_Validators *ValidatorsTransactor) ResetSlashingMultiplier(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "resetSlashingMultiplier")
}

// ResetSlashingMultiplier is a paid mutator transaction binding the contract method 0xb8f93943.
//
// Solidity: function resetSlashingMultiplier() returns()
func (_Validators *ValidatorsSession) ResetSlashingMultiplier() (*types.Transaction, error) {
	return _Validators.Contract.ResetSlashingMultiplier(&_Validators.TransactOpts)
}

// ResetSlashingMultiplier is a paid mutator transaction binding the contract method 0xb8f93943.
//
// Solidity: function resetSlashingMultiplier() returns()
func (_Validators *ValidatorsTransactorSession) ResetSlashingMultiplier() (*types.Transaction, error) {
	return _Validators.Contract.ResetSlashingMultiplier(&_Validators.TransactOpts)
}

// SetCommissionUpdateDelay is a paid mutator transaction binding the contract method 0x6c620d90.
//
// Solidity: function setCommissionUpdateDelay(uint256 delay) returns()
func (_Validators *ValidatorsTransactor) SetCommissionUpdateDelay(opts *bind.TransactOpts, delay *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setCommissionUpdateDelay", delay)
}

// SetCommissionUpdateDelay is a paid mutator transaction binding the contract method 0x6c620d90.
//
// Solidity: function setCommissionUpdateDelay(uint256 delay) returns()
func (_Validators *ValidatorsSession) SetCommissionUpdateDelay(delay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetCommissionUpdateDelay(&_Validators.TransactOpts, delay)
}

// SetCommissionUpdateDelay is a paid mutator transaction binding the contract method 0x6c620d90.
//
// Solidity: function setCommissionUpdateDelay(uint256 delay) returns()
func (_Validators *ValidatorsTransactorSession) SetCommissionUpdateDelay(delay *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetCommissionUpdateDelay(&_Validators.TransactOpts, delay)
}

// SetGroupLockedGoldRequirements is a paid mutator transaction binding the contract method 0x5a61d15b.
//
// Solidity: function setGroupLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactor) SetGroupLockedGoldRequirements(opts *bind.TransactOpts, value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setGroupLockedGoldRequirements", value, duration)
}

// SetGroupLockedGoldRequirements is a paid mutator transaction binding the contract method 0x5a61d15b.
//
// Solidity: function setGroupLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsSession) SetGroupLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetGroupLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetGroupLockedGoldRequirements is a paid mutator transaction binding the contract method 0x5a61d15b.
//
// Solidity: function setGroupLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetGroupLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetGroupLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetMaxGroupSize is a paid mutator transaction binding the contract method 0xe1497ff7.
//
// Solidity: function setMaxGroupSize(uint256 size) returns(bool)
func (_Validators *ValidatorsTransactor) SetMaxGroupSize(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMaxGroupSize", size)
}

// SetMaxGroupSize is a paid mutator transaction binding the contract method 0xe1497ff7.
//
// Solidity: function setMaxGroupSize(uint256 size) returns(bool)
func (_Validators *ValidatorsSession) SetMaxGroupSize(size *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMaxGroupSize(&_Validators.TransactOpts, size)
}

// SetMaxGroupSize is a paid mutator transaction binding the contract method 0xe1497ff7.
//
// Solidity: function setMaxGroupSize(uint256 size) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetMaxGroupSize(size *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMaxGroupSize(&_Validators.TransactOpts, size)
}

// SetMembershipHistoryLength is a paid mutator transaction binding the contract method 0xeff2ea3f.
//
// Solidity: function setMembershipHistoryLength(uint256 length) returns(bool)
func (_Validators *ValidatorsTransactor) SetMembershipHistoryLength(opts *bind.TransactOpts, length *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setMembershipHistoryLength", length)
}

// SetMembershipHistoryLength is a paid mutator transaction binding the contract method 0xeff2ea3f.
//
// Solidity: function setMembershipHistoryLength(uint256 length) returns(bool)
func (_Validators *ValidatorsSession) SetMembershipHistoryLength(length *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMembershipHistoryLength(&_Validators.TransactOpts, length)
}

// SetMembershipHistoryLength is a paid mutator transaction binding the contract method 0xeff2ea3f.
//
// Solidity: function setMembershipHistoryLength(uint256 length) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetMembershipHistoryLength(length *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetMembershipHistoryLength(&_Validators.TransactOpts, length)
}

// SetNextCommissionUpdate is a paid mutator transaction binding the contract method 0x86d81a5a.
//
// Solidity: function setNextCommissionUpdate(uint256 commission) returns()
func (_Validators *ValidatorsTransactor) SetNextCommissionUpdate(opts *bind.TransactOpts, commission *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setNextCommissionUpdate", commission)
}

// SetNextCommissionUpdate is a paid mutator transaction binding the contract method 0x86d81a5a.
//
// Solidity: function setNextCommissionUpdate(uint256 commission) returns()
func (_Validators *ValidatorsSession) SetNextCommissionUpdate(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetNextCommissionUpdate(&_Validators.TransactOpts, commission)
}

// SetNextCommissionUpdate is a paid mutator transaction binding the contract method 0x86d81a5a.
//
// Solidity: function setNextCommissionUpdate(uint256 commission) returns()
func (_Validators *ValidatorsTransactorSession) SetNextCommissionUpdate(commission *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetNextCommissionUpdate(&_Validators.TransactOpts, commission)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Validators *ValidatorsTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Validators *ValidatorsSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.SetRegistry(&_Validators.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Validators *ValidatorsTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Validators.Contract.SetRegistry(&_Validators.TransactOpts, registryAddress)
}

// SetSlashingMultiplierResetPeriod is a paid mutator transaction binding the contract method 0x6ab951a0.
//
// Solidity: function setSlashingMultiplierResetPeriod(uint256 value) returns()
func (_Validators *ValidatorsTransactor) SetSlashingMultiplierResetPeriod(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setSlashingMultiplierResetPeriod", value)
}

// SetSlashingMultiplierResetPeriod is a paid mutator transaction binding the contract method 0x6ab951a0.
//
// Solidity: function setSlashingMultiplierResetPeriod(uint256 value) returns()
func (_Validators *ValidatorsSession) SetSlashingMultiplierResetPeriod(value *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetSlashingMultiplierResetPeriod(&_Validators.TransactOpts, value)
}

// SetSlashingMultiplierResetPeriod is a paid mutator transaction binding the contract method 0x6ab951a0.
//
// Solidity: function setSlashingMultiplierResetPeriod(uint256 value) returns()
func (_Validators *ValidatorsTransactorSession) SetSlashingMultiplierResetPeriod(value *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetSlashingMultiplierResetPeriod(&_Validators.TransactOpts, value)
}

// SetValidatorLockedGoldRequirements is a paid mutator transaction binding the contract method 0x76c0a9ed.
//
// Solidity: function setValidatorLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactor) SetValidatorLockedGoldRequirements(opts *bind.TransactOpts, value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setValidatorLockedGoldRequirements", value, duration)
}

// SetValidatorLockedGoldRequirements is a paid mutator transaction binding the contract method 0x76c0a9ed.
//
// Solidity: function setValidatorLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsSession) SetValidatorLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetValidatorLockedGoldRequirements is a paid mutator transaction binding the contract method 0x76c0a9ed.
//
// Solidity: function setValidatorLockedGoldRequirements(uint256 value, uint256 duration) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetValidatorLockedGoldRequirements(value *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorLockedGoldRequirements(&_Validators.TransactOpts, value, duration)
}

// SetValidatorScoreParameters is a paid mutator transaction binding the contract method 0xcb8f98e0.
//
// Solidity: function setValidatorScoreParameters(uint256 exponent, uint256 adjustmentSpeed) returns(bool)
func (_Validators *ValidatorsTransactor) SetValidatorScoreParameters(opts *bind.TransactOpts, exponent *big.Int, adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "setValidatorScoreParameters", exponent, adjustmentSpeed)
}

// SetValidatorScoreParameters is a paid mutator transaction binding the contract method 0xcb8f98e0.
//
// Solidity: function setValidatorScoreParameters(uint256 exponent, uint256 adjustmentSpeed) returns(bool)
func (_Validators *ValidatorsSession) SetValidatorScoreParameters(exponent *big.Int, adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorScoreParameters(&_Validators.TransactOpts, exponent, adjustmentSpeed)
}

// SetValidatorScoreParameters is a paid mutator transaction binding the contract method 0xcb8f98e0.
//
// Solidity: function setValidatorScoreParameters(uint256 exponent, uint256 adjustmentSpeed) returns(bool)
func (_Validators *ValidatorsTransactorSession) SetValidatorScoreParameters(exponent *big.Int, adjustmentSpeed *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.SetValidatorScoreParameters(&_Validators.TransactOpts, exponent, adjustmentSpeed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validators *ValidatorsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validators *ValidatorsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TransferOwnership(&_Validators.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Validators *ValidatorsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TransferOwnership(&_Validators.TransactOpts, newOwner)
}

// UpdateBlsPublicKey is a paid mutator transaction binding the contract method 0xbfdb7417.
//
// Solidity: function updateBlsPublicKey(bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactor) UpdateBlsPublicKey(opts *bind.TransactOpts, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateBlsPublicKey", blsPublicKey, blsPop)
}

// UpdateBlsPublicKey is a paid mutator transaction binding the contract method 0xbfdb7417.
//
// Solidity: function updateBlsPublicKey(bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsSession) UpdateBlsPublicKey(blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateBlsPublicKey(&_Validators.TransactOpts, blsPublicKey, blsPop)
}

// UpdateBlsPublicKey is a paid mutator transaction binding the contract method 0xbfdb7417.
//
// Solidity: function updateBlsPublicKey(bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactorSession) UpdateBlsPublicKey(blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateBlsPublicKey(&_Validators.TransactOpts, blsPublicKey, blsPop)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xe7f03766.
//
// Solidity: function updateCommission() returns()
func (_Validators *ValidatorsTransactor) UpdateCommission(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateCommission")
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xe7f03766.
//
// Solidity: function updateCommission() returns()
func (_Validators *ValidatorsSession) UpdateCommission() (*types.Transaction, error) {
	return _Validators.Contract.UpdateCommission(&_Validators.TransactOpts)
}

// UpdateCommission is a paid mutator transaction binding the contract method 0xe7f03766.
//
// Solidity: function updateCommission() returns()
func (_Validators *ValidatorsTransactorSession) UpdateCommission() (*types.Transaction, error) {
	return _Validators.Contract.UpdateCommission(&_Validators.TransactOpts)
}

// UpdateEcdsaPublicKey is a paid mutator transaction binding the contract method 0x4e06fd8a.
//
// Solidity: function updateEcdsaPublicKey(address account, address signer, bytes ecdsaPublicKey) returns(bool)
func (_Validators *ValidatorsTransactor) UpdateEcdsaPublicKey(opts *bind.TransactOpts, account common.Address, signer common.Address, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateEcdsaPublicKey", account, signer, ecdsaPublicKey)
}

// UpdateEcdsaPublicKey is a paid mutator transaction binding the contract method 0x4e06fd8a.
//
// Solidity: function updateEcdsaPublicKey(address account, address signer, bytes ecdsaPublicKey) returns(bool)
func (_Validators *ValidatorsSession) UpdateEcdsaPublicKey(account common.Address, signer common.Address, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateEcdsaPublicKey(&_Validators.TransactOpts, account, signer, ecdsaPublicKey)
}

// UpdateEcdsaPublicKey is a paid mutator transaction binding the contract method 0x4e06fd8a.
//
// Solidity: function updateEcdsaPublicKey(address account, address signer, bytes ecdsaPublicKey) returns(bool)
func (_Validators *ValidatorsTransactorSession) UpdateEcdsaPublicKey(account common.Address, signer common.Address, ecdsaPublicKey []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdateEcdsaPublicKey(&_Validators.TransactOpts, account, signer, ecdsaPublicKey)
}

// UpdatePublicKeys is a paid mutator transaction binding the contract method 0x713ea0f3.
//
// Solidity: function updatePublicKeys(address account, address signer, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactor) UpdatePublicKeys(opts *bind.TransactOpts, account common.Address, signer common.Address, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updatePublicKeys", account, signer, ecdsaPublicKey, blsPublicKey, blsPop)
}

// UpdatePublicKeys is a paid mutator transaction binding the contract method 0x713ea0f3.
//
// Solidity: function updatePublicKeys(address account, address signer, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsSession) UpdatePublicKeys(account common.Address, signer common.Address, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdatePublicKeys(&_Validators.TransactOpts, account, signer, ecdsaPublicKey, blsPublicKey, blsPop)
}

// UpdatePublicKeys is a paid mutator transaction binding the contract method 0x713ea0f3.
//
// Solidity: function updatePublicKeys(address account, address signer, bytes ecdsaPublicKey, bytes blsPublicKey, bytes blsPop) returns(bool)
func (_Validators *ValidatorsTransactorSession) UpdatePublicKeys(account common.Address, signer common.Address, ecdsaPublicKey []byte, blsPublicKey []byte, blsPop []byte) (*types.Transaction, error) {
	return _Validators.Contract.UpdatePublicKeys(&_Validators.TransactOpts, account, signer, ecdsaPublicKey, blsPublicKey, blsPop)
}

// UpdateValidatorScoreFromSigner is a paid mutator transaction binding the contract method 0xc0c6ad6f.
//
// Solidity: function updateValidatorScoreFromSigner(address signer, uint256 uptime) returns()
func (_Validators *ValidatorsTransactor) UpdateValidatorScoreFromSigner(opts *bind.TransactOpts, signer common.Address, uptime *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateValidatorScoreFromSigner", signer, uptime)
}

// UpdateValidatorScoreFromSigner is a paid mutator transaction binding the contract method 0xc0c6ad6f.
//
// Solidity: function updateValidatorScoreFromSigner(address signer, uint256 uptime) returns()
func (_Validators *ValidatorsSession) UpdateValidatorScoreFromSigner(signer common.Address, uptime *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateValidatorScoreFromSigner(&_Validators.TransactOpts, signer, uptime)
}

// UpdateValidatorScoreFromSigner is a paid mutator transaction binding the contract method 0xc0c6ad6f.
//
// Solidity: function updateValidatorScoreFromSigner(address signer, uint256 uptime) returns()
func (_Validators *ValidatorsTransactorSession) UpdateValidatorScoreFromSigner(signer common.Address, uptime *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateValidatorScoreFromSigner(&_Validators.TransactOpts, signer, uptime)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Validators *ValidatorsFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Validators.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "CommissionUpdateDelaySet":
		event, err = _Validators.ParseCommissionUpdateDelaySet(log)
	case "GroupLockedGoldRequirementsSet":
		event, err = _Validators.ParseGroupLockedGoldRequirementsSet(log)
	case "MaxGroupSizeSet":
		event, err = _Validators.ParseMaxGroupSizeSet(log)
	case "MembershipHistoryLengthSet":
		event, err = _Validators.ParseMembershipHistoryLengthSet(log)
	case "OwnershipTransferred":
		event, err = _Validators.ParseOwnershipTransferred(log)
	case "RegistrySet":
		event, err = _Validators.ParseRegistrySet(log)
	case "ValidatorAffiliated":
		event, err = _Validators.ParseValidatorAffiliated(log)
	case "ValidatorBlsPublicKeyUpdated":
		event, err = _Validators.ParseValidatorBlsPublicKeyUpdated(log)
	case "ValidatorDeaffiliated":
		event, err = _Validators.ParseValidatorDeaffiliated(log)
	case "ValidatorDeregistered":
		event, err = _Validators.ParseValidatorDeregistered(log)
	case "ValidatorEcdsaPublicKeyUpdated":
		event, err = _Validators.ParseValidatorEcdsaPublicKeyUpdated(log)
	case "ValidatorEpochPaymentDistributed":
		event, err = _Validators.ParseValidatorEpochPaymentDistributed(log)
	case "ValidatorGroupCommissionUpdateQueued":
		event, err = _Validators.ParseValidatorGroupCommissionUpdateQueued(log)
	case "ValidatorGroupCommissionUpdated":
		event, err = _Validators.ParseValidatorGroupCommissionUpdated(log)
	case "ValidatorGroupDeregistered":
		event, err = _Validators.ParseValidatorGroupDeregistered(log)
	case "ValidatorGroupMemberAdded":
		event, err = _Validators.ParseValidatorGroupMemberAdded(log)
	case "ValidatorGroupMemberRemoved":
		event, err = _Validators.ParseValidatorGroupMemberRemoved(log)
	case "ValidatorGroupMemberReordered":
		event, err = _Validators.ParseValidatorGroupMemberReordered(log)
	case "ValidatorGroupRegistered":
		event, err = _Validators.ParseValidatorGroupRegistered(log)
	case "ValidatorLockedGoldRequirementsSet":
		event, err = _Validators.ParseValidatorLockedGoldRequirementsSet(log)
	case "ValidatorRegistered":
		event, err = _Validators.ParseValidatorRegistered(log)
	case "ValidatorScoreParametersSet":
		event, err = _Validators.ParseValidatorScoreParametersSet(log)
	case "ValidatorScoreUpdated":
		event, err = _Validators.ParseValidatorScoreUpdated(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// ValidatorsCommissionUpdateDelaySetIterator is returned from FilterCommissionUpdateDelaySet and is used to iterate over the raw logs and unpacked data for CommissionUpdateDelaySet events raised by the Validators contract.
type ValidatorsCommissionUpdateDelaySetIterator struct {
	Event *ValidatorsCommissionUpdateDelaySet // Event containing the contract specifics and raw log

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
func (it *ValidatorsCommissionUpdateDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsCommissionUpdateDelaySet)
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
		it.Event = new(ValidatorsCommissionUpdateDelaySet)
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
func (it *ValidatorsCommissionUpdateDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsCommissionUpdateDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsCommissionUpdateDelaySet represents a CommissionUpdateDelaySet event raised by the Validators contract.
type ValidatorsCommissionUpdateDelaySet struct {
	Delay *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCommissionUpdateDelaySet is a free log retrieval operation binding the contract event 0xf2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b.
//
// Solidity: event CommissionUpdateDelaySet(uint256 delay)
func (_Validators *ValidatorsFilterer) FilterCommissionUpdateDelaySet(opts *bind.FilterOpts) (*ValidatorsCommissionUpdateDelaySetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "CommissionUpdateDelaySet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsCommissionUpdateDelaySetIterator{contract: _Validators.contract, event: "CommissionUpdateDelaySet", logs: logs, sub: sub}, nil
}

// WatchCommissionUpdateDelaySet is a free log subscription operation binding the contract event 0xf2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b.
//
// Solidity: event CommissionUpdateDelaySet(uint256 delay)
func (_Validators *ValidatorsFilterer) WatchCommissionUpdateDelaySet(opts *bind.WatchOpts, sink chan<- *ValidatorsCommissionUpdateDelaySet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "CommissionUpdateDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsCommissionUpdateDelaySet)
				if err := _Validators.contract.UnpackLog(event, "CommissionUpdateDelaySet", log); err != nil {
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

// ParseCommissionUpdateDelaySet is a log parse operation binding the contract event 0xf2da07d08fd8dc9c5dcf87ad6f540e306f884a47dd8de14b718a4d5395f1ca9b.
//
// Solidity: event CommissionUpdateDelaySet(uint256 delay)
func (_Validators *ValidatorsFilterer) ParseCommissionUpdateDelaySet(log types.Log) (*ValidatorsCommissionUpdateDelaySet, error) {
	event := new(ValidatorsCommissionUpdateDelaySet)
	if err := _Validators.contract.UnpackLog(event, "CommissionUpdateDelaySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsGroupLockedGoldRequirementsSetIterator is returned from FilterGroupLockedGoldRequirementsSet and is used to iterate over the raw logs and unpacked data for GroupLockedGoldRequirementsSet events raised by the Validators contract.
type ValidatorsGroupLockedGoldRequirementsSetIterator struct {
	Event *ValidatorsGroupLockedGoldRequirementsSet // Event containing the contract specifics and raw log

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
func (it *ValidatorsGroupLockedGoldRequirementsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsGroupLockedGoldRequirementsSet)
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
		it.Event = new(ValidatorsGroupLockedGoldRequirementsSet)
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
func (it *ValidatorsGroupLockedGoldRequirementsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsGroupLockedGoldRequirementsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsGroupLockedGoldRequirementsSet represents a GroupLockedGoldRequirementsSet event raised by the Validators contract.
type ValidatorsGroupLockedGoldRequirementsSet struct {
	Value    *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGroupLockedGoldRequirementsSet is a free log retrieval operation binding the contract event 0x999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4.
//
// Solidity: event GroupLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) FilterGroupLockedGoldRequirementsSet(opts *bind.FilterOpts) (*ValidatorsGroupLockedGoldRequirementsSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "GroupLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsGroupLockedGoldRequirementsSetIterator{contract: _Validators.contract, event: "GroupLockedGoldRequirementsSet", logs: logs, sub: sub}, nil
}

// WatchGroupLockedGoldRequirementsSet is a free log subscription operation binding the contract event 0x999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4.
//
// Solidity: event GroupLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) WatchGroupLockedGoldRequirementsSet(opts *bind.WatchOpts, sink chan<- *ValidatorsGroupLockedGoldRequirementsSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "GroupLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsGroupLockedGoldRequirementsSet)
				if err := _Validators.contract.UnpackLog(event, "GroupLockedGoldRequirementsSet", log); err != nil {
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

// ParseGroupLockedGoldRequirementsSet is a log parse operation binding the contract event 0x999f7ee1917e6d7ea08360edfe9250cda3eda859c38dcb71a92623665de64dd4.
//
// Solidity: event GroupLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) ParseGroupLockedGoldRequirementsSet(log types.Log) (*ValidatorsGroupLockedGoldRequirementsSet, error) {
	event := new(ValidatorsGroupLockedGoldRequirementsSet)
	if err := _Validators.contract.UnpackLog(event, "GroupLockedGoldRequirementsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsMaxGroupSizeSetIterator is returned from FilterMaxGroupSizeSet and is used to iterate over the raw logs and unpacked data for MaxGroupSizeSet events raised by the Validators contract.
type ValidatorsMaxGroupSizeSetIterator struct {
	Event *ValidatorsMaxGroupSizeSet // Event containing the contract specifics and raw log

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
func (it *ValidatorsMaxGroupSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsMaxGroupSizeSet)
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
		it.Event = new(ValidatorsMaxGroupSizeSet)
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
func (it *ValidatorsMaxGroupSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsMaxGroupSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsMaxGroupSizeSet represents a MaxGroupSizeSet event raised by the Validators contract.
type ValidatorsMaxGroupSizeSet struct {
	Size *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMaxGroupSizeSet is a free log retrieval operation binding the contract event 0x603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d5.
//
// Solidity: event MaxGroupSizeSet(uint256 size)
func (_Validators *ValidatorsFilterer) FilterMaxGroupSizeSet(opts *bind.FilterOpts) (*ValidatorsMaxGroupSizeSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "MaxGroupSizeSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsMaxGroupSizeSetIterator{contract: _Validators.contract, event: "MaxGroupSizeSet", logs: logs, sub: sub}, nil
}

// WatchMaxGroupSizeSet is a free log subscription operation binding the contract event 0x603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d5.
//
// Solidity: event MaxGroupSizeSet(uint256 size)
func (_Validators *ValidatorsFilterer) WatchMaxGroupSizeSet(opts *bind.WatchOpts, sink chan<- *ValidatorsMaxGroupSizeSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "MaxGroupSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsMaxGroupSizeSet)
				if err := _Validators.contract.UnpackLog(event, "MaxGroupSizeSet", log); err != nil {
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

// ParseMaxGroupSizeSet is a log parse operation binding the contract event 0x603fe12c33c253a23da1680aa453dc70c3a0ee07763569bd5f602406ebd4e5d5.
//
// Solidity: event MaxGroupSizeSet(uint256 size)
func (_Validators *ValidatorsFilterer) ParseMaxGroupSizeSet(log types.Log) (*ValidatorsMaxGroupSizeSet, error) {
	event := new(ValidatorsMaxGroupSizeSet)
	if err := _Validators.contract.UnpackLog(event, "MaxGroupSizeSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsMembershipHistoryLengthSetIterator is returned from FilterMembershipHistoryLengthSet and is used to iterate over the raw logs and unpacked data for MembershipHistoryLengthSet events raised by the Validators contract.
type ValidatorsMembershipHistoryLengthSetIterator struct {
	Event *ValidatorsMembershipHistoryLengthSet // Event containing the contract specifics and raw log

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
func (it *ValidatorsMembershipHistoryLengthSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsMembershipHistoryLengthSet)
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
		it.Event = new(ValidatorsMembershipHistoryLengthSet)
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
func (it *ValidatorsMembershipHistoryLengthSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsMembershipHistoryLengthSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsMembershipHistoryLengthSet represents a MembershipHistoryLengthSet event raised by the Validators contract.
type ValidatorsMembershipHistoryLengthSet struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMembershipHistoryLengthSet is a free log retrieval operation binding the contract event 0x1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a.
//
// Solidity: event MembershipHistoryLengthSet(uint256 length)
func (_Validators *ValidatorsFilterer) FilterMembershipHistoryLengthSet(opts *bind.FilterOpts) (*ValidatorsMembershipHistoryLengthSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "MembershipHistoryLengthSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsMembershipHistoryLengthSetIterator{contract: _Validators.contract, event: "MembershipHistoryLengthSet", logs: logs, sub: sub}, nil
}

// WatchMembershipHistoryLengthSet is a free log subscription operation binding the contract event 0x1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a.
//
// Solidity: event MembershipHistoryLengthSet(uint256 length)
func (_Validators *ValidatorsFilterer) WatchMembershipHistoryLengthSet(opts *bind.WatchOpts, sink chan<- *ValidatorsMembershipHistoryLengthSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "MembershipHistoryLengthSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsMembershipHistoryLengthSet)
				if err := _Validators.contract.UnpackLog(event, "MembershipHistoryLengthSet", log); err != nil {
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

// ParseMembershipHistoryLengthSet is a log parse operation binding the contract event 0x1c75c7fb3ee9d13d8394372d8c7cdf1702fa947faa03f6ccfa500f787b09b48a.
//
// Solidity: event MembershipHistoryLengthSet(uint256 length)
func (_Validators *ValidatorsFilterer) ParseMembershipHistoryLengthSet(log types.Log) (*ValidatorsMembershipHistoryLengthSet, error) {
	event := new(ValidatorsMembershipHistoryLengthSet)
	if err := _Validators.contract.UnpackLog(event, "MembershipHistoryLengthSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Validators contract.
type ValidatorsOwnershipTransferredIterator struct {
	Event *ValidatorsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ValidatorsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsOwnershipTransferred)
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
		it.Event = new(ValidatorsOwnershipTransferred)
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
func (it *ValidatorsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsOwnershipTransferred represents a OwnershipTransferred event raised by the Validators contract.
type ValidatorsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validators *ValidatorsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ValidatorsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsOwnershipTransferredIterator{contract: _Validators.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Validators *ValidatorsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ValidatorsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsOwnershipTransferred)
				if err := _Validators.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Validators *ValidatorsFilterer) ParseOwnershipTransferred(log types.Log) (*ValidatorsOwnershipTransferred, error) {
	event := new(ValidatorsOwnershipTransferred)
	if err := _Validators.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Validators contract.
type ValidatorsRegistrySetIterator struct {
	Event *ValidatorsRegistrySet // Event containing the contract specifics and raw log

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
func (it *ValidatorsRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsRegistrySet)
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
		it.Event = new(ValidatorsRegistrySet)
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
func (it *ValidatorsRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsRegistrySet represents a RegistrySet event raised by the Validators contract.
type ValidatorsRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Validators *ValidatorsFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*ValidatorsRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsRegistrySetIterator{contract: _Validators.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Validators *ValidatorsFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *ValidatorsRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsRegistrySet)
				if err := _Validators.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Validators *ValidatorsFilterer) ParseRegistrySet(log types.Log) (*ValidatorsRegistrySet, error) {
	event := new(ValidatorsRegistrySet)
	if err := _Validators.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorAffiliatedIterator is returned from FilterValidatorAffiliated and is used to iterate over the raw logs and unpacked data for ValidatorAffiliated events raised by the Validators contract.
type ValidatorsValidatorAffiliatedIterator struct {
	Event *ValidatorsValidatorAffiliated // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorAffiliatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorAffiliated)
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
		it.Event = new(ValidatorsValidatorAffiliated)
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
func (it *ValidatorsValidatorAffiliatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorAffiliatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorAffiliated represents a ValidatorAffiliated event raised by the Validators contract.
type ValidatorsValidatorAffiliated struct {
	Validator common.Address
	Group     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAffiliated is a free log retrieval operation binding the contract event 0x91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d.
//
// Solidity: event ValidatorAffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) FilterValidatorAffiliated(opts *bind.FilterOpts, validator []common.Address, group []common.Address) (*ValidatorsValidatorAffiliatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorAffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorAffiliatedIterator{contract: _Validators.contract, event: "ValidatorAffiliated", logs: logs, sub: sub}, nil
}

// WatchValidatorAffiliated is a free log subscription operation binding the contract event 0x91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d.
//
// Solidity: event ValidatorAffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) WatchValidatorAffiliated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorAffiliated, validator []common.Address, group []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorAffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorAffiliated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorAffiliated", log); err != nil {
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

// ParseValidatorAffiliated is a log parse operation binding the contract event 0x91ef92227057e201e406c3451698dd780fe7672ad74328591c88d281af31581d.
//
// Solidity: event ValidatorAffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) ParseValidatorAffiliated(log types.Log) (*ValidatorsValidatorAffiliated, error) {
	event := new(ValidatorsValidatorAffiliated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorAffiliated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorBlsPublicKeyUpdatedIterator is returned from FilterValidatorBlsPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for ValidatorBlsPublicKeyUpdated events raised by the Validators contract.
type ValidatorsValidatorBlsPublicKeyUpdatedIterator struct {
	Event *ValidatorsValidatorBlsPublicKeyUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorBlsPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorBlsPublicKeyUpdated)
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
		it.Event = new(ValidatorsValidatorBlsPublicKeyUpdated)
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
func (it *ValidatorsValidatorBlsPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorBlsPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorBlsPublicKeyUpdated represents a ValidatorBlsPublicKeyUpdated event raised by the Validators contract.
type ValidatorsValidatorBlsPublicKeyUpdated struct {
	Validator    common.Address
	BlsPublicKey []byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterValidatorBlsPublicKeyUpdated is a free log retrieval operation binding the contract event 0x36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee33.
//
// Solidity: event ValidatorBlsPublicKeyUpdated(address indexed validator, bytes blsPublicKey)
func (_Validators *ValidatorsFilterer) FilterValidatorBlsPublicKeyUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorBlsPublicKeyUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorBlsPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorBlsPublicKeyUpdatedIterator{contract: _Validators.contract, event: "ValidatorBlsPublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorBlsPublicKeyUpdated is a free log subscription operation binding the contract event 0x36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee33.
//
// Solidity: event ValidatorBlsPublicKeyUpdated(address indexed validator, bytes blsPublicKey)
func (_Validators *ValidatorsFilterer) WatchValidatorBlsPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorBlsPublicKeyUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorBlsPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorBlsPublicKeyUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorBlsPublicKeyUpdated", log); err != nil {
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

// ParseValidatorBlsPublicKeyUpdated is a log parse operation binding the contract event 0x36a1aabe506bbe8802233cbb9aad628e91269e77077c953f9db3e02d7092ee33.
//
// Solidity: event ValidatorBlsPublicKeyUpdated(address indexed validator, bytes blsPublicKey)
func (_Validators *ValidatorsFilterer) ParseValidatorBlsPublicKeyUpdated(log types.Log) (*ValidatorsValidatorBlsPublicKeyUpdated, error) {
	event := new(ValidatorsValidatorBlsPublicKeyUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorBlsPublicKeyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorDeaffiliatedIterator is returned from FilterValidatorDeaffiliated and is used to iterate over the raw logs and unpacked data for ValidatorDeaffiliated events raised by the Validators contract.
type ValidatorsValidatorDeaffiliatedIterator struct {
	Event *ValidatorsValidatorDeaffiliated // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorDeaffiliatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorDeaffiliated)
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
		it.Event = new(ValidatorsValidatorDeaffiliated)
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
func (it *ValidatorsValidatorDeaffiliatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorDeaffiliatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorDeaffiliated represents a ValidatorDeaffiliated event raised by the Validators contract.
type ValidatorsValidatorDeaffiliated struct {
	Validator common.Address
	Group     common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeaffiliated is a free log retrieval operation binding the contract event 0x71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c.
//
// Solidity: event ValidatorDeaffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) FilterValidatorDeaffiliated(opts *bind.FilterOpts, validator []common.Address, group []common.Address) (*ValidatorsValidatorDeaffiliatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorDeaffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorDeaffiliatedIterator{contract: _Validators.contract, event: "ValidatorDeaffiliated", logs: logs, sub: sub}, nil
}

// WatchValidatorDeaffiliated is a free log subscription operation binding the contract event 0x71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c.
//
// Solidity: event ValidatorDeaffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) WatchValidatorDeaffiliated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorDeaffiliated, validator []common.Address, group []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}
	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorDeaffiliated", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorDeaffiliated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorDeaffiliated", log); err != nil {
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

// ParseValidatorDeaffiliated is a log parse operation binding the contract event 0x71815121f0622b31a3e7270eb28acb9fd10825ff418c9a18591f617bb8a31a6c.
//
// Solidity: event ValidatorDeaffiliated(address indexed validator, address indexed group)
func (_Validators *ValidatorsFilterer) ParseValidatorDeaffiliated(log types.Log) (*ValidatorsValidatorDeaffiliated, error) {
	event := new(ValidatorsValidatorDeaffiliated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorDeaffiliated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorDeregisteredIterator is returned from FilterValidatorDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorDeregistered events raised by the Validators contract.
type ValidatorsValidatorDeregisteredIterator struct {
	Event *ValidatorsValidatorDeregistered // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorDeregistered)
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
		it.Event = new(ValidatorsValidatorDeregistered)
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
func (it *ValidatorsValidatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorDeregistered represents a ValidatorDeregistered event raised by the Validators contract.
type ValidatorsValidatorDeregistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeregistered is a free log retrieval operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorDeregistered(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorDeregisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorDeregistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorDeregisteredIterator{contract: _Validators.contract, event: "ValidatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorDeregistered is a free log subscription operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorDeregistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorDeregistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorDeregistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorDeregistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
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

// ParseValidatorDeregistered is a log parse operation binding the contract event 0x51407fafe7ef9bec39c65a12a4885a274190991bf1e9057fcc384fc77ff1a7f0.
//
// Solidity: event ValidatorDeregistered(address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorDeregistered(log types.Log) (*ValidatorsValidatorDeregistered, error) {
	event := new(ValidatorsValidatorDeregistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorDeregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorEcdsaPublicKeyUpdatedIterator is returned from FilterValidatorEcdsaPublicKeyUpdated and is used to iterate over the raw logs and unpacked data for ValidatorEcdsaPublicKeyUpdated events raised by the Validators contract.
type ValidatorsValidatorEcdsaPublicKeyUpdatedIterator struct {
	Event *ValidatorsValidatorEcdsaPublicKeyUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorEcdsaPublicKeyUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorEcdsaPublicKeyUpdated)
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
		it.Event = new(ValidatorsValidatorEcdsaPublicKeyUpdated)
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
func (it *ValidatorsValidatorEcdsaPublicKeyUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorEcdsaPublicKeyUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorEcdsaPublicKeyUpdated represents a ValidatorEcdsaPublicKeyUpdated event raised by the Validators contract.
type ValidatorsValidatorEcdsaPublicKeyUpdated struct {
	Validator      common.Address
	EcdsaPublicKey []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorEcdsaPublicKeyUpdated is a free log retrieval operation binding the contract event 0x213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8.
//
// Solidity: event ValidatorEcdsaPublicKeyUpdated(address indexed validator, bytes ecdsaPublicKey)
func (_Validators *ValidatorsFilterer) FilterValidatorEcdsaPublicKeyUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorEcdsaPublicKeyUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorEcdsaPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorEcdsaPublicKeyUpdatedIterator{contract: _Validators.contract, event: "ValidatorEcdsaPublicKeyUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorEcdsaPublicKeyUpdated is a free log subscription operation binding the contract event 0x213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8.
//
// Solidity: event ValidatorEcdsaPublicKeyUpdated(address indexed validator, bytes ecdsaPublicKey)
func (_Validators *ValidatorsFilterer) WatchValidatorEcdsaPublicKeyUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorEcdsaPublicKeyUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorEcdsaPublicKeyUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorEcdsaPublicKeyUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorEcdsaPublicKeyUpdated", log); err != nil {
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

// ParseValidatorEcdsaPublicKeyUpdated is a log parse operation binding the contract event 0x213377eec2c15b21fa7abcbb0cb87a67e893cdb94a2564aa4bb4d380869473c8.
//
// Solidity: event ValidatorEcdsaPublicKeyUpdated(address indexed validator, bytes ecdsaPublicKey)
func (_Validators *ValidatorsFilterer) ParseValidatorEcdsaPublicKeyUpdated(log types.Log) (*ValidatorsValidatorEcdsaPublicKeyUpdated, error) {
	event := new(ValidatorsValidatorEcdsaPublicKeyUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorEcdsaPublicKeyUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorEpochPaymentDistributedIterator is returned from FilterValidatorEpochPaymentDistributed and is used to iterate over the raw logs and unpacked data for ValidatorEpochPaymentDistributed events raised by the Validators contract.
type ValidatorsValidatorEpochPaymentDistributedIterator struct {
	Event *ValidatorsValidatorEpochPaymentDistributed // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorEpochPaymentDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorEpochPaymentDistributed)
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
		it.Event = new(ValidatorsValidatorEpochPaymentDistributed)
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
func (it *ValidatorsValidatorEpochPaymentDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorEpochPaymentDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorEpochPaymentDistributed represents a ValidatorEpochPaymentDistributed event raised by the Validators contract.
type ValidatorsValidatorEpochPaymentDistributed struct {
	Validator        common.Address
	ValidatorPayment *big.Int
	Group            common.Address
	GroupPayment     *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterValidatorEpochPaymentDistributed is a free log retrieval operation binding the contract event 0x6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975.
//
// Solidity: event ValidatorEpochPaymentDistributed(address indexed validator, uint256 validatorPayment, address indexed group, uint256 groupPayment)
func (_Validators *ValidatorsFilterer) FilterValidatorEpochPaymentDistributed(opts *bind.FilterOpts, validator []common.Address, group []common.Address) (*ValidatorsValidatorEpochPaymentDistributedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorEpochPaymentDistributed", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorEpochPaymentDistributedIterator{contract: _Validators.contract, event: "ValidatorEpochPaymentDistributed", logs: logs, sub: sub}, nil
}

// WatchValidatorEpochPaymentDistributed is a free log subscription operation binding the contract event 0x6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975.
//
// Solidity: event ValidatorEpochPaymentDistributed(address indexed validator, uint256 validatorPayment, address indexed group, uint256 groupPayment)
func (_Validators *ValidatorsFilterer) WatchValidatorEpochPaymentDistributed(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorEpochPaymentDistributed, validator []common.Address, group []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorEpochPaymentDistributed", validatorRule, groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorEpochPaymentDistributed)
				if err := _Validators.contract.UnpackLog(event, "ValidatorEpochPaymentDistributed", log); err != nil {
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

// ParseValidatorEpochPaymentDistributed is a log parse operation binding the contract event 0x6f5937add2ec38a0fa4959bccd86e3fcc2aafb706cd3e6c0565f87a7b36b9975.
//
// Solidity: event ValidatorEpochPaymentDistributed(address indexed validator, uint256 validatorPayment, address indexed group, uint256 groupPayment)
func (_Validators *ValidatorsFilterer) ParseValidatorEpochPaymentDistributed(log types.Log) (*ValidatorsValidatorEpochPaymentDistributed, error) {
	event := new(ValidatorsValidatorEpochPaymentDistributed)
	if err := _Validators.contract.UnpackLog(event, "ValidatorEpochPaymentDistributed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupCommissionUpdateQueuedIterator is returned from FilterValidatorGroupCommissionUpdateQueued and is used to iterate over the raw logs and unpacked data for ValidatorGroupCommissionUpdateQueued events raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdateQueuedIterator struct {
	Event *ValidatorsValidatorGroupCommissionUpdateQueued // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupCommissionUpdateQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupCommissionUpdateQueued)
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
		it.Event = new(ValidatorsValidatorGroupCommissionUpdateQueued)
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
func (it *ValidatorsValidatorGroupCommissionUpdateQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupCommissionUpdateQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupCommissionUpdateQueued represents a ValidatorGroupCommissionUpdateQueued event raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdateQueued struct {
	Group           common.Address
	Commission      *big.Int
	ActivationBlock *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupCommissionUpdateQueued is a free log retrieval operation binding the contract event 0x557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a152.
//
// Solidity: event ValidatorGroupCommissionUpdateQueued(address indexed group, uint256 commission, uint256 activationBlock)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupCommissionUpdateQueued(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupCommissionUpdateQueuedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupCommissionUpdateQueued", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupCommissionUpdateQueuedIterator{contract: _Validators.contract, event: "ValidatorGroupCommissionUpdateQueued", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupCommissionUpdateQueued is a free log subscription operation binding the contract event 0x557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a152.
//
// Solidity: event ValidatorGroupCommissionUpdateQueued(address indexed group, uint256 commission, uint256 activationBlock)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupCommissionUpdateQueued(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupCommissionUpdateQueued, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupCommissionUpdateQueued", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupCommissionUpdateQueued)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdateQueued", log); err != nil {
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

// ParseValidatorGroupCommissionUpdateQueued is a log parse operation binding the contract event 0x557d39a57520d9835859d4b7eda805a7f4115a59c3a374eeed488436fc62a152.
//
// Solidity: event ValidatorGroupCommissionUpdateQueued(address indexed group, uint256 commission, uint256 activationBlock)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupCommissionUpdateQueued(log types.Log) (*ValidatorsValidatorGroupCommissionUpdateQueued, error) {
	event := new(ValidatorsValidatorGroupCommissionUpdateQueued)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdateQueued", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupCommissionUpdatedIterator is returned from FilterValidatorGroupCommissionUpdated and is used to iterate over the raw logs and unpacked data for ValidatorGroupCommissionUpdated events raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdatedIterator struct {
	Event *ValidatorsValidatorGroupCommissionUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupCommissionUpdated)
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
		it.Event = new(ValidatorsValidatorGroupCommissionUpdated)
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
func (it *ValidatorsValidatorGroupCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupCommissionUpdated represents a ValidatorGroupCommissionUpdated event raised by the Validators contract.
type ValidatorsValidatorGroupCommissionUpdated struct {
	Group      common.Address
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupCommissionUpdated is a free log retrieval operation binding the contract event 0x815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c.
//
// Solidity: event ValidatorGroupCommissionUpdated(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupCommissionUpdated(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupCommissionUpdatedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupCommissionUpdated", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupCommissionUpdatedIterator{contract: _Validators.contract, event: "ValidatorGroupCommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupCommissionUpdated is a free log subscription operation binding the contract event 0x815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c.
//
// Solidity: event ValidatorGroupCommissionUpdated(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupCommissionUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupCommissionUpdated, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupCommissionUpdated", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupCommissionUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdated", log); err != nil {
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

// ParseValidatorGroupCommissionUpdated is a log parse operation binding the contract event 0x815d292dbc1a08dfb3103aabb6611233dd2393903e57bdf4c5b3db91198a826c.
//
// Solidity: event ValidatorGroupCommissionUpdated(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupCommissionUpdated(log types.Log) (*ValidatorsValidatorGroupCommissionUpdated, error) {
	event := new(ValidatorsValidatorGroupCommissionUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupCommissionUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupDeregisteredIterator is returned from FilterValidatorGroupDeregistered and is used to iterate over the raw logs and unpacked data for ValidatorGroupDeregistered events raised by the Validators contract.
type ValidatorsValidatorGroupDeregisteredIterator struct {
	Event *ValidatorsValidatorGroupDeregistered // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupDeregistered)
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
		it.Event = new(ValidatorsValidatorGroupDeregistered)
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
func (it *ValidatorsValidatorGroupDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupDeregistered represents a ValidatorGroupDeregistered event raised by the Validators contract.
type ValidatorsValidatorGroupDeregistered struct {
	Group common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupDeregistered is a free log retrieval operation binding the contract event 0xae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a8669.
//
// Solidity: event ValidatorGroupDeregistered(address indexed group)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupDeregistered(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupDeregisteredIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupDeregistered", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupDeregisteredIterator{contract: _Validators.contract, event: "ValidatorGroupDeregistered", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupDeregistered is a free log subscription operation binding the contract event 0xae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a8669.
//
// Solidity: event ValidatorGroupDeregistered(address indexed group)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupDeregistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupDeregistered, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupDeregistered", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupDeregistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupDeregistered", log); err != nil {
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

// ParseValidatorGroupDeregistered is a log parse operation binding the contract event 0xae7e034b0748a10a219b46074b20977a9170bf4027b156c797093773619a8669.
//
// Solidity: event ValidatorGroupDeregistered(address indexed group)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupDeregistered(log types.Log) (*ValidatorsValidatorGroupDeregistered, error) {
	event := new(ValidatorsValidatorGroupDeregistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupDeregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupMemberAddedIterator is returned from FilterValidatorGroupMemberAdded and is used to iterate over the raw logs and unpacked data for ValidatorGroupMemberAdded events raised by the Validators contract.
type ValidatorsValidatorGroupMemberAddedIterator struct {
	Event *ValidatorsValidatorGroupMemberAdded // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupMemberAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupMemberAdded)
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
		it.Event = new(ValidatorsValidatorGroupMemberAdded)
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
func (it *ValidatorsValidatorGroupMemberAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupMemberAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupMemberAdded represents a ValidatorGroupMemberAdded event raised by the Validators contract.
type ValidatorsValidatorGroupMemberAdded struct {
	Group     common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMemberAdded is a free log retrieval operation binding the contract event 0xbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad.
//
// Solidity: event ValidatorGroupMemberAdded(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupMemberAdded(opts *bind.FilterOpts, group []common.Address, validator []common.Address) (*ValidatorsValidatorGroupMemberAddedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupMemberAdded", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupMemberAddedIterator{contract: _Validators.contract, event: "ValidatorGroupMemberAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMemberAdded is a free log subscription operation binding the contract event 0xbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad.
//
// Solidity: event ValidatorGroupMemberAdded(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupMemberAdded(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupMemberAdded, group []common.Address, validator []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupMemberAdded", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupMemberAdded)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberAdded", log); err != nil {
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

// ParseValidatorGroupMemberAdded is a log parse operation binding the contract event 0xbdf7e616a6943f81e07a7984c9d4c00197dc2f481486ce4ffa6af52a113974ad.
//
// Solidity: event ValidatorGroupMemberAdded(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupMemberAdded(log types.Log) (*ValidatorsValidatorGroupMemberAdded, error) {
	event := new(ValidatorsValidatorGroupMemberAdded)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupMemberRemovedIterator is returned from FilterValidatorGroupMemberRemoved and is used to iterate over the raw logs and unpacked data for ValidatorGroupMemberRemoved events raised by the Validators contract.
type ValidatorsValidatorGroupMemberRemovedIterator struct {
	Event *ValidatorsValidatorGroupMemberRemoved // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupMemberRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupMemberRemoved)
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
		it.Event = new(ValidatorsValidatorGroupMemberRemoved)
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
func (it *ValidatorsValidatorGroupMemberRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupMemberRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupMemberRemoved represents a ValidatorGroupMemberRemoved event raised by the Validators contract.
type ValidatorsValidatorGroupMemberRemoved struct {
	Group     common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMemberRemoved is a free log retrieval operation binding the contract event 0xc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa57.
//
// Solidity: event ValidatorGroupMemberRemoved(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupMemberRemoved(opts *bind.FilterOpts, group []common.Address, validator []common.Address) (*ValidatorsValidatorGroupMemberRemovedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupMemberRemoved", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupMemberRemovedIterator{contract: _Validators.contract, event: "ValidatorGroupMemberRemoved", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMemberRemoved is a free log subscription operation binding the contract event 0xc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa57.
//
// Solidity: event ValidatorGroupMemberRemoved(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupMemberRemoved(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupMemberRemoved, group []common.Address, validator []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupMemberRemoved", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupMemberRemoved)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberRemoved", log); err != nil {
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

// ParseValidatorGroupMemberRemoved is a log parse operation binding the contract event 0xc7666a52a66ff601ff7c0d4d6efddc9ac20a34792f6aa003d1804c9d4d5baa57.
//
// Solidity: event ValidatorGroupMemberRemoved(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupMemberRemoved(log types.Log) (*ValidatorsValidatorGroupMemberRemoved, error) {
	event := new(ValidatorsValidatorGroupMemberRemoved)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupMemberReorderedIterator is returned from FilterValidatorGroupMemberReordered and is used to iterate over the raw logs and unpacked data for ValidatorGroupMemberReordered events raised by the Validators contract.
type ValidatorsValidatorGroupMemberReorderedIterator struct {
	Event *ValidatorsValidatorGroupMemberReordered // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupMemberReorderedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupMemberReordered)
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
		it.Event = new(ValidatorsValidatorGroupMemberReordered)
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
func (it *ValidatorsValidatorGroupMemberReorderedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupMemberReorderedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupMemberReordered represents a ValidatorGroupMemberReordered event raised by the Validators contract.
type ValidatorsValidatorGroupMemberReordered struct {
	Group     common.Address
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupMemberReordered is a free log retrieval operation binding the contract event 0x38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c6.
//
// Solidity: event ValidatorGroupMemberReordered(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupMemberReordered(opts *bind.FilterOpts, group []common.Address, validator []common.Address) (*ValidatorsValidatorGroupMemberReorderedIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupMemberReordered", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupMemberReorderedIterator{contract: _Validators.contract, event: "ValidatorGroupMemberReordered", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupMemberReordered is a free log subscription operation binding the contract event 0x38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c6.
//
// Solidity: event ValidatorGroupMemberReordered(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupMemberReordered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupMemberReordered, group []common.Address, validator []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupMemberReordered", groupRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupMemberReordered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberReordered", log); err != nil {
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

// ParseValidatorGroupMemberReordered is a log parse operation binding the contract event 0x38819cc49a343985b478d72f531a35b15384c398dd80fd191a14662170f895c6.
//
// Solidity: event ValidatorGroupMemberReordered(address indexed group, address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupMemberReordered(log types.Log) (*ValidatorsValidatorGroupMemberReordered, error) {
	event := new(ValidatorsValidatorGroupMemberReordered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupMemberReordered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorGroupRegisteredIterator is returned from FilterValidatorGroupRegistered and is used to iterate over the raw logs and unpacked data for ValidatorGroupRegistered events raised by the Validators contract.
type ValidatorsValidatorGroupRegisteredIterator struct {
	Event *ValidatorsValidatorGroupRegistered // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorGroupRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorGroupRegistered)
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
		it.Event = new(ValidatorsValidatorGroupRegistered)
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
func (it *ValidatorsValidatorGroupRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorGroupRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorGroupRegistered represents a ValidatorGroupRegistered event raised by the Validators contract.
type ValidatorsValidatorGroupRegistered struct {
	Group      common.Address
	Commission *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorGroupRegistered is a free log retrieval operation binding the contract event 0xbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc.
//
// Solidity: event ValidatorGroupRegistered(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) FilterValidatorGroupRegistered(opts *bind.FilterOpts, group []common.Address) (*ValidatorsValidatorGroupRegisteredIterator, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorGroupRegistered", groupRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorGroupRegisteredIterator{contract: _Validators.contract, event: "ValidatorGroupRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorGroupRegistered is a free log subscription operation binding the contract event 0xbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc.
//
// Solidity: event ValidatorGroupRegistered(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) WatchValidatorGroupRegistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorGroupRegistered, group []common.Address) (event.Subscription, error) {

	var groupRule []interface{}
	for _, groupItem := range group {
		groupRule = append(groupRule, groupItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorGroupRegistered", groupRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorGroupRegistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorGroupRegistered", log); err != nil {
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

// ParseValidatorGroupRegistered is a log parse operation binding the contract event 0xbf4b45570f1907a94775f8449817051a492a676918e38108bb762e991e6b58dc.
//
// Solidity: event ValidatorGroupRegistered(address indexed group, uint256 commission)
func (_Validators *ValidatorsFilterer) ParseValidatorGroupRegistered(log types.Log) (*ValidatorsValidatorGroupRegistered, error) {
	event := new(ValidatorsValidatorGroupRegistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorGroupRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorLockedGoldRequirementsSetIterator is returned from FilterValidatorLockedGoldRequirementsSet and is used to iterate over the raw logs and unpacked data for ValidatorLockedGoldRequirementsSet events raised by the Validators contract.
type ValidatorsValidatorLockedGoldRequirementsSetIterator struct {
	Event *ValidatorsValidatorLockedGoldRequirementsSet // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorLockedGoldRequirementsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorLockedGoldRequirementsSet)
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
		it.Event = new(ValidatorsValidatorLockedGoldRequirementsSet)
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
func (it *ValidatorsValidatorLockedGoldRequirementsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorLockedGoldRequirementsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorLockedGoldRequirementsSet represents a ValidatorLockedGoldRequirementsSet event raised by the Validators contract.
type ValidatorsValidatorLockedGoldRequirementsSet struct {
	Value    *big.Int
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValidatorLockedGoldRequirementsSet is a free log retrieval operation binding the contract event 0x62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389.
//
// Solidity: event ValidatorLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) FilterValidatorLockedGoldRequirementsSet(opts *bind.FilterOpts) (*ValidatorsValidatorLockedGoldRequirementsSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorLockedGoldRequirementsSetIterator{contract: _Validators.contract, event: "ValidatorLockedGoldRequirementsSet", logs: logs, sub: sub}, nil
}

// WatchValidatorLockedGoldRequirementsSet is a free log subscription operation binding the contract event 0x62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389.
//
// Solidity: event ValidatorLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) WatchValidatorLockedGoldRequirementsSet(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorLockedGoldRequirementsSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorLockedGoldRequirementsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorLockedGoldRequirementsSet)
				if err := _Validators.contract.UnpackLog(event, "ValidatorLockedGoldRequirementsSet", log); err != nil {
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

// ParseValidatorLockedGoldRequirementsSet is a log parse operation binding the contract event 0x62d947118dd4c1f5ece7f787a9cad4e1127d14d403b71133e95792b473bf8389.
//
// Solidity: event ValidatorLockedGoldRequirementsSet(uint256 value, uint256 duration)
func (_Validators *ValidatorsFilterer) ParseValidatorLockedGoldRequirementsSet(log types.Log) (*ValidatorsValidatorLockedGoldRequirementsSet, error) {
	event := new(ValidatorsValidatorLockedGoldRequirementsSet)
	if err := _Validators.contract.UnpackLog(event, "ValidatorLockedGoldRequirementsSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the Validators contract.
type ValidatorsValidatorRegisteredIterator struct {
	Event *ValidatorsValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorRegistered)
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
		it.Event = new(ValidatorsValidatorRegistered)
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
func (it *ValidatorsValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorRegistered represents a ValidatorRegistered event raised by the Validators contract.
type ValidatorsValidatorRegistered struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_Validators *ValidatorsFilterer) FilterValidatorRegistered(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorRegisteredIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorRegisteredIterator{contract: _Validators.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_Validators *ValidatorsFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorRegistered, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorRegistered", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorRegistered)
				if err := _Validators.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0xd09501348473474a20c772c79c653e1fd7e8b437e418fe235d277d2c88853251.
//
// Solidity: event ValidatorRegistered(address indexed validator)
func (_Validators *ValidatorsFilterer) ParseValidatorRegistered(log types.Log) (*ValidatorsValidatorRegistered, error) {
	event := new(ValidatorsValidatorRegistered)
	if err := _Validators.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorScoreParametersSetIterator is returned from FilterValidatorScoreParametersSet and is used to iterate over the raw logs and unpacked data for ValidatorScoreParametersSet events raised by the Validators contract.
type ValidatorsValidatorScoreParametersSetIterator struct {
	Event *ValidatorsValidatorScoreParametersSet // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorScoreParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorScoreParametersSet)
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
		it.Event = new(ValidatorsValidatorScoreParametersSet)
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
func (it *ValidatorsValidatorScoreParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorScoreParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorScoreParametersSet represents a ValidatorScoreParametersSet event raised by the Validators contract.
type ValidatorsValidatorScoreParametersSet struct {
	Exponent        *big.Int
	AdjustmentSpeed *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorScoreParametersSet is a free log retrieval operation binding the contract event 0x4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c536.
//
// Solidity: event ValidatorScoreParametersSet(uint256 exponent, uint256 adjustmentSpeed)
func (_Validators *ValidatorsFilterer) FilterValidatorScoreParametersSet(opts *bind.FilterOpts) (*ValidatorsValidatorScoreParametersSetIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorScoreParametersSet")
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorScoreParametersSetIterator{contract: _Validators.contract, event: "ValidatorScoreParametersSet", logs: logs, sub: sub}, nil
}

// WatchValidatorScoreParametersSet is a free log subscription operation binding the contract event 0x4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c536.
//
// Solidity: event ValidatorScoreParametersSet(uint256 exponent, uint256 adjustmentSpeed)
func (_Validators *ValidatorsFilterer) WatchValidatorScoreParametersSet(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorScoreParametersSet) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorScoreParametersSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorScoreParametersSet)
				if err := _Validators.contract.UnpackLog(event, "ValidatorScoreParametersSet", log); err != nil {
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

// ParseValidatorScoreParametersSet is a log parse operation binding the contract event 0x4b48724280029c2ea7a445c9cea30838525342e7a9ea9468f630b52e75d6c536.
//
// Solidity: event ValidatorScoreParametersSet(uint256 exponent, uint256 adjustmentSpeed)
func (_Validators *ValidatorsFilterer) ParseValidatorScoreParametersSet(log types.Log) (*ValidatorsValidatorScoreParametersSet, error) {
	event := new(ValidatorsValidatorScoreParametersSet)
	if err := _Validators.contract.UnpackLog(event, "ValidatorScoreParametersSet", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ValidatorsValidatorScoreUpdatedIterator is returned from FilterValidatorScoreUpdated and is used to iterate over the raw logs and unpacked data for ValidatorScoreUpdated events raised by the Validators contract.
type ValidatorsValidatorScoreUpdatedIterator struct {
	Event *ValidatorsValidatorScoreUpdated // Event containing the contract specifics and raw log

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
func (it *ValidatorsValidatorScoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsValidatorScoreUpdated)
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
		it.Event = new(ValidatorsValidatorScoreUpdated)
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
func (it *ValidatorsValidatorScoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsValidatorScoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsValidatorScoreUpdated represents a ValidatorScoreUpdated event raised by the Validators contract.
type ValidatorsValidatorScoreUpdated struct {
	Validator  common.Address
	Score      *big.Int
	EpochScore *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterValidatorScoreUpdated is a free log retrieval operation binding the contract event 0xedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9.
//
// Solidity: event ValidatorScoreUpdated(address indexed validator, uint256 score, uint256 epochScore)
func (_Validators *ValidatorsFilterer) FilterValidatorScoreUpdated(opts *bind.FilterOpts, validator []common.Address) (*ValidatorsValidatorScoreUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "ValidatorScoreUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsValidatorScoreUpdatedIterator{contract: _Validators.contract, event: "ValidatorScoreUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorScoreUpdated is a free log subscription operation binding the contract event 0xedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9.
//
// Solidity: event ValidatorScoreUpdated(address indexed validator, uint256 score, uint256 epochScore)
func (_Validators *ValidatorsFilterer) WatchValidatorScoreUpdated(opts *bind.WatchOpts, sink chan<- *ValidatorsValidatorScoreUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "ValidatorScoreUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsValidatorScoreUpdated)
				if err := _Validators.contract.UnpackLog(event, "ValidatorScoreUpdated", log); err != nil {
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

// ParseValidatorScoreUpdated is a log parse operation binding the contract event 0xedf9f87e50e10c533bf3ae7f5a7894ae66c23e6cbbe8773d7765d20ad6f995e9.
//
// Solidity: event ValidatorScoreUpdated(address indexed validator, uint256 score, uint256 epochScore)
func (_Validators *ValidatorsFilterer) ParseValidatorScoreUpdated(log types.Log) (*ValidatorsValidatorScoreUpdated, error) {
	event := new(ValidatorsValidatorScoreUpdated)
	if err := _Validators.contract.UnpackLog(event, "ValidatorScoreUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
