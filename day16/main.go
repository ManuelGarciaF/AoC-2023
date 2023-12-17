package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	layout := parseInput(os.Args[1])

	fmt.Println("Part1:", solvePart1(layout, LaserHead{
		Coord: Coord{0, -1},
		dir:   RIGHT,
	}))
	fmt.Println("Part2:", solvePart2(layout))
}

type Coord struct {
	x, y int
}

func (c Coord) Add(other Coord) Coord {
	return Coord{
		x: c.x + other.x,
		y: c.y + other.y,
	}
}

type Direction int

const (
	UP Direction = iota
	RIGHT
	DOWN
	LEFT
)

var offsets = map[Direction]Coord{
	UP:    {-1, 0},
	RIGHT: {0, 1},
	DOWN:  {1, 0},
	LEFT:  {0, -1},
}

type LaserHead struct {
	Coord
	dir Direction
}

// '/' mirror
var mirrorOneDirMap = map[Direction]Direction{
	UP:    RIGHT,
	RIGHT: UP,
	DOWN:  LEFT,
	LEFT:  DOWN,
}

// '\' mirror
var mirrorTwoDirMap = map[Direction]Direction{
	UP:    LEFT,
	RIGHT: DOWN,
	DOWN:  RIGHT,
	LEFT:  UP,
}

func solvePart1(layout [][]byte, startingPos LaserHead) int {
	lasers := []LaserHead{startingPos}

	// Layout is square
	maxX, maxY := len(layout), len(layout[1])

	energizedTiles := make(map[Coord]Direction)

	for len(lasers) > 0 {
		for i := 0; i < len(lasers); i++ {
			if _, ok := energizedTiles[lasers[i].Coord]; !ok {
				energizedTiles[lasers[i].Coord] = lasers[i].dir
			}

			// Move laser forwards
			lasers[i].Coord = lasers[i].Coord.Add(offsets[lasers[i].dir])

			if lasers[i].x < 0 || lasers[i].x >= maxX || lasers[i].y < 0 || lasers[i].y >= maxY { // If out of bounds
				// Remove laser
				lasers = append(lasers[:i], lasers[i+1:]...)
				continue
			}

			switch layout[lasers[i].x][lasers[i].y] {
			case '/':
				lasers[i].dir = mirrorOneDirMap[lasers[i].dir]
			case '\\':
				lasers[i].dir = mirrorTwoDirMap[lasers[i].dir]
			case '|':
				if lasers[i].dir == LEFT || lasers[i].dir == RIGHT {
					lasers[i].dir = UP
					lasers = append(lasers, LaserHead{ // Make another laser going down
						Coord: lasers[i].Coord,
						dir:   DOWN,
					})
				}
			case '-':
				if lasers[i].dir == UP || lasers[i].dir == DOWN {
					lasers[i].dir = LEFT
					lasers = append(lasers, LaserHead{ // Make another laser going right
						Coord: lasers[i].Coord,
						dir:   RIGHT,
					})
				}
			}

			// If there has already been a laser going this direction here
			if dir, ok := energizedTiles[lasers[i].Coord]; ok && dir == lasers[i].dir {
				// Remove laser
				lasers = append(lasers[:i], lasers[i+1:]...)
				continue
			}
		}
	}

	return len(energizedTiles) - 1 // This includes an extra tile at (0, -1)
}

// Bruteforce time
func solvePart2(layout [][]byte) int {
	// Layout is square
	maxX, maxY := len(layout), len(layout[1])

	maxEnergized := 0
	for _, dir := range []Direction{UP, LEFT, DOWN, RIGHT} {
		if dir == RIGHT || dir == LEFT {
			y := -1 // dir == RIGHT
			if dir == LEFT {
				y = maxY
			}
			for x := 0; x < maxX; x++ {
				maxEnergized = max(maxEnergized, solvePart1(layout, LaserHead{
					Coord: Coord{x, y},
					dir:   dir,
				}))
			}
		} else {
			x := -1 // dir == DOWN
			if dir == UP {
				x = maxX
			}
			for y := 0; y < maxY; y++ {
				maxEnergized = max(maxEnergized, solvePart1(layout, LaserHead{
					Coord: Coord{x, y},
					dir:   dir,
				}))
			}
		}
	}
	return maxEnergized
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseInput(inputPath string) [][]byte {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	image := make([][]byte, 0)
	for scanner.Scan() {
		bytes := slices.Clone(scanner.Bytes())
		image = append(image, bytes)
	}
	return image
}
