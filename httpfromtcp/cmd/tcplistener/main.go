package main

import (
	"fmt"
	"io"
	"net"
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
	listener, err := net.Listen("tcp", ":42069")

	if err != nil {
		panic(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
		}

		fmt.Println("Connection accepted from", conn.RemoteAddr())

		linesChannel := getLinesChannel(conn)
		for line := range(linesChannel) {
			fmt.Printf("%s\n", line)
		}

		fmt.Println("Connection to", conn.RemoteAddr(), "closed")
	}
}
