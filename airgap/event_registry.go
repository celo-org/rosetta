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
	"strings"
)

var eventRegistry = make(map[string]map[string]*CeloEvent)

// FromString returns the CeloEvent that matches the given string
// Events are represented as "Contract.EventName"
func EventFromString(celoEventStr string) (*CeloEvent, error) {
	parts := strings.Split(celoEventStr, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("Invalid event string: %s", celoEventStr)
	}
	m, ok := eventRegistry[parts[0]]
	if !ok {
		return nil, fmt.Errorf("Invalid event string: %s", celoEventStr)
	}
	cm, ok := m[parts[1]]
	if !ok {
		return nil, fmt.Errorf("Invalid event string: %s", celoEventStr)
	}
	return cm, nil
}

func registerEvent(contract string, name string, topicParsers []topicParser) *CeloEvent {
	ce := &CeloEvent{
		Name:         name,
		Contract:     contract,
		topicParsers: topicParsers,
	}
	// register the event in the eventRegistry
	if eventRegistry[contract] == nil {
		eventRegistry[contract] = make(map[string]*CeloEvent)
	}
	eventRegistry[contract][name] = ce
	return ce
}
