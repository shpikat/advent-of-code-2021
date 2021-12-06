package day06

import (
	"container/ring"
	"strconv"
	"strings"
)

const (
	breedingPeriod    = 7
	preBreedingPeriod = 2
)

func part1(input string) (int, error) {
	return calculatePopulation(input, 80)
}

func part2(input string) (int, error) {
	return calculatePopulation(input, 256)
}

func calculatePopulation(input string, days int) (int, error) {
	ages := [breedingPeriod]int{}
	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		ages[v] += 1
	}

	maturePopulation := ring.New(breedingPeriod)
	for _, age := range ages {
		maturePopulation.Value = age
		maturePopulation = maturePopulation.Next()
	}
	youngPopulation := ring.New(preBreedingPeriod)
	for i := 0; i < preBreedingPeriod; i++ {
		youngPopulation.Value = 0
		youngPopulation = youngPopulation.Next()
	}

	for i := 0; i < days; i++ {
		day7to6 := youngPopulation.Value.(int)
		youngPopulation.Value = maturePopulation.Value
		maturePopulation = maturePopulation.Move(1)
		youngPopulation = youngPopulation.Move(1)
		day6 := maturePopulation.Prev()
		day6.Value = day6.Value.(int) + day7to6
	}

	sum := 0
	for i := 0; i < breedingPeriod; i++ {
		sum += maturePopulation.Value.(int)
		maturePopulation = maturePopulation.Next()
	}
	for i := 0; i < preBreedingPeriod; i++ {
		sum += youngPopulation.Value.(int)
		youngPopulation = youngPopulation.Next()
	}

	return sum, nil
}
