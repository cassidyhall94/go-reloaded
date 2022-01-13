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
	result := lowerrunes
	for i := range result {
		if result[i] >= rune(48) && result[i] <= rune(57) {
			return []rune{}
		}
		if result[i] >= rune(65) && result[i] <= rune(90) {
			result[i] = result[i] + 32
		}
	}
	return result
}

func upper(upperrunes []rune) []rune {
	result := upperrunes
	for i := range result {
		if result[i] >= rune(97) && result[i] <= rune(122) {
			result[i] = result[i] - 32
		}
	}
	return result
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
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(46) {
			result = append(result, rune(46))
		}
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(44) {
			return []rune{44, 32}
		}
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(33) {
			return []rune{33, 32}
		}
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(63) {
			return []rune{63, 32}
		}
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(58) {
			return []rune{58, 32}
		}
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(59) {
			return []rune{59, 32}
		}
		if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(39) {
			return []rune{39, 32}
		}
	//	if puncrunes[i] == rune(32) && puncrunes[i+1] == rune(46) && puncrunes[i+2] == rune(46) && puncrunes[i+3] == rune(46) && puncrunes[i+4] == rune(32) {
	//		return []rune{46, 46, 46, 32}
	//	}
	}
	result = append(result, rune(32))
	return []rune{}
}

func apos(puncrunes []rune) ([]rune, error) {
	return nil, nil
}
