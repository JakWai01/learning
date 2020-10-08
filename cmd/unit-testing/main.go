package main

func main() {

}

func sum(x, y int) int {

	// This is not covered -> coverage would get smaller
	// if x > 100 {
	// 	return 2
	// }
	return x + y
}
