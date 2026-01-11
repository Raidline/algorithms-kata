package algorithms

import (
	"algos/kata/algorithms/utils"
	"testing"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[string, int](3)

	utils.AssertValueWithError(t, "LRU", 0, func() (int, error) {
		return lru.Get("foo")
	})

	lru.Update("foo", 69)

	utils.AssertValueWithError(t, "LRU", 69, func() (int, error) {
		return lru.Get("foo")
	})

	lru.Update("bar", 420)
	utils.AssertValueWithError(t, "LRU", 420, func() (int, error) {
		return lru.Get("bar")
	})

	lru.Update("baz", 1337)
	utils.AssertValueWithError(t, "LRU", 1337, func() (int, error) {
		return lru.Get("baz")
	})

	lru.Update("ball", 69420)

	utils.AssertValueWithError(t, "LRU", 69420, func() (int, error) {
		return lru.Get("ball")
	})

	utils.AssertValueWithError(t, "LRU", 0, func() (int, error) {
		return lru.Get("foo")
	})

	utils.AssertValueWithError(t, "LRU", 420, func() (int, error) {
		return lru.Get("bar")
	})
	lru.Update("foo", 69)

	utils.AssertValueWithError(t, "LRU", 420, func() (int, error) {
		return lru.Get("bar")
	})

	utils.AssertValueWithError(t, "LRU", 69, func() (int, error) {
		return lru.Get("foo")
	})

	// shouldn't of been deleted, but since bar was get'd, bar was added to the
	// front of the list, so baz became the end
	utils.AssertValueWithError(t, "LRU", 0, func() (int, error) {
		return lru.Get("baz")
	})
}
