package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// in other languages, strings are made up of characters
	// in go, the concept of character is called a rune
	const s1 = "abc"
	const s2 = "สวัสดี"

	// strings are equivalent to a slice of bytes([]byte)
	// so len funcion gives the lenght of raw bytes stored within
	fmt.Printf("length of %s: %d\n", s1, len(s1))
	fmt.Printf("length of %s: %d\n", s2, len(s2))

	// indexing a string produces raw byte values at each index
	// we are printing the hex value of each byte
	for i := 0; i < len(s1); i++ {
		fmt.Printf("s1[%d] = %x\n", i, s1[i])
	}
	fmt.Println()

	for i := 0; i < len(s2); i++ {
		fmt.Printf("s2[%d] = %x\n", i, s2[i])
	}
	fmt.Println()

	// strings are treated by range loop specially decoding each rune
	// and offset of the rune in the sequence of bytes
	for idx, runevalue := range s1 {
		fmt.Printf("%#U of %6s starts at %d\n", runevalue, s1, idx)
	}

	for idx, runevalue := range s2 {
		fmt.Printf("%#U of %6s starts at %d\n", runevalue, s2, idx)
	}
	fmt.Println()

	// some Thai charactes represented by UTF-8 code points that can take multiple bytes
	// so the result of the count may be surprising
	fmt.Println("Rune count of s1(latin):", utf8.RuneCountInString(s1))
	fmt.Println("Rune count of s2(thai):", utf8.RuneCountInString(s2))

	// decoding manually
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s2); i += w {
		runeValue, width := utf8.DecodeRuneInString(s2[i:])
		fmt.Printf("%#U of %s starts at %d\n", runeValue, s2, i);
		w = width

		examineRune(runeValue)
	} 
}

func examineRune(r rune) {
	// values enclosed in single quotes are rune literals
	// rune values can be compared directly to a rune literal by == operator
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
