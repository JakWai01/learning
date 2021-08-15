package main

import "fmt"

func main() {

	// // unbuffered channel
	// words := make(chan string)

	// go func() {
	// 	words <- "Jakob"
	// 	// This is not displayed because the channel won't wait for a second item to be added unless its buffered
	// 	words <- "Waibel"
	// }()

	// msg := <-words
	// fmt.Println(msg)

	// // unbuffered channel
	// words := make(chan string, 2)

	// // goroutine sending information to channel
	// go func() {
	// 	words <- "Jakob"
	// }()

	// // goroutine receiving information
	// go func() { fmt.Println(<-words) }()

	// // this is necessary because otherwhise the code will have ended before the second goroutine received anything so it can't be executed properly
	// time.Sleep(2 * time.Second)

	// use waitgroup instead of time.Sleep
	// with waitgroups we can successfully pass information between goroutines
	// var wg sync.WaitGroup
	// words := make(chan string)

	// wg.Add(1)

	// go func() {
	// 	words <- "Jakob"

	// }()

	// // We just need to wait for this one to finish
	// go func() {
	// 	fmt.Println(<-words)
	// 	wg.Done()
	// }()

	// wg.Wait()

	buff := make(chan string, 2)

	buff <- "Jakob"
	buff <- "Waibel"

	for i := 0; i < 2; i++ {

		fmt.Println(<-buff)

	}

}
