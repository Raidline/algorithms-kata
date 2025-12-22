package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}

	executeTest(t, 69, foo, true)
	executeTest(t, 1336, foo, false)
	executeTest(t, 69420, foo, true)
	executeTest(t, 69421, foo, false)
	executeTest(t, 1, foo, true)
	executeTest(t, 0, foo, false)
}

func executeTest(t *testing.T, needle int, haystack []int, want bool) {
	res := BinarySearch(haystack, needle)
	if res != want {
		printErr(t, needle, res, want)
	}
}

func printErr(t *testing.T, params ...any) {
	t.Errorf("Testing BinarySearch for : %s, got %d, want %d", params[0], params[1], params[2])
	t.FailNow()
}
