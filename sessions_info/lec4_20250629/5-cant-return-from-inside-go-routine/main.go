package main

import (
	"fmt"
	"time"
)

func doSomething(a int, b int) int {
	result := a * b
	time.Sleep(2 * time.Second)

	fmt.Println("Did something")
	return result
}

func doSomethingAndReturnResultInSharedVariable(a int, b int, result *int) {
	time.Sleep(2 * time.Second)
	*result = a * b
}

func main() {
	//  regular functions return values this way:
	result := doSomething(2, 4)
	fmt.Println("regular function result: ", result)

	// -----------------

	// go routines cannot do that!
	// result = go doSomething(2, 4)
	// fmt.Println(result)

	// -----------------

	// What if You Try to Use Shared Variables?
	var a int
	fmt.Println("share variable value before: ", a)
	go doSomethingAndReturnResultInSharedVariable(2, 4, &a)
	fmt.Println("share variable value after: ", a)
	time.Sleep(3 * time.Second)
	fmt.Println("share variable value after sleep: ", a)

	/*
		Problems with Shared Variables:

		Race Conditions:
		- Both goroutine and main can access result at the same time, possibly before it’s set!

		Synchronization:
		- How do you know when the goroutine has set the value? What if main checks before it’s ready?

		Go will warn you: “Detected data race!”
	*/
}
