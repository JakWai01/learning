// /*
// UDPDaytimeClient
// */

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// )

// func main() {

// 	if len(os.Args) != 2 {
// 		fmt.Fprintf("Usage of: %s host:port", os.Args[0])
// 		os.Exit(1)
// 	}

// 	service := os.Args[1]

// 	udpAddr, err := net.ResolveUDPAddr("up4", service)
// 	checkError(err)

// 	conn, err := net.DialUDP("udp", nil, udpAddr)
// 	checkError(err)

// 	_, err := conn.Write([]byte("anything"))
// 	checkError(err)

// 	var buf [512]byte
// 	n, err := conn.Read(buf[0:])
// 	checkError(err)

// 	fmt.Println(string(buf[0:n]))

// 	os.Exit(0)

// }

// func checkError(err error) {

// 	log.Fatal(err)
// 	os.Exit(1)
// }
