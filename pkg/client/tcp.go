/*
tcpClient
*/

package client

import (
	"fmt"
	"net"
)

// TCPClient makes TCP Request
type TCPClient struct {
	laddr string
}

// NewTCPClient initializes laddr
func NewTCPClient(laddr string) string {
	return laddr
}

// Open sends request to TCP server with laddr as argument
func (s *TCPClient) Open() error {

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
