package practice_problems

func longestSubstring(str string, k int) int {
	windowStart, windowEnd, longestSubtringLength, lastLength := 0, 1, 0, 0
	chars := []rune(str)
	distintChars := map[rune]int{}

	for windowEnd <= len(str) {
		lastLength = windowEnd - windowStart - 1
		distintChars[chars[windowEnd-1]]++

		if len(distintChars) > k {
			// Reassing if new max
			if lastLength > longestSubtringLength {
				longestSubtringLength = lastLength
			}

			// Once max distinct chars reached, delete
			distintChars[chars[windowStart]]--
			if distintChars[chars[windowStart]] == 0 {
				delete(distintChars, chars[windowStart])
			}
			windowStart++
		}
		windowEnd++
	}

	lastLength = windowEnd - windowStart - 1
	if lastLength > longestSubtringLength {
		longestSubtringLength = lastLength
	}

	return longestSubtringLength
}
