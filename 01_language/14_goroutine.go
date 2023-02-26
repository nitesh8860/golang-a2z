package main

import "fmt"

func main() {

	// create channel
	ch := make(chan string)

	// function call with goroutine
	go receiveData(ch)
	go receiveData(ch)

	fmt.Println("No data. Receive Operation Blocked")

	// send data to the channel
	ch <- "Data 1 Received. Receive Operation Successful"
	ch <- "Data 2 Received. Receive Operation Successful"

}

func receiveData(ch chan string) {

	// receive data from the channel
	fmt.Println(<-ch)

}
