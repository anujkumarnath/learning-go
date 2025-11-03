package main

import (
	"fmt"
	// "time"
)

func goroutine(channel chan string) {
	fmt.Println("I am a goroutine that communicates via a channel")
	channel <- "Done running go routine."
}

func main() {
	/* unbuffered channel - the read/write is synchronous
	 * i.e., the writer will be paused until the reader reads the message
	 */
	channel := make(chan string)
	go goroutine(channel)
	msg := <- channel


	/*
		// Uncomment to see the wrong "Sleep" way to wait for goroutines
		fmt.Println("Before running goroutine")

		go func() {
			fmt.Println("Goroutine running...");
		}()

		// otherwise the program will exit before goroutine finishes executing
		fmt.Println("Sleeping for second for goroutine ...");
		time.Sleep(1 * time.Second)
		fmt.Println("After running Goroutine 1")
	*/

	fmt.Println(msg)
}

