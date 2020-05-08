package airgap

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

func (args *TxArgs) MarshallMap() (map[string]interface{}, error) {
	data, err := args.MarshalJSON()
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

func (args *TxArgs) UnmarshallMap(input map[string]interface{}) error {
	data, err := json.Marshal(input)
	if err != nil {
		return err
	}

	return args.UnmarshalJSON(data)
}

// func (args *TxArgs) MarshallMap() (map[string]interface{}, error) {
// 	output := make(map[string]interface{})

// 	output["from"] = args.From.Hex()
// 	output["args"] = args.Args

// 	if args.Value != nil {
// 		output["value"] = args.Value.String()
// 	}

// 	if args.To != nil {
// 		output["to"] = args.To.Hex()
// 	}

// 	if args.Method != nil {
// 		output["method"] = args.Method.String()
// 	}

// 	return output, nil
// }

// func (args *TxArgs) UnmarshallMap(input map[string]interface{}) error {
// 	var values struct {
// 		From   string        `mapstructure:"from"`
// 		To     *string       `mapstructure:"to,omitempty"`
// 		Value  *string       `mapstructure:"value,omitempty"`
// 		Method *string       `mapstructure:"method,omitempty"`
// 		Args   []interface{} `mapstructure:"args,omitempty"`
// 	}

// 	err := mapstructure.Decode(input, &values)
// 	if err != nil {
// 		return err
// 	}

// 	args.From = common.HexToAddress(values.From)
// 	args.Args = values.Args

// 	if values.To != nil {
// 		val := common.HexToAddress(*values.To)
// 		args.To = &val
// 	}

// 	if values.Value != nil {
// 		val, ok := new(big.Int).SetString(*values.Value, 10)
// 		if !ok {
// 			return fmt.Errorf("'value': can't convert to big.Int")
// 		}
// 		args.Value = val
// 	}

// 	if values.Method != nil {
// 		m, err := MethodFromString(*values.Method)
// 		if err != nil {
// 			return fmt.Errorf("'method': invalid method %s", *values.Method)
// 		}
// 		args.Method = m
// 	}

// 	return nil
// }

func (args TxArgs) MarshalJSON() ([]byte, error) {
	var data struct {
		From   common.Address  `json:"from"`
		Value  *string         `json:"value,omitempty"`
		To     *common.Address `json:"to,omitempty"`
		Method *string         `json:"method,omitempty"`
		Args   []interface{}   `json:"args,omitempty"`
	}

	data.From = args.From
	data.Value = bigIntToString(args.Value)
	data.To = args.To
	if args.Method != nil {
		var str = args.Method.String()
		data.Method = &str
	}
	data.Args = args.Args

	return json.Marshal(data)
}

func (args *TxArgs) UnmarshalJSON(b []byte) error {
	var err error
	var data struct {
		From   common.Address  `json:"from"`
		Value  *string         `json:"value,omitempty"`
		To     *common.Address `json:"to,omitempty"`
		Method *string         `json:"method,omitempty"`
		Args   []interface{}   `json:"args,omitempty"`
	}

	if err = json.Unmarshal(b, &data); err != nil {
		return err
	}

	args.From = data.From
	args.Value, err = stringToBigInt(data.Value)
	if err != nil {
		return err
	}
	args.To = data.To
	args.Method, err = stringToMethod(data.Method)
	if err != nil {
		return err
	}
	args.Args = data.Args

	return nil
}
