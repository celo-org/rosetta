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

	"github.com/celo-org/celo-blockchain/core/types"
	"github.com/celo-org/celo-blockchain/log"
	"github.com/celo-org/kliento/client"
	"github.com/celo-org/rosetta/db"
	"github.com/celo-org/rosetta/internal/utils"
	"github.com/celo-org/rosetta/service"
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
	err = group.Wait()
	if err == context.Canceled {
		return nil
	}
	return err
}
