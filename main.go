package main

import (
	"flag"
	"fmt"
	"sync"
)

func main() {

	// so kann man beim start vom Programm parameter mitgeben
	author := flag.String("author", "Jakob Waibel", "Author of the code")
	flag.Parse()
	fmt.Println(*author)

	// waitgroup
	wg := new(sync.WaitGroup)
	wg.Add(2)

	a := 1
	b := 1
	go add(a, b, wg)
	go substract(a, b, wg)

	wg.Wait()
}

func add(a, b int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return a + b
}

func substract(a, b int, wg *sync.WaitGroup) int {
	defer wg.Done()
	return a - b
}
