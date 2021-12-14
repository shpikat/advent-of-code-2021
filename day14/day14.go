package day14

import (
	"math"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	polymer, rules := readInput(input)

	return polymerize(polymer, rules, 10), nil
}

func part2(input string) (int, error) {
	polymer, rules := readInput(input)

	return polymerize(polymer, rules, 40), nil
}

func polymerize(polymer string, rules map[string][2]string, steps int) int {
	pairs := map[string]int{}
	for i := 0; i < len(polymer)-1; i++ {
		pair := polymer[i : i+2]
		pairs[pair]++
	}

	for i := 0; i < steps; i++ {
		newPairs := make(map[string]int, len(pairs))
		for pair, count := range pairs {
			for _, p := range rules[pair] {
				newPairs[p] += count
			}
		}
		pairs = newPairs
	}

	elements := map[rune]int{}
	for pair, count := range pairs {
		for _, r := range pair {
			elements[r] += count
		}
	}

	// All the elements are counted twice except for the first and the last ones, bringing them on par now.
	elements[rune(polymer[0])]++
	elements[rune(polymer[len(polymer)-1])]++

	min := math.MaxInt
	max := 0
	for _, count := range elements {
		min = utils.Min(min, count)
		max = utils.Max(max, count)
	}

	return (max - min) / 2
}

func readInput(input string) (string, map[string][2]string) {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")

	polymer := blocks[0]
	rules := make(map[string][2]string, len(blocks[1]))
	for _, line := range strings.Split(blocks[1], "\n") {
		pairToSplit := line[0:2]
		elementToInsert := line[len(line)-1:]
		rules[pairToSplit] = [...]string{
			pairToSplit[:1] + elementToInsert,
			elementToInsert + pairToSplit[1:],
		}
	}

	return polymer, rules
}
