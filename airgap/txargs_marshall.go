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

	"github.com/ethereum/go-ethereum/common"
)

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
