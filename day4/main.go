package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// bytes, err := os.ReadFile("./inputSample")
	bytes, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println(solvePart1(string(bytes)))
	fmt.Println(solvePart2(string(bytes)))

}

func solvePart1(input string) int {
	cards := strings.Split(input, "\n")
	cards = cards[:len(cards)-1] // Remove last empty line

	sum := 0

	for _, card := range cards {
		_, card, _ = strings.Cut(card, ": ")
		cardNumbersStr, winningNumbersStr, _ := strings.Cut(card, " | ")
		cardNumbers := parseNumList(cardNumbersStr)
		winningNumbers := parseNumList(winningNumbersStr)
		slices.Sort(winningNumbers)

		points := 0
		for _, cardNumber := range cardNumbers {
			if _, found := slices.BinarySearch(winningNumbers, cardNumber); found {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}
		sum += points

	}

	return sum
}

func solvePart2(input string) int {
	cards := strings.Split(input, "\n")
	cards = cards[:len(cards)-1] // Remove last empty line

	totalCards := 0
	toProcess := make([]int, 0, len(cards)) // Queue for cards to process
	toProcess = append(toProcess, nextNInts(0, len(cards))...)

	// Calculate values for each card once
	cardValues := make([]int, 0, len(cards))
	for _, card := range cards {
		cardValues = append(cardValues, correctNumbers(card))
	}

	for len(toProcess) > 0 {
		index := toProcess[0]
		toProcess = append(toProcess, nextNInts(index+1, cardValues[index])...)
		totalCards++
		toProcess = toProcess[1:] // Remove first element
	}
	return totalCards
}

func correctNumbers(card string) int {
	_, card, _ = strings.Cut(card, ": ")
	cardNumbersStr, winningNumbersStr, _ := strings.Cut(card, " | ")
	cardNumbers := parseNumList(cardNumbersStr)
	winningNumbers := parseNumList(winningNumbersStr)
	slices.Sort(winningNumbers)

	sum := 0
	for _, cardNumber := range cardNumbers {
		if _, found := slices.BinarySearch(winningNumbers, cardNumber); found {
			sum++
		}
	}
	return sum
}

func parseNumList(str string) []int {
	strs := strings.Split(str, " ")
	ints := make([]int, 0, len(strs))

	for _, numStr := range strs {
		if numStr != "" {
			n, err := strconv.Atoi(numStr)
			if err != nil {
				panic("Input not in expected format")
			}

			ints = append(ints, n)
		}
	}

	return ints
}

func nextNInts(start, n int) []int {
	s := make([]int, 0, n)
	for i := start; i < (n + start); i++ {
		s = append(s, i)
	}
	return s
}
