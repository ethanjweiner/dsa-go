package practice_problems

import (
	"github.com/ethanjweiner/dsa-go/utils"
)

// TODO: Use sliding window w/ a left & right, continually dividing the product
// TODO: Goal is to avoid unnecessarily computing things twice
// TODO: Write down this technique. This idea of minimizing computation and asking the question:
// TODO: How can I avoid duplicatee processing? What step can I take (here it's division)

// Traditional Approach

// func subarrayProductLessThanK(numbers []int, target int) (subarrays [][]int) {
// 	for i := range numbers {
// 		product := 1
// 		subarray := []int{}

// 		for j := i; j < len(numbers); j++ {
// 			product *= numbers[j]
// 			if product >= target {
// 				break
// 			}
// 			subarray = append(subarray, numbers[j])
// 			newSubarray := []int{}
// 			copy(newSubarray, subarray)
// 			subarrays = append(subarrays, subarray)
// 		}
// 	}

// 	return subarrays
// }

func subarrayProductLessThanK(numbers []int, target int) (subarrays [][]int) {
	product := 1
	left := 0

	for right := 0; right < len(numbers); right++ {
		product *= numbers[right]

		for product >= target && left < len(numbers) {
			product /= numbers[left]
			left++
		}

		subarrays = append(subarrays, utils.ContiguousSubarraysFromEnd(numbers, left, right+1)...)
	}

	return
}
