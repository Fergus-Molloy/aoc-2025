package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func addToDial(d, v int) (int, int) {
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

func Pt1(inp string) (int, error) {
	inp = strings.TrimRight(inp, "\n")
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
					dial, _ = addToDial(dial, 100)
					value -= 100
				} else {
					dial, _ = addToDial(dial, -100)
					value += 100
				}
			}
		}

		dial, _ = addToDial(dial, value)

		if dial == 0 {
			zeroCount += 1
		}

	}

	return zeroCount, nil
}

func Pt2(inp string) (int, error) {
	inp = strings.TrimRight(inp, "\n")
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
		var i int
		var zeros int
		if absValue > 100 {
			for range absValue / 100 {
				if value > 0 {
					dial, i = addToDial(dial, 100)
					value -= 100
				} else {
					dial, i = addToDial(dial, -100)
					value += 100
				}
				zeros += i
			}
		}

		dial, i = addToDial(dial, value)
		zeros += i

		zeroCount += zeros
	}

	return zeroCount, nil
}
