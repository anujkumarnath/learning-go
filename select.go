package main

import (
	"fmt"
	"time"
)

// select allows to wait on multiple channles

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(){
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()
	
	// the loop will run twice - 0, 1
	// there will be message in either c1 or c2
	// after that if we try to read, there will be a deadlock
	for range 2 {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
