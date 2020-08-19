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

	// you actually don't have to declare a data type.
	// go is guessing which type it is, could be wrong. So rather declare it.
	var newnum = 12
	var newstring = "Hallo"

	// Prints datatype
	fmt.Printf("%T", newnum)
	fmt.Println()
	fmt.Printf("%T", newstring)
	fmt.Println()

	// we can also just declare without var
	qwe := 6
	fmt.Printf("%T", qwe)
	fmt.Println()

	test := true
	fmt.Println(test)

	// all data types got default values e.g uint64 = 0
	var numbers uint64
	fmt.Println(numbers)

	// format ; das format wird im string gespeichert und dann geprinted
	var out string = fmt.Sprintf("Number: %07d is cool", 45)
	fmt.Println(out)

	// escapes: \n (newline) \t (tab)
}
