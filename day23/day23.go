package day23

import (
	"strings"
)

func part1(input string) (int, error) {
	burrow := readInput(input, nil)

	return burrow.findLeastEnergyRequired()
}

func part2(input string) (int, error) {
	const folded = `
  #D#C#B#A#
  #D#B#A#C#
`
	burrow := readInput(input, strings.Split(strings.TrimSpace(folded), "\n"))

	return burrow.findLeastEnergyRequired()
}

func readInput(input string, folded []string) Burrow {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var rooms []string
	if folded == nil {
		rooms = lines[2 : len(lines)-1]
	} else {
		rooms = []string{
			lines[2],
			folded[0],
			folded[1],
			lines[3],
		}
	}

	burrow := NewBurrow(len(rooms))
	for place := range rooms {
		for i, s := range strings.Split(strings.Trim(strings.TrimSpace(rooms[place]), "#"), "#") {
			burrow.rooms[i].set(place, Amphipod(s[0]-'A'))
		}
	}

	return burrow
}
