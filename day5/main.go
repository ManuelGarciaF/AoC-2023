package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type RangeMapping struct {
	dstRangeStart int
	srcRangeStart int
	rangeLen      int
}

func (r RangeMapping) convert(input int) int {
	return r.dstRangeStart + input - r.srcRangeStart
}

func (r RangeMapping) inRange(input int) bool {
	return input >= r.srcRangeStart && input < (r.srcRangeStart+r.rangeLen)
}

type Mapping []RangeMapping

func (m Mapping) convert(input int) int {
	for _, r := range m {
		if r.inRange(input) {
			return r.convert(input)
		}
	}
	return input // Maps to the same number
}

func main() {
	// bytes, err := os.ReadFile("./inputSample")
	bytes, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePart1(string(bytes)))
	fmt.Println(solvePart2(string(bytes)))
}

func solvePart1(input string) int {
	seeds, mappings := parseInput(input)
	lowestLocation := -1
	for _, seed := range seeds {
		// println("seed: ", seed)
		mappedVal := seed
		for _, mapping := range mappings {
			mappedVal = mapping.convert(mappedVal)
			// println("\tmap", i, ":", mappedVal)
		}
		if mappedVal < lowestLocation || lowestLocation < 0 {
			lowestLocation = mappedVal
		}
	}

	return lowestLocation
}

func solvePart2(input string) int {
	seedRanges, mappings := parseInput(input)


	lowestLocation := -1
	for _, seed := range seeds {
		// println("seed: ", seed)
		mappedVal := seed
		for _, mapping := range mappings {
			mappedVal = mapping.convert(mappedVal)
			// println("\tmap", i, ":", mappedVal)
		}
		if mappedVal < lowestLocation || lowestLocation < 0 {
			lowestLocation = mappedVal
		}
	}

	return lowestLocation
}

func parseInput(input string) (seeds []int, mappings []Mapping) {
	input = input[:len(input)-1] // Remove trailing newline
	blocks := strings.Split(input, "\n\n")

	// Get seed numbers
	seedStrs := strings.Split(blocks[0], " ")
	seedStrs = seedStrs[1:] // Ignore "seeds:" part
	seeds = sliceAtoi(seedStrs)

	// Make maps
	mappings = make([]Mapping, 0, len(blocks[1:]))
	for _, block := range blocks[1:] {
		mapping := make(Mapping, 0, len(block))

		lines := strings.Split(block, "\n")
		lines = lines[1:] // Ignore first line
		for _, line := range lines {
			numbers := sliceAtoi(strings.Split(line, " "))
			mapping = append(mapping, RangeMapping{
				dstRangeStart: numbers[0],
				srcRangeStart: numbers[1],
				rangeLen:      numbers[2],
			})
		}
		mappings = append(mappings, mapping)
	}

	return
}

func sliceAtoi(strs []string) []int {
	ints := make([]int, 0, len(strs))
	for _, str := range strs {
		n, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			panic("Input not in expected format")
		}
		ints = append(ints, n)
	}
	return ints
}
