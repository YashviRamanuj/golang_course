# resouces and links
## 01 June, 2025, https://team.shiksha/

### 🎯 Objectives of today's lecture:
Understand Go’s ecosystem and why it’s unique.

Learn basic syntax and usage patterns.

Build your first CLI app – “Hello GoBlogr”.

Explore data structures & error handling via a simple word counter.


### Introduction to Go

🎯 Key Points:
- Go Language: Created by Google in 2007, open-sourced in 2009.
- Modern & pragmatic: Designed for simplicity, concurrency, and fast compile times.
- Compiled & statically typed: Great for performance and reliability.
- Strong community: Used at Google, Uber, Dropbox, and more!

### 🚀 Why Go?
- Speed: Compiled to machine code, quick execution.
- Simplicity: Clean syntax, explicit error handling.
- Concurrency: Goroutines & channels built-in for modern multi-core systems.
- Static binaries: Easy deployment, no runtime dependency hell.

### ⚖️ Strengths & Weaknesses
STRENGTH
- Simplicity & readability
- Concurrency support (goroutines)
- Fast compilation & execution
WEAKNESS
- No generics (until Go 1.18)
- Limited metaprogramming
- Simpler error handling can be verbose



### how to install go

- ubuntu https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-ubuntu-18-04
- window https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-windows-10
- macos - https://www.digitalocean.com/community/tutorials/how-to-install-go-and-set-up-a-local-programming-environment-on-macos

If a version manager is needed, I usually use gvm.

### ✍️ Syntax Basics
Variables (var, :=), types.

Functions (func), parameters, returns.

Control flow (if, for, switch).

@hello.go file


```
<!-- run -->
go run main.go

<!-- build: -->
go build -o goblogr
./goblogr

```



### Data Structures & Error Handling with Simple CLI Word Counter (10 mins)
🧰 Key Data Structures
Int
Float
```
- uint8 unsigned 8-bit integers (0 to 255)
- uint16 unsigned 16-bit integers (0 to 65535)
- uint32 unsigned 32-bit integers (0 to 4294967295)
- uint64 unsigned 64-bit integers (0 to 18446744073709551615)
- int8 signed 8-bit integers (-128 to 127)
- int16 signed 16-bit integers (-32768 to 32767)
- int32 signed 32-bit integers (-2147483648 to 2147483647)
- int64 signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

Floats and complex numbers also come in varying sizes:
- float32 IEEE-754 32-bit floating-point numbers
- float64 IEEE-754 64-bit floating-point numbers
- complex64 complex numbers with float32 real and imaginary parts
- complex128 complex numbers with float64 real and imaginary parts

There are also a couple of alias number types, which assign useful names to specific data types:
- byte alias for uint8
- rune alias for int32
```
Slices: dynamic arrays ([]int, []string).
Maps: key-value pairs (map[string]int).

🚨 Error Handling (No exceptions, explicit errors)
Check err after every operation!

⚡️ Mini Project: Word Counter CLI
Prompt user for input (or read from file – optional).

Count words and print.


### Homework
- Read about GOPATH, GOROOT, packages 

- optional - can read about concurrency vs parallelism