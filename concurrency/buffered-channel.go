package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}

	/* buffered channel with capacity 5 */
	channel := make(chan int, 5) 
	
	for _, num := range nums {
		channel <- num
	}

	/* this marks the end of the buffer
	 * and the following range call knows when to stop
	 * Try commenting this our
	 */
	close(channel)

	for result := range channel {
		fmt.Println(result)
	}

}
