package day10

import (
	"sort"
	"strings"
)

const (
	openings = "([{<"
	closings = ")]}>"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var errorPoints = [...]int{3, 57, 1197, 25137}

	score := 0
	for _, line := range lines {
		expected := NewStack()
		for _, ch := range []byte(line) {
			index := strings.IndexByte(openings, ch)
			if index < 0 {
				if len(expected) == 0 || expected.MustPop() != ch {
					score += errorPoints[strings.IndexByte(closings, ch)]
					break
				}
			} else {
				expected.Push(closings[index])
			}
		}
	}

	return score, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	scores := make([]int, 0, len(lines))

lines:
	for _, line := range lines {
		expected := NewStack()
		for _, ch := range []byte(line) {
			index := strings.IndexByte(openings, ch)
			if index < 0 {
				if len(expected) == 0 || expected.MustPop() != ch {
					continue lines
				}
			} else {
				expected.Push(closings[index])
			}
		}

		score := 0
		for len(expected) > 0 {
			score = score*5 + strings.IndexByte(closings, expected.MustPop()) + 1
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)
	median := scores[len(scores)/2]

	return median, nil
}

type Stack []byte

func NewStack() Stack {
	return make([]byte, 0, 16)
}
func (s *Stack) Push(c byte) {
	*s = append(*s, c)
}

func (s *Stack) MustPop() (next byte) {
	last := len(*s) - 1
	next = (*s)[last]
	*s = (*s)[:last]
	return
}
