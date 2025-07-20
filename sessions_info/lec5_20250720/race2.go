package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	mu      sync.Mutex
)

func increment() {
	mu.Lock()
	counter++
	mu.Unlock()
}

func main() {
	for i := 0; i < 1; i++ {
		go increment()
	}
	fmt.Println(counter)
}
