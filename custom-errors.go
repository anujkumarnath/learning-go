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
	// errors.As is a more advanced version of errors.Is
	// It check that a given error (or any error in its chain)
	// matches a specific error type and converts to a value of that type, returning true
	// If there's no match it returns false
	// errors.As updates the ae pointer to point to the argsError struct
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("the error doesn't match argsError")
	}
}
