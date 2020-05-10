package airgap

import (
	"fmt"
	"math/big"

	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/ethereum/go-ethereum/common"
)

var ReleaseGold = "ReleaseGold"

// values taken from contract method names for ABI usage
var (
	// Accounts
	CreateAccount              = registerMethod(wrapper.AccountsRegistryId.String(), "createAccount", nil)
	AuthorizeVoteSigner        = registerMethod(wrapper.AccountsRegistryId.String(), "authorizeVoteSigner", []argParser{addressParser, bytesParser})
	AuthorizeAttestationSigner = registerMethod(wrapper.AccountsRegistryId.String(), "authorizeAttestationSigner", []argParser{addressParser, bytesParser})
	AuthorizeValidatorSigner   = registerMethod(wrapper.AccountsRegistryId.String(), "authorizeValidatorSigner", []argParser{addressParser, bytesParser})

	// Locked Gold
	LockGold     = registerMethod(wrapper.LockedGoldRegistryId.String(), "lock", nil)
	UnlockGold   = registerMethod(wrapper.LockedGoldRegistryId.String(), "unlock", []argParser{bigIntParser})
	RelockGold   = registerMethod(wrapper.LockedGoldRegistryId.String(), "relock", []argParser{bigIntParser, bigIntParser})
	WithdrawGold = registerMethod(wrapper.LockedGoldRegistryId.String(), "withdraw", []argParser{bigIntParser})

	// Election
	Vote               = registerMethod(wrapper.ElectionRegistryId.String(), "vote", []argParser{addressParser, bigIntParser})
	ActivateVotes      = registerMethod(wrapper.ElectionRegistryId.String(), "activate", []argParser{addressParser, addressParser})
	RevokePendingVotes = registerMethod(wrapper.ElectionRegistryId.String(), "revokePending", []argParser{addressParser, addressParser, bigIntParser})
	RevokeActiveVotes  = registerMethod(wrapper.ElectionRegistryId.String(), "revokeActive", []argParser{addressParser, addressParser, bigIntParser})

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

func (cm *CeloMethod) CreateTxArgs(to *common.Address, from common.Address, value *big.Int, args ...interface{}) (*TxArgs, error) {
	parsedArgs, err := cm.SerializeArguments(args...)
	if err != nil {
		return nil, err
	}
	return &TxArgs{
		To:     to,
		From:   from,
		Value:  value,
		Method: cm,
		Args:   parsedArgs,
	}, nil
}

func (cm *CeloMethod) SerializeArguments(args ...interface{}) ([]interface{}, error) {
	out := make([]interface{}, len(args))
	if len(args) != len(cm.argParsers) {
		return nil, fmt.Errorf("Received %d args; expected %d", len(args), len(cm.argParsers))
	}

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
	parsedArgs := make([]interface{}, len(cm.argParsers))
	if len(values) != len(cm.argParsers) {
		return nil, fmt.Errorf("Received %d args; expected %d", len(values), len(cm.argParsers))
	}

	var err error
	for i, value := range values {
		parsedArgs[i], err = cm.argParsers[i](value)
		if err != nil {
			return nil, fmt.Errorf("bad argument: idx=%d error=%w", i, err)
		}
	}

	return parsedArgs, nil
}
