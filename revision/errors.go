package main

import (
	"fmt"
	"errors"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}

	return arg + 3, nil
}

var ErrorOutOfTea = errors.New("no more tea available")
var ErrorPower = errors.New("can't boil water")

func makeTea(arg int) error {
	if arg == 2 {
		return ErrorOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrorPower)
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
			if errors.Is(err, ErrorOutOfTea) {
				fmt.Println("We shoudl buy new tea")
			} else if errors.Is(err, ErrorPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unkown error: %s\n", err)
			}
			continue
		}

		fmt.Println("Tea is ready!")
	}
		
}
