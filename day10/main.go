package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type PipeLayout [][]byte

func (m PipeLayout) getCoord(c Coord) byte {
	if c.x < 0 || c.x >= len(m) || c.y < 0 || c.y >= len(m[c.x]) {
		return ' '
	}
	return m[c.x][c.y]
}

type Coord struct {
	x, y int
}

func (c Coord) add(other Coord) Coord {
	return Coord{c.x + other.x, c.y + other.y}
}
func (c Coord) sub(other Coord) Coord {
	return Coord{c.x - other.x, c.y - other.y}
}

// Offsets
var NORTH = Coord{-1, 0}
var EAST = Coord{0, 1}
var SOUTH = Coord{1, 0}
var WEST = Coord{0, -1}

func main() {
	inputPath := os.Args[1]
	pipes := parseInput(inputPath)

	// Part 1
	fmt.Println("Part1:", solvePart1(pipes))

	// Part 2
	fmt.Println("Part2:", solvePart2(pipes))
}

func solvePart1(pipes PipeLayout) int {
	start := startPos(pipes)
	currPos := start
	next := nextPosition(pipes, currPos, Coord{-1, -1}) // Use a dummy value for prevPos
	distances := make([]int, 0)
	distances = append(distances, 0)
	for i := 1; next != start; i++ {
		currPos, next = next, nextPosition(pipes, next, currPos)
		distances = append(distances, i)
	}
	// The furthest point will be at half of the loop
	return distances[len(distances)/2]
}

func solvePart2(pipes PipeLayout) int {
	// Find loop
	start := startPos(pipes)
	currPos := start
	next := nextPosition(pipes, currPos, Coord{-1, -1}) // Use a dummy value for prevPos
	loop := make([]Coord, 0)
	loop = append(loop, start)
	for i := 1; next != start; i++ {
		currPos, next = next, nextPosition(pipes, next, currPos)
		loop = append(loop, currPos)
	}

	// Shoelace formula
	sum := 0
	for i := 0; i < len(loop); i++ {
		p1, p2 := loop[i], loop[(i+1)%len(loop)]
		sum += (p1.y + p2.y) * (p1.x - p2.x)
	}
	area := sum / 2
	if area < 0 {
		area *= -1
	}
	// This area considers paths next to each other to have a 1 unit gap
	// we can find the interior points with Pick's theorem
	// A = i + b/2 - 1
	return area - (len(loop)/2) + 1
}

// Possible connections from a given pipe
var possibleConnections = map[byte][]Coord{
	'S': {NORTH, EAST, SOUTH, WEST},
	'|': {NORTH, SOUTH},
	'-': {EAST, WEST},
	'L': {NORTH, EAST},
	'J': {NORTH, WEST},
	'7': {SOUTH, WEST},
	'F': {SOUTH, EAST},
	'.': {},
	' ': {},
}
var inverseDirections = map[Coord]Coord{
	NORTH: SOUTH,
	EAST:  WEST,
	SOUTH: NORTH,
	WEST:  EAST,
}

func nextPosition(pipes PipeLayout, currPos Coord, prevPos Coord) Coord {
	for _, dir := range possibleConnections[pipes.getCoord(currPos)] {
		next := currPos.add(dir)
		if next == prevPos {
			continue
		}
		// If possible direction can connect back
		if slices.Index(possibleConnections[pipes.getCoord(next)], inverseDirections[dir]) != -1 {
			return next
		}
	}
	panic("No next position found")
}

func startPos(pipes PipeLayout) Coord {
	for x, line := range pipes {
		for y, char := range line {
			if char == 'S' {
				return Coord{x, y}
			}
		}
	}
	panic("No start position found")
}

func parseInput(inputPath string) PipeLayout {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := make(PipeLayout, 0, 150)
	for scanner.Scan() {
		bytes := slices.Clone(scanner.Bytes()) // Weird stuff was happening
		lines = append(lines, bytes)
	}
	return lines
}

func graphGrid(grid [][]Cell) {
	for _, line := range grid {
		for _, cell := range line {
			if cell == UNDECIDED {
				fmt.Print(" ")
			} else if cell == PIPE {
				fmt.Print(".")
			} else {
				fmt.Print("X")
			}
		}
		fmt.Println()
	}
}
