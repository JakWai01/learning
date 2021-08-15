/*
SimpleEchoServer
*/
//
// import (
// 	"fmt"
// 	"net"
// 	"os"
// 	"log"
// )

// func main() {

// 	service := "1201"
// 	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
// 	checkError(err)

// 	listener, err := net.ListenTCP("tcp", tcpAddr)
// 	checkError(err)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			continue
// 		}

// 		go handleClient(conn)

// 	}
// }

// func handleClient() {
// 	var buf [512]byte

// 	defer conn.Close()

// 	for {
// 		n, err := conn.Read(buf[0:])
// 		if err != nil {
// 			continue
// 		}
// 		fmt.Println(string(buf[0:]))
// 		_, err2 := conn.Write(buf(0:n))
// 		if err2 != nil {
// 			return
// 		}

// 	}
// }

// func CheckError(err error) {

// 	log.Fatal(err)
// }