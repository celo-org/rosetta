package airgap

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

func (tx Transaction) MarshalJSON() ([]byte, error) {
	var data struct {
		From                common.Address  `json:"from"`
		Nonce               uint64          `json:"nonce"`
		GasPrice            *string         `json:"gasPrice,omitempty"`
		GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient,omitempty"`
		GatewayFee          *string         `json:"gatewayFee,omitempty"`
		FeeCurrency         *common.Address `json:"feeCurrency,omitempty"`
		To                  common.Address  `json:"to"`
		Data                string          `json:"data"`
		Value               *string         `json:"value,omitempty"`
		Gas                 uint64          `json:"gas"`
		ChainId             string          `json:"chainId"`
		V                   *string         `json:"v,omitempty"`
		R                   *string         `json:"r,omitempty"`
		S                   *string         `json:"s,omitempty"`
	}

	data.From = tx.From
	data.Nonce = tx.Nonce
	data.GasPrice = bigIntToString(tx.GasPrice)
	data.GatewayFeeRecipient = tx.GatewayFeeRecipient
	data.GatewayFee = bigIntToString(tx.GatewayFee)
	data.FeeCurrency = tx.FeeCurrency
	data.To = tx.To
	data.Data = common.Bytes2Hex(tx.Data)
	data.Value = bigIntToString(tx.Value)
	data.Gas = tx.Gas
	data.ChainId = *bigIntToString(tx.ChainId)
	data.V = bigIntToString(tx.V)
	data.R = bigIntToString(tx.R)
	data.S = bigIntToString(tx.S)

	return json.Marshal(data)
}

func (tx *Transaction) UnmarshalJSON(b []byte) error {
	var err error
	var data struct {
		From                common.Address  `json:"from"`
		Nonce               uint64          `json:"nonce"`
		GasPrice            *string         `json:"gasPrice,omitempty"`
		GatewayFeeRecipient *common.Address `json:"gatewayFeeRecipient,omitempty"`
		GatewayFee          *string         `json:"gatewayFee,omitempty"`
		FeeCurrency         *common.Address `json:"feeCurrency,omitempty"`
		To                  common.Address  `json:"to"`
		Data                string          `json:"data"`
		Value               *string         `json:"value,omitempty"`
		Gas                 uint64          `json:"gas"`
		ChainId             string          `json:"chainId"`
		V                   *string         `json:"v,omitempty"`
		R                   *string         `json:"r,omitempty"`
		S                   *string         `json:"s,omitempty"`
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
	tx.GatewayFeeRecipient = data.GatewayFeeRecipient
	tx.GatewayFee, err = stringToBigInt(data.GatewayFee)
	if err != nil {
		return err
	}
	tx.FeeCurrency = data.FeeCurrency
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
	tx.V, err = stringToBigInt(data.V)
	if err != nil {
		return err
	}
	tx.R, err = stringToBigInt(data.R)
	if err != nil {
		return err
	}
	tx.S, err = stringToBigInt(data.S)
	if err != nil {
		return err
	}

	return nil
}
