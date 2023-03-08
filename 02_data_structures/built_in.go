package main

import "fmt"

func main() {
	fmt.Println("-------01 Array-------")
	// create a 2 dimensional array
	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	// access the values of 2d array
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}

	// Slice
	fmt.Println("-------02 Slice-------")
	primeNumbers := []int{2, 3}

	// add elements 5, 7 to the slice
	primeNumbers = append(primeNumbers, 5, 7)
	fmt.Println("Prime Numbers:", primeNumbers)

	// Map
	fmt.Println("-------03 Map-------")
	// create a map
	flowerColor := map[string]string{"Sunflower": "Yellow", "Jasmine": "White", "Hibiscus": "Red"}
	fmt.Println(flowerColor)
	// access value for key Sunflower
	fmt.Println(flowerColor["Sunflower"]) // Yellow
	flowerColor["Sunflower"] = "Dark Yellow"
	fmt.Println(flowerColor["Sunflower"])

	fmt.Println("-------04 Struct-------")
	// create a Struct
	type Rectangle struct {
		length  int
		breadth int
	}

	// declare instance rect1 and defining the struct
	rect := Rectangle{22, 12}

	// access the length of the struct
	fmt.Println("Length:", rect.length)

	// access the breadth of the struct
	fmt.Println("Breadth:", rect.breadth)

	area := rect.length * rect.breadth
	fmt.Println("Area:", area)
}
