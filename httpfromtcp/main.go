package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"errors"
	"strings"
)

const inputFilePath = "messages.txt"

func getLinesChannel(file io.ReadCloser) <-chan string {

	linesChannel := make(chan string)

	go func() {
		defer file.Close()
		defer close(linesChannel)

		currentLine := ""

		for {
			buffer := make([]byte, 8)
			readBytes, err := file.Read(buffer)
			if err != nil {
				if currentLine != "" {
					linesChannel <- currentLine
					currentLine = ""
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("errors: %s\n", err.Error())
				break
			}

			text := string(buffer[:readBytes])
			parts := strings.Split(text, "\n")

			for i := 0; i < len(parts) - 1; i++ {
				linesChannel <- currentLine+parts[i]
				currentLine = ""
			}

			currentLine += parts[len(parts) - 1]
		}
	}()

	return linesChannel
}

func main() {
	file, err := os.Open(inputFilePath)

	if err != nil {
		log.Fatal("could not open %s: %s", inputFilePath, err)
	}

	linesChannel := getLinesChannel(file)
	for line := range(linesChannel) {
		fmt.Printf("read: %s\n", line)
	}
}
