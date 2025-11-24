package main

import (
	"fmt"
)

func main() {
	// buffered channel with buffer size of one
	// can send to it 1 message even wihtout having a receiver
	// it won't block if we send only 1 msg
	// if we try to write more to it, there should be a receiver
	// to consume the message, otherwise deadlock will occur
	ch := make(chan string, 1)
	ch <- "hello"
	msg := <-ch
	fmt.Println(msg)
}
