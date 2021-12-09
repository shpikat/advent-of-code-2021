package day09

import (
	"sort"
	"strings"
)

func part1(input string) (int, error) {
	points := readInput(input)

	sum := 0
	for i := range points {
		for j := range points[i] {
			current := points[i][j]
			if (j == 0 || points[i][j-1] > current) &&
				(j == len(points[i])-1 || points[i][j+1] > current) &&
				(i == 0 || points[i-1][j] > current) &&
				(i == len(points)-1 || points[i+1][j] > current) {
				sum += current + 1
			}
		}
	}

	return sum, nil
}

func part2(input string) (int, error) {
	points := readInput(input)

	neighbours := []func(Index) Index{
		func(v Index) Index {
			return Index{v.i + 1, v.j}
		},
		func(v Index) Index {
			return Index{v.i - 1, v.j}
		},
		func(v Index) Index {
			return Index{v.i, v.j - 1}
		},
		func(v Index) Index {
			return Index{v.i, v.j + 1}
		},
	}

	preAllocationLength := len(points)
	basins := make([]int, 0, preAllocationLength)
	for i := range points {
		for j := range points[i] {
			current := Index{i, j}
			if current.get(points) != 9 {
				stack := NewStack()
				stack.Push(current)
				current.set(points, 9)
				size := 1
				for stack.HasMore() {
					next := stack.MustPop()
					for _, f := range neighbours {
						neighbour := f(next)
						if neighbour.within(points) {
							p := neighbour.get(points)
							if p != 9 {
								stack.Push(neighbour)
								neighbour.set(points, 9)
								size++
							}
						}
					}
				}
				basins = append(basins, size)
			}
		}
	}

	//TODO  implement selection algorithm
	sort.Ints(basins)
	mul := 1
	for _, basin := range basins[len(basins)-3:] {
		mul *= basin
	}
	return mul, nil
}

type Index struct {
	i, j int
}

func (v Index) get(points [][]int) int {
	return points[v.i][v.j]
}

func (v Index) set(points [][]int, value int) {
	points[v.i][v.j] = value
}

func (v Index) within(points [][]int) bool {
	return v.i >= 0 && v.j >= 0 && v.i < len(points) && v.j < len(points[v.i])
}

type Stack struct {
	data []Index
}

func NewStack() Stack {
	return Stack{make([]Index, 0, 16)}
}
func (s *Stack) Push(v Index) {
	s.data = append(s.data, v)
}

func (s *Stack) MustPop() (next Index) {
	last := len(s.data) - 1
	next = s.data[last]
	s.data = s.data[:last]
	return
}

func (s Stack) HasMore() bool {
	return len(s.data) > 0
}

func readInput(input string) [][]int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	points := make([][]int, len(lines))
	for i, line := range lines {
		points[i] = make([]int, len(line))
		for j, ch := range line {
			points[i][j] = int(ch - '0')
		}
	}
	return points
}
