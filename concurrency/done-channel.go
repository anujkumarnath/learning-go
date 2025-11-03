package main

import (
	"fmt"
	"time"
)


func doWork() chan bool {
	done := make(chan bool)
	go func(done <- chan bool) {
		for {
			select {
			case <- done:
				return
			default:
				fmt.Println("DOING WORK")
			}
		}
	}(done)
	return done;
}

func main() {
	/* Assume our application's intended run time is second,
	 * So, we may not want to keep the go routine running till the end of the application
	 * To stop the goroutine, we need to stop it
	 * Here the "done" channel helps
	 */
	done := doWork();

	time.Sleep(1 * time.Second);
	close(done)
	fmt.Println("Stopping doWork")

	time.Sleep(5 * time.Second)
	fmt.Println("Exiting application")
}
