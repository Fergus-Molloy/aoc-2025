package day1

import (
	"fmt"
	"strconv"
	"strings"
)

func Pt1(inp string) int {
	inp = strings.TrimRight(inp, "\n")
	var sum int
	for l := range strings.SplitSeq(inp, "\n") {
		v, err := strconv.Atoi(l)
		if err != nil {
			fmt.Printf("could not parse %s as int\n", l)
			v = 0
		}
		sum += v
	}
	return sum
}
