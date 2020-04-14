package monitor

import (
	"context"
	"math/big"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

func HeaderListener(ctx context.Context, headers chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock *big.Int) error {
	logger = logger.New("pipe", "header_listener")

	sub, err := cc.Eth.SubscribeNewHead(ctx, headers)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()

	lastBlock, err := lastNodeBlockNumber(ctx, cc)
	if err != nil {
		return err
	}

	if lastBlock.Cmp(startBlock) > 0 {
		logger.Info("Start fetching old blocks", "start", startBlock, "end", lastBlock)
		if err = fetchHeaderRange(ctx, headers, cc, logger, startBlock, lastBlock); err != nil {
			return err
		}
		logger.Info("Finished fetching old blocks", "start", startBlock, "end", lastBlock)
	}

	select {
	case err := <-sub.Err():
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func fetchHeaderRange(ctx context.Context, headers chan<- *types.Header, cc *client.CeloClient, logger log.Logger, startBlock, endBlock *big.Int) error {

	count := 0

	// TODO(Alec): Add concurrency here

	for i := new(big.Int).Set(startBlock); i.Cmp(endBlock) <= 0; i.Add(i, big.NewInt(1)) {
		h, err := cc.Eth.HeaderByNumber(ctx, i)
		if err != nil {
			return err
		}

		logger.Trace("Block Fetched", "block", i)
		count++
		if count == 10000 {
			logger.Info("Fetched 10000 Blocks", "from", new(big.Int).Sub(i, big.NewInt(int64(count-1))), "to", i)
			count = 0
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		case headers <- h:
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
