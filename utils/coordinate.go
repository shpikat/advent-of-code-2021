package utils

import (
	"fmt"
	"strings"
)

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Add(coordinate Coordinate) Coordinate {
	return Coordinate{c.X + coordinate.X, c.Y + coordinate.Y}
}

func (c Coordinate) Subtract(coordinate Coordinate) Coordinate {
	return Coordinate{c.X - coordinate.X, c.Y - coordinate.Y}
}

func (c Coordinate) String() string {
	return fmt.Sprintf("(%d, %d)", c.X, c.Y)
}

type Plane struct {
	points map[Coordinate]Void
	Width  int
	Height int
}

func NewPlane(capacity int) Plane {
	return Plane{points: make(map[Coordinate]Void, capacity)}
}

func (p *Plane) Add(coordinate Coordinate) (added bool) {
	if !p.Has(coordinate) {
		(*p).points[coordinate] = void
		(*p).Width = Max(p.Width, coordinate.X)
		(*p).Height = Max(p.Height, coordinate.Y)
		added = true
	}
	return
}

func (p *Plane) Remove(coordinate Coordinate) (removed bool) {
	removed = p.Has(coordinate)
	delete((*p).points, coordinate)
	return
}

func (p Plane) Has(coordinate Coordinate) (has bool) {
	_, has = p.points[coordinate]
	return
}

func (p Plane) Size() int {
	return len(p.points)
}

func (p Plane) DoFunc(fn func(p Coordinate)) {
	for point := range p.points {
		fn(point)
	}
}

func (p Plane) ToString() string {
	// For code simplicity create a redundant slice instead of reusing the first row
	empty := make([]byte, p.Width)
	for i := range empty {
		empty[i] = '.'
	}
	grid := make([][]byte, p.Height)
	for i := range grid {
		grid[i] = make([]byte, p.Width)
		copy(grid[i], empty)
	}

	for p := range p.points {
		grid[p.Y][p.X] = '#'
	}

	lines := make([]string, len(grid))
	for i := range grid {
		lines[i] = string(grid[i])
	}

	return strings.Join(lines, "\n")
}
