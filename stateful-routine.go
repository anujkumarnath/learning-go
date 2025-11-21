package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync/atomic"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// var readOps atomic.Uint64
	// var writeOps atomic.Uint64
	var readOps uint64
	var writeOps uint64

	reads := make(chan readOp)
	writes := make(chan writeOp)

	go func() {
		// this goroutines owns the state and it updates it on request
		// instead of doing it in diffrent goroutines
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// spawning 100 goroutines to read the state concurrently
	for range 100 {
		go func() {
			for {
				read := readOp{
					key: rand.Intn(5),
					resp: make(chan int),
				}
				reads <- read
				<-read.resp
				// readOps.Add(1)
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}	
		}()
	}

	// writing from 100 goroutines concurrently
	for range 100 {
		go func() {
			for {
				write := writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool),
				}
			
				writes <- write
				<-write.resp

				// writeOps.Add(1)
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// waiting 1 second before exiting main
	// above go routines run indefinitely
	// and are leaking
	time.Sleep(time.Second)

	// readOpsFinal := readOps.Load()
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)

	// writeOpsFinal := writeOps.Load()
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
