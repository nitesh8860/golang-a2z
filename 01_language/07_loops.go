// Program to print the first 5 natural numbers

package main

import "fmt"

func main() {

	// for loop
	fmt.Println("---------For loop---------")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	// While loop
	number := 1
	fmt.Println("---------While loop---------")
	for number <= 5 {
		fmt.Println(number)
		number++
	}

	// Do While loop
	number = 1
	fmt.Println("---------Do While loop---------")

	// loop that runs infinitely
	for {

		// condition to terminate the loop
		if number > 5 {
			break
		}

		fmt.Printf("%d\n", number)
		number++

	}

	// Range loops
	numbers := [5]int{21, 24, 27, 30, 33}
	fmt.Println("---------loop over range---------")
	// use range to iterate over the elements of array
	for index, item := range numbers {
		fmt.Printf("numbers[%d] = %d \n", index, item)
	}

	// loop over maps
	fmt.Println("---------loop over maps---------")
	// create a map
	subjectMarks := map[string]float32{"Java": 80, "Python": 81, "Golang": 85}
	fmt.Println("Marks obtained:")

	// use for range to iterate through the key-value pair
	for subject, marks := range subjectMarks {
		fmt.Println(subject, ":", marks)
	}
}
