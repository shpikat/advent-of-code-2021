package day19

import (
	"errors"
	"strconv"
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

const guaranteedOverlappingBeacons = 12

func part1(input string) (int, error) {
	scanners, err := readInput(input)
	if err != nil {
		return 0, err
	}

	orientations := findOrientations(scanners)

	beacons := make(Coordinates, len(scanners)*len(scanners[0].beacons))
	for i, scanner := range scanners {
		orientation := orientations[i]
		for _, beacon := range scanner.beacons {
			beacons.Add(orientation.translate(beacon))
		}
	}

	return len(beacons), nil
}

func part2(input string) (int, error) {
	scanners, err := readInput(input)
	if err != nil {
		return 0, err
	}

	orientations := findOrientations(scanners)

	max := 0
	for i := 0; i < len(orientations)-1; i++ {
		origin1 := orientations[i].origin
		for j := i + 1; j < len(orientations); j++ {
			origin2 := orientations[j].origin
			max = utils.Max(max, origin1.getManhattanDistance(origin2))
		}
	}

	return max, nil
}

func findOrientations(scanners []Scanner) map[int]Orientation {
	beaconsByDistances := make(map[int]map[int][]int, len(scanners))
	unoriented := make(map[int]Scanner, len(scanners))

	for n, s := range scanners {
		distances := make(map[int][]int, getConnectionsCount(len(s.beacons)))
		for i := 0; i < len(s.beacons)-1; i++ {
			for j := i + 1; j < len(s.beacons); j++ {
				distance := s.beacons[i].getManhattanDistance(s.beacons[j])
				distances[distance] = append(distances[distance], i, j)
			}
		}
		beaconsByDistances[n] = distances

		unoriented[n] = scanners[n]
	}

	orientations := make(map[int]Orientation, len(scanners))
	orientations[0] = Orientation{
		origin: Coordinate{},
		translate: func(coordinate Coordinate) Coordinate {
			return coordinate
		},
	}
	const first = 0
	next := Stack{first}
	delete(unoriented, first)
	for next.HasMore() {
		n := next.Pop()
		baseScanner := scanners[n]

		found := make([]int, 0, 16)
	scan:
		for _, scanner := range unoriented {
			count := 0
			commonBeaconsBase := utils.IntSet{}
			commonBeacons := utils.IntSet{}
			for distance, beacons := range beaconsByDistances[scanner.n] {
				baseBeacons, exist := beaconsByDistances[baseScanner.n][distance]
				if exist {
					count++
					for _, b := range beacons {
						commonBeacons.Add(b)
					}
					for _, b := range baseBeacons {
						commonBeaconsBase.Add(b)
					}
				}
			}

			if count >= getConnectionsCount(guaranteedOverlappingBeacons) {
				baseBeacons := make(Coordinates, len(commonBeaconsBase))
				for b := range commonBeaconsBase {
					c := baseScanner.beacons[b]
					baseBeacons.Add(c)
				}

				for i := range commonBeaconsBase {
					for j := range commonBeacons {
						for _, fn := range allPossibleTranslations {
							origin := fn(scanner.beacons[j]).getVector(baseScanner.beacons[i])
							count := 0
							for b := range commonBeacons {
								if baseBeacons.Has(origin.addVector(fn(scanner.beacons[b]))) {
									count++
									if count >= guaranteedOverlappingBeacons {
										baseOrientation := orientations[baseScanner.n]
										orientations[scanner.n] = Orientation{
											origin: baseOrientation.translate(origin),
											translate: func(c Coordinate) Coordinate {
												return baseOrientation.translate(origin.addVector(fn(c)))
											},
										}
										found = append(found, scanner.n)

										continue scan
									}
								}
							}
						}
					}
				}
			}
		}

		next.PushAll(found)
		for _, n := range found {
			delete(unoriented, n)
		}
	}

	return orientations
}

func getConnectionsCount(n int) int {
	return (n - 1) * n / 2
}

type Coordinate struct {
	x, y, z int
}

func (c Coordinate) getManhattanDistance(target Coordinate) int {
	return abs(target.x-c.x) + abs(target.y-c.y) + abs(target.z-c.z)
}

func (c Coordinate) getVector(target Coordinate) Coordinate {
	return Coordinate{target.x - c.x, target.y - c.y, target.z - c.z}
}

func (c Coordinate) addVector(target Coordinate) Coordinate {
	return Coordinate{target.x + c.x, target.y + c.y, target.z + c.z}
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

type Scanner struct {
	n       int
	beacons []Coordinate
}

type Orientation struct {
	origin    Coordinate
	translate func(Coordinate) Coordinate
}

// Calculated and inlined
var allPossibleTranslations = []func(Coordinate) Coordinate{
	func(c Coordinate) Coordinate { return c },
	func(c Coordinate) Coordinate { return Coordinate{-c.y, c.x, c.z} },
	func(c Coordinate) Coordinate { return Coordinate{-c.x, -c.y, c.z} },
	func(c Coordinate) Coordinate { return Coordinate{c.y, -c.x, c.z} },
	func(c Coordinate) Coordinate { return Coordinate{c.y, -c.z, -c.x} },
	func(c Coordinate) Coordinate { return Coordinate{c.z, c.y, -c.x} },
	func(c Coordinate) Coordinate { return Coordinate{-c.y, c.z, -c.x} },
	func(c Coordinate) Coordinate { return Coordinate{-c.z, -c.y, -c.x} },
	func(c Coordinate) Coordinate { return Coordinate{-c.z, c.x, -c.y} },
	func(c Coordinate) Coordinate { return Coordinate{-c.x, -c.z, -c.y} },
	func(c Coordinate) Coordinate { return Coordinate{c.z, -c.x, -c.y} },
	func(c Coordinate) Coordinate { return Coordinate{c.x, c.z, -c.y} },
	func(c Coordinate) Coordinate { return Coordinate{-c.x, c.y, -c.z} },
	func(c Coordinate) Coordinate { return Coordinate{-c.y, -c.x, -c.z} },
	func(c Coordinate) Coordinate { return Coordinate{c.x, -c.y, -c.z} },
	func(c Coordinate) Coordinate { return Coordinate{c.y, c.x, -c.z} },
	func(c Coordinate) Coordinate { return Coordinate{c.y, c.z, c.x} },
	func(c Coordinate) Coordinate { return Coordinate{-c.z, c.y, c.x} },
	func(c Coordinate) Coordinate { return Coordinate{-c.y, -c.z, c.x} },
	func(c Coordinate) Coordinate { return Coordinate{c.z, -c.y, c.x} },
	func(c Coordinate) Coordinate { return Coordinate{c.z, c.x, c.y} },
	func(c Coordinate) Coordinate { return Coordinate{-c.x, c.z, c.y} },
	func(c Coordinate) Coordinate { return Coordinate{-c.z, -c.x, c.y} },
	func(c Coordinate) Coordinate { return Coordinate{c.x, -c.z, c.y} },
}

func readInput(input string) ([]Scanner, error) {
	sections := strings.Split(strings.TrimSpace(input), "\n\n")
	scanners := make([]Scanner, len(sections))
	for i, section := range sections {
		lines := strings.Split(strings.TrimSpace(section), "\n")
		n, err := strconv.Atoi(lines[0][len("--- scanner ") : len(lines[0])-len(" ---")])
		if err != nil {
			return nil, err
		}

		if i != n {
			return nil, errors.New("Scanner number is out of order: " + strconv.Itoa(n))
		}

		beacons := make([]Coordinate, len(lines)-1)
		for j, line := range lines[1:] {
			a := [3]int{}
			for k, s := range strings.Split(line, ",") {
				a[k], err = strconv.Atoi(s)
				if err != nil {
					return nil, err
				}
			}
			beacons[j] = Coordinate{a[0], a[1], a[2]}
		}

		scanners[n] = Scanner{n, beacons}
	}
	return scanners, nil
}
