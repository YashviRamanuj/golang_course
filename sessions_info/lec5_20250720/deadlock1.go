package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    var mu1, mu2 sync.Mutex

    wg.Add(2)
    go func() {
        defer wg.Done()
        mu1.Lock()
        fmt.Println("goroutine 1 acquired lock 1")
        mu2.Lock()
        fmt.Println("goroutine 1 acquired lock 2")
    }()
    go func() {
        defer wg.Done()
        mu2.Lock()
        fmt.Println("goroutine 2 acquired lock 2")
        mu1.Lock()
        fmt.Println("goroutine 2 acquired lock 1")
    }()
    wg.Wait()
}


/* 

How can we handle deadlocks?

Acquiring locks in a consistent order
Avoiding deadlocks can be facilitated by standardizing the sequence in which all goroutines obtain locks. In the introductory example given previously, for instance, a deadlock can be avoided if both goroutines consistently acquire the lock on mu1 before the lock on mu2.


Using Timeout
Timeouts can be used in lock acquisition attempts as another means of breaking deadlocks. 
In this approach, a goroutine might presume a stalemate has occurred and take necessary action, such as releasing the locks it now has and retrying at a later time, if it has been unable to acquire a lock within a particular length of time.



Using a channel-based approach
If you need to control who has access to a shared resource but don’t want to deal with locks, use channels instead. 
Requests for resources are communicated over channels between goroutines; the receiving goroutine then decides whether or not to grant the request.
Locks aren’t necessary thanks to this method, and deadlocks are less likely to occur as a result.


Using a dedicated goroutine
A dedicated goroutine can be used to manage access to shared resources, eliminating the need for several goroutines to compete for locks on the same objects. 
This method, getting locks and preventing deadlocks is handled by a single goroutine.

more: https://medium.com/@trinad536/deadlocks-in-go-f4ae0ecd05f6
*/