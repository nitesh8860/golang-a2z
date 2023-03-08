// Program to illustrate call by reference

package main

import "fmt"

// call by value
func callByValue(num int) {

	fmt.Println(num)

}

// call by reference
func callByReference(num *int) {

	fmt.Println(*num) // dereference

}

func main() {

	var number int = 30

	// passing value
	callByValue(number)

	// passing a reference (address)
	callByReference(&number)

}
