package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"slices"
)

func main() {
	image := parseInput(os.Args[1])

	fmt.Println("Part1: ", solvePart1(image))

	image = parseInput(os.Args[1])
	fmt.Println("Part2: ", solvePart2(image))
}

func solvePart1(image Image) int {
	image = roll(image, NORTH)
	sum := 0
	for i, line := range image {
		load := len(image[0]) - i
		for _, char := range line {
			if char == 'O' {
				sum += load
			}
		}
	}
	return sum
}

const CYCLE_LIMIT = 1000000000 // 1e9

func solvePart2(image Image) int {
	prevRolls := []Image{image.clone()}
	var cyclesLeft int
	for i, loopFound := 0, false; i < CYCLE_LIMIT && !loopFound; i++ {
		image = cycle(image)

		for j, prev := range prevRolls {
			if reflect.DeepEqual(image, prev) { // Found a loop
				loopLength := len(prevRolls) - j
				cyclesLeft = (CYCLE_LIMIT - i - 1) % loopLength
				loopFound = true
			}
		}
		prevRolls = append(prevRolls, image.clone())
	}
	for i := 0; i < cyclesLeft; i++ {
		image = cycle(image)
	}

	sum := 0
	for i, line := range image {
		load := len(image[0]) - i
		for _, char := range line {
			if char == 'O' {
				sum += load
			}
		}
	}
	return sum
}

func cycle(image Image) Image {
	for _, dir := range []Direction{NORTH, WEST, SOUTH, EAST} {
		image = roll(image, dir)
	}
	return image
}

type Image [][]byte

func (i Image) clone() Image {
	clone := make(Image, len(i))
	for j := 0; j < len(i); j++ {
		clone[j] = slices.Clone(i[j])
	}
	return clone
}

type Direction int

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

var opposite = map[Direction]Direction{
	NORTH: SOUTH,
	EAST:  WEST,
	SOUTH: NORTH,
	WEST:  EAST,
}

func roll(image Image, dir Direction) Image {
	side := len(image) // The image is square
	for i := 0; i < side; i++ {
		// Pointers to the elements of the current line
		// Use opposite because we need to go backwards relative to the roll direction
		line := getLine(image, i, opposite[dir])

		lastSolid := -1
		for j := 0; j < side; j++ {
			switch *line[j] {
			case '#':
				lastSolid = j
			case 'O':
				nextSpot := lastSolid + 1
				if nextSpot == j {
					lastSolid = nextSpot
					continue
				}
				*line[nextSpot] = 'O'
				*line[j] = '.'
				lastSolid = nextSpot
			}
		}
	}

	return image
}

// Creates a slice of pointers to elements in the image in the direction given
func getLine(image Image, i int, dir Direction) []*byte {
	line := make([]*byte, len(image))
	// Which indices we use depends on the direction
	for j := 0; j < len(image); j++ {
		switch dir {
		case NORTH:
			line[(len(image)-1)-j] = &image[j][i]
		case EAST:
			line[j] = &image[i][j]
		case SOUTH:
			line[j] = &image[j][i]
		case WEST:
			line[(len(image)-1)-j] = &image[i][j]
		}
	}
	return line
}

func parseInput(inputPath string) Image {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	image := make(Image, 0)
	for scanner.Scan() {
		bytes := slices.Clone(scanner.Bytes())
		image = append(image, bytes)
	}
	return image
}
