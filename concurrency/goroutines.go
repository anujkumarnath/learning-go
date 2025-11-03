package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Before invoking goroutine")

	go func() {
		fmt.Println("Goroutine running...");
	}()

	/* otherwise the program will exit before the go routine finishes executing */
	fmt.Println("Sleeping for second...");
	time.Sleep(1 * time.Second)

	fmt.Println("Before invoking goroutine")
}


