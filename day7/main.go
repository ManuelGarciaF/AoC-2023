package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"unicode"
)

type Play struct {
	hand Hand
	bid  uint64
}

type Hand struct {
	cards    []Card // Ordered by card value
	handType HandType
}

type Card int

const (
	// ints from 2 to 9 and
	T Card = iota + 10
	J
	Q
	K
	A
)

type HandType int

const (
	highCard HandType = iota + 1
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

func main() {
	// plays := parseInput("./sample")
	plays := parseInput("./input")
	for _, play := range plays {
		fmt.Println(play.hand)
	}

	fmt.Println(solvePart1(plays))
}

func solvePart1(plays []Play) uint64 {
	slices.SortFunc(plays, func(a, b Play) int {
		if a.hand.handType == b.hand.handType {
			// Compare cards until one is different
			for i := 0; i < len(a.hand.cards); i++ {
				if a.hand.cards[i] != b.hand.cards[i] {
					return int(a.hand.cards[i]) - int(b.hand.cards[i])
				}
			}
		}
		return int(a.hand.handType) - int(b.hand.handType)
	})
	var total uint64 = 0
	for i, play := range plays {
		total += play.bid * uint64(i+1)
	}

	return total
}

func parseInput(inputPath string) []Play {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\S+`)

	plays := make([]Play, 0)
	for scanner.Scan() {
		tokens := re.FindAllString(scanner.Text(), 2)
		bid, err := strconv.Atoi(tokens[1])
		if err != nil {
			panic("Input not in expected format")
		}

		plays = append(plays, Play{
			hand: handFromString(tokens[0]),
			bid:  uint64(bid),
		})
	}

	return plays
}

func handFromString(str string) Hand {
	if len(str) != 5 {
		panic("Hand not made of 5 chars")
	}
	cards := make([]Card, 0)
	for _, char := range str {
		switch {
		case unicode.IsDigit(char):
			n, _ := strconv.Atoi(string(char))
			cards = append(cards, Card(n))
			break
		case char == 'T':
			cards = append(cards, T)
			break
		case char == 'J':
			cards = append(cards, J)
			break
		case char == 'Q':
			cards = append(cards, Q)
			break
		case char == 'K':
			cards = append(cards, K)
			break
		case char == 'A':
			cards = append(cards, A)
			break
		}
	}

	return Hand{
		cards:    cards,
		handType: getHandType(cards),
	}
}

func getHandType(cards []Card) HandType {
	counts := make(map[Card]int)
	for _, card := range cards {
		counts[card]++
	}

	countValues := mapValues(counts)
	slices.SortFunc(countValues, func(a, b int) int {
		return cmp.Compare(b, a)
	})

	switch countValues[0] {
	case 1:
		return highCard
	case 2:
		switch countValues[1] {
		case 1:
			return onePair
		case 2:
			return twoPair
		}
	case 3:
		switch countValues[1] {
		case 1:
			return threeOfAKind
		case 2:
			return fullHouse
		}
	case 4:
		return fourOfAKind
	case 5:
		return fiveOfAKind
	}

	panic("Unreachable")

}

func mapValues(m map[Card]int) []int {
	values := make([]int, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}
