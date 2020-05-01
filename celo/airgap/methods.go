package airgap

import (
	"fmt"
	"strings"

	"github.com/celo-org/rosetta/celo/wrapper"
)

type CeloMethod struct {
	// Name of the abi method
	Name string
	// Registry id of contract where the method is defined
	Contract wrapper.RegistryKey
}

var methodsMap = make(map[wrapper.RegistryKey]map[string]*CeloMethod)

// FromString returns the CeloMethod that matches the given string
// Methods are represented as "Contract.Name"
func FromString(celoMethodStr string) (*CeloMethod, error) {
	parts := strings.Split(celoMethodStr, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	m, ok := methodsMap[(wrapper.RegistryKey)(parts[0])]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	cm, ok := m[parts[1]]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	return cm, nil
}

func registerMethod(contract wrapper.RegistryKey, name string) *CeloMethod {
	cm := &CeloMethod{
		Name:     name,
		Contract: contract,
	}
	// register the method in the methodsMap
	if methodsMap[contract] == nil {
		methodsMap[contract] = make(map[string]*CeloMethod)
	}
	methodsMap[contract][name] = cm
	return cm
}

// values taken from contract method names for ABI usage
var (
	CreateAccount       = registerMethod(wrapper.AccountsRegistryId, "createAccount")
	AuthorizeVoteSigner = registerMethod(wrapper.AccountsRegistryId, "authorizeVoteSigner")
	LockGold            = registerMethod(wrapper.LockedGoldRegistryId, "lock")
	UnlockGold          = registerMethod(wrapper.LockedGoldRegistryId, "unlock")
	RelockGold          = registerMethod(wrapper.LockedGoldRegistryId, "relock")
	WithdrawGold        = registerMethod(wrapper.LockedGoldRegistryId, "withdraw")
	Vote                = registerMethod(wrapper.ElectionRegistryId, "vote")
	ActivateVotes       = registerMethod(wrapper.ElectionRegistryId, "activate")
	RevokePendingVotes  = registerMethod(wrapper.ElectionRegistryId, "revokePending")
	RevokeActiveVotes   = registerMethod(wrapper.ElectionRegistryId, "revokeActive")
)

func (tt *CeloMethod) String() string { return fmt.Sprintf("%s.%s", tt.Contract, tt.Name) }
