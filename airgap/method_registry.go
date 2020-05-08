package airgap

import (
	"fmt"
	"strings"

	"github.com/celo-org/rosetta/celo/wrapper"
)

var methodRegistry = make(map[wrapper.RegistryKey]map[string]*CeloMethod)

// FromString returns the CeloMethod that matches the given string
// Methods are represented as "Contract.Name"
func MethodFromString(celoMethodStr string) (*CeloMethod, error) {
	parts := strings.Split(celoMethodStr, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	m, ok := methodRegistry[(wrapper.RegistryKey)(parts[0])]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	cm, ok := m[parts[1]]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	return cm, nil
}

func registerMethod(contract wrapper.RegistryKey, name string, argParsers []argParser) *CeloMethod {
	cm := &CeloMethod{
		Name:       name,
		Contract:   contract,
		argParsers: argParsers,
	}
	// register the method in the methodRegistry
	if methodRegistry[contract] == nil {
		methodRegistry[contract] = make(map[string]*CeloMethod)
	}
	methodRegistry[contract][name] = cm
	return cm
}
