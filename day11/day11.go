package day11

import (
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

const levelsCount = 10

var neighbourhoodDeltas = createNeighbourhoodDeltas()

func part1(input string) (int, error) {
	const steps = 100

	grid := readInput(input)

	totalFlashes := 0
	for s := 0; s < steps; s++ {
		totalFlashes += grid.Step()
	}

	return totalFlashes, nil
}

func part2(input string) (int, error) {
	const totalOctopuses = 100

	grid := readInput(input)

	step := 0
	for {
		// Assume there's obviously no need to calculate the first step if all
		// octopuses are lit at the moment of entering the cavern (i.e. step 0)
		step++
		if grid.Step() == totalOctopuses {
			return step, nil
		}
	}
}

type Grid struct {
	levels         [10]utils.IntSet
	zeroLevelIndex int
}

func NewGrid() (grid Grid) {
	for i := range grid.levels {
		grid.levels[i] = make(utils.IntSet)
	}
	return
}

func (g *Grid) Set(row, column, level int) {
	g.levels[level].Add(pack(row, column))
}

func (g *Grid) Step() int {
	if g.zeroLevelIndex == 0 {
		g.zeroLevelIndex = levelsCount - 1
	} else {
		g.zeroLevelIndex--
	}

	flashed := make([]int, 0, len(g.levels[g.zeroLevelIndex]))
	for octopus := range g.levels[g.zeroLevelIndex] {
		flashed = append(flashed, octopus)
	}
	for len(flashed) != 0 {
		energyGained := make(map[int]int, len(flashed)*8)
		for _, octopus := range flashed {
			// can be improved with signed packed integer arithmetics
			row, column := unpack(octopus)
			for _, delta := range neighbourhoodDeltas {
				neighbour := pack(row+delta[0], column+delta[1])
				energyGained[neighbour] += 1
			}
		}

		movingToLevel := [levelsCount + 1][]int{}
		for level := 1; level <= 9; level++ {
			index := (g.zeroLevelIndex + level) % levelsCount
			leavingThisLevel := make([]int, 0, len(g.levels[index]))
			for octopus := range g.levels[index] {
				energy, hasGained := energyGained[octopus]
				if hasGained {
					leavingThisLevel = append(leavingThisLevel, octopus)
					levelToMoveTo := utils.Min(level+energy, levelsCount)
					movingToLevel[levelToMoveTo] = append(movingToLevel[levelToMoveTo], octopus)
				}
			}
			for _, octopus := range leavingThisLevel {
				g.levels[index].Remove(octopus)
			}
			for _, octopus := range movingToLevel[level] {
				g.levels[index].Add(octopus)
			}
		}
		flashed = movingToLevel[levelsCount]
		for _, octopus := range flashed {
			g.levels[g.zeroLevelIndex].Add(octopus)
		}
	}
	return len(g.levels[g.zeroLevelIndex])
}

func pack(row, column int) int {
	return row<<8 | column
}

func unpack(packed int) (int, int) {
	return packed >> 8, packed & 0x0FF
}

func createNeighbourhoodDeltas() (neighbourhood [8][2]int) {
	index := 0
	for row := -1; row <= 1; row++ {
		for column := -1; column <= 1; column++ {
			if !(row == 0 && column == 0) {
				neighbourhood[index] = [2]int{row, column}
				index++
			}
		}
	}
	return
}

func readInput(input string) Grid {
	grid := NewGrid()
	for i, line := range strings.Split(strings.TrimSpace(input), "\n") {
		for j, ch := range line {
			level := int(ch - '0')
			grid.Set(i, j, level)
		}
	}
	return grid
}
