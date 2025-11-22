package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Go signal notification works by sending os.Signal values
	// on a specified channel
	// NOTE: the channel should be buffered
	sigs := make(chan os.Signal, 1)
	
	// signal.Notify registers the channel to receive notification of
	// specified signals
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// We could receive from sigs here in the main function,
	// but let’s see how this could also be done in a separate goroutine,
	// to demonstrate a more realistic scenario of graceful shutdown.
	done := make(chan bool, 1)

	go func() {
		// this goroutine executes a blocking recieive for signals
		// to proceed further and tell another gouroutine(main)
		// that it can finish
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("Awaiting signals")
	<-done
	fmt.Println("exiting")
}
