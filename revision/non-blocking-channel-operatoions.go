package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking send
	// if there is no sender, default will execute
	select {
	case msg := <-messages:
		fmt.Println("message from channel:", msg)
	default:
		fmt.Println("did'nt receive any message")
	}

	msg := "Hi"

	// non-blocking receive
	// if no receiver is there for messages channel, default will execute
	select {
	case messages <- msg:
		fmt.Println("message sent")
	default:
		fmt.Println("no message sent")
	}

	// multiway select
	select {
	case msg := <-messages:
		fmt.Println("message received", msg)
	case sig := <-signals:
		fmt.Println("signal received:", sig)
	default:
		fmt.Println("no activity")
	}
	
}
