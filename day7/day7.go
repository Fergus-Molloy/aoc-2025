package day7

import (
	_ "embed"
	"maps"
	"strings"
)

//go:embed 7.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

func Pt1() (int, error) {
	beamIdx := make([]int, 0)
	lines := strings.Split(inp, "\n")
	beamIdx = append(beamIdx, strings.IndexAny(lines[0], "S"))

	count := 0
	for _, l := range lines {
		newIdx := make(map[int]struct{})
		for _, i := range beamIdx {
			if l[i] == '^' {
				_, found := newIdx[i-1]
				if !found {
					newIdx[i-1] = struct{}{}
				}
				_, found = newIdx[i+1]
				if !found {
					newIdx[i+1] = struct{}{}
				}

				count += 1
			} else {
				_, found := newIdx[i]
				if !found {
					newIdx[i] = struct{}{}
				}
			}
		}
		beamIdx = make([]int, 0, len(newIdx))
		for idx := range newIdx {
			beamIdx = append(beamIdx, idx)
		}
	}

	return count, nil
}

func Pt2() (int, error) {
	beamIdx := make(map[int]int, 0)
	lines := strings.Split(inp, "\n")
	beamIdx[strings.IndexAny(lines[0], "S")] = 1

	for _, l := range lines {
		newIdx := make(map[int]int)
		for i, c := range beamIdx {
			if l[i] == '^' {
				v, found := newIdx[i-1]
				if !found {
					newIdx[i-1] = c
				} else {
					newIdx[i-1] = v + c
				}
				v, found = newIdx[i+1]
				if !found {
					newIdx[i+1] = c
				} else {
					newIdx[i+1] = v + c
				}
			} else {
				v, found := newIdx[i]
				if !found {
					newIdx[i] = c
				} else {
					newIdx[i] = v + c
				}
			}
		}
		beamIdx = make(map[int]int)
		maps.Copy(beamIdx, newIdx)
	}

	count := 0
	for _, v := range beamIdx {
		count += v
	}

	return count, nil
}
