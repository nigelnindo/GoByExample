package main

import "fmt"

func main() {

	// Go has tyype inference for initialized variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c) // Comma seperate to space logged variables in STDOUT

	// Variables zero valued if declared without initialization
	var e int
	fmt.Println(e)

	// := syntax is short form for declaring and initializing a variable
	// equal to var f string = "apple"
	f := "apple"
	fmt.Println(f)
}
