package main

import (
	"fmt"
	"net"
	"os"
)

const (
	Port = ":8080"
	ExitError = 1
	Message = "OK\n"
)

func main() {
	// Listen for incoming connections on port 8080
	listener, err := net.Listen("tcp", Port)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(ExitError)
	}
	defer listener.Close()
	fmt.Println("TCP Server listening on port", Port)

	for {
		// Accept incoming connections
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			continue
		}
		fmt.Println("Client connected:", conn.RemoteAddr().String())

		handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	// Send the message to the client
	_, err := conn.Write([]byte(Message))
	if err != nil {
		fmt.Println("Error writing:", err.Error())
		return
	}
}

