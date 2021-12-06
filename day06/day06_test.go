package day06

import (
	"testing"

	"github.com/shpikat/advent-of-code-2021/internal"
)

const (
	sample1 = `3,4,3,1,2`

	part1Sample = 5934
	part1Answer = 358214

	part2Sample = 26984457539
	part2Answer = 1622533344325
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part1Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part1Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part1(tc.input)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", sample1, part2Sample},
		{"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), part2Answer},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := part2(tc.input)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
