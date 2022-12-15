package utils

import "math"

func Min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
