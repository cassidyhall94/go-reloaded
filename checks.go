package goreloaded

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
				splitS = removeIndex(splitS, i)
				counter++
			}
		}
	}
	return strings.Join(splitS, " ")
}

func checkLower(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		if isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(low)") || isLower(splitS[i-1]) && strings.Contains(splitS[i], "(low)") {
			result := lower([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		} else if isUpper(splitS[i-1]) && strings.Contains(splitS[i], "(low,") || isLower(splitS[i-1]) && strings.Contains(splitS[i], "(low,") {
			indexn := splitS[i+1]
			n := indexn[:len(indexn)-1]
			// n := splitS[i+1][0 : len(splitS[i+1])-1]
			N, _ := strconv.Atoi(n)
			for i := N; i > N; i-- {
				result := lower([]rune(string(splitS[i-N])))
				splitS[i-N] = string(result)
				splitS = removeIndex(splitS, i)
				splitS = removeIndex(splitS, i+1)
			}
		}
	}
	return strings.Join(splitS, " ")
}

func checkCap(s string) string {
	splitS := strings.Split(s, " ")

	for i := range splitS {
		if i == 0 {
			continue
		} else if len(splitS) == i {
			break
		}
		if isCap(splitS[i-1]) && strings.Contains(splitS[i], "(cap)") {
			result := cap([]rune(splitS[i-1]))
			splitS[i-1] = string(result)
			splitS = removeIndex(splitS, i)
		} else if isCap(splitS[i-1]) && strings.Contains(splitS[i], "(cap,") {
			n := splitS[i+1]
			N, _ := strconv.Atoi(n)
			for i := N; i > N; i-- {
				result := cap([]rune(string(splitS[i-N])))
				splitS[i-N] = string(result)
				splitS = removeIndex(splitS, i)
				splitS = removeIndex(splitS, i+1)
			}
		}
	}
	return strings.Join(splitS, " ")
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
