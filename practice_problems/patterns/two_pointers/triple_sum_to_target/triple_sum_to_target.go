package practice_problems

import (
	"math"
	"sort"
)

// Solution 1: Search from left -> right
// func searchTriplet(numbers []int, target int) int {
// 	// Guard clause if no triplets exist
// 	if len(numbers) < 3 {
// 		return 0
// 	}

// 	sort.Ints(numbers)
// 	bestSum := numbers[0] + numbers[1] + numbers[2]

// 	for idx := 0; idx <= len(numbers)-3; idx++ {
// 		x := numbers[idx]
// 		bestRemainingSum := searchPair(numbers[idx+1:], target-x)
// 		currentSum := x + bestRemainingSum
// 		if abs(currentSum-target) < abs(bestSum-target) {
// 			bestSum = currentSum
// 		} else if currentSum-target >= abs(bestSum-target) {
// 			break
// 		}
// 	}

// 	return bestSum
// }

func abs(num int) int {
	if num < 0 {
		return -1 * num
	}

	return num
}

// func searchPair(numbers []int, target int) int {
// 	bestSum := numbers[0] + numbers[1]

// 	for idx := 0; idx <= len(numbers)-2; idx++ {
// 		y := numbers[idx]
// 		z := closestNumberToTarget(numbers[idx+1:], target-y)
// 		currentSum := y + z
// 		if abs(currentSum-target) < abs(bestSum-target) {
// 			bestSum = currentSum
// 		} else if currentSum-target >= abs(bestSum-target) {
// 			break
// 		}
// 	}

// 	return bestSum
// }

func closestNumberToTarget(numbers []int, target int) int {
	bestNum := numbers[0]

	for idx := 0; idx <= len(numbers)-1; idx++ {
		z := numbers[idx]
		if abs(z-target) < abs(bestNum-target) {
			bestNum = z
		} else if z-target >= abs(bestNum-target) {
			break
		}
	}

	return bestNum
}

// Solution 2: Search in a more efficient way
func searchTriplet(numbers []int, target int) int {
	if len(numbers) < 3 {
		return 0
	}

	sort.Ints(numbers)

	smallestDiff := math.MaxInt

	for idx := 0; idx <= len(numbers)-3; idx++ {
		x := numbers[idx]
		currentDiff := target - x - searchPair(numbers, target-x, idx+1)

		if currentDiff == 0 {
			return target
		}

		if (abs(currentDiff) < abs(smallestDiff)) || (abs(currentDiff) == abs(smallestDiff) && currentDiff > smallestDiff) {
			smallestDiff = currentDiff
		}
	}

	return target - smallestDiff
}

func searchPair(numbers []int, target int, left int) int {
	right := len(numbers) - 1

	smallestDiff := math.MaxInt

	for left < right {
		currentDiff := target - (numbers[left] + numbers[right])

		if currentDiff == 0 {
			return target
		}

		// Keep adjusting `currentDiff` until it is no longer becoming smaller
		if (abs(currentDiff) < abs(smallestDiff)) || (abs(currentDiff) == abs(smallestDiff) && currentDiff > smallestDiff) {
			smallestDiff = currentDiff
		}

		if currentDiff > 0 {
			left++
		} else {
			right--
		}
	}

	return target - smallestDiff
}
