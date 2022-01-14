package main

import (
	"reflect"
	"testing"
)

func Test_checkhex(t *testing.T) {
	test1 := checkhex("42 (hex)")
	if !reflect.DeepEqual(test1, "66") {
		t.Logf("checkhex(\"42 (hex)\") failed, wanted \"%v\", got \"%v\"", "66", test1)
		t.Fail()
	}
}
