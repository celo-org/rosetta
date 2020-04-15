package wrapper

import (
	"errors"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/contract"
)

type GasPriceMinimumWrapper struct {
	contract *contract.GasPriceMinimum
}

var (
	ErrGasPriceMinimumNotDeployed = errors.New("GasPriceMinimum Not Deployed")
)

func NewGasPriceMinimum(celoClient *client.CeloClient, address commom.Address) (*GasPriceMinimumWrapper, error) {
	gasPriceMinimum, err := contract.NewGasPriceMinimum(address, celoClient.Eth)
	err = client.WrapRpcError(err)
	if err != nil {
		return nil, err
	}

	return &GasPriceMinimumWrapper{
		contract: gasPriceMinimum,
	}, nil
}

func (w *GasPriceMinimumWrapper) Contract() *contract.GasPriceMinimum {
	return w.contract
}
