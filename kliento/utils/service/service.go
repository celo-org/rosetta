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

package service

import (
	"context"
	"errors"
	"sync"
)

var ErrAlreadyRunning = errors.New("Service already running")

type Service interface {
	// Name retrieves the name of the service, that will be used
	// to identify the service in log messages
	Name() string

	// Running indicates if the service is currently running
	Running() bool

	// Start runs the service and blocks until the service finishes,
	// returns an error when service failed
	Start(ctx context.Context) error
}

// monitor HTTP Service
// Geth Node Service
// Monitor Service (package monitor)

func RunServices(ctx context.Context, services ...Service) error {
	sm := NewServiceManager(ctx)
	for _, srv := range services {
		sm.Add(srv)
	}
	return sm.Wait()
}

type ServiceManager struct {
	wg             sync.WaitGroup
	errorCollector ErrorCollector
	ctx            context.Context
	stopAll        context.CancelFunc
}

func NewServiceManager(ctx context.Context) *ServiceManager {

	ctx, stopAll := context.WithCancel(ctx)
	return &ServiceManager{
		ctx:     ctx,
		stopAll: stopAll,
	}

}

func (sm *ServiceManager) Add(srv Service) {
	sm.wg.Add(1)
	go func() {
		defer sm.wg.Done()
		err := srv.Start(sm.ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) {
				sm.errorCollector.Add(err)
			}
			sm.stopAll()
		}
	}()
}

func (sm *ServiceManager) Wait() error {
	sm.wg.Wait()
	return sm.errorCollector.Error()
}
