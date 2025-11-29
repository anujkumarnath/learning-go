package main

import "fmt"

func main() {
	messages := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			if msg, more := <-messages; more {
				fmt.Println("received message", msg)	
			} else {
				fmt.Println("no more messages")
				done <- true
				return
			}
		}
	}()

	for i := range 10 {
		messages <- i+1
		fmt.Println("sent message", i+1)
	}

	close(messages)
	fmt.Println("sent all messages")

	<-done

	_, ok := <-messages
	fmt.Println("received more messages:", ok)
}
