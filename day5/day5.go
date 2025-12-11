package day5

import (
	"cmp"
	_ "embed"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

//go:embed 5.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

type freshRange struct {
	start, end int
}

func newFreshRange(s string) (freshRange, error) {
	parts := strings.Split(s, "-")
	if len(parts) != 2 {
		return freshRange{}, fmt.Errorf("range not formatted correctly, expected 2 parts got %d", len(parts))
	}

	start, err := strconv.Atoi(parts[0])
	if err != nil {
		return freshRange{}, fmt.Errorf("start: %w", err)
	}

	end, err := strconv.Atoi(parts[1])
	if err != nil {
		return freshRange{}, fmt.Errorf("start: %w", err)
	}

	// ensure start is always less than end
	if start < end {
		return freshRange{start, end}, nil
	} else {
		return freshRange{end, start}, nil
	}
}

func (f freshRange) inRange(i int) bool {
	return i >= f.start && i <= f.end
}

func mergeRanges(a, b freshRange) (freshRange, bool) {
	if a.start <= b.end && b.start <= a.end {
		n := freshRange{min(a.start, b.start), max(a.end, b.end)}
		return n, true
	}

	return freshRange{}, false
}

func Pt1() (int, error) {
	parts := strings.Split(inp, "\n\n")
	if len(parts) != 2 {
		return 0, fmt.Errorf("input not formatted correctly, expected 2 parts got %d", len(parts))
	}

	rangeLines := strings.Split(parts[0], "\n")
	ranges := make([]freshRange, 0, len(rangeLines))
	for _, r := range rangeLines {
		freshRange, err := newFreshRange(r)
		if err != nil {
			return 0, err
		}
		ranges = append(ranges, freshRange)
	}

	var count int
	for idStr := range strings.SplitSeq(parts[1], "\n") {
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return 0, err
		}
		for _, r := range ranges {
			if r.inRange(id) {
				count += 1
				break
			}
		}
	}

	return count, nil
}

func Pt2() (int, error) {
	return CountFreshIds(inp)
}

func CountFreshIds(inp string) (int, error) {
	parts := strings.Split(inp, "\n\n")
	if len(parts) != 2 {
		return 0, fmt.Errorf("input not formatted correctly, expected 2 parts got %d", len(parts))
	}

	rangeLines := strings.Split(parts[0], "\n")
	ranges := make([]freshRange, 0)
	for _, r := range rangeLines {
		newRange, err := newFreshRange(r)
		if err != nil {
			return 0, err
		}

		merged := false
		for i := range ranges {
			var n freshRange
			n, merged = mergeRanges(newRange, ranges[i])
			if merged {
				ranges[i] = n
				break
			}
		}
		if merged {
			ranges = reduceRanges(ranges)
		} else {
			ranges = append(ranges, newRange)
		}
	}

	var count int
	for _, r := range ranges {
		count += r.end - r.start + 1
	}

	return count, nil
}

func reduceRanges(ranges []freshRange) []freshRange {
	slices.SortFunc(ranges, func(a, b freshRange) int {
		return cmp.Compare(a.start, b.start)
	})
	merged := true
	for merged {
		merged = false
		for i := 0; i < len(ranges)-1; i++ {
			var n freshRange
			n, merged := mergeRanges(ranges[i], ranges[i+1])
			if merged {
				oldRanges := ranges[i+2:]
				ranges = append(ranges[:i], n)
				ranges = append(ranges, oldRanges...)
				break
			}
		}
	}

	return ranges
}
