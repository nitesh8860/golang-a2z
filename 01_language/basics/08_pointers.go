package main

import "fmt"

func update(num *int) {
	// dereference the pointer
	*num = 30
}

func main() {
	var num int = 5

	// create the pointer variable
	var ptr *int = &num
	// print the address of the variable
	fmt.Println(ptr)
	// print the value stored at pointer address
	fmt.Println(*ptr)

	var number = 55

	// function call
	update(&number)

	fmt.Println("The number is", number)
}
