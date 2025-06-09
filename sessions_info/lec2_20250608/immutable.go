package main

import "fmt"

func main() {
	c := "Sammy says, \"I like to use the `fmt` package\"." // interpreted string literal
	fmt.Println(c)	
	c = `Sammy says, "I like to use the \"fmt\" package\".` // raw string literal
	fmt.Println(c)


	var scaling string 
	fmt.Scanln(&scaling) // Fit
	if string.lower(scaling) == "fit" {

	} else if scaling == "fill" {
	
	}
}
	
