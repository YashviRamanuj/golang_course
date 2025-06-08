package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println("Enter a sentence:")

    var input string
    fmt.Scanln(&input)

    words := strings.Fields(input)
    fmt.Println("Word count:", len(words))
}
