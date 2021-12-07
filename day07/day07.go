package day07

import (
	"sort"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	crabs, err := readInput(input)
	if err != nil {
		return 0, err
	}

	sort.Ints(crabs)
	median := crabs[len(crabs)/2]

	fuel := 0
	for _, n := range crabs {
		fuel += utils.Abs(n - median)
	}

	return fuel, nil
}

func part2(input string) (int, error) {
	crabs, err := readInput(input)
	if err != nil {
		return 0, err
	}

	sum := 0
	for _, c := range crabs {
		sum += c
	}
	meanFloored := int(float32(sum) / float32(len(crabs)))
	meanCeiled := meanFloored + 1

	fuelFloored := 0
	fuelCeiled := 0
	for _, c := range crabs {
		fuelFloored += calculateFuel(c, meanFloored)
		fuelCeiled += calculateFuel(c, meanCeiled)
	}

	return utils.Min(fuelFloored, fuelCeiled), nil
}

func calculateFuel(from, to int) int {
	n := utils.Abs(from - to)
	return (n + 1) * n / 2
}

func readInput(input string) ([]int, error) {
	splits := strings.Split(strings.TrimSpace(input), ",")
	numbers := make([]int, len(splits))
	for i, s := range splits {
		v, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		numbers[i] = v
	}
	return numbers, nil
}
