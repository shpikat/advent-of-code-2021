package day16

import (
	"errors"
	"math/big"
	"strings"
)

func part1(input string) (int, error) {
	bits, offset, err := readInput(input)
	if err != nil {
		return 0, err
	}

	packet, _ := Decode(bits, offset)
	sum := packet.GetVersionSum()
	return sum, nil
}

func part2(input string) (int, error) {
	bits, offset, err := readInput(input)
	if err != nil {
		return 0, err
	}

	packet, _ := Decode(bits, offset)
	result := packet.Calculate()
	return int(result.Int64()), nil
}

func readInput(input string) (*big.Int, uint, error) {
	input = strings.TrimSpace(input)

	bits, ok := new(big.Int).SetString(input, 16)
	if !ok {
		return nil, 0, errors.New("failed to convert the value " + input)
	}
	offset := uint(len(input) * 4)
	return bits, offset, nil
}
