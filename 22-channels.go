package main

import "fmt"

func main() {

	// create new channel 
	messages := make(chan string)

	// send value to the channel
	go func() { messages <- "ping" }()

	// receive value from channel
	msg := <- messages

	fmt.Println(msg)
}