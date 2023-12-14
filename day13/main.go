package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type InputImage [][]byte

func main() {
	images := parseInput(os.Args[1])

	fmt.Println("Part1: ", solve(images, false))
	fmt.Println("Part2: ", solve(images, true))
}

func solve(images []InputImage, allowSmudges bool) int {
	sum := 0
	for _, image := range images {
		val := imageSummary(image, allowSmudges)
		sum += val
	}
	return sum
}

func imageSummary(image InputImage, allowSmudges bool) int {
	// Iterate over rows
	for x := 0; x < len(image)-1; x++ {
		if rowSymm(image, x, allowSmudges) {
			return (x + 1) * 100
		}
	}

	// Transpose image
	transposed := make(InputImage, len(image[0]))
	for i := 0; i < len(transposed); i++ {
		transposed[i] = make([]byte, len(image))
	}
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[0]); j++ {
			transposed[j][i] = image[i][j]
		}
	}

	// Check "columns", rows of the transposed image
	for x := 0; x < len(transposed)-1; x++ {
		if rowSymm(transposed, x, allowSmudges) {
			return x + 1
		}
	}
	panic("No symmetry found")
}

func rowSymm(image InputImage, x int, allowSmudges bool) bool {
	perfect := true
	for offset := 0; offset < len(image); offset++ {
		// Check if we are still in bounds
		if x-offset < 0 || x+offset+1 >= len(image) {
			break
		}
		// If lines are not symmetric
		if string(image[x-offset]) != string(image[x+offset+1]) {
			if !perfect { // If there was already an error
				return false
			}
			for i := 0; i < len(image[0]); i++ { // Check there is only one error
				if image[x-offset][i] != image[x+offset+1][i] {
					if !perfect { // If there was already an error
						return false
					}
					perfect = false
				}
			}
		}
	}
	// Use perfects if !allowSmudges, ignore perfects if allowSmudges
	if (perfect && !allowSmudges) || (!perfect && allowSmudges) {
		return true
	}
	return false
}

func parseInput(inputPath string) []InputImage {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	images := make([]InputImage, 0)
	newImage := true
	for scanner.Scan() {
		if scanner.Text() == "" {
			newImage = true
			continue
		}
		if newImage {
			image := make(InputImage, 0)
			images = append(images, image)
			newImage = false
		}
		bytes := slices.Clone(scanner.Bytes()) // Weird stuff was happening
		images[len(images)-1] = append(images[len(images)-1], bytes)
	}
	return images
}
