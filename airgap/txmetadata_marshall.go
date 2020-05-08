package airgap

import (
	"encoding/json"

	"github.com/ethereum/go-ethereum/common"
)

func (tm TxMetadata) MarshalJSON() ([]byte, error) {
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
	}

	data.From = tm.From
	data.Nonce = tm.Nonce
	data.GasPrice = bigIntToString(tm.GasPrice)
	data.GatewayFeeRecipient = tm.GatewayFeeRecipient
	data.GatewayFee = bigIntToString(tm.GatewayFee)
	data.FeeCurrency = tm.FeeCurrency
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
	tm.GatewayFeeRecipient = data.GatewayFeeRecipient
	tm.GatewayFee, err = stringToBigInt(data.GatewayFee)
	if err != nil {
		return err
	}
	tm.FeeCurrency = data.FeeCurrency
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
