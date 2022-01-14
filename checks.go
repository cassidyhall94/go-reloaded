package main

import "strings"

// func to look for (hex) and trigger the hex function
// loop through string and replace number before (hex) to dec instead using hex function

func checkhex(s string) string {
	splitS := strings.Split(s, " ")

	return strings.Join(splitS, " ")
}
