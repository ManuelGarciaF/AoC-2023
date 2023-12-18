package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// x=0, y=0 at first character of first line
type NumberCoord struct {
	x, y int
	endX int
	n    int
}

type SymbolCoord struct {
	x, y   int
	symbol byte
}

func main() {
	// bytes, err := os.ReadFile("./inputSample")
	bytes, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println("Part 1 solution:", solvePart1(string(bytes)))
	fmt.Println("Part 2 solution:", solvePart2(string(bytes)))

}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove empty line at end.

	numbers, symbols := parseInput(lines)

	sum := 0
	for _, number := range numbers {
		if hasSymbolAdjacent(number, symbols) {
			sum += number.n
		}
	}

	return sum
}

func hasSymbolAdjacent(number NumberCoord, symbols []SymbolCoord) bool {
	for _, symbol := range symbols {
		if nextTo(number, symbol) {
			// fmt.Printf("%d has symbol %c next to it\n", number.n, symbol.symbol)
			return true
		}
		// Since symbols is ordered by y coord, we can stop if symbol.y - number.y > 1
		diff := symbol.y - number.y
		if diff > 1 {
			break
		}
	}
	return false
}

func nextTo(n NumberCoord, s SymbolCoord) bool {
	xOk := s.x >= (n.x-1) && s.x <= (n.endX+1)
	yOk := s.y >= (n.y-1) && s.y <= (n.y+1)
	return xOk && yOk
}

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1] // Remove empty line at end.

	numbers, symbols := parseInput(lines)

	sum := 0
	for _, symbol := range symbols {
		if symbol.symbol != byte('*') {
			continue
		}
		gearRatio, isGear := findGearRatio(symbol, numbers)
		if isGear {
			sum += gearRatio
		}
	}

	return sum
}

func findGearRatio(symbol SymbolCoord, numbers []NumberCoord) (gearRatio int, isGear bool) {
	numbersFound := 0
	gearRatio = 1
	for _, number := range numbers {
		if nextTo(number, symbol) {
			gearRatio *= number.n
			numbersFound++
			if numbersFound == 2 {
				return gearRatio, true
			}
		}
		// Since numbers is ordered by y coord, we can stop if number.y - symbol.y > 1
		diff := number.y - symbol.y
		if diff > 1 {
			break
		}
	}
	return 0, false
}

func parseInput(lines []string) ([]NumberCoord, []SymbolCoord) {
	numbers := make([]NumberCoord, 0)
	symbols := make([]SymbolCoord, 0)

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			switch {
			case line[x] == '.':
				break
			case unicode.IsDigit(rune(line[x])):
				n, length := parseNum(line[x:])
				endX := x + (length - 1)
				numbers = append(numbers, NumberCoord{x: x, y: y, endX: endX, n: n})
				x += length - 1
				break
			default: // Symbol
				symbols = append(symbols, SymbolCoord{x: x, y: y, symbol: line[x]})
				break
			}
		}
	}
	return numbers, symbols
}

func parseNum(str string) (n, length int) {
	length = 0
	for _, char := range str {
		if !unicode.IsDigit(char) {
			break
		}
		length++
	}
	n, _ = strconv.Atoi(str[:length])
	return n, length
}
