package day17

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	_, ty, err := readInput(input)
	if err != nil {
		return 0, err
	}

	// setting trajectory is a mirrored rising one, with the maximal velocity the probe will just touch the lower edge of the target area
	vy0 := -ty[0] - 1
	// the deceleration is at the constant rate, the number of the steps is one away from the absolute value of the starting velocity
	n := vy0 + 1
	// the velocity at the highest point is 0, using the formula for the arithmetic progression
	maxY := vy0 * n / 2
	return maxY, nil
}

func part2(input string) (int, error) {
	tx, ty, err := readInput(input)
	if err != nil {
		return 0, err
	}

	stopStepMin := int(math.Ceil(findHorizontalVelocityToStopAtPosition(tx[0])))
	stopStepMax := int(math.Floor(findHorizontalVelocityToStopAtPosition(tx[1])))

	velocities := Velocities{}

	for n := 1; n <= -ty[0]*2; n++ {
		var x1 int
		if n >= stopStepMin {
			x1 = stopStepMin
		} else {
			x1 = int(math.Ceil(findVelocityForPositionAndStep(tx[0], n)))
		}
		var x2 int
		if n >= stopStepMax {
			x2 = stopStepMax
		} else {
			x2 = int(math.Floor(findVelocityForPositionAndStep(tx[1], n)))
		}
		for x := x1; x <= x2; x++ {
			y1 := int(math.Ceil(findVelocityForPositionAndStep(ty[0], n)))
			y2 := int(math.Floor(findVelocityForPositionAndStep(ty[1], n)))
			for y := y1; y <= y2; y++ {
				// Alternatively, squash values into int and use utils.IntSet
				velocities.Add(x, y)
			}
		}
	}

	return len(velocities), nil
}

func findHorizontalVelocityToStopAtPosition(position int) float64 {
	// find the positive root of quadratic equation
	c := 1 - 2*position
	d := 1 - 4*c
	return (-1 + math.Sqrt(float64(d))) / 2
}

func findVelocityForPositionAndStep(position int, step int) float64 {
	// Find the first item using the formula for the sum of the arithmetic progression
	s := float64(step)
	return (float64(position*2)/s + s - 1) / 2
}

type Velocity struct {
	x, y int
}

type Velocities map[Velocity]utils.Void

var void utils.Void

func (s *Velocities) Add(x, y int) {
	(*s)[Velocity{x, y}] = void
}

var pattern = regexp.MustCompile(`^target area: x=(-?\d+)\.\.(-?\d+), y=(-?\d+)\.\.(-?\d+)$`)

func readInput(input string) ([2]int, [2]int, error) {
	values := [4]int{}
	for i, s := range pattern.FindStringSubmatch(strings.TrimSpace(input))[1:] {
		v, err := strconv.Atoi(s)
		if err != nil {
			return [2]int{}, [2]int{}, err
		}
		values[i] = v
	}

	return [2]int{values[0], values[1]}, [2]int{values[2], values[3]}, nil
}
