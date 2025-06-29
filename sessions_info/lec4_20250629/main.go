//go:build practical
// +build practical

package main

import (
	"fmt"
	"sync"
	"time"
)

func sendEmail(email string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond) // simulate sending
	results <- fmt.Sprintf("Email sent to %s", email)
}

func main() {
	emails := []string{"a@example.com", "b@example.com", "c@example.com"}
	var wg sync.WaitGroup
	results := make(chan string, len(emails))

	for _, email := range emails {
		wg.Add(1)
		go sendEmail(email, &wg, results)
	}

	wg.Wait()
	close(results)

	for res := range results {
		fmt.Println(res)
	}
}
