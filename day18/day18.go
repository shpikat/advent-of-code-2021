package day18

import (
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	numbers, err := readInput(input)
	if err != nil {
		return 0, err
	}

	sum := numbers[0]
	for _, n := range numbers[1:] {
		sum.Add(n)
	}

	return sum.GetMagnitude(), nil
}

func part2(input string) (int, error) {
	numbers, err := readInput(input)
	if err != nil {
		return 0, err
	}

	max := 0
	for _, a := range numbers {
		for _, b := range numbers {
			if a != b {
				max = utils.Max(max, Copy(a).Add(Copy(b)).GetMagnitude())
			}
		}
	}
	return max, nil
}

// SnailfishNumber is a type used to store the snailfish number.
//
// For simplicity most of the methods are recursive, as the depth of the
// recursion should never exceed the level of 4. The reduce method is iterative
// for better flow control.
type SnailfishNumber struct {
	left, right *SnailfishNumber
	value       *int
}

// Add makes a a part of n, possibly updating both of them during reduction.
// In case the original values must be preserved, the copy of them must be made
// beforehand.
func (n *SnailfishNumber) Add(a *SnailfishNumber) *SnailfishNumber {
	old := *n
	*n = SnailfishNumber{left: &old, right: a}
	n.reduce()
	return n
}

func (n *SnailfishNumber) reduce() {
	isChanged := true
	for isChanged {
		isChanged = false

		explode := findExplode([]*SnailfishNumber{n}, 4)
		var firstRegularNumberToTheLeft *int
		var addToFirstRegularNumberToTheRight *int

		stack := NewStack()
		stack.Push(n)
		for len(stack) > 0 {
			current := stack.MustPop()
			if current == explode {
				// explode
				if firstRegularNumberToTheLeft != nil {
					*firstRegularNumberToTheLeft += *explode.left.value
				}
				addToFirstRegularNumberToTheRight = explode.right.value

				(*explode).left = nil
				(*explode).right = nil
				(*explode).value = new(int)
				isChanged = true
			} else if current.value == nil {
				if current.right != nil {
					stack.Push(current.right)
				}
				if current.left != nil {
					stack.Push(current.left)
				}
			} else {
				if addToFirstRegularNumberToTheRight != nil {
					*current.value += *addToFirstRegularNumberToTheRight
					break
				} else if explode == nil && *current.value >= 10 {
					// split only when there's nothing to explode
					left := *current.value / 2
					right := *current.value - left
					(*current).left = &SnailfishNumber{value: &left}
					(*current).right = &SnailfishNumber{value: &right}
					(*current).value = nil
					isChanged = true
					break
				}
				firstRegularNumberToTheLeft = current.value
			}
		}
	}
}

func findExplode(nodes []*SnailfishNumber, depth int) *SnailfishNumber {
	if len(nodes) == 0 {
		return nil
	}
	if depth == 0 {
		for _, node := range nodes {
			if node.value == nil {
				// If it's not the literal value--it's the node
				return node
			}
		}
		return nil
	}
	next := make([]*SnailfishNumber, 0, 16)
	for _, node := range nodes {
		if node.left != nil {
			next = append(next, node.left)
		}
		if node.right != nil {
			next = append(next, node.right)
		}
	}
	return findExplode(next, depth-1)
}

func (n SnailfishNumber) GetMagnitude() int {
	if n.value == nil {
		return 3*n.left.GetMagnitude() + 2*n.right.GetMagnitude()
	} else {
		return *n.value
	}
}

func Copy(n *SnailfishNumber) *SnailfishNumber {
	c := SnailfishNumber{}
	if n.left != nil {
		c.left = Copy(n.left)
	}
	if n.right != nil {
		c.right = Copy(n.right)
	}
	if n.value != nil {
		c.value = new(int)
		*c.value = *n.value
	}

	return &c
}

func (n SnailfishNumber) String() string {
	if n.value == nil {
		return "[" + n.left.String() + "," + n.right.String() + "]"
	} else {
		return strconv.Itoa(*n.value)
	}
}

func readInput(input string) ([]*SnailfishNumber, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	numbers := make([]*SnailfishNumber, len(lines))
	for i, line := range lines {
		n, _, err := readSnailfishNumber(line, 0)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}

func readSnailfishNumber(input string, i int) (*SnailfishNumber, int, error) {
	if input[i] == '[' {
		left, i, err := readSnailfishNumber(input, i+1)
		if err != nil {
			return nil, 0, err
		}
		right, i, err := readSnailfishNumber(input, i+1)
		if err != nil {
			return nil, 0, err
		}
		return &SnailfishNumber{left: left, right: right}, i + 1, nil
	} else {
		mark := i
		for input[i] >= '0' && input[i] <= '9' {
			i++
		}
		v, err := strconv.Atoi(input[mark:i])
		if err != nil {
			return nil, 0, err
		}
		return &SnailfishNumber{value: &v}, i, nil
	}
}

type Stack []*SnailfishNumber

func NewStack() Stack {
	return make([]*SnailfishNumber, 0, 16)
}
func (s *Stack) Push(e *SnailfishNumber) {
	*s = append(*s, e)
}

func (s *Stack) MustPop() (next *SnailfishNumber) {
	last := len(*s) - 1
	next = (*s)[last]
	*s = (*s)[:last]
	return
}
