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

func AssertArrays(t *testing.T, test string, want []int, got []int) {
	if len(want) != len(got) {
		t.Errorf("Testing %s, want Len : %d, got Len: %d", test, len(want), len(got))
		t.FailNow()
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("Testing %s, want %d, got %d at position %d", test, v, got[i], i)
			t.FailNow()
		}
	}
}
