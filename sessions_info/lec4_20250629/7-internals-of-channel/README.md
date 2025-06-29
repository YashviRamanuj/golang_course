
## Go Channels: Core Implementation Crux

### 1. `hchan` – Channel’s Heart  
```go
type hchan struct {
  qcount   uint           // items in buffer
  dataqsiz uint           // buffer size
  buf      unsafe.Pointer // circular buffer
  elemsize uint16         // size of element
  elemtype *_type         // element’s Go type
  sendx    uint           // next write index
  recvx    uint           // next read index
  recvq    waitq          // waiting receivers
  sendq    waitq          // waiting senders
  closed   uint32         // closed flag
  lock     mutex          // protects entire hchan
}
```

2. Wait Queues & `sudog`

- `waitq`: doubly-linked list of blocked goroutines on send/receive.

- `sudog`: per-goroutine node holding:
    - `g` (goroutine pointer)
    - `elem` (for direct handoff)
    - queue links


3. Buffered Send

```go

func send(ch *hchan, elem unsafe.Pointer) {
  ch.lock.Lock()
  if ch.qcount < ch.dataqsiz {
    copyToBuf(ch, elem)          // buf[sendx]=elem
    enqueueIfWaiting(&ch.recvq)  // wake one recv
  } else {
    enqueueSender(&ch.sendq)     // park in sendq
  }
  ch.lock.Unlock()
}
```

4. Buffered Receive
```go
func recv(ch *hchan, out *unsafe.Pointer) {
  ch.lock.Lock()
  if ch.qcount > 0 {
    *out = copyFromBuf(ch)       // read buf[recvx]
    enqueueIfWaiting(&ch.sendq)  // wake one sender
  } else {
    enqueueReceiver(&ch.recvq)   // park in recvq
  }
  ch.lock.Unlock()
}
```

5. Unbuffered / Direct Handoff
- Fast path: if sendq or recvq nonempty, skip buffer entirely.
- Direct copy: sender writes into receiver’s stack (sudog.elem), wakes it.
- Advantage: no extra lock/unlock or buffer copy → lower latency.

6. Parking & Waking
- Park: wrap goroutine in sudog, push to sendq/recvq, unblock scheduler.
- Wake: dequeue one sudog, set its state runnable, scheduler resumes it.

7. Synchronization Guarantees
- Unbuffered: send & receive rendezvous.
- Buffered: send blocks when buffer full; receive blocks when empty.
- No explicit user locks needed—Go runtime’s mutex and queues handle it.

8. Key Takeaways
- Channels ≠ shared memory: share data by communication.
- hchan + waitq + sudog form the internal plumbing.
- Optimizations (direct handoff) improve throughput.


Understanding internals guides correct usage and performance tuning.

```go
// Example: direct handoff (pseudo-logic)
if ch.sendq.notEmpty() {
  sg := dequeue(&ch.sendq)      // receiver’s sudog
  *sg.elem = copy(elem)         // place into receiver’s stack
  wake(goroutineOf(sg))
} else {
  // fallback to buffered send
}
```


reference: https://medium.com/@ravikumar19997/exploring-the-depths-of-golang-channels-a-comprehensive-guide-53e1a97cafe6
