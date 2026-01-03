package algorithms

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{9, 420, 7, 4, 69, 3, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}

	QuickSort(arr)

	for i := range expected {
		if arr[i] != expected[i] {
			t.Errorf("Testing QuickSort, got %d, want %d at position %d", arr[i], expected[i], i)
			t.FailNow()
		}
	}
}
