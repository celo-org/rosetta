package utils

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

func AddressIndexOf(slice []common.Address, item common.Address) (*big.Int, error) {
	for idx, currItem := range slice {
		if currItem == item {
			return big.NewInt(int64(idx)), nil
		}
	}
	return nil, fmt.Errorf("Item not found in slice")
}
