package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	steps := parseInput(os.Args[1])

	fmt.Println("Part1:", solvePart1(steps))
	fmt.Println("Part2:", solvePart2(steps))
}

func solvePart1(steps []string) int {
	sum := 0
	for _, step := range steps {
		sum += int(hash(step))
	}
	return sum
}

func hash(s string) byte {
	sum := 0
	for _, c := range s {
		if c == '\n' { // Ignore newlines
			continue
		}
		sum += int(c)
		sum *= 17
		sum %= 256
	}
	return byte(sum)
}

type Lens struct {
	label string
	focalLength int
}

func solvePart2(steps []string) int {
	opRE := regexp.MustCompile(`(.+)([-=])(.*)`)

	boxes := make(map[byte][]Lens)

	for _, instruction := range steps {
		tokens := opRE.FindStringSubmatch(instruction)
		label := tokens[1]
		op := tokens[2]

		switch op {
		case "=":
			focalLength, err := strconv.Atoi(tokens[3])
			if err != nil {
				panic("Found invalid focal length: " + err.Error())
			}
			// Search for lens in box
			found := false
			for i, lens := range boxes[hash(label)] {
				if lens.label == label {
					boxes[hash(label)][i].focalLength = focalLength
					found = true
					break
				}
			}
			if !found {
				boxes[hash(label)] = append(boxes[hash(label)], Lens{label, focalLength})
			}
		case "-":
			for i, lens := range boxes[hash(label)] {
				if lens.label == label {
					boxes[hash(label)] = append(boxes[hash(label)][:i], boxes[hash(label)][i+1:]...) // Remove that element
					break
				}
			}
		}
	}

	power := 0
	for boxNum := 0; boxNum < 256; boxNum++ {
		for slotNum, lens := range boxes[byte(boxNum)] {
			power += (boxNum+1) * (slotNum+1) * lens.focalLength
		}
	}

	return power
}

func parseInput(inputPath string) []string {
	bytes, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	return strings.Split(string(bytes), ",")
}
