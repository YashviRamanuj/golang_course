package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment(wg *sync.WaitGroup) {
	mu.Lock()
	counter++
	mu.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println(counter)
}
