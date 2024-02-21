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

func (tx Transaction) MarshalJSON() ([]byte, error) {
	var data struct {
		From      common.Address `json:"from"`
		Nonce     uint64         `json:"nonce"`
		GasPrice  *string        `json:"gasPrice,omitempty"`
		To        common.Address `json:"to"`
		Data      string         `json:"data"`
		Value     *string        `json:"value,omitempty"`
		Gas       uint64         `json:"gas"`
		ChainId   string         `json:"chainId"`
		Signature string         `json:"signature"`
	}

	data.From = tx.From
	data.Nonce = tx.Nonce
	data.GasPrice = bigIntToString(tx.GasPrice)
	data.To = tx.To
	data.Data = common.Bytes2Hex(tx.Data)
	data.Value = bigIntToString(tx.Value)
	data.Gas = tx.Gas
	data.ChainId = *bigIntToString(tx.ChainId)
	data.Signature = common.Bytes2Hex(tx.Signature)

	return json.Marshal(data)
}

func (tx *Transaction) UnmarshalJSON(b []byte) error {
	var err error
	var data struct {
		From      common.Address `json:"from"`
		Nonce     uint64         `json:"nonce"`
		GasPrice  *string        `json:"gasPrice,omitempty"`
		To        common.Address `json:"to"`
		Data      string         `json:"data"`
		Value     *string        `json:"value,omitempty"`
		Gas       uint64         `json:"gas"`
		ChainId   string         `json:"chainId"`
		Signature string         `json:"signature"`
	}

	if err = json.Unmarshal(b, &data); err != nil {
		return err
	}

	tx.TxMetadata = &TxMetadata{}

	tx.From = data.From
	tx.Nonce = data.Nonce
	tx.GasPrice, err = stringToBigInt(data.GasPrice)
	if err != nil {
		return err
	}
	tx.To = data.To
	tx.Data = common.Hex2Bytes(data.Data)
	tx.Value, err = stringToBigInt(data.Value)
	if err != nil {
		return err
	}
	tx.Gas = data.Gas
	tx.ChainId, err = stringToBigInt(&data.ChainId)
	if err != nil {
		return err
	}
	tx.Signature = common.Hex2Bytes(data.Signature)

	return nil
}
