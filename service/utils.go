package service

import (
	"context"
	"errors"
	"strings"
	"sync"
	"time"
)

type ErrorCollector struct {
	errors []error
	mu     sync.Mutex
}

func NewErrorCollector() *ErrorCollector {
	return &ErrorCollector{}
}

func (ec *ErrorCollector) Add(err error) {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	ec.errors = append(ec.errors, err)
}

func (ec *ErrorCollector) Error() error {
	ec.mu.Lock()
	defer ec.mu.Unlock()
	if len(ec.errors) == 0 {
		return nil
	}

	if len(ec.errors) == 1 {
		return ec.errors[0]
	}

	var errStr strings.Builder
	errStr.WriteString("ErrorGroup: ")
	for _, err := range ec.errors {
		errStr.WriteString(err.Error())
		errStr.WriteString("\n")
	}

	return errors.New(errStr.String())
}

type delayedService struct {
	srv     Service
	delay   time.Duration
	running RunningLock
}

func WithDelay(srv Service, delay time.Duration) Service {
	return &delayedService{
		srv:   srv,
		delay: delay,
	}
}

func (ds *delayedService) Name() string {
	return ds.srv.Name()
}

func (ds *delayedService) Running() bool {
	return ds.running.Running()
}

func (ds *delayedService) Start(ctx context.Context) error {
	select {
	case <-time.After(ds.delay):
	case <-ctx.Done():
		return nil
	}

	return ds.srv.Start(ctx)
}

type ServieFactory func() (Service, error)

type lazyService struct {
	factory ServieFactory
	srv     Service
	name    string
}

func LazyService(name string, factory func() (Service, error)) *lazyService {
	return &lazyService{
		name:    name,
		factory: factory,
	}
}

func (ls *lazyService) Name() string {
	if ls.srv == nil {
		return ls.name
	}
	return ls.srv.Name()
}

func (ls *lazyService) Running() bool {
	if ls.srv == nil {
		return false
	}
	return ls.srv.Running()
}

func (ls *lazyService) Start(ctx context.Context) error {
	// TODO improve
	if ls.srv != nil {
		return ErrAlreadyRunning
	}

	srv, err := ls.factory()
	if err != nil {
		return err
	}

	ls.srv = srv

	return ls.srv.Start(ctx)
}
