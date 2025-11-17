package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// increments the counter by 1
		wg.Add(1)
		go func() {
			// decrements the counter by 1
			defer wg.Done()
			worker(i)
		}()
	}

	// blocks untill all the gorutines started by wg are done
	wg.Wait()
}
