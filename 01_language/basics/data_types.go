package main

import "fmt"

func main() {
	var say string = "Hello "
	place := "World! "
	const times = 3000
	fmt.Print(say + place)
	fmt.Println(times)

	// basic data types: int, float, complex, string, bool

	// byte data type = unsigned integer 8 bytes data type = uint8
	var a byte = 100
	var b byte = 'B'

	fmt.Println(a)
	fmt.Println(b)

	// rune data type = signed integer 32 = int32 = UTF-8
	var str string = "ABCD"
	r_array := []rune(str)
	fmt.Printf("Array of rune values for A, B, C and D: %U\n", r_array)

	// declaring and initializing a Unicode character
	emoji := '😀'
	// %T represents the type of the variable num
	fmt.Printf("Data type of %c is %T and the rune value is %U \n", emoji, emoji, emoji)

}
