package practice_problems

// func removeDuplicates(sortedNumbers []int) []int {
// 	if len(sortedNumbers) <= 1 {
// 		return sortedNumbers
// 	}

// 	p1, p2 := 0, 1

// 	// NB: `len` re-evaluates per loop
// 	for p2 < len(sortedNumbers) {
// 		if sortedNumbers[p1] == sortedNumbers[p2] {
// 			sortedNumbers, _ = utils.Delete(sortedNumbers, p2)
// 		} else {
// 			p1++
// 			p2++
// 		}
// 	}

// 	return sortedNumbers
// }

func removeDuplicates(sortedNumbers []int) []int {
	if len(sortedNumbers) <= 1 {
		return sortedNumbers
	}

	nextNonDuplicate := 1

	for idx := range sortedNumbers {
		// Only change the next insertion position once we find a non-duplicate
		if sortedNumbers[idx] != sortedNumbers[nextNonDuplicate-1] {
			sortedNumbers[nextNonDuplicate] = sortedNumbers[idx]
			nextNonDuplicate++
		}
	}

	return sortedNumbers[:nextNonDuplicate]
}
