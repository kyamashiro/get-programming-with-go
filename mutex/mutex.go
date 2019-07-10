package main

import "sync"

func main() {
	var mu sync.Mutex

	mu.Lock()
	// returnする前にmutexをunlockする
	defer mu.Unlock()
}
