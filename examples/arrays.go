package main

import "fmt"

func main() {
	//creattes an array that holds exactly 5 ints
	//the type of elements and length of array are both part of the array's type
	var a [5]int
	fmt.Println("emp: ", a)

	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	fmt.Println("len: ", len(a))

	//short hand syntax for declaring and initializing array
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl: ", b)

	//the above can also be declared as
	// compiler throws warning that the type can be inferred
	var c [5]int = [5]int{6, 7, 8, 9, 10}
	fmt.Println("lcd: ", c)
}
