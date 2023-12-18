package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/ManuelGarciaF/AoC-2023/utils"
)

func main() {
	rows := parseInput(os.Args[1], false)
	fmt.Println("Part1: ", solve(rows))
	rows = parseInput(os.Args[1], true)
	fmt.Println("Part2: ", solve(rows))
}

type Row struct {
	dir    utils.Direction
	meters int
}

func solve(rows []Row) int {
	dug := []utils.Coord{{X: 0, Y: 0}}

	vertices := 0
	curr := utils.Coord{X: 0, Y: 0}
	for _, row := range rows {
		newCoord := curr.Add(utils.Offsets[row.dir].Scale(row.meters))
		vertices += row.meters
		dug = append(dug, newCoord)
		curr = newCoord
	}

	// Shoelace formula
	sum := 0
	for i := 0; i < len(dug); i++ {
		p1, p2 := dug[i], dug[(i+1)%len(dug)]
		sum += (p1.Y + p2.Y) * (p1.X - p2.X)
	}
	poligonArea := sum / 2
	if poligonArea < 0 {
		poligonArea *= -1
	}
	// Pick's theorem
	interiorPoints := poligonArea - (vertices / 2) + 1
	area := interiorPoints + (vertices)

	return area
}

func parseInput(inputPath string, useColors bool) []Row {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := make([]Row, 0)

	re := regexp.MustCompile(`([UDLR]) (\d+) (\(#[0-9a-f]{6}\))`)

	for scanner.Scan() {
		tokens := re.FindStringSubmatch(scanner.Text())

		var dir utils.Direction
		var meters int
		if useColors {
			color := tokens[3][2 : len(tokens[3])-1]
			uMeters, err := strconv.ParseUint(color[:5], 16, 64)
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
			meters = int(uMeters)
			switch color[5] {
			case '0':
				dir = utils.RIGHT
			case '1':
				dir = utils.DOWN
			case '2':
				dir = utils.LEFT
			case '3':
				dir = utils.UP
			}
		} else {
			switch tokens[1] {
			case "U":
				dir = utils.UP
			case "D":
				dir = utils.DOWN
			case "L":
				dir = utils.LEFT
			case "R":
				dir = utils.RIGHT
			}
			meters, err = strconv.Atoi(tokens[2])
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
		}

		rows = append(rows, Row{
			dir:    dir,
			meters: meters,
		})
	}

	return rows
}
