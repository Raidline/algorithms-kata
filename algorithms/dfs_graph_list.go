package algorithms

import (
	"algos/kata/algorithms/model"
	"errors"
)

func walkGraphList(graph model.WeightedAdjacencyList, curr int, needle int,
	seen []bool, path *[]int) bool {

	if seen[curr] {
		return false
	}

	//pre
	seen[curr] = true

	*path = append(*path, curr)
	if curr == needle {
		return true
	}
	//recurse
	for _, edge := range graph.Graph[curr] {

		if walkGraphList(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	//post
	*path = (*path)[:len(*path)-1]
	return false
}

func DfsGraphList(graph model.WeightedAdjacencyList, source int, needle int) ([]int, error) {
	seen := make([]bool, len(graph.Graph))
	path := make([]int, 0)

	walkGraphList(graph, source, needle, seen, &path)

	if len(path) != 0 {
		return path, nil
	}

	return []int{}, errors.New("Not found")
}
