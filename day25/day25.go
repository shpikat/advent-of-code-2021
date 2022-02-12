package day25

import (
	"strings"
)

const (
	EastFacing  = '>'
	SouthFacing = 'v'
	Empty       = '.'
	JustMoved   = 'x'
)

func part1(input string) (int, error) {
	seafloor := readInput(input)

	rightEdgeIndex := len(seafloor[0]) - 1
	bottomEdgeIndex := len(seafloor) - 1
	topEdgeSnapshot := make([]byte, len(seafloor[0]))

	step := 0
	hasChanges := true

	for hasChanges {
		step++
		hasChanges = false
		for i := range seafloor {
			mustMoveFromRightToLeft := seafloor[i][rightEdgeIndex] == EastFacing && seafloor[i][0] == Empty
			for j := 0; j < rightEdgeIndex; j++ {
				if seafloor[i][j] == EastFacing && seafloor[i][j+1] == Empty {
					seafloor[i][j], seafloor[i][j+1] = Empty, EastFacing
					j++
					hasChanges = true
				}
			}
			if mustMoveFromRightToLeft {
				seafloor[i][rightEdgeIndex], seafloor[i][0] = Empty, EastFacing
				hasChanges = true
			}
		}

		copy(topEdgeSnapshot, seafloor[0])
		for i := 0; i < bottomEdgeIndex; i++ {
			for j := range seafloor[i] {
				if seafloor[i][j] == JustMoved {
					seafloor[i][j] = SouthFacing
				} else if seafloor[i][j] == SouthFacing && seafloor[i+1][j] == Empty {
					seafloor[i][j], seafloor[i+1][j] = Empty, JustMoved
					hasChanges = true
				}
			}
		}
		for j := range seafloor[bottomEdgeIndex] {
			if seafloor[bottomEdgeIndex][j] == JustMoved {
				seafloor[bottomEdgeIndex][j] = SouthFacing
			} else if seafloor[bottomEdgeIndex][j] == SouthFacing && topEdgeSnapshot[j] == Empty {
				seafloor[bottomEdgeIndex][j], seafloor[0][j] = Empty, SouthFacing
				hasChanges = true
			}
		}
	}

	return step, nil
}

func readInput(input string) [][]byte {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	seafloor := make([][]byte, len(lines))
	for i := range lines {
		seafloor[i] = []byte(lines[i])
	}
	return seafloor
}
