package api

import (
	"math/big"
	"strconv"
	"testing"

	"github.com/celo-org/rosetta/celo"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
	gs "github.com/onsi/gomega/gstruct"
	"github.com/onsi/gomega/types"
)

func MatchOperation(account common.Address, value int, currency Currency, status OperationResult, kind OperationKind) types.GomegaMatcher {
	return gs.MatchFields(gs.IgnoreExtras, gs.Fields{
		"Account": Equal(NewAccountIdentifier(account)),
		"Amount": gs.MatchAllFields(gs.Fields{
			"Value":    Equal(strconv.Itoa(value)),
			"Currency": Equal(currency),
			"Metadata": BeNil(),
		}),
		"Status": Equal(status.String()),
		"Type":   Equal(kind.String()),
	})
}

func TestMapTxHashesToTransaction(t *testing.T) {
	g := NewGomegaWithT(t)

	txHashes := []common.Hash{common.HexToHash("1"), common.HexToHash("2"), common.HexToHash("3")}

	getHash := func(t Transaction) common.Hash { return common.HexToHash(t.TransactionIdentifier.Hash) }
	g.Expect(MapTxHashesToTransaction(txHashes)).To(And(
		HaveLen(len(txHashes)),
		ConsistOf(
			WithTransform(getHash, Equal(txHashes[0])),
			WithTransform(getHash, Equal(txHashes[1])),
			WithTransform(getHash, Equal(txHashes[2])),
		),
	))

}

func TestTransferToOperations(t *testing.T) {
	g := NewGomegaWithT(t)

	transfer := celo.Transfer{
		From:  celo.Account{Address: common.HexToAddress("1"), SubAccount: celo.Main},
		To:    celo.Account{Address: common.HexToAddress("2"), SubAccount: celo.Main},
		Value: big.NewInt(10000),
	}

	g.Expect(TransferToOperations(50, &transfer)).To(And(
		HaveLen(2),
		ConsistOf(
			MatchOperation(common.HexToAddress("1"), -10000, CeloGold, OperationSuccess, OpKindTransfer),
			MatchOperation(common.HexToAddress("2"), 10000, CeloGold, OperationSuccess, OpKindTransfer),
		),
	))
}

func TestGasDetailsToOperations(t *testing.T) {
	g := NewGomegaWithT(t)

	gasDetail := map[common.Address]*big.Int{
		common.HexToAddress("1"): big.NewInt(-1000),
		common.HexToAddress("2"): big.NewInt(800),
		common.HexToAddress("3"): big.NewInt(200),
	}

	g.Expect(GasDetailsToOperations(gasDetail)).To(And(
		HaveLen(3),
		ContainElement(MatchOperation(common.HexToAddress("1"), -1000, CeloGold, OperationSuccess, OpKindFee)),
		ContainElement(MatchOperation(common.HexToAddress("2"), 800, CeloGold, OperationSuccess, OpKindFee)),
		ContainElement(MatchOperation(common.HexToAddress("3"), 200, CeloGold, OperationSuccess, OpKindFee)),
	))
}
