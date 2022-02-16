package day15

import (
	"container/heap"
	"errors"
	"strings"
)

func part1(input string) (int, error) {
	risks := readInput(input)

	return findLowestRiskPath(risks)
}

func part2(input string) (int, error) {
	risks := readInput(input)

	fullRisks := expand(risks, 5)
	return findLowestRiskPath(fullRisks)
}

func findLowestRiskPath(risks [][]int) (int, error) {
	size := len(risks)

	const undefined = 0
	cumulativeRisks := make([][]int, size)
	for i := range cumulativeRisks {
		cumulativeRisks[i] = make([]int, size)
	}

	queue := NewPriorityQueue()
	heap.Push(&queue, Position{0, 0, 0})

	destination := size - 1
	for queue.Len() > 0 {
		current := heap.Pop(&queue).(Position)
		if current.row == destination && current.column == destination {
			return current.risk, nil
		}

		for _, delta := range [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			row := current.row + delta[0]
			if row >= 0 && row < size {
				column := current.column + delta[1]
				if column >= 0 && column < size {
					nextRisk := cumulativeRisks[current.row][current.column] + risks[row][column]
					if cumulativeRisks[row][column] == undefined || nextRisk < cumulativeRisks[row][column] {
						cumulativeRisks[row][column] = nextRisk
						heap.Push(&queue, Position{nextRisk, row, column})
					}
				}
			}
		}
	}

	return 0, errors.New("no path found")
}

func expand(original [][]int, times int) [][]int {
	tile := len(original)
	expanded := make([][]int, tile*times)
	for i := range expanded {
		expanded[i] = make([]int, tile*times)
	}

	for countForRows := 0; countForRows < times; countForRows++ {
		offsetRows := countForRows * tile
		for i := 0; i < tile; i++ {
			for countForColumns := 0; countForColumns < times; countForColumns++ {
				offsetColumns := countForColumns * tile
				for j := 0; j < tile; j++ {
					newRisk := original[i][j] + countForRows + countForColumns
					if newRisk >= 10 {
						newRisk -= 9
					}
					expanded[offsetRows+i][offsetColumns+j] = newRisk
				}
			}
		}
	}
	return expanded
}

func readInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j, ch := range line {
			grid[i][j] = int(ch - '0')
		}
	}

	return grid
}
