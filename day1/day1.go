package day1

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed 1.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

func Pt1() (int, error) {
	dial := 50
	zeroCount := 0

	for l := range strings.SplitSeq(inp, "\n") {
		dir := l[0]
		value, err := strconv.Atoi(l[1:])
		if err != nil {
			return 0, fmt.Errorf("Failed to convert str to int: %s", l[1:])
		}
		absValue := value
		if dir == 'L' {
			value *= -1
		}

		// account for multiple rotations of dial in one value
		if absValue > 100 {
			for range absValue / 100 {
				if value > 0 {
					dial, _ = incrementDial(dial, 100)
					value -= 100
				} else {
					dial, _ = incrementDial(dial, -100)
					value += 100
				}
			}
		}

		dial, _ = incrementDial(dial, value)

		if dial == 0 {
			zeroCount += 1
		}
	}

	return zeroCount, nil
}

func Pt2() (int, error) {
	dial := 50
	zeroCount := 0

	for l := range strings.SplitSeq(inp, "\n") {
		dir := l[0]
		value, err := strconv.Atoi(l[1:])
		if err != nil {
			return 0, fmt.Errorf("Failed to convert str to int: %s", l[1:])
		}
		if dir == 'L' {
			value *= -1
		}

		var zeros int
		dial, zeros = addToDial(dial, value)

		zeroCount += zeros
	}

	return zeroCount, nil
}

// addToDial simulate turning the dial as many clicks as passed in
func addToDial(d, v int) (int, int) {
	absValue := v
	if absValue < 0 {
		absValue = -absValue
	}

	zeros := (absValue / 100)
	newD := d + (v % 100)

	switch {
	case newD == 0:
		if d == 0 {
			return newD, zeros
		} else {
			return newD, zeros + 1
		}
	case newD > 99:
		return newD - 100, zeros + 1
	case newD < 0:
		newD += 100
		if d == 0 {
			return newD, zeros
		} else {
			return newD, zeros + 1
		}
	default:
		return newD, zeros
	}
}

// incrementDial simulate turning the dial no more than 100 clicks
func incrementDial(d, v int) (int, int) {
	dial := d + v
	crossed := false
	if dial >= 100 {
		dial -= 100
		crossed = true
	}
	if dial <= -1 {
		dial += 100
		crossed = true
	}

	if (crossed && d != 0) || dial == 0 {
		return dial, 1
	}

	return dial, 0
}
