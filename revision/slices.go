package main

import (
	"fmt"
	"slices"
)

func main() {
	var s []string
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = make([]string, 3)
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "d")
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, "e")
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	s = append(s, []string{"f", "g"}...)
	fmt.Println("s:", s, "len:", len(s), "cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[:3]
	fmt.Println("l:", l)

	m := s[3:6]
	fmt.Println("m:", m)

	n := s[:10]
	fmt.Println("n:", n)

	fmt.Println(slices.Equal(c, s))

	twoD := make([][]int, 3)
	for i := range 3 {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen) 
		for j := range innerLen {
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)
}
