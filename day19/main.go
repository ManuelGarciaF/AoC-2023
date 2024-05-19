package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	workflows, parts := parseInput(os.Args[1])

	start := time.Now()
	fmt.Println("Part1:", solvePart1(workflows, parts))
	end := time.Now()
	fmt.Println("Time it took to solve part1: ", end.Sub(start))
	// fmt.Println("Part2:", solvePart2(workflows))
}

type Workflow struct {
	rules       []Rule
	defaultNext string
}

type Rule struct {
	stat      Stat
	condition Condition
	value     int
	next      string
}

type Stat int

const (
	X Stat = iota
	M
	A
	S
)

type Condition int

const (
	GT Condition = iota
	LT
)

type Part struct {
	x, m, a, s int
}

func (w Workflow) Run(part Part) string {
	for _, rule := range w.rules {
		if rule.IsPassedBy(part) {
			return rule.next
		}
	}
	return w.defaultNext
}

func (r Rule) IsPassedBy(part Part) bool {
	switch r.condition {
	case GT:
		return part.GetStat(r.stat) > r.value
	case LT:
		return part.GetStat(r.stat) < r.value
	}
	return false
}

func (p Part) GetStat(stat Stat) int {
	switch stat {
	case X:
		return p.x
	case M:
		return p.m
	case A:
		return p.a
	case S:
		return p.s
	}
	panic("Invalid Stat")
}

func (p Part) Total() int {
	return p.x + p.m + p.a + p.s
}

func solvePart1(workflows map[string]Workflow, parts []Part) int {
	sum := 0
	for _, part := range parts {
		curr := workflows["in"]
		next := ""
		for {
			next = curr.Run(part)
			if next == "A" {
				sum += part.Total()
				break
			}
			if next == "R" {
				break
			}
			curr = workflows[next]
		}
	}
	return sum
}

const MIN_STAT, MAX_STAT = 1, 4000

func solvePart2(workflows map[string]Workflow) int {
	return 0
}

func parseInput(inputPath string) (map[string]Workflow, []Part) {
	file, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	workflowRe := regexp.MustCompile(`([a-z]+){(([xmas][<>]\d+:[a-zAR]+,)+)([a-zAR]+)}`)
	RulesRe := regexp.MustCompile(`([xmas])([<>])(\d+):([a-zAR]+),`)
	PartRe := regexp.MustCompile(`{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}`)

	workflows := make(map[string]Workflow)
	parts := make([]Part, 0)

	scanningParts := false
	for scanner.Scan() {
		if scanner.Text() == "" {
			scanningParts = true
			continue
		}

		if !scanningParts { // Parse workflows
			rules := make([]Rule, 0)

			tokens := workflowRe.FindStringSubmatch(scanner.Text())
			name := tokens[1]
			rulesStr := tokens[2]
			defaultNext := tokens[4]

			rulesTokens := RulesRe.FindAllStringSubmatch(rulesStr, -1)
			for _, ruleTokens := range rulesTokens {
				stat := Stat(0)
				switch ruleTokens[1] {
				case "x":
					stat = X
				case "m":
					stat = M
				case "a":
					stat = A
				case "s":
					stat = S
				}

				condition := Condition(0)
				switch ruleTokens[2] {
				case "<":
					condition = LT
				case ">":
					condition = GT
				}

				value, err := strconv.Atoi(ruleTokens[3])
				if err != nil {
					panic("Input not in expected format: " + err.Error())
				}

				next := ruleTokens[4]

				rules = append(rules, Rule{stat, condition, value, next})
			}
			workflows[name] = Workflow{rules, defaultNext}

		} else { // Parse parts
			tokens := PartRe.FindStringSubmatch(scanner.Text())
			x, err := strconv.Atoi(tokens[1])
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
			m, err := strconv.Atoi(tokens[2])
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
			a, err := strconv.Atoi(tokens[3])
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
			s, err := strconv.Atoi(tokens[4])
			if err != nil {
				panic("Input not in expected format: " + err.Error())
			}
			parts = append(parts, Part{x, m, a, s})
		}
	}
	return workflows, parts
}
