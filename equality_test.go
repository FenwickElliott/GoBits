package test

import "testing"

var a = 0
var b = 0

func TestEquality(t *testing.T) {
	if a != b {
		t.Error("not equal")
	}
}
