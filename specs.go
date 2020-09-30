package main

import "fmt"

func main() {

	// This should act like a newline
	/* This should act like a space */

	fmt.Println("Hello World!")

	type Student struct {
		firstname string
	}

	p1 := Student{firstname: "Jakob"}
	fmt.Println(p1)

	names := make(map[string]int)

	names["Jakob"] = 19
	fmt.Println(names["Jakob"])
}
