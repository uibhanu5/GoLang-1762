package main

import (
	"fmt"
	"sync"
)

type MutexInt struct {
	v  int
	mu sync.Mutex
}

func main() {
	mi := MutexInt{}

	var wg sync.WaitGroup

	numGoroutines := 5

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			mi.mu.Lock()
			mi.v++
			fmt.Printf("Goroutine incremented, new value: %d\n", mi.v)
			mi.mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Printf("Final shared variable value: %d\n", mi.v)
}
