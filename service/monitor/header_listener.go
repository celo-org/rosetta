package monitor

import (
	"context"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/service"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

func HeaderListener(ctx context.Context, headers chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock *big.Int) error {
	logger = logger.New("pipe", "header_listener")

	// Subscriptions also have a built in buffer/queue of size 2000.
	// So, 3000 new headers can be produced while we're fetching old
	// blocks before the subscription overflows.
	newHeaders := make(chan *types.Header, 1000)
	defer close(newHeaders)

	syncMode := func(sub ethereum.Subscription) error {
		for {
			select {
			case err := <-sub.Err():
				return err
			case <-ctx.Done():
				return ctx.Err()
			case h := <-newHeaders:
				logger.Info("New Header", "block", h.Number.Int64())
				if err := sendToProcessor(ctx, headers, h); err != nil {
					return err
				}
			}
		}
	}

	catchUpMode := func() error {
		lastBlock, err := lastNodeBlockNumber(ctx, cc)
		if err != nil {
			return err
		}

		if lastBlock.Cmp(startBlock) > 0 {
			logger.Info("Start fetching old blocks", "start", startBlock, "end", lastBlock)
			if err = fetchHeaderRange(ctx, headers, cc, logger, startBlock, lastBlock); err != nil {
				return err
			}
			logger.Info("Finished fetching old blocks", "start", startBlock, "end", lastBlock, "new", len(newHeaders))
		}
		return nil
	}

	for {
		sub, err := cc.Eth.SubscribeNewHead(ctx, newHeaders)
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		if err := catchUpMode(); err != nil {
			return err
		}

		if err := syncMode(sub); err != nil {
			if err == rpc.ErrSubscriptionQueueOverflow {
				logger.Error("Subscription Queue Overflowed", "New headers in buffer", len(newHeaders))
				if len(newHeaders) > 0 {
					var h *types.Header
					for h = range newHeaders {
						if err := sendToProcessor(ctx, headers, h); err != nil {
							return err
						}
					}
					startBlock = new(big.Int).Add(h.Number, big.NewInt(1))
				}
				continue
			}
			return err
		}
		return nil
	}
}

func sendToProcessor(ctx context.Context, headers chan<- *types.Header, h *types.Header) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case headers <- h:
	}
	return nil
}

func fetchHeaderRange(ctx context.Context, headers chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock, endBlock *big.Int) error {
	var wg sync.WaitGroup
	var errorCollector service.ErrorCollector
	batchSize := 20

	count := 0

	for i := new(big.Int).Set(startBlock); i.Cmp(endBlock) <= 0; i.Add(i, big.NewInt(int64(batchSize))) {
		remaining := new(big.Int).Add(new(big.Int).Sub(endBlock, i), big.NewInt(1))
		if remaining.Cmp(big.NewInt(int64(batchSize))) < 0 {
			batchSize = int(remaining.Int64())
		}

		mu := &sync.Mutex{}
		batch := make([]*types.Header, batchSize)
		wg.Add(batchSize)
		for j := 0; j < batchSize; j++ {
			go func(index int) {
				defer wg.Done()
				h, err := cc.Eth.HeaderByNumber(ctx, new(big.Int).Add(i, big.NewInt(int64(index))))
				if err != nil {
					errorCollector.Add(err)
				}
				mu.Lock()
				batch[index] = h
				mu.Unlock()
			}(j)
		}
		wg.Wait()

		if err := errorCollector.Error(); err != nil {
			return err
		}

		for _, h := range batch {
			if err := sendToProcessor(ctx, headers, h); err != nil {
				return err
			}
			logger.Trace("Block Fetched", "block", h.Number)
			count++
			if count == 10000 {
				logger.Info("Fetched 10000 Blocks", "from", new(big.Int).Sub(h.Number, big.NewInt(int64(count-1))), "to", h.Number)
				count = 0
			}
		}
	}
	return nil
}

func lastNodeBlockNumber(ctx context.Context, cc *client.CeloClient) (*big.Int, error) {
	latest, err := cc.Eth.HeaderByNumber(ctx, nil)
	if err != nil {
		return nil, err
	}
	return latest.Number, nil
}
