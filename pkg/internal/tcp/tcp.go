package main

import (
	"fmt"
	"net"
	"sync"
)

func handleConnection(conn net.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		chunk := make([]byte, 1024)
		readBytes, err := conn.Read(chunk)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Printf("Received data: %v", string(chunk[:readBytes]))
	}
}

func InitTCP(wg *sync.WaitGroup) {
	defer wg.Done()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	var wg_tcp sync.WaitGroup
	for {
		fmt.Println("Listening on 8080")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			return
		}
		fmt.Println("Accepted connection")
		wg_tcp.Add(1)
		go handleConnection(conn, &wg_tcp)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go InitTCP(&wg)
	wg.Wait()
}
