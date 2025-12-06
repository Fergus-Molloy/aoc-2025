package day4

import (
	_ "embed"
	"fmt"
	"iter"
	"slices"
	"strings"
)

//go:embed 4.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

type point struct {
	x, y int
}

func Pt1() (int, error) {
	count := 0
	movableRolls := make(map[point]struct{})
	grid := strings.Split(inp, "\n")
	for y, line := range grid {
		for x, r := range line {
			rolls := countRolls(x, y, grid)
			if rolls < 4 && r == '@' {
				_, counted := movableRolls[point{x, y}]
				if !counted {
					movableRolls[point{x, y}] = struct{}{}
					count += 1
				}
			}
		}
	}

	return count, nil
}

func Pt2() (int, error) {
	count := 0
	grid := strings.Split(inp, "\n")
	newGrid := make([]string, len(grid))
	for !slices.Equal(grid, newGrid) {
		movableRolls := make(map[point]struct{})
		copy(newGrid, grid)
		for y, line := range newGrid {
			for x, r := range line {
				rolls := countRolls(x, y, newGrid)
				if rolls < 4 && r == '@' {
					_, counted := movableRolls[point{x, y}]
					if !counted {
						movableRolls[point{x, y}] = struct{}{}
						count += 1
					}
				}
			}
		}
		grid = removeRolls(movableRolls, newGrid)
	}

	return count, nil
}

func countRolls(x, y int, grid []string) int {
	count := 0
	for p := range getValidCoord(x, y, len(grid[0]), len(grid)) {
		if grid[p.y][p.x] == '@' {
			count += 1
		}
	}
	return count
}

func getValidCoord(x, y, w, h int) iter.Seq[point] {
	return func(yield func(point) bool) {
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				// skip invalid coords
				switch {
				case i < 0:
					continue
				case i >= w:
					continue
				case j < 0:
					continue
				case j >= h:
					continue
				case i == x && j == y: // not counting middle point towards total
					continue
				}
				if !yield(point{i, j}) {
					return
				}
			}
		}
	}
}

func removeRolls(f map[point]struct{}, grid []string) []string {
	g := make([]string, len(grid))
	copy(g, grid)

	for p := range f {
		x := p.x
		y := p.y
		line := []rune(g[y])
		line[x] = '.'
		g[y] = string(line)
	}

	return g
}

func printFound(f map[point]struct{}, grid []string) {
	g := make([]string, len(grid))
	copy(g, grid)

	for p := range f {
		x := p.x
		y := p.y
		line := []rune(g[y])
		line[x] = 'x'
		g[y] = string(line)
	}

	for _, line := range g {
		fmt.Println(line)
	}
}
