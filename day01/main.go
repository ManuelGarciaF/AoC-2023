package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(content), "\n")
	sum := 0
	for c, line := range lines {
		digits := findDigits(line)
		if len(digits) > 0 {
			val, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
			sum += val
			fmt.Println("Line: ", c, " Value: ", val, " sum: ", sum)
		}
	}
	fmt.Println(sum)
}

func findDigits(str string) string {
	filtered := ""
	for i := 0; i < len(str); i++ {
		if isDigit(str[i]) {
			filtered += string(str[i])
		} else if val, n, ok := startsWithWordNumber(str[i:]); ok {
			filtered += val
			i += n - 1
		}
	}

	return filtered
}

var digits = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func startsWithWordNumber(str string) (digitNumber string, wordLength int, ok bool) {
	for word, digit := range digits {
		if strings.HasPrefix(str, word) {
			return digit, len(word), true
		}
	}
	return "", 0, false
}

func isDigit(c byte) bool {
	_, err := strconv.Atoi(string(c))
	return err == nil
}
