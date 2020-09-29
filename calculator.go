package main

import (
	"fmt"
	"flag"

)

func main() {
	
	a := flag.Float64("a", 1, "First parameter")
	b := flag.Float64("b", 2, "Second parameter")
	c := flag.String("c", "Add", "Operation (Add, Substract, Multiply, Divide)")
	flag.Parse()
 
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	
	switch(*c) {
	case "Add": fmt.Println(add(*a, *b))
	case "Substract": fmt.Println(substract(*a,*b))
	case "Multiply": fmt.Println(multiply(*a,*b))
	case "Divide": fmt.Println(divide(*a, *b))
	default: fmt.Println("This is not a valid operation")
	}
	
}

func add(a, b float64) float64{
 	return (a+b)
 }

func substract(a, b float64) float64 {
	 return a-b
 }

 func multiply(a, b float64) float64  {
 	return a*b
 }

 func divide(a, b float64) float64 {
 	return a/b
 }