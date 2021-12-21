package day20

import (
	"strings"
)

func part1(input string) (int, error) {
	algorithm, image := readInput(input)

	count := algorithm.enhance(image, 2)

	return count, nil
}

func part2(input string) (int, error) {
	algorithm, image := readInput(input)

	count := algorithm.enhance(image, 50)

	return count, nil
}

type Algorithm [8]uint64

func (a Algorithm) enhance(input [][]byte, steps int) int {
	const neighboursMask = 0b111
	infiniteValue := 0
	for step := 0; step < steps; step++ {
		emptyInputValue := infiniteValue * neighboursMask
		infiniteValue = a.get(infiniteValue * 0b111111111)
		emptyOutputValue := infiniteValue * neighboursMask

		rows := len(input)
		l := len(input[0])

		output := make([][]byte, rows+2)
		output[0] = make([]byte, l+2)
		value := emptyOutputValue
		for i := 0; i < l; i++ {
			empty := emptyInputValue<<6 | emptyInputValue<<3
			index := empty | int(input[0][i])
			value = (value<<1)&0b110 | a.get(index)
			output[0][i] = byte(value)
		}
		value = (value<<1)&0b110 | infiniteValue
		output[0][l] = byte(value)
		value = (value<<1)&0b110 | infiniteValue
		output[0][l+1] = byte(value)

		output[1] = make([]byte, l+2)
		value = emptyOutputValue
		for i := 0; i < l; i++ {
			index := emptyInputValue<<6 | int(input[0][i])<<3 | int(input[1][i])
			value = (value<<1)&0b110 | a.get(index)
			output[1][i] = byte(value)
		}
		value = (value<<1)&0b110 | infiniteValue
		output[1][l] = byte(value)
		value = (value<<1)&0b110 | infiniteValue
		output[1][l+1] = byte(value)

		for j := 2; j < rows; j++ {
			output[j] = make([]byte, l+2)
			value = emptyOutputValue
			for i := 0; i < l; i++ {
				index := int(input[j-2][i])<<6 | int(input[j-1][i])<<3 | int(input[j][i])
				value = (value<<1)&0b110 | a.get(index)
				output[j][i] = byte(value)
			}
			value = (value<<1)&0b110 | infiniteValue
			output[j][l] = byte(value)
			value = (value<<1)&0b110 | infiniteValue
			output[j][l+1] = byte(value)
		}

		output[rows] = make([]byte, l+2)
		value = emptyOutputValue
		for i := 0; i < l; i++ {
			index := int(input[rows-2][i])<<6 | int(input[rows-1][i])<<3 | emptyInputValue
			value = (value<<1)&0b110 | a.get(index)
			output[rows][i] = byte(value)
		}
		value = (value<<1)&0b110 | infiniteValue
		output[rows][l] = byte(value)
		value = (value<<1)&0b110 | infiniteValue
		output[rows][l+1] = byte(value)

		output[rows+1] = make([]byte, l+2)
		value = emptyOutputValue
		for i := 0; i < l; i++ {
			index := int(input[rows-1][i])<<6 | emptyInputValue<<3 | emptyInputValue
			value = (value<<1)&0b110 | a.get(index)
			output[rows+1][i] = byte(value)
		}
		value = (value<<1)&0b110 | infiniteValue
		output[rows+1][l] = byte(value)
		value = (value<<1)&0b110 | infiniteValue
		output[rows+1][l+1] = byte(value)

		input = output
	}

	count := 0
	for i := range input {
		for j := range input[i] {
			count += int(input[i][j] & 0b001)
		}
	}
	return count
}

func (a *Algorithm) set(i int) {
	const mask = 1<<6 - 1
	word := i >> 6
	bit := i & mask
	(*a)[word] |= 1 << bit
}

func (a Algorithm) get(i int) int {
	const mask = 1<<6 - 1
	word := i >> 6
	bit := i & mask
	return int((a[word] >> bit) & 1)
}

func readInput(input string) (algorithm Algorithm, image [][]byte) {
	blocks := strings.Split(input, "\n\n")
	for i, ch := range strings.TrimSpace(blocks[0]) {
		if ch == '#' {
			algorithm.set(i)
		}
	}
	lines := strings.Split(strings.TrimSpace(blocks[1]), "\n")
	image = make([][]byte, len(lines))
	for i, line := range lines {
		l := len(line)
		image[i] = make([]byte, l+2)
		value := 0
		for j, ch := range line {
			value = (value << 1) & 0b110
			if ch == '#' {
				value |= 0b001
			}
			image[i][j] = byte(value)
		}
		image[i][l] = byte((value << 1) & 0b110)
		image[i][l+1] = byte((value << 2) & 0b100)
	}
	return
}
