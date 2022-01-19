package main

import (
	"reflect"
	"testing"
)

func Test_checkhex(t *testing.T) {
	test1, err := checkhex("42 (hex)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, "66") {
		t.Logf("checkhex(\"42 (hex)\") failed, wanted \"%v\", got \"%v\"", "66", test1)
		t.Fail()
	}

	test2, err := checkhex("42 (hox)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, "42 (hox)") {
		t.Logf("checkhex(\"42 (hox)\") failed, wanted \"%v\", got \"%v\"", "42 (hox)", test2)
		t.Fail()
	}

	test3, err := checkhex("42 ()")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test3, "42 ()") {
		t.Logf("checkhex(\"42 ()\") failed, wanted \"%v\", got \"%v\"", "42 ()", test3)
		t.Fail()
	}

	test4, err := checkhex("42 hex")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test4, "42 hex") {
		t.Logf("checkhex(\"42 hex\") failed, wanted \"%v\", got \"%v\"", "42 (hex)", test4)
		t.Fail()
	}

	test5, err := checkhex("xx (hex)")
	if err == nil {
		t.Logf("xx (hex) failed, wanted err, got test5 == %v", test5)
		t.Fail()
	}

	test6, err := checkhex("1e (hex)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test6, "30") {
		t.Logf("checkhex(\"1e (hex)\") failed, wanted \"%v\", got \"%v\"", "30", test6)
		t.Fail()
	}

	test7In := "Simply add 42 (hex) and 10 (up) and you will see the result is 68"
	test7Want := "Simply add 66 and 10 (up) and you will see the result is 68"
	test7, err := checkhex(test7In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkhex(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}

	test8In := "Simply add 42 (hex) and 10 (up) and you 1e (hex) will see the result is 68"
	test8Want := "Simply add 66 and 10 (up) and you 30 will see the result is 68"
	test8, err := checkhex(test8In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test8, test8Want) {
		t.Logf("checkhex(\"%s\") failed, wanted \"%s\", got \"%s\"", test8In, test8Want, test8)
		t.Fail()
	}
}

func Test_checkbin(t *testing.T) {
	test1, err := checkbin("10 (bin)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test1, "2") {
		t.Logf("checkbin(\"10 (bin)\") failed, wanted \"%v\", got \"%v\"", "42", test1)
		t.Fail()
	}

	test2, err := checkbin("42 (hox)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test2, "42 (hox)") {
		t.Logf("checkbin(\"42 (hox)\") failed, wanted \"%v\", got \"%v\"", "42 (hox)", test2)
		t.Fail()
	}

	test3, err := checkbin("42 ()")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test3, "42 ()") {
		t.Logf("checkbin(\"42 ()\") failed, wanted \"%v\", got \"%v\"", "42 ()", test3)
		t.Fail()
	}

	test4, err := checkbin("42 bin")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test4, "42 bin") {
		t.Logf("checkbin(\"42 bin\") failed, wanted \"%v\", got \"%v\"", "42 (bin)", test4)
		t.Fail()
	}

	test5, err := checkbin("xx (bin)")
	if err == nil {
		t.Logf("xx (bin) failed, wanted err, got test5 == %v", test5)
		t.Fail()
	}

	test6, err := checkbin("110 (bin)")
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test6, "6") {
		t.Logf("checkbin(\"110 (bin)\") failed, wanted \"%v\", got \"%v\"", "6", test6)
		t.Fail()
	}

	test7In := "Simply add 10 (bin) and 66 you will see the result is 68"
	test7Want := "Simply add 2 and 66 you will see the result is 68"
	test7, err := checkbin(test7In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkbin(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}

	test8In := "Simply add 100110 (bin) and 11110 (bin) and you will see the result is 68"
	test8Want := "Simply add 38 and 30 and you will see the result is 68"
	test8, err := checkbin(test8In)
	if err != nil {
		t.Log(err.Error())
		t.Fail()
	} else if !reflect.DeepEqual(test8, test8Want) {
		t.Logf("checkbin(\"%s\") failed, wanted \"%s\", got \"%s\"", test8In, test8Want, test8)
		t.Fail()
	}
}

func Test_checkupper(t *testing.T) {
	test1 := checkupper("hey (up)")
	if !reflect.DeepEqual(test1, "HEY") {
		t.Logf("checkupper(\"hey (up)\") failed, wanted \"%v\", got \"%v\"", "HEY", test1)
		t.Fail()
	}

	// test2 := checkupper("Ready, set, go (up, 2) !")
	// if !reflect.DeepEqual(test2, "Ready, SET, GO !") {
	// 	t.Logf("checkupper(\"Ready, set, go (up, 2) !\") failed, wanted \"%v\", got \"%v\"", "Ready, SET, GO !", test2)
	// 	t.Fail()
	// }

	test3 := checkupper("42 (up)")
	if !reflect.DeepEqual(test3, "42 (up)") {
		t.Logf("checkupper(\"42 (up)\") failed, wanted \"%v\", got \"%v\"", "42 (up)", test3)
		t.Fail()
	}

	test4 := checkupper("hey up")
	if !reflect.DeepEqual(test4, "hey up") {
		t.Logf("checkupper(\"hey up\") failed, wanted \"%v\", got \"%v\"", "hey up", test4)
		t.Fail()
	}

	test6 := checkupper("HEY (up)")
	if !reflect.DeepEqual(test6, "HEY (up)") {
		t.Logf("checkupper(\"HEY (up)\") failed, wanted \"%v\", got \"%v\"", "HEY (up)", test6)
		t.Fail()
	}

	test7In := "it was the best of times, it was the worst of times (up) ,"
	test7Want := "it was the best of times, it was the worst of TIMES ,"
	test7 := checkupper(test7In)
	if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checkupper(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}
}

func Test_checklower(t *testing.T) {
	test1 := checklower("HEY (low)")
	if !reflect.DeepEqual(test1, "hey") {
		t.Logf("checklower(\"HEY (low)\") failed, wanted \"%v\", got \"%v\"", "hey", test1)
		t.Fail()
	}

	test2 := checklower("Ready, set, GO (low) !")
	if !reflect.DeepEqual(test2, "Ready, set, go !") {
		t.Logf("checklower(\"Ready, set, GO (low) !\") failed, wanted \"%v\", got \"%v\"", "Ready, set, go !", test2)
		t.Fail()
	}

	test3 := checklower("42 (low)")
	if !reflect.DeepEqual(test3, "42 (low)") {
		t.Logf("checklower(\"42 (low)\") failed, wanted \"%v\", got \"%v\"", "42 (low)", test3)
		t.Fail()
	}

	test4 := checklower("hey low")
	if !reflect.DeepEqual(test4, "hey low") {
		t.Logf("checklower(\"hey low\") failed, wanted \"%v\", got \"%v\"", "hey low", test4)
		t.Fail()
	}

	test6 := checklower("hey (low)")
	if !reflect.DeepEqual(test6, "hey (low)") {
		t.Logf("checklower(\"hey (low)\") failed, wanted \"%v\", got \"%v\"", "hey (low)", test6)
		t.Fail()
	}

	test7In := "it was the best of times, it was the worst of TIMES (low) ,"
	test7Want := "it was the best of times, it was the worst of times ,"
	test7 := checklower(test7In)
	if !reflect.DeepEqual(test7, test7Want) {
		t.Logf("checklower(\"%s\") failed, wanted \"%s\", got \"%s\"", test7In, test7Want, test7)
		t.Fail()
	}
}

func Test_checkN(t *testing.T) {
	test1 := checkN("HEY THERE (low, 2)")
	if !reflect.DeepEqual(test1, "hey there") {
		t.Logf("checkN(\"HEY THERE (low, 2)\") failed, wanted \"%v\", got \"%v\"", "hey there", test1)
		t.Fail()
	}
}
