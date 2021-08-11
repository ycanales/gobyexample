package main

import (
	"fmt"
	"time"
)

func main() {

	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// receive a value every 200ms
	limiter := time.Tick(200 * time.Millisecond)

	// wait for "limiter" channel before serving a request.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// allow bursts of 3 requests
	burstyLimiter := make(chan time.Time, 3)

	// fill up the channel to represent allowed bursting.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// try to add a value to "burstyLimiter"
	// every 200ms up to its limit of 3.
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// simulate 5 more incoming requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
        <-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}