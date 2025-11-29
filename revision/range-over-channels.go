package main

import "fmt"

func main() {
	queue := make(chan int, 2)
	queue <- 1
	queue <- 2
	
	// this is important to stop ranging over the channel
	// otherwise will cause deadlock
	close(queue)

	for num := range queue {
		fmt.Println(num)
	}
}
