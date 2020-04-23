package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/log"
)

func ExitOnError(err error) {
	if err != nil {
		log.Crit("Unknown error", "err", err)
	}
}

func PrettyPrint(value interface{}) {
	marshallRes, err := json.MarshalIndent(value, "", " ")
	ExitOnError(err)
	fmt.Println(string(marshallRes))
}

// WaitUntil returns a bool indicating whether the check succeeded before the timeout.
func WaitUntil(ctx context.Context, retryPeriod, timeout time.Duration, check func() bool) bool {
	waitCtx, stop := context.WithTimeout(ctx, timeout)
	defer stop()
	for {
		select {
		case <-time.After(retryPeriod):
			if check() {
				return true
			}
		case <-waitCtx.Done():
			return false
		}
	}
}
