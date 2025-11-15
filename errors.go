package main

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	// putting the last argument nil signifies that there is no error
	return arg + 3, nil
}

// A "sentinal" error is a predclared variable that is used to signify
// a specific error condition
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return  ErrOutOfTea
	} else if arg == 4 {
		// errors can be wrapped with high level errors to add context
		// example - with the use of %w verb in fmt.Errorf 
		// wrapped errors create a ligical chain
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	for i := range 5 {
		if err := makeTea(i); err != nil {
			// errors.Is checks if a given error matches a specific value
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println(i, ": we should buy more tea", err)
			} else if errors.Is(err, ErrPower) {
				fmt.Println(i, ": no power", err)
			} else {
				fmt.Printf("%d: unknown error: %s\n", i, err)
			}
			continue
		}

		fmt.Println(i, ": tea is ready!")
	}
}
