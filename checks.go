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

func checkupper(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		if isLower(splitS[i-1]) && strings.Contains(splitS[i], "(up") {
			result := upper([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
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

func checkcap(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		if isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(cap)") {
			result := cap([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		}
	}
	return strings.Join(splitS, " ")
}

// loop through string, looking for (low/up/cap, n) and when found it implements the respective func for n strings before it

func checkN(N string) string {
	splitS := strings.Split(N, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i+1 {
			break
		}
		if IsNumeric(splitS[i+1]) && strings.Contains(splitS[i], ", ") {
			if strings.Contains(splitS[i], "(low") {
				result := lower([]rune(splitS[i-1]))
				splitS[i] = string(result)
				splitS = removeIndex(splitS, i)
				splitS = removeIndex(splitS, i+1)
			}
			if strings.Contains(splitS[i], "(up") {
				result := upper([]rune(splitS[i-1]))
				splitS[i] = string(result)
				splitS = removeIndex(splitS, i)
				splitS = removeIndex(splitS, i+1)
			}
			// if strings.Contains(splitS[i], "(cap") {
			// 	result := cap([]rune(splitS[i-1]))
			// 	splitS[i] = string(result)
			// 	splitS = removeIndex(splitS, i)
			// 	splitS = removeIndex(splitS, i+1)
			// }

		}
	}
	return strings.Join(splitS, " ")
}

// func isCap(s string) bool {
// 	counter := 0
// 	for _, stringrune := range s {
// 		if (stringrune >= rune(65)) && (stringrune <= rune(90)) {
// 			continue
// 		}
// 			if (stringrune[i+1] >= rune(97)) && (stringrune[i+1] <= rune(122)) {
// 			counter++
// 		} else {
// 			return false
// 		}
// 	}
// 	return true
// }

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
