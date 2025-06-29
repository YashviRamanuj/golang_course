package main

import (
	"fmt"
	util "lec4/util"
	"sync"
)

func notify() {
	var wg sync.WaitGroup
	wg.Add(1) // We have one thing to wait for

	go func() {
		defer wg.Done() // Mark this as done when finished
		util.SendNotification()
	}()

	wg.Wait() // Wait for all tasks to finish
	fmt.Println("All notifications sent. Post published!")
}

func main() {
	notify()
	// OR
	afterPublish()
}

/* Theory:
Wait groups are created using the sync package, and they provide three essential methods: Add(), Done(), and Wait().

Add() is used to add the number of goroutines that need to be waited upon.
Done() is called by each goroutine when it finishes its work, decrementing the internal counter of the wait group.
Wait() is used to block the execution of the goroutine until all the goroutines have called Done().
*/
