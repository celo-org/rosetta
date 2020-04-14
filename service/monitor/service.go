package monitor

import (
	"context"
	"errors"
	"math/big"

	"sync"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/service"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

type monitorService struct {
	running service.RunningLock
	cc      *client.CeloClient
	db      db.RosettaDB
	logger  log.Logger
}

const srvName = "celo-monitor"

func NewMonitorService(cc *client.CeloClient, db db.RosettaDB) *monitorService {
	return &monitorService{
		cc:     cc,
		db:     db,
		logger: log.New("srv", srvName),
	}
}

// Name retrieves the name of the service, that will be used
// to identify the service in log messages
func (ms *monitorService) Name() string {
	return srvName
}

// Running indicates if the service is currently running
func (ms *monitorService) Running() bool {
	return ms.running.Running()
}

// Start runs the service and blocks until the service finishes,
// returns an error when service failed
func (ms *monitorService) Start(ctx context.Context) error {
	if err := ms.running.EnableOrFail(); err != nil {
		return err
	}
	defer ms.running.Disable()

	startBlock, err := ms.db.LastPersistedBlock(ctx)
	if err != nil {
		return err
	}

	ms.logger.Info("Resuming operation from last persisted  block", "block", startBlock)

	startBlock.Add(startBlock, big.NewInt(1))

	ctx, stopAll := context.WithCancel(ctx)

	var wg sync.WaitGroup
	var errorCollector service.ErrorCollector

	headerCh := make(chan *types.Header)
	changeSetsCh := make(chan *db.BlockChangeSet)

	wg.Add(3)

	// 1st. Listen for Headers
	go func() {
		defer wg.Done()
		err := HeaderListener(ctx, headerCh, ms.cc, ms.logger, startBlock)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				errorCollector.Add(err)
			}
			stopAll()
		}
	}()

	// 2nd. Process Headers
	go func() {
		defer wg.Done()
		err := BlockProcessor(ctx, headerCh, changeSetsCh, ms.cc, ms.db, ms.logger)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				errorCollector.Add(err)
			}
			stopAll()
		}
	}()

	// 3rd. Store Changes into DB
	go func() {
		defer wg.Done()
		err := ProcessChanges(ctx, changeSetsCh, ms.db, ms.logger)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				errorCollector.Add(err)
			}
			stopAll()
		}
	}()

	wg.Wait()
	return errorCollector.Error()
}
