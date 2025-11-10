package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n - 1)
}

func main() {
	num := 5
	fmt.Printf("Factorial of %d is %d\n", num, fact(num))

	// annonymous functions can be recursive too
	// but need to be declared explicitly with var to store the
	// function before it is defined
	var fib func(n int) int

	// fib := func(n int) int {} won't work
	fib = func (n int) int {
		if n < 2 {
			return n
		}

		return fib(n - 1) + fib(n - 2)
	}

	fmt.Println("Fib", num, " is", fib(num))
}
