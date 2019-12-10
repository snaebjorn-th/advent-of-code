package main

import (
	"fmt"

	"github.com/snaebjorn-th/advent-of-code/utils"
)

func fuelRequired(mass int) int {
	return (mass / 3) - 2
}

func fuelRequiredRecursive(mass int) int {
	mass = fuelRequired(mass)
	if mass < 0 {
		return 0
	}
	return mass + fuelRequiredRecursive(mass)
}

func main() {
	var input = utils.ReadInput(utils.GetFilename())
	var numbers = utils.ExtractInts(input)

	var sum = 0
	for _, num := range numbers {
		sum += fuelRequired(num)
	}

	fmt.Printf("Part 1: %d\n", sum)

	var sum2 = 0
	for _, num := range numbers {
		sum2 += fuelRequiredRecursive(num)
	}

	fmt.Printf("Part 2: %d\n", sum2)
}
