package main

import (
	"fmt"
	"os"
	"io"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("./messages.txt")
	check(err)

	buffer := make([]byte, 8)

	for {
		readBytes, err := file.Read(buffer)
		if err == io.EOF {
			break
		} else {
			check(err)
		}
		fmt.Printf("read: %s\n", buffer[:readBytes])
	}
}
