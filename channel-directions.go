package main

import "fmt"

// write only channel, trying to read from ti would
// cause compile time error
func ping(pings chan<- string, msg string){
	pings <- msg
}

// pings channel is a read only channel, trying to write
// to it would cause a compile time error
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg 
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
