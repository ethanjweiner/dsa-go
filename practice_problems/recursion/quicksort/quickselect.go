package practice_problems

import "golang.org/x/exp/constraints"

// Only accepts values < n
func FindNthLowest[T constraints.Ordered](slice []T, n int) T {
	sliceCopy := make([]T, len(slice))
	copy(sliceCopy, slice)
	return FindNthLowestHelper(sliceCopy, n-1, 0, len(slice)-1)
}

func FindNthLowestHelper[T constraints.Ordered](slice []T, indexToFind int, leftPointer int, rightPointer int) T {
	// Base case #1: Only 1 element left
	if rightPointer <= leftPointer {
		return slice[leftPointer]
	}

	pivotIndex := Partition(slice, leftPointer, rightPointer)

	switch {
	case pivotIndex > indexToFind:
		return FindNthLowestHelper(slice, indexToFind, leftPointer, pivotIndex-1)
	case pivotIndex < indexToFind:
		return FindNthLowestHelper(slice, indexToFind, pivotIndex+1, rightPointer)
	}

	// Base case #2: Newly retrieved pivot index is the index we are searching for
	return slice[pivotIndex]
}
