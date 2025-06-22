package main

import "fmt"

func PrintAny(value interface{}) {
    fmt.Println(value)
}

func main() {
    PrintAny(42)
    PrintAny("hello")
    PrintAny(true)
}