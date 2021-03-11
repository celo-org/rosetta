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

package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/celo-org/celo-blockchain/log"
	"github.com/coinbase/rosetta-sdk-go/fetcher"
)

func ExitOnError(err error) {
	if err != nil {
		log.Crit("Unknown error", "err", err)
	}
}

func ExitOnFetcherError(err *fetcher.Error) {
	if err != nil {
		ExitOnError(err.Err)
	}
}

func PrettyPrint(value interface{}) {
	marshallRes, err := json.MarshalIndent(value, "", " ")
	ExitOnError(err)
	fmt.Println(string(marshallRes))
}

// WaitUntil returns a bool indicating whether the check succeeded before the timeout.
func WaitUntil(retryPeriod, timeout time.Duration, check func() bool) bool {
	ctx, stop := context.WithTimeout(context.Background(), timeout)
	defer stop()
	for {
		select {
		case <-time.After(retryPeriod):
			if check() {
				return true
			}
		case <-ctx.Done():
			return false
		}
	}
}
