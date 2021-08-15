package main

import (
	"flag"
	"log"

	"github.com/JakWai01/golang-tutorial/pkg/client"
)

func main() {
	listenAddress := flag.String("laddr", "192.168.178.38:30036", "Listen address.")

	flag.Parse()

	log.Println(*listenAddress)

	tcpClient := client.NewTCPClient(*listenAddress)

	// Open instances
	if err := tcpClient.Open(); err != nil {
		log.Fatal("could not open tcpClient", err)
	}
}
