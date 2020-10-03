/*
 tcpServer
*/

package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

// TCPServer creates TCP server
type TCPServer struct {
	laddr string
}

// NewTCPServer initializes laddr
func NewTCPServer(laddr string) string {
	return laddr
}

// Open waits for requests
func (s *TCPServer) Open() error {

	tcpAddr, err := net.ResolveTCPAddr("tcp", s.laddr)
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
