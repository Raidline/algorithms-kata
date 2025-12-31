package algorithms

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	expected := []int{3, 4, 7, 9, 42, 69, 420}

	BubbleSort(arr)

	for i := range arr {
		if arr[i] != expected[i] {
			t.Errorf("Testing BubbleSort, got %d, want %d at position %d", arr[i], expected[i], i)
			t.FailNow()
		}
	}
}
