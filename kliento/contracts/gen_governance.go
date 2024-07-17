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

// GovernanceMetaData contains all meta data concerning the Governance contract.
var GovernanceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"test\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApproverSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"concurrentProposals\",\"type\":\"uint256\"}],\"name\":\"ConcurrentProposalsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"functionId\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"ConstitutionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"dequeueFrequency\",\"type\":\"uint256\"}],\"name\":\"DequeueFrequencySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionStageDuration\",\"type\":\"uint256\"}],\"name\":\"ExecutionStageDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"HotfixApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"HotfixExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"HotfixPrepared\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"}],\"name\":\"HotfixWhitelisted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDeposit\",\"type\":\"uint256\"}],\"name\":\"MinDepositSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baselineQuorumFactor\",\"type\":\"uint256\"}],\"name\":\"ParticipationBaselineQuorumFactorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baselineUpdateFactor\",\"type\":\"uint256\"}],\"name\":\"ParticipationBaselineUpdateFactorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"participationBaseline\",\"type\":\"uint256\"}],\"name\":\"ParticipationBaselineUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"participationFloor\",\"type\":\"uint256\"}],\"name\":\"ParticipationFloorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProposalDequeued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExpired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"transactionCount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"deposit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"revokedUpvotes\",\"type\":\"uint256\"}],\"name\":\"ProposalUpvoteRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"upvotes\",\"type\":\"uint256\"}],\"name\":\"ProposalUpvoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoteRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"name\":\"ProposalVoteRevokedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"}],\"name\":\"ProposalVoted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"name\":\"ProposalVotedV2\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"queueExpiry\",\"type\":\"uint256\"}],\"name\":\"QueueExpirySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"referendumStageDuration\",\"type\":\"uint256\"}],\"name\":\"ReferendumStageDurationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"RegistrySet\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[],\"name\":\"approver\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"blsKey\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"blsPop\",\"type\":\"bytes\"}],\"name\":\"checkProofOfPossession\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"concurrentProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dequeueFrequency\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dequeued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"emptyIndices\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"aNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"bDenominator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exponent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_decimals\",\"type\":\"uint256\"}],\"name\":\"fractionMulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getBlockNumberFromHeader\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getEpochNumberOfBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getEpochSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getParentSealBitmap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"getVerifiedSealBitmapFromHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"}],\"name\":\"hashHeader\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hotfixes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"preparedEpoch\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastDequeue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"minQuorumSize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minQuorumSizeInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numberValidatorsInCurrentSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"numberValidatorsInSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"queueExpiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"refundedDeposits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractIRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"}],\"name\":\"setRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stageDurations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"approval\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referendum\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"execution\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromCurrentSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"validatorSignerAddressFromSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getVersionNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"registryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_concurrentProposals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minDeposit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_queueExpiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_dequeueFrequency\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referendumStageDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"executionStageDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"participationBaseline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"participationFloor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baselineUpdateFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baselineQuorumFactor\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_approver\",\"type\":\"address\"}],\"name\":\"setApprover\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_concurrentProposals\",\"type\":\"uint256\"}],\"name\":\"setConcurrentProposals\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minDeposit\",\"type\":\"uint256\"}],\"name\":\"setMinDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_queueExpiry\",\"type\":\"uint256\"}],\"name\":\"setQueueExpiry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_dequeueFrequency\",\"type\":\"uint256\"}],\"name\":\"setDequeueFrequency\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"referendumStageDuration\",\"type\":\"uint256\"}],\"name\":\"setReferendumStageDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"executionStageDuration\",\"type\":\"uint256\"}],\"name\":\"setExecutionStageDuration\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"participationBaseline\",\"type\":\"uint256\"}],\"name\":\"setParticipationBaseline\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"participationFloor\",\"type\":\"uint256\"}],\"name\":\"setParticipationFloor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baselineUpdateFactor\",\"type\":\"uint256\"}],\"name\":\"setBaselineUpdateFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"baselineQuorumFactor\",\"type\":\"uint256\"}],\"name\":\"setBaselineQuorumFactor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionId\",\"type\":\"bytes4\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"setConstitution\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destinations\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"dataLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"string\",\"name\":\"descriptionUrl\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lesser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"greater\",\"type\":\"uint256\"}],\"name\":\"upvote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposalStage\",\"outputs\":[{\"internalType\":\"enumProposals.Stage\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lesser\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"greater\",\"type\":\"uint256\"}],\"name\":\"revokeUpvote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"enumProposals.VoteValue\",\"name\":\"value\",\"type\":\"uint8\"}],\"name\":\"vote\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"name\":\"votePartially\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"revokeVotes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"approveHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"whitelister\",\"type\":\"address\"}],\"name\":\"isHotfixWhitelistedBy\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"whitelistHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"prepareHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"address[]\",\"name\":\"destinations\",\"type\":\"address[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[]\",\"name\":\"dataLengths\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"executeHotfix\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isVoting\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getReferendumStageDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getExecutionStageDuration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getParticipationParameters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getProposalTransaction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isApproved\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoteTotals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVoteRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQueueLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getUpvotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getQueue\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getDequeue\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getUpvoteRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getMostRecentReferendumProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"hotfixWhitelistValidatorTally\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"isHotfixPassing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getHotfixRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"dequeueProposalsIfReady\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"dequeueProposalIfReady\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isProposalDequeued\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isQueued\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isProposalPassing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"isDequeuedProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isDequeuedProposalExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"isQueuedProposalExpired\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"destination\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionId\",\"type\":\"bytes4\"}],\"name\":\"getConstitution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAmountOfGoldUsedForVoting\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// GovernanceABI is the input ABI used to generate the binding from.
// Deprecated: Use GovernanceMetaData.ABI instead.
var GovernanceABI = GovernanceMetaData.ABI

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// ParseGovernanceABI parses the ABI
func ParseGovernanceABI() (*abi.ABI, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return &parsed, nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// Approver is a free data retrieval call binding the contract method 0x141a8dd8.
//
// Solidity: function approver() view returns(address)
func (_Governance *GovernanceCaller) Approver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "approver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Approver is a free data retrieval call binding the contract method 0x141a8dd8.
//
// Solidity: function approver() view returns(address)
func (_Governance *GovernanceSession) Approver() (common.Address, error) {
	return _Governance.Contract.Approver(&_Governance.CallOpts)
}

// Approver is a free data retrieval call binding the contract method 0x141a8dd8.
//
// Solidity: function approver() view returns(address)
func (_Governance *GovernanceCallerSession) Approver() (common.Address, error) {
	return _Governance.Contract.Approver(&_Governance.CallOpts)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Governance *GovernanceCaller) CheckProofOfPossession(opts *bind.CallOpts, sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "checkProofOfPossession", sender, blsKey, blsPop)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Governance *GovernanceSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Governance.Contract.CheckProofOfPossession(&_Governance.CallOpts, sender, blsKey, blsPop)
}

// CheckProofOfPossession is a free data retrieval call binding the contract method 0x23f0ab65.
//
// Solidity: function checkProofOfPossession(address sender, bytes blsKey, bytes blsPop) view returns(bool)
func (_Governance *GovernanceCallerSession) CheckProofOfPossession(sender common.Address, blsKey []byte, blsPop []byte) (bool, error) {
	return _Governance.Contract.CheckProofOfPossession(&_Governance.CallOpts, sender, blsKey, blsPop)
}

// ConcurrentProposals is a free data retrieval call binding the contract method 0x1201a0fb.
//
// Solidity: function concurrentProposals() view returns(uint256)
func (_Governance *GovernanceCaller) ConcurrentProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "concurrentProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConcurrentProposals is a free data retrieval call binding the contract method 0x1201a0fb.
//
// Solidity: function concurrentProposals() view returns(uint256)
func (_Governance *GovernanceSession) ConcurrentProposals() (*big.Int, error) {
	return _Governance.Contract.ConcurrentProposals(&_Governance.CallOpts)
}

// ConcurrentProposals is a free data retrieval call binding the contract method 0x1201a0fb.
//
// Solidity: function concurrentProposals() view returns(uint256)
func (_Governance *GovernanceCallerSession) ConcurrentProposals() (*big.Int, error) {
	return _Governance.Contract.ConcurrentProposals(&_Governance.CallOpts)
}

// DequeueFrequency is a free data retrieval call binding the contract method 0x77d26a2a.
//
// Solidity: function dequeueFrequency() view returns(uint256)
func (_Governance *GovernanceCaller) DequeueFrequency(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "dequeueFrequency")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DequeueFrequency is a free data retrieval call binding the contract method 0x77d26a2a.
//
// Solidity: function dequeueFrequency() view returns(uint256)
func (_Governance *GovernanceSession) DequeueFrequency() (*big.Int, error) {
	return _Governance.Contract.DequeueFrequency(&_Governance.CallOpts)
}

// DequeueFrequency is a free data retrieval call binding the contract method 0x77d26a2a.
//
// Solidity: function dequeueFrequency() view returns(uint256)
func (_Governance *GovernanceCallerSession) DequeueFrequency() (*big.Int, error) {
	return _Governance.Contract.DequeueFrequency(&_Governance.CallOpts)
}

// Dequeued is a free data retrieval call binding the contract method 0xadd004df.
//
// Solidity: function dequeued(uint256 ) view returns(uint256)
func (_Governance *GovernanceCaller) Dequeued(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "dequeued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dequeued is a free data retrieval call binding the contract method 0xadd004df.
//
// Solidity: function dequeued(uint256 ) view returns(uint256)
func (_Governance *GovernanceSession) Dequeued(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.Dequeued(&_Governance.CallOpts, arg0)
}

// Dequeued is a free data retrieval call binding the contract method 0xadd004df.
//
// Solidity: function dequeued(uint256 ) view returns(uint256)
func (_Governance *GovernanceCallerSession) Dequeued(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.Dequeued(&_Governance.CallOpts, arg0)
}

// EmptyIndices is a free data retrieval call binding the contract method 0xaa2feb83.
//
// Solidity: function emptyIndices(uint256 ) view returns(uint256)
func (_Governance *GovernanceCaller) EmptyIndices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "emptyIndices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EmptyIndices is a free data retrieval call binding the contract method 0xaa2feb83.
//
// Solidity: function emptyIndices(uint256 ) view returns(uint256)
func (_Governance *GovernanceSession) EmptyIndices(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.EmptyIndices(&_Governance.CallOpts, arg0)
}

// EmptyIndices is a free data retrieval call binding the contract method 0xaa2feb83.
//
// Solidity: function emptyIndices(uint256 ) view returns(uint256)
func (_Governance *GovernanceCallerSession) EmptyIndices(arg0 *big.Int) (*big.Int, error) {
	return _Governance.Contract.EmptyIndices(&_Governance.CallOpts, arg0)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Governance *GovernanceCaller) FractionMulExp(opts *bind.CallOpts, aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "fractionMulExp", aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)

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
func (_Governance *GovernanceSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Governance.Contract.FractionMulExp(&_Governance.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// FractionMulExp is a free data retrieval call binding the contract method 0xec683072.
//
// Solidity: function fractionMulExp(uint256 aNumerator, uint256 aDenominator, uint256 bNumerator, uint256 bDenominator, uint256 exponent, uint256 _decimals) view returns(uint256, uint256)
func (_Governance *GovernanceCallerSession) FractionMulExp(aNumerator *big.Int, aDenominator *big.Int, bNumerator *big.Int, bDenominator *big.Int, exponent *big.Int, _decimals *big.Int) (*big.Int, *big.Int, error) {
	return _Governance.Contract.FractionMulExp(&_Governance.CallOpts, aNumerator, aDenominator, bNumerator, bDenominator, exponent, _decimals)
}

// GetAmountOfGoldUsedForVoting is a free data retrieval call binding the contract method 0x66547163.
//
// Solidity: function getAmountOfGoldUsedForVoting(address account) view returns(uint256)
func (_Governance *GovernanceCaller) GetAmountOfGoldUsedForVoting(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getAmountOfGoldUsedForVoting", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOfGoldUsedForVoting is a free data retrieval call binding the contract method 0x66547163.
//
// Solidity: function getAmountOfGoldUsedForVoting(address account) view returns(uint256)
func (_Governance *GovernanceSession) GetAmountOfGoldUsedForVoting(account common.Address) (*big.Int, error) {
	return _Governance.Contract.GetAmountOfGoldUsedForVoting(&_Governance.CallOpts, account)
}

// GetAmountOfGoldUsedForVoting is a free data retrieval call binding the contract method 0x66547163.
//
// Solidity: function getAmountOfGoldUsedForVoting(address account) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetAmountOfGoldUsedForVoting(account common.Address) (*big.Int, error) {
	return _Governance.Contract.GetAmountOfGoldUsedForVoting(&_Governance.CallOpts, account)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Governance *GovernanceCaller) GetBlockNumberFromHeader(opts *bind.CallOpts, header []byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getBlockNumberFromHeader", header)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Governance *GovernanceSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Governance.Contract.GetBlockNumberFromHeader(&_Governance.CallOpts, header)
}

// GetBlockNumberFromHeader is a free data retrieval call binding the contract method 0x8a883626.
//
// Solidity: function getBlockNumberFromHeader(bytes header) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetBlockNumberFromHeader(header []byte) (*big.Int, error) {
	return _Governance.Contract.GetBlockNumberFromHeader(&_Governance.CallOpts, header)
}

// GetConstitution is a free data retrieval call binding the contract method 0x97b9eba6.
//
// Solidity: function getConstitution(address destination, bytes4 functionId) view returns(uint256)
func (_Governance *GovernanceCaller) GetConstitution(opts *bind.CallOpts, destination common.Address, functionId [4]byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getConstitution", destination, functionId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConstitution is a free data retrieval call binding the contract method 0x97b9eba6.
//
// Solidity: function getConstitution(address destination, bytes4 functionId) view returns(uint256)
func (_Governance *GovernanceSession) GetConstitution(destination common.Address, functionId [4]byte) (*big.Int, error) {
	return _Governance.Contract.GetConstitution(&_Governance.CallOpts, destination, functionId)
}

// GetConstitution is a free data retrieval call binding the contract method 0x97b9eba6.
//
// Solidity: function getConstitution(address destination, bytes4 functionId) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetConstitution(destination common.Address, functionId [4]byte) (*big.Int, error) {
	return _Governance.Contract.GetConstitution(&_Governance.CallOpts, destination, functionId)
}

// GetDequeue is a free data retrieval call binding the contract method 0x6de8a63b.
//
// Solidity: function getDequeue() view returns(uint256[])
func (_Governance *GovernanceCaller) GetDequeue(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getDequeue")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDequeue is a free data retrieval call binding the contract method 0x6de8a63b.
//
// Solidity: function getDequeue() view returns(uint256[])
func (_Governance *GovernanceSession) GetDequeue() ([]*big.Int, error) {
	return _Governance.Contract.GetDequeue(&_Governance.CallOpts)
}

// GetDequeue is a free data retrieval call binding the contract method 0x6de8a63b.
//
// Solidity: function getDequeue() view returns(uint256[])
func (_Governance *GovernanceCallerSession) GetDequeue() ([]*big.Int, error) {
	return _Governance.Contract.GetDequeue(&_Governance.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Governance *GovernanceCaller) GetEpochNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getEpochNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Governance *GovernanceSession) GetEpochNumber() (*big.Int, error) {
	return _Governance.Contract.GetEpochNumber(&_Governance.CallOpts)
}

// GetEpochNumber is a free data retrieval call binding the contract method 0x9a7b3be7.
//
// Solidity: function getEpochNumber() view returns(uint256)
func (_Governance *GovernanceCallerSession) GetEpochNumber() (*big.Int, error) {
	return _Governance.Contract.GetEpochNumber(&_Governance.CallOpts)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCaller) GetEpochNumberOfBlock(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getEpochNumberOfBlock", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetEpochNumberOfBlock(&_Governance.CallOpts, blockNumber)
}

// GetEpochNumberOfBlock is a free data retrieval call binding the contract method 0x3b1eb4bf.
//
// Solidity: function getEpochNumberOfBlock(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetEpochNumberOfBlock(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetEpochNumberOfBlock(&_Governance.CallOpts, blockNumber)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Governance *GovernanceCaller) GetEpochSize(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getEpochSize")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Governance *GovernanceSession) GetEpochSize() (*big.Int, error) {
	return _Governance.Contract.GetEpochSize(&_Governance.CallOpts)
}

// GetEpochSize is a free data retrieval call binding the contract method 0xdf4da461.
//
// Solidity: function getEpochSize() view returns(uint256)
func (_Governance *GovernanceCallerSession) GetEpochSize() (*big.Int, error) {
	return _Governance.Contract.GetEpochSize(&_Governance.CallOpts)
}

// GetExecutionStageDuration is a free data retrieval call binding the contract method 0x30a095d0.
//
// Solidity: function getExecutionStageDuration() view returns(uint256)
func (_Governance *GovernanceCaller) GetExecutionStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getExecutionStageDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExecutionStageDuration is a free data retrieval call binding the contract method 0x30a095d0.
//
// Solidity: function getExecutionStageDuration() view returns(uint256)
func (_Governance *GovernanceSession) GetExecutionStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetExecutionStageDuration(&_Governance.CallOpts)
}

// GetExecutionStageDuration is a free data retrieval call binding the contract method 0x30a095d0.
//
// Solidity: function getExecutionStageDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) GetExecutionStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetExecutionStageDuration(&_Governance.CallOpts)
}

// GetHotfixRecord is a free data retrieval call binding the contract method 0x0e0b78ae.
//
// Solidity: function getHotfixRecord(bytes32 hash) view returns(bool, bool, uint256)
func (_Governance *GovernanceCaller) GetHotfixRecord(opts *bind.CallOpts, hash [32]byte) (bool, bool, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getHotfixRecord", hash)

	if err != nil {
		return *new(bool), *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHotfixRecord is a free data retrieval call binding the contract method 0x0e0b78ae.
//
// Solidity: function getHotfixRecord(bytes32 hash) view returns(bool, bool, uint256)
func (_Governance *GovernanceSession) GetHotfixRecord(hash [32]byte) (bool, bool, *big.Int, error) {
	return _Governance.Contract.GetHotfixRecord(&_Governance.CallOpts, hash)
}

// GetHotfixRecord is a free data retrieval call binding the contract method 0x0e0b78ae.
//
// Solidity: function getHotfixRecord(bytes32 hash) view returns(bool, bool, uint256)
func (_Governance *GovernanceCallerSession) GetHotfixRecord(hash [32]byte) (bool, bool, *big.Int, error) {
	return _Governance.Contract.GetHotfixRecord(&_Governance.CallOpts, hash)
}

// GetMostRecentReferendumProposal is a free data retrieval call binding the contract method 0x283aaefb.
//
// Solidity: function getMostRecentReferendumProposal(address account) view returns(uint256)
func (_Governance *GovernanceCaller) GetMostRecentReferendumProposal(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getMostRecentReferendumProposal", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMostRecentReferendumProposal is a free data retrieval call binding the contract method 0x283aaefb.
//
// Solidity: function getMostRecentReferendumProposal(address account) view returns(uint256)
func (_Governance *GovernanceSession) GetMostRecentReferendumProposal(account common.Address) (*big.Int, error) {
	return _Governance.Contract.GetMostRecentReferendumProposal(&_Governance.CallOpts, account)
}

// GetMostRecentReferendumProposal is a free data retrieval call binding the contract method 0x283aaefb.
//
// Solidity: function getMostRecentReferendumProposal(address account) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetMostRecentReferendumProposal(account common.Address) (*big.Int, error) {
	return _Governance.Contract.GetMostRecentReferendumProposal(&_Governance.CallOpts, account)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Governance *GovernanceCaller) GetParentSealBitmap(opts *bind.CallOpts, blockNumber *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getParentSealBitmap", blockNumber)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Governance *GovernanceSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Governance.Contract.GetParentSealBitmap(&_Governance.CallOpts, blockNumber)
}

// GetParentSealBitmap is a free data retrieval call binding the contract method 0xfae8db0a.
//
// Solidity: function getParentSealBitmap(uint256 blockNumber) view returns(bytes32)
func (_Governance *GovernanceCallerSession) GetParentSealBitmap(blockNumber *big.Int) ([32]byte, error) {
	return _Governance.Contract.GetParentSealBitmap(&_Governance.CallOpts, blockNumber)
}

// GetParticipationParameters is a free data retrieval call binding the contract method 0xc805956d.
//
// Solidity: function getParticipationParameters() view returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCaller) GetParticipationParameters(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getParticipationParameters")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetParticipationParameters is a free data retrieval call binding the contract method 0xc805956d.
//
// Solidity: function getParticipationParameters() view returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceSession) GetParticipationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetParticipationParameters(&_Governance.CallOpts)
}

// GetParticipationParameters is a free data retrieval call binding the contract method 0xc805956d.
//
// Solidity: function getParticipationParameters() view returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCallerSession) GetParticipationParameters() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetParticipationParameters(&_Governance.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns(address, uint256, uint256, uint256, string, uint256, bool)
func (_Governance *GovernanceCaller) GetProposal(opts *bind.CallOpts, proposalId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, string, *big.Int, bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposal", proposalId)

	if err != nil {
		return *new(common.Address), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(string), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns(address, uint256, uint256, uint256, string, uint256, bool)
func (_Governance *GovernanceSession) GetProposal(proposalId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, string, *big.Int, bool, error) {
	return _Governance.Contract.GetProposal(&_Governance.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns(address, uint256, uint256, uint256, string, uint256, bool)
func (_Governance *GovernanceCallerSession) GetProposal(proposalId *big.Int) (common.Address, *big.Int, *big.Int, *big.Int, string, *big.Int, bool, error) {
	return _Governance.Contract.GetProposal(&_Governance.CallOpts, proposalId)
}

// GetProposalStage is a free data retrieval call binding the contract method 0x582ae53b.
//
// Solidity: function getProposalStage(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceCaller) GetProposalStage(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposalStage", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetProposalStage is a free data retrieval call binding the contract method 0x582ae53b.
//
// Solidity: function getProposalStage(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceSession) GetProposalStage(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.GetProposalStage(&_Governance.CallOpts, proposalId)
}

// GetProposalStage is a free data retrieval call binding the contract method 0x582ae53b.
//
// Solidity: function getProposalStage(uint256 proposalId) view returns(uint8)
func (_Governance *GovernanceCallerSession) GetProposalStage(proposalId *big.Int) (uint8, error) {
	return _Governance.Contract.GetProposalStage(&_Governance.CallOpts, proposalId)
}

// GetProposalTransaction is a free data retrieval call binding the contract method 0xd704f0c5.
//
// Solidity: function getProposalTransaction(uint256 proposalId, uint256 index) view returns(uint256, address, bytes)
func (_Governance *GovernanceCaller) GetProposalTransaction(opts *bind.CallOpts, proposalId *big.Int, index *big.Int) (*big.Int, common.Address, []byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getProposalTransaction", proposalId, index)

	if err != nil {
		return *new(*big.Int), *new(common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// GetProposalTransaction is a free data retrieval call binding the contract method 0xd704f0c5.
//
// Solidity: function getProposalTransaction(uint256 proposalId, uint256 index) view returns(uint256, address, bytes)
func (_Governance *GovernanceSession) GetProposalTransaction(proposalId *big.Int, index *big.Int) (*big.Int, common.Address, []byte, error) {
	return _Governance.Contract.GetProposalTransaction(&_Governance.CallOpts, proposalId, index)
}

// GetProposalTransaction is a free data retrieval call binding the contract method 0xd704f0c5.
//
// Solidity: function getProposalTransaction(uint256 proposalId, uint256 index) view returns(uint256, address, bytes)
func (_Governance *GovernanceCallerSession) GetProposalTransaction(proposalId *big.Int, index *big.Int) (*big.Int, common.Address, []byte, error) {
	return _Governance.Contract.GetProposalTransaction(&_Governance.CallOpts, proposalId, index)
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() view returns(uint256[], uint256[])
func (_Governance *GovernanceCaller) GetQueue(opts *bind.CallOpts) ([]*big.Int, []*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getQueue")

	if err != nil {
		return *new([]*big.Int), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() view returns(uint256[], uint256[])
func (_Governance *GovernanceSession) GetQueue() ([]*big.Int, []*big.Int, error) {
	return _Governance.Contract.GetQueue(&_Governance.CallOpts)
}

// GetQueue is a free data retrieval call binding the contract method 0x01fce27e.
//
// Solidity: function getQueue() view returns(uint256[], uint256[])
func (_Governance *GovernanceCallerSession) GetQueue() ([]*big.Int, []*big.Int, error) {
	return _Governance.Contract.GetQueue(&_Governance.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint256)
func (_Governance *GovernanceCaller) GetQueueLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getQueueLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint256)
func (_Governance *GovernanceSession) GetQueueLength() (*big.Int, error) {
	return _Governance.Contract.GetQueueLength(&_Governance.CallOpts)
}

// GetQueueLength is a free data retrieval call binding the contract method 0xb8f77005.
//
// Solidity: function getQueueLength() view returns(uint256)
func (_Governance *GovernanceCallerSession) GetQueueLength() (*big.Int, error) {
	return _Governance.Contract.GetQueueLength(&_Governance.CallOpts)
}

// GetReferendumStageDuration is a free data retrieval call binding the contract method 0xad78c109.
//
// Solidity: function getReferendumStageDuration() view returns(uint256)
func (_Governance *GovernanceCaller) GetReferendumStageDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getReferendumStageDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReferendumStageDuration is a free data retrieval call binding the contract method 0xad78c109.
//
// Solidity: function getReferendumStageDuration() view returns(uint256)
func (_Governance *GovernanceSession) GetReferendumStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetReferendumStageDuration(&_Governance.CallOpts)
}

// GetReferendumStageDuration is a free data retrieval call binding the contract method 0xad78c109.
//
// Solidity: function getReferendumStageDuration() view returns(uint256)
func (_Governance *GovernanceCallerSession) GetReferendumStageDuration() (*big.Int, error) {
	return _Governance.Contract.GetReferendumStageDuration(&_Governance.CallOpts)
}

// GetUpvoteRecord is a free data retrieval call binding the contract method 0xcd845a76.
//
// Solidity: function getUpvoteRecord(address account) view returns(uint256, uint256)
func (_Governance *GovernanceCaller) GetUpvoteRecord(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getUpvoteRecord", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUpvoteRecord is a free data retrieval call binding the contract method 0xcd845a76.
//
// Solidity: function getUpvoteRecord(address account) view returns(uint256, uint256)
func (_Governance *GovernanceSession) GetUpvoteRecord(account common.Address) (*big.Int, *big.Int, error) {
	return _Governance.Contract.GetUpvoteRecord(&_Governance.CallOpts, account)
}

// GetUpvoteRecord is a free data retrieval call binding the contract method 0xcd845a76.
//
// Solidity: function getUpvoteRecord(address account) view returns(uint256, uint256)
func (_Governance *GovernanceCallerSession) GetUpvoteRecord(account common.Address) (*big.Int, *big.Int, error) {
	return _Governance.Contract.GetUpvoteRecord(&_Governance.CallOpts, account)
}

// GetUpvotes is a free data retrieval call binding the contract method 0x98f42702.
//
// Solidity: function getUpvotes(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCaller) GetUpvotes(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getUpvotes", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUpvotes is a free data retrieval call binding the contract method 0x98f42702.
//
// Solidity: function getUpvotes(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceSession) GetUpvotes(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetUpvotes(&_Governance.CallOpts, proposalId)
}

// GetUpvotes is a free data retrieval call binding the contract method 0x98f42702.
//
// Solidity: function getUpvotes(uint256 proposalId) view returns(uint256)
func (_Governance *GovernanceCallerSession) GetUpvotes(proposalId *big.Int) (*big.Int, error) {
	return _Governance.Contract.GetUpvotes(&_Governance.CallOpts, proposalId)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Governance *GovernanceCaller) GetVerifiedSealBitmapFromHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVerifiedSealBitmapFromHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Governance *GovernanceSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.GetVerifiedSealBitmapFromHeader(&_Governance.CallOpts, header)
}

// GetVerifiedSealBitmapFromHeader is a free data retrieval call binding the contract method 0x4b2c2f44.
//
// Solidity: function getVerifiedSealBitmapFromHeader(bytes header) view returns(bytes32)
func (_Governance *GovernanceCallerSession) GetVerifiedSealBitmapFromHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.GetVerifiedSealBitmapFromHeader(&_Governance.CallOpts, header)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCaller) GetVersionNumber(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVersionNumber")

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
func (_Governance *GovernanceSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVersionNumber(&_Governance.CallOpts)
}

// GetVersionNumber is a free data retrieval call binding the contract method 0x54255be0.
//
// Solidity: function getVersionNumber() pure returns(uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCallerSession) GetVersionNumber() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVersionNumber(&_Governance.CallOpts)
}

// GetVoteRecord is a free data retrieval call binding the contract method 0x5f115a85.
//
// Solidity: function getVoteRecord(address account, uint256 index) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCaller) GetVoteRecord(opts *bind.CallOpts, account common.Address, index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVoteRecord", account, index)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetVoteRecord is a free data retrieval call binding the contract method 0x5f115a85.
//
// Solidity: function getVoteRecord(address account, uint256 index) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_Governance *GovernanceSession) GetVoteRecord(account common.Address, index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteRecord(&_Governance.CallOpts, account, index)
}

// GetVoteRecord is a free data retrieval call binding the contract method 0x5f115a85.
//
// Solidity: function getVoteRecord(address account, uint256 index) view returns(uint256, uint256, uint256, uint256, uint256, uint256)
func (_Governance *GovernanceCallerSession) GetVoteRecord(account common.Address, index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteRecord(&_Governance.CallOpts, account, index)
}

// GetVoteTotals is a free data retrieval call binding the contract method 0xe41db455.
//
// Solidity: function getVoteTotals(uint256 proposalId) view returns(uint256, uint256, uint256)
func (_Governance *GovernanceCaller) GetVoteTotals(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "getVoteTotals", proposalId)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetVoteTotals is a free data retrieval call binding the contract method 0xe41db455.
//
// Solidity: function getVoteTotals(uint256 proposalId) view returns(uint256, uint256, uint256)
func (_Governance *GovernanceSession) GetVoteTotals(proposalId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteTotals(&_Governance.CallOpts, proposalId)
}

// GetVoteTotals is a free data retrieval call binding the contract method 0xe41db455.
//
// Solidity: function getVoteTotals(uint256 proposalId) view returns(uint256, uint256, uint256)
func (_Governance *GovernanceCallerSession) GetVoteTotals(proposalId *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Governance.Contract.GetVoteTotals(&_Governance.CallOpts, proposalId)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Governance *GovernanceCaller) HashHeader(opts *bind.CallOpts, header []byte) ([32]byte, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "hashHeader", header)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Governance *GovernanceSession) HashHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.HashHeader(&_Governance.CallOpts, header)
}

// HashHeader is a free data retrieval call binding the contract method 0x67960e91.
//
// Solidity: function hashHeader(bytes header) view returns(bytes32)
func (_Governance *GovernanceCallerSession) HashHeader(header []byte) ([32]byte, error) {
	return _Governance.Contract.HashHeader(&_Governance.CallOpts, header)
}

// HotfixWhitelistValidatorTally is a free data retrieval call binding the contract method 0x81d4728d.
//
// Solidity: function hotfixWhitelistValidatorTally(bytes32 hash) view returns(uint256)
func (_Governance *GovernanceCaller) HotfixWhitelistValidatorTally(opts *bind.CallOpts, hash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "hotfixWhitelistValidatorTally", hash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HotfixWhitelistValidatorTally is a free data retrieval call binding the contract method 0x81d4728d.
//
// Solidity: function hotfixWhitelistValidatorTally(bytes32 hash) view returns(uint256)
func (_Governance *GovernanceSession) HotfixWhitelistValidatorTally(hash [32]byte) (*big.Int, error) {
	return _Governance.Contract.HotfixWhitelistValidatorTally(&_Governance.CallOpts, hash)
}

// HotfixWhitelistValidatorTally is a free data retrieval call binding the contract method 0x81d4728d.
//
// Solidity: function hotfixWhitelistValidatorTally(bytes32 hash) view returns(uint256)
func (_Governance *GovernanceCallerSession) HotfixWhitelistValidatorTally(hash [32]byte) (*big.Int, error) {
	return _Governance.Contract.HotfixWhitelistValidatorTally(&_Governance.CallOpts, hash)
}

// Hotfixes is a free data retrieval call binding the contract method 0x45a78499.
//
// Solidity: function hotfixes(bytes32 ) view returns(bool executed, bool approved, uint256 preparedEpoch)
func (_Governance *GovernanceCaller) Hotfixes(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Executed      bool
	Approved      bool
	PreparedEpoch *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "hotfixes", arg0)

	outstruct := new(struct {
		Executed      bool
		Approved      bool
		PreparedEpoch *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Executed = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Approved = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.PreparedEpoch = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Hotfixes is a free data retrieval call binding the contract method 0x45a78499.
//
// Solidity: function hotfixes(bytes32 ) view returns(bool executed, bool approved, uint256 preparedEpoch)
func (_Governance *GovernanceSession) Hotfixes(arg0 [32]byte) (struct {
	Executed      bool
	Approved      bool
	PreparedEpoch *big.Int
}, error) {
	return _Governance.Contract.Hotfixes(&_Governance.CallOpts, arg0)
}

// Hotfixes is a free data retrieval call binding the contract method 0x45a78499.
//
// Solidity: function hotfixes(bytes32 ) view returns(bool executed, bool approved, uint256 preparedEpoch)
func (_Governance *GovernanceCallerSession) Hotfixes(arg0 [32]byte) (struct {
	Executed      bool
	Approved      bool
	PreparedEpoch *big.Int
}, error) {
	return _Governance.Contract.Hotfixes(&_Governance.CallOpts, arg0)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Governance *GovernanceCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Governance *GovernanceSession) Initialized() (bool, error) {
	return _Governance.Contract.Initialized(&_Governance.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Governance *GovernanceCallerSession) Initialized() (bool, error) {
	return _Governance.Contract.Initialized(&_Governance.CallOpts)
}

// IsApproved is a free data retrieval call binding the contract method 0x7910867b.
//
// Solidity: function isApproved(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCaller) IsApproved(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isApproved", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApproved is a free data retrieval call binding the contract method 0x7910867b.
//
// Solidity: function isApproved(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceSession) IsApproved(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsApproved(&_Governance.CallOpts, proposalId)
}

// IsApproved is a free data retrieval call binding the contract method 0x7910867b.
//
// Solidity: function isApproved(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCallerSession) IsApproved(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsApproved(&_Governance.CallOpts, proposalId)
}

// IsDequeuedProposal is a free data retrieval call binding the contract method 0x152b4834.
//
// Solidity: function isDequeuedProposal(uint256 proposalId, uint256 index) view returns(bool)
func (_Governance *GovernanceCaller) IsDequeuedProposal(opts *bind.CallOpts, proposalId *big.Int, index *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isDequeuedProposal", proposalId, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDequeuedProposal is a free data retrieval call binding the contract method 0x152b4834.
//
// Solidity: function isDequeuedProposal(uint256 proposalId, uint256 index) view returns(bool)
func (_Governance *GovernanceSession) IsDequeuedProposal(proposalId *big.Int, index *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposal(&_Governance.CallOpts, proposalId, index)
}

// IsDequeuedProposal is a free data retrieval call binding the contract method 0x152b4834.
//
// Solidity: function isDequeuedProposal(uint256 proposalId, uint256 index) view returns(bool)
func (_Governance *GovernanceCallerSession) IsDequeuedProposal(proposalId *big.Int, index *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposal(&_Governance.CallOpts, proposalId, index)
}

// IsDequeuedProposalExpired is a free data retrieval call binding the contract method 0x6f2ab693.
//
// Solidity: function isDequeuedProposalExpired(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCaller) IsDequeuedProposalExpired(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isDequeuedProposalExpired", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDequeuedProposalExpired is a free data retrieval call binding the contract method 0x6f2ab693.
//
// Solidity: function isDequeuedProposalExpired(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceSession) IsDequeuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsDequeuedProposalExpired is a free data retrieval call binding the contract method 0x6f2ab693.
//
// Solidity: function isDequeuedProposalExpired(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCallerSession) IsDequeuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsDequeuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsHotfixPassing is a free data retrieval call binding the contract method 0x344944cf.
//
// Solidity: function isHotfixPassing(bytes32 hash) view returns(bool)
func (_Governance *GovernanceCaller) IsHotfixPassing(opts *bind.CallOpts, hash [32]byte) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isHotfixPassing", hash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHotfixPassing is a free data retrieval call binding the contract method 0x344944cf.
//
// Solidity: function isHotfixPassing(bytes32 hash) view returns(bool)
func (_Governance *GovernanceSession) IsHotfixPassing(hash [32]byte) (bool, error) {
	return _Governance.Contract.IsHotfixPassing(&_Governance.CallOpts, hash)
}

// IsHotfixPassing is a free data retrieval call binding the contract method 0x344944cf.
//
// Solidity: function isHotfixPassing(bytes32 hash) view returns(bool)
func (_Governance *GovernanceCallerSession) IsHotfixPassing(hash [32]byte) (bool, error) {
	return _Governance.Contract.IsHotfixPassing(&_Governance.CallOpts, hash)
}

// IsHotfixWhitelistedBy is a free data retrieval call binding the contract method 0x3fa5fed0.
//
// Solidity: function isHotfixWhitelistedBy(bytes32 hash, address whitelister) view returns(bool)
func (_Governance *GovernanceCaller) IsHotfixWhitelistedBy(opts *bind.CallOpts, hash [32]byte, whitelister common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isHotfixWhitelistedBy", hash, whitelister)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsHotfixWhitelistedBy is a free data retrieval call binding the contract method 0x3fa5fed0.
//
// Solidity: function isHotfixWhitelistedBy(bytes32 hash, address whitelister) view returns(bool)
func (_Governance *GovernanceSession) IsHotfixWhitelistedBy(hash [32]byte, whitelister common.Address) (bool, error) {
	return _Governance.Contract.IsHotfixWhitelistedBy(&_Governance.CallOpts, hash, whitelister)
}

// IsHotfixWhitelistedBy is a free data retrieval call binding the contract method 0x3fa5fed0.
//
// Solidity: function isHotfixWhitelistedBy(bytes32 hash, address whitelister) view returns(bool)
func (_Governance *GovernanceCallerSession) IsHotfixWhitelistedBy(hash [32]byte, whitelister common.Address) (bool, error) {
	return _Governance.Contract.IsHotfixWhitelistedBy(&_Governance.CallOpts, hash, whitelister)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Governance *GovernanceCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Governance *GovernanceSession) IsOwner() (bool, error) {
	return _Governance.Contract.IsOwner(&_Governance.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Governance *GovernanceCallerSession) IsOwner() (bool, error) {
	return _Governance.Contract.IsOwner(&_Governance.CallOpts)
}

// IsProposalPassing is a free data retrieval call binding the contract method 0x27621321.
//
// Solidity: function isProposalPassing(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCaller) IsProposalPassing(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isProposalPassing", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposalPassing is a free data retrieval call binding the contract method 0x27621321.
//
// Solidity: function isProposalPassing(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceSession) IsProposalPassing(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsProposalPassing(&_Governance.CallOpts, proposalId)
}

// IsProposalPassing is a free data retrieval call binding the contract method 0x27621321.
//
// Solidity: function isProposalPassing(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCallerSession) IsProposalPassing(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsProposalPassing(&_Governance.CallOpts, proposalId)
}

// IsQueued is a free data retrieval call binding the contract method 0xc73a6d78.
//
// Solidity: function isQueued(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCaller) IsQueued(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isQueued", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQueued is a free data retrieval call binding the contract method 0xc73a6d78.
//
// Solidity: function isQueued(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceSession) IsQueued(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueued(&_Governance.CallOpts, proposalId)
}

// IsQueued is a free data retrieval call binding the contract method 0xc73a6d78.
//
// Solidity: function isQueued(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCallerSession) IsQueued(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueued(&_Governance.CallOpts, proposalId)
}

// IsQueuedProposalExpired is a free data retrieval call binding the contract method 0xc134b2fc.
//
// Solidity: function isQueuedProposalExpired(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCaller) IsQueuedProposalExpired(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isQueuedProposalExpired", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQueuedProposalExpired is a free data retrieval call binding the contract method 0xc134b2fc.
//
// Solidity: function isQueuedProposalExpired(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceSession) IsQueuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsQueuedProposalExpired is a free data retrieval call binding the contract method 0xc134b2fc.
//
// Solidity: function isQueuedProposalExpired(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCallerSession) IsQueuedProposalExpired(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.IsQueuedProposalExpired(&_Governance.CallOpts, proposalId)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address account) view returns(bool)
func (_Governance *GovernanceCaller) IsVoting(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "isVoting", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address account) view returns(bool)
func (_Governance *GovernanceSession) IsVoting(account common.Address) (bool, error) {
	return _Governance.Contract.IsVoting(&_Governance.CallOpts, account)
}

// IsVoting is a free data retrieval call binding the contract method 0x5f8dd649.
//
// Solidity: function isVoting(address account) view returns(bool)
func (_Governance *GovernanceCallerSession) IsVoting(account common.Address) (bool, error) {
	return _Governance.Contract.IsVoting(&_Governance.CallOpts, account)
}

// LastDequeue is a free data retrieval call binding the contract method 0xc0aee5f4.
//
// Solidity: function lastDequeue() view returns(uint256)
func (_Governance *GovernanceCaller) LastDequeue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "lastDequeue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastDequeue is a free data retrieval call binding the contract method 0xc0aee5f4.
//
// Solidity: function lastDequeue() view returns(uint256)
func (_Governance *GovernanceSession) LastDequeue() (*big.Int, error) {
	return _Governance.Contract.LastDequeue(&_Governance.CallOpts)
}

// LastDequeue is a free data retrieval call binding the contract method 0xc0aee5f4.
//
// Solidity: function lastDequeue() view returns(uint256)
func (_Governance *GovernanceCallerSession) LastDequeue() (*big.Int, error) {
	return _Governance.Contract.LastDequeue(&_Governance.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_Governance *GovernanceCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "minDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_Governance *GovernanceSession) MinDeposit() (*big.Int, error) {
	return _Governance.Contract.MinDeposit(&_Governance.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x41b3d185.
//
// Solidity: function minDeposit() view returns(uint256)
func (_Governance *GovernanceCallerSession) MinDeposit() (*big.Int, error) {
	return _Governance.Contract.MinDeposit(&_Governance.CallOpts)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCaller) MinQuorumSize(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "minQuorumSize", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.MinQuorumSize(&_Governance.CallOpts, blockNumber)
}

// MinQuorumSize is a free data retrieval call binding the contract method 0xe50e652d.
//
// Solidity: function minQuorumSize(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCallerSession) MinQuorumSize(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.MinQuorumSize(&_Governance.CallOpts, blockNumber)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Governance *GovernanceCaller) MinQuorumSizeInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "minQuorumSizeInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Governance *GovernanceSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.MinQuorumSizeInCurrentSet(&_Governance.CallOpts)
}

// MinQuorumSizeInCurrentSet is a free data retrieval call binding the contract method 0x7385e5da.
//
// Solidity: function minQuorumSizeInCurrentSet() view returns(uint256)
func (_Governance *GovernanceCallerSession) MinQuorumSizeInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.MinQuorumSizeInCurrentSet(&_Governance.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Governance *GovernanceCaller) NumberValidatorsInCurrentSet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "numberValidatorsInCurrentSet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Governance *GovernanceSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInCurrentSet(&_Governance.CallOpts)
}

// NumberValidatorsInCurrentSet is a free data retrieval call binding the contract method 0x87ee8a0f.
//
// Solidity: function numberValidatorsInCurrentSet() view returns(uint256)
func (_Governance *GovernanceCallerSession) NumberValidatorsInCurrentSet() (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInCurrentSet(&_Governance.CallOpts)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCaller) NumberValidatorsInSet(opts *bind.CallOpts, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "numberValidatorsInSet", blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInSet(&_Governance.CallOpts, blockNumber)
}

// NumberValidatorsInSet is a free data retrieval call binding the contract method 0x9b2b592f.
//
// Solidity: function numberValidatorsInSet(uint256 blockNumber) view returns(uint256)
func (_Governance *GovernanceCallerSession) NumberValidatorsInSet(blockNumber *big.Int) (*big.Int, error) {
	return _Governance.Contract.NumberValidatorsInSet(&_Governance.CallOpts, blockNumber)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governance *GovernanceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governance *GovernanceSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Governance *GovernanceCallerSession) Owner() (common.Address, error) {
	return _Governance.Contract.Owner(&_Governance.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Governance *GovernanceCaller) ProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Governance *GovernanceSession) ProposalCount() (*big.Int, error) {
	return _Governance.Contract.ProposalCount(&_Governance.CallOpts)
}

// ProposalCount is a free data retrieval call binding the contract method 0xda35c664.
//
// Solidity: function proposalCount() view returns(uint256)
func (_Governance *GovernanceCallerSession) ProposalCount() (*big.Int, error) {
	return _Governance.Contract.ProposalCount(&_Governance.CallOpts)
}

// ProposalExists is a free data retrieval call binding the contract method 0x1374b22d.
//
// Solidity: function proposalExists(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCaller) ProposalExists(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "proposalExists", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalExists is a free data retrieval call binding the contract method 0x1374b22d.
//
// Solidity: function proposalExists(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceSession) ProposalExists(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.ProposalExists(&_Governance.CallOpts, proposalId)
}

// ProposalExists is a free data retrieval call binding the contract method 0x1374b22d.
//
// Solidity: function proposalExists(uint256 proposalId) view returns(bool)
func (_Governance *GovernanceCallerSession) ProposalExists(proposalId *big.Int) (bool, error) {
	return _Governance.Contract.ProposalExists(&_Governance.CallOpts, proposalId)
}

// QueueExpiry is a free data retrieval call binding the contract method 0x8e905ed6.
//
// Solidity: function queueExpiry() view returns(uint256)
func (_Governance *GovernanceCaller) QueueExpiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "queueExpiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QueueExpiry is a free data retrieval call binding the contract method 0x8e905ed6.
//
// Solidity: function queueExpiry() view returns(uint256)
func (_Governance *GovernanceSession) QueueExpiry() (*big.Int, error) {
	return _Governance.Contract.QueueExpiry(&_Governance.CallOpts)
}

// QueueExpiry is a free data retrieval call binding the contract method 0x8e905ed6.
//
// Solidity: function queueExpiry() view returns(uint256)
func (_Governance *GovernanceCallerSession) QueueExpiry() (*big.Int, error) {
	return _Governance.Contract.QueueExpiry(&_Governance.CallOpts)
}

// RefundedDeposits is a free data retrieval call binding the contract method 0x60b4d34d.
//
// Solidity: function refundedDeposits(address ) view returns(uint256)
func (_Governance *GovernanceCaller) RefundedDeposits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "refundedDeposits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundedDeposits is a free data retrieval call binding the contract method 0x60b4d34d.
//
// Solidity: function refundedDeposits(address ) view returns(uint256)
func (_Governance *GovernanceSession) RefundedDeposits(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.RefundedDeposits(&_Governance.CallOpts, arg0)
}

// RefundedDeposits is a free data retrieval call binding the contract method 0x60b4d34d.
//
// Solidity: function refundedDeposits(address ) view returns(uint256)
func (_Governance *GovernanceCallerSession) RefundedDeposits(arg0 common.Address) (*big.Int, error) {
	return _Governance.Contract.RefundedDeposits(&_Governance.CallOpts, arg0)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Governance *GovernanceCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Governance *GovernanceSession) Registry() (common.Address, error) {
	return _Governance.Contract.Registry(&_Governance.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_Governance *GovernanceCallerSession) Registry() (common.Address, error) {
	return _Governance.Contract.Registry(&_Governance.CallOpts)
}

// StageDurations is a free data retrieval call binding the contract method 0x0f717e42.
//
// Solidity: function stageDurations() view returns(uint256 approval, uint256 referendum, uint256 execution)
func (_Governance *GovernanceCaller) StageDurations(opts *bind.CallOpts) (struct {
	Approval   *big.Int
	Referendum *big.Int
	Execution  *big.Int
}, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "stageDurations")

	outstruct := new(struct {
		Approval   *big.Int
		Referendum *big.Int
		Execution  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Approval = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Referendum = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Execution = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// StageDurations is a free data retrieval call binding the contract method 0x0f717e42.
//
// Solidity: function stageDurations() view returns(uint256 approval, uint256 referendum, uint256 execution)
func (_Governance *GovernanceSession) StageDurations() (struct {
	Approval   *big.Int
	Referendum *big.Int
	Execution  *big.Int
}, error) {
	return _Governance.Contract.StageDurations(&_Governance.CallOpts)
}

// StageDurations is a free data retrieval call binding the contract method 0x0f717e42.
//
// Solidity: function stageDurations() view returns(uint256 approval, uint256 referendum, uint256 execution)
func (_Governance *GovernanceCallerSession) StageDurations() (struct {
	Approval   *big.Int
	Referendum *big.Int
	Execution  *big.Int
}, error) {
	return _Governance.Contract.StageDurations(&_Governance.CallOpts)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Governance *GovernanceCaller) ValidatorSignerAddressFromCurrentSet(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "validatorSignerAddressFromCurrentSet", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Governance *GovernanceSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromCurrentSet(&_Governance.CallOpts, index)
}

// ValidatorSignerAddressFromCurrentSet is a free data retrieval call binding the contract method 0x123633ea.
//
// Solidity: function validatorSignerAddressFromCurrentSet(uint256 index) view returns(address)
func (_Governance *GovernanceCallerSession) ValidatorSignerAddressFromCurrentSet(index *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromCurrentSet(&_Governance.CallOpts, index)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Governance *GovernanceCaller) ValidatorSignerAddressFromSet(opts *bind.CallOpts, index *big.Int, blockNumber *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Governance.contract.Call(opts, &out, "validatorSignerAddressFromSet", index, blockNumber)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Governance *GovernanceSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromSet(&_Governance.CallOpts, index, blockNumber)
}

// ValidatorSignerAddressFromSet is a free data retrieval call binding the contract method 0x5d180adb.
//
// Solidity: function validatorSignerAddressFromSet(uint256 index, uint256 blockNumber) view returns(address)
func (_Governance *GovernanceCallerSession) ValidatorSignerAddressFromSet(index *big.Int, blockNumber *big.Int) (common.Address, error) {
	return _Governance.Contract.ValidatorSignerAddressFromSet(&_Governance.CallOpts, index, blockNumber)
}

// Approve is a paid mutator transaction binding the contract method 0x5d35a3d9.
//
// Solidity: function approve(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactor) Approve(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "approve", proposalId, index)
}

// Approve is a paid mutator transaction binding the contract method 0x5d35a3d9.
//
// Solidity: function approve(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceSession) Approve(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Approve(&_Governance.TransactOpts, proposalId, index)
}

// Approve is a paid mutator transaction binding the contract method 0x5d35a3d9.
//
// Solidity: function approve(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactorSession) Approve(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Approve(&_Governance.TransactOpts, proposalId, index)
}

// ApproveHotfix is a paid mutator transaction binding the contract method 0xb0f99842.
//
// Solidity: function approveHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactor) ApproveHotfix(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "approveHotfix", hash)
}

// ApproveHotfix is a paid mutator transaction binding the contract method 0xb0f99842.
//
// Solidity: function approveHotfix(bytes32 hash) returns()
func (_Governance *GovernanceSession) ApproveHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ApproveHotfix(&_Governance.TransactOpts, hash)
}

// ApproveHotfix is a paid mutator transaction binding the contract method 0xb0f99842.
//
// Solidity: function approveHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactorSession) ApproveHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ApproveHotfix(&_Governance.TransactOpts, hash)
}

// DequeueProposalIfReady is a paid mutator transaction binding the contract method 0xbd3a7d0f.
//
// Solidity: function dequeueProposalIfReady(uint256 proposalId) returns(bool isProposalDequeued)
func (_Governance *GovernanceTransactor) DequeueProposalIfReady(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "dequeueProposalIfReady", proposalId)
}

// DequeueProposalIfReady is a paid mutator transaction binding the contract method 0xbd3a7d0f.
//
// Solidity: function dequeueProposalIfReady(uint256 proposalId) returns(bool isProposalDequeued)
func (_Governance *GovernanceSession) DequeueProposalIfReady(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.DequeueProposalIfReady(&_Governance.TransactOpts, proposalId)
}

// DequeueProposalIfReady is a paid mutator transaction binding the contract method 0xbd3a7d0f.
//
// Solidity: function dequeueProposalIfReady(uint256 proposalId) returns(bool isProposalDequeued)
func (_Governance *GovernanceTransactorSession) DequeueProposalIfReady(proposalId *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.DequeueProposalIfReady(&_Governance.TransactOpts, proposalId)
}

// DequeueProposalsIfReady is a paid mutator transaction binding the contract method 0x3bb0ed2b.
//
// Solidity: function dequeueProposalsIfReady() returns()
func (_Governance *GovernanceTransactor) DequeueProposalsIfReady(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "dequeueProposalsIfReady")
}

// DequeueProposalsIfReady is a paid mutator transaction binding the contract method 0x3bb0ed2b.
//
// Solidity: function dequeueProposalsIfReady() returns()
func (_Governance *GovernanceSession) DequeueProposalsIfReady() (*types.Transaction, error) {
	return _Governance.Contract.DequeueProposalsIfReady(&_Governance.TransactOpts)
}

// DequeueProposalsIfReady is a paid mutator transaction binding the contract method 0x3bb0ed2b.
//
// Solidity: function dequeueProposalsIfReady() returns()
func (_Governance *GovernanceTransactorSession) DequeueProposalsIfReady() (*types.Transaction, error) {
	return _Governance.Contract.DequeueProposalsIfReady(&_Governance.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x5601eaea.
//
// Solidity: function execute(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactor) Execute(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "execute", proposalId, index)
}

// Execute is a paid mutator transaction binding the contract method 0x5601eaea.
//
// Solidity: function execute(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceSession) Execute(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, proposalId, index)
}

// Execute is a paid mutator transaction binding the contract method 0x5601eaea.
//
// Solidity: function execute(uint256 proposalId, uint256 index) returns(bool)
func (_Governance *GovernanceTransactorSession) Execute(proposalId *big.Int, index *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Execute(&_Governance.TransactOpts, proposalId, index)
}

// ExecuteHotfix is a paid mutator transaction binding the contract method 0xcf48eb94.
//
// Solidity: function executeHotfix(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, bytes32 salt) returns()
func (_Governance *GovernanceTransactor) ExecuteHotfix(opts *bind.TransactOpts, values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "executeHotfix", values, destinations, data, dataLengths, salt)
}

// ExecuteHotfix is a paid mutator transaction binding the contract method 0xcf48eb94.
//
// Solidity: function executeHotfix(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, bytes32 salt) returns()
func (_Governance *GovernanceSession) ExecuteHotfix(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ExecuteHotfix(&_Governance.TransactOpts, values, destinations, data, dataLengths, salt)
}

// ExecuteHotfix is a paid mutator transaction binding the contract method 0xcf48eb94.
//
// Solidity: function executeHotfix(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, bytes32 salt) returns()
func (_Governance *GovernanceTransactorSession) ExecuteHotfix(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, salt [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.ExecuteHotfix(&_Governance.TransactOpts, values, destinations, data, dataLengths, salt)
}

// Initialize is a paid mutator transaction binding the contract method 0xaf203110.
//
// Solidity: function initialize(address registryAddress, address _approver, uint256 _concurrentProposals, uint256 _minDeposit, uint256 _queueExpiry, uint256 _dequeueFrequency, uint256 referendumStageDuration, uint256 executionStageDuration, uint256 participationBaseline, uint256 participationFloor, uint256 baselineUpdateFactor, uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactor) Initialize(opts *bind.TransactOpts, registryAddress common.Address, _approver common.Address, _concurrentProposals *big.Int, _minDeposit *big.Int, _queueExpiry *big.Int, _dequeueFrequency *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int, participationBaseline *big.Int, participationFloor *big.Int, baselineUpdateFactor *big.Int, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "initialize", registryAddress, _approver, _concurrentProposals, _minDeposit, _queueExpiry, _dequeueFrequency, referendumStageDuration, executionStageDuration, participationBaseline, participationFloor, baselineUpdateFactor, baselineQuorumFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0xaf203110.
//
// Solidity: function initialize(address registryAddress, address _approver, uint256 _concurrentProposals, uint256 _minDeposit, uint256 _queueExpiry, uint256 _dequeueFrequency, uint256 referendumStageDuration, uint256 executionStageDuration, uint256 participationBaseline, uint256 participationFloor, uint256 baselineUpdateFactor, uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceSession) Initialize(registryAddress common.Address, _approver common.Address, _concurrentProposals *big.Int, _minDeposit *big.Int, _queueExpiry *big.Int, _dequeueFrequency *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int, participationBaseline *big.Int, participationFloor *big.Int, baselineUpdateFactor *big.Int, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, registryAddress, _approver, _concurrentProposals, _minDeposit, _queueExpiry, _dequeueFrequency, referendumStageDuration, executionStageDuration, participationBaseline, participationFloor, baselineUpdateFactor, baselineQuorumFactor)
}

// Initialize is a paid mutator transaction binding the contract method 0xaf203110.
//
// Solidity: function initialize(address registryAddress, address _approver, uint256 _concurrentProposals, uint256 _minDeposit, uint256 _queueExpiry, uint256 _dequeueFrequency, uint256 referendumStageDuration, uint256 executionStageDuration, uint256 participationBaseline, uint256 participationFloor, uint256 baselineUpdateFactor, uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactorSession) Initialize(registryAddress common.Address, _approver common.Address, _concurrentProposals *big.Int, _minDeposit *big.Int, _queueExpiry *big.Int, _dequeueFrequency *big.Int, referendumStageDuration *big.Int, executionStageDuration *big.Int, participationBaseline *big.Int, participationFloor *big.Int, baselineUpdateFactor *big.Int, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Initialize(&_Governance.TransactOpts, registryAddress, _approver, _concurrentProposals, _minDeposit, _queueExpiry, _dequeueFrequency, referendumStageDuration, executionStageDuration, participationBaseline, participationFloor, baselineUpdateFactor, baselineQuorumFactor)
}

// PrepareHotfix is a paid mutator transaction binding the contract method 0x9cb02dfc.
//
// Solidity: function prepareHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactor) PrepareHotfix(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "prepareHotfix", hash)
}

// PrepareHotfix is a paid mutator transaction binding the contract method 0x9cb02dfc.
//
// Solidity: function prepareHotfix(bytes32 hash) returns()
func (_Governance *GovernanceSession) PrepareHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.PrepareHotfix(&_Governance.TransactOpts, hash)
}

// PrepareHotfix is a paid mutator transaction binding the contract method 0x9cb02dfc.
//
// Solidity: function prepareHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactorSession) PrepareHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.PrepareHotfix(&_Governance.TransactOpts, hash)
}

// Propose is a paid mutator transaction binding the contract method 0x65bbdaa0.
//
// Solidity: function propose(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, string descriptionUrl) payable returns(uint256)
func (_Governance *GovernanceTransactor) Propose(opts *bind.TransactOpts, values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, descriptionUrl string) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "propose", values, destinations, data, dataLengths, descriptionUrl)
}

// Propose is a paid mutator transaction binding the contract method 0x65bbdaa0.
//
// Solidity: function propose(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, string descriptionUrl) payable returns(uint256)
func (_Governance *GovernanceSession) Propose(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, descriptionUrl string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, values, destinations, data, dataLengths, descriptionUrl)
}

// Propose is a paid mutator transaction binding the contract method 0x65bbdaa0.
//
// Solidity: function propose(uint256[] values, address[] destinations, bytes data, uint256[] dataLengths, string descriptionUrl) payable returns(uint256)
func (_Governance *GovernanceTransactorSession) Propose(values []*big.Int, destinations []common.Address, data []byte, dataLengths []*big.Int, descriptionUrl string) (*types.Transaction, error) {
	return _Governance.Contract.Propose(&_Governance.TransactOpts, values, destinations, data, dataLengths, descriptionUrl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governance.Contract.RenounceOwnership(&_Governance.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Governance *GovernanceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Governance.Contract.RenounceOwnership(&_Governance.TransactOpts)
}

// RevokeUpvote is a paid mutator transaction binding the contract method 0xaf108a0e.
//
// Solidity: function revokeUpvote(uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactor) RevokeUpvote(opts *bind.TransactOpts, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "revokeUpvote", lesser, greater)
}

// RevokeUpvote is a paid mutator transaction binding the contract method 0xaf108a0e.
//
// Solidity: function revokeUpvote(uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceSession) RevokeUpvote(lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.RevokeUpvote(&_Governance.TransactOpts, lesser, greater)
}

// RevokeUpvote is a paid mutator transaction binding the contract method 0xaf108a0e.
//
// Solidity: function revokeUpvote(uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactorSession) RevokeUpvote(lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.RevokeUpvote(&_Governance.TransactOpts, lesser, greater)
}

// RevokeVotes is a paid mutator transaction binding the contract method 0x9381ab25.
//
// Solidity: function revokeVotes() returns(bool)
func (_Governance *GovernanceTransactor) RevokeVotes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "revokeVotes")
}

// RevokeVotes is a paid mutator transaction binding the contract method 0x9381ab25.
//
// Solidity: function revokeVotes() returns(bool)
func (_Governance *GovernanceSession) RevokeVotes() (*types.Transaction, error) {
	return _Governance.Contract.RevokeVotes(&_Governance.TransactOpts)
}

// RevokeVotes is a paid mutator transaction binding the contract method 0x9381ab25.
//
// Solidity: function revokeVotes() returns(bool)
func (_Governance *GovernanceTransactorSession) RevokeVotes() (*types.Transaction, error) {
	return _Governance.Contract.RevokeVotes(&_Governance.TransactOpts)
}

// SetApprover is a paid mutator transaction binding the contract method 0x3156560e.
//
// Solidity: function setApprover(address _approver) returns()
func (_Governance *GovernanceTransactor) SetApprover(opts *bind.TransactOpts, _approver common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setApprover", _approver)
}

// SetApprover is a paid mutator transaction binding the contract method 0x3156560e.
//
// Solidity: function setApprover(address _approver) returns()
func (_Governance *GovernanceSession) SetApprover(_approver common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetApprover(&_Governance.TransactOpts, _approver)
}

// SetApprover is a paid mutator transaction binding the contract method 0x3156560e.
//
// Solidity: function setApprover(address _approver) returns()
func (_Governance *GovernanceTransactorSession) SetApprover(_approver common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetApprover(&_Governance.TransactOpts, _approver)
}

// SetBaselineQuorumFactor is a paid mutator transaction binding the contract method 0x04acaec9.
//
// Solidity: function setBaselineQuorumFactor(uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactor) SetBaselineQuorumFactor(opts *bind.TransactOpts, baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setBaselineQuorumFactor", baselineQuorumFactor)
}

// SetBaselineQuorumFactor is a paid mutator transaction binding the contract method 0x04acaec9.
//
// Solidity: function setBaselineQuorumFactor(uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceSession) SetBaselineQuorumFactor(baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineQuorumFactor(&_Governance.TransactOpts, baselineQuorumFactor)
}

// SetBaselineQuorumFactor is a paid mutator transaction binding the contract method 0x04acaec9.
//
// Solidity: function setBaselineQuorumFactor(uint256 baselineQuorumFactor) returns()
func (_Governance *GovernanceTransactorSession) SetBaselineQuorumFactor(baselineQuorumFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineQuorumFactor(&_Governance.TransactOpts, baselineQuorumFactor)
}

// SetBaselineUpdateFactor is a paid mutator transaction binding the contract method 0x5c759394.
//
// Solidity: function setBaselineUpdateFactor(uint256 baselineUpdateFactor) returns()
func (_Governance *GovernanceTransactor) SetBaselineUpdateFactor(opts *bind.TransactOpts, baselineUpdateFactor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setBaselineUpdateFactor", baselineUpdateFactor)
}

// SetBaselineUpdateFactor is a paid mutator transaction binding the contract method 0x5c759394.
//
// Solidity: function setBaselineUpdateFactor(uint256 baselineUpdateFactor) returns()
func (_Governance *GovernanceSession) SetBaselineUpdateFactor(baselineUpdateFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineUpdateFactor(&_Governance.TransactOpts, baselineUpdateFactor)
}

// SetBaselineUpdateFactor is a paid mutator transaction binding the contract method 0x5c759394.
//
// Solidity: function setBaselineUpdateFactor(uint256 baselineUpdateFactor) returns()
func (_Governance *GovernanceTransactorSession) SetBaselineUpdateFactor(baselineUpdateFactor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetBaselineUpdateFactor(&_Governance.TransactOpts, baselineUpdateFactor)
}

// SetConcurrentProposals is a paid mutator transaction binding the contract method 0xc8d8d2b5.
//
// Solidity: function setConcurrentProposals(uint256 _concurrentProposals) returns()
func (_Governance *GovernanceTransactor) SetConcurrentProposals(opts *bind.TransactOpts, _concurrentProposals *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setConcurrentProposals", _concurrentProposals)
}

// SetConcurrentProposals is a paid mutator transaction binding the contract method 0xc8d8d2b5.
//
// Solidity: function setConcurrentProposals(uint256 _concurrentProposals) returns()
func (_Governance *GovernanceSession) SetConcurrentProposals(_concurrentProposals *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConcurrentProposals(&_Governance.TransactOpts, _concurrentProposals)
}

// SetConcurrentProposals is a paid mutator transaction binding the contract method 0xc8d8d2b5.
//
// Solidity: function setConcurrentProposals(uint256 _concurrentProposals) returns()
func (_Governance *GovernanceTransactorSession) SetConcurrentProposals(_concurrentProposals *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConcurrentProposals(&_Governance.TransactOpts, _concurrentProposals)
}

// SetConstitution is a paid mutator transaction binding the contract method 0xed385274.
//
// Solidity: function setConstitution(address destination, bytes4 functionId, uint256 threshold) returns()
func (_Governance *GovernanceTransactor) SetConstitution(opts *bind.TransactOpts, destination common.Address, functionId [4]byte, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setConstitution", destination, functionId, threshold)
}

// SetConstitution is a paid mutator transaction binding the contract method 0xed385274.
//
// Solidity: function setConstitution(address destination, bytes4 functionId, uint256 threshold) returns()
func (_Governance *GovernanceSession) SetConstitution(destination common.Address, functionId [4]byte, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConstitution(&_Governance.TransactOpts, destination, functionId, threshold)
}

// SetConstitution is a paid mutator transaction binding the contract method 0xed385274.
//
// Solidity: function setConstitution(address destination, bytes4 functionId, uint256 threshold) returns()
func (_Governance *GovernanceTransactorSession) SetConstitution(destination common.Address, functionId [4]byte, threshold *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetConstitution(&_Governance.TransactOpts, destination, functionId, threshold)
}

// SetDequeueFrequency is a paid mutator transaction binding the contract method 0x8018556e.
//
// Solidity: function setDequeueFrequency(uint256 _dequeueFrequency) returns()
func (_Governance *GovernanceTransactor) SetDequeueFrequency(opts *bind.TransactOpts, _dequeueFrequency *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setDequeueFrequency", _dequeueFrequency)
}

// SetDequeueFrequency is a paid mutator transaction binding the contract method 0x8018556e.
//
// Solidity: function setDequeueFrequency(uint256 _dequeueFrequency) returns()
func (_Governance *GovernanceSession) SetDequeueFrequency(_dequeueFrequency *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetDequeueFrequency(&_Governance.TransactOpts, _dequeueFrequency)
}

// SetDequeueFrequency is a paid mutator transaction binding the contract method 0x8018556e.
//
// Solidity: function setDequeueFrequency(uint256 _dequeueFrequency) returns()
func (_Governance *GovernanceTransactorSession) SetDequeueFrequency(_dequeueFrequency *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetDequeueFrequency(&_Governance.TransactOpts, _dequeueFrequency)
}

// SetExecutionStageDuration is a paid mutator transaction binding the contract method 0x6643ac58.
//
// Solidity: function setExecutionStageDuration(uint256 executionStageDuration) returns()
func (_Governance *GovernanceTransactor) SetExecutionStageDuration(opts *bind.TransactOpts, executionStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setExecutionStageDuration", executionStageDuration)
}

// SetExecutionStageDuration is a paid mutator transaction binding the contract method 0x6643ac58.
//
// Solidity: function setExecutionStageDuration(uint256 executionStageDuration) returns()
func (_Governance *GovernanceSession) SetExecutionStageDuration(executionStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetExecutionStageDuration(&_Governance.TransactOpts, executionStageDuration)
}

// SetExecutionStageDuration is a paid mutator transaction binding the contract method 0x6643ac58.
//
// Solidity: function setExecutionStageDuration(uint256 executionStageDuration) returns()
func (_Governance *GovernanceTransactorSession) SetExecutionStageDuration(executionStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetExecutionStageDuration(&_Governance.TransactOpts, executionStageDuration)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _minDeposit) returns()
func (_Governance *GovernanceTransactor) SetMinDeposit(opts *bind.TransactOpts, _minDeposit *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setMinDeposit", _minDeposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _minDeposit) returns()
func (_Governance *GovernanceSession) SetMinDeposit(_minDeposit *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinDeposit(&_Governance.TransactOpts, _minDeposit)
}

// SetMinDeposit is a paid mutator transaction binding the contract method 0x8fcc9cfb.
//
// Solidity: function setMinDeposit(uint256 _minDeposit) returns()
func (_Governance *GovernanceTransactorSession) SetMinDeposit(_minDeposit *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetMinDeposit(&_Governance.TransactOpts, _minDeposit)
}

// SetParticipationBaseline is a paid mutator transaction binding the contract method 0x3db9dd9a.
//
// Solidity: function setParticipationBaseline(uint256 participationBaseline) returns()
func (_Governance *GovernanceTransactor) SetParticipationBaseline(opts *bind.TransactOpts, participationBaseline *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setParticipationBaseline", participationBaseline)
}

// SetParticipationBaseline is a paid mutator transaction binding the contract method 0x3db9dd9a.
//
// Solidity: function setParticipationBaseline(uint256 participationBaseline) returns()
func (_Governance *GovernanceSession) SetParticipationBaseline(participationBaseline *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationBaseline(&_Governance.TransactOpts, participationBaseline)
}

// SetParticipationBaseline is a paid mutator transaction binding the contract method 0x3db9dd9a.
//
// Solidity: function setParticipationBaseline(uint256 participationBaseline) returns()
func (_Governance *GovernanceTransactorSession) SetParticipationBaseline(participationBaseline *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationBaseline(&_Governance.TransactOpts, participationBaseline)
}

// SetParticipationFloor is a paid mutator transaction binding the contract method 0x1c65bc61.
//
// Solidity: function setParticipationFloor(uint256 participationFloor) returns()
func (_Governance *GovernanceTransactor) SetParticipationFloor(opts *bind.TransactOpts, participationFloor *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setParticipationFloor", participationFloor)
}

// SetParticipationFloor is a paid mutator transaction binding the contract method 0x1c65bc61.
//
// Solidity: function setParticipationFloor(uint256 participationFloor) returns()
func (_Governance *GovernanceSession) SetParticipationFloor(participationFloor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationFloor(&_Governance.TransactOpts, participationFloor)
}

// SetParticipationFloor is a paid mutator transaction binding the contract method 0x1c65bc61.
//
// Solidity: function setParticipationFloor(uint256 participationFloor) returns()
func (_Governance *GovernanceTransactorSession) SetParticipationFloor(participationFloor *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetParticipationFloor(&_Governance.TransactOpts, participationFloor)
}

// SetQueueExpiry is a paid mutator transaction binding the contract method 0x2c052355.
//
// Solidity: function setQueueExpiry(uint256 _queueExpiry) returns()
func (_Governance *GovernanceTransactor) SetQueueExpiry(opts *bind.TransactOpts, _queueExpiry *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setQueueExpiry", _queueExpiry)
}

// SetQueueExpiry is a paid mutator transaction binding the contract method 0x2c052355.
//
// Solidity: function setQueueExpiry(uint256 _queueExpiry) returns()
func (_Governance *GovernanceSession) SetQueueExpiry(_queueExpiry *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetQueueExpiry(&_Governance.TransactOpts, _queueExpiry)
}

// SetQueueExpiry is a paid mutator transaction binding the contract method 0x2c052355.
//
// Solidity: function setQueueExpiry(uint256 _queueExpiry) returns()
func (_Governance *GovernanceTransactorSession) SetQueueExpiry(_queueExpiry *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetQueueExpiry(&_Governance.TransactOpts, _queueExpiry)
}

// SetReferendumStageDuration is a paid mutator transaction binding the contract method 0xcea69e74.
//
// Solidity: function setReferendumStageDuration(uint256 referendumStageDuration) returns()
func (_Governance *GovernanceTransactor) SetReferendumStageDuration(opts *bind.TransactOpts, referendumStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setReferendumStageDuration", referendumStageDuration)
}

// SetReferendumStageDuration is a paid mutator transaction binding the contract method 0xcea69e74.
//
// Solidity: function setReferendumStageDuration(uint256 referendumStageDuration) returns()
func (_Governance *GovernanceSession) SetReferendumStageDuration(referendumStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetReferendumStageDuration(&_Governance.TransactOpts, referendumStageDuration)
}

// SetReferendumStageDuration is a paid mutator transaction binding the contract method 0xcea69e74.
//
// Solidity: function setReferendumStageDuration(uint256 referendumStageDuration) returns()
func (_Governance *GovernanceTransactorSession) SetReferendumStageDuration(referendumStageDuration *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.SetReferendumStageDuration(&_Governance.TransactOpts, referendumStageDuration)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Governance *GovernanceTransactor) SetRegistry(opts *bind.TransactOpts, registryAddress common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "setRegistry", registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Governance *GovernanceSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetRegistry(&_Governance.TransactOpts, registryAddress)
}

// SetRegistry is a paid mutator transaction binding the contract method 0xa91ee0dc.
//
// Solidity: function setRegistry(address registryAddress) returns()
func (_Governance *GovernanceTransactorSession) SetRegistry(registryAddress common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SetRegistry(&_Governance.TransactOpts, registryAddress)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Governance *GovernanceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Governance.Contract.TransferOwnership(&_Governance.TransactOpts, newOwner)
}

// Upvote is a paid mutator transaction binding the contract method 0x57333978.
//
// Solidity: function upvote(uint256 proposalId, uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactor) Upvote(opts *bind.TransactOpts, proposalId *big.Int, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "upvote", proposalId, lesser, greater)
}

// Upvote is a paid mutator transaction binding the contract method 0x57333978.
//
// Solidity: function upvote(uint256 proposalId, uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceSession) Upvote(proposalId *big.Int, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Upvote(&_Governance.TransactOpts, proposalId, lesser, greater)
}

// Upvote is a paid mutator transaction binding the contract method 0x57333978.
//
// Solidity: function upvote(uint256 proposalId, uint256 lesser, uint256 greater) returns(bool)
func (_Governance *GovernanceTransactorSession) Upvote(proposalId *big.Int, lesser *big.Int, greater *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Upvote(&_Governance.TransactOpts, proposalId, lesser, greater)
}

// Vote is a paid mutator transaction binding the contract method 0xbbb2eab9.
//
// Solidity: function vote(uint256 proposalId, uint256 index, uint8 value) returns(bool)
func (_Governance *GovernanceTransactor) Vote(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int, value uint8) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "vote", proposalId, index, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbbb2eab9.
//
// Solidity: function vote(uint256 proposalId, uint256 index, uint8 value) returns(bool)
func (_Governance *GovernanceSession) Vote(proposalId *big.Int, index *big.Int, value uint8) (*types.Transaction, error) {
	return _Governance.Contract.Vote(&_Governance.TransactOpts, proposalId, index, value)
}

// Vote is a paid mutator transaction binding the contract method 0xbbb2eab9.
//
// Solidity: function vote(uint256 proposalId, uint256 index, uint8 value) returns(bool)
func (_Governance *GovernanceTransactorSession) Vote(proposalId *big.Int, index *big.Int, value uint8) (*types.Transaction, error) {
	return _Governance.Contract.Vote(&_Governance.TransactOpts, proposalId, index, value)
}

// VotePartially is a paid mutator transaction binding the contract method 0x2edfd12e.
//
// Solidity: function votePartially(uint256 proposalId, uint256 index, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes) returns(bool)
func (_Governance *GovernanceTransactor) VotePartially(opts *bind.TransactOpts, proposalId *big.Int, index *big.Int, yesVotes *big.Int, noVotes *big.Int, abstainVotes *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "votePartially", proposalId, index, yesVotes, noVotes, abstainVotes)
}

// VotePartially is a paid mutator transaction binding the contract method 0x2edfd12e.
//
// Solidity: function votePartially(uint256 proposalId, uint256 index, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes) returns(bool)
func (_Governance *GovernanceSession) VotePartially(proposalId *big.Int, index *big.Int, yesVotes *big.Int, noVotes *big.Int, abstainVotes *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.VotePartially(&_Governance.TransactOpts, proposalId, index, yesVotes, noVotes, abstainVotes)
}

// VotePartially is a paid mutator transaction binding the contract method 0x2edfd12e.
//
// Solidity: function votePartially(uint256 proposalId, uint256 index, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes) returns(bool)
func (_Governance *GovernanceTransactorSession) VotePartially(proposalId *big.Int, index *big.Int, yesVotes *big.Int, noVotes *big.Int, abstainVotes *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.VotePartially(&_Governance.TransactOpts, proposalId, index, yesVotes, noVotes, abstainVotes)
}

// WhitelistHotfix is a paid mutator transaction binding the contract method 0xb15f0f58.
//
// Solidity: function whitelistHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactor) WhitelistHotfix(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "whitelistHotfix", hash)
}

// WhitelistHotfix is a paid mutator transaction binding the contract method 0xb15f0f58.
//
// Solidity: function whitelistHotfix(bytes32 hash) returns()
func (_Governance *GovernanceSession) WhitelistHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.WhitelistHotfix(&_Governance.TransactOpts, hash)
}

// WhitelistHotfix is a paid mutator transaction binding the contract method 0xb15f0f58.
//
// Solidity: function whitelistHotfix(bytes32 hash) returns()
func (_Governance *GovernanceTransactorSession) WhitelistHotfix(hash [32]byte) (*types.Transaction, error) {
	return _Governance.Contract.WhitelistHotfix(&_Governance.TransactOpts, hash)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// TryParseLog attempts to parse a log. Returns the parsed log, evenName and whether it was succesfull
func (_Governance *GovernanceFilterer) TryParseLog(log types.Log) (eventName string, event interface{}, ok bool, err error) {
	eventName, ok, err = _Governance.contract.LogEventName(log)
	if err != nil || !ok {
		return "", nil, false, err
	}

	switch eventName {
	case "ApproverSet":
		event, err = _Governance.ParseApproverSet(log)
	case "ConcurrentProposalsSet":
		event, err = _Governance.ParseConcurrentProposalsSet(log)
	case "ConstitutionSet":
		event, err = _Governance.ParseConstitutionSet(log)
	case "DequeueFrequencySet":
		event, err = _Governance.ParseDequeueFrequencySet(log)
	case "ExecutionStageDurationSet":
		event, err = _Governance.ParseExecutionStageDurationSet(log)
	case "HotfixApproved":
		event, err = _Governance.ParseHotfixApproved(log)
	case "HotfixExecuted":
		event, err = _Governance.ParseHotfixExecuted(log)
	case "HotfixPrepared":
		event, err = _Governance.ParseHotfixPrepared(log)
	case "HotfixWhitelisted":
		event, err = _Governance.ParseHotfixWhitelisted(log)
	case "MinDepositSet":
		event, err = _Governance.ParseMinDepositSet(log)
	case "OwnershipTransferred":
		event, err = _Governance.ParseOwnershipTransferred(log)
	case "ParticipationBaselineQuorumFactorSet":
		event, err = _Governance.ParseParticipationBaselineQuorumFactorSet(log)
	case "ParticipationBaselineUpdateFactorSet":
		event, err = _Governance.ParseParticipationBaselineUpdateFactorSet(log)
	case "ParticipationBaselineUpdated":
		event, err = _Governance.ParseParticipationBaselineUpdated(log)
	case "ParticipationFloorSet":
		event, err = _Governance.ParseParticipationFloorSet(log)
	case "ProposalApproved":
		event, err = _Governance.ParseProposalApproved(log)
	case "ProposalDequeued":
		event, err = _Governance.ParseProposalDequeued(log)
	case "ProposalExecuted":
		event, err = _Governance.ParseProposalExecuted(log)
	case "ProposalExpired":
		event, err = _Governance.ParseProposalExpired(log)
	case "ProposalQueued":
		event, err = _Governance.ParseProposalQueued(log)
	case "ProposalUpvoteRevoked":
		event, err = _Governance.ParseProposalUpvoteRevoked(log)
	case "ProposalUpvoted":
		event, err = _Governance.ParseProposalUpvoted(log)
	case "ProposalVoteRevoked":
		event, err = _Governance.ParseProposalVoteRevoked(log)
	case "ProposalVoteRevokedV2":
		event, err = _Governance.ParseProposalVoteRevokedV2(log)
	case "ProposalVoted":
		event, err = _Governance.ParseProposalVoted(log)
	case "ProposalVotedV2":
		event, err = _Governance.ParseProposalVotedV2(log)
	case "QueueExpirySet":
		event, err = _Governance.ParseQueueExpirySet(log)
	case "ReferendumStageDurationSet":
		event, err = _Governance.ParseReferendumStageDurationSet(log)
	case "RegistrySet":
		event, err = _Governance.ParseRegistrySet(log)
	}
	if err != nil {
		return "", nil, false, err
	}

	return eventName, event, ok, nil
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Governance *GovernanceTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Governance.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Governance *GovernanceSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Governance.Contract.Fallback(&_Governance.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Governance *GovernanceTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Governance.Contract.Fallback(&_Governance.TransactOpts, calldata)
}

// GovernanceApproverSetIterator is returned from FilterApproverSet and is used to iterate over the raw logs and unpacked data for ApproverSet events raised by the Governance contract.
type GovernanceApproverSetIterator struct {
	Event *GovernanceApproverSet // Event containing the contract specifics and raw log

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
func (it *GovernanceApproverSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceApproverSet)
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
		it.Event = new(GovernanceApproverSet)
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
func (it *GovernanceApproverSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceApproverSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceApproverSet represents a ApproverSet event raised by the Governance contract.
type GovernanceApproverSet struct {
	Approver common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproverSet is a free log retrieval operation binding the contract event 0xa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e1.
//
// Solidity: event ApproverSet(address indexed approver)
func (_Governance *GovernanceFilterer) FilterApproverSet(opts *bind.FilterOpts, approver []common.Address) (*GovernanceApproverSetIterator, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ApproverSet", approverRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceApproverSetIterator{contract: _Governance.contract, event: "ApproverSet", logs: logs, sub: sub}, nil
}

// WatchApproverSet is a free log subscription operation binding the contract event 0xa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e1.
//
// Solidity: event ApproverSet(address indexed approver)
func (_Governance *GovernanceFilterer) WatchApproverSet(opts *bind.WatchOpts, sink chan<- *GovernanceApproverSet, approver []common.Address) (event.Subscription, error) {

	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ApproverSet", approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceApproverSet)
				if err := _Governance.contract.UnpackLog(event, "ApproverSet", log); err != nil {
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

// ParseApproverSet is a log parse operation binding the contract event 0xa03757d836cb0b61c0fbba2147f5d51d6071ff3dd5bf6963beb55563d64878e1.
//
// Solidity: event ApproverSet(address indexed approver)
func (_Governance *GovernanceFilterer) ParseApproverSet(log types.Log) (*GovernanceApproverSet, error) {
	event := new(GovernanceApproverSet)
	if err := _Governance.contract.UnpackLog(event, "ApproverSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceConcurrentProposalsSetIterator is returned from FilterConcurrentProposalsSet and is used to iterate over the raw logs and unpacked data for ConcurrentProposalsSet events raised by the Governance contract.
type GovernanceConcurrentProposalsSetIterator struct {
	Event *GovernanceConcurrentProposalsSet // Event containing the contract specifics and raw log

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
func (it *GovernanceConcurrentProposalsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceConcurrentProposalsSet)
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
		it.Event = new(GovernanceConcurrentProposalsSet)
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
func (it *GovernanceConcurrentProposalsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceConcurrentProposalsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceConcurrentProposalsSet represents a ConcurrentProposalsSet event raised by the Governance contract.
type GovernanceConcurrentProposalsSet struct {
	ConcurrentProposals *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterConcurrentProposalsSet is a free log retrieval operation binding the contract event 0x85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a6.
//
// Solidity: event ConcurrentProposalsSet(uint256 concurrentProposals)
func (_Governance *GovernanceFilterer) FilterConcurrentProposalsSet(opts *bind.FilterOpts) (*GovernanceConcurrentProposalsSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ConcurrentProposalsSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceConcurrentProposalsSetIterator{contract: _Governance.contract, event: "ConcurrentProposalsSet", logs: logs, sub: sub}, nil
}

// WatchConcurrentProposalsSet is a free log subscription operation binding the contract event 0x85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a6.
//
// Solidity: event ConcurrentProposalsSet(uint256 concurrentProposals)
func (_Governance *GovernanceFilterer) WatchConcurrentProposalsSet(opts *bind.WatchOpts, sink chan<- *GovernanceConcurrentProposalsSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ConcurrentProposalsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceConcurrentProposalsSet)
				if err := _Governance.contract.UnpackLog(event, "ConcurrentProposalsSet", log); err != nil {
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

// ParseConcurrentProposalsSet is a log parse operation binding the contract event 0x85399b9b2595eb13c392e1638ca77730156cd8d48d4733df4db068e4aa6b93a6.
//
// Solidity: event ConcurrentProposalsSet(uint256 concurrentProposals)
func (_Governance *GovernanceFilterer) ParseConcurrentProposalsSet(log types.Log) (*GovernanceConcurrentProposalsSet, error) {
	event := new(GovernanceConcurrentProposalsSet)
	if err := _Governance.contract.UnpackLog(event, "ConcurrentProposalsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceConstitutionSetIterator is returned from FilterConstitutionSet and is used to iterate over the raw logs and unpacked data for ConstitutionSet events raised by the Governance contract.
type GovernanceConstitutionSetIterator struct {
	Event *GovernanceConstitutionSet // Event containing the contract specifics and raw log

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
func (it *GovernanceConstitutionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceConstitutionSet)
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
		it.Event = new(GovernanceConstitutionSet)
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
func (it *GovernanceConstitutionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceConstitutionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceConstitutionSet represents a ConstitutionSet event raised by the Governance contract.
type GovernanceConstitutionSet struct {
	Destination common.Address
	FunctionId  [4]byte
	Threshold   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConstitutionSet is a free log retrieval operation binding the contract event 0x60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef.
//
// Solidity: event ConstitutionSet(address indexed destination, bytes4 indexed functionId, uint256 threshold)
func (_Governance *GovernanceFilterer) FilterConstitutionSet(opts *bind.FilterOpts, destination []common.Address, functionId [][4]byte) (*GovernanceConstitutionSetIterator, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ConstitutionSet", destinationRule, functionIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceConstitutionSetIterator{contract: _Governance.contract, event: "ConstitutionSet", logs: logs, sub: sub}, nil
}

// WatchConstitutionSet is a free log subscription operation binding the contract event 0x60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef.
//
// Solidity: event ConstitutionSet(address indexed destination, bytes4 indexed functionId, uint256 threshold)
func (_Governance *GovernanceFilterer) WatchConstitutionSet(opts *bind.WatchOpts, sink chan<- *GovernanceConstitutionSet, destination []common.Address, functionId [][4]byte) (event.Subscription, error) {

	var destinationRule []interface{}
	for _, destinationItem := range destination {
		destinationRule = append(destinationRule, destinationItem)
	}
	var functionIdRule []interface{}
	for _, functionIdItem := range functionId {
		functionIdRule = append(functionIdRule, functionIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ConstitutionSet", destinationRule, functionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceConstitutionSet)
				if err := _Governance.contract.UnpackLog(event, "ConstitutionSet", log); err != nil {
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

// ParseConstitutionSet is a log parse operation binding the contract event 0x60c5b4756af49d7b071b00dbf0f87af605cce11896ecd3b760d19f0f9d3fbcef.
//
// Solidity: event ConstitutionSet(address indexed destination, bytes4 indexed functionId, uint256 threshold)
func (_Governance *GovernanceFilterer) ParseConstitutionSet(log types.Log) (*GovernanceConstitutionSet, error) {
	event := new(GovernanceConstitutionSet)
	if err := _Governance.contract.UnpackLog(event, "ConstitutionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceDequeueFrequencySetIterator is returned from FilterDequeueFrequencySet and is used to iterate over the raw logs and unpacked data for DequeueFrequencySet events raised by the Governance contract.
type GovernanceDequeueFrequencySetIterator struct {
	Event *GovernanceDequeueFrequencySet // Event containing the contract specifics and raw log

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
func (it *GovernanceDequeueFrequencySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDequeueFrequencySet)
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
		it.Event = new(GovernanceDequeueFrequencySet)
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
func (it *GovernanceDequeueFrequencySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceDequeueFrequencySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceDequeueFrequencySet represents a DequeueFrequencySet event raised by the Governance contract.
type GovernanceDequeueFrequencySet struct {
	DequeueFrequency *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterDequeueFrequencySet is a free log retrieval operation binding the contract event 0x391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c34.
//
// Solidity: event DequeueFrequencySet(uint256 dequeueFrequency)
func (_Governance *GovernanceFilterer) FilterDequeueFrequencySet(opts *bind.FilterOpts) (*GovernanceDequeueFrequencySetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "DequeueFrequencySet")
	if err != nil {
		return nil, err
	}
	return &GovernanceDequeueFrequencySetIterator{contract: _Governance.contract, event: "DequeueFrequencySet", logs: logs, sub: sub}, nil
}

// WatchDequeueFrequencySet is a free log subscription operation binding the contract event 0x391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c34.
//
// Solidity: event DequeueFrequencySet(uint256 dequeueFrequency)
func (_Governance *GovernanceFilterer) WatchDequeueFrequencySet(opts *bind.WatchOpts, sink chan<- *GovernanceDequeueFrequencySet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "DequeueFrequencySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceDequeueFrequencySet)
				if err := _Governance.contract.UnpackLog(event, "DequeueFrequencySet", log); err != nil {
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

// ParseDequeueFrequencySet is a log parse operation binding the contract event 0x391e82aae76c653cd640ad1b6028e2ee39ca4f29b30152e3d0a9ddd7e1169c34.
//
// Solidity: event DequeueFrequencySet(uint256 dequeueFrequency)
func (_Governance *GovernanceFilterer) ParseDequeueFrequencySet(log types.Log) (*GovernanceDequeueFrequencySet, error) {
	event := new(GovernanceDequeueFrequencySet)
	if err := _Governance.contract.UnpackLog(event, "DequeueFrequencySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceExecutionStageDurationSetIterator is returned from FilterExecutionStageDurationSet and is used to iterate over the raw logs and unpacked data for ExecutionStageDurationSet events raised by the Governance contract.
type GovernanceExecutionStageDurationSetIterator struct {
	Event *GovernanceExecutionStageDurationSet // Event containing the contract specifics and raw log

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
func (it *GovernanceExecutionStageDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceExecutionStageDurationSet)
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
		it.Event = new(GovernanceExecutionStageDurationSet)
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
func (it *GovernanceExecutionStageDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceExecutionStageDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceExecutionStageDurationSet represents a ExecutionStageDurationSet event raised by the Governance contract.
type GovernanceExecutionStageDurationSet struct {
	ExecutionStageDuration *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterExecutionStageDurationSet is a free log retrieval operation binding the contract event 0x7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d8.
//
// Solidity: event ExecutionStageDurationSet(uint256 executionStageDuration)
func (_Governance *GovernanceFilterer) FilterExecutionStageDurationSet(opts *bind.FilterOpts) (*GovernanceExecutionStageDurationSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ExecutionStageDurationSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceExecutionStageDurationSetIterator{contract: _Governance.contract, event: "ExecutionStageDurationSet", logs: logs, sub: sub}, nil
}

// WatchExecutionStageDurationSet is a free log subscription operation binding the contract event 0x7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d8.
//
// Solidity: event ExecutionStageDurationSet(uint256 executionStageDuration)
func (_Governance *GovernanceFilterer) WatchExecutionStageDurationSet(opts *bind.WatchOpts, sink chan<- *GovernanceExecutionStageDurationSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ExecutionStageDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceExecutionStageDurationSet)
				if err := _Governance.contract.UnpackLog(event, "ExecutionStageDurationSet", log); err != nil {
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

// ParseExecutionStageDurationSet is a log parse operation binding the contract event 0x7819c8059302d1d66abc7fe228ecc02214e0bc5c529956c13717aabefce937d8.
//
// Solidity: event ExecutionStageDurationSet(uint256 executionStageDuration)
func (_Governance *GovernanceFilterer) ParseExecutionStageDurationSet(log types.Log) (*GovernanceExecutionStageDurationSet, error) {
	event := new(GovernanceExecutionStageDurationSet)
	if err := _Governance.contract.UnpackLog(event, "ExecutionStageDurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceHotfixApprovedIterator is returned from FilterHotfixApproved and is used to iterate over the raw logs and unpacked data for HotfixApproved events raised by the Governance contract.
type GovernanceHotfixApprovedIterator struct {
	Event *GovernanceHotfixApproved // Event containing the contract specifics and raw log

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
func (it *GovernanceHotfixApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixApproved)
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
		it.Event = new(GovernanceHotfixApproved)
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
func (it *GovernanceHotfixApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixApproved represents a HotfixApproved event raised by the Governance contract.
type GovernanceHotfixApproved struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHotfixApproved is a free log retrieval operation binding the contract event 0x36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f.
//
// Solidity: event HotfixApproved(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) FilterHotfixApproved(opts *bind.FilterOpts, hash [][32]byte) (*GovernanceHotfixApprovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixApproved", hashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixApprovedIterator{contract: _Governance.contract, event: "HotfixApproved", logs: logs, sub: sub}, nil
}

// WatchHotfixApproved is a free log subscription operation binding the contract event 0x36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f.
//
// Solidity: event HotfixApproved(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) WatchHotfixApproved(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixApproved, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixApproved", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixApproved)
				if err := _Governance.contract.UnpackLog(event, "HotfixApproved", log); err != nil {
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

// ParseHotfixApproved is a log parse operation binding the contract event 0x36bc158cba244a94dc9b8c08d327e8f7e3c2ab5f1925454c577527466f04851f.
//
// Solidity: event HotfixApproved(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) ParseHotfixApproved(log types.Log) (*GovernanceHotfixApproved, error) {
	event := new(GovernanceHotfixApproved)
	if err := _Governance.contract.UnpackLog(event, "HotfixApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceHotfixExecutedIterator is returned from FilterHotfixExecuted and is used to iterate over the raw logs and unpacked data for HotfixExecuted events raised by the Governance contract.
type GovernanceHotfixExecutedIterator struct {
	Event *GovernanceHotfixExecuted // Event containing the contract specifics and raw log

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
func (it *GovernanceHotfixExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixExecuted)
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
		it.Event = new(GovernanceHotfixExecuted)
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
func (it *GovernanceHotfixExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixExecuted represents a HotfixExecuted event raised by the Governance contract.
type GovernanceHotfixExecuted struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterHotfixExecuted is a free log retrieval operation binding the contract event 0x708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b63.
//
// Solidity: event HotfixExecuted(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) FilterHotfixExecuted(opts *bind.FilterOpts, hash [][32]byte) (*GovernanceHotfixExecutedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixExecuted", hashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixExecutedIterator{contract: _Governance.contract, event: "HotfixExecuted", logs: logs, sub: sub}, nil
}

// WatchHotfixExecuted is a free log subscription operation binding the contract event 0x708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b63.
//
// Solidity: event HotfixExecuted(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) WatchHotfixExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixExecuted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixExecuted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixExecuted)
				if err := _Governance.contract.UnpackLog(event, "HotfixExecuted", log); err != nil {
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

// ParseHotfixExecuted is a log parse operation binding the contract event 0x708a7934acb657a77a617b1fcd5f6d7d9ad592b72934841bff01acefd10f9b63.
//
// Solidity: event HotfixExecuted(bytes32 indexed hash)
func (_Governance *GovernanceFilterer) ParseHotfixExecuted(log types.Log) (*GovernanceHotfixExecuted, error) {
	event := new(GovernanceHotfixExecuted)
	if err := _Governance.contract.UnpackLog(event, "HotfixExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceHotfixPreparedIterator is returned from FilterHotfixPrepared and is used to iterate over the raw logs and unpacked data for HotfixPrepared events raised by the Governance contract.
type GovernanceHotfixPreparedIterator struct {
	Event *GovernanceHotfixPrepared // Event containing the contract specifics and raw log

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
func (it *GovernanceHotfixPreparedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixPrepared)
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
		it.Event = new(GovernanceHotfixPrepared)
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
func (it *GovernanceHotfixPreparedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixPreparedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixPrepared represents a HotfixPrepared event raised by the Governance contract.
type GovernanceHotfixPrepared struct {
	Hash  [32]byte
	Epoch *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterHotfixPrepared is a free log retrieval operation binding the contract event 0x6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e0.
//
// Solidity: event HotfixPrepared(bytes32 indexed hash, uint256 indexed epoch)
func (_Governance *GovernanceFilterer) FilterHotfixPrepared(opts *bind.FilterOpts, hash [][32]byte, epoch []*big.Int) (*GovernanceHotfixPreparedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixPrepared", hashRule, epochRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixPreparedIterator{contract: _Governance.contract, event: "HotfixPrepared", logs: logs, sub: sub}, nil
}

// WatchHotfixPrepared is a free log subscription operation binding the contract event 0x6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e0.
//
// Solidity: event HotfixPrepared(bytes32 indexed hash, uint256 indexed epoch)
func (_Governance *GovernanceFilterer) WatchHotfixPrepared(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixPrepared, hash [][32]byte, epoch []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var epochRule []interface{}
	for _, epochItem := range epoch {
		epochRule = append(epochRule, epochItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixPrepared", hashRule, epochRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixPrepared)
				if err := _Governance.contract.UnpackLog(event, "HotfixPrepared", log); err != nil {
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

// ParseHotfixPrepared is a log parse operation binding the contract event 0x6f184ec313435b3307a4fe59e2293381f08419a87214464c875a2a247e8af5e0.
//
// Solidity: event HotfixPrepared(bytes32 indexed hash, uint256 indexed epoch)
func (_Governance *GovernanceFilterer) ParseHotfixPrepared(log types.Log) (*GovernanceHotfixPrepared, error) {
	event := new(GovernanceHotfixPrepared)
	if err := _Governance.contract.UnpackLog(event, "HotfixPrepared", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceHotfixWhitelistedIterator is returned from FilterHotfixWhitelisted and is used to iterate over the raw logs and unpacked data for HotfixWhitelisted events raised by the Governance contract.
type GovernanceHotfixWhitelistedIterator struct {
	Event *GovernanceHotfixWhitelisted // Event containing the contract specifics and raw log

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
func (it *GovernanceHotfixWhitelistedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceHotfixWhitelisted)
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
		it.Event = new(GovernanceHotfixWhitelisted)
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
func (it *GovernanceHotfixWhitelistedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceHotfixWhitelistedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceHotfixWhitelisted represents a HotfixWhitelisted event raised by the Governance contract.
type GovernanceHotfixWhitelisted struct {
	Hash        [32]byte
	Whitelister common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterHotfixWhitelisted is a free log retrieval operation binding the contract event 0xf6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f.
//
// Solidity: event HotfixWhitelisted(bytes32 indexed hash, address whitelister)
func (_Governance *GovernanceFilterer) FilterHotfixWhitelisted(opts *bind.FilterOpts, hash [][32]byte) (*GovernanceHotfixWhitelistedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "HotfixWhitelisted", hashRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceHotfixWhitelistedIterator{contract: _Governance.contract, event: "HotfixWhitelisted", logs: logs, sub: sub}, nil
}

// WatchHotfixWhitelisted is a free log subscription operation binding the contract event 0xf6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f.
//
// Solidity: event HotfixWhitelisted(bytes32 indexed hash, address whitelister)
func (_Governance *GovernanceFilterer) WatchHotfixWhitelisted(opts *bind.WatchOpts, sink chan<- *GovernanceHotfixWhitelisted, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "HotfixWhitelisted", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceHotfixWhitelisted)
				if err := _Governance.contract.UnpackLog(event, "HotfixWhitelisted", log); err != nil {
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

// ParseHotfixWhitelisted is a log parse operation binding the contract event 0xf6d22d0b43a6753880b8f9511b82b86cd0fe349cd580bbe6a25b6dc063ef496f.
//
// Solidity: event HotfixWhitelisted(bytes32 indexed hash, address whitelister)
func (_Governance *GovernanceFilterer) ParseHotfixWhitelisted(log types.Log) (*GovernanceHotfixWhitelisted, error) {
	event := new(GovernanceHotfixWhitelisted)
	if err := _Governance.contract.UnpackLog(event, "HotfixWhitelisted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceMinDepositSetIterator is returned from FilterMinDepositSet and is used to iterate over the raw logs and unpacked data for MinDepositSet events raised by the Governance contract.
type GovernanceMinDepositSetIterator struct {
	Event *GovernanceMinDepositSet // Event containing the contract specifics and raw log

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
func (it *GovernanceMinDepositSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceMinDepositSet)
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
		it.Event = new(GovernanceMinDepositSet)
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
func (it *GovernanceMinDepositSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceMinDepositSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceMinDepositSet represents a MinDepositSet event raised by the Governance contract.
type GovernanceMinDepositSet struct {
	MinDeposit *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMinDepositSet is a free log retrieval operation binding the contract event 0xc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba81.
//
// Solidity: event MinDepositSet(uint256 minDeposit)
func (_Governance *GovernanceFilterer) FilterMinDepositSet(opts *bind.FilterOpts) (*GovernanceMinDepositSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "MinDepositSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceMinDepositSetIterator{contract: _Governance.contract, event: "MinDepositSet", logs: logs, sub: sub}, nil
}

// WatchMinDepositSet is a free log subscription operation binding the contract event 0xc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba81.
//
// Solidity: event MinDepositSet(uint256 minDeposit)
func (_Governance *GovernanceFilterer) WatchMinDepositSet(opts *bind.WatchOpts, sink chan<- *GovernanceMinDepositSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "MinDepositSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceMinDepositSet)
				if err := _Governance.contract.UnpackLog(event, "MinDepositSet", log); err != nil {
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

// ParseMinDepositSet is a log parse operation binding the contract event 0xc50a7f0bdf88c216b2541b0bdea26f22305460e39ffc672ec1a7501732c5ba81.
//
// Solidity: event MinDepositSet(uint256 minDeposit)
func (_Governance *GovernanceFilterer) ParseMinDepositSet(log types.Log) (*GovernanceMinDepositSet, error) {
	event := new(GovernanceMinDepositSet)
	if err := _Governance.contract.UnpackLog(event, "MinDepositSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Governance contract.
type GovernanceOwnershipTransferredIterator struct {
	Event *GovernanceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *GovernanceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceOwnershipTransferred)
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
		it.Event = new(GovernanceOwnershipTransferred)
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
func (it *GovernanceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceOwnershipTransferred represents a OwnershipTransferred event raised by the Governance contract.
type GovernanceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*GovernanceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceOwnershipTransferredIterator{contract: _Governance.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Governance *GovernanceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *GovernanceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceOwnershipTransferred)
				if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseOwnershipTransferred(log types.Log) (*GovernanceOwnershipTransferred, error) {
	event := new(GovernanceOwnershipTransferred)
	if err := _Governance.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceParticipationBaselineQuorumFactorSetIterator is returned from FilterParticipationBaselineQuorumFactorSet and is used to iterate over the raw logs and unpacked data for ParticipationBaselineQuorumFactorSet events raised by the Governance contract.
type GovernanceParticipationBaselineQuorumFactorSetIterator struct {
	Event *GovernanceParticipationBaselineQuorumFactorSet // Event containing the contract specifics and raw log

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
func (it *GovernanceParticipationBaselineQuorumFactorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationBaselineQuorumFactorSet)
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
		it.Event = new(GovernanceParticipationBaselineQuorumFactorSet)
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
func (it *GovernanceParticipationBaselineQuorumFactorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationBaselineQuorumFactorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationBaselineQuorumFactorSet represents a ParticipationBaselineQuorumFactorSet event raised by the Governance contract.
type GovernanceParticipationBaselineQuorumFactorSet struct {
	BaselineQuorumFactor *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterParticipationBaselineQuorumFactorSet is a free log retrieval operation binding the contract event 0xddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc.
//
// Solidity: event ParticipationBaselineQuorumFactorSet(uint256 baselineQuorumFactor)
func (_Governance *GovernanceFilterer) FilterParticipationBaselineQuorumFactorSet(opts *bind.FilterOpts) (*GovernanceParticipationBaselineQuorumFactorSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationBaselineQuorumFactorSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationBaselineQuorumFactorSetIterator{contract: _Governance.contract, event: "ParticipationBaselineQuorumFactorSet", logs: logs, sub: sub}, nil
}

// WatchParticipationBaselineQuorumFactorSet is a free log subscription operation binding the contract event 0xddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc.
//
// Solidity: event ParticipationBaselineQuorumFactorSet(uint256 baselineQuorumFactor)
func (_Governance *GovernanceFilterer) WatchParticipationBaselineQuorumFactorSet(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationBaselineQuorumFactorSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationBaselineQuorumFactorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationBaselineQuorumFactorSet)
				if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineQuorumFactorSet", log); err != nil {
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

// ParseParticipationBaselineQuorumFactorSet is a log parse operation binding the contract event 0xddfdbe55eaaa70fe2b8bc82a9b0734c25cabe7cb6f1457f9644019f0b5ff91fc.
//
// Solidity: event ParticipationBaselineQuorumFactorSet(uint256 baselineQuorumFactor)
func (_Governance *GovernanceFilterer) ParseParticipationBaselineQuorumFactorSet(log types.Log) (*GovernanceParticipationBaselineQuorumFactorSet, error) {
	event := new(GovernanceParticipationBaselineQuorumFactorSet)
	if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineQuorumFactorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceParticipationBaselineUpdateFactorSetIterator is returned from FilterParticipationBaselineUpdateFactorSet and is used to iterate over the raw logs and unpacked data for ParticipationBaselineUpdateFactorSet events raised by the Governance contract.
type GovernanceParticipationBaselineUpdateFactorSetIterator struct {
	Event *GovernanceParticipationBaselineUpdateFactorSet // Event containing the contract specifics and raw log

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
func (it *GovernanceParticipationBaselineUpdateFactorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationBaselineUpdateFactorSet)
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
		it.Event = new(GovernanceParticipationBaselineUpdateFactorSet)
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
func (it *GovernanceParticipationBaselineUpdateFactorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationBaselineUpdateFactorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationBaselineUpdateFactorSet represents a ParticipationBaselineUpdateFactorSet event raised by the Governance contract.
type GovernanceParticipationBaselineUpdateFactorSet struct {
	BaselineUpdateFactor *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterParticipationBaselineUpdateFactorSet is a free log retrieval operation binding the contract event 0x8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f.
//
// Solidity: event ParticipationBaselineUpdateFactorSet(uint256 baselineUpdateFactor)
func (_Governance *GovernanceFilterer) FilterParticipationBaselineUpdateFactorSet(opts *bind.FilterOpts) (*GovernanceParticipationBaselineUpdateFactorSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationBaselineUpdateFactorSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationBaselineUpdateFactorSetIterator{contract: _Governance.contract, event: "ParticipationBaselineUpdateFactorSet", logs: logs, sub: sub}, nil
}

// WatchParticipationBaselineUpdateFactorSet is a free log subscription operation binding the contract event 0x8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f.
//
// Solidity: event ParticipationBaselineUpdateFactorSet(uint256 baselineUpdateFactor)
func (_Governance *GovernanceFilterer) WatchParticipationBaselineUpdateFactorSet(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationBaselineUpdateFactorSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationBaselineUpdateFactorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationBaselineUpdateFactorSet)
				if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdateFactorSet", log); err != nil {
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

// ParseParticipationBaselineUpdateFactorSet is a log parse operation binding the contract event 0x8dedb4d995dd500718c7c5f6a077fba6153a7ee063da961d9fcab90ff528ac1f.
//
// Solidity: event ParticipationBaselineUpdateFactorSet(uint256 baselineUpdateFactor)
func (_Governance *GovernanceFilterer) ParseParticipationBaselineUpdateFactorSet(log types.Log) (*GovernanceParticipationBaselineUpdateFactorSet, error) {
	event := new(GovernanceParticipationBaselineUpdateFactorSet)
	if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdateFactorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceParticipationBaselineUpdatedIterator is returned from FilterParticipationBaselineUpdated and is used to iterate over the raw logs and unpacked data for ParticipationBaselineUpdated events raised by the Governance contract.
type GovernanceParticipationBaselineUpdatedIterator struct {
	Event *GovernanceParticipationBaselineUpdated // Event containing the contract specifics and raw log

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
func (it *GovernanceParticipationBaselineUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationBaselineUpdated)
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
		it.Event = new(GovernanceParticipationBaselineUpdated)
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
func (it *GovernanceParticipationBaselineUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationBaselineUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationBaselineUpdated represents a ParticipationBaselineUpdated event raised by the Governance contract.
type GovernanceParticipationBaselineUpdated struct {
	ParticipationBaseline *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterParticipationBaselineUpdated is a free log retrieval operation binding the contract event 0x51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c5.
//
// Solidity: event ParticipationBaselineUpdated(uint256 participationBaseline)
func (_Governance *GovernanceFilterer) FilterParticipationBaselineUpdated(opts *bind.FilterOpts) (*GovernanceParticipationBaselineUpdatedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationBaselineUpdated")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationBaselineUpdatedIterator{contract: _Governance.contract, event: "ParticipationBaselineUpdated", logs: logs, sub: sub}, nil
}

// WatchParticipationBaselineUpdated is a free log subscription operation binding the contract event 0x51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c5.
//
// Solidity: event ParticipationBaselineUpdated(uint256 participationBaseline)
func (_Governance *GovernanceFilterer) WatchParticipationBaselineUpdated(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationBaselineUpdated) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationBaselineUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationBaselineUpdated)
				if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdated", log); err != nil {
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

// ParseParticipationBaselineUpdated is a log parse operation binding the contract event 0x51131d2820f04a6b6edd20e22a07d5bf847e265a3906e85256fca7d6043417c5.
//
// Solidity: event ParticipationBaselineUpdated(uint256 participationBaseline)
func (_Governance *GovernanceFilterer) ParseParticipationBaselineUpdated(log types.Log) (*GovernanceParticipationBaselineUpdated, error) {
	event := new(GovernanceParticipationBaselineUpdated)
	if err := _Governance.contract.UnpackLog(event, "ParticipationBaselineUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceParticipationFloorSetIterator is returned from FilterParticipationFloorSet and is used to iterate over the raw logs and unpacked data for ParticipationFloorSet events raised by the Governance contract.
type GovernanceParticipationFloorSetIterator struct {
	Event *GovernanceParticipationFloorSet // Event containing the contract specifics and raw log

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
func (it *GovernanceParticipationFloorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceParticipationFloorSet)
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
		it.Event = new(GovernanceParticipationFloorSet)
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
func (it *GovernanceParticipationFloorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceParticipationFloorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceParticipationFloorSet represents a ParticipationFloorSet event raised by the Governance contract.
type GovernanceParticipationFloorSet struct {
	ParticipationFloor *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterParticipationFloorSet is a free log retrieval operation binding the contract event 0x122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b5.
//
// Solidity: event ParticipationFloorSet(uint256 participationFloor)
func (_Governance *GovernanceFilterer) FilterParticipationFloorSet(opts *bind.FilterOpts) (*GovernanceParticipationFloorSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ParticipationFloorSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceParticipationFloorSetIterator{contract: _Governance.contract, event: "ParticipationFloorSet", logs: logs, sub: sub}, nil
}

// WatchParticipationFloorSet is a free log subscription operation binding the contract event 0x122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b5.
//
// Solidity: event ParticipationFloorSet(uint256 participationFloor)
func (_Governance *GovernanceFilterer) WatchParticipationFloorSet(opts *bind.WatchOpts, sink chan<- *GovernanceParticipationFloorSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ParticipationFloorSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceParticipationFloorSet)
				if err := _Governance.contract.UnpackLog(event, "ParticipationFloorSet", log); err != nil {
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

// ParseParticipationFloorSet is a log parse operation binding the contract event 0x122a37b609e0f758b6a23c43506cb567017a4d18b21fa91312fb42b44975a5b5.
//
// Solidity: event ParticipationFloorSet(uint256 participationFloor)
func (_Governance *GovernanceFilterer) ParseParticipationFloorSet(log types.Log) (*GovernanceParticipationFloorSet, error) {
	event := new(GovernanceParticipationFloorSet)
	if err := _Governance.contract.UnpackLog(event, "ParticipationFloorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalApprovedIterator is returned from FilterProposalApproved and is used to iterate over the raw logs and unpacked data for ProposalApproved events raised by the Governance contract.
type GovernanceProposalApprovedIterator struct {
	Event *GovernanceProposalApproved // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalApproved)
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
		it.Event = new(GovernanceProposalApproved)
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
func (it *GovernanceProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalApproved represents a ProposalApproved event raised by the Governance contract.
type GovernanceProposalApproved struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalApproved(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalApprovedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalApproved", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalApprovedIterator{contract: _Governance.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *GovernanceProposalApproved, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalApproved", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalApproved)
				if err := _Governance.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
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

// ParseProposalApproved is a log parse operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalApproved(log types.Log) (*GovernanceProposalApproved, error) {
	event := new(GovernanceProposalApproved)
	if err := _Governance.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalDequeuedIterator is returned from FilterProposalDequeued and is used to iterate over the raw logs and unpacked data for ProposalDequeued events raised by the Governance contract.
type GovernanceProposalDequeuedIterator struct {
	Event *GovernanceProposalDequeued // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalDequeuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalDequeued)
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
		it.Event = new(GovernanceProposalDequeued)
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
func (it *GovernanceProposalDequeuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalDequeuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalDequeued represents a ProposalDequeued event raised by the Governance contract.
type GovernanceProposalDequeued struct {
	ProposalId *big.Int
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalDequeued is a free log retrieval operation binding the contract event 0x3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543.
//
// Solidity: event ProposalDequeued(uint256 indexed proposalId, uint256 timestamp)
func (_Governance *GovernanceFilterer) FilterProposalDequeued(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalDequeuedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalDequeued", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalDequeuedIterator{contract: _Governance.contract, event: "ProposalDequeued", logs: logs, sub: sub}, nil
}

// WatchProposalDequeued is a free log subscription operation binding the contract event 0x3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543.
//
// Solidity: event ProposalDequeued(uint256 indexed proposalId, uint256 timestamp)
func (_Governance *GovernanceFilterer) WatchProposalDequeued(opts *bind.WatchOpts, sink chan<- *GovernanceProposalDequeued, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalDequeued", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalDequeued)
				if err := _Governance.contract.UnpackLog(event, "ProposalDequeued", log); err != nil {
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

// ParseProposalDequeued is a log parse operation binding the contract event 0x3e069fb74dcf5fbc07740b0d40d7f7fc48e9c0ca5dc3d19eb34d2e05d74c5543.
//
// Solidity: event ProposalDequeued(uint256 indexed proposalId, uint256 timestamp)
func (_Governance *GovernanceFilterer) ParseProposalDequeued(log types.Log) (*GovernanceProposalDequeued, error) {
	event := new(GovernanceProposalDequeued)
	if err := _Governance.contract.UnpackLog(event, "ProposalDequeued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Governance contract.
type GovernanceProposalExecutedIterator struct {
	Event *GovernanceProposalExecuted // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExecuted)
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
		it.Event = new(GovernanceProposalExecuted)
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
func (it *GovernanceProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExecuted represents a ProposalExecuted event raised by the Governance contract.
type GovernanceProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalExecuted(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalExecutedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExecutedIterator{contract: _Governance.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExecuted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExecuted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExecuted)
				if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalExecuted(log types.Log) (*GovernanceProposalExecuted, error) {
	event := new(GovernanceProposalExecuted)
	if err := _Governance.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalExpiredIterator is returned from FilterProposalExpired and is used to iterate over the raw logs and unpacked data for ProposalExpired events raised by the Governance contract.
type GovernanceProposalExpiredIterator struct {
	Event *GovernanceProposalExpired // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalExpired)
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
		it.Event = new(GovernanceProposalExpired)
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
func (it *GovernanceProposalExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalExpired represents a ProposalExpired event raised by the Governance contract.
type GovernanceProposalExpired struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExpired is a free log retrieval operation binding the contract event 0x88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a2643.
//
// Solidity: event ProposalExpired(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) FilterProposalExpired(opts *bind.FilterOpts, proposalId []*big.Int) (*GovernanceProposalExpiredIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalExpired", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalExpiredIterator{contract: _Governance.contract, event: "ProposalExpired", logs: logs, sub: sub}, nil
}

// WatchProposalExpired is a free log subscription operation binding the contract event 0x88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a2643.
//
// Solidity: event ProposalExpired(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) WatchProposalExpired(opts *bind.WatchOpts, sink chan<- *GovernanceProposalExpired, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalExpired", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalExpired)
				if err := _Governance.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
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

// ParseProposalExpired is a log parse operation binding the contract event 0x88e53c486703527139dfc8d97a1e559d9bd93d3f9d52cda4e06564111e7a2643.
//
// Solidity: event ProposalExpired(uint256 indexed proposalId)
func (_Governance *GovernanceFilterer) ParseProposalExpired(log types.Log) (*GovernanceProposalExpired, error) {
	event := new(GovernanceProposalExpired)
	if err := _Governance.contract.UnpackLog(event, "ProposalExpired", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the Governance contract.
type GovernanceProposalQueuedIterator struct {
	Event *GovernanceProposalQueued // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalQueued)
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
		it.Event = new(GovernanceProposalQueued)
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
func (it *GovernanceProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalQueued represents a ProposalQueued event raised by the Governance contract.
type GovernanceProposalQueued struct {
	ProposalId       *big.Int
	Proposer         common.Address
	TransactionCount *big.Int
	Deposit          *big.Int
	Timestamp        *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address indexed proposer, uint256 transactionCount, uint256 deposit, uint256 timestamp)
func (_Governance *GovernanceFilterer) FilterProposalQueued(opts *bind.FilterOpts, proposalId []*big.Int, proposer []common.Address) (*GovernanceProposalQueuedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalQueued", proposalIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalQueuedIterator{contract: _Governance.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address indexed proposer, uint256 transactionCount, uint256 deposit, uint256 timestamp)
func (_Governance *GovernanceFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GovernanceProposalQueued, proposalId []*big.Int, proposer []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalQueued", proposalIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalQueued)
				if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
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

// ParseProposalQueued is a log parse operation binding the contract event 0x1bfe527f3548d9258c2512b6689f0acfccdd0557d80a53845db25fc57e93d8fe.
//
// Solidity: event ProposalQueued(uint256 indexed proposalId, address indexed proposer, uint256 transactionCount, uint256 deposit, uint256 timestamp)
func (_Governance *GovernanceFilterer) ParseProposalQueued(log types.Log) (*GovernanceProposalQueued, error) {
	event := new(GovernanceProposalQueued)
	if err := _Governance.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalUpvoteRevokedIterator is returned from FilterProposalUpvoteRevoked and is used to iterate over the raw logs and unpacked data for ProposalUpvoteRevoked events raised by the Governance contract.
type GovernanceProposalUpvoteRevokedIterator struct {
	Event *GovernanceProposalUpvoteRevoked // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalUpvoteRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalUpvoteRevoked)
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
		it.Event = new(GovernanceProposalUpvoteRevoked)
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
func (it *GovernanceProposalUpvoteRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalUpvoteRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalUpvoteRevoked represents a ProposalUpvoteRevoked event raised by the Governance contract.
type GovernanceProposalUpvoteRevoked struct {
	ProposalId     *big.Int
	Account        common.Address
	RevokedUpvotes *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterProposalUpvoteRevoked is a free log retrieval operation binding the contract event 0x7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114.
//
// Solidity: event ProposalUpvoteRevoked(uint256 indexed proposalId, address indexed account, uint256 revokedUpvotes)
func (_Governance *GovernanceFilterer) FilterProposalUpvoteRevoked(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalUpvoteRevokedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalUpvoteRevoked", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalUpvoteRevokedIterator{contract: _Governance.contract, event: "ProposalUpvoteRevoked", logs: logs, sub: sub}, nil
}

// WatchProposalUpvoteRevoked is a free log subscription operation binding the contract event 0x7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114.
//
// Solidity: event ProposalUpvoteRevoked(uint256 indexed proposalId, address indexed account, uint256 revokedUpvotes)
func (_Governance *GovernanceFilterer) WatchProposalUpvoteRevoked(opts *bind.WatchOpts, sink chan<- *GovernanceProposalUpvoteRevoked, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalUpvoteRevoked", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalUpvoteRevoked)
				if err := _Governance.contract.UnpackLog(event, "ProposalUpvoteRevoked", log); err != nil {
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

// ParseProposalUpvoteRevoked is a log parse operation binding the contract event 0x7dc46237a819c9171a9c037ec98928e563892905c4d23373ca0f3f500f4ed114.
//
// Solidity: event ProposalUpvoteRevoked(uint256 indexed proposalId, address indexed account, uint256 revokedUpvotes)
func (_Governance *GovernanceFilterer) ParseProposalUpvoteRevoked(log types.Log) (*GovernanceProposalUpvoteRevoked, error) {
	event := new(GovernanceProposalUpvoteRevoked)
	if err := _Governance.contract.UnpackLog(event, "ProposalUpvoteRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalUpvotedIterator is returned from FilterProposalUpvoted and is used to iterate over the raw logs and unpacked data for ProposalUpvoted events raised by the Governance contract.
type GovernanceProposalUpvotedIterator struct {
	Event *GovernanceProposalUpvoted // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalUpvotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalUpvoted)
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
		it.Event = new(GovernanceProposalUpvoted)
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
func (it *GovernanceProposalUpvotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalUpvotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalUpvoted represents a ProposalUpvoted event raised by the Governance contract.
type GovernanceProposalUpvoted struct {
	ProposalId *big.Int
	Account    common.Address
	Upvotes    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalUpvoted is a free log retrieval operation binding the contract event 0xd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d62.
//
// Solidity: event ProposalUpvoted(uint256 indexed proposalId, address indexed account, uint256 upvotes)
func (_Governance *GovernanceFilterer) FilterProposalUpvoted(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalUpvotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalUpvoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalUpvotedIterator{contract: _Governance.contract, event: "ProposalUpvoted", logs: logs, sub: sub}, nil
}

// WatchProposalUpvoted is a free log subscription operation binding the contract event 0xd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d62.
//
// Solidity: event ProposalUpvoted(uint256 indexed proposalId, address indexed account, uint256 upvotes)
func (_Governance *GovernanceFilterer) WatchProposalUpvoted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalUpvoted, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalUpvoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalUpvoted)
				if err := _Governance.contract.UnpackLog(event, "ProposalUpvoted", log); err != nil {
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

// ParseProposalUpvoted is a log parse operation binding the contract event 0xd19965d25ef670a1e322fbf05475924b7b12d81fd6b96ab718b261782efb3d62.
//
// Solidity: event ProposalUpvoted(uint256 indexed proposalId, address indexed account, uint256 upvotes)
func (_Governance *GovernanceFilterer) ParseProposalUpvoted(log types.Log) (*GovernanceProposalUpvoted, error) {
	event := new(GovernanceProposalUpvoted)
	if err := _Governance.contract.UnpackLog(event, "ProposalUpvoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalVoteRevokedIterator is returned from FilterProposalVoteRevoked and is used to iterate over the raw logs and unpacked data for ProposalVoteRevoked events raised by the Governance contract.
type GovernanceProposalVoteRevokedIterator struct {
	Event *GovernanceProposalVoteRevoked // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalVoteRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalVoteRevoked)
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
		it.Event = new(GovernanceProposalVoteRevoked)
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
func (it *GovernanceProposalVoteRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalVoteRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalVoteRevoked represents a ProposalVoteRevoked event raised by the Governance contract.
type GovernanceProposalVoteRevoked struct {
	ProposalId *big.Int
	Account    common.Address
	Value      *big.Int
	Weight     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVoteRevoked is a free log retrieval operation binding the contract event 0xb59283e3d5436f05576bddef72ddbfb6344c216ed6ea6d7ced2e9bbb94c661ab.
//
// Solidity: event ProposalVoteRevoked(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) FilterProposalVoteRevoked(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalVoteRevokedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalVoteRevoked", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalVoteRevokedIterator{contract: _Governance.contract, event: "ProposalVoteRevoked", logs: logs, sub: sub}, nil
}

// WatchProposalVoteRevoked is a free log subscription operation binding the contract event 0xb59283e3d5436f05576bddef72ddbfb6344c216ed6ea6d7ced2e9bbb94c661ab.
//
// Solidity: event ProposalVoteRevoked(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) WatchProposalVoteRevoked(opts *bind.WatchOpts, sink chan<- *GovernanceProposalVoteRevoked, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalVoteRevoked", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalVoteRevoked)
				if err := _Governance.contract.UnpackLog(event, "ProposalVoteRevoked", log); err != nil {
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

// ParseProposalVoteRevoked is a log parse operation binding the contract event 0xb59283e3d5436f05576bddef72ddbfb6344c216ed6ea6d7ced2e9bbb94c661ab.
//
// Solidity: event ProposalVoteRevoked(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) ParseProposalVoteRevoked(log types.Log) (*GovernanceProposalVoteRevoked, error) {
	event := new(GovernanceProposalVoteRevoked)
	if err := _Governance.contract.UnpackLog(event, "ProposalVoteRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalVoteRevokedV2Iterator is returned from FilterProposalVoteRevokedV2 and is used to iterate over the raw logs and unpacked data for ProposalVoteRevokedV2 events raised by the Governance contract.
type GovernanceProposalVoteRevokedV2Iterator struct {
	Event *GovernanceProposalVoteRevokedV2 // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalVoteRevokedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalVoteRevokedV2)
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
		it.Event = new(GovernanceProposalVoteRevokedV2)
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
func (it *GovernanceProposalVoteRevokedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalVoteRevokedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalVoteRevokedV2 represents a ProposalVoteRevokedV2 event raised by the Governance contract.
type GovernanceProposalVoteRevokedV2 struct {
	ProposalId   *big.Int
	Account      common.Address
	YesVotes     *big.Int
	NoVotes      *big.Int
	AbstainVotes *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalVoteRevokedV2 is a free log retrieval operation binding the contract event 0x6791653c96b4863b3768c664e9a03e1094ae334ba9d35862627ceeebd1abcc1f.
//
// Solidity: event ProposalVoteRevokedV2(uint256 indexed proposalId, address indexed account, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes)
func (_Governance *GovernanceFilterer) FilterProposalVoteRevokedV2(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalVoteRevokedV2Iterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalVoteRevokedV2", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalVoteRevokedV2Iterator{contract: _Governance.contract, event: "ProposalVoteRevokedV2", logs: logs, sub: sub}, nil
}

// WatchProposalVoteRevokedV2 is a free log subscription operation binding the contract event 0x6791653c96b4863b3768c664e9a03e1094ae334ba9d35862627ceeebd1abcc1f.
//
// Solidity: event ProposalVoteRevokedV2(uint256 indexed proposalId, address indexed account, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes)
func (_Governance *GovernanceFilterer) WatchProposalVoteRevokedV2(opts *bind.WatchOpts, sink chan<- *GovernanceProposalVoteRevokedV2, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalVoteRevokedV2", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalVoteRevokedV2)
				if err := _Governance.contract.UnpackLog(event, "ProposalVoteRevokedV2", log); err != nil {
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

// ParseProposalVoteRevokedV2 is a log parse operation binding the contract event 0x6791653c96b4863b3768c664e9a03e1094ae334ba9d35862627ceeebd1abcc1f.
//
// Solidity: event ProposalVoteRevokedV2(uint256 indexed proposalId, address indexed account, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes)
func (_Governance *GovernanceFilterer) ParseProposalVoteRevokedV2(log types.Log) (*GovernanceProposalVoteRevokedV2, error) {
	event := new(GovernanceProposalVoteRevokedV2)
	if err := _Governance.contract.UnpackLog(event, "ProposalVoteRevokedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalVotedIterator is returned from FilterProposalVoted and is used to iterate over the raw logs and unpacked data for ProposalVoted events raised by the Governance contract.
type GovernanceProposalVotedIterator struct {
	Event *GovernanceProposalVoted // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalVotedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalVoted)
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
		it.Event = new(GovernanceProposalVoted)
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
func (it *GovernanceProposalVotedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalVotedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalVoted represents a ProposalVoted event raised by the Governance contract.
type GovernanceProposalVoted struct {
	ProposalId *big.Int
	Account    common.Address
	Value      *big.Int
	Weight     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalVoted is a free log retrieval operation binding the contract event 0xf3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) FilterProposalVoted(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalVotedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalVoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalVotedIterator{contract: _Governance.contract, event: "ProposalVoted", logs: logs, sub: sub}, nil
}

// WatchProposalVoted is a free log subscription operation binding the contract event 0xf3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) WatchProposalVoted(opts *bind.WatchOpts, sink chan<- *GovernanceProposalVoted, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalVoted", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalVoted)
				if err := _Governance.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
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

// ParseProposalVoted is a log parse operation binding the contract event 0xf3709dc32cf1356da6b8a12a5be1401aeb00989556be7b16ae566e65fef7a9df.
//
// Solidity: event ProposalVoted(uint256 indexed proposalId, address indexed account, uint256 value, uint256 weight)
func (_Governance *GovernanceFilterer) ParseProposalVoted(log types.Log) (*GovernanceProposalVoted, error) {
	event := new(GovernanceProposalVoted)
	if err := _Governance.contract.UnpackLog(event, "ProposalVoted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceProposalVotedV2Iterator is returned from FilterProposalVotedV2 and is used to iterate over the raw logs and unpacked data for ProposalVotedV2 events raised by the Governance contract.
type GovernanceProposalVotedV2Iterator struct {
	Event *GovernanceProposalVotedV2 // Event containing the contract specifics and raw log

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
func (it *GovernanceProposalVotedV2Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceProposalVotedV2)
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
		it.Event = new(GovernanceProposalVotedV2)
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
func (it *GovernanceProposalVotedV2Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceProposalVotedV2Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceProposalVotedV2 represents a ProposalVotedV2 event raised by the Governance contract.
type GovernanceProposalVotedV2 struct {
	ProposalId   *big.Int
	Account      common.Address
	YesVotes     *big.Int
	NoVotes      *big.Int
	AbstainVotes *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalVotedV2 is a free log retrieval operation binding the contract event 0x683ebddc94b5b0a7dae3d1b6c168ad05684fcfd831b24ecb5ea9ecdf5524d028.
//
// Solidity: event ProposalVotedV2(uint256 indexed proposalId, address indexed account, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes)
func (_Governance *GovernanceFilterer) FilterProposalVotedV2(opts *bind.FilterOpts, proposalId []*big.Int, account []common.Address) (*GovernanceProposalVotedV2Iterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ProposalVotedV2", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceProposalVotedV2Iterator{contract: _Governance.contract, event: "ProposalVotedV2", logs: logs, sub: sub}, nil
}

// WatchProposalVotedV2 is a free log subscription operation binding the contract event 0x683ebddc94b5b0a7dae3d1b6c168ad05684fcfd831b24ecb5ea9ecdf5524d028.
//
// Solidity: event ProposalVotedV2(uint256 indexed proposalId, address indexed account, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes)
func (_Governance *GovernanceFilterer) WatchProposalVotedV2(opts *bind.WatchOpts, sink chan<- *GovernanceProposalVotedV2, proposalId []*big.Int, account []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ProposalVotedV2", proposalIdRule, accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceProposalVotedV2)
				if err := _Governance.contract.UnpackLog(event, "ProposalVotedV2", log); err != nil {
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

// ParseProposalVotedV2 is a log parse operation binding the contract event 0x683ebddc94b5b0a7dae3d1b6c168ad05684fcfd831b24ecb5ea9ecdf5524d028.
//
// Solidity: event ProposalVotedV2(uint256 indexed proposalId, address indexed account, uint256 yesVotes, uint256 noVotes, uint256 abstainVotes)
func (_Governance *GovernanceFilterer) ParseProposalVotedV2(log types.Log) (*GovernanceProposalVotedV2, error) {
	event := new(GovernanceProposalVotedV2)
	if err := _Governance.contract.UnpackLog(event, "ProposalVotedV2", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceQueueExpirySetIterator is returned from FilterQueueExpirySet and is used to iterate over the raw logs and unpacked data for QueueExpirySet events raised by the Governance contract.
type GovernanceQueueExpirySetIterator struct {
	Event *GovernanceQueueExpirySet // Event containing the contract specifics and raw log

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
func (it *GovernanceQueueExpirySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceQueueExpirySet)
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
		it.Event = new(GovernanceQueueExpirySet)
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
func (it *GovernanceQueueExpirySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceQueueExpirySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceQueueExpirySet represents a QueueExpirySet event raised by the Governance contract.
type GovernanceQueueExpirySet struct {
	QueueExpiry *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterQueueExpirySet is a free log retrieval operation binding the contract event 0x4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d7.
//
// Solidity: event QueueExpirySet(uint256 queueExpiry)
func (_Governance *GovernanceFilterer) FilterQueueExpirySet(opts *bind.FilterOpts) (*GovernanceQueueExpirySetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "QueueExpirySet")
	if err != nil {
		return nil, err
	}
	return &GovernanceQueueExpirySetIterator{contract: _Governance.contract, event: "QueueExpirySet", logs: logs, sub: sub}, nil
}

// WatchQueueExpirySet is a free log subscription operation binding the contract event 0x4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d7.
//
// Solidity: event QueueExpirySet(uint256 queueExpiry)
func (_Governance *GovernanceFilterer) WatchQueueExpirySet(opts *bind.WatchOpts, sink chan<- *GovernanceQueueExpirySet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "QueueExpirySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceQueueExpirySet)
				if err := _Governance.contract.UnpackLog(event, "QueueExpirySet", log); err != nil {
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

// ParseQueueExpirySet is a log parse operation binding the contract event 0x4ecbf0bb0701615e2d6f9b0a0996056c959fe359ce76aa7ce06c5f1d57dae4d7.
//
// Solidity: event QueueExpirySet(uint256 queueExpiry)
func (_Governance *GovernanceFilterer) ParseQueueExpirySet(log types.Log) (*GovernanceQueueExpirySet, error) {
	event := new(GovernanceQueueExpirySet)
	if err := _Governance.contract.UnpackLog(event, "QueueExpirySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceReferendumStageDurationSetIterator is returned from FilterReferendumStageDurationSet and is used to iterate over the raw logs and unpacked data for ReferendumStageDurationSet events raised by the Governance contract.
type GovernanceReferendumStageDurationSetIterator struct {
	Event *GovernanceReferendumStageDurationSet // Event containing the contract specifics and raw log

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
func (it *GovernanceReferendumStageDurationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceReferendumStageDurationSet)
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
		it.Event = new(GovernanceReferendumStageDurationSet)
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
func (it *GovernanceReferendumStageDurationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceReferendumStageDurationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceReferendumStageDurationSet represents a ReferendumStageDurationSet event raised by the Governance contract.
type GovernanceReferendumStageDurationSet struct {
	ReferendumStageDuration *big.Int
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterReferendumStageDurationSet is a free log retrieval operation binding the contract event 0x90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed6.
//
// Solidity: event ReferendumStageDurationSet(uint256 referendumStageDuration)
func (_Governance *GovernanceFilterer) FilterReferendumStageDurationSet(opts *bind.FilterOpts) (*GovernanceReferendumStageDurationSetIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "ReferendumStageDurationSet")
	if err != nil {
		return nil, err
	}
	return &GovernanceReferendumStageDurationSetIterator{contract: _Governance.contract, event: "ReferendumStageDurationSet", logs: logs, sub: sub}, nil
}

// WatchReferendumStageDurationSet is a free log subscription operation binding the contract event 0x90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed6.
//
// Solidity: event ReferendumStageDurationSet(uint256 referendumStageDuration)
func (_Governance *GovernanceFilterer) WatchReferendumStageDurationSet(opts *bind.WatchOpts, sink chan<- *GovernanceReferendumStageDurationSet) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "ReferendumStageDurationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceReferendumStageDurationSet)
				if err := _Governance.contract.UnpackLog(event, "ReferendumStageDurationSet", log); err != nil {
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

// ParseReferendumStageDurationSet is a log parse operation binding the contract event 0x90290eb9b27055e686a69fb810bada5381e544d07b8270021da2d355a6c96ed6.
//
// Solidity: event ReferendumStageDurationSet(uint256 referendumStageDuration)
func (_Governance *GovernanceFilterer) ParseReferendumStageDurationSet(log types.Log) (*GovernanceReferendumStageDurationSet, error) {
	event := new(GovernanceReferendumStageDurationSet)
	if err := _Governance.contract.UnpackLog(event, "ReferendumStageDurationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GovernanceRegistrySetIterator is returned from FilterRegistrySet and is used to iterate over the raw logs and unpacked data for RegistrySet events raised by the Governance contract.
type GovernanceRegistrySetIterator struct {
	Event *GovernanceRegistrySet // Event containing the contract specifics and raw log

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
func (it *GovernanceRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceRegistrySet)
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
		it.Event = new(GovernanceRegistrySet)
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
func (it *GovernanceRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceRegistrySet represents a RegistrySet event raised by the Governance contract.
type GovernanceRegistrySet struct {
	RegistryAddress common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterRegistrySet is a free log retrieval operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Governance *GovernanceFilterer) FilterRegistrySet(opts *bind.FilterOpts, registryAddress []common.Address) (*GovernanceRegistrySetIterator, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Governance.contract.FilterLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return &GovernanceRegistrySetIterator{contract: _Governance.contract, event: "RegistrySet", logs: logs, sub: sub}, nil
}

// WatchRegistrySet is a free log subscription operation binding the contract event 0x27fe5f0c1c3b1ed427cc63d0f05759ffdecf9aec9e18d31ef366fc8a6cb5dc3b.
//
// Solidity: event RegistrySet(address indexed registryAddress)
func (_Governance *GovernanceFilterer) WatchRegistrySet(opts *bind.WatchOpts, sink chan<- *GovernanceRegistrySet, registryAddress []common.Address) (event.Subscription, error) {

	var registryAddressRule []interface{}
	for _, registryAddressItem := range registryAddress {
		registryAddressRule = append(registryAddressRule, registryAddressItem)
	}

	logs, sub, err := _Governance.contract.WatchLogs(opts, "RegistrySet", registryAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceRegistrySet)
				if err := _Governance.contract.UnpackLog(event, "RegistrySet", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseRegistrySet(log types.Log) (*GovernanceRegistrySet, error) {
	event := new(GovernanceRegistrySet)
	if err := _Governance.contract.UnpackLog(event, "RegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
