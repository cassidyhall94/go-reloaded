package main

import (
	"strconv"
)

func hex(hexrunes []rune) ([]rune, error) {
	num, err := strconv.ParseInt(string(hexrunes), 16, 64)
	if err != nil {
		return nil, err
	}
	return []rune(strconv.FormatInt(num, 10)), nil
}

func bin(binrunes []rune) ([]rune, error) {
	num, err := strconv.ParseInt(string(binrunes), 2, 64)
	if err != nil {
		return nil, err
	}
	return []rune(strconv.FormatInt(num, 10)), nil
}

func IsAlpha(s string) bool {
	counter := 0
	for _, stringrune := range s {
		if ((stringrune >= rune(65)) && (stringrune <= rune(90))) || ((stringrune >= rune(97)) && (stringrune <= rune(122))) || (stringrune >= rune(48)) && (stringrune <= rune(57)) {
			counter++
		} else {
			return false
		}
	}
	return true
}

func cap(caprunes []rune) []rune {
	result := caprunes
	if string(result) != string([]rune{}) {
		if result[0] >= 97 && result[0] <= 122 {
			result[0] -= 32
		}
	}
	for i, stringrune := range caprunes {
		if i != 0 {
			if IsAlpha(string(result[i-1])) {
				if stringrune >= 65 && stringrune <= 90 {
					result[i] += 32
				}
			}
			if !(IsAlpha(string(result[i-1]))) {
				if stringrune >= 97 && stringrune <= 122 {
					result[i] -= 32
				}
			}
		}
	}
	return result
}

func lower(lowerrunes []rune) []rune {
	for i := range lowerrunes {
		if lowerrunes[i] >= rune(65) && lowerrunes[i] <= rune(90) {
			lowerrunes[i] = lowerrunes[i] + 32
		}
	}
	return lowerrunes
}

func upper(upperrunes []rune) []rune {
	for i := range upperrunes {
		if upperrunes[i] >= rune(97) && upperrunes[i] <= rune(122) {
			upperrunes[i] = upperrunes[i] - 32
		}
	}
	return upperrunes
}

func a(arunes []rune) []rune {
	result := arunes
	for i := range result {
		if result[i] == rune(97) {
			return []rune{'a', 'n'}
		}
		if result[i] == rune(65) {
			return []rune{'A', 'n'}
		}
	}
	return []rune{}
}

// starting with an empty slice of runes as result
// for each element in the rune slice, if there's a space, ignore it.
// if there's a puncuation mark, append to result
// add space to end of result and return result

func punc(puncrunes []rune) []rune {
	result := []rune{}
	for i := range puncrunes {
		if puncrunes[i] == rune(46) {
			result = append(result, rune(46))
		}
		if puncrunes[i] == rune(44) {
			result = append(result, rune(44))
		}
		if puncrunes[i] == rune(33) {
			result = append(result, rune(33))
		}
		if puncrunes[i] == rune(63) {
			result = append(result, rune(63))
		}
		if puncrunes[i] == rune(58) {
			result = append(result, rune(58))
		}
		if puncrunes[i] == rune(59) {
			result = append(result, rune(59))
		}
	}
	if string(result) != "" {
		result = append(result, rune(32))
	}
	return result
}

// starting with an empty slice of runes as result
// for each element in the rune slice, if there's a space, ignore it.
// if there's a puncuation mark, append to result
// add space to end of result and return result

func apos(aposrunes []rune) []rune {
	result := []rune{}
	for i := range aposrunes {
		if aposrunes[i] == rune(39) {
			result = append(result, rune(39))
		}
	}
	// result = append(result)
	return result
}

func acute(acuterunes []rune) []rune {
	result := []rune{}
	for i := range acuterunes {
		if acuterunes[i] == '´' {
			result = append(result, '´')
		}
	}
	// result = append(result)
	return result
}