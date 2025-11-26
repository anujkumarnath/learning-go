package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		time.Sleep(time.Second)
		ch1 <- "msg1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "msg2"
	}()

	for range 2 {
		select {
		case msg := <-ch1:
			fmt.Println("msg received from chan 1:", msg)
		case msg := <-ch2:
			fmt.Println("msg received from chan 2:", msg)
		}
	}

	fmt.Println("Exiting main")
}
