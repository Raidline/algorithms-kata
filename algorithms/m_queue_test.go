package algorithms

import (
	"runtime/debug"
	"testing"
)

func TestQueue(t *testing.T) {
	list := NewQueue()

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	assertValue(t, 5, func() int { return list.Deque() })
	assertValue(t, 2, func() int { return list.Length })

	list.Enqueue(11)

	assertValue(t, 7, func() int { return list.Deque() })
	assertValue(t, 9, func() int { return list.Deque() })
	assertValue(t, 11, func() int { return list.Peek() })
	assertValue(t, 11, func() int { return list.Deque() })
	assertValue(t, -1, func() int { return list.Deque() }) //should be nil but not using pointers as values
	assertValue(t, 0, func() int { return list.Length })

	list.Enqueue(69)
	assertValue(t, 69, func() int { return list.Peek() })
	assertValue(t, 1, func() int { return list.Length })
}

func assertValue(t *testing.T, want int, op func() int) {
	got := op()
	if want != got {
		t.Errorf("Testing Queue, got %d, want %d", got, want)
		debug.PrintStack()
		t.FailNow()
	}
}
