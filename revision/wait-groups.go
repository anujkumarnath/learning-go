package main

import (
	"fmt"
	"sync"
)

func worker(id int) {
	fmt.Println("Worker:", id, "running")
}

func main() {
	numWorkers := 10
	var wg sync.WaitGroup

	for i := range numWorkers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker(i+1)			
		}()
	}

	wg.Wait()
}
