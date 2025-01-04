package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"sync"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default to 9000 if not set
	}

	// Start a separate HTTP server for health checks
	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "OK")
		})
		http.ListenAndServe(":8080", nil)
	}()

	listener, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
	defer listener.Close()

	fmt.Println("TCP connection listening on port", port)

	var wg sync.WaitGroup

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("error establishing connection: %v\n", err)
			continue
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
			fmt.Printf("error reading from client: %v\n", err)
			return
		}

		_, err = conn.Write(buff[:n])
		if err != nil {
			fmt.Printf("error writing to client: %v\n", err)
			return
		}
	}
}
