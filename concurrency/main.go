package main

import (
	"fmt"
)

/* main function itself is a goroutine */
func main () {
	channelA := make(chan string)
	channelB := make(chan string)

	/* it's a goroutine - starts with the 'go' keyword */
	go func() {
		channelA <- "message from go routine A"
	}()

	/* it's another goroutine */
	go func() {
		channelB <- "message from go routine B"
	}()

	/*
		slelect is a blocking thing.
		It waits for response from one channel,
		and moves ahead when any one of them is received
		If both channel has message, it choses one at random
	*/
	select {
	case msgFromChannelA := <- channelA:
		fmt.Println(msgFromChannelA)
	case msgFromChannelB := <- channelB:
		fmt.Println(msgFromChannelB)
	}

}
