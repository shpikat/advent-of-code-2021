package utils

import "fmt"

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
