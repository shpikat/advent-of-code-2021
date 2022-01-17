package day18

import (
	"testing"

	"github.com/shpikat/advent-of-code-2021/internal"
)

const (
	sample1 = `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

	part1Sample = 4140
	part1Answer = 3216

	part2Sample = 3993
	part2Answer = 4643
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

func Benchmark(b *testing.B) {
	input := internal.ReadInput(b, "./testdata/input.txt")
	parts := []struct {
		name   string
		fn     func(input string) (int, error)
		answer int
	}{
		{"part1", part1, part1Answer},
		{"part2", part2, part2Answer},
	}

	for _, part := range parts {
		b.Run(part.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				got, err := part.fn(input)
				if err != nil {
					b.Errorf("error: %v", err)
				}
				if got != part.answer {
					b.Errorf("got: %v, answer: %v", got, part.answer)
				}
			}
		})
	}
}
