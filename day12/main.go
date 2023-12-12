package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type SpringRow struct {
	layout []byte
	groups []int
}

func main() {
	rows := parseInput(os.Args[1])

	fmt.Println("Part1: ", solvePart1(rows))
}

func solvePart1(rows []SpringRow) int {
	sum := 0
	for _, row := range rows {
		sum += possibleConfigurations(row)
	}

	return sum
}

// Bruteforce with bitmasks
func possibleConfigurations(row SpringRow) int {
	sum := 0
	unknowns := make([]int, 0, len(row.layout))
	for i, c := range row.layout {
		if c == '?' {
			unknowns = append(unknowns, i)
			// Set to . to begin with
			row.layout[i] = '.'
		}
	}

	possible := uint(math.Pow(2, float64(len(unknowns))))
	for i:= uint(0); i < possible; i++ {
		// Set the unknowns to the current configuration
		for j, unknown := range unknowns {
			mask := uint(1 << uint(j))
			// fmt.Printf("i: %03b, mask: %03b\n", i, mask)
			if i & mask == mask { // If the bit is 1
				row.layout[unknown] = '#'
			} else {
				row.layout[unknown] = '.'
			}
		}
		// fmt.Println("i: ", i, "layout: ", string(row.layout))
		if isPossible(row) {
			sum++
		}
	}

	return sum
}

func isPossible(row SpringRow) bool {
	re := regexp.MustCompile(`#+`)
	groups := re.FindAllString(string(row.layout), -1)
	lens := make([]int, 0, len(groups))
	for _, group := range groups {
		lens = append(lens, len(group))
	}

	return slices.Equal(row.groups, lens)
}

func parseInput(inputPath string) []SpringRow {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rows := make([]SpringRow, 0)

	re := regexp.MustCompile(`([.#\?]+) ((\d,?)+)`)

	for scanner.Scan() {
		groups := re.FindStringSubmatch(scanner.Text())
		layout := []byte(groups[1])

		springGroups := make([]int, 0)
		strs := strings.Split(groups[2], ",")
		for _, str := range strs {
			n, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			springGroups = append(springGroups, n)
		}

		rows = append(rows, SpringRow{layout, springGroups})
	}
	return rows
}
