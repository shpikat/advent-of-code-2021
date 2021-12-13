package day13

import (
	"errors"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	dots, instructions := readInput(input)

	plane, err := createPlane(dots)
	if err != nil {
		return 0, err
	}

	firstLine := instructions[:strings.IndexByte(instructions, '\n')]
	err = fold(&plane, firstLine)
	if err != nil {
		return 0, err
	}

	return plane.Size(), nil
}

func part2(input string) (string, error) {
	dots, instructions := readInput(input)

	plane, err := createPlane(dots)
	if err != nil {
		return "", err
	}

	for _, line := range strings.Split(instructions, "\n") {
		err = fold(&plane, line)
		if err != nil {
			return "", err
		}
	}

	return plane.ToString(), nil
}

func readInput(input string) (string, string) {
	page1 := strings.Split(strings.TrimSpace(input), "\n\n")

	dots := page1[0]
	instructions := page1[1]
	return dots, instructions
}

func createPlane(input string) (utils.Plane, error) {
	plane := utils.NewPlane(len(input))
	for _, line := range strings.Split(input, "\n") {
		values := strings.Split(line, ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			return utils.Plane{}, err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return utils.Plane{}, err
		}
		plane.Add(utils.Coordinate{X: x, Y: y})
	}
	return plane, nil
}

func fold(plane *utils.Plane, line string) error {
	const offset = len("fold along ")
	foldAlong := line[offset]
	foldAt, err := strconv.Atoi(line[offset+2:])
	if err != nil {
		return err
	}

	// Just two small cases, we can duplicate most of the code with no worries at all
	if foldAlong == 'x' {
		toTheRight := make([]utils.Coordinate, 0, plane.Size())

		plane.DoFunc(func(p utils.Coordinate) {
			if p.X > foldAt {
				toTheRight = append(toTheRight, p)
			}
		})

		for _, p := range toTheRight {
			plane.Remove(p)
			plane.Add(utils.Coordinate{X: foldAt - (p.X - foldAt), Y: p.Y})
		}
		(*plane).Width = foldAt
	} else if foldAlong == 'y' {
		toTheBottom := make([]utils.Coordinate, 0, plane.Size())
		plane.DoFunc(func(p utils.Coordinate) {
			if p.Y > foldAt {
				toTheBottom = append(toTheBottom, p)
			}
		})

		for _, p := range toTheBottom {
			plane.Remove(p)
			plane.Add(utils.Coordinate{X: p.X, Y: foldAt - (p.Y - foldAt)})
		}
		(*plane).Height = foldAt
	} else {
		return errors.New("unexpected coordinate to fold along: " + string(foldAlong))
	}

	return nil
}
