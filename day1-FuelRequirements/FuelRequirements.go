package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}

func ExtractInts(s string) []int {
	fs := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsDigit(r) && r != '-'
	})
	ints := make([]int, 0, len(fs))
	for _, w := range fs {
		i, err := strconv.Atoi(w)
		if err == nil {
			ints = append(ints, i)
		}
	}
	return ints
}

func FuelRequired(mass int) int {
	return (mass / 3) - 2
}

func FuelRequiredRecursive(mass int) int {
	mass = FuelRequired(mass)
	if mass < 0 {
		return 0
	}
	return mass + FuelRequiredRecursive(mass)
}

func main() {
	var input = ReadInput(os.Args[1])
	var numbers = ExtractInts(input)

	var sum = 0
	for _, num := range numbers {
		sum += FuelRequired(num)
	}

	fmt.Printf("Part 1: %d\n", sum)

	var sum2 = 0
	for _, num := range numbers {
		sum2 += FuelRequiredRecursive(num)
	}

	fmt.Printf("Part 2: %d\n", sum2)
}
