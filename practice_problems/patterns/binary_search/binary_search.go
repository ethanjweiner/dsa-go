package practice_problems

/*
# PROBLEM

- Input: Array of `sortedNumbers` and a `numberToFind`
- Output: The index of that number, or -1
- Edge Case: Start = end -> middle index = start - end / 2 = 0 / 2 = 0 (good)

# ALGORITHM 1: INDEX UPDATES

Given an array of `sortedNumbers` and `numberToFind`:
- Kick off the binary search with `sortedNumbers`, `start` and `end` set to 0 and length - 1

Given an array of `sortedNumbers`, a `numberToFind` a `start`, & `end``:
- If start > end -> return -1 (never reached)
- Find the `middleIndex`: floor(`end - start` / 2)
- Find the element at the `middleIndex`
- If element is `numberToFind`:
	- Return `middleIndex`:
- Else if element < numberToFind:
	- Binary search left
- Else:
	- Binary search right

Time Complexity:

# ALGORITHM 2: SLICING

*/

// Recursive

// Time Complexity: O(log N) (halving)
// Space Complexity: O(log N), assuming no tail recursion (call stack size max log N levels)

// func IndexOf(sortedNumbers []int, numberToFind int) int {
// 	return BinarySearch(sortedNumbers, numberToFind, 0, len(sortedNumbers)-1)
// }

func BinarySearch(sortedNumbers []int, numberToFind int, start int, end int) int {
	if start > end {
		return -1
	}

	middleIndex := (start + end) / 2
	middleElement := sortedNumbers[middleIndex]

	if numberToFind < middleElement {
		return BinarySearch(sortedNumbers, numberToFind, start, middleIndex-1)
	}

	if numberToFind > middleElement {
		return BinarySearch(sortedNumbers, numberToFind, middleIndex+1, end)
	}

	// Number has been found
	return middleIndex
}

// Iterative
// Time Complexity: O(log N) (halving)
// Space Complexity: O(1)

func IndexOf(sortedNumbers []int, numberToFind int) int {
	start, end := 0, len(sortedNumbers)-1

	for start <= end {
		middleIndex := (start + end) / 2
		middleElement := sortedNumbers[middleIndex]

		if numberToFind < middleElement {
			end = middleIndex - 1
		} else if numberToFind > middleElement {
			start = middleIndex + 1
		} else {
			return middleIndex
		}
	}

	return -1
}
