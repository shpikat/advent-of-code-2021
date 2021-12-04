package day04

import (
	"errors"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	numbers, boards, err := readInput(input)
	if err != nil {
		return 0, err
	}

	for _, number := range numbers {
		for _, board := range boards {
			if board.Mark(number) {
				return number * board.GetSumUnmarked(), nil
			}
		}
	}

	return 0, errors.New("no winning board found")
}

func part2(input string) (int, error) {
	numbers, boards, err := readInput(input)
	if err != nil {
		return 0, err
	}

	boardsInPlay := utils.IntSet{}
	for i := range boards {
		boardsInPlay.Add(i)
	}

	for _, number := range numbers {
		for i, board := range boards {
			if board.Mark(number) && boardsInPlay.Remove(i) && len(boardsInPlay) == 0 {
				return number * board.GetSumUnmarked(), nil
			}
		}
	}

	return 0, errors.New("no winning board found")
}

func readInput(input string) ([]int, []Board, error) {
	blocks := strings.Split(strings.TrimSpace(input), "\n\n")

	numbers, err := readNumbers(blocks[0])
	if err != nil {
		return nil, nil, err
	}

	boards := make([]Board, len(blocks)-1)
	for b, block := range blocks[1:] {
		boards[b] = NewBoard()
		for i, row := range strings.Split(block, "\n") {
			for j, value := range strings.Fields(row) {
				number, err := strconv.Atoi(value)
				if err != nil {
					return nil, nil, err
				}
				boards[b].Set(i, j, number)
			}
		}
	}
	return numbers, boards, nil
}

func readNumbers(s string) ([]int, error) {
	splits := strings.Split(s, ",")
	numbers := make([]int, len(splits))
	for i, value := range splits {
		n, err := strconv.Atoi(value)
		if err != nil {
			return nil, err
		}
		numbers[i] = n
	}
	return numbers, nil
}
