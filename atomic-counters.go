package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// using atomic.Uint64 to increment this atomically
	// across multiple go routines to avoid rance conditions
	// using a normal int will cause indeterministic behaviour
	// which can be observed from different values of num
	// in different execution of the program

	var ops atomic.Uint64
	var wg sync.WaitGroup
	var num int = 0
	
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				// ops.Add() method is used to atomically increment the unisigned int
				ops.Add(1)
				num++
			}
			fmt.Println("Worker done")
		}()
	}

	wg.Wait()

	// ops.Load() is even safe to read when other goroutines are using it
	fmt.Println("ops:", ops.Load(), "num:", num)
}
