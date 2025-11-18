package main

import (
	"fmt"
	"sync"
)

// to update the struct from multiple goroutines safely
// we need to use mutex(to synchronize access)
// NOTE: mutexes shouldn't be copied, so if we want to pass the struct to
// another function, we should pass the pointer and not the value 
type Container struct {
	mu       sync.Mutex
	counters map[string]int
}

// Lock the mutex before accessing the counters
// Unlock it at the end of function using defer
func (c *Container) inc(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counters[name]++
}

func main() {
	c := Container{
		// zero value of mutex is usable as-is,
		// no intialization needed
		counters: map[string]int{"a": 0, "b": 0},
	}

	var wg sync.WaitGroup
	
	doIncrement := func(name string, n int) {
		defer wg.Done()
		for range n {
			c.inc(name)
		}
	}

	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)

	/* This is also completely valid:
	wg.Add(1)
	go doIncrement("a", 1000)

	wg.Add(1)
	go doIncrement("a", 1000)

	wg.Add(1)
	go doIncrement("b", 1000)
	*/

	wg.Wait()
	fmt.Println(c.counters)
}
