package main

import (
	"fmt"
)

// equivalent of defining nums []int
// but cannot use them interchangably
func variadic(nums ...int) {
	fmt.Println("nums:", nums, "len:", len(nums))
}

func main() {
	variadic(1)
	variadic(1, 2)
	variadic(1, 2, 3)

	// will not work
	// variadic([]int{ 1, 2, 4})

	nums := []int{5, 6, 7, 8, 9, 10}
	variadic(nums...)
}
