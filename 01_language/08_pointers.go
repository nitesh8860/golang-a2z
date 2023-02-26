package main

import "fmt"

func update(num *int) {
	// dereference the pointer
	*num = 30
}

type Weather struct {
	city        string
	temperature int
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

	// struct example
	// create struct variable
	weather := Weather{"California", 20}
	fmt.Println("Initial Weather:", weather)

	// create struct type pointer
	wtr := &weather

	// change value of temperature to 25
	wtr.temperature = 25

	fmt.Println("Updated Weather:", weather)
}
