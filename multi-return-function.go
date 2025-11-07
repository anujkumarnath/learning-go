package main

import "fmt"

func main() {
	a, b := vals()
	fmt.Println("a:", a, "b:", b)

	// can ignore one value by using the _ 'blank identifier'
	_, c := vals()
	fmt.Println("c:", c)
}

func vals() (int, int) {
	return 1, 2
}
