// Copyright 2023 Celo Org
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

package service

import (
	"math/big"

	"github.com/celo-org/celo-blockchain/consensus/istanbul"
	"github.com/celo-org/celo-blockchain/params"
)

type ChainParameters struct {
	ChainId       *big.Int
	EpochSize     uint64
	IsGingerbread func(*big.Int) bool
}

func NewChainParametersFromConfig(config *params.ChainConfig) *ChainParameters {
	return &ChainParameters{
		ChainId:       config.ChainID,
		EpochSize:     config.Istanbul.Epoch,
		IsGingerbread: config.IsGingerbread,
	}
}

func (cp *ChainParameters) IsLastBlockOfEpoch(blockNumber uint64) bool {
	return istanbul.IsLastBlockOfEpoch(blockNumber, cp.EpochSize)
}
