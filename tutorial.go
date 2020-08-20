package main

import "fmt"

func main() {
	/*
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
			fmt.Printf("What year were you born?")
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
			fmt.Printf("You will be %d years old at the end of 2020", -input+2020)

			// Arithmetic operators * / + - % ()

			var num1 float64 = 8
			var num2 int = 4
			answer := num1 / float64(num2)

		fmt.Printf("%g", answer)



			// Conditions and Boolean expressions

			x := 5
			y := 6.5
			// <, > , <=, >= ...
			val := float64(x)+1.5 == y
			fmt.Printf("%t", val)

			// && || !



			// if, else if, else

			// else cannot be between else if (logisch)

			name := "Jakobus"
			fmt.Println("Before if")
			if name == "Jakob" {
				fmt.Println("inside if")
			} else if name == "Jakobus" {
				fmt.Println("FUCK OFF")
			} else {
				fmt.Println("WAIT, you aren't Jakob!")
			}
			fmt.Println("After if")

	*/
	x := 3
	for x < 5 {
		fmt.Println(x)
		x++
	}

	for x := 0; x <= 5; x++ {
		fmt.Println(x)
	}

	// break continue

}
