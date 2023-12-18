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

func (r RangeMapping) inverseConvert(input int) int {
	return r.srcRangeStart + input - r.dstRangeStart
}

func (r RangeMapping) inRange(input int) bool {
	return input >= r.srcRangeStart && input < (r.srcRangeStart+r.rangeLen)
}

func (r RangeMapping) inDstRange(input int) bool {
	return input >= r.dstRangeStart && input < (r.dstRangeStart+r.rangeLen)
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

func (m Mapping) inverseConvert(input int) int {
	for _, r := range m {
		if r.inDstRange(input) {
			return r.inverseConvert(input)
		}
	}
	return input // Maps to the same number
}

type Range struct {
	start  int
	length int
}

func (r Range) inRange(input int) bool {
	return input >= r.start && input < (r.start+r.length)
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
	seedInfo, mappings := parseInput(input)
	seedRanges := parseSeedRanges(seedInfo)

	location := 0
	for !locationInSeeds(location, seedRanges, mappings) { // Until location is in the seed ranges
		location++
	}

	return location
}

func parseSeedRanges(seedInfo []int) []Range {
	ranges := make([]Range, 0, len(seedInfo)/2)
	for i := 0; i < len(seedInfo); i += 2 {
		ranges = append(ranges, Range{
			start:  seedInfo[i],
			length: seedInfo[i+1],
		})
	}
	return ranges
}

func locationInSeeds(location int, seedRanges []Range, mappings []Mapping) bool {
	seed := seedFromLocation(location, mappings)

	for _, seedRange := range seedRanges {
		if seedRange.inRange(seed) {
			return true
		}
	}
	return false
}

func seedFromLocation(location int, mappings []Mapping) int {
	value := location
	// Iterate over mappings backwards
	for i := len(mappings) - 1; i >= 0; i-- {
		value = mappings[i].inverseConvert(value)
	}
	return value
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
