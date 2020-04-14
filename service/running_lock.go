package service

import "sync"

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
