package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Race struct {
	time, record uint64
}

func (r Race) distanceWith(prepTime uint64) uint64 {
	runTime := r.time - prepTime
	return prepTime * runTime // prepTime == speed in mm/ms
}

func main() {
	// races := parseInput("./inputSample")
	races := parsePart1("./input")
	fmt.Printf("%+v\n", races)
	fmt.Println(solvePart1(races))

	race := parsePart2("./input")
	fmt.Printf("%+v\n", race)
	fmt.Println(solvePart2(race))
}

func solvePart1(races []Race) uint64 {
	// By doing some kinematics, we know the max distance is at timelimit/2,
	// we can check either side until we don't break the record

	var answer uint64 = 1
	for _, race := range races {
		var sum uint64 = 0

		// >= prepTime than max distance
		for prepTime := race.time / 2; race.distanceWith(prepTime) > race.record; prepTime++ {
			sum++
		}
		// < prepTime than max distance
		for prepTime := (race.time / 2) - 1; race.distanceWith(prepTime) > race.record; prepTime-- {
			sum++
		}
		answer *= sum
	}
	return answer
}

func solvePart2(race Race) uint64 {
	var sum uint64 = 0

	// >= prepTime than max distance
	for prepTime := race.time / 2; race.distanceWith(prepTime) > race.record; prepTime++ {
		sum++
	}
	// < prepTime than max distance
	for prepTime := (race.time / 2) - 1; race.distanceWith(prepTime) > race.record; prepTime-- {
		sum++
	}

	return sum
}

func parsePart1(inputPath string) []Race {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Kind of overkill for this problem but wanted to try regex
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Scan first line
	re := regexp.MustCompile(`\d+`)
	line := scanner.Text()
	times := re.FindAllString(line, -1)
	scanner.Scan() // Second line
	line = scanner.Text()
	records := re.FindAllString(line, -1)

	races := make([]Race, 0, len(times))
	for i := 0; i < len(times); i++ {
		time, err := strconv.Atoi(times[i])
		if err != nil {
			panic("Input not in expected format")
		}
		record, err := strconv.Atoi(records[i])
		if err != nil {
			panic("Input not in expected format")
		}
		races = append(races, Race{
			time:   uint64(time),
			record: uint64(record),
		})
	}

	return races
}

func parsePart2(inputPath string) Race {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Kind of overkill for this problem but wanted to try regex
	scanner := bufio.NewScanner(file)
	scanner.Scan() // Scan first line
	re := regexp.MustCompile(`\d(\d|\s)*`)
	line := scanner.Text()
	timeStr := re.FindString(line)
	scanner.Scan() // Second line
	line = scanner.Text()
	recordStr := re.FindString(line)

	time, err := strconv.Atoi(strings.ReplaceAll(timeStr, " ", ""))
	if err != nil {
		panic("Input not in expected format")
	}
	record, err := strconv.Atoi(strings.ReplaceAll(recordStr, " ", ""))
	if err != nil {
		panic("Input not in expected format")
	}
	race := Race{
		time:   uint64(time),
		record: uint64(record),
	}

	return race
}
