package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		fmt.Println("Should run next")
	}()

	fmt.Println("Should run first")
	time.Sleep(1 * time.Second)
}
