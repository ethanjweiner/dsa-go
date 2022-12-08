package practice_problems

import "golang.org/x/exp/constraints"

func Quicksort[T constraints.Ordered](slice []T) {
	if len(slice) <= 1 {
		return
	}

	pivotIndex := Partition(slice, 0, len(slice)-1)

	// Alternative option: Pass pointers to quicksort (mutations affect same slice)
	Quicksort(slice[:pivotIndex])
	Quicksort(slice[pivotIndex+1:])
}
