package main

import (
	_ "embed"
	"fmt"

	"fergus.molloy.xyz/aoc2025/day1"
)

//go:embed inputs/1
var inp1 string

func main() {
	fmt.Println(day1.Pt1(inp1))
}
