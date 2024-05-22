package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	fmt.Println("Listening on 8080")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("Error accepting: ", err.Error())
		return
	}
	chunk := make([]byte, 1024)
	readBytes, err := conn.Read(chunk)
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
		return
	}
	fmt.Printf("Received data: %v", string(chunk[:readBytes]))
}
