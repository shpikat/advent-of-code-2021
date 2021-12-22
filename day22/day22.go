package day22

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	instructions, err := readInput(input)
	if err != nil {
		return 0, err
	}

	bootRegion := Cuboid{
		Range{-50, 50},
		Range{-50, 50},
		Range{-50, 50},
	}
	grid := make(Grid, len(instructions)*8)
	for _, instruction := range instructions {
		c, intersects := bootRegion.intersect(instruction.cuboid)
		if intersects {
			if instruction.state {
				grid.turnOn(c)
			} else {
				grid.turnOff(c)
			}
		}
	}

	return grid.getCount(), nil
}

func part2(input string) (int, error) {
	instructions, err := readInput(input)
	if err != nil {
		return 0, err
	}

	grid := make(Grid, len(instructions)*8)
	for _, instruction := range instructions {
		if instruction.state {
			grid.turnOn(instruction.cuboid)
		} else {
			grid.turnOff(instruction.cuboid)
		}
	}

	return grid.getCount(), nil
}

type Range struct {
	start, end int
}

func (r Range) overlap(o Range) (Range, bool) {
	start := utils.Max(r.start, o.start)
	end := utils.Min(r.end, o.end)
	if start <= end {
		return Range{start, end}, true
	} else {
		return Range{}, false
	}
}

func (r Range) String() string {
	return fmt.Sprintf("%d..%d", r.start, r.end)
}

type Cuboid struct {
	x, y, z Range
}

func (c Cuboid) intersect(i Cuboid) (Cuboid, bool) {
	x, overlaps := c.x.overlap(i.x)
	if !overlaps {
		return Cuboid{}, false
	}
	y, overlaps := c.y.overlap(i.y)
	if !overlaps {
		return Cuboid{}, false
	}
	z, overlaps := c.z.overlap(i.z)
	if !overlaps {
		return Cuboid{}, false
	}
	return Cuboid{x, y, z}, true
}

func (c Cuboid) subtract(s Cuboid) (subtracted []Cuboid) {
	x := Range{s.x.start - 1, s.x.end + 1}
	if c.x.start <= x.start {
		subtracted = append(subtracted, Cuboid{
			Range{c.x.start, x.start},
			c.y,
			c.z,
		})
	}
	if x.end <= c.x.end {
		subtracted = append(subtracted, Cuboid{
			Range{x.end, c.x.end},
			c.y,
			c.z,
		})
	}
	y := Range{s.y.start - 1, s.y.end + 1}
	if c.y.start <= y.start {
		subtracted = append(subtracted, Cuboid{
			s.x,
			Range{c.y.start, y.start},
			c.z,
		})
	}
	if y.end <= c.y.end {
		subtracted = append(subtracted, Cuboid{
			s.x,
			Range{y.end, c.y.end},
			c.z,
		})
	}
	z := Range{s.z.start - 1, s.z.end + 1}
	if c.z.start <= z.start {
		subtracted = append(subtracted, Cuboid{
			s.x,
			s.y,
			Range{c.z.start, z.start},
		})
	}
	if z.end <= c.z.end {
		subtracted = append(subtracted, Cuboid{
			s.x,
			s.y,
			Range{z.end, c.z.end},
		})
	}

	return
}

func (c Cuboid) volume() int {
	return (c.x.end - c.x.start + 1) * (c.y.end - c.y.start + 1) * (c.z.end - c.z.start + 1)
}

func (c Cuboid) String() string {
	return fmt.Sprintf("%v, %v, %v", c.x, c.y, c.z)
}

type Grid map[Cuboid]utils.Void

var void utils.Void

func (g *Grid) add(c Cuboid) {
	(*g)[c] = void
}

func (g *Grid) remove(c Cuboid) {
	delete(*g, c)
}

func (g *Grid) turnOn(cuboid Cuboid) {
	adding := make(Grid, 16)
	adding.add(cuboid)
	for c := range *g {
		// fast path is checking for covering cuboid
		// some indexing like R-tree should be even faster, but this is fast enough
		_, canPossiblyIntersect := c.intersect(cuboid)
		if canPossiblyIntersect {
			var remove, add []Cuboid
			for a := range adding {
				intersection, intersects := c.intersect(a)
				if intersects {
					remove = append(remove, a)
					add = append(add, a.subtract(intersection)...)
				}
			}
			for _, r := range remove {
				adding.remove(r)
			}
			for _, a := range add {
				adding.add(a)
			}
			if len(adding) == 0 {
				break
			}
		}
	}
	for c := range adding {
		g.add(c)
	}
}

func (g *Grid) turnOff(cuboid Cuboid) {
	var remove, add []Cuboid
	for c := range *g {
		intersection, intersects := c.intersect(cuboid)
		if intersects {
			remove = append(remove, c)
			add = append(add, c.subtract(intersection)...)
		}
	}
	for _, c := range remove {
		g.remove(c)
	}
	for _, c := range add {
		g.add(c)
	}
}

func (g Grid) getCount() (count int) {
	for c := range g {
		count += c.volume()
	}
	return
}

type Instruction struct {
	state  bool
	cuboid Cuboid
}

func readInput(input string) ([]Instruction, error) {
	const coordinatesInInstruction = 6

	lines := strings.Split(strings.TrimSpace(input), "\n")
	instructions := make([]Instruction, len(lines))
	pattern := regexp.MustCompile(`^\w+ x=(-?\d+)\.\.(-?\d+),y=(-?\d+)\.\.(-?\d+),z=(-?\d+)\.\.(-?\d+)$`)
	numbers := [coordinatesInInstruction]int{}
	for i, line := range lines {
		state := strings.HasPrefix(line, "on")
		submatches := pattern.FindStringSubmatch(line)
		for j := 0; j < coordinatesInInstruction; j++ {
			var err error
			numbers[j], err = strconv.Atoi(submatches[j+1])
			if err != nil {
				return nil, err
			}
		}
		instructions[i] = Instruction{
			state,
			Cuboid{
				Range{numbers[0], numbers[1]},
				Range{numbers[2], numbers[3]},
				Range{numbers[4], numbers[5]},
			},
		}
	}
	return instructions, nil
}
