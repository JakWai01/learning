package main

import (
	"flag"
	"log"

	"github.com/JakWai01/golang-tutorial/pkg/server"
)

func main() {
	listenAddress := flag.String("laddr", "0.0.0.0:8080", "Listen address.")

	flag.Parse()

	log.Println(*listenAddress)

	tcpServer := server.NewTCPServer(*listenAddress)

	// Open instances
	if err := tcpServer.Open(); err != nil {
		log.Fatal("could not open tcpServer", err)
	}
}
