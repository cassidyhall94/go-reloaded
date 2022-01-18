package main

import (
	"strings"
)

// func to look for (hex) and trigger the hex function
// loop through string and replace number before (hex) to dec instead using hex function

func checkhex(s string) (string, error) {
	splitS := strings.Split(s, " ")
	hexvar := "(hex)"

	for i := range splitS {
		if len(splitS) == i+3 {
			break
		} else if splitS[i] == hexvar {
			result, err := hex([]rune(splitS[i-1]))
			if err != nil {
				return "", err
			}
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " "), nil
}

func checkbin(s string) (string, error) {
	splitS := strings.Split(s, " ")
	binvar := "(bin)"

	for i := range splitS {
		if len(splitS) == i+3 {
			break
		} else if splitS[i] == binvar {
			result, err := bin([]rune(splitS[i-1]))
			if err != nil {
				return "", err
			}
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " "), nil
}

// checkupper looks for any instances of (up) and replaces the string before with uppercase letters
// checkupper looks for any instances of (up, n) and replaces n strings before with uppercase letters

func checkupper(s string) string {
	splitS := strings.Split(s, " ") // splits string by a space

	for i := range splitS { 
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		// we extract n from (up, n) in the string using a var called n
		// instead of i, we use n to represent a range of indexes of the string that is transformed
		// when we want to get our range of splitS, we use splitS[i-n:i]
		if isLower(strings.Join(splitS[i-n:i], "")) && strings.Contains(splitS[i], "(up)") {
			result := upper([]rune(splitS[i-n:i]))
			splitS[i-n:i] = string(result)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
}

func isLower(s string) bool {
	counter := 0
	for _, stringrune := range s {
		if (stringrune >= rune(97)) && (stringrune <= rune(122)) {
			counter++
		} else {
			return false
		}
	}
	return true
}

func checklower(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		if isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(low)") {
			result := lower([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
}

func isUpper(s string) bool {
	counter := 0
	for _, stringrune := range []rune(s) {
		if (stringrune >= rune(65)) && (stringrune <= rune(90)) {
			counter++
		} else {
			return false
		}
	}
	return true
}

func removeIndex(s []string, index int) []string {
	ret := make([]string, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
