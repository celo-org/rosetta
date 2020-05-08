package airgap

import (
	"encoding/json"
	"fmt"
	"math/big"
)

func MarshallToMap(input interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	var output map[string]interface{}
	err = json.Unmarshal(data, &output)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func UnmarshallFromMap(input map[string]interface{}, output interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, output)
}

func bigIntToString(input *big.Int) *string {
	if input == nil {
		return nil
	}
	out := input.String()
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
