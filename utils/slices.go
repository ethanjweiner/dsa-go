package utils

import (
	"fmt"
)

func Sum(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}

func Delete[T any](slice []T, indexToDelete int) (newSlice []T, err error) {
	if indexToDelete >= len(slice) {
		err = fmt.Errorf("deletion index %d is out of the bounds of the slice", indexToDelete)
		return
	}

	newSlice = append(slice[:indexToDelete], slice[indexToDelete+1:]...)
	return
}

func ContiguousSubarraysFromStart[T any](slice []T, start int, end int) (subarrays [][]T) {
	for i := start + 1; i <= end; i++ {
		subarrays = append(subarrays, slice[start:i])
	}

	return
}

func ContiguousSubarraysFromEnd[T any](slice []T, start int, end int) (subarrays [][]T) {
	for i := end - 1; i >= start; i-- {
		subarrays = append(subarrays, slice[i:end])
	}

	return
}
