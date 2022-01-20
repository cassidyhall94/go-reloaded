package main

import (
	"goreloaded"
	"reflect"
	"testing"
)

func Test_Hex(t *testing.T) {
	test1, err := goreloaded.Hex([]rune{'6', 'f'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, []rune{'1', '1', '1'}) {
		t.Logf("Hex('6f') failed, wanted %v, got %v", []rune{'1', '1', '1'}, test1)
		t.Fail()
	}

	test2, err := goreloaded.Hex([]rune{'4', '2'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, []rune{'6', '6'}) {
		t.Logf("Hex('42') failed, wanted %v, got %v", []rune{'6', '6'}, test2)
		t.Fail()
	}

	test3, err := goreloaded.Hex([]rune{'a', 'c'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test3, []rune{'1', '7', '2'}) {
		t.Logf("Hex('ac') failed, wanted %v, got %v", []rune{'1', '7', '2'}, test3)
		t.Fail()
	}

	test4, err := goreloaded.Hex([]rune{'z', 'x'})
	if err == nil {
		t.Logf("Hex([]rune{'z', 'x'}) failed, wanted err, got test4 == %v", test4)
		t.Fail()
	}
}

func Test_Bin(t *testing.T) {
	test1, err := goreloaded.Bin([]rune{'1', '1', '0', '0', '0', '0'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, []rune{'4', '8'}) {
		t.Logf("Bin('110000') failed, wanted %v, got %v", []rune{'4', '8'}, test1)
		t.Fail()
	}

	test2, err := goreloaded.Bin([]rune{'1', '0'})
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, []rune{'2'}) {
		t.Logf("Bin('10') failed, wanted %v, got %v", []rune{'2'}, test2)
		t.Fail()
	}

	test3, err := goreloaded.Bin([]rune{'4', '5'})
	if err == nil {
		t.Logf("Bin([]rune{'4', '5'}) failed, wanted err, got test3 == %v", test3)
		t.Fail()
	}
}

func Test_Cap(t *testing.T) {
	test1 := goreloaded.Cap([]rune{'i', 't'})
	if !reflect.DeepEqual(test1, []rune{'I', 't'}) {
		t.Logf("Cap([]rune{'i', 't'}) failed, wanted %v, got %v", []rune{'I', 't'}, test1)
		t.Fail()
	}

	test2 := goreloaded.Cap([]rune{'1'})
	if !reflect.DeepEqual(test2, []rune{'1'}) {
		t.Logf("Cap([]rune{'1'}) failed, wanted %v, got %v", []rune{'1'}, test2)
		t.Fail()
	}

	test3 := goreloaded.Cap([]rune{'t', 'i', 'm', 'e'})
	if !reflect.DeepEqual(test3, []rune{'T', 'i', 'm', 'e'}) {
		t.Logf("Cap([]rune{'t', 'i', 'm', 'e'}) failed, wanted %v, got %v", []rune{'T', 'i', 'm', 'e'}, test3)
		t.Fail()
	}
}

func Test_Lower(t *testing.T) {
	test1 := goreloaded.Lower([]rune{'A'})
	if !reflect.DeepEqual(test1, []rune{'a'}) {
		t.Logf("Lower([]rune{'A'}) failed, wanted %v, got %v", []rune{'a'}, test1)
		t.Fail()
	}

	test2 := goreloaded.Lower([]rune{'T', 'H', 'E'})
	if !reflect.DeepEqual(test2, []rune{'t', 'h', 'e'}) {
		t.Logf("Lower([]rune{'T', 'H', 'E'}) failed, wanted %v, got %v", []rune{'t', 'h', 'e'}, test2)
		t.Fail()
	}

	test3 := goreloaded.Lower([]rune{'!'})
	if !reflect.DeepEqual(test3, []rune{'!'}) {
		t.Logf("Lower([]rune{'!'}) failed, wanted %v, got %v", []rune{'!'}, test3)
		t.Fail()
	}
}

func Test_Upper(t *testing.T) {
	test1 := goreloaded.Upper([]rune{'a'})
	if !reflect.DeepEqual(test1, []rune{'A'}) {
		t.Logf("Upper([]rune{'a'}) failed, wanted %v, got %v", []rune{'A'}, test1)
		t.Fail()
	}

	test2 := goreloaded.Upper([]rune{'1'})
	if !reflect.DeepEqual(test2, []rune{'1'}) {
		t.Logf("Upper([]rune{'1'}) failed, wanted %v, got %v", []rune{'1'}, test2)
		t.Fail()
	}

	test3 := goreloaded.Upper([]rune{'t', 'i', 'm', 'e'})
	if !reflect.DeepEqual(test3, []rune{'T', 'I', 'M', 'E'}) {
		t.Logf("Upper([]rune{'t', 'i', 'm', 'e'}) failed, wanted %v, got %v", []rune{'T', 'I', 'M', 'E'}, test3)
		t.Fail()
	}
}

func Test_A(t *testing.T) {
	test1 := goreloaded.A([]rune{'a'})
	if !reflect.DeepEqual(test1, []rune{'a', 'n'}) {
		t.Logf("a([]rune{'a'}) failed, wanted %v, got %v", []rune{'a', 'n'}, test1)
		t.Fail()
	}

	test2 := goreloaded.A([]rune{'1'})
	if !reflect.DeepEqual(test2, []rune{}) {
		t.Logf("a([]rune{'1'}) failed, wanted %v, got %v", []rune{}, test2)
		t.Fail()
	}

	test3 := goreloaded.A([]rune{'A'})
	if !reflect.DeepEqual(test3, []rune{'A', 'n'}) {
		t.Logf("a([]rune{'A'}) failed, wanted %v, got %v", []rune{'A', 'n'}, test3)
		t.Fail()
	}
}

func Test_Punc(t *testing.T) {
	test1 := goreloaded.Punc([]rune{' ', ','})
	if !reflect.DeepEqual(test1, []rune{',', ' '}) {
		t.Logf("Punc([]rune{' ', ','}) failed, wanted %v, got %v", []rune{',', ' '}, test1)
		t.Fail()
	}

	test2 := goreloaded.Punc([]rune{' ', '.', '.', '.', ' '})
	if !reflect.DeepEqual(test2, []rune{'.', '.', '.', ' '}) {
		t.Logf("Punc([]rune{' ', '.', '.', '.', ' '}) failed, wanted %v, got %v", []rune{'.', '.', '.', ' '}, test2)
		t.Fail()
	}

	test3 := goreloaded.Punc([]rune{' ', '.'})
	if !reflect.DeepEqual(test3, []rune{'.', ' '}) {
		t.Logf("Punc([]rune{' ', '.'}) failed, wanted %v, got %v", []rune{'.', ' '}, test3)
		t.Fail()
	}

	test4 := goreloaded.Punc([]rune{' ', '!'})
	if !reflect.DeepEqual(test4, []rune{'!', ' '}) {
		t.Logf("Punc([]rune{' ', '!'}) failed, wanted %v, got %v", []rune{'!', ' '}, test4)
		t.Fail()
	}

	test5 := goreloaded.Punc([]rune{' ', '?'})
	if !reflect.DeepEqual(test5, []rune{'?', ' '}) {
		t.Logf("Punc([]rune{' ', '?'}) failed, wanted %v, got %v", []rune{'?', ' '}, test5)
		t.Fail()
	}

	test6 := goreloaded.Punc([]rune{' ', ':'})
	if !reflect.DeepEqual(test6, []rune{':', ' '}) {
		t.Logf("Punc([]rune{' ', ':'}) failed, wanted %v, got %v", []rune{':', ' '}, test6)
		t.Fail()
	}

	test7 := goreloaded.Punc([]rune{' ', ';'})
	if !reflect.DeepEqual(test7, []rune{';', ' '}) {
		t.Logf("Punc([]rune{' ', ';'}) failed, wanted %v, got %v", []rune{';', ' '}, test7)
		t.Fail()
	}

	test8 := goreloaded.Punc([]rune{' ', '!', '?', ' '})
	if !reflect.DeepEqual(test8, []rune{'!', '?', ' '}) {
		t.Logf("Punc([]rune{' ', '!', '?', ' '}) failed, wanted %v, got %v", []rune{'!', '?', ' '}, test8)
		t.Fail()
	}

	test9 := goreloaded.Punc([]rune{' ', 'a'})
	if !reflect.DeepEqual(test9, []rune{}) {
		t.Logf("Punc([]rune{' ', 'a'}) failed, wanted %v, got %v", []rune{}, test9)
		t.Fail()
	}
}

func Test_Apos(t *testing.T) {
	test1 := goreloaded.Apos([]rune{32, 39, 39, 32})
	if !reflect.DeepEqual(test1, []rune{39, 39}) {
		t.Logf("Apos([]rune{' ', ''', ''', ' '}) failed, wanted %v, got %v", []rune{39, 39}, test1)
		t.Fail()
	}

	test2 := goreloaded.Apos([]rune{39, 39, 32})
	if !reflect.DeepEqual(test2, []rune{39, 39}) {
		t.Logf("Apos([]rune{''', ''', ' '}) failed, wanted %v, got %v", []rune{39, 39}, test2)
		t.Fail()
	}

	test3 := goreloaded.Apos([]rune{32, 39, 39})
	if !reflect.DeepEqual(test3, []rune{39, 39}) {
		t.Logf("Apos([]rune{' ', ''', '''}) failed, wanted %v, got %v", []rune{39, 39}, test3)
		t.Fail()
	}

	test4 := goreloaded.Apos([]rune{32, 'a'})
	if !reflect.DeepEqual(test4, []rune{}) {
		t.Logf("Apos([]rune{' ', 'a'}) failed, wanted %v, got %v", []rune{}, test4)
		t.Fail()
	}
}

func Test_Acute(t *testing.T) {
	test1 := goreloaded.Acute([]rune{32, '´', '´', 32})
	if !reflect.DeepEqual(test1, []rune{'´', '´'}) {
		t.Logf("Acute([]rune{' ', '´', '´', ' '}) failed, wanted %v, got %v", []rune{'´', '´'}, test1)
		t.Fail()
	}

	test2 := goreloaded.Acute([]rune{'´', '´', 32})
	if !reflect.DeepEqual(test2, []rune{'´', '´'}) {
		t.Logf("Acute([]rune{'´', '´', ' '}) failed, wanted %v, got %v", []rune{'´', '´'}, test2)
		t.Fail()
	}

	test3 := goreloaded.Acute([]rune{32, '´', '´'})
	if !reflect.DeepEqual(test3, []rune{'´', '´'}) {
		t.Logf("Acute([]rune{' ', '´', '´'}) failed, wanted %v, got %v", []rune{'´', '´'}, test3)
		t.Fail()
	}

	test4 := goreloaded.Acute([]rune{32, 'a'})
	if !reflect.DeepEqual(test4, []rune{}) {
		t.Logf("Acute([]rune{' ', 'a'}) failed, wanted %v, got %v", []rune{}, test4)
		t.Fail()
	}
}
