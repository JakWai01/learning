package main

import "fmt"

func main() {

	// if a variable isn't used -> error
	// only numbers, letters and underscores in the variable names
	var name string = "Hello World!"
	name = "Jakob"
	name = "Paula"
	var number int = 12
	// there are signed integers and unsigned integers
	var unsignednumber uint8 = 255
	var signednumber int8 = -128

	var floaty float32 = 1.2
	// uint (32 or 64 bits)
	// int (same size as uint)
	// float32 or float64
	// complex64 (complex numbers)
	// complex128 (complex numbers)
	// Strings
	// boleans
	fmt.Println(floaty, name)
	fmt.Println(floaty)
	fmt.Println(unsignednumber)
	fmt.Println(signednumber)
	fmt.Println(number)
	fmt.Println(name)

}
