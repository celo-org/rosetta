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

func (tm TxMetadata) MarshalJSON() ([]byte, error) {
	var data struct {
		From     common.Address `json:"from"`
		Nonce    uint64         `json:"nonce"`
		GasPrice *string        `json:"gasPrice,omitempty"`
		To       common.Address `json:"to"`
		Data     string         `json:"data"`
		Value    *string        `json:"value,omitempty"`
		Gas      uint64         `json:"gas"`
		ChainId  string         `json:"chainId"`
	}

	data.From = tm.From
	data.Nonce = tm.Nonce
	data.GasPrice = bigIntToString(tm.GasPrice)
	data.To = tm.To
	data.Data = common.Bytes2Hex(tm.Data)
	data.Value = bigIntToString(tm.Value)
	data.Gas = tm.Gas
	data.ChainId = *bigIntToString(tm.ChainId)

	return json.Marshal(data)
}

func (tm *TxMetadata) UnmarshalJSON(b []byte) error {
	var err error
	var data struct {
		From     common.Address `json:"from"`
		Nonce    uint64         `json:"nonce"`
		GasPrice *string        `json:"gasPrice,omitempty"`
		To       common.Address `json:"to"`
		Data     string         `json:"data"`
		Value    *string        `json:"value,omitempty"`
		Gas      uint64         `json:"gas"`
		ChainId  string         `json:"chainId"`
	}

	if err = json.Unmarshal(b, &data); err != nil {
		return err
	}

	tm.From = data.From
	tm.Nonce = data.Nonce
	tm.GasPrice, err = stringToBigInt(data.GasPrice)
	if err != nil {
		return err
	}
	tm.To = data.To
	tm.Data = common.Hex2Bytes(data.Data)
	tm.Value, err = stringToBigInt(data.Value)
	if err != nil {
		return err
	}
	tm.Gas = data.Gas
	tm.ChainId, err = stringToBigInt(&data.ChainId)
	if err != nil {
		return err
	}

	return nil
}
