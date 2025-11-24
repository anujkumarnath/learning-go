package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func(done chan<- bool) {
		fmt.Println("Worker started")
		time.Sleep(5 * time.Second)
		done <- true
	}(done)

	fmt.Println("Waiting for worker to complete...")
	<-done
	fmt.Println("Worker done!")
}
