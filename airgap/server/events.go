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

package server

import (
	"context"
	"fmt"

	"github.com/celo-org/rosetta/airgap"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var serverCallEventDefinitions = []*airgap.CeloEvent{
	airgap.EpochRewardsDistributedToVoters,
}

func hydrateEvents(srvCtx ServerContext, events []*airgap.CeloEvent) (map[*airgap.CeloEvent]airGapServerEvent, error) {
	abis := make(map[string]*abi.ABI)
	for id, abiFactory := range abiFactoryMap {
		abi, err := abiFactory()
		if err != nil {
			return nil, err
		}
		abis[id] = abi
	}

	mappedEvents := make(map[*airgap.CeloEvent]airGapServerEvent)
	for _, event := range events {
		abi, ok := abis[event.Contract]
		if !ok {
			return nil, fmt.Errorf("Missing abi mapping for %s", event.Contract)
		}
		evt, ok := abi.Events[event.Name]
		if !ok {
			return nil, fmt.Errorf("Missing event data for %s", event.Name)
		}

		mappedEvents[event] = airgapEventFactory(srvCtx, evt, event)
	}
	return mappedEvents, nil
}

func airgapEventFactory(srvCtx ServerContext, evt abi.Event, event *airgap.CeloEvent) airGapServerEvent {
	return func(ctx context.Context, restTopics [][]common.Hash) [][]common.Hash {
		topic0 := evt.ID()
		var topics [][]common.Hash
		topics = append(topics, []common.Hash{topic0})
		topics = append(topics, restTopics...)

		return topics
	}
}
