package utils

import (
	"algos/kata/algorithms/model"
	"testing"
)

func TestList(list model.List, t *testing.T) {
	list.Append(5)
	list.Append(7)
	list.Append(9)

	if v := list.Get(2); v != 9 {
		printErr(t, "Get", v, 9)
	}

	if v := list.RemoveAt(1); v != 7 {
		printErr(t, "RemoveAt", v, 7)
	}

	if v := list.Length(); v != 2 {
		printErr(t, "Length", v, 2)
	}

	list.Append(11)

	if v := list.RemoveAt(1); v != 9 {
		printErr(t, "RemoveAt", v, 9)
	}

	if v := list.Remove(1); v != -1 {
		printErr(t, "Remove", v, -1)
	}

	if v := list.RemoveAt(0); v != 5 {
		printErr(t, "RemoveAt", v, 5)
	}

	if v := list.RemoveAt(0); v != 11 {
		printErr(t, "RemoveAt", v, 11)
	}

	if v := list.Length(); v != 0 {
		printErr(t, "Length", v, 0)
	}

	list.Prepend(5)
	list.Prepend(7)
	list.Prepend(9)

	if v := list.Get(2); v != 5 {
		printErr(t, "Get", v, 5)
	}

	if v := list.Get(0); v != 9 {
		printErr(t, "Get", v, 9)
	}

	if v := list.Remove(9); v != 9 {
		printErr(t, "Remove", v, 9)
	}

	if v := list.Length(); v != 2 {
		printErr(t, "Length", v, 2)
	}

	if v := list.Get(0); v != 7 {
		printErr(t, "Get", v, 7)
	}
}

func printErr(t *testing.T, params ...any) {
	t.Errorf("Testing List %s, got %d, want %d", params[0], params[1], params[2])
	t.FailNow()
}
