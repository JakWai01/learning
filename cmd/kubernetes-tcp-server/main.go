package main

import (
	"flag"
	"log"
)

func main() {
	laddr := flag.String("laddr", "0.0.0.0:1927", "Listen address.")

	flag.Parse()

	log.Println(*laddr)
}
