package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	redTotal   = 12
	greenTotal = 13
	blueTotal  = 14
)

type Color int

const (
	Red Color = iota
	Blue
	Green
)

type Game []Reveal

type Reveal []Pull

type Pull struct {
	color Color
	n     int
}

func main() {
	games := parseInput("./input")
	// games := parseInput("./inputSample")

	fmt.Println(solvePart1(games))
}

func solvePart1(games map[int]Game) int {
	sum := 0

	for id, game := range games {
		possible := true
		for _, reveal := range game {
			red, green, blue := 0, 0, 0

			for _, pull := range reveal {
				switch pull.color {
				case Red:
					red += pull.n
					break
				case Green:
					green += pull.n
					break
				case Blue:
					blue += pull.n
					break
				}
			}
			if red > redTotal || blue > blueTotal || green > greenTotal {
				fmt.Printf("Marking game %d as impossible with r: %d, g: %d, b: %d\n", id, red, green, blue)
				possible = false
				break
			}
		}
		if possible {
			fmt.Printf("Game %d is possible\n", id)
			sum += id
		}
	}
	return sum
}

func parseInput(path string) map[int]Game {
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	gameStrings := strings.Split(string(bytes), "\n")
	gameStrings = gameStrings[:len(gameStrings)-1] // Remove last empty line

	games := make(map[int]Game)

	for lineN, gameString := range gameStrings {
		games[lineN+1] = parseGame(gameString)
	}

	return games
}

func parseGame(str string) Game {
	_, str, ok := strings.Cut(str, ": ")
	if !ok {
		panic("input not in expected format")
	}
	revealStrings := strings.Split(str, "; ")

	reveals := make([]Reveal, 0, len(revealStrings))

	for _, revealString := range revealStrings {
		reveals = append(reveals, parseReveal(revealString))
	}

	return reveals
}

func parseReveal(str string) Reveal {
	pulls := strings.Split(str, ", ")

	reveal := make(Reveal, 0, len(pulls))

	for _, pullStr := range pulls {
		reveal = append(reveal, parsePull(pullStr))
	}

	return reveal
}

func parsePull(str string) Pull {

	nStr, colorStr, ok := strings.Cut(str, " ")
	if !ok {
		panic("input not in expected format")
	}
	n, err := strconv.Atoi(nStr)
	if err != nil {
		panic("input not in expected format")
	}

	return Pull{color: parseColorStr(colorStr), n: n}
}

func parseColorStr(str string) Color {
	switch str {
	case "red":
		return Red
	case "blue":
		return Blue
	case "green":
		return Green
	default:
		return -1
	}
}
