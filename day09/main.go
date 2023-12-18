package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	// series := parseInput("./sample")
	series := parseInput("./input")

	fmt.Println(solvePart1(series))
	fmt.Println(solvePart2(series))
}

func solvePart1(series [][]int) int {
	sum := 0
	for _, sequence := range series {
		sum += nextValue(sequence)
	}
	return sum
}

func solvePart2(series [][]int) int {
	sum := 0
	for _, sequence := range series {
		sum += previousValue(sequence)
	}
	return sum
}

func nextValue(sequence []int) int {
	differences := calculateDifferences(sequence)

	// Calculate next digit bottom up
	value := 0
	for currDiff := len(differences) - 1; currDiff >= 0; currDiff-- {
		value += differences[currDiff][len(differences[currDiff])-1]
	}

	return value
}

func previousValue(sequence []int) int {
	differences := calculateDifferences(sequence)

	// Calculate previous digit bottom up
	value := 0
	for currDiff := len(differences) - 1; currDiff >= 0; currDiff-- {
		// fmt.Printf("%d, %v\n", value, differences[currDiff])
		value = differences[currDiff][0] - value
	}

	return value
}

func calculateDifferences(sequence []int) [][]int {
	differences := make([][]int, 1)
	differences[0] = sequence

	// Until the last difference is not all 0s
	for !all(differences[len(differences)-1], func(n int) bool { return n == 0 }) {
		lastDiff := differences[len(differences)-1]
		newDiff := make([]int, len(lastDiff)-1)
		for i := 0; i < len(lastDiff)-1; i++ {
			newDiff[i] = lastDiff[i+1] - lastDiff[i]
		}
		differences = append(differences, newDiff)
	}
	return differences
}

func parseInput(inputPath string) [][]int {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	re := regexp.MustCompile(`-?\d+`)
	series := make([][]int, 0)
	for scanner.Scan() {
		strs := re.FindAllString(scanner.Text(), -1)
		ints := make([]int, 0, len(strs))
		for _, str := range strs {
			n, _ := strconv.Atoi(str)
			ints = append(ints, n)
		}
		series = append(series, ints)
	}

	return series
}

func all[T any](ts []T, pred func(T) bool) bool {
	for _, t := range ts {
		if !pred(t) {
			return false
		}
	}
	return true
}
