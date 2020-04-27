package wrapper

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func CallOptsFromTxOpts(txOpts *bind.TransactOpts) *bind.CallOpts {
	return &bind.CallOpts{
		From:    txOpts.From,
		Context: txOpts.Context,
	}
}
