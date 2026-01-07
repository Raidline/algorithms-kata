package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestMinHeap(t *testing.T) {
	heap := NewHeap()

	utils.AssertValue(t, "MinHeap", 0, func() int { return heap.Length() })

	heap.Insert(5)
	heap.Insert(3)
	heap.Insert(69)
	heap.Insert(420)
	heap.Insert(4)
	heap.Insert(1)
	heap.Insert(8)
	heap.Insert(7)

	utils.AssertValue(t, "MinHeap", 8, func() int { return heap.Length() })
	utils.AssertValue(t, "MinHeap", 1, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 3, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 4, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 5, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 4, func() int { return heap.Length() })
	utils.AssertValue(t, "MinHeap", 7, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 8, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 69, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 420, func() int { return heap.Delete() })
	utils.AssertValue(t, "MinHeap", 0, func() int { return heap.Length() })
}
