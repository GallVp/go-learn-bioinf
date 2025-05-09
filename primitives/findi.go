package primitives

import (
	"golang.org/x/exp/slices"
)

func findIndicesOfMin(numbers []int) []int {

	var indices []int

	minVal := slices.Min(numbers)
	for i, v := range numbers {
		if v == minVal {
			indices = append(indices, i)
		}
	}
	return indices
}
