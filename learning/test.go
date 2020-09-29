//package main

// import (
// 	"bufio"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net"
// 	"sync"
// )

// func main() {

// 	// so kann man beim start vom Programm parameter mitgeben
// 	author := flag.String("author", "Jakob Waibel", "Author of the code")
// 	flag.Parse()
// 	fmt.Println(*author)

// 	// waitgroup
// 	wg := new(sync.WaitGroup)
// 	wg.Add(2)

// 	a := 1
// 	b := 1
// 	go add(a, b, wg)
// 	go substract(a, b, wg)

// 	wg.Wait()

// 	// net
// 	conn, err := net.Dial("tcp", "golang.org:80")
// 	if err != nil {
// 		log.Fatal()
// 	}
// 	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
// 	status, err := bufio.NewReader(conn).ReadString('\n')
// 	fmt.Println(status)

// }

// func add(a, b int, wg *sync.WaitGroup) int {
// 	defer wg.Done()
// 	log.Print("Additions has been done!")
// 	return a + b
// }

// func substract(a, b int, wg *sync.WaitGroup) int {
// 	defer wg.Done()
// 	log.Print("Substraction has been done!")
// 	return a - b
// }
