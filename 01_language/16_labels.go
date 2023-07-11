package main

import "fmt"

func main() {
LOOP1:
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue LOOP1
		}
		fmt.Println(i)
	}
	// output will be 0 to 9 with 5 missing as it went to continue at LOOP label

LOOP2:
	for i := 0; i < 10; i++ {
		if i == 5 {
			goto LOOP2
		}
		fmt.Println(i)
	}
	// it will reset
}
