package main

import (
	"reflect"
	"testing"
)

// https://git.learn.01founders.co/root/public/src/branch/master/subjects/go-reloaded/audit

func Test_checkHex(t *testing.T) {
	test1, err := checkHex("42 (hex)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, "66") {
		t.Logf("checkhex(\"42 (hex)\") failed, wanted \"%v\", got \"%v\"", "66", test1)
		t.Fail()
	}

	test2, err := checkHex("42 (hox)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, "42 (hox)") {
		t.Logf("checkhex(\"42 (hox)\") failed, wanted \"%v\", got \"%v\"", "42 (hox)", test2)
		t.Fail()
	}

	test3, err := checkHex("42 ()")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test3, "42 ()") {
		t.Logf("checkhex(\"42 ()\") failed, wanted \"%v\", got \"%v\"", "42 ()", test3)
		t.Fail()
	}

	test4, err := checkHex("42 hex")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test4, "42 hex") {
		t.Logf("checkhex(\"42 hex\") failed, wanted \"%v\", got \"%v\"", "42 (hex)", test4)
		t.Fail()
	}

	test5, err := checkHex("xx (hex)")
	if err == nil {
		t.Logf("xx (hex) failed, wanted err, got test5 == %v", test5)
		t.Fail()
	}

	test6, err := checkHex("1e (hex)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test6, "30") {
		t.Logf("checkhex(\"1e (hex)\") failed, wanted \"%v\", got \"%v\"", "30", test6)
		t.Fail()
	}

	test7In := "Simply add 42 (hex) and 10 (up) and you will see the result is 68"
	test7Want := "Simply add 66 and 10 (up) and you will see the result is 68"
	test7, err := checkHex(test7In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkhex(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}

	test8In := "Simply add 42 (hex) and 10 (up) and you 1e (hex) will see the result is 68"
	test8Want := "Simply add 66 and 10 (up) and you 30 will see the result is 68"
	test8, err := checkHex(test8In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test8, test8Want) {
		t.Logf("checkhex(\"%s\") failed, wanted \"%s\", got \"%s\"", test8In, test8Want, test8)
		t.Fail()
	}
}

func Test_checkBin(t *testing.T) {
	test1, err := checkBin("10 (bin)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, "2") {
		t.Logf("checkbin(\"10 (bin)\") failed, wanted \"%v\", got \"%v\"", "42", test1)
		t.Fail()
	}

	test2, err := checkBin("42 (hox)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, "42 (hox)") {
		t.Logf("checkbin(\"42 (hox)\") failed, wanted \"%v\", got \"%v\"", "42 (hox)", test2)
		t.Fail()
	}

	test3, err := checkBin("42 ()")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test3, "42 ()") {
		t.Logf("checkbin(\"42 ()\") failed, wanted \"%v\", got \"%v\"", "42 ()", test3)
		t.Fail()
	}

	test4, err := checkBin("42 bin")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test4, "42 bin") {
		t.Logf("checkbin(\"42 bin\") failed, wanted \"%v\", got \"%v\"", "42 (bin)", test4)
		t.Fail()
	}

	test5, err := checkBin("xx (bin)")
	if err == nil {
		t.Logf("xx (bin) failed, wanted err, got test5 == %v", test5)
		t.Fail()
	}

	test6, err := checkBin("110 (bin)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test6, "6") {
		t.Logf("checkbin(\"110 (bin)\") failed, wanted \"%v\", got \"%v\"", "6", test6)
		t.Fail()
	}

	test7In := "Simply add 10 (bin) and 66 you will see the result is 68"
	test7Want := "Simply add 2 and 66 you will see the result is 68"
	test7, err := checkBin(test7In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkbin(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}

	test8In := "Simply add 100110 (bin) and 11110 (bin) and you will see the result is 68"
	test8Want := "Simply add 38 and 30 and you will see the result is 68"
	test8, err := checkBin(test8In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test8, test8Want) {
		t.Logf("checkbin(\"%s\") failed, wanted \"%s\", got \"%s\"", test8In, test8Want, test8)
		t.Fail()
	}
}

func Test_checkUpper(t *testing.T) {
	test1 := checkUpper("hey (up)")
	if !reflect.DeepEqual(test1, "HEY") {
		t.Logf("checkupper(\"hey (up)\") failed, wanted \"%v\", got \"%v\"", "HEY", test1)
		t.Fail()
	}

	test2 := checkUpper("Ready, set, go (up, 2)")
	if !reflect.DeepEqual(test2, "Ready, SET, GO") {
		t.Logf("checkupper(\"Ready, set, go (up, 2) !\") failed, wanted \"%v\", got \"%v\"", "Ready, SET, GO", test2)
		t.Fail()
	}

	test3 := checkUpper("42 (up)")
	if !reflect.DeepEqual(test3, "42 (up)") {
		t.Logf("checkupper(\"42 (up)\") failed, wanted \"%v\", got \"%v\"", "42 (up)", test3)
		t.Fail()
	}

	test4 := checkUpper("hey up")
	if !reflect.DeepEqual(test4, "hey up") {
		t.Logf("checkupper(\"hey up\") failed, wanted \"%v\", got \"%v\"", "hey up", test4)
		t.Fail()
	}

	test5 := checkUpper("hey there (up, 2)")
	if !reflect.DeepEqual(test5, "HEY THERE") {
		t.Logf("checkupper(\"hey there (up, 2)\") failed, wanted \"%v\", got \"%v\"", "HEY THERE", test5)
		t.Fail()
	}

	test6 := checkUpper("HEY (up)")
	if !reflect.DeepEqual(test6, "HEY") {
		t.Logf("checkupper(\"HEY (up)\") failed, wanted \"%v\", got \"%v\"", "HEY", test6)
		t.Fail()
	}

	test7In := "it was the best of times, it was the worst of times (up) ,"
	test7Want := "it was the best of times, it was the worst of TIMES ,"
	test7 := checkUpper(test7In)
	if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkupper(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}
}

func Test_checkLower(t *testing.T) {
	test1 := checkLower("HEY (low)")
	if !reflect.DeepEqual(test1, "hey") {
		t.Logf("checkLower(\"HEY (low)\") failed, wanted \"%v\", got \"%v\"", "hey", test1)
		t.Fail()
	}

	test2 := checkLower("Ready, set, GO (low) !")
	if !reflect.DeepEqual(test2, "Ready, set, go !") {
		t.Logf("checkLower(\"Ready, set, GO (low) !\") failed, wanted \"%v\", got \"%v\"", "Ready, set, go !", test2)
		t.Fail()
	}

	test3 := checkLower("42 (low)")
	if !reflect.DeepEqual(test3, "42 (low)") {
		t.Logf("checkLower(\"42 (low)\") failed, wanted \"%v\", got \"%v\"", "42 (low)", test3)
		t.Fail()
	}

	test4 := checkLower("hey low")
	if !reflect.DeepEqual(test4, "hey low") {
		t.Logf("checkLower(\"hey low\") failed, wanted \"%v\", got \"%v\"", "hey low", test4)
		t.Fail()
	}

	test5 := checkLower("HEY THERE (low, 2)")
	if !reflect.DeepEqual(test5, "hey there") {
		t.Logf("checkLower(\"HEY THERE (low, 2)\") failed, wanted \"%v\", got \"%v\"", "hey there", test5)
		t.Fail()
	}

	test6 := checkLower("hey (low)")
	if !reflect.DeepEqual(test6, "hey") {
		t.Logf("checkLower(\"hey (low)\") failed, wanted \"%v\", got \"%v\"", "hey", test6)
		t.Fail()
	}

	test7In := "it was the best of times, it was the worst of TIMES (low) ,"
	test7Want := "it was the best of times, it was the worst of times ,"
	test7 := checkLower(test7In)
	if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkLower(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}
}

func Test_checkCap(t *testing.T) {
	test1 := checkCap("hey (cap)")
	if !reflect.DeepEqual(test1, "Hey") {
		t.Logf("checkCap(\"hey (cap)\") failed, wanted \"%v\", got \"%v\"", "Hey", test1)
		t.Fail()
	}

	test2 := checkCap("Ready, set, go (cap) !")
	if !reflect.DeepEqual(test2, "Ready, set, Go !") {
		t.Logf("checkCap(\"Ready, set, go (cap) !\") failed, wanted \"%v\", got \"%v\"", "Ready, set, Go !", test2)
		t.Fail()
	}

	test3 := checkCap("42 (cap)")
	if !reflect.DeepEqual(test3, "42 (cap)") {
		t.Logf("checkCap(\"42 (cap)\") failed, wanted \"%v\", got \"%v\"", "42 (cap)", test3)
		t.Fail()
	}

	test4 := checkCap("hey cap")
	if !reflect.DeepEqual(test4, "hey cap") {
		t.Logf("checkCap(\"hey cap\") failed, wanted \"%v\", got \"%v\"", "hey cap", test4)
		t.Fail()
	}

	test5 := checkCap("hey there (cap, 2)")
	if !reflect.DeepEqual(test5, "Hey There") {
		t.Logf("checkCap(\"hey there (cap, 2)\") failed, wanted \"%v\", got \"%v\"", "Hey There", test5)
		t.Fail()
	}

	test6 := checkCap("Hey (cap)")
	if !reflect.DeepEqual(test6, "Hey") {
		t.Logf("checkCap(\"Hey (cap)\") failed, wanted \"%v\", got \"%v\"", "Hey", test6)
		t.Fail()
	}

	test7In := "it (cap) was the best of times,"
	test7Want := "It was the best of times,"
	test7 := checkCap(test7In)
	if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkCap(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}
}
