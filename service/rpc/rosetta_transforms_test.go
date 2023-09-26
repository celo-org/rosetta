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
	"math/big"
	"strconv"
	"testing"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/rosetta/analyzer"

	rosettaTypes "github.com/coinbase/rosetta-sdk-go/types"
	. "github.com/onsi/gomega"
	gs "github.com/onsi/gomega/gstruct"
	gtypes "github.com/onsi/gomega/types"
)

func MatchOperation(account common.Address, value int, currency *rosettaTypes.Currency, status OperationResult, kind analyzer.OperationType) gtypes.GomegaMatcher {
	return gs.PointTo(gs.MatchFields(gs.IgnoreExtras, gs.Fields{
		"Account": gs.PointTo(Equal(NewAccountIdentifier(account, nil))),
		"Amount": gs.PointTo(gs.MatchAllFields(gs.Fields{
			"Value":    Equal(strconv.Itoa(value)),
			"Currency": Equal(currency),
			"Metadata": BeNil(),
		})),
		"Status": Equal(status.String()),
		"Type":   Equal(kind.String()),
	}))
}

func TestMapTxHashesToTransaction(t *testing.T) {
	RegisterTestingT(t)

	txHashes := []common.Hash{common.HexToHash("1"), common.HexToHash("2"), common.HexToHash("3")}

	getHashesFromIds := func(txIds []*rosettaTypes.TransactionIdentifier) []common.Hash {
		txHashes := make([]common.Hash, len(txIds))
		for i, t := range txIds {
			txHashes[i] = common.HexToHash(t.Hash)
		}
		return txHashes
	}

	Ω(MapTxHashesToTransaction(txHashes)).Should(And(
		HaveLen(len(txHashes)),
		WithTransform(getHashesFromIds, ConsistOf(txHashes)),
	))
}

func TestTransferToOperations(t *testing.T) {
	RegisterTestingT(t)

	aop := analyzer.NewTransfer(common.HexToAddress("1"), common.HexToAddress("2"), big.NewInt(10000), true)

	Ω(OperationsFromAnalyzer(aop, 50)).Should(And(
		HaveLen(2),
		ConsistOf(
			MatchOperation(common.HexToAddress("1"), -10000, CeloGold, OperationSuccess, analyzer.OpTransfer),
			MatchOperation(common.HexToAddress("2"), 10000, CeloGold, OperationSuccess, analyzer.OpTransfer),
		),
	))
}

func TestGasDetailsToOperations(t *testing.T) {
	RegisterTestingT(t)

	gasDetail := analyzer.NewFee(map[common.Address]*big.Int{
		common.HexToAddress("1"): big.NewInt(-1000),
		common.HexToAddress("2"): big.NewInt(800),
		common.HexToAddress("3"): big.NewInt(200),
	})

	Ω(OperationsFromAnalyzer(gasDetail, 0)).Should(And(
		HaveLen(3),
		ConsistOf(
			MatchOperation(common.HexToAddress("1"), -1000, CeloGold, OperationSuccess, analyzer.OpFee),
			MatchOperation(common.HexToAddress("2"), 800, CeloGold, OperationSuccess, analyzer.OpFee),
			MatchOperation(common.HexToAddress("3"), 200, CeloGold, OperationSuccess, analyzer.OpFee),
		),
	))
}

func TestOperationsFromAnalyzer_RelatedOpsCounter(t *testing.T) {
	RegisterTestingT(t)

	gasDetail := analyzer.NewFee(map[common.Address]*big.Int{
		common.HexToAddress("1"): big.NewInt(-1000),
		common.HexToAddress("2"): big.NewInt(800),
		common.HexToAddress("3"): big.NewInt(200),
	})

	getRelatedOps := func(op *rosettaTypes.Operation) []int64 {
		ops := make([]int64, len(op.RelatedOperations))
		for i, relatedOp := range op.RelatedOperations {
			ops[i] = relatedOp.Index
		}
		return ops
	}

	Ω(OperationsFromAnalyzer(gasDetail, 0)).Should(
		ConsistOf(
			WithTransform(getRelatedOps, BeEmpty()),
			WithTransform(getRelatedOps, Equal([]int64{0})),
			WithTransform(getRelatedOps, Equal([]int64{0, 1})),
		),
	)

	Ω(OperationsFromAnalyzer(gasDetail, 5)).Should(
		ConsistOf(
			WithTransform(getRelatedOps, BeEmpty()),
			WithTransform(getRelatedOps, Equal([]int64{5})),
			WithTransform(getRelatedOps, Equal([]int64{5, 6})),
		),
	)
}
