package algorithms

import (
	"algos/kata/algorithms/model"
	"math"
	"slices"
)

// this is a straighforward impl of the algoritm just to get the real grasp of it
// the runtime can be improved if we used a minHeap for this.
// why? the minHeap will always have the min value at the top and by adding to the heap the seen nodes we do not have to keep the seen array
// seeing if the node has been visited and getting the miniumn is O(V^2)
// by using the minheap we can get that to O(log(V)) -> V being the edges
func DijkstraList(source int, sink int, arr model.WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr.Graph))
	dists := make([]int, len(arr.Graph))
	prev := make([]int, len(arr.Graph))

	fillIntArray(prev, -1)
	fillIntArray(dists, math.MaxInt)

	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)
		seen[curr] = true

		for _, edge := range arr.Graph[curr] {
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight

			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}
		}
	}

	out := make([]int, 0)
	curr := sink

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	slices.Reverse(out)
	return out
}

func hasUnvisited(seen []bool, dists []int) bool {
	for i, v := range seen {
		if !v && dists[i] < math.MaxInt {
			return true
		}
	}

	return false
}

// returns the index of the lowestVisited item
func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDist := math.MaxInt

	for i := range seen {
		if seen[i] {
			continue
		}

		if lowestDist > dists[i] {
			lowestDist = dists[i]
			idx = i
		}
	}

	return idx
}
