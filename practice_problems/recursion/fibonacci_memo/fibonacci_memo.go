package practice_problems

import "fmt"

func Fib(n int) int {
	return fibMemo(n, make(map[int]int))
}

func fibMemo(n int, memo map[int]int) int {
	fmt.Println("Call!")
	if n == 0 || n == 1 {
		return n
	}

	if _, ok := memo[n]; !ok {
		memo[n] = fibMemo(n-2, memo) + fibMemo(n-1, memo)
	}

	return memo[n]
}
