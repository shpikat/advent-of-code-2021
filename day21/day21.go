package day21

import (
	"strings"

	"github.com/shpikat/advent-of-code-2021/utils"
)

func part1(input string) (int, error) {
	position1, position2 := readInput(input)

	return playWithDeterministicDice(Player{position: position1}, Player{position: position2}), nil
}

func part2(input string) (int, error) {
	position1, position2 := readInput(input)

	return playWithDiracDice(Player{position: position1}, Player{position: position2}), nil
}

type Player struct {
	position int
	score    int
}

func (p *Player) makeMove(dice int) int {
	(*p).position += dice
	if p.position > 10 {
		(*p).position -= 10
	}
	(*p).score += p.position
	return p.score
}

func playWithDeterministicDice(p1 Player, p2 Player) int {
	dice := 6
	rolls := 3

	for {
		if p1.makeMove(dice) >= 1000 {
			return p2.score * rolls
		}
		dice, rolls = rollDeterministicDice(dice, rolls)

		if p2.makeMove(dice) >= 1000 {
			return p1.score * rolls
		}
		dice, rolls = rollDeterministicDice(dice, rolls)
	}
}

func rollDeterministicDice(dice int, rolls int) (int, int) {
	// as long as every next sum of three is different by 9, and we need only the lowest digit,
	// in the end we can just subtract 1 to get what we need
	if dice == 0 {
		dice = 9
	} else {
		dice--
	}
	return dice, rolls + 3
}

func playWithDiracDice(p1 Player, p2 Player) int {
	wins := rollQuantumDiceRecursivelyWithCache(p1, p2, make(map[[2]Player][2]int, 10*10*21*21))

	return utils.Max(wins[0], wins[1])
}

// Actual values
//var waysToThrowSumOfThree = [10]int{
//	0, 0, 0, 1, 3, 6, 7, 6, 3, 1,
//}
var waysToThrowSumOfThree = createSums()

func createSums() (sums [10]int) {
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				sums[i+j+k] += 1
			}
		}
	}
	return
}

func rollQuantumDiceRecursivelyWithCache(player1, player2 Player, cache map[[2]Player][2]int) [2]int {
	wins, hasInCache := cache[[2]Player{player1, player2}]
	if !hasInCache {
		for dice := 3; dice <= 9; dice++ {
			quantumPlayer1 := player1
			if quantumPlayer1.makeMove(dice) >= 21 {
				wins[0] += waysToThrowSumOfThree[dice]
			} else {
				w := rollQuantumDiceRecursivelyWithCache(player2, quantumPlayer1, cache)
				wins[0] += w[1] * waysToThrowSumOfThree[dice]
				wins[1] += w[0] * waysToThrowSumOfThree[dice]
			}
		}
		cache[[2]Player{player1, player2}] = wins
	}
	return wins
}

func readInput(input string) (int, int) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	return getLastDigit(lines[0]), getLastDigit(lines[1])
}

func getLastDigit(input string) int {
	return int(input[len(input)-1] - '0')
}
