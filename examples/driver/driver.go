package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/big"
	"os"

	"github.com/celo-org/rosetta/celo"
	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/celo/wrapper"
	"github.com/celo-org/rosetta/internal/config"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/k0kubun/pp"
)

var ctx = context.Background()

func main() {
	log.Root().SetHandler(log.LvlFilterHandler(log.LvlInfo, log.StreamHandler(os.Stderr, log.TerminalFormat(true))))

	// FetchEveryXBlocksByHash()
	// CheckHeaderHash()
	// DriverRegistryErrors()
	DriverTransactionTransfer()
}

func DriverTransactionTransfer() {
	cc := CeloClient()

	txHash := common.HexToHash("0xc21b6442885f4375e323097d567c8a406c1589a30b984c768de536add1d7b2e1")

	transfers, err := cc.Debug.TransactionTransfers(ctx, txHash)
	if err != nil {
		panic(err)
	}

	pp.Print(transfers)
}

func DriverRegistryErrors() {
	cc := CeloClient()

	registry, err := wrapper.NewRegistry(cc)
	if err != nil {
		panic(err)
	}

	// // Find where Registry Was deployed
	// value := BinarySearch(0, 226008, func(n uint) bool {
	// 	_, err := registry.GetAddressForString(&bind.CallOpts{
	// 		BlockNumber: new(big.Int).SetUint64(uint64(n)),
	// 	}, "StableToken")
	// 	return err != wrapper.ErrRegistryNotDeployed
	// })
	// fmt.Println(value)

	// Answer is 670 it already exist, so deployed on 669

	// registry.GetUpdatesOnBlock((ctx, 669))

	// Find where Registry Was deployed
	value := BinarySearch(670, 226008, func(n uint) bool {
		_, err := registry.GetAddressForString(&bind.CallOpts{
			BlockNumber: new(big.Int).SetUint64(uint64(n)),
		}, "Governance")
		return err != client.ErrContractNotDeployed
	})
	fmt.Println(value)
}

func TxContextDriver() {
	cc := CeloClient()

	txHash := common.HexToHash("0xd6ab1c883179b677d2120c7a0d2f2a32351bd735b5c76386d13b2c23eb33ce4c")
	blockNumber := big.NewInt(222130)

	header, err := cc.Eth.HeaderByNumber(ctx, blockNumber)
	if err != nil {
		panic(err)
	}

	tx, _, err := cc.Eth.TransactionByHash(ctx, txHash)
	if err != nil {
		panic(err)
	}

	receipt, err := cc.Eth.TransactionReceipt(ctx, tx.Hash())
	if err != nil {
		panic(err)
	}

	txTracer := celo.NewTxTracer(
		ctx,
		cc,
		header,
		tx,
		receipt,
	)

	gasDetail, err := txTracer.GasDetail()
	pp.Print(gasDetail)
	pp.Print(err)
}

func CheckHeaderHash() {
	celo := CeloClient()
	lastBlock, err := celo.Eth.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	block, err := celo.Eth.HeaderByHash(context.Background(), lastBlock.ParentHash)
	if err != nil {
		panic(err)
	}

	log.Info("Block", "num", block.Number, "hash", block.Hash().Hex(), "correctHash", lastBlock.ParentHash.Hex(), "mixDigest", lastBlock.MixDigest.Hex())
}

func FetchEveryXBlocksByHash() {
	step := big.NewInt(1)
	celo, err := client.Dial(config.FornoAlfajoresUrl)
	if err != nil {
		panic(err)
	}

	lastBlock, err := celo.Eth.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	log.Info("LastBlock", "num", lastBlock.Number, "hash", lastBlock.Hash().Hex(), "parentHash", lastBlock.ParentHash.Hex())

	currBlock := lastBlock
	initial := currBlock.Number
	for {
		nextNumber := new(big.Int).Sub(currBlock.Number, step)

		logger := log.New("num", nextNumber, "distance", new(big.Int).Sub(initial, nextNumber))

		logger.Info("Fetching By Number")
		block, err := celo.Eth.ExtendedHeaderByNumber(context.Background(), nextNumber)
		if err != nil {
			logger.Info("Error Fetching By Number", "err", err)
			os.Exit(1)
		}

		blockHash := block.Hash()
		logger = logger.New("hash", blockHash.Hex())
		logger.Info("Fetching By hash")

		block, err = celo.Eth.ExtendedHeaderByHash(context.Background(), blockHash)
		if err != nil {
			logger.Info("Error Fetching By Hash", "err", err)
			os.Exit(1)
		}
		currBlock = &block.Header
	}

}

func FetchAllBlockByHash() {
	celo, err := client.Dial(config.FornoAlfajoresUrl)
	if err != nil {
		panic(err)
	}

	lastBlock, err := celo.Eth.HeaderByNumber(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	log.Info("LastBlock", "num", lastBlock.Number, "hash", lastBlock.Hash)

	currBlock := lastBlock
	for {
		log.Info("Fetching", "num", currBlock.Number.Int64()-1, "hash", currBlock.ParentHash.Hex())
		parentBlock, err := celo.Eth.HeaderByHash(context.Background(), currBlock.ParentHash)
		if err != nil {
			log.Info("Error Fetching", "num", currBlock.Number.Int64()-1, "hash", currBlock.ParentHash.Hex(), "err", err)
			block, err := celo.Eth.ExtendedHeaderByNumber(context.Background(), new(big.Int).Sub(currBlock.Number, big.NewInt(1)))
			if err != nil {
				panic(err)
			}
			MustPPHeader(block)

			os.Exit(1)
		}
		currBlock = parentBlock
	}
}

//885363
//885671

func MustPPHeader(block *ethclient.ExtendedHeader) {
	str, err := json.MarshalIndent(block, "  ", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(str))
}

func CeloClient() *client.CeloClient {
	celo, err := client.Dial(config.FornoRC0Url)
	if err != nil {
		panic(err)
	}
	return celo
}

func BinarySearch(low uint, right uint, test func(n uint) bool) uint {
	lowValue := test(low)
	rightValue := test(right)
	if lowValue == rightValue {
		panic("Bad Initial Values")
	}

	for low+1 != right {
		mid := (low + right) / 2
		fmt.Printf("low=%d  mid=%d  right=%d\n", low, mid, right)

		if test(mid) == lowValue {
			low = mid
		} else {
			right = mid
		}
	}
	return right
}
