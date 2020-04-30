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

	"github.com/celo-org/rosetta/db"
	"github.com/ethereum/go-ethereum/log"
)

func ProcessChanges(ctx context.Context, changes <-chan *db.BlockChangeSet, dbWriter db.RosettaDBWriter, logger log.Logger) error {
	logger = logger.New("pipe", "persister")
	var count uint
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case changeSet := <-changes:
			if err := dbWriter.ApplyChanges(ctx, changeSet); err != nil {
				return err
			}

			count++
			if count == 1000 {
				count = 0
				logger.Info("Stored 1000 blocks", "block", changeSet.BlockNumber, "registryUpdates", len(changeSet.RegistryChanges))
			}
		}
	}
}
