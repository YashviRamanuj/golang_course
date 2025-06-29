// Suppose you want to process a list of numbers by squaring them, using multiple workers:

package main

import "fmt"

func worker(id int, in <-chan int, out chan<- int) {
	for num := range in {
		fmt.Printf("Worker %d processing %d\n", id, num)
		out <- num * num
	}
}

func caller() {
	in := make(chan int)
	out := make(chan int)

	// start 3 workers
	for i := 0; i < 3; i++ {
		go worker(i, in, out)
	}

	nums := []int{1, 2, 3, 4, 5}
	go func() {
		for _, n := range nums {
			in <- n
		}
		close(in)
	}()

	for range nums {
		fmt.Println(<-out)
	}
}
