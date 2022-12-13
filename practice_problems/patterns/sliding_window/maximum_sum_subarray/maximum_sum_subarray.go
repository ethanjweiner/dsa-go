package practice_problems

// Solution 1: Compute sum externally
// func MaximumSumSubarray(numbers []int, k int) int {
// 	maxSum := utils.Sum(numbers[:k])
// 	currentSum := maxSum

// 	for windowStart := 1; windowStart <= len(numbers)-k; windowStart++ {
// 		currentSum -= numbers[windowStart-1]
// 		currentSum += numbers[windowStart+k-1]

// 		if currentSum > maxSum {
// 			maxSum = currentSum
// 		}
// 	}

// 	return maxSum
// }

// Solution 2: Compute sum internally
func MaximumSumSubarray(numbers []int, k int) int {
	maxSum, currentSum, windowStart := 0, 0, 0

	for windowEnd := 0; windowEnd < len(numbers); windowEnd++ {
		currentSum += numbers[windowEnd]

		if windowEnd >= k-1 {
			if currentSum > maxSum {
				maxSum = currentSum
			}

			currentSum -= numbers[windowStart]
			windowStart++
		}
	}

	return maxSum
}
