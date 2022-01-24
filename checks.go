package main

import (
	"strconv"
	"strings"
)

// func to look for (hex) and trigger the hex function
// loop through string and replace number before (hex) to dec instead using hex function

func checkHex(s string) (string, error) {
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

func checkBin(s string) (string, error) {
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

func checkUpper(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i-1 || len(splitS) == i {
			break
		}
		if isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(up)") || isLower(splitS[i-1]) && strings.Contains(splitS[i], "(up)") {
			result := upper([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		} else if isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(up,") || isLower(splitS[i-1]) && strings.Contains(splitS[i], "(up,") {
			indexn := splitS[i+1]
			n := indexn[:len(indexn)-1]
			N, _ := strconv.Atoi(n)
			counter := 0
			for j := len(splitS[:i]); counter < N; j-- {
				result := upper([]rune(splitS[j-1]))
				splitS[j-1] = string(result)
				counter++
			}
			splitS = removeIndex(splitS, i)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
}

func checkLower(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i-1 || len(splitS) == i {
			break
		}
		if isLower(splitS[i-1]) && strings.Contains(splitS[i], "(low)") || isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(low)") {
			result := lower([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		} else if isLower(splitS[i-1]) && strings.Contains(splitS[i], "(low,") || isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(low,") {
			indexn := splitS[i+1]
			n := indexn[:len(indexn)-1]
			N, _ := strconv.Atoi(n)
			counter := 0
			for j := len(splitS[:i]); counter < N; j-- {
				result := lower([]rune(splitS[j-1]))
				splitS[j-1] = string(result)
				counter++
			}
			splitS = removeIndex(splitS, i)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
}

func checkCap(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i || len(splitS) == i-1 {
			break
		}
		if isCap(splitS[i-1]) && strings.Contains(splitS[i], "(cap)") {
			result := cap([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		} else if isCap(splitS[i-1]) && strings.Contains(splitS[i], "(cap,") {
			indexn := splitS[i+1]
			n := indexn[:len(indexn)-1]
			N, _ := strconv.Atoi(n)
			counter := 0
			for j := len(splitS[:i]); counter < N; j-- {
				result := cap([]rune(splitS[j-1]))
				splitS[j-1] = string(result)
				counter++
			}
			splitS = removeIndex(splitS, i)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
}

func checkA(s string) string {
	// check if 'a' needs to be transformed to 'an'
	// if there is a vowel or an 'h' after the 'a' then transform
	// loop through string to find an 'a' or 'A', then loop through string again to see if i+1 is a vowel/h
	// transform the 'a' and join the string together to print full transformed string
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		if isA(splitS[i]) && isVowel(splitS[i+1]) {
			result := a([]rune(splitS[i]))
			splitS[i] = string(result)
		}
	}
	return strings.Join(splitS, " ")
}

func checkPunc(s string) string {
	// remove spaces before the mark and ensure there's a space after the mark
	// loop through string by []rune to find marks with space in front of them and remove it
	srunes := []rune(s)
	removed := 0
	for i, r := range srunes {
		if isPunc(r) {
			if srunes[i-(1+removed)] == ' ' {
				srunes = removeIndexrune(srunes, i-(1+removed))
				removed++
			}
		}
	}
	for i, r := range srunes {
		if isPunc(r) {
			if srunes[i+1] != ' ' {
				srunes = insert(srunes, rune(32), i+1)
			}
		}
	}
	return string(srunes)
}

func isPunc(s rune) bool {
	switch s {

	case rune(46):
		return true
	case rune(44):
		return true
	case rune(33):
		return true
	case rune(63):
		return true
	case rune(58):
		return true
	case rune(59):
		return true
	default:
		return false
	}
}

func insert(a []rune, c rune, i int) []rune {
	return append(a[:i], append([]rune{c}, a[i:]...)...)
}

func isCap(s string) bool {
	counter := 0
	for _, stringrune := range s {
		if (stringrune >= rune(65)) && (stringrune <= rune(90)) {
			continue
		}
		if (stringrune >= rune(97)) && (stringrune <= rune(122)) {
			counter++
		} else {
			return false
		}
	}
	return true
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

func isUpper(s string) bool {
	counter := 0
	for _, stringrune := range s {
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

func removeIndexrune(s []rune, index int) []rune {
	ret := make([]rune, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func IsNumeric(s string) bool {
	counter := 0
	for _, stringrune := range s {
		if (stringrune >= rune(48)) && (stringrune <= rune(57)) {
			counter++
		} else {
			return false
		}
	}
	return true
}

func isA(s string) bool {
	counter := 0
	for _, stringrune := range s {
		if (stringrune == rune(97)) || (stringrune == rune(65)) {
			counter++
		} else {
			return false
		}
	}
	return true
}

func isVowel(s string) bool {
	switch string([]rune(s)[0]) {
	case "o":
		return true
	case "a":
		return true
	case "i":
		return true
	case "u":
		return true
	case "e":
		return true
	case "h":
		return true
	case "y":
		return true
	case "O":
		return true
	case "A":
		return true
	case "I":
		return true
	case "U":
		return true
	case "E":
		return true
	case "H":
		return true
	case "Y":
		return true
	default:
		return false
	}
}
