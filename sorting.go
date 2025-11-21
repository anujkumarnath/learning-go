package main

import (
	"fmt"
	"slices"
)

func main() {
	strs := []string{"john", "John", "Doe", "doe", "abc", "124" }
	slices.Sort(strs)
	fmt.Println(strs)

	ints := []int{4, 5, 1}
	slices.Sort(ints)
	fmt.Println(ints)

	fmt.Println("Is sorted:", slices.IsSorted(ints))
}
