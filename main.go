package main

import (
	_ "embed"
	"fmt"

	"fergus.molloy.xyz/aoc2025/day1"
	"fergus.molloy.xyz/aoc2025/day2"
	"fergus.molloy.xyz/aoc2025/day3"
)

func main() {
	fmt.Println(day1.Pt1())
	fmt.Println(day1.Pt2())
	fmt.Println(day2.Pt1())
	fmt.Println(day2.Pt2())
	fmt.Println(day3.Pt1())
	fmt.Println(day3.Pt2())
}
