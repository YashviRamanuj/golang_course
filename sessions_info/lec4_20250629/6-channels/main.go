package main

import "fmt"

func doSomething(a, b int, result chan int) {
	product := a * b
	result <- product // Send the result into the channel
}

func main() {
	resultChan := make(chan int)

	go doSomething(2, 4, resultChan)
	product := <-resultChan // Wait for and receive the result
	fmt.Printf("Product: %d\n", product)

	// --------------- Closing channel ---------------
	ch := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v) // reads until channel is closed
	}

	// --------------- Real world channel usecase ---------------
	caller()

}

/*
----------------
Channel theory
----------------
Channels are basically like a locked queue (a “First-In-First-Out” data structure)

Channels in Go let different parts of your program (goroutines) safely send data to each other, a bit like passing messages through a pipe.

ch <- 42: "send 42 into channel ch"
<-ch : "receive from channel ch"

----------------

Types of Channels:
1. Unbuffered Channels (default):
- Send blocks until the receiver picks it up.
- Good for synchronization ("don’t move ahead until someone gets my message").

2. Buffered Channels:
- Can "store" a certain number of values without waiting for a receiver.
- Created as: ch := make(chan int, 3) (can hold 3 ints).
- Sender only blocks when buffer is full; receiver blocks if empty.
ch := make(chan int, 2)
ch <- 1 // does not block
ch <- 2 // does not block
// ch <- 3 // would block, because buffer size is 2

fmt.Println(<-ch) // prints 1
fmt.Println(<-ch) // prints 2

----------------
If one goroutine wants to receive data from another, but the sender doesn’t provide it, the channel will make the waiting goroutine wait indefinitely, and vice versa.

----------------
closing channel :close(ch)

Receivers can check if channel is closed:
value, ok := <-ch
if !ok {
    fmt.Println("channel closed")
}

*/
