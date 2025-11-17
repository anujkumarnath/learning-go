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

	limiter := time.Tick(200 * time.Millisecond)

	// by blocking on a receive from the limiter channel before serving
	// each request, we limit ourselves to 1 request every 200 milliseconds
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	burstyLimiter := make(chan time.Time, 3)

	// first 3 requests will be served immediately because we are not blocking them
	for range 3 {
		burstyLimiter <- time.Now()
	}

	// remaining requests will be server at 2s interval because we are blocking the
	// sends with waiting on the channel t returned by time.Tick by ranging on it
	go func() {
		for t := range time.Tick(2 * time.Second) {
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	fmt.Println("-------------------------------------------------------------")

	// first 3 reuqests will process immediately, and rest will be processed
	// at 2s intervals
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
