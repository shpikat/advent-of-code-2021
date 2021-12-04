package day04

import "github.com/shpikat/advent-of-code-2021/utils"

const (
	boardRows    = 5
	boardColumns = 5
)

type Board struct {
	rows    [boardRows]utils.IntSet
	columns [boardColumns]utils.IntSet
}

func NewBoard() (b Board) {
	for i := range b.rows {
		b.rows[i] = utils.IntSet{}
	}
	for i := range b.columns {
		b.columns[i] = utils.IntSet{}
	}
	return
}

func (b *Board) Set(row, column, value int) {
	(*b).rows[row].Add(value)
	(*b).columns[column].Add(value)
}

func (b *Board) Mark(number int) (hasWon bool) {
	for _, row := range b.rows {
		if row.Remove(number) && len(row) == 0 {
			hasWon = true
		}
	}
	for _, column := range b.columns {
		if column.Remove(number) && len(column) == 0 {
			hasWon = true
		}
	}
	return
}

func (b Board) GetSumUnmarked() (sum int) {
	for _, row := range b.rows {
		for number := range row {
			sum += number
		}
	}
	return
}
