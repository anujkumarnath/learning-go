package main

import (
	"fmt"
	"time"
)

func main() {
	pings := make(chan int, 1)
	pongs := make(chan int, 1)

	for i := range 10 {
		fmt.Println("-> ping seq:", i)
		ping(pings, i)
		pong(pings, pongs)
		fmt.Println("<- pong seq:", <-pongs)
	}
	
}

func ping(pings chan<- int, i int) {
	pings <- i
}

func pong(pings <-chan int, pongs chan<- int) {
	time.Sleep(100 * time.Millisecond)
	seq := <-pings
	pongs <- seq 
}
