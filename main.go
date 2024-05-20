package main

import (
	"fmt"
	"log"
	"net"
	"strings"
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
			log.Println("[ERROR]\tError reading using RESP: "+err.Error())	
			return 
		}

		if value.typ != "array" {
			fmt.Println("[ERROR]\tInvalid request, expected array")
		continue
		}

		if len(value.array) == 0 {
			fmt.Println("[ERROR]\tInvalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		writer := NewWriter(conn)

		handler, ok := Handlers[command]
		if !ok {
			fmt.Println("[ERROR]\tInvalid command: ", command)
			writer.Write(Value{typ: "string", str: ""})
			continue
		}

		writer.Write(handler(args))
	}

}