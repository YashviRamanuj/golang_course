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
	counter++
}

func main() {
	for i := 0; i < 1; i++ {
		go increment()
	}
	fmt.Println(counter)
}

// go run --race 1.go
