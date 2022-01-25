package main

import (
	"reflect"
	"testing"
)

func Test_Hex(t *testing.T) {
	test1, err := hex([]rune{'6', 'f'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, []rune{'1', '1', '1'}) {
		t.Logf("Hex('6f') failed, wanted %v, got %v", []rune{'1', '1', '1'}, test1)
		t.Fail()
	}

	test2, err := hex([]rune{'4', '2'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, []rune{'6', '6'}) {
		t.Logf("Hex('42') failed, wanted %v, got %v", []rune{'6', '6'}, test2)
		t.Fail()
	}

	test3, err := hex([]rune{'a', 'c'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test3, []rune{'1', '7', '2'}) {
		t.Logf("Hex('ac') failed, wanted %v, got %v", []rune{'1', '7', '2'}, test3)
		t.Fail()
	}

	test4, err := hex([]rune{'z', 'x'})
	if err == nil {
		t.Logf("hex([]rune{'z', 'x'}) failed, wanted err, got test4 == %v", test4)
		t.Fail()
	}
}

func Test_Bin(t *testing.T) {
	test1, err := bin([]rune{'1', '1', '0', '0', '0', '0'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, []rune{'4', '8'}) {
		t.Logf("Bin('110000') failed, wanted %v, got %v", []rune{'4', '8'}, test1)
		t.Fail()
	}

	test2, err := bin([]rune{'1', '0'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, []rune{'2'}) {
		t.Logf("Bin('10') failed, wanted %v, got %v", []rune{'2'}, test2)
		t.Fail()
	}

	test3, err := bin([]rune{'4', '5'})
	if err == nil {
		t.Logf("bin([]rune{'4', '5'}) failed, wanted err, got test3 == %v", test3)
		t.Fail()
	}
}

func Test_Cap(t *testing.T) {
	test1 := cap([]rune{'i', 't'})
	if !reflect.DeepEqual(test1, []rune{'I', 't'}) {
		t.Logf("cap([]rune{'i', 't'}) failed, wanted %v, got %v", []rune{'I', 't'}, test1)
		t.Fail()
	}

	test2 := cap([]rune{'1'})
	if !reflect.DeepEqual(test2, []rune{'1'}) {
		t.Logf("cap([]rune{'1'}) failed, wanted %v, got %v", []rune{'1'}, test2)
		t.Fail()
	}

	test3 := cap([]rune{'t', 'i', 'm', 'e'})
	if !reflect.DeepEqual(test3, []rune{'T', 'i', 'm', 'e'}) {
		t.Logf("cap([]rune{'t', 'i', 'm', 'e'}) failed, wanted %v, got %v", []rune{'T', 'i', 'm', 'e'}, test3)
		t.Fail()
	}
}

func Test_Lower(t *testing.T) {
	test1 := lower([]rune{'A'})
	if !reflect.DeepEqual(test1, []rune{'a'}) {
		t.Logf("lower([]rune{'A'}) failed, wanted %v, got %v", []rune{'a'}, test1)
		t.Fail()
	}

	test2 := lower([]rune{'T', 'H', 'E'})
	if !reflect.DeepEqual(test2, []rune{'t', 'h', 'e'}) {
		t.Logf("lower([]rune{'T', 'H', 'E'}) failed, wanted %v, got %v", []rune{'t', 'h', 'e'}, test2)
		t.Fail()
	}

	test3 := lower([]rune{'!'})
	if !reflect.DeepEqual(test3, []rune{'!'}) {
		t.Logf("lower([]rune{'!'}) failed, wanted %v, got %v", []rune{'!'}, test3)
		t.Fail()
	}
}

func Test_Upper(t *testing.T) {
	test1 := upper([]rune{'a'})
	if !reflect.DeepEqual(test1, []rune{'A'}) {
		t.Logf("upper([]rune{'a'}) failed, wanted %v, got %v", []rune{'A'}, test1)
		t.Fail()
	}

	test2 := upper([]rune{'1'})
	if !reflect.DeepEqual(test2, []rune{'1'}) {
		t.Logf("upper([]rune{'1'}) failed, wanted %v, got %v", []rune{'1'}, test2)
		t.Fail()
	}

	test3 := upper([]rune{'t', 'i', 'm', 'e'})
	if !reflect.DeepEqual(test3, []rune{'T', 'I', 'M', 'E'}) {
		t.Logf("upper([]rune{'t', 'i', 'm', 'e'}) failed, wanted %v, got %v", []rune{'T', 'I', 'M', 'E'}, test3)
		t.Fail()
	}
}

func Test_A(t *testing.T) {
	test1 := a([]rune{'a'})
	if !reflect.DeepEqual(test1, []rune{'a', 'n'}) {
		t.Logf("a([]rune{'a'}) failed, wanted %v, got %v", []rune{'a', 'n'}, test1)
		t.Fail()
	}

	test2 := a([]rune{'1'})
	if !reflect.DeepEqual(test2, []rune{}) {
		t.Logf("a([]rune{'1'}) failed, wanted %v, got %v", []rune{}, test2)
		t.Fail()
	}

	test3 := a([]rune{'A'})
	if !reflect.DeepEqual(test3, []rune{'A', 'n'}) {
		t.Logf("a([]rune{'A'}) failed, wanted %v, got %v", []rune{'A', 'n'}, test3)
		t.Fail()
	}
}

// func Test_Punc(t *testing.T) {
// 	test1 := punc([]rune{' ', ','})
// 	if !reflect.DeepEqual(test1, []rune{',', ' '}) {
// 		t.Logf("punc([]rune{' ', ','}) failed, wanted %v, got %v", []rune{',', ' '}, test1)
// 		t.Fail()
// 	}

// 	test2 := punc([]rune{' ', '.', '.', '.', ' '})
// 	if !reflect.DeepEqual(test2, []rune{'.', '.', '.', ' '}) {
// 		t.Logf("punc([]rune{' ', '.', '.', '.', ' '}) failed, wanted %v, got %v", []rune{'.', '.', '.', ' '}, test2)
// 		t.Fail()
// 	}

// 	test3 := punc([]rune{' ', '.'})
// 	if !reflect.DeepEqual(test3, []rune{'.', ' '}) {
// 		t.Logf("punc([]rune{' ', '.'}) failed, wanted %v, got %v", []rune{'.', ' '}, test3)
// 		t.Fail()
// 	}

// 	test4 := punc([]rune{' ', '!'})
// 	if !reflect.DeepEqual(test4, []rune{'!', ' '}) {
// 		t.Logf("punc([]rune{' ', '!'}) failed, wanted %v, got %v", []rune{'!', ' '}, test4)
// 		t.Fail()
// 	}

// 	test5 := punc([]rune{' ', '?'})
// 	if !reflect.DeepEqual(test5, []rune{'?', ' '}) {
// 		t.Logf("punc([]rune{' ', '?'}) failed, wanted %v, got %v", []rune{'?', ' '}, test5)
// 		t.Fail()
// 	}

// 	test6 := punc([]rune{' ', ':'})
// 	if !reflect.DeepEqual(test6, []rune{':', ' '}) {
// 		t.Logf("punc([]rune{' ', ':'}) failed, wanted %v, got %v", []rune{':', ' '}, test6)
// 		t.Fail()
// 	}

// 	test7 := punc([]rune{' ', ';'})
// 	if !reflect.DeepEqual(test7, []rune{';', ' '}) {
// 		t.Logf("punc([]rune{' ', ';'}) failed, wanted %v, got %v", []rune{';', ' '}, test7)
// 		t.Fail()
// 	}

// 	test8 := punc([]rune{' ', '!', '?', ' '})
// 	if !reflect.DeepEqual(test8, []rune{'!', '?', ' '}) {
// 		t.Logf("punc([]rune{' ', '!', '?', ' '}) failed, wanted %v, got %v", []rune{'!', '?', ' '}, test8)
// 		t.Fail()
// 	}

// 	test9 := punc([]rune{' ', 'a'})
// 	if !reflect.DeepEqual(test9, []rune{}) {
// 		t.Logf("punc([]rune{' ', 'a'}) failed, wanted %v, got %v", []rune{}, test9)
// 		t.Fail()
// 	}
// }

// func Test_Apos(t *testing.T) {
// 	test1 := apos([]rune{32, 39, 39, 32})
// 	if !reflect.DeepEqual(test1, []rune{39, 39}) {
// 		t.Logf("apos([]rune{' ', ''', ''', ' '}) failed, wanted %v, got %v", []rune{39, 39}, test1)
// 		t.Fail()
// 	}

// 	test2 := apos([]rune{39, 39, 32})
// 	if !reflect.DeepEqual(test2, []rune{39, 39}) {
// 		t.Logf("apos([]rune{''', ''', ' '}) failed, wanted %v, got %v", []rune{39, 39}, test2)
// 		t.Fail()
// 	}

// 	test3 := apos([]rune{32, 39, 39})
// 	if !reflect.DeepEqual(test3, []rune{39, 39}) {
// 		t.Logf("apos([]rune{' ', ''', '''}) failed, wanted %v, got %v", []rune{39, 39}, test3)
// 		t.Fail()
// 	}

// 	test4 := apos([]rune{32, 'a'})
// 	if !reflect.DeepEqual(test4, []rune{}) {
// 		t.Logf("apos([]rune{' ', 'a'}) failed, wanted %v, got %v", []rune{}, test4)
// 		t.Fail()
// 	}
// }

// func Test_Acute(t *testing.T) {
// 	test1 := acute([]rune{32, '´', '´', 32})
// 	if !reflect.DeepEqual(test1, []rune{'´', '´'}) {
// 		t.Logf("acute([]rune{' ', '´', '´', ' '}) failed, wanted %v, got %v", []rune{'´', '´'}, test1)
// 		t.Fail()
// 	}

// 	test2 := acute([]rune{'´', '´', 32})
// 	if !reflect.DeepEqual(test2, []rune{'´', '´'}) {
// 		t.Logf("acute([]rune{'´', '´', ' '}) failed, wanted %v, got %v", []rune{'´', '´'}, test2)
// 		t.Fail()
// 	}

// 	test3 := acute([]rune{32, '´', '´'})
// 	if !reflect.DeepEqual(test3, []rune{'´', '´'}) {
// 		t.Logf("acute([]rune{' ', '´', '´'}) failed, wanted %v, got %v", []rune{'´', '´'}, test3)
// 		t.Fail()
// 	}

// 	test4 := acute([]rune{32, 'a'})
// 	if !reflect.DeepEqual(test4, []rune{}) {
// 		t.Logf("acute([]rune{' ', 'a'}) failed, wanted %v, got %v", []rune{}, test4)
// 		t.Fail()
// 	}
// }
