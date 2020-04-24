package db

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	. "github.com/onsi/gomega"
)

var _ RosettaDB = (*rosettaSqlDb)(nil)

func TestApplyChanges(t *testing.T) {
	RegisterTestingT(t)

	celoDb, err := NewSqliteDb(":memory:")
	Ω(err).ShouldNot(HaveOccurred())

	blockNumber := big.NewInt(10)
	changeSet := BlockChangeSet{
		BlockNumber: blockNumber,
	}

	err = celoDb.ApplyChanges(context.Background(), &changeSet)
	Ω(err).ShouldNot(HaveOccurred())

	storedBlockNumber, err := celoDb.LastPersistedBlock(context.Background())
	Ω(err).ShouldNot(HaveOccurred())

	Ω(storedBlockNumber.String()).To(Equal(blockNumber.String()))
}

func TestRegisterContract(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	// celoDb, err := NewSqliteDb("/tmp/prueba.db")
	celoDb, err := NewSqliteDb(":memory:")
	Ω(err).ShouldNot(HaveOccurred())

	// Changes
	// Governance 0x34   (10, 4)
	// Governance 0x34   (15, 4)

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber: big.NewInt(10),
		RegistryChanges: []RegistryChange{
			{TxIndex: 4, Contract: "Governance", NewAddress: common.HexToAddress("0x34")},
		},
	})
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber: big.NewInt(15),
		RegistryChanges: []RegistryChange{
			{TxIndex: 4, Contract: "Governance", NewAddress: common.HexToAddress("0x111")},
		},
	})
	Ω(err).ShouldNot(HaveOccurred())

	var addr common.Address

	t.Run("Before", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(2), 8, "Governance")
		Ω(err).Should(Equal(ErrContractNotFound))
	})

	t.Run("Same Block, Before Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(10), 4, "Governance")
		Ω(err).Should(Equal(ErrContractNotFound))
	})

	t.Run("Same Block & Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(10), 5, "Governance")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x34")))
	})

	t.Run("Same Block & After Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(10), 7, "Governance")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x34")))
	})

	t.Run("After Block & Before Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(11), 3, "Governance")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x34")))
	})

	t.Run("On Next Change", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(15), 5, "Governance")
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x111")))
	})

	t.Run("After Last Persisted Change", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.RegistryAddressStartOf(ctx, big.NewInt(16), 1, "Governance")
		Ω(err).Should(Equal(ErrFutureBlock))
	})
}

func TestGasPriceMinimum(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	celoDb, err := NewSqliteDb(":memory:")
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber:     big.NewInt(10),
		GasPriceMinimum: big.NewInt(50000),
	})
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber:     big.NewInt(15),
		GasPriceMinimum: big.NewInt(100000),
	})
	Ω(err).ShouldNot(HaveOccurred())

	var gpm *big.Int

	t.Run("Before", func(t *testing.T) {
		RegisterTestingT(t)
		gpm, err = celoDb.GasPriceMinimumFor(ctx, big.NewInt(10))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(gpm.Uint64()).Should(Equal(uint64(0)))
	})

	t.Run("Same Block", func(t *testing.T) {
		RegisterTestingT(t)
		gpm, err = celoDb.GasPriceMinimumFor(ctx, big.NewInt(11))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(gpm.Uint64()).Should(Equal(uint64(50000)))
	})

	t.Run("After Block", func(t *testing.T) {
		RegisterTestingT(t)
		gpm, err = celoDb.GasPriceMinimumFor(ctx, big.NewInt(15))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(gpm.Uint64()).Should(Equal(uint64(50000)))
	})

	t.Run("On Next Change", func(t *testing.T) {
		RegisterTestingT(t)
		gpm, err = celoDb.GasPriceMinimumFor(ctx, big.NewInt(16))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(gpm.Uint64()).Should(Equal(uint64(100000)))
	})

	t.Run("After Last Persisted Change", func(t *testing.T) {
		RegisterTestingT(t)
		gpm, err = celoDb.GasPriceMinimumFor(ctx, big.NewInt(17))
		Ω(err).Should(Equal(ErrFutureBlock))
	})

}

func TestTobinTax(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	celoDb, err := NewSqliteDb(":memory:")
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber: big.NewInt(10),
		TobinTax:    big.NewInt(50000),
	})
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber: big.NewInt(15),
		TobinTax:    big.NewInt(100000),
	})
	Ω(err).ShouldNot(HaveOccurred())

	var tobinTax *big.Int

	t.Run("Before", func(t *testing.T) {
		RegisterTestingT(t)
		tobinTax, err = celoDb.TobinTaxFor(ctx, big.NewInt(9))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(tobinTax.Uint64()).Should(Equal(uint64(0)))
	})

	t.Run("Same Block", func(t *testing.T) {
		RegisterTestingT(t)
		tobinTax, err = celoDb.TobinTaxFor(ctx, big.NewInt(10))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(tobinTax.Uint64()).Should(Equal(uint64(50000)))
	})

	t.Run("After Block", func(t *testing.T) {
		RegisterTestingT(t)
		tobinTax, err = celoDb.TobinTaxFor(ctx, big.NewInt(11))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(tobinTax.Uint64()).Should(Equal(uint64(50000)))
	})

	t.Run("On Next Change", func(t *testing.T) {
		RegisterTestingT(t)
		tobinTax, err = celoDb.TobinTaxFor(ctx, big.NewInt(15))
		Ω(err).ShouldNot(HaveOccurred())
		Ω(tobinTax.Uint64()).Should(Equal(uint64(100000)))
	})

	t.Run("After Last Persisted Change", func(t *testing.T) {
		RegisterTestingT(t)
		tobinTax, err = celoDb.TobinTaxFor(ctx, big.NewInt(17))
		Ω(err).Should(Equal(ErrFutureBlock))
	})

}

func TestGasPriceMinimum_VeryLargeNumber(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	// celoDb, err := NewSqliteDb("/tmp/prueba2.db")
	celoDb, err := NewSqliteDb(":memory:")
	Ω(err).ShouldNot(HaveOccurred())

	value := new(big.Int).SetBytes([]byte{255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255, 255})
	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber:     big.NewInt(10),
		GasPriceMinimum: value,
	})
	Ω(err).ShouldNot(HaveOccurred())

	var gpm *big.Int

	gpm, err = celoDb.GasPriceMinimumFor(ctx, big.NewInt(11))
	Ω(err).ShouldNot(HaveOccurred())
	Ω(gpm.String()).Should(Equal(value.String()))
}

func TestSetCarbonOffsetPartner(t *testing.T) {
	RegisterTestingT(t)
	ctx := context.Background()

	celoDb, err := NewSqliteDb(":memory:")
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber: big.NewInt(10),
		CarbonOffsetPartnerChange: CarbonOffsetPartnerChange{
			TxIndex: 4,
			Address: common.HexToAddress("0x34"),
		},
	})
	Ω(err).ShouldNot(HaveOccurred())

	err = celoDb.ApplyChanges(ctx, &BlockChangeSet{
		BlockNumber: big.NewInt(15),
		CarbonOffsetPartnerChange: CarbonOffsetPartnerChange{
			TxIndex: 4,
			Address: common.HexToAddress("0x111"),
		},
	})
	Ω(err).ShouldNot(HaveOccurred())

	var addr common.Address

	t.Run("Before", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(2), 8)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.ZeroAddress))
	})

	t.Run("Same Block, Before Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(10), 4)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.ZeroAddress))
	})

	t.Run("Same Block & Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(10), 5)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x34")))
	})

	t.Run("Same Block & After Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(10), 7)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x34")))
	})

	t.Run("After Block & Before Tx", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(11), 3)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x34")))
	})

	t.Run("On Next Change", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(15), 5)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(addr).Should(Equal(common.HexToAddress("0x111")))
	})

	t.Run("After Last Persisted Change", func(t *testing.T) {
		RegisterTestingT(t)
		addr, err = celoDb.CarbonOffsetPartnerStartOf(ctx, big.NewInt(16), 1)
		Ω(err).Should(Equal(ErrFutureBlock))
	})
}
