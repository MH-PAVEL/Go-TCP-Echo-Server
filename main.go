// Package main implements a simple TCP echo server in Go.
// The server listens for incoming connections on a specified port,
// reads data from clients, and echoes the same data back to them.
package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Parse command-line arguments to get the port number
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Please provide a port number.")
		return
	}
	port := ":" + arguments[1]

	// Start the TCP server by listening on the specified port
	// net.Listen creates a server socket and binds it to the port
	l, err := net.Listen("tcp4", port)
	if err != nil {
		panic(err)
	}

	// Print server status
	fmt.Printf("Server is listening at port %s\n", port)

	// Ensure the listener is closed when the function exits
	// This prevents resource leaks and frees up the port
	defer l.Close()

	// Main server loop: continuously accept new connections
	// Each connection is handled in a separate goroutine for concurrency
	for {
		c, err := l.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(c)
	}
}

// handleConnection manages a single client connection.
// It reads data from the client and echoes it back until the connection is closed.
// This function runs in a separate goroutine for each client connection.
func handleConnection(c net.Conn) {
	// Log the client's remote address 
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	
	// Initialize buffers for reading and storing data
	packet := make([]byte, 4096)
	tmp := make([]byte, 4096)
	
	// Ensure the connection is closed when this function exits
	// This sends a TCP FIN packet to gracefully close the connection
	defer c.Close()
	
	// Read data from the client in a loop until connection is closed
	for {
		// Read data into the temporary buffer
		_, err := c.Read(tmp)
		if err != nil {
			// Check if the error is due to end-of-file (client closed connection)
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		// Append the read data to the main packet buffer
		packet = append(packet, tmp...)
	}
	
	// Echo all received data back to the client
	c.Write(packet)
}