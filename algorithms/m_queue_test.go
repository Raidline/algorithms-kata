package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestQueue(t *testing.T) {
	test := "Queue"
	list := NewQueue()

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	utils.AssertValue(t, test, 5, func() int { return list.Deque() })
	utils.AssertValue(t, test, 2, func() int { return list.Length })

	list.Enqueue(11)

	utils.AssertValue(t, test, 7, func() int { return list.Deque() })
	utils.AssertValue(t, test, 9, func() int { return list.Deque() })
	utils.AssertValue(t, test, 11, func() int { return list.Peek() })
	utils.AssertValue(t, test, 11, func() int { return list.Deque() })
	utils.AssertValue(t, test, -1, func() int { return list.Deque() }) //should be nil but not using pointers as values
	utils.AssertValue(t, test, 0, func() int { return list.Length })

	list.Enqueue(69)
	utils.AssertValue(t, test, 69, func() int { return list.Peek() })
	utils.AssertValue(t, test, 1, func() int { return list.Length })
}
