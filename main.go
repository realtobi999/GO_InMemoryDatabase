package main

import (
	"fmt"
	"log"
	"net"
)

const PORT = 6379

func main() {
	log.Printf("[EVENT]\tListening on port: ':%v'", PORT)

	// start a server
	server, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%v", PORT))
	if err != nil {
		log.Fatal(err)
	}

	// listen for connections
	conn, err := server.Accept()
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			log.Println("[ERROR] Error reading using RESP: "+err.Error())	
			return 
		}

		fmt.Println(value)

		conn.Write([]byte("+OK\r\n"))
	}

}