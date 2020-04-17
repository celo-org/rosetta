package monitor

import (
	"context"
	"math/big"
	"sync"

	"github.com/celo-org/rosetta/celo/client"
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

	lastFetchedBlock, err := listener.syncOldBlocks(startBlock)
	if err != nil {
		return err
	}

	nextBlock := new(big.Int).Add(big.NewInt(1), lastFetchedBlock)
	if err := listener.syncNewBlocks(nextBlock); err != nil {
		return err
	}

	return nil
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
		if err = listener.syncBlockRange(startBlock, lastBlock); err != nil {
			return nil, err
		}
		listener.logger.Info("Finished Catching Up", "start", startBlock, "end", lastBlock)

		currBlock = lastBlock
		lastBlock, err = listener.lastNodeBlockNumber()
		if err != nil {
			return nil, err
		}
	}
	return currBlock, nil
}

func (listener *listener) syncNewBlocks(startBlock *big.Int) error {
	listener.logger.Info("Listening For New Headers")
	defer listener.logger.Info("Stopped Listening For New Headers")

	subscriptionHeaders := make(chan *types.Header)
	var closeOnce sync.Once
	defer closeOnce.Do(func() { close(subscriptionHeaders) })

	sub, err := listener.cc.Eth.SubscribeNewHead(listener.ctx, subscriptionHeaders)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	// After stating the subscription we fetch block gap between lastFetched and node's head
	// by doing it after, we assure we will get all blocks in order
	lastBlock, err := listener.lastNodeBlockNumber()
	if err != nil {
		return err
	}
	if lastBlock.Cmp(startBlock) > 0 {
		listener.logger.Info("Closing gap with header subscription", "from", startBlock, "to", lastBlock)
		if err = listener.syncBlockRange(startBlock, lastBlock); err != nil {
			return err
		}
	}

	// After closing the gap, start reading from the subscription
	for {
		select {
		case err := <-sub.Err():
			if err == rpc.ErrSubscriptionQueueOverflow {
				listener.logger.Warn("Header Subscription overflowed. Flushing remaining header & restart")
				closeOnce.Do(func() { close(subscriptionHeaders) })
				for h := range subscriptionHeaders {
					if err := listener.writeHeader(h); err != nil {
						return err
					}
				}
				return nil
			}
			return err
		case <-listener.ctx.Done():
			return listener.ctx.Err()
		case h := <-subscriptionHeaders:
			if h.Number.Int64()%1000 == 0 {
				listener.logger.Info("Fetched 1000 New Blocks", "block", h.Number)
			}
			if err := listener.writeHeader(h); err != nil {
				return err
			}
		}
	}
}

// syncBlockRange will fetch & write all blocks in range [start,end] incluse
func (listener *listener) syncBlockRange(start, end *big.Int) error {

	for curr := new(big.Int).Set(start); curr.Cmp(end) <= 0; curr.Add(curr, big.NewInt(1)) {
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
