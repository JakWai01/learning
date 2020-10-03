// package main

// /*
// UDPDaytimeServer
// */

// import (
// 	"log"
// 	"net"
// 	"os"
// 	"time"
// )

// func main() {

// 	service := ":1200"

// 	udpAddr, err := net.ResolveUDPAddr("up4", service)
// 	checkError(err)

// 	conn, err := net.ListenUDP("udp", udpAddr)
// 	checkError(err)

// 	for {
// 		handleClient()
// 	}
// }

// func handleClient() {

// 	var buf [512]byte

// 	_, arr, err := conn.ReadFromUDP(buf[0:])

// 	if err != nil {
// 		return
// 	}

// 	daytime := time.Now().String()

// 	conn.WriteToUDP([]byte(daytime), addr)
// }

// func checkError(err error) {

// 	log.Fatal(err)
// 	os.Exit(1)
// }
