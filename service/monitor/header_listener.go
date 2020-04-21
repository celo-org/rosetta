package monitor

import (
	"context"
	"math/big"
	"sync"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
)

type listener struct {
	ctx     context.Context
	headers chan<- *types.Header
	cc      *client.CeloClient
	logger  log.Logger
}

const (
	maxAllowedGap = 50
)

func HeaderListener(ctx context.Context, headers chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock *big.Int) error {
	listener := &listener{
		logger:  logger.New("pipe", "header_listener"),
		ctx:     ctx,
		cc:      cc,
		headers: headers,
	}

	var currBlock *big.Int = startBlock
	var err error
	for {
		currBlock, err = listener.syncOldBlocks(startBlock)
		if err != nil {
			return err
		}

		currBlock, err = listener.syncNewBlocks(new(big.Int).Add(utils.Big1, currBlock))
		if err != nil {
			return err
		}
	}
}

// syncOldBlocks will fetch all blocks until until the distance between node head and last fetched block is less than maxDistance
func (listener *listener) syncOldBlocks(startBlock *big.Int) (*big.Int, error) {
	maxGap := big.NewInt(maxAllowedGap)

	currBlock := startBlock
	lastBlock, err := listener.lastNodeBlockNumber()
	if err != nil {
		return nil, err
	}

	// While we are behind more than `maxGap` we keep on syncing by range
	for new(big.Int).Sub(lastBlock, currBlock).Cmp(maxGap) > 0 {
		listener.logger.Info("Fetching Old Blocks", "start", currBlock, "end", lastBlock)
		if err = listener.syncBlockRange(currBlock, lastBlock); err != nil {
			return nil, err
		}
		listener.logger.Info("Finished Catching Up", "start", currBlock, "end", lastBlock)

		// next Range
		currBlock = new(big.Int).Add(lastBlock, big.NewInt(1))
		lastBlock, err = listener.lastNodeBlockNumber()
		if err != nil {
			return nil, err
		}
	}
	return currBlock, nil
}

func (listener *listener) syncNewBlocks(startBlock *big.Int) (*big.Int, error) {
	listener.logger.Info("Listening For New Headers")
	defer listener.logger.Info("Stopped Listening For New Headers")

	subscriptionHeaders := make(chan *types.Header, 1000)
	var closeOnce sync.Once
	defer closeOnce.Do(func() { close(subscriptionHeaders) })

	sub, err := listener.cc.Eth.SubscribeNewHead(listener.ctx, subscriptionHeaders)
	if err != nil {
		return nil, err
	}
	defer sub.Unsubscribe()

	nextHeader := func() (*types.Header, error) {
		select {
		case err := <-sub.Err():
			return nil, err
		case <-listener.ctx.Done():
			return nil, listener.ctx.Err()
		case h := <-subscriptionHeaders:
			return h, nil
		}
	}

	// Need to check if we are missing blocks, and if so close the gap
	firstNewBlock, err := nextHeader()
	if err != nil {
		return nil, err
	}

	// if there's a gap, first fetch all old blocks
	endGapBlock := new(big.Int).Sub(firstNewBlock.Number, big.NewInt(1))
	if endGapBlock.Cmp(startBlock) > 0 {
		listener.logger.Info("Closing gap with header subscription", "from", startBlock, "to", endGapBlock)
		if err = listener.syncBlockRange(startBlock, endGapBlock); err != nil {
			return nil, err
		}
	}

	// write the first block from the subscription
	if err := listener.writeHeader(firstNewBlock); err != nil {
		return nil, err
	}

	// After closing the gap, start reading from the subscription
	for {
		header, err := nextHeader()

		// If overflowed, first flush the subscription
		if err == rpc.ErrSubscriptionQueueOverflow {
			listener.logger.Error("Header Subscription overflowed. Flushing remaining header")
			closeOnce.Do(func() { close(subscriptionHeaders) })

			var lastFetchedBlock *big.Int
			for h := range subscriptionHeaders {
				if nestedErr := listener.writeHeader(h); nestedErr != nil {
					return nil, nestedErr
				}
				lastFetchedBlock = h.Number
			}
			return lastFetchedBlock, nil
		}

		if err != nil {
			return nil, err
		}

		if err := listener.writeHeader(header); err != nil {
			return nil, err
		}
	}
}

// syncBlockRange will fetch & write all blocks in range [start,end] incluse
func (listener *listener) syncBlockRange(start, end *big.Int) error {

	for curr := new(big.Int).Set(start); end.Cmp(curr) > 0; curr.Add(curr, utils.Big1) {
		h, err := listener.cc.Eth.HeaderByNumber(listener.ctx, curr)
		if err != nil {
			return err
		}
		if err = listener.writeHeader(h); err != nil {
			return err
		}
	}
	return nil
}

func (listener *listener) lastNodeBlockNumber() (*big.Int, error) {
	latest, err := listener.cc.Eth.HeaderByNumber(listener.ctx, nil)
	if err != nil {
		return nil, err
	}
	return latest.Number, nil
}

func (listener *listener) writeHeader(h *types.Header) error {
	select {
	case <-listener.ctx.Done():
		return listener.ctx.Err()
	case listener.headers <- h:
	}
	return nil
}
