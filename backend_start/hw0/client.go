package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Address = "localhost:8080"
	ExitError = 1
	Response = "OK\n"
)

func main() {
	// Connect to the TCP server on localhost:8080
	conn, err := net.Dial("tcp", Address)
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(ExitError)
	}
	defer conn.Close()
	fmt.Println("Connected to TCP server on", Address)

	for {
		// Read the server's response
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading response:", err.Error())
			return
		}
		if string(buffer[:n]) == Response {
			fmt.Printf("Server response: %s", string(buffer[:n]))
			return
		}
	}
}
