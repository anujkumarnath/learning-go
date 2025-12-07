package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"errors"
)

const inputFilePath = "messages.txt"

func main() {
	file, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatal("could not open %s: %s", inputFilePath, err)
	}

	defer file.Close()

	for {
		buffer := make([]byte, 8)
		readBytes, err := file.Read(buffer)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			fmt.Printf("errors: %s\n", err.Error())
			break
		}

		fmt.Printf("read: %s\n", string(buffer[:readBytes]))
	}
}
