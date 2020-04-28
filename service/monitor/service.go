package monitor

import (
	"context"

	"github.com/celo-org/rosetta/celo/client"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/celo-org/rosetta/service"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"
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
	startBlock = utils.Inc(startBlock)

	// headerCh is buffered so that sends from the header
	// listener never have to wait for the preceding block
	// to finish processing.
	headerCh := make(chan *types.Header, 10)
	changeSetsCh := make(chan *db.BlockChangeSet)

	group, ctx := errgroup.WithContext(ctx)
	group.Go(func() error { return HeaderListener(ctx, headerCh, ms.cc, ms.logger, startBlock) })
	group.Go(func() error { return BlockProcessor(ctx, headerCh, changeSetsCh, ms.cc, ms.db, ms.logger) })
	group.Go(func() error { return ProcessChanges(ctx, changeSetsCh, ms.db, ms.logger) })
	return group.Wait()
}
