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

package airgap

import (
	"fmt"
	"math/big"

	"github.com/celo-org/kliento/registry"
	"github.com/ethereum/go-ethereum/common"
)

var ReleaseGold = "ReleaseGold"

// values taken from contract method names for ABI usage
var (
	// Accounts
	CreateAccount              = registerMethod(registry.AccountsContractID.String(), "createAccount", nil)
	AuthorizeVoteSigner        = registerMethod(registry.AccountsContractID.String(), "authorizeVoteSigner", []argParser{addressParser, bytesParser})
	AuthorizeAttestationSigner = registerMethod(registry.AccountsContractID.String(), "authorizeAttestationSigner", []argParser{addressParser, bytesParser})
	AuthorizeValidatorSigner   = registerMethod(registry.AccountsContractID.String(), "authorizeValidatorSigner", []argParser{addressParser, bytesParser})
	IsAccount                  = registerMethod(registry.AccountsContractID.String(), "isAccount", []argParser{addressParser})
	GetVoteSigner              = registerMethod(registry.AccountsContractID.String(), "getVoteSigner", []argParser{addressParser})

	// Locked Gold
	LockGold     = registerMethod(registry.LockedGoldContractID.String(), "lock", nil)
	UnlockGold   = registerMethod(registry.LockedGoldContractID.String(), "unlock", []argParser{bigIntParser})
	RelockGold   = registerMethod(registry.LockedGoldContractID.String(), "relock", []argParser{bigIntParser, bigIntParser})
	WithdrawGold = registerMethod(registry.LockedGoldContractID.String(), "withdraw", []argParser{bigIntParser})

	// Election
	Vote                            = registerMethod(registry.ElectionContractID.String(), "vote", []argParser{addressParser, bigIntParser})
	ActivateVotes                   = registerMethod(registry.ElectionContractID.String(), "activate", []argParser{addressParser})
	RevokePendingVotes              = registerMethod(registry.ElectionContractID.String(), "revokePending", []argParser{addressParser, addressParser, bigIntParser})
	RevokeActiveVotes               = registerMethod(registry.ElectionContractID.String(), "revokeActive", []argParser{addressParser, addressParser, bigIntParser})
	CanReceiveVotes                 = registerMethod(registry.ElectionContractID.String(), "canReceiveVotes", []argParser{addressParser, bigIntParser})
	GetEpochNumber                  = registerMethod(registry.ElectionContractID.String(), "getEpochNumber", []argParser{})
	GetEpochSize                    = registerMethod(registry.ElectionContractID.String(), "getEpochSize", []argParser{})
	GetEpochNumberOfBlock           = registerMethod(registry.ElectionContractID.String(), "getEpochNumberOfBlock", []argParser{bigIntParser})
	GetGroupsVotedForByAccount      = registerMethod(registry.ElectionContractID.String(), "getGroupsVotedForByAccount", []argParser{addressParser})
	GetActiveVotesForGroup          = registerMethod(registry.ElectionContractID.String(), "getActiveVotesForGroup", []argParser{addressParser})
	GetActiveVotesForGroupByAccount = registerMethod(registry.ElectionContractID.String(), "getActiveVotesForGroupByAccount", []argParser{addressParser, addressParser})

	// ReleaseGold
	ReleaseGoldWithdraw = registerMethod(ReleaseGold, "withdraw", []argParser{bigIntParser})

	// Proxy
	ReleaseGoldCreateAccount              = registerMethod(ReleaseGold, "createAccount", nil)
	ReleaseGoldLockGold                   = registerMethod(ReleaseGold, "lockGold", []argParser{bigIntParser})
	ReleaseGoldUnlockGold                 = registerMethod(ReleaseGold, "unlockGold", []argParser{bigIntParser})
	ReleaseGoldRelockGold                 = registerMethod(ReleaseGold, "relockGold", []argParser{bigIntParser, bigIntParser})
	ReleaseGoldWithdrawGold               = registerMethod(ReleaseGold, "withdrawLockedGold", []argParser{bigIntParser})
	ReleaseGoldAuthorizeVoteSigner        = registerMethod(ReleaseGold, "authorizeVoteSigner", []argParser{addressParser, bytesParser})
	ReleaseGoldAuthorizeAttestationSigner = registerMethod(ReleaseGold, "authorizeAttestationSigner", []argParser{addressParser, bytesParser})
	ReleaseGoldAuthorizeValidatorSigner   = registerMethod(ReleaseGold, "authorizeValidatorSigner", []argParser{addressParser, bytesParser})
	ReleaseGoldRevokePendingVotes         = registerMethod(ReleaseGold, "revokePending", []argParser{addressParser, addressParser, bigIntParser})
	ReleaseGoldRevokeActiveVotes          = registerMethod(ReleaseGold, "revokeActive", []argParser{addressParser, addressParser, bigIntParser})
	ReleaseGoldMaxDistribution            = registerMethod(ReleaseGold, "maxDistribution", nil)
	ReleaseGoldTotalWithdrawn             = registerMethod(ReleaseGold, "totalWithdrawn", nil)

	// StableToken
	StableTokenBalanceOf = registerMethod(registry.StableTokenContractID.String(), "balanceOf", []argParser{addressParser})
	StableTokenTransfer  = registerMethod(registry.StableTokenContractID.String(), "transfer", []argParser{addressParser, bigIntParser})
)

// Represents a CeloMethod that can be called with the AirgapClient
// DO NOT CREATE them, instead use `MethodFromString()`
type CeloMethod struct {
	// Name of the abi method
	Name string
	// Registry id of contract where the method is defined
	Contract string

	argParsers []argParser
}

func (cm *CeloMethod) String() string { return fmt.Sprintf("%s.%s", cm.Contract, cm.Name) }

func (cm *CeloMethod) SerializeArguments(args ...interface{}) ([]interface{}, error) {
	if len(args) != len(cm.argParsers) {
		return nil, fmt.Errorf("Received %d args; expected %d", len(args), len(cm.argParsers))
	}
	out := make([]interface{}, len(args))

	for i, arg := range args {
		switch v := arg.(type) {
		case common.Address:
			out[i] = v.Hex()
		case *big.Int:
			out[i] = v.String()
		case []byte:
			out[i] = common.Bytes2Hex(v)
		default:
			out[i] = v
		}
	}
	return out, nil
}

func (cm *CeloMethod) DeserializeArguments(values ...interface{}) ([]interface{}, error) {
	if len(values) != len(cm.argParsers) {
		return nil, fmt.Errorf("Received %d args; expected %d", len(values), len(cm.argParsers))
	}
	parsedArgs := make([]interface{}, len(cm.argParsers))

	var err error
	for i, value := range values {
		parsedArgs[i], err = cm.argParsers[i](value)
		if err != nil {
			return nil, fmt.Errorf("bad argument: idx=%d error=%w", i, err)
		}
	}

	return parsedArgs, nil
}
