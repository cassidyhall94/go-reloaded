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

func cap(caprunes []rune) ([]rune, error) {
	result := caprunes
	// if caprunes != []rune{} {
	// 	if result[0] >= 97 && result[0] <= 122 {
	// 		result[0] -= 32
	// 	}
	// }
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
	return result, nil
}

func lower(lowerrunes []rune) ([]rune, error) {
	result := lowerrunes
	for i := range result {
		if result[i] >= rune(65) && result[i] <= rune(90) {
			result[i] = result[i] + 32
		}
	}
	return result, nil
}

func upper(upperrunes []rune) ([]rune, error) {
	result := upperrunes
	for i := range result {
		if result[i] >= rune(97) && result[i] <= rune(122) {
			result[i] = result[i] - 32
		}
	}
	return result, nil
}

func a(arunes []rune) ([]rune, error) {
	return nil, nil
}

func punc(puncrunes []rune) ([]rune, error) {
	return nil, nil
}

func apos(puncrunes []rune) ([]rune, error) {
	return nil, nil
}