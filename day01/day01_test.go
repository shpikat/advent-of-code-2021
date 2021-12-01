package day01

import (
	"testing"

	"github.com/shpikat/advent-of-code-2021/internal"
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{
			"sample 1",
			`
199
200
208
210
200
207
240
269
260
263
`, 7,
		},
		{
			"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), 1527,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := solvePart1(tc.input)
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
		{
			"sample 1",
			`
199
200
208
210
200
207
240
269
260
263
`, 5,
		},
		{
			"puzzle input", internal.ReadInput(t, "./testdata/input.txt"), 1575,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := solvePart2(tc.input)
			if err != nil {
				t.Errorf("error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
