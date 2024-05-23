package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

// 4. keep listening on the same connection
func handleConnection(conn net.Conn, wg *sync.WaitGroup, ctx context.Context) {
	defer wg.Done()
	fmt.Println("New connection")
	connect_ch := make(chan bool)
	go func(connect_ch chan bool) {
		for {
			chunk := make([]byte, 1024)
			readBytes, err := conn.Read(chunk)
			if err != nil {
				fmt.Println("Error reading:", err.Error())
				connect_ch <- true
				return
			}

			fmt.Printf("Received data: %v", string(chunk[:readBytes]))
		}
	}(connect_ch)

	select {
	case <-connect_ch:
		fmt.Println("Closing connection, client disconnected")
	case <-ctx.Done():
		fmt.Println("Closing connection, context cancelled")
	}

	if err := conn.Close(); err != nil {
		fmt.Println("Error closing connection:", err.Error())
	}
}

func InitTCP(wg *sync.WaitGroup, ctx context.Context) {

	defer wg.Done()

	var wg_tcp sync.WaitGroup
	// 1. create a listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", 8080))
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}

	// 3. keep listening for new connections
	go func() {
		// 2. start accepting connections
		for {
			conn, err := listener.Accept()
			if err != nil {
				fmt.Println("Error accepting:", err.Error())
				return
			}
			wg_tcp.Add(1)
			go handleConnection(conn, &wg_tcp, ctx)
		}
	}()

	// 4. wait for a signal to cancel
	<-ctx.Done()
	fmt.Println("Closing TCP server, context cancelled")
	if err := listener.Close(); err != nil {
		fmt.Println("Error closing listener:", err.Error())
	}
	wg_tcp.Wait()
	fmt.Println("TCP server closed")

}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go InitTCP(&wg, ctx)

	// Wait for a signal to cancel
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	fmt.Println("Received signal to cancel")
	cancel()
	wg.Wait()
}
