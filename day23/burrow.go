package day23

import (
	"math"
	"sort"

	"github.com/shpikat/advent-of-code-2021/utils"
)

const (
	Empty  = 0xF
	Amber  = 0
	Bronze = 1
	Copper = 2
	Desert = 3

	NumberOfTypes = 4

	HallwayLength = 11

	Mask    = 0xF
	LenBits = 4
)

var (
	costs = [NumberOfTypes]int{
		Amber:  1,
		Bronze: 10,
		Copper: 100,
		Desert: 1000,
	}
	roomEntrances = [NumberOfTypes]int{
		Amber:  2,
		Bronze: 4,
		Copper: 6,
		Desert: 8,
	}

	// There are no stops right outside the room
	hallwayStops = []int{0, 1, 3, 5, 7, 9, 10}
)

type Amphipod uint

type Room uint64

type Hallway struct {
	Room
}

type Burrow struct {
	hallway  Hallway
	rooms    [NumberOfTypes]Room
	roomSize int
}

type Move struct {
	Burrow
	energy int
}

func (r *Room) set(place int, a Amphipod) {
	shift := place * LenBits
	t := uint64(*r) & ^(Mask << shift)
	t |= uint64(a) << shift
	*r = Room(t)
}

func (r Room) get(place int) Amphipod {
	return Amphipod((r >> (place * LenBits)) & Mask)
}

func (h Hallway) hasWay(start, end int) bool {
	if start != end {
		var from, to int
		if start < end {
			from, to = start+1, end+1
		} else {
			from, to = end, start
		}
		// Can be further optimized by having a table of masks to compare the range against
		for p := h.Room >> (from * LenBits); p != h.Room>>(to*LenBits); p >>= LenBits {
			if p&Mask != Empty {
				return false
			}
		}
	}
	return true
}

func NewBurrow(roomSize int) Burrow {
	return Burrow{
		hallway:  Hallway{math.MaxUint64},
		roomSize: roomSize,
	}
}

func (b Burrow) findLeastEnergyRequired() (int, error) {
	energies := make(map[Burrow]int, 100_000)

	final := b.getFinalOrganization()

	leastEnergy := math.MaxInt

	stack := NewStack()
	stack.Push([]Move{{b, 0}})

	for len(stack) != 0 {
		current := stack.MustPop()
		if current.energy < leastEnergy {
			previous, exists := energies[current.Burrow]
			if !exists || previous > current.energy {
				energies[current.Burrow] = current.energy
				if current.Burrow == final {
					leastEnergy = utils.Min(leastEnergy, current.energy)
				} else {
					stack.Push(current.getNextMoves())
				}
			}
		}
	}

	return leastEnergy, nil
}

func (b Burrow) getFinalOrganization() Burrow {
	b.hallway = Hallway{math.MaxUint64}
	for i := range b.rooms {
		a := Amphipod(i)
		for place := 0; place < b.roomSize; place++ {
			b.rooms[i].set(place, a)
		}
	}
	return b
}

func (m Move) getNextMoves() []Move {
	var moves []Move

	// Try to move from the hallway into their destinations first
	for location := 0; location < HallwayLength; location++ {
		a := m.hallway.get(location)
		if a != Empty {
			var p int
			for p = m.roomSize - 1; p >= 0 && m.rooms[a].get(p) == a; p-- {
			}
			if p >= 0 && m.rooms[a].get(p) == Empty && m.hallway.hasWay(location, roomEntrances[a]) {
				moves = append(moves, m.moveToOwnRoom(location, p))
			}
		}
	}

	// If the room is not the destination, move to all the available locations in the hallway
rooms:
	for i, room := range m.rooms {
		owner := Amphipod(i)

		for place := 0; place < m.roomSize; place++ {
			if room.get(place) != owner {
				for p := 0; p <= place; p++ {
					if room.get(p) != Empty {
						roomExit := sort.SearchInts(hallwayStops, roomEntrances[i])
						for left := roomExit - 1; left >= 0 && m.hallway.get(hallwayStops[left]) == Empty; left-- {
							moves = append(moves, m.moveToHall(owner, p, hallwayStops[left]))
						}
						for right := roomExit; right < len(hallwayStops) && m.hallway.get(hallwayStops[right]) == Empty; right++ {
							moves = append(moves, m.moveToHall(owner, p, hallwayStops[right]))
						}
						continue rooms
					}
				}
			}
		}
	}

	return moves
}

func (m Move) moveToOwnRoom(location int, place int) Move {
	a := m.hallway.get(location)
	m.hallway.set(location, Empty)
	m.rooms[a].set(place, a)
	m.energy += (utils.Abs(location-roomEntrances[a]) + place + 1) * costs[a]
	return m
}

func (m Move) moveToHall(room Amphipod, place int, location int) Move {
	a := m.rooms[room].get(place)
	m.rooms[room].set(place, Empty)
	m.hallway.set(location, a)
	m.energy += (utils.Abs(location-roomEntrances[room]) + place + 1) * costs[a]
	return m
}
