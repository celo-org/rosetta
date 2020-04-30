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

package rpc

import (
	"testing"

	"github.com/coinbase/rosetta-sdk-go/types"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

// func MatchOperation(account common.Address, value int, currency *types.Currency, status OperationResult, kind analyzer.OperationType) gtypes.GomegaMatcher {
// 	return gs.MatchFields(gs.IgnoreExtras, gs.Fields{
// 		"Account": Equal(NewAccountIdentifier(account, nil)),
// 		"Amount": gs.MatchAllFields(gs.Fields{
// 			"Value":    Equal(strconv.Itoa(value)),
// 			"Currency": Equal(currency),
// 			"Metadata": BeNil(),
// 		}),
// 		"Status": Equal(status.String()),
// 		"Type":   Equal(kind.String()),
// 	})
// }

func TestMapTxHashesToTransaction(t *testing.T) {
	RegisterTestingT(t)

	txHashes := []common.Hash{common.HexToHash("1"), common.HexToHash("2"), common.HexToHash("3")}

	getHashesFromIds := func(txIds []*types.TransactionIdentifier) []common.Hash {
		txHashes := make([]common.Hash, len(txIds))
		for i, t := range txIds {
			txHashes[i] = common.HexToHash(t.Hash)
		}
		return txHashes
	}

	Ω(MapTxHashesToTransaction(txHashes)).To(And(
		HaveLen(len(txHashes)),
		WithTransform(getHashesFromIds, ConsistOf(txHashes)),
	))
}

// func TestTransferToOperations(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	aop := analyzer.NewTransfer(common.HexToAddress("1"), common.HexToAddress("2"), big.NewInt(10000), true)

// 	g.Expect(OperationsFromAnalyzer(aop, 50)).To(And(
// 		HaveLen(2),
// 		ConsistOf(
// 			MatchOperation(common.HexToAddress("1"), -10000, CeloGold, OperationSuccess, analyzer.OpTransfer),
// 			MatchOperation(common.HexToAddress("2"), 10000, CeloGold, OperationSuccess, analyzer.OpTransfer),
// 		),
// 	))
// }

// func TestGasDetailsToOperations(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	gasDetail := analyzer.NewFee(map[common.Address]*big.Int{
// 		common.HexToAddress("1"): big.NewInt(-1000),
// 		common.HexToAddress("2"): big.NewInt(800),
// 		common.HexToAddress("3"): big.NewInt(200),
// 	})

// 	g.Expect(OperationsFromAnalyzer(gasDetail, 0)).To(And(
// 		HaveLen(3),
// 		ContainElement(MatchOperation(common.HexToAddress("1"), -1000, CeloGold, OperationSuccess, analyzer.OpFee)),
// 		ContainElement(MatchOperation(common.HexToAddress("2"), 800, CeloGold, OperationSuccess, analyzer.OpFee)),
// 		ContainElement(MatchOperation(common.HexToAddress("3"), 200, CeloGold, OperationSuccess, analyzer.OpFee)),
// 	))
// }
