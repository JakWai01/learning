/*
 tcpServer
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {

	var port = flag.String("port", "8080", "Assign to which port to listen to")
	flag.Parse()

	service := *port

	fmt.Println(service)

	tcpAddr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:"+string(service))
	checkError(err)

	ln, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := ln.Accept()
		checkError(err)

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	var buf [512]byte

	defer conn.Close()

	for {
		n, err := conn.Read(buf[0:])

		if err != nil {
			if err == io.EOF {
				break
			}
		}

		fmt.Println(buf[0:n])

		_, err2 := conn.Write(buf[0:n])
		checkError(err2)
	}
}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
