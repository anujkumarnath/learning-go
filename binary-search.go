package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2

		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}

	return -1
}

func main() {
	nums := []int{1, 2, 5, 13, 15, 23, 24}

	fmt.Println("Numbers: ", nums)

	var target int

	fmt.Print("Enter the number to search: ")
	fmt.Scan(&target)

	index := search(nums, target)

	if index == -1 {
		fmt.Println("The number isn't found!")
	} else {
		fmt.Println("The number is found at index:", index)
	}

}
