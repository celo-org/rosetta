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

package airgap

import (
	"fmt"

	"github.com/celo-org/celo-blockchain/common"
	"github.com/celo-org/kliento/registry"
)

var (
	// Election
	EpochRewardsDistributedToVoters = registerEvent(registry.ElectionContractID.String(), "EpochRewardsDistributedToVoters", []topicParser{addressTopicParser})
	// StableToken
	StableTokenTransferred = registerEvent(registry.StableTokenContractID.String(), "Transfer", []topicParser{addressTopicParser, addressTopicParser})
)

type CeloEvent struct {
	// Name of the abi method
	Name string
	// Registry id of contract where the method is defined
	Contract string

	topicParsers []topicParser
}

func (cm *CeloEvent) String() string { return fmt.Sprintf("%s.%s", cm.Contract, cm.Name) }

func (ce *CeloEvent) DeserializeTopics(values ...[]interface{}) ([][]common.Hash, error) {
	if len(values) > len(ce.topicParsers) {
		return nil, fmt.Errorf("Received %d topics; expected at most %d", len(values), len(ce.topicParsers))
	}
	parsedTopics := make([][]common.Hash, len(values))

	var err error
	for i, topicGroup := range values {
		parsedTopicGroup := make([]common.Hash, len(topicGroup))
		for j, topic := range topicGroup {
			parsedTopicGroup[j], err = ce.topicParsers[i](topic)
			if err != nil {
				return nil, fmt.Errorf("bad argument: idx=%d,%d error=%w", i, j, err)
			}
		}
		parsedTopics[i] = parsedTopicGroup
	}

	return parsedTopics, nil
}
