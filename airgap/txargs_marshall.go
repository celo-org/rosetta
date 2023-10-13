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
	"encoding/json"

	"github.com/celo-org/celo-blockchain/common"
)

type txArgsRawData struct {
	From        common.Address  `json:"from"`
	Value       *string         `json:"value,omitempty"`
	To          *common.Address `json:"to,omitempty"`
	Method      *string         `json:"method,omitempty"`
	Args        []interface{}   `json:"args,omitempty"`
	ArgsEncoded *bool           `json:"args_encoded,omitempty"`
}

type callParamsRawData struct {
	txArgsRawData
	BlockNumber *string `json:"block_number,omitempty"`
}

func (data *txArgsRawData) transform(args *TxArgs) error {
	var err error
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
	args.ArgsEncoded = data.ArgsEncoded
	return nil
}

func (args *TxArgs) transform() *txArgsRawData {
	var data txArgsRawData
	data.From = args.From
	data.Value = bigIntToString(args.Value)
	data.To = args.To
	if args.Method != nil {
		var str = args.Method.String()
		data.Method = &str
	}
	data.Args = args.Args
	data.ArgsEncoded = args.ArgsEncoded
	return &data
}

func (args TxArgs) MarshalJSON() ([]byte, error) {
	return json.Marshal(args.transform())
}

func (params CallParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(callParamsRawData{
		txArgsRawData: *params.TxArgs.transform(),
		BlockNumber:   bigIntToString(params.BlockNumber),
	})
}

func (args *TxArgs) UnmarshalJSON(b []byte) error {
	var err error

	var data txArgsRawData
	if err = json.Unmarshal(b, &data); err != nil {
		return err
	}
	return data.transform(args)
}

func (params *CallParams) UnmarshalJSON(b []byte) error {
	var err error
	var data callParamsRawData

	if err = json.Unmarshal(b, &data); err != nil {
		return err
	}

	err = data.txArgsRawData.transform(&params.TxArgs)
	if err != nil {
		return err
	}

	params.BlockNumber, err = stringToBigInt(data.BlockNumber)
	return err
}
