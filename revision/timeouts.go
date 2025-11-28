package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "hi"
	}()

	select {
	case msg := <-ch1:
		fmt.Println("message from goroutine 1:", msg)
	case t := <- time.After(2 * time.Second):
		fmt.Println("timeout before goroutine 1 could finish:", t)
	}

	ch2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "hi"
	}()

	select {
	case msg := <-ch2:
		fmt.Println("message from goroutine 2:", msg)
	case t := <- time.After(2 * time.Second):
		fmt.Println("timeout before goroutine 2 could finish:", t)
	}

}
