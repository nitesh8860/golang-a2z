package main

import "fmt"

// recover function to handle panic
func handlePanic() {

	// detect if panic occurs or not
	a := recover()
	if a != nil {
		fmt.Println("recovering from panic, panic msg is: ", a)
	}

}

func main() {
	defer fmt.Println("One") // 3rd

	defer handlePanic() // recover

	defer fmt.Println("Two")   // 2nd
	defer fmt.Println("Three") // 1st

	panic("Ending the program because panic was called") // panic

}
