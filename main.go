package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
	}
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Printf("Error starting server on port %s: %v\n", port, err)
		os.Exit(1)
	}
	fmt.Printf("Listening on port %s\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("connection from ", conn.RemoteAddr())
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	fmt.Printf("Handling connection from %s\n", conn.RemoteAddr())

	// Create a buffer for reading data
	buffer := make([]byte, 1024)

	for {
		// Read from the connection
		n, err := conn.Read(buffer)
		if err == io.EOF {
			fmt.Printf("Connection closed by client %s\n", conn.RemoteAddr())
			break
		}
		if err != nil {
			fmt.Printf("Error reading from client %s: %v\n", conn.RemoteAddr(), err)
			break
		}

		// Echo data back to the client
		_, writeErr := conn.Write(buffer[:n])
		if writeErr != nil {
			fmt.Printf("Error writing to client %s: %v\n", conn.RemoteAddr(), writeErr)
			break
		}
	}

	fmt.Printf("Successfully echoed data to %s\n", conn.RemoteAddr())
}
