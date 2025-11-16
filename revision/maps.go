package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v2 := m["k2"]
	fmt.Println("v2:", v2)

	v3 := m["k3"]
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	clear(m)
	fmt.Println("map:", m, "len:", len(m))

	_, present := m["k2"]
	fmt.Println("present:", present)

	m2 := map[string]int{ "one": 1, "two": 2 }
	fmt.Println("inline declared and intialized m2:", m2)

	m3 := map[string]int{ "one": 1, "two": 2 }

	fmt.Println(maps.Equal(m2, m3))
}
