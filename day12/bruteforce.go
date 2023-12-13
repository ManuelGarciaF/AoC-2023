package main

import (
	"fmt"
	"math"
	"regexp"
	"slices"
)

func solvePart1Bruteforce(rows []SpringRow) int {
	sum := 0
	for _, row := range rows {
		val := possibleConfigurationsBruteforce(row)
		fmt.Print(val, ", ")
		sum += val
	}
	fmt.Println()

	return sum
}

// Bruteforce with bitmasks
func possibleConfigurationsBruteforce(row SpringRow) int {
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
