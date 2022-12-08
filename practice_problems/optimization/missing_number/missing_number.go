package practice_problems

import "sort"

func MissingNumber(integers []int) int {
	integersCopy := make([]int, len(integers))
	copy(integersCopy, integers) // O(n)
	sort.Ints(integers)          // O(n log n)

	// Return first index which
	for idx, num := range integers {
		if idx != num {
			return idx
		}
	}

	return -1 // No missing number found
}
