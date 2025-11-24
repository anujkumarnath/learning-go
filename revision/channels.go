package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)

	go func() {
		ch <- "hello"
	}()

	msg := <-ch
	fmt.Println(msg)
}
