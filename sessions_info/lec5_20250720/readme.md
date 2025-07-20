1. race
2. leak
3. deadlock
4. M:N scheduler
5. channel internals
6. goroutine stack management
7. memory model
8. context propagation
9. performance pitfalls & optimization tips
10. debugging & profiling concurrency


# 🧠 Go Scheduler Internals – Key Concepts

- **M:N scheduler**
  - G (goroutine) → lightweight thread
  - M (machine) → OS thread
  - P (processor) → context to run goroutines
  - GOMAXPROCS → number of active Ps (usually = CPU cores)

- **How it works**
  - Each P has a local run queue (list of runnable Gs)
  - Global run queue for overflow goroutines
  - M needs a P to run Go code
  - If goroutine blocks (I/O, mutex), M can yield P to another M

- **Work stealing**
  - If P's queue is empty → steal half Gs from random other P
  - Helps keep CPUs busy
  - Prevents one P from being idle while others are overloaded

- **Preemption**
  - Old Go: only cooperatively (safe-points like function calls)
  - Go 1.14+: async preemption added
  - Sysmon thread monitors long-running goroutines
  - Sends signal to force goroutine to yield if hogging CPU
  - Ensures fairness, helps garbage collector pause all Gs

- **Implications**
  - Tight infinite loops don’t freeze program → get preempted
  - Goroutines can be paused and resumed anywhere safely
  - Developers don’t control scheduling – runtime handles it

- **Debugging tools**
  - `runtime.Gosched()` → manually yield (rarely needed)
  - `go tool trace` → visualize goroutine activity on timeline
  - `runtime.NumGoroutine()` → current goroutine count


# 📬 Go Channels Internals – Key Concepts

- **What is a channel?**
  - Pipe for communication between goroutines
  - Built-in queue + lock + wait lists
  - Type-safe (e.g. `chan int`, `chan string`)

- **Buffered vs Unbuffered**
  - **Unbuffered (`make(chan T)`)**
    - No storage
    - Send/receive must meet at same time → handshake
    - Strong sync: sender waits for receiver
  - **Buffered (`make(chan T, N)`)**
    - Circular buffer of size N
    - Send succeeds if buffer has space
    - Receive succeeds if buffer has values
    - Weaker sync: send may complete before receive

- **Blocking behavior**
  - Send blocks if: buffer full or no receiver (unbuffered)
  - Receive blocks if: buffer empty or no sender (unbuffered)
  - `select` helps avoid blocking forever

- **Channel internals (`hchan` struct)**
  - Circular buffer for storage
  - `sendq` and `recvq` = lists of waiting goroutines
  - Mutex used internally to protect access
  - Direct memory handoff in unbuffered (sender writes to receiver’s stack)

- **Select statement**
  - Wait on multiple channels at once
  - Runtime randomizes case order → fairness
  - First ready case runs, others ignored
  - Blocks if no case is ready (unless `default` is present)

- **Closing and nil channels**
  - Send on closed channel → panic
  - Receive from closed → gets zero value + `ok=false`
  - Nil channel → blocks forever on send/receive
  - Useful trick: disable `select` case by using `nil`

- **Memory model**
  - **Unbuffered** send → happens-before receive
  - **Buffered** send → happens-before that value’s receive
  - Unbuffered = sync point; buffered = async queue

- **Gotchas**
  - Only sender should `close()`
  - `range ch` blocks unless channel is closed
  - Channel ops have some overhead (lock + wake-up)


# 🧵 Goroutine Stack Management – Key Concepts

- **Starts small**
  - New goroutine stack ~2 KB (vs. ~1 MB for OS thread)
  - Saves memory when running many goroutines

- **Grows automatically**
  - If out of space → runtime allocates bigger stack
  - Copies old stack → resumes execution
  - Shrinks during GC if unused

- **No shared stack**
  - Each goroutine has its own stack
  - Local vars in different goroutines are fully isolated

- **Performance**
  - Stack resizing is efficient but not free
  - Avoid huge local vars or deep recursion

- **Debugging**
  - No manual control over stack size
  - GC & runtime handle stack pointers safely


# 🧠 Go Memory Model – Key Concepts

- **Happens-before**
  - A → B if A’s writes visible to B
  - Critical for avoiding races

- **Data race**
  - Two goroutines access same var, at least one writes, no sync
  - Result: non-deterministic, buggy behavior

- **Synchronization tools**
  - Channels: send happens-before receive
  - Mutexes: Unlock happens-before next Lock
  - `sync/atomic`: low-level, for single vars
  - WaitGroups: Wait happens-after Done

- **Race Detector**
  - Run with `-race` flag
  - Finds unsynchronized access, shows code locations

- **Rules of thumb**
  - No shared writes without sync
  - Channels and mutexes define memory visibility


# 🧭 Context Propagation – Key Concepts

- **What is context?**
  - Carries cancel signal + deadline + metadata
  - Passed down function calls, across goroutines

- **Types**
  - `Background()` → root context
  - `WithCancel`, `WithTimeout`, `WithDeadline` → create child ctx

- **Usage pattern**
  - Pass `ctx` to functions/goroutines
  - Use `select { case <-ctx.Done() }` to exit on cancel/timeout

- **Why useful**
  - Cancel multiple goroutines at once
  - Avoids goroutine leaks
  - Clean shutdowns

- **Common pattern**
  - `ctx, cancel := context.WithCancel(...)`
  - `defer cancel()` in main function or handler


# 🐢 Performance Pitfalls & Optimization Tips

- **Too many goroutines**
  - CPU-bound: too many = preemption overhead
  - Fix: limit to GOMAXPROCS, use worker pool

- **Too fine-grained concurrency**
  - Small tasks → goroutine overhead > benefit
  - Fix: batch work, coarse-grain parallelism

- **Channel overhead**
  - Channels ≠ zero cost
  - Fix: reduce chatter, combine messages, or use atomics

- **Contention**
  - One lock/channel hit by all → bottleneck
  - Fix: shard data, use atomic ops, reduce shared access

- **Unbounded spawning**
  - Tasks arrive faster than finish → memory blow-up
  - Fix: use buffered channels, semaphores, bounded pools

- **Benchmark**
  - Always compare: sequential vs. concurrent
  - Measure CPU, goroutines, time


# 🚨 Deadlocks, Livelocks, Leaks – Key Concepts

- **Deadlock**
  - All goroutines waiting, no progress
  - Often cyclic dependency (e.g. send→receive loops)

- **Livelock**
  - Goroutines active, but spinning without progress
  - Often due to retry loops, backoff without randomness

- **Goroutine leaks**
  - Goroutines waiting forever (e.g. on channel never closed)
  - `runtime.NumGoroutine()` keeps growing

- **Fixes**
  - Use context to signal cancellation
  - Always close channels or check for done
  - Avoid infinite `range ch` unless sure channel will close

- **Debugging**
  - `pprof` → see blocked/stuck goroutines
  - Use logs + trace to find where goroutines are stuck


# 🔍 Debugging & Profiling Concurrency

- **Race Detector**
  - `go run -race`, `go test -race`
  - Catches unsynced access

- **pprof**
  - `/debug/pprof/goroutine` → goroutine dumps
  - Profiles: CPU, memory, blocking, mutex

- **Trace**
  - `go tool trace` → visual timeline
  - Shows scheduling, blocks, GC, syscalls

- **Block & Mutex profiling**
  - `SetBlockProfileRate(1)` → where goroutines block
  - `SetMutexProfileFraction(1)` → contested locks

- **Best practices**
  - Monitor `NumGoroutine()`
  - Add timeouts, context
  - Simulate edge cases in tests



# resources 
- https://dev.to/l_walid/concurrency-gotchas-in-go-25lc