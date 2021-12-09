package day09

import (
	"container/heap"
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

	largestBasins := NewTopK(3)
	for i := range points {
		for j := range points[i] {
			current := Index{i, j}
			if current.get(points) != 9 {
				stack := NewStack()
				stack.Push(current)
				current.set(points, 9)
				size := 1
				for len(stack) > 0 {
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
				largestBasins.Add(size)
			}
		}
	}

	mul := 1
	for _, v := range largestBasins.Drain() {
		mul *= v
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

type Stack []Index

func NewStack() Stack {
	return make([]Index, 0, 16)
}
func (s *Stack) Push(v Index) {
	*s = append(*s, v)
}

func (s *Stack) MustPop() (next Index) {
	last := len(*s) - 1
	next = (*s)[last]
	*s = (*s)[:last]
	return
}

type IntTopK []int

func NewTopK(k int) (topK IntTopK) {
	topK = make([]int, 0, k+1)
	heap.Init(&topK)
	return
}

func (t *IntTopK) Add(x int) {
	if len(*t) < 4 {
		heap.Push(t, x)
	} else {
		// Implementation actually puts lower values on top, so we can easily replace it
		(*t)[0] = x
		heap.Fix(t, 0)
	}
}

func (t *IntTopK) Drain() (topK []int) {
	if len(*t) > 3 {
		heap.Pop(t)
	}
	topK = make([]int, 0, 3)
	for len(*t) > 0 {
		topK = append(topK, heap.Pop(t).(int))
	}
	return topK
}

func (t IntTopK) Len() int           { return len(t) }
func (t IntTopK) Less(i, j int) bool { return t[i] < t[j] }
func (t IntTopK) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }

func (t *IntTopK) Push(x interface{}) {
	*t = append(*t, x.(int))
}
func (t *IntTopK) Pop() (element interface{}) {
	old := *t
	last := len(old) - 1
	element = old[last]
	*t = old[:last]
	return
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
