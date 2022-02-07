package day19

import "github.com/shpikat/advent-of-code-2021/utils"

type Coordinates map[Coordinate]utils.Void

var void utils.Void

func (s *Coordinates) Add(c Coordinate) {
	(*s)[c] = void
}

func (s Coordinates) Has(c Coordinate) (has bool) {
	_, has = s[c]
	return
}

type Stack []int

func (s Stack) HasMore() bool {
	return len(s) != 0
}

func (s *Stack) Pop() int {
	lastElement := len(*s) - 1
	popped := (*s)[lastElement]
	*s = (*s)[:lastElement]
	return popped
}

func (s *Stack) PushAll(elements []int) {
	*s = append(*s, elements...)
}
