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
			if count == 100 {
				count = 0
				logger.Info("Stored 100 blocks", "block", changeSet.BlockNumber, "registryUpdates", len(changeSet.RegistryChanges))
			}
		}
	}
}
