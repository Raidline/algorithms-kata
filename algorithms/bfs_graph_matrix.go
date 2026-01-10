package algorithms

import (
	"algos/kata/algorithms/model"
	"errors"
	"slices"
)

func BfsGraph(graph model.WeightedAdjacencyMatrix, source int, needle int) ([]int, error) {
	seen := make([]bool, len(graph.Graph))
	prev := make([]int, len(graph.Graph))

	fillIntArray(prev, -1)

	seen[source] = true
	q := NewQueue()
	q.Enqueue(source)

	for q.Length > 0 {

		curr := q.Deque()

		if curr == needle {
			break
		}

		seen[curr] = true

		for i, ch := range graph.Graph[curr] {
			if ch == 0 {
				continue
			}

			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = curr

			q.Enqueue(i)
		}
	}

	curr := needle
	out := make([]int, 0)

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	if len(out) == 0 {
		return []int{}, errors.New("Not found")
	}

	out = append(out, source)
	slices.Reverse(out)

	return out, nil
}

func fillIntArray(arr []int, v int) {
	for i := range arr {
		arr[i] = v
	}
}
