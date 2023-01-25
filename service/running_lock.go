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

package service

import (
	"errors"
	"sync"
)

var ErrAlreadyRunning = errors.New("Service already running")

type RunningLock struct {
	running bool
	mux     sync.Mutex
}

func (rl *RunningLock) Running() bool {
	rl.mux.Lock()
	defer rl.mux.Unlock()
	return rl.running
}

func (rl *RunningLock) Disable() {
	rl.mux.Lock()
	defer rl.mux.Unlock()
	rl.running = false
}

func (rl *RunningLock) EnableOrFail() error {
	rl.mux.Lock()
	if rl.running {
		rl.mux.Unlock()
		return ErrAlreadyRunning
	}
	rl.running = true
	rl.mux.Unlock()
	return nil
}
