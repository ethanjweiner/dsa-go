package practice_problems

func LongestConsecutiveSequence(integers []int) int {
	set := map[int]bool{}

	for _, integer := range integers {
		set[integer] = true
	}

	maxSequenceLength := 0

	for _, integer := range integers {
		if _, present := set[integer-1]; present {
			continue
		}

		currentSequenceLength := 1
		sequenceCandidate := integer + 1

		for {
			if _, present := set[sequenceCandidate]; !present {
				break
			}
			sequenceCandidate++
			currentSequenceLength++
		}

		if currentSequenceLength > maxSequenceLength {
			maxSequenceLength = currentSequenceLength
		}
	}

	return maxSequenceLength
}

// func LongestConsecutiveSequence(integers []int) int {
// 	min := math.MaxInt
// 	max := math.MinInt
// 	set := map[int]bool{}

// 	for _, integer := range integers {
// 		if integer < min {
// 			min = integer
// 		}

// 		if integer > max {
// 			max = integer
// 		}

// 		set[integer] = true
// 	}

// 	maxSequenceLength := 0
// 	currentSequenceLength := 0

// 	for num := min; num <= max; num++ {
// 		if _, ok := set[num]; ok {
// 			currentSequenceLength++
// 			continue
// 		}

// 		if currentSequenceLength > maxSequenceLength {
// 			maxSequenceLength = currentSequenceLength
// 		}

// 		currentSequenceLength = 0
// 	}

// 	return maxSequenceLength
// }
