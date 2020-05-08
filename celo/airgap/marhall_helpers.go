package airgap

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func bigIntToString(input *big.Int) *string {
	if input == nil {
		return nil
	}
	out := input.String()
	return &out
}

func addressToString(input *common.Address) *string {
	if input == nil {
		return nil
	}
	out := input.Hex()
	return &out
}

func stringToBigInt(input *string) (*big.Int, error) {
	if input == nil {
		return nil, nil
	}

	out, ok := new(big.Int).SetString(*input, 10)
	if !ok {
		return nil, fmt.Errorf("Invalid bigInt format: %s", *input)
	}
	return out, nil
}

func stringToMethod(input *string) (*CeloMethod, error) {
	if input == nil {
		return nil, nil
	}
	out, err := MethodFromString(*input)
	if err != nil {
		return nil, err
	}
	return out, nil
}
