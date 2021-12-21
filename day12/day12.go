package day12

import (
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

const (
	start = "start"
	end   = "end"
)

func part1(input string) (int, error) {
	caves := readInput(input)

	return caves.ExploreVisitingSmallOnce(utils.StringSet{}, start), nil
}

func part2(input string) (int, error) {
	caves := readInput(input)

	return caves.ExploreVisitingOneSmallTwice(utils.StringSet{}, start, true), nil
}

type Caves map[string][]string

func (c *Caves) AddConnection(name1 string, name2 string) {
	cave1, exists := (*c)[name1]
	if !exists {
		cave1 = make([]string, 0, 8)
	}
	cave2, exists := (*c)[name2]
	if !exists {
		cave2 = make([]string, 0, 8)
	}
	(*c)[name1] = append(cave1, name2)
	(*c)[name2] = append(cave2, name1)
}

func (c Caves) ExploreVisitingSmallOnce(visited utils.StringSet, name string) (count int) {
	if name == end {
		count = 1
	} else {
		newVisited := visited.Copy()
		newVisited.Add(name)
		for _, next := range c[name] {
			if !isSmallCave(next) || !newVisited.Has(next) {
				count += c.ExploreVisitingSmallOnce(newVisited, next)
			}
		}
	}
	return
}

func (c Caves) ExploreVisitingOneSmallTwice(visited utils.StringSet, name string, canVisitSmall bool) (count int) {
	if name == end {
		count = 1
	} else {
		newVisited := visited.Copy()
		newVisited.Add(name)
		for _, next := range c[name] {
			if !isSmallCave(next) || !newVisited.Has(next) {
				count += c.ExploreVisitingOneSmallTwice(newVisited, next, canVisitSmall)
			} else if canVisitSmall && next != start && next != end {
				count += c.ExploreVisitingOneSmallTwice(newVisited, next, false)
			}
		}
	}
	return
}

func isSmallCave(name string) bool {
	return 'a' <= name[0] && name[0] <= 'z'
}

func readInput(input string) Caves {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	caves := make(Caves, len(lines))
	for _, line := range lines {
		connected := strings.Split(line, "-")
		caves.AddConnection(connected[0], connected[1])
	}
	return caves
}
