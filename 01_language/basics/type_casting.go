package main

import "fmt"

func main() {

	var intValue int = 2

	// Explicit Typecasting
	var floatValue float32 = float32(intValue)

	fmt.Printf("Integer Value is %d\n", intValue)
	fmt.Printf("Float Value is %f", floatValue)

}

// Go does not support implicit typecasting
// e.g. var number int = 4.34 - Go will throw error
