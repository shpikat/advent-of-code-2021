package day13

import (
	"strings"
	"testing"

	"github.com/shpikat/advent-of-code-2021/internal"
)

const (
	sample1 = `
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`

	part1Sample = 17
	part1Answer = 747

	part2Sample = `
#####
#...#
#...#
#...#
#####
.....
.....
`
	part2Answer = `
.##..###..#..#.####.###...##..#..#.#..#.
#..#.#..#.#..#....#.#..#.#..#.#..#.#..#.
#..#.#..#.####...#..#..#.#....#..#.####.
####.###..#..#..#...###..#....#..#.#..#.
#..#.#.#..#..#.#....#....#..#.#..#.#..#.
#..#.#..#.#..#.####.#.....##...##..#..#.
`
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
		want  string
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
			if strings.TrimSpace(got) != strings.TrimSpace(tc.want) {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
