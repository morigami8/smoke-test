package main

import (
	"fmt"
	"io"
	"net"
	"sync"
)

func main() {

	listener, err := net.Listen("tcp", "0.0.0.0:9000")

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

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	defer wg.Done()

	buff := make([]byte, 1024)

	for {
		n, err := conn.Read(buff)

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Printf("error reading from client: %v", err)
			return
		}

		_, err = conn.Write(buff[:n])

		if err != nil {
			fmt.Printf("error writing to client: %v", err)
			return
		}

	}

}
