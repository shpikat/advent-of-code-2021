package day03

import (
	"strings"
)

func part1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	nBits := len(lines[0])
	counts := make([]int, nBits)
	for _, line := range lines {
		offset := 0
		for {
			i := strings.IndexByte(line[offset:], '1')
			if i < 0 {
				break
			} else {
				offset += i + 1
				counts[nBits-offset]++
			}
		}
	}

	median := len(lines) / 2

	gammaRate := 0
	mask := -1
	for i, count := range counts {
		if count > median {
			gammaRate |= 1 << i
		}
		mask <<= 1
	}
	epsilonRate := ^(gammaRate | mask)

	return gammaRate * epsilonRate, nil
}

func part2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	var trie BinaryTrie
	for _, line := range lines {
		trie.add(line)
	}

	oxygenGeneratorRating := trie.findFunc(func(node BinaryTrie) int {
		if node.counts[0] > node.counts[1] {
			return 0
		}
		return 1
	})
	co2ScrubberRating := trie.findFunc(func(node BinaryTrie) int {
		if node.counts[0] <= node.counts[1] {
			return 0
		}
		return 1
	})

	return oxygenGeneratorRating * co2ScrubberRating, nil
}

type BinaryTrie struct {
	children [2]*BinaryTrie
	counts   [2]int
}

func (head *BinaryTrie) add(value string) {
	node := &head
	for _, ch := range value {
		if *node == nil {
			*node = &BinaryTrie{}
		}
		digit := ch - '0'
		(*node).counts[digit] += 1
		node = &(*node).children[digit]
	}
}

func (head BinaryTrie) findFunc(next func(node BinaryTrie) int) (result int) {
	node := &head
	for node != nil {
		var digit int
		if node.counts[0] == 0 && node.counts[1] == 1 {
			digit = 1
		} else if node.counts[0] == 1 && node.counts[1] == 0 {
			digit = 0
		} else {
			digit = next(*node)
		}
		result = result<<1 | digit
		node = node.children[digit]
	}
	return
}
