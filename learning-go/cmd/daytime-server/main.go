package main

/*
DaytimeServer
*/

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {

	service := ":1200"

	tcpAddr, err := net.ListenTCP("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func checkError() {
	if err != nil {
		fmt.Println(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}
