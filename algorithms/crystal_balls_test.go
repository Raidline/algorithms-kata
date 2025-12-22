package algorithms

import (
	"math/rand"
	"testing"
)

func TestCrystalBalls(t *testing.T) {
	idx := rand.Intn(10000)
	data := fillArray(10000, false)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	executeCrystalTest(t, data, idx)
	executeCrystalTest(t, fillArray(821, false), -1)
}

func executeCrystalTest(t *testing.T, haystack []bool, want int) {
	res := TwoCristalBalls(haystack)
	if res != want {
		printTwoCristalBallsErr(t, res, want)
	}
}

func printTwoCristalBallsErr(t *testing.T, params ...any) {
	t.Errorf("Testing TwoCristalBalls, got %d, want %d", params[0], params[1])
	t.FailNow()
}

func fillArray(len int, value bool) []bool {
	values := make([]bool, len)
	for i := range len {
		values[i] = value
	}

	return values
}
