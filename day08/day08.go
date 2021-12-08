package day08

import (
	"strings"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	count := 0
	for _, line := range lines {
		splits := strings.Split(line, " | ")
		values := strings.Fields(splits[1])
		for _, v := range values {
			switch len(v) {
			case 2, 4, 3, 7:
				count++
			}
		}
	}

	return count, nil
}

func part2(input string) (int, error) {
	type entry struct {
		patterns, output []string
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	entries := make([]entry, len(lines))
	for i, line := range lines {
		splits := strings.Split(line, " | ")
		entries[i] = entry{
			strings.Fields(splits[0]),
			strings.Fields(splits[1]),
		}
	}

	definiteSegmentsCountToNumber := map[int]int{2: 1, 3: 7, 4: 4, 7: 8}

	sum := 0
	for _, e := range entries {
		numbers := [10]digit{}
		signals := make(map[digit]int, 10)

		// Find 1, 4, 7, 8
		for _, p := range e.patterns {
			n, exists := definiteSegmentsCountToNumber[len(p)]
			if exists {
				d := getDigit(p)
				numbers[n] = d
				signals[d] = n
			}
		}

		// Find 0, 6, 9
		for _, p := range e.patterns {
			if len(p) == 6 {
				d := getDigit(p)
				var n int
				switch {
				case d&numbers[4] == numbers[4]:
					n = 9
				case d&numbers[7] == numbers[7]:
					n = 0
				default:
					n = 6
				}
				numbers[n] = d
				signals[d] = n
			}
		}

		// Find 2, 3, 5
		segmentE := numbers[8] ^ numbers[9]
		for _, p := range e.patterns {
			if len(p) == 5 {
				d := getDigit(p)
				var n int
				switch {
				case d&segmentE == segmentE:
					n = 2
				case d&numbers[1] == numbers[1]:
					n = 3
				default:
					n = 5
				}
				numbers[n] = d
				signals[d] = n
			}
		}

		output := 0
		for _, v := range e.output {
			output = output*10 + signals[getDigit(v)]
		}
		sum += output
	}

	return sum, nil
}

type digit uint8

func getDigit(input string) (bits digit) {
	for _, ch := range input {
		bits |= 1 << (ch - 'a')
	}
	return
}
