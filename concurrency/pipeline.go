package main

import "fmt"

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			fmt.Println("Sending", num, "to next stage")
			out <- num
		}
		close(out)
	}()

	return out
}

func square(in <- chan int) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			squared := num * num
			fmt.Println("Sending sqaured num", squared, "to next stage")
			out <- squared
		}
		close(out)
	}()

	return out
}

func main() {
	nums := []int{1, 5, 2, 6, 8, 9}

	fmt.Println("Input numbers:", nums)

	/* Stage 1: convert slice to channel */
	dataChannel := sliceToChannel(nums);

	/* Stage 2: convert to string */
	finalChannel := square(dataChannel);

	/* Stage 3: print the squared numbers */
	for n := range finalChannel {
		fmt.Println("Processing output:", n)
	}

}
