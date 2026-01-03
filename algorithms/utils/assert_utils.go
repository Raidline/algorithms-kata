package utils

import (
	"runtime/debug"
	"testing"
)

func AssertValue(t *testing.T, test string, want int, op func() int) {
	got := op()
	if want != got {
		t.Errorf("Testing %s, got %d, want %d", test, got, want)
		debug.PrintStack()
		t.FailNow()
	}
}
