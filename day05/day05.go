package day05

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

var pattern = regexp.MustCompile(`^(\d+),(\d+) -> (\d+),(\d+)$`)

func part1(input string) (int, error) {
	lines, err := readInput(input)
	if err != nil {
		return 0, err
	}

	points := map[utils.Coordinate]int{}
	for _, line := range lines {
		if line.from.X == line.to.X || line.from.Y == line.to.Y {
			line.UpdatePoints(points)
		}
	}

	count := 0
	for _, value := range points {
		if value > 1 {
			count++
		}
	}

	return count, nil
}

func part2(input string) (int, error) {
	lines, err := readInput(input)
	if err != nil {
		return 0, err
	}

	points := map[utils.Coordinate]int{}
	for _, line := range lines {
		line.UpdatePoints(points)
	}

	count := 0
	for _, value := range points {
		if value > 1 {
			count++
		}
	}

	return count, nil
}

type Line struct {
	from, to utils.Coordinate
}

func (l Line) UpdatePoints(points map[utils.Coordinate]int) {
	unitVector := l.getUnitVector()
	point := l.from
	for {
		points[point] += 1
		if point == l.to {
			break
		}
		point = point.Add(unitVector)
	}
}

// getUnitVector calculates Euclidean unit vector for the given cases of 0, 45 and 90 degrees
func (l Line) getUnitVector() utils.Coordinate {
	vector := l.to.Subtract(l.from)
	length := utils.Max(utils.Abs(vector.X), utils.Abs(vector.Y))
	return utils.Coordinate{X: vector.X / length, Y: vector.Y / length}
}

func readInput(input string) ([]Line, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	clouds := make([]Line, len(lines))
	for i, line := range lines {
		values := [4]int{}
		for j, submatch := range pattern.FindStringSubmatch(line)[1:] {
			value, err := strconv.Atoi(submatch)
			if err != nil {
				return nil, err
			}
			values[j] = value
		}
		clouds[i] = Line{
			utils.Coordinate{X: values[0], Y: values[1]},
			utils.Coordinate{X: values[2], Y: values[3]},
		}
	}
	return clouds, nil
}
