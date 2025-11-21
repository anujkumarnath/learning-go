package main

import (
	"fmt"
)

// recover provides a way to recover from a panic

func mayPanic() {
	panic("a problem")
}

func main() {
	// recover must be called inside a defer function
	defer func() {
		// the return value of recover() is the
		// error raised in the panic
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	// this won't run because of the panic
	// execution of main stops at the point of panic
	// and resumes in the deferred closure
	fmt.Println("After mayPanic()")
}
