package day01

import (
	"strconv"
	"strings"
)

func solvePart1(input string) (int, error) {
	splits := strings.Split(strings.TrimSpace(input), "\n")

	previous := 0
	count := 0
	for i, s := range splits {
		value, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		if i != 0 && value > previous {
			count++
		}
		previous = value
	}

	return count, nil
}

func solvePart2(input string) (int, error) {
	splits := strings.Split(strings.TrimSpace(input), "\n")

	const windowSize = 3
	windows := make([]int, len(splits))
	count := 0
	for i, s := range splits {
		value, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}

		window := i - windowSize
		for j := i; j >= 0 && j > window; j-- {
			windows[j] += value
		}

		if window >= 0 && windows[window+1] > windows[window] {
			count++
		}
	}

	return count, nil
}
