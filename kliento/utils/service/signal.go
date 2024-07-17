package service

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// WithExitSignals returns a context that will stop whenever
// the process receive a SIGINT / SIGTERM
func WithExitSignals(parentCtx context.Context) context.Context {
	ctx, stop := context.WithCancel(parentCtx)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press CTRL-C to stop the process")
	go func() {
		select {
		case <-parentCtx.Done():
		case sig := <-sigs:
			fmt.Printf("Got Signal, Shutting down... signal: %s", sig.String())
		}
		stop()
	}()
	return ctx
}
