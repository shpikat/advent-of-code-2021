package day16

import (
	"testing"

	"github.com/shpikat/advent-of-code-2021/internal"
)

const (
	part1Answer = 889
	part2Answer = 739303923668
)

func TestPart1(t *testing.T) {
	testCases := []struct {
		name  string
		input string
		want  int
	}{
		{"sample 1", `8A004A801A8002F478`, 16},
		{"sample 2", `620080001611562C8802118E34`, 12},
		{"sample 3", `C0015000016115A2E0802F182340`, 23},
		{"sample 4", `A0016C880162017C3686B18A3D4780`, 31},
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
		{"sum of 1 and 2", `C200B40A82`, 3},
		{"product of 6 and 9", `04005AC33890`, 54},
		{"minumum of 7, 8, 9", `880086C3E88112`, 7},
		{"maximum of 7, 8, 9", `CE00C43D881120`, 9},
		{"5 < 15 ?", `D8005AC2A8F0`, 1},
		{"5 > 15 ?", `F600BC2D8F`, 0},
		{"5 == 15 ?", `9C005AC2F8F0`, 0},
		{"1+3 == 2*2 ?", `9C0141080250320F1802104A08`, 1},
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
