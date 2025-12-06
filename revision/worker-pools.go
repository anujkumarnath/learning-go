package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int , results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d started job %d\n", id, j)
		time.Sleep(1 * time.Second)
		fmt.Printf("Worker %d finished job %d\n", id, j)
		results <- j * 2
	}
}

func main() {
	numJobs := 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	numWorkers := 2

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	for j := range numJobs {
		jobs <- j+1
		fmt.Println("Submitted job:", j+1)
	}

	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	
}
