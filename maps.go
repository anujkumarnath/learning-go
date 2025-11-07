package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string] int)

	fmt.Println("emp:", m, "len:", len(m))

	m["abc"] = 10
	m["def"] = 12

	fmt.Println("emp:", m, "len:", len(m))

	fmt.Println("existing key's value:", m["abc"])

	// finding the value of a non-existing key will return 'zero' value
	fmt.Println("non existing key's value:", m["mno"])
	
	// existence check vs actual zero check
	_, present := m["mno"]
	fmt.Println("mno present?", present)

	_, present = m["abc"]
	fmt.Println("abc present?", present)

	// deleting a key
	delete(m, "def")
	fmt.Println("map:", m, "len:", len(m))

	// deleting a no existing key
	// no panic/error
	delete(m, "hjk")
	fmt.Println("map:", m, "len:", len(m))

	// clear all entries
	clear(m)
	fmt.Println("emp:", m, "len:", len(m))

	// single line assignment
	s := map[int]string { 0: "abc", 1: "def" }
	fmt.Println(s)

	// maps equeality check
	t := map[int]string { 0: "abc", 1: "def" }
	u := map[int]string { 0: "abc", 1: "defg" }

	fmt.Println("s == t", maps.Equal(s, t), "s == u", maps.Equal(s, u))
}
