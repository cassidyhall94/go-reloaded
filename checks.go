package main

import "strings"

// how to nest functions
//func main() {
//	fmt.Println(mult(2, 3, 2, 4))
//}
//
//func mult(a, b, c, d int) int {
//	sum2 := add(c, d)
//	return add(a, b) * sum2
//}
//
//func add(e, f int) int {
//	sum := e + f
//	return sum
//}


// func to look for (hex) and trigger the hex function
// loop through string and replace number before (hex) to dec instead using hex function

func checkhex(s string) string {
	splitS := strings.Split(s, " ")

	return strings.Join(splitS, " ")
}
