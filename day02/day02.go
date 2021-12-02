package day02

import (
	"strconv"
	"strings"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	hor := 0
	depth := 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		command := splits[0]
		value, err := strconv.Atoi(splits[1])
		if err != nil {
			return 0, err
		}
		switch command {
		case "forward":
			hor += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return hor * depth, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	hor := 0
	depth := 0
	aim := 0
	for _, line := range lines {
		splits := strings.Split(line, " ")
		command := splits[0]
		value, err := strconv.Atoi(splits[1])
		if err != nil {
			return 0, err
		}
		switch command {
		case "forward":
			hor += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return hor * depth, nil
}
