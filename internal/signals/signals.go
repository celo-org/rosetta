package signals

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ethereum/go-ethereum/log"
)

func WatchForExitSignals() <-chan bool {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	log.Info("Press CTRL-C to stop the process")
	go func() {
		sig := <-sigs
		log.Info("Got Signal, Shutting down...", "signal", sig)
		done <- true
	}()

	return done
}
