package day24

const length = 14

// pen-and-paper kind of solution: the input was simplified through search-and-replace, there was absolutely no need for coding this routine by that time
var solution = []InputCondition{
	{0, 13, 7},
	{1, 4, 3},
	{3, 2, 8},
	{5, 10, 4},
	{6, 9, 1},
	{8, 7, 2},
	{11, 12, 5},
}

func part1(_ string) (int, error) {
	number := Number{}
	for _, c := range solution {
		number[c.greaterValueIndex] = 9
		number[c.lowerValueIndex] = 9 - c.positiveDelta
	}
	return number.getValue(), nil
}

func part2(_ string) (int, error) {
	number := Number{}
	for _, c := range solution {
		number[c.lowerValueIndex] = 1
		number[c.greaterValueIndex] = 1 + c.positiveDelta
	}
	return number.getValue(), nil
}

type InputCondition struct {
	greaterValueIndex, lowerValueIndex, positiveDelta int
}

type Number [length]int

func (n Number) getValue() int {
	value := 0
	for _, digit := range n {
		value = value*10 + digit
	}
	return value
}
