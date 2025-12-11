package day6

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed 6.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

func parseProblemsPt1(s string) ([][]int, []string, error) {
	lines := strings.Split(s, "\n")
	finalLine := strings.Split(lines[len(lines)-1], " ")
	operators := make([]string, 0, len(finalLine))

	for _, op := range finalLine {
		if op != "" {
			operators = append(operators, op)
		}
	}

	problems := make([][]int, 0, len(lines))
	for _, l := range lines[:len(lines)-1] {
		ints := strings.Split(l, " ")

		line := make([]int, 0)
		for _, i := range ints {
			if i == "" {
				continue
			}

			v, err := strconv.Atoi(strings.TrimSpace(i))
			if err != nil {
				return nil, nil, err
			}

			line = append(line, v)
		}
		problems = append(problems, line)
	}

	if len(problems[0]) != len(operators) {
		return nil, nil, fmt.Errorf("number of numbers is different to operator count nums[%v] ops[%v]", len(problems[0]), len(operators))
	}

	return problems, operators, nil
}

func Pt1() (int, error) {
	numbers, operators, err := parseProblemsPt1(inp)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := range len(numbers[0]) {
		op := operators[i]

		var value int
		switch op {
		case "+":
			for j := range numbers {
				value += numbers[j][i]
			}
		case "*":
			value = 1
			for j := range numbers {
				value *= numbers[j][i]
			}
		}
		total += value
	}

	return total, nil
}

func parseProblemsPt2(s string) ([][]int, []string, error) {
	lines := strings.Split(s, "\n")
	numbers := lines[:len(lines)-1]
	finalLine := strings.Split(lines[len(lines)-1], " ")
	operators := make([]string, 0, len(finalLine))

	for _, op := range finalLine {
		if op != "" {
			operators = append(operators, op)
		}
	}

	problems := make([][]int, 0, len(lines))
	set := make([]int, 0)
	for i := range numbers[0] {
		if verticalSpace(numbers, i) {
			problems = append(problems, set)
			set = make([]int, 0)
			continue
		}

		str := ""
		for _, l := range numbers {
			str += string(l[i])
		}

		v, err := strconv.Atoi(strings.TrimSpace(str))
		if err != nil {
			return nil, nil, err
		}
		set = append(set, v)
	}
	problems = append(problems, set)

	if len(problems) != len(operators) {
		return nil, nil, fmt.Errorf("number of numbers is different to operator count nums[%v] ops[%v]", len(problems), len(operators))
	}

	return problems, operators, nil
}

func verticalSpace(numbers []string, i int) bool {
	spaceCount := 0
	for _, l := range numbers {
		if l[i] == ' ' {
			spaceCount += 1
		}
	}
	return spaceCount == len(numbers)
}

func Pt2() (int, error) {
	numbers, operators, err := parseProblemsPt2(inp)
	if err != nil {
		return 0, err
	}

	total := 0
	for i := range numbers {
		op := operators[i]

		var value int
		switch op {
		case "+":
			for j := range len(numbers[i]) {
				value += numbers[i][j]
			}
		case "*":
			value = 1
			for j := range len(numbers[i]) {
				value *= numbers[i][j]
			}
		}
		total += value
	}

	return total, nil
}
