/*
tcpClient
*/

package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	laddr := flag.String("laddr", "0.0.0.0:8080", "TCP address")

	service := *laddr

	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	checkError(err)

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	_, err2 := conn.Write([]byte("anything"))
	checkError(err2)

	var buf [512]byte

	n, err := conn.Read(buf[0:])
	checkError(err)

	fmt.Println(buf[0:n])

	os.Exit(0)

}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
