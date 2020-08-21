package main

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


							x := 3
							for x < 5 {
								fmt.Println(x)
								x++
							}

							for x := 0; x <= 5; x++ {
								fmt.Println(x)
							}

							// break continue


							// switch statement, kein Fall through
							ans := 2
							switch ans {
							case 1, -1:
								fmt.Println("one")

							case 2, -2:
								fmt.Println("two")

							default:
								fmt.Println("not a case")
							}

							switch {
							case ans > 0:
								fmt.Println("greater than 0")
							case ans < 0:
								fmt.Println("less than 0")
							default:
								fmt.Println("zero")
							}



						// Arrays

						var arr [5]string
						arr[0] = "Jakob"
						arr[3] = "Jakobus"
						fmt.Println(arr)

					arr := [3]int{4, 5, 6}
					fmt.Println(arr)
					fmt.Println(len(arr))
					sum := 0

					for i := 0; i < len(arr); i++ {
						sum += arr[i]
					}

					fmt.Println(sum)

					arr2D := [2][2]int{{1, 2}, {3, 4}}
					fmt.Println(arr2D)



				var x [5]int = [5]int{1, 2, 3, 4, 5}
				var s []int = x[1:3]
				// Wie viele ELemente könnten potentiell noch rechts im slice sein
				fmt.Println(cap(s))
				fmt.Println(s[1:])



			var a []int = []int{5, 6, 7, 8, 9}
			fmt.Println(cap(a[:3]))
			b := append(a, 10)
			c := make([]int, 5)
			fmt.Println(b)
			fmt.Println(c)



		var a []int = []int{1, 3, 4, 56, 7, 12, 4, 6}

		for i := 0; i < len(a); i++ {
			fmt.Println(a[i])
		}

		for i, element := range a {
			fmt.Printf("%d: %d\n", i, element)
		}

		for _, element := range a {
			fmt.Printf("%d\n", element)
		}

	*/
}
