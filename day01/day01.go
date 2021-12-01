package day01

import (
	"math"
	"strconv"
	"strings"
)

func part1(input string) (int, error) {
	splits := strings.Split(strings.TrimSpace(input), "\n")

	previous := math.MaxInt
	count := 0
	for _, s := range splits {
		value, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		if value > previous {
			count++
		}
		previous = value
	}

	return count, nil
}

func part2(input string) (int, error) {
	splits := strings.Split(strings.TrimSpace(input), "\n")

	const windowSize = 3
	windowSum := 0
	values := make([]int, len(splits))
	count := 0
	for i := 0; i < windowSize; i++ {
		value, err := strconv.Atoi(splits[i])
		if err != nil {
			return 0, err
		}
		values[i] = value
		windowSum += value
	}
	for i := windowSize; i < len(splits); i++ {
		value, err := strconv.Atoi(splits[i])
		if err != nil {
			return 0, err
		}
		values[i] = value

		previous := windowSum
		windowSum += value - values[i-windowSize]
		if windowSum > previous {
			count++
		}
	}

	return count, nil
}
