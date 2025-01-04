package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {

	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Printf("error connecting to listener: %v", err)
	}

	defer listener.Close()

	fmt.Println("TCP connection listening on port 9000")

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("error establishing connection: %v", err)
		}

		wg.Add(1)
		go handleConnection(conn, &wg)
	}

	wg.Wait()
}
