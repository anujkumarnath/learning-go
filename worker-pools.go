package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// creating 2 worker goroutines
	for w := 1; w <= 2; w++ {
		go worker(w, jobs, results)
	}

	// creating 5 jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	// indicating we are done producing jobs
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}

