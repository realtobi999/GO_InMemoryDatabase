package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const PORT = 6379

func main() {
	fmt.Printf("[EVENT]\tListening on port: ':%v'", PORT)

	// Start  a server
	server, err := net.Listen("tcp", fmt.Sprintf(":%v", PORT))
	if err != nil {
		log.Fatal(err)
	}

	// Listen for connections
	conn, err := server.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		buf := make([]byte, 1024)

		// read message from client
		_, err := conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}

			fmt.Println("[ERROR]\tError reading from client: "+err.Error())
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}

}