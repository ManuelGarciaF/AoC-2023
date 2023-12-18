package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/ManuelGarciaF/AoC-2023/utils"
)

func main() {
	heatLosses, sideLen := parseInput(os.Args[1])

	fmt.Println("Part1: ", solve(heatLosses, sideLen, 0, 3))
	fmt.Println("Part2: ", solve(heatLosses, sideLen, 3, 10))
}

type Node struct {
	utils.Coord
	prevDir       utils.Direction
	straightSteps int
}

func solve(heatLosses map[utils.Coord]int, sideLen int, minStraight, maxStraight int) int {
	// Dijkstra's algorithm
	costs := make(map[Node]int, len(heatLosses))
	cameFrom := make(map[Node]Node, len(heatLosses))
	frontier := utils.NewPriorityQueue[Node]()

	start := Node{
		Coord:         utils.Coord{X: 0, Y: 0}, // Start at top-left
		prevDir:       -1,                      // No incoming direction
		straightSteps: 0,
	}

	frontier.PushItem(start, 0)

	costs[start] = 0

	goal := utils.Coord{X: sideLen - 1, Y: sideLen - 1}

	lastNode := start
	for !frontier.IsEmpty() {
		current := frontier.PopItem()
		if current.Coord == goal {
			lastNode = current
			break
		}

		for _, neighbour := range neighbours(heatLosses, current.Coord) {
			currDir := utils.DirFromOffset[neighbour.Sub(current.Coord)]
			turning := current.prevDir != currDir
			// Don't allow turning before maxStraight steps
			if !turning && current.straightSteps+1 >= maxStraight {
				continue
			}
			// Don't allow turning back
			if utils.InverseDir[currDir] == current.prevDir {
				continue
			}
			// Avoid turning before minStraight steps
			if turning && current.straightSteps < minStraight && current.prevDir != -1 {
				continue
			}

			newStraightSteps := 0
			if !turning {
				newStraightSteps = current.straightSteps + 1
			}
			cost := costs[current] + heatLosses[neighbour]
			next := Node{
				Coord:         neighbour,
				prevDir:       currDir,
				straightSteps: newStraightSteps,
			}
			if _, ok := costs[next]; !ok || cost < costs[next] {
				costs[next] = cost
				frontier.PushItem(next, cost)
				cameFrom[next] = current
			}
		}
	}

	// We have a reverse path to the goal in cameFrom
	current := lastNode
	heatLoss := 0
	for current.Coord != start.Coord {
		heatLoss += heatLosses[current.Coord]
		current = cameFrom[current]
	}

	return heatLoss
}

func neighbours(m map[utils.Coord]int, c utils.Coord) []utils.Coord {
	neighbours := make([]utils.Coord, 0, 4)
	for _, offset := range utils.Offsets {
		nCoord := c.Add(offset)
		if _, ok := m[nCoord]; ok {
			neighbours = append(neighbours, nCoord)
		}
	}
	return neighbours
}

func parseInput(inputPath string) (heatLosses map[utils.Coord]int, sideLen int) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	heatLosses = make(map[utils.Coord]int)
	x := 0
	for scanner.Scan() {
		for y, c := range scanner.Text() { // For char in line
			n, err := strconv.Atoi(string(c))
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
			heatLosses[utils.Coord{X: x, Y: y}] = n
		}
		x++
	}
	sideLen = x
	return
}
