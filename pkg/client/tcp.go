/*
tcpClient
*/

package client

import (
	"fmt"
	"net"
	"os"
)

// TCPServer consists of laddr
type TCPServer struct {
	laddr string
}

// NewTCPServer initializes laddr
func NewTCPServer(laddr string) string {
	return laddr
}

// Open initializes TCP server with laddr
func (s *TCPServer) Open() error {

	tcpAddr, err := net.ResolveTCPAddr("tcp", s.laddr)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err2 := conn.Write([]byte("Hello World"))
	checkError(err2)

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(buf[0:n])

	os.Exit(0)

	return nil

}

func checkError(err error) error {
	return err
}

// func main() {
// 	laddr := flag.String("laddr", "0.0.0.0:8080", "TCP address")

// 	service := *laddr

// 	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
// 	checkError(err)

// 	conn, err := net.DialTCP("tcp", nil, tcpAddr)
// 	checkError(err)

// 	_, err2 := conn.Write([]byte("anything"))
// 	checkError(err2)

// 	var buf [512]byte

// 	n, err := conn.Read(buf[0:])
// 	checkError(err)

// 	fmt.Println(buf[0:n])

// 	os.Exit(0)

// }

// func checkError(err error) {

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
