package main

import (
	"fmt"
	"errors"
)

type argsError struct {
	arg int
	message string
}

// add Error() method to argsError makes it implement the error interface
func (e argsError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, argsError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	_, err := f(42)
	fmt.Println(err)
	var ae *argsError
	// err is an interface; it may hold a wrapped error, a *argsError, or anything else.
	// errors.As walks the error chain and looks for a value assignable to *argsError.
	// When it finds one, it does `ae = (pointer to that concrete argsError value)`
	// and returns true; otherwise it returns false. It simply extracts the underlying
	// concrete value and writes it into ae.

	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("the error doesn't match argsError")
	}
}
