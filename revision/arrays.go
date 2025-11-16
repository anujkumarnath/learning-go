package main

import (
	"fmt"
)

func main() {
	var arr [5]int
	fmt.Println(arr)

	arr[4] = 20

	fmt.Println("len", len(arr))

	a := [4]int{0, 1, 2, 3}
	fmt.Println(a)

	var b [2]int = [2]int{0, 2}
	fmt.Println(b)
	c := [...]int{1,2,3,4,4,5,5}
	fmt.Println(c)

	d := [...]int{0, 1, 2, 3: 3, 7: 8, 9}
	fmt.Println(d)

	e := [9]int{0, 1, 2, 3: 3, 7: 8, 9}
	fmt.Println(e)

	var twoD [2][3]int
	for i := range 2 {
		for j := range 3 {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{0, 1, 2},
		{3, 4, 5},
	}
	fmt.Println("2d:", twoD)
}
