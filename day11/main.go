package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type Coord struct {
	x, y int
}

type Image struct {
	stars []Coord
	xMax  int
	yMax  int
}

func main() {
	canvas := parseInput(os.Args[1])
	fmt.Println("Part1:", solve(canvas, 1))
	fmt.Println("Part2:", solve(canvas, 1e6 - 1)) // Add 999999
}

func solve(canvas [][]byte, expansionSize int) int {
	image := findStars(canvas)
	emptyRows, emptyCols := findEmptyLines(canvas)
	fmt.Printf("rows: %v, cols: %v\n", emptyRows, emptyCols)

	sum := 0
	for i, star := range image.stars {
		if i == len(image.stars)-1 { // Skip last star
			break
		}
		for j := i + 1; j < len(image.stars); j++ {
			sum += distanceWithExpansion(star, image.stars[j], emptyRows, emptyCols, expansionSize)
		}
	}

	return sum
}

func findEmptyLines(canvas [][]byte) ([]int, []int) {
	emptyRows := make([]int, 0)
	emptyCols := make([]int, 0)
	for x := 0; x < len(canvas); x++ {
		// All chars of row are '.'
		if all(canvas[x], func(b byte) bool { return b == '.' }) {
			emptyRows = append(emptyRows, x)
		}
	}
	// Colums
	for y := 0; y < len(canvas[0]); y++ {
		// Check
		allEmpty := true
		for x := 0; x < len(canvas); x++ {
			if canvas[x][y] != '.' {
				allEmpty = false
				break
			}
		}
		if allEmpty {
			emptyCols = append(emptyCols, y)
		}
	}
	return emptyRows, emptyCols
}

func distanceWithExpansion(a, b Coord, emptyRows, emptyCols []int, expansionSize int) int {
	dy := abs(a.y - b.y)
	dx := abs(a.x - b.x)
	// If there is an empty row in between, add expansion
	for _, row := range emptyRows {
		if (a.x < row && row < b.x) || (b.x < row && row < a.x) {
			dx += expansionSize
		}
	}
	// If there is an empty column in between, add expansion
	for _, col := range emptyCols {
		if (a.y < col && col < b.y) || (b.y < col && col < a.y) {
			dy += expansionSize
		}
	}
	return dx + dy
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func findStars(canvas [][]byte) Image {
	stars := make([]Coord, 0)
	for x := 0; x < len(canvas); x++ {
		for y := 0; y < len(canvas[x]); y++ {
			if canvas[x][y] == '#' {
				stars = append(stars, Coord{x, y})
			}
		}
	}
	return Image{stars, len(canvas), len(canvas[0])}
}

func parseInput(inputPath string) [][]byte {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	stars := make([][]byte, 0)
	for scanner.Scan() {
		line := slices.Clone(scanner.Bytes())
		stars = append(stars, line)
	}
	return stars
}

func all[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
