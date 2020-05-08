package airgap

import (
	"fmt"
	"strings"
)

var methodRegistry = make(map[string]map[string]*CeloMethod)

// FromString returns the CeloMethod that matches the given string
// Methods are represented as "Contract.Name"
func MethodFromString(celoMethodStr string) (*CeloMethod, error) {
	parts := strings.Split(celoMethodStr, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	m, ok := methodRegistry[parts[0]]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	cm, ok := m[parts[1]]
	if !ok {
		return nil, fmt.Errorf("Invalid method string: %s", celoMethodStr)
	}
	return cm, nil
}

func registerMethod(contract string, name string, argParsers []argParser) *CeloMethod {
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
