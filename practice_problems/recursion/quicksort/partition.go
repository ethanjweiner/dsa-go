package practice_problems

import "golang.org/x/exp/constraints"

// Inserts the rightmost element into the proper array position
func Partition[T constraints.Ordered](slice []T, leftPointer int, rightPointer int) int {
	pivotIndex := rightPointer
	pivot := slice[pivotIndex]
	rightPointer--

	// Until the pointers cross
	for {
		// Move left pointer to right to value > pivot
		for slice[leftPointer] < pivot && leftPointer < len(slice) {
			leftPointer++
		}

		for slice[rightPointer] > pivot && rightPointer > 0 {
			rightPointer--
		}

		if leftPointer >= rightPointer {
			break
		}

		slice[leftPointer], slice[rightPointer] = slice[rightPointer], slice[leftPointer]
		leftPointer++
	}

	slice[pivotIndex] = slice[leftPointer]
	slice[leftPointer] = pivot

	return leftPointer
}
