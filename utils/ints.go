package utils

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a < 0 {
		// Correct enough for AoC, not expecting math.MinInt
		return -a
	}
	return a
}
