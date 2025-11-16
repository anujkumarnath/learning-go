package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const s = "สวัสดี"
	const l = "abc"
	fmt.Println("s:", s, "len:", len(s))
	fmt.Println("l:", l, "len:", len(l))
	
	fmt.Println("Rune count of s:", utf8.RuneCountInString(s))
	fmt.Println("Rune count of l:", utf8.RuneCountInString(l))

	for idx, runeValue := range s {
		fmt.Printf("%s -> idx: %d, runeValue: %x\n", s, idx, runeValue)
	}

	for idx, runeValue := range l {
		fmt.Printf("%s -> idx: %d, runeValue: %x\n", l, idx, runeValue)
	}

	fmt.Println("\nUsing DecodeRuneInString:")
	for i := 0; i < len(s); {
		runeValue, width := utf8.DecodeRuneInString(s[i:])	
		fmt.Printf("%s -> idx: %d, runeValue: %#U\n", s, i, runeValue)
		i += width
	
		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
