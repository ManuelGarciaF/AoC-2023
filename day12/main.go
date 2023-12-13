package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type SpringRow struct {
	layout []byte
	groups []int
}

func main() {
	rows := parseInput(os.Args[1])

	fmt.Println("Part1: ", solve(rows, false))
	fmt.Println("Part2: ", solve(rows, true))
}

type Cache map[string]int

func (c Cache) get(layout []byte, groups []int) (int, bool) {
	key := fmt.Sprint(string(layout), groups)
	val, ok := c[key]
	return val, ok
}

func (c Cache) set(layout []byte, groups []int, result int) {
	key := fmt.Sprint(string(layout), groups)
	c[key] = result
}

func solve(rows []SpringRow, expand bool) int {
	// Memoization for the recursive calls
	cache := make(Cache, 0)

	sum := 0
	for _, row := range rows {
		if expand {
			row = expandRow(row)
		}
		val := possibleConfigurations(row.layout, row.groups, cache)
		sum += val
	}
	return sum
}

func possibleConfigurations(layout []byte, groups []int, cache Cache) int {
	// If there are no more groups and there are no '#' left it's valid
	if len(groups) == 0 {
		if i := bytes.IndexByte(layout, '#'); i != -1 {
			return 0
		}
		return 1
	}

	firstGroup := groups[0]
	otherGroups := groups[1:]
	// If the remaining groups don't fit, then it's invalid.
	if len(layout) < firstGroup {
		return 0
	}

	// Try to place first group in all possible offsets
	possibilities := 0
	for i := 0; i < len(layout)-(firstGroup-1); i++ {
		groupSlot := layout[i : i+firstGroup]
		// There can't be any '.' in the slot nor can there be a '#' right next to the slot.
		if index := bytes.IndexByte(groupSlot, '.'); index != -1 {
			continue
		}
		if i+firstGroup < len(layout) && layout[i+firstGroup] == '#' {
			continue
		}
		// If there is a '#' to the left of the slot, we've gone too far.
		if index := bytes.IndexByte(layout[:i], '#'); index != -1 {
			break
		}

		remainingLayout := layout[i+firstGroup:]
		// If we are not at the end, remove an extra char for the separator
		if len(remainingLayout) > 0 {
			remainingLayout = remainingLayout[1:]
		}

		val, ok := cache.get(remainingLayout, otherGroups)
		if !ok {
			val = possibleConfigurations(remainingLayout, otherGroups, cache)
			cache.set(remainingLayout, otherGroups, val)
		}
		possibilities += val

	}

	return possibilities
}

const EXPAND_FACTOR = 5

func expandRow(row SpringRow) SpringRow {
	layout := make([]byte, 0, len(row.layout)*EXPAND_FACTOR)
	groups := make([]int, 0, len(row.groups)*EXPAND_FACTOR)
	for i := 0; i < EXPAND_FACTOR; i++ {
		layout = append(layout, row.layout...)
		layout = append(layout, '?')
		groups = append(groups, row.groups...)
	}
	// Added one too many '?'
	layout = layout[:len(layout)-1]
	return SpringRow{layout, groups}
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
