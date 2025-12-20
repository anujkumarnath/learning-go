package main

import (
	"fmt"
	"net"
	"bufio"
	"os"
)

func main() {
	serverAddr := "localhost:42069"

	udpAddr, err := net.ResolveUDPAddr("udp", serverAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error resolving UDP address: %v\n", err)
		os.Exit(1)
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error dialing UDP: %v\n", err)
	}
	defer conn.Close()

	fmt.Printf("Sending to %s. Type your message and press Enter to send. Press Ctrl+C to exit.\n", serverAddr)

	//var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ");
		var delim byte = '\n'
		message, err := reader.ReadString(delim)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
			os.Exit(1)
		}

		_, err = conn.Write([]byte(message))
		if err != nil {
			panic(err)
		}

		fmt.Printf("Message sent: %s", message)

	}

}
