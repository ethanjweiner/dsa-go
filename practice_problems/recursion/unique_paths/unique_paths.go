package practice_problems

import "fmt"

// Overlapping subproblem: Some paths from each function call overlap
/*
|x|1| |
+-+-+-+
|2| | |
+-+-+-+
| | |y|
+-+-+-+

1 & 2 can both take the same path to y (start 1 downward)
Solution: Memoize the # of unique paths from for a given row/col combo
Potential solution: create unique string key
*/

func UniquePaths(rows int, columns int, memo map[string]int) int {
	if rows == 1 || columns == 1 {
		return 1
	}

	key := fmt.Sprintf("%d %d", rows, columns)

	if _, ok := memo[key]; !ok {
		memo[key] = UniquePaths(rows-1, columns, memo) + UniquePaths(rows, columns-1, memo)
	}

	return memo[key]
}
