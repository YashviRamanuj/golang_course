package main


/*
2. Goroutine Leaks
Goroutine leaks occur when goroutines remain blocked indefinitely, consuming resources and potentially leading to memory exhaustion. 

This often happens when goroutines are waiting on channels that are never closed or when exit conditions are not properly managed.


solution:

Ensure that all goroutines have a clear exit strategy. 
Use context cancellation to signal goroutines to terminate gracefully.
*/

ctx, cancel := context.WithCancel(context.Background())
defer cancel()

go func() {
    select {
    case <-ctx.Done():
        // Cleanup and exit
    }
}()

