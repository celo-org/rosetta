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
	// max numbers of blocks we allow before starting a subscription
	maxAllowedGap = 50
)

func HeaderListener(ctx context.Context, headers chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock *big.Int) error {
	listener := &listener{
		logger:  logger.New("pipe", "header_listener"),
		ctx:     ctx,
		cc:      cc,
		headers: headers,
	}

	var nextBlock *big.Int = startBlock
	var err error
	for {
		nextBlock, err = listener.syncOldBlocks(nextBlock)
		if err != nil {
			return err
		}

		nextBlock, err = listener.syncNewBlocks(nextBlock)
		if err != nil {
			return err
		}
	}
}

// syncOldBlocks will fetch all blocks until until the distance between node head and last fetched block is less than maxDistance
// returns the next fetchable block
func (listener *listener) syncOldBlocks(firstBlockToFetch *big.Int) (*big.Int, error) {
	maxGap := big.NewInt(maxAllowedGap)

	nextBlock := firstBlockToFetch
	lastBlockInRange, err := listener.lastNodeBlockNumber()
	if err != nil {
		return nil, err
	}

	// While we are behind more than `maxGap` we keep on syncing by range
	for new(big.Int).Sub(lastBlockInRange, nextBlock).Cmp(maxGap) > 0 {
		listener.logger.Info("BatchFetchMode:Start", "start", nextBlock, "end", lastBlockInRange)
		if err = listener.syncBlockRange(nextBlock, lastBlockInRange); err != nil {
			return nil, err
		}
		listener.logger.Info("BatchFetchMode:Finish", "start", nextBlock, "end", lastBlockInRange)

		// next Range
		nextBlock = utils.Inc(lastBlockInRange)
		lastBlockInRange, err = listener.lastNodeBlockNumber()
		if err != nil {
			return nil, err
		}
	}
	return nextBlock, nil
}

// syncNewBlocks will continuously fetch all new blocks on the node
// If consumer is too slow, it will overflow, and exit with the next fetchable block number
func (listener *listener) syncNewBlocks(firstBlockToFetch *big.Int) (*big.Int, error) {
	subscriptionHeaders := make(chan *types.Header, 100)
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

	var lastWrittenBlock *big.Int
	write := func(h *types.Header) error {
		if err := listener.writeHeader(h); err != nil {
			return err
		}
		lastWrittenBlock = h.Number
		return nil
	}

	listener.logger.Info("SubscriptionFetchMode:Start", "start", firstBlockToFetch)
	defer func() {
		listener.logger.Info("SubscriptionFetchMode:Finish", "start", firstBlockToFetch, "end", lastWrittenBlock)
	}()

	// Need to check if we are missing blocks, and if so close the gap
	firstSubscriptionBlock, err := nextHeader()
	if err != nil {
		return nil, err
	}

	// if there's a gap, first fetch all old blocks
	if firstSubscriptionBlock.Number.Cmp(firstBlockToFetch) > 0 {
		endGapBlock := utils.Dec(firstSubscriptionBlock.Number)
		listener.logger.Info("Gap found: closing gap with header subscription", "from", firstBlockToFetch, "to", endGapBlock)
		if err = listener.syncBlockRange(firstBlockToFetch, endGapBlock); err != nil {
			return nil, err
		}
	}

	// write the first block from the subscription
	if err := write(firstSubscriptionBlock); err != nil {
		return nil, err
	}

	// After closing the gap, start reading from the subscription
	for {
		header, err := nextHeader()

		// If overflowed, first flush the subscription
		if err == rpc.ErrSubscriptionQueueOverflow {
			listener.logger.Warn("Header Subscription overflowed. Flushing remaining headers")
			closeOnce.Do(func() { close(subscriptionHeaders) })

			for h := range subscriptionHeaders {
				if nestedErr := write(h); nestedErr != nil {
					return nil, nestedErr
				}
			}

			return utils.Inc(lastWrittenBlock), nil
		}

		if err != nil {
			return nil, err
		}

		if err := write(header); err != nil {
			return nil, err
		}
	}
}

// syncBlockRange will fetch & write all blocks in range [start,end] inclusive
func (listener *listener) syncBlockRange(start, end *big.Int) error {

	// curr <= end => curr.Cmp(end) <= 0
	for curr := new(big.Int).Set(start); curr.Cmp(end) <= 0; curr.Add(curr, utils.Big1) {
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
