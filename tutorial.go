package main

/*
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



		var mp map[string]int = map[string]int{
			"apple":  5,
			"pear":   6,
			"orange": 9,
		}
		fmt.Println(mp["apple"])
		// so kann man einfach elemente hinzufügen
		mp["tim"] = 900

		val, ok := mp["apple"]
		fmt.Println(val, ok)
		fmt.Println(mp)

		// -> nur arrays benutzen wenn die reihenfolge wichtig ist


	// Functions (func main) function sind methods

	ans1, ans2 := add(1, 2)
	fmt.Println(ans1, ans2)

}

// man kann auch mehrere Sachen returnen, dann zum Beispiel (int, int)
func test(a int, b int) (int, int) {

	return a + b, a - b

}

func add(x, y int) (z1 int, z2 int) {
	// defer happens at return (also beim return passiert das defer)
	defer fmt.Println("hello")
	z1 = x + y
	z2 = x - y
	fmt.Println("before return")
	return
}



func test2(myFunc func(int) int) {
	fmt.Println(myFunc(7))
}

func returnFunc(x string) func() {
	return func() {
		fmt.Println(x)
	}
}
func main() {
	// extrem effiziente anwendung, das ist im Endeffekt die Methode und der Methodcall in einem
	//test := func(x int) int {
	//	return -x
	//}
	// functions can be passed around
	//test3 := func(x int) int {
	//	return x * 7
	//}
	//test2(test3)
	returnFunc("hello")()
	x := returnFunc("goodbye")
	x()
}
*/

//func main() {
/*var x int = 5
y := x
y = 7
fmt.Println(x, y)*/
//var x map[string]int = map[string]int{"hello": 3}
// y ist sozusagend nur ein anderer Name für x und deshalb wird x auch geändert
/*y := x
y["y"] = 100

fmt.Println(x, y)
*/
// Arrays sind immutable. Sclices und Maps nicht (glaube ich)
/*var x [2]int = [2]int{3, 4}
y := x
y[0] = 100
fmt.Println(x, y)

*/
/*
	x := 7
	y := &x
	// Gibt einem die reference von x (die Speicherreservierung oder den Pointer)
	fmt.Println(x, y)
	// * = dereference
	*y = 8
	fmt.Println(x, y)

	toChange := "hello"
	changeValue2(toChange)
	fmt.Println(toChange)
}

func changeValue(str *string) {
	*str = "chagend!"
}

func changeValue2(str string) {
	str = "changed!"
}

*/
/*
type Point struct {
	x int32
	y int32
}

type Circle struct {
	radius float64
	center *Point
}

func changeX(pt Point) {
	pt.x = 100
}
func main() {
	var p1 Point = Point{1, 2}
	var p2 Point = Point{-5, 7}
	fmt.Println(p1.x)
	fmt.Println(p1.y, p2.y)
	changeX(p1)
	fmt.Println(p1)
	c1 := Circle{4.56, &Point{4, 5}}
	fmt.Println(c1.center.x)
}
*/

//Methods
/*
type Student struct {
	name   string
	grades []int
	age    int
}

func (s *Student) getMaxGrade() int {
	curMax := 0

	for _, v := range s.grades {
		if curMax < v {
			curMax = v
		}
	}
	return curMax
}
func (s *Student) setAge(age int) {
	s.age = age
}
func main() {
	s1 := Student{"Jakob", []int{70, 90, 80, 85}, 19}
	s2 := Student{"Joe", []int{70, 32, 12, 51}, 19}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)
	average := s1.getAverageGrade()
	fmt.Println(average)
	average2 := s2.getAverageGrade()
	fmt.Println(average2)
	currentMax := s1.getMaxGrade()
	fmt.Println(currentMax)
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

*/
