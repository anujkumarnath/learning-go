package main

import (
	"fmt"
	"slices"
)

func main() {
	var s[] string
	fmt.Println("uinit:", s, s == nil, len(s) == 0)

	// cannot write to this empty slice (size 0)
	// this will fail 
	// s[0] = "a"

	// Slice with non-zero length
	s = make([]string, 3) // zero valued slice with capacity 3
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "bc"
	s[2] = "def"

	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// by default the capacity of the slice is equal to its size
	// but we can increase it by using aapend call
	s = append(s, "ghi")
	s = append(s, "jkl")

	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	// copying a slice
	// make a empty slice of the same size
	c := make([]string, len(s)) 

	// copy(dest, src)
	copy(c, s)

	fmt.Println("copy:", c, "len:", len(c), "cap:", cap(c))

	src := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// the slice(:) operator
	// src[start, end BEFORE]
	l := src[3:7]
	fmt.Println(l)

	// all before 7th index
	m := src[:7]
	fmt.Println(m)

	// all starting index 3
	n := src[3:]
	fmt.Println(n)

	// copy all
	o := src[:]
	fmt.Println(o)

	// slice package utility functions
	if slices.Equal(c, s) {
		fmt.Println("c == s")
	}

	// multidimentional array
	// inner arrays can have dynamic size unlike arrays
	twoD := make([][]int, 3)
	for i := range 3 {
		innerLenght := i + 1
		twoD[i] = make([]int, innerLenght)
		for j := range innerLenght {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)

}
