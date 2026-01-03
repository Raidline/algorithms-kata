package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestStack(t *testing.T) {
	list := NewStack()
	test := "Stack"

	list.Push(5)
	list.Push(7)
	list.Push(9)

	utils.AssertValue(t, test, 9, func() int { return list.Pop() })
	utils.AssertValue(t, test, 2, func() int { return list.Length })

	list.Push(11)
	utils.AssertValue(t, test, 11, func() int { return list.Pop() })
	utils.AssertValue(t, test, 7, func() int { return list.Pop() })
	utils.AssertValue(t, test, 5, func() int { return list.Peek() })
	utils.AssertValue(t, test, 5, func() int { return list.Pop() })
	utils.AssertValue(t, test, -1, func() int { return list.Pop() })

	list.Push(69)
	utils.AssertValue(t, test, 69, func() int { return list.Peek() })
	utils.AssertValue(t, test, 1, func() int { return list.Length })
}
