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

	for {
		sub, err := catchUpWithNode(ctx, headers, newHeaders, cc, logger, startBlock)
		if err != nil {
			return err
		}
		defer sub.Unsubscribe()

		if err := keepUpWithNode(ctx, headers, newHeaders, cc, logger, startBlock, sub); err != nil {
			if err == rpc.ErrSubscriptionQueueOverflow {
				logger.Info("Restarting Header Listener")
				sub.Unsubscribe()
				continue
			}
			return err
		}
		return nil
	}
}

func catchUpWithNode(ctx context.Context, headers, newHeaders chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock *big.Int) (ethereum.Subscription, error) {
	// Start subscription right away so we don't miss any new headers.
	sub, err := cc.Eth.SubscribeNewHead(ctx, newHeaders)
	if err != nil {
		return nil, err
	}

	// Check which block the node is on.
	lastBlock, err := lastNodeBlockNumber(ctx, cc)
	if err != nil {
		return nil, err
	}

	// If the node is ahead, catch-up with it.
	if lastBlock.Cmp(startBlock) > 0 {
		logger.Info("Catching Up...", "start", startBlock, "end", lastBlock)
		if err = fetchHeaderRange(ctx, headers, cc, logger, startBlock, lastBlock); err != nil {
			return nil, err
		}
		logger.Info("Finished Catching Up", "start", startBlock, "end", lastBlock)
	} else {
		logger.Info("Already Caught Up")
	}

	return sub, nil
}

func keepUpWithNode(ctx context.Context, headers chan<- *types.Header, newHeaders <-chan *types.Header, cc *client.CeloClient, logger log.Logger, startBlock *big.Int, sub ethereum.Subscription) error {
	logger.Info("Listening For New Headers")
	defer logger.Info("Stopped Listening For New Headers")
	if len(newHeaders) > 0 {
		logger.Debug("New Headers Were Produced And Buffered During Catch Up", "amount", len(newHeaders))
	}

	for {
		select {
		case err := <-sub.Err():
			return handleSubscriptionOverflow(ctx, headers, newHeaders, logger, startBlock, err)
		case <-ctx.Done():
			return ctx.Err()
		case h := <-newHeaders:
			if h.Number.Int64()%1000 == 0 {
				logger.Info("Fetched 1000 New Blocks", "block", h.Number)
			}
			if err := sendToProcessor(ctx, headers, h); err != nil {
				return err
			}
		}
	}
}

func handleSubscriptionOverflow(ctx context.Context, headers chan<- *types.Header, newHeaders <-chan *types.Header, logger log.Logger, startBlock *big.Int, _err error) error {
	if _err == rpc.ErrSubscriptionQueueOverflow {
		logger.Warn("Max new headers queued during catch-up, subscription will need to restart.")
		if len(newHeaders) > 0 {
			logger.Info("Flushing Subscription Buffer", "length", len(newHeaders))
			var h *types.Header
			for h = range newHeaders {
				if err := sendToProcessor(ctx, headers, h); err != nil {
					return err
				}
			}
			startBlock.Add(h.Number, big.NewInt(1))
		}
	}
	return _err
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
