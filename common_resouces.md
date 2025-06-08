# resouces and links
## 01 June, 2025, https://team.shiksha/

### tips for new programmers
- Break problems into tiny pieces when overwhelmed, very imp skill even when you're expienced.
- Exploit student discounts and AI tools as much as possible: github student developer pack -> github copilot free
- Don't replace AI code with yours, try to augment AI in your code. 
- Be in communities and explore and follow your curiosity, ask around people are generally helpful.
- Be nice and kind, it definetly pays in the long run.
- Practise to solve problems as much as possible, brain is a muscle that u can train. DSA, projects, hackathons, open source projects.
- Invest in yourself : time, energy and money as much as possible.


### Advice for these session series

- Do hands on every week along w these sessions.


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
Slices: dynamic arrays ([]int, []string).

Maps: key-value pairs (map[string]int).

🚨 Error Handling (No exceptions, explicit errors)
Check err after every operation!

⚡️ Mini Project: Word Counter CLI
Prompt user for input (or read from file – optional).

Count words and print.


### Book recommendations

- Not necessary but can try if want to go deeper in any topic: https://assets.digitalocean.com/books/how-to-code-in-go.pdf 





