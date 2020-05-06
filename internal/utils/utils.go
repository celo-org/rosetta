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
	"errors"

	"github.com/ethereum/go-ethereum/core/types"
)

var ErrNotImplemented = errors.New("Not implemented")

var ProxyEventIds = []string{
	"0xab64f92ab780ecbf4f3866f57cee465ff36c89450dcce20237ca7a8d81fb7d13", // ImplementationSet(address)
	"0x50146d0e3c60aa1d17a70635b05494f864e86144a2201275021014fbf08bafe2", // OwnerSet(address)
}

// RemoveProxyLogs filters out logs from Proxy events and returns a list
// of Logs that are safe to Parse using the implementation contract's abi.
func RemoveProxyLogs(logs []*types.Log) []*types.Log {
	safeLogs := make([]*types.Log, 0, len(logs))

	isProxyEvent := func(id string) bool {
		for _, proxyEventId := range ProxyEventIds {
			if id == proxyEventId {
				return true
			}
		}
		return false
	}

	for _, log := range logs {
		if len(log.Topics) > 0 && !isProxyEvent(log.Topics[0].Hex()) {
			safeLogs = append(safeLogs, log)
		}
	}
	return safeLogs
}
