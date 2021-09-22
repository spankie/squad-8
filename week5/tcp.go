package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	// 65535
	// 127.0.0.1
	// 192.168.0.1
	// 255.255.255.255
	service := "localhost:1200" // localhost:1200, 127.0.0.1, 0.0.0.0
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	fmt.Printf("listening on %v:%v\n", tcpAddr.IP, tcpAddr.Port)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept() // waits for someone to connect
		if err != nil {
			continue
		}
		fmt.Println("someone connected")
		daytime := time.Now().String()
		conn.Write([]byte(daytime)) // don't care about return value
		conn.Close()                // we're finished with this client
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}