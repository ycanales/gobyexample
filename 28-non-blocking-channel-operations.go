package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// no message available on 'messages',
	// then 'default' case is selected
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 'messages' has no buffer nor receiver
	// then 'default' case is selected.
	msg := "hi"
	select {
	case messages <-msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}