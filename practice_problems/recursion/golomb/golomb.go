package practice_problems

func Golomb(n int, memo map[int]int) int {
	if n == 1 {
		return 1
	}

	if _, ok := memo[n]; !ok {
		memo[n] = 1 + Golomb(n-Golomb(Golomb(n-1, memo), memo), memo)
	}

	return memo[n]
}
