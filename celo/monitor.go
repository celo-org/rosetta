package celo

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/contract"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

type DB struct {
	// TODO(Alec): add db
}

type Monitor struct {
	startBlock *big.Int
	stop       <-chan struct{}
	db         *DB
}

type BlockResult struct {
	header *types.Header
	logs   []types.Log
}

func NewMonitor(cc *client.CeloClient, startBlock *big.Int) (*Monitor, error) {

	if startBlock == nil {
		startBlock = big.NewInt(0)
	}

	done := make(chan struct{})
	stop := make(chan struct{})
	errors := make(chan error)

	go func() {
		defer close(done)
		select {
		case err := <-errors:
			log.Error("Monitor Error", "error", err)
			return
		case <-stop:
			return
		}
	}()

	headers := getHeaders(cc, done, errors, startBlock)
	results := processHeaders(cc, done, errors, headers)
	db := storeResults(done, errors, results)

	return &Monitor{
		startBlock: startBlock,
		stop:       stop,
		db:         db,
	}, nil
}

func getHeaders(cc *client.CeloClient, done <-chan struct{}, errors chan<- error, startBlock *big.Int) <-chan *types.Header {
	ctx := context.Background()

	headers := make(chan *types.Header)

	synced := make(chan struct{})

	sync := func() {
		//TODO(Alec): Add concurrency.

		defer func() {
			synced <- struct{}{}
		}()

		// First, get most recent header so we have the current block height.
		latest, err := cc.Eth.HeaderByNumber(ctx, nil)
		log.Info("latest", "BlockNum", latest.Number.Int64())
		if err != nil {
			select {
			case errors <- err:
			case <-done:
			}
			return
		}

		// Then, get all previous headers.
		for i := startBlock; i.Cmp(latest.Number) < 0; i.Add(i, big.NewInt(1)) {
			h, err := cc.Eth.HeaderByNumber(ctx, i)
			if err != nil {
				select {
				case errors <- err:
				case <-done:
				}
				return
			}

			select {
			case headers <- h:
			case <-done:
				return
			}
		}

		select {
		case headers <- latest:
		case <-done:
		}
	}

	go func() {
		defer close(synced)
		defer close(headers)

		go sync()

		select {
		case <-synced:
		case <-done:
			return
		}

		sub, err := cc.Eth.SubscribeNewHead(context.Background(), headers)
		if err != nil {
			select {
			case errors <- err:
			case <-done:
			}
		} else {
			defer sub.Unsubscribe()

			log.Info("NewHeadSubscription started")

			select {
			case err := <-sub.Err():
				select {
				case errors <- err:
				case <-done:
					return
				}
			case <-done:
				return
			}
		}
	}()

	return headers
}

func processHeaders(cc *client.CeloClient, done <-chan struct{}, errors chan<- error, headers <-chan *types.Header) <-chan *BlockResult {
	results := make(chan *BlockResult)

	registryABI, err := contract.ParseRegistryABI()
	if err != nil {
		select {
		case errors <- err:
		case <-done:
		}
		return nil
	}
	registryUpdatedTopic := registryABI.Events["RegistryUpdated"].ID()

	process := func(h *types.Header) *BlockResult {
		//TODO(Alec): filter logs from other addresses too.

		logs, err := cc.Eth.FilterLogs(context.Background(), ethereum.FilterQuery{
			FromBlock: h.Number,
			ToBlock:   h.Number,
			Addresses: []common.Address{params.RegistrySmartContractAddress},
			Topics:    [][]common.Hash{[]common.Hash{registryUpdatedTopic}},
		})
		if err != nil {
			select {
			case errors <- err:
			case <-done:
			}
			return nil
		}

		for _, l := range logs {
			if l.Removed {
				continue
			}

			if l.Topics[0] == registryUpdatedTopic {
				registryId := l.Topics[1]
				newAddress := l.Topics[2]
				txIndex := l.TxIndex // index of the transaction in the block

				log.Info("Core Contract Address Changed", "registryId", registryId, "newAddress", newAddress, "txIndex", txIndex)
			}
		}

		return &BlockResult{
			header: h,
			logs:   logs,
		}
	}

	go func() {
		defer close(results)
		for h := range headers {
			res := process(h)
			select {
			case results <- res:
			case <-done:
				return
			}
		}
	}()

	return results
}

func storeResults(done <-chan struct{}, errors chan<- error, results <-chan *BlockResult) *DB {
	db := &DB{}

	go func() {
		for {
			select {
			case res := <-results:
				//TODO(Alec): write to db
				log.Info("Storing Header", "number", res.header.Number.Int64())
				if len(res.logs) > 0 {
					//log.Info("Storing Header", "number", res.header.Number.Int64())
					log.Info("Storing Logs", "logs", res.logs)
				}
				//log.Info("Storing Header", "header", res.header)
			case <-done:
				return
			}
		}
	}()

	return db
}
