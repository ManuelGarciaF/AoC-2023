package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Node struct {
	left  string
	right string
}

func main() {
	// nodes, instructions := parseInput("./sample")
	// nodes, instructions := parseInput("./sample2")
	nodes, instructions := parseInput("./input")

	// fmt.Println(solvePart1(nodes, instructions))
	fmt.Println(solvePart2(nodes, instructions))
}

func solvePart1(nodes map[string]Node, instructions string) int {
	steps := 0
	for i, currLocation := 0, "AAA"; currLocation != "ZZZ"; i++ {
		i %= len(instructions)
		currLocation = step(nodes[currLocation], instructions[i])
		steps++
	}
	return steps
}
func solvePart2Dumb(nodes map[string]Node, instructions string) int {
	steps := 0
	positions := make([]string, 0)
	// Filter only positions that end with A
	for name := range nodes {
		if name[2] == 'A' {
			positions = append(positions, name)
		}
	}

	for instructionStep := 0; !all(positions, func(p string) bool { return p[2] == 'Z' }); instructionStep++ {
		instructionStep %= len(instructions)

		if steps%1e6 == 0 {
			fmt.Println(steps)
		}

		switch instructions[instructionStep] {
		case 'L':
			for i, p := range positions {
				positions[i] = nodes[p].left
			}
		case 'R':
			for i, p := range positions {
				positions[i] = nodes[p].right
			}
		default:
			panic("Unknown instruction")
		}
		steps++
	}

	return steps
}

func solvePart2(nodes map[string]Node, instructions string) int {
	loopLengths := make([]int, 0)
	// Filter only positions that end with A
	for name := range nodes {
		if name[2] != 'A' {
			continue
		}
		currLocation := name
		// Go to the first Z
		i := 0
		for ; currLocation[2] != 'Z'; i++ {
			i %= len(instructions)
			currLocation = step(nodes[currLocation], instructions[i])
		}
		// Advance a step
		currLocation = step(nodes[currLocation], instructions[0])
		i++
		loopLen := 1
		// Check length to next repetition
		for ; currLocation[2] != 'Z'; i++ {
			i %= len(instructions)
			currLocation = step(nodes[currLocation], instructions[i])
			loopLen++
		}

		loopLengths = append(loopLengths, loopLen)
	}
	// Since for some unexplained reason, all paths are on closed regular loops, we can just take the LCM of all loop lengths
	fmt.Println("Loops are: ", loopLengths)
	return lcm(loopLengths)
}

// Why do you make me implement all this
func lcm(ns []int) int {
	lcm := 1
	for _, n := range ns {
		lcm = lcm2(lcm, n)
	}
	return lcm
}

func lcm2(a, b int) int {
	return (a / gcd(a, b)) * b
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func step(curr Node, instruction byte) string {
	switch instruction {
	case 'L':
		return curr.left
	case 'R':
		return curr.right
	default:
		panic("Unknown instruction")
	}
}

func parseInput(inputPath string) (nodes map[string]Node, instructions string) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan() // First line for instructions
	instructions = scanner.Text()
	scanner.Scan() // Ignore empty line

	re := regexp.MustCompile(`([0-9A-Z]{3}) = \(([0-9A-Z]{3}), ([0-9A-Z]{3})\)`)

	nodes = make(map[string]Node, 0)
	for scanner.Scan() {
		tokens := re.FindStringSubmatch(scanner.Text())
		nodes[tokens[1]] = Node{
			left:  tokens[2],
			right: tokens[3],
		}
	}

	return
}

func all[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
