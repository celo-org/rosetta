package rpc

// import (
// 	"strconv"

// 	"github.com/celo-org/rosetta/analyzer"
// 	"github.com/coinbase/rosetta-sdk-go/types"
// 	"github.com/ethereum/go-ethereum/common"
// 	. "github.com/onsi/gomega"
// 	gs "github.com/onsi/gomega/gstruct"
// 	gtypes "github.com/onsi/gomega/types"
// )

// func MatchOperation(account common.Address, value int, currency *types.Currency, status OperationResult, kind analyzer.OperationType) gtypes.GomegaMatcher {
// 	return gs.MatchFields(gs.IgnoreExtras, gs.Fields{
// 		"Account": Equal(NewAccountIdentifier(account)),
// 		"Amount": gs.MatchAllFields(gs.Fields{
// 			"Value":    Equal(strconv.Itoa(value)),
// 			"Currency": Equal(currency),
// 			"Metadata": BeNil(),
// 		}),
// 		"Status": Equal(status.String()),
// 		"Type":   Equal(kind.String()),
// 	})
// }

// // func TestMapTxHashesToTransaction(t *testing.T) {
// 	g := NewGomegaWithT(t)

// 	txHashes := []common.Hash{common.HexToHash("1"), common.HexToHash("2"), common.HexToHash("3")}

// 	getHash := func(t types.TransactionIdentifier) common.Hash { return common.HexToHash(t.Hash) }
// 	g.Expect(MapTxHashesToTransaction(txHashes)).To(And(
// 		HaveLen(len(txHashes)),
// 		ConsistOf(
// 			WithTransform(getHash, Equal(txHashes[0])),
// 			WithTransform(getHash, Equal(txHashes[1])),
// 			WithTransform(getHash, Equal(txHashes[2])),
// 		),
// 	))

// }

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
