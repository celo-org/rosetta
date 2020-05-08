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
