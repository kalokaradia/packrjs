package watcher

import (
	"sync"
	"time"
)

var (
	debounceTimer *time.Timer
	mu            sync.Mutex
)

// StartDebounce starts a debounce timer
func StartDebounce(delay time.Duration, callback func()) {
	mu.Lock()
	defer mu.Unlock()
	
	if debounceTimer != nil {
		debounceTimer.Stop()
	}
	debounceTimer = time.AfterFunc(delay, callback)
}

// StopDebounce stops the current debounce timer
func StopDebounce() {
	mu.Lock()
	defer mu.Unlock()
	
	if debounceTimer != nil {
		debounceTimer.Stop()
	}
}
