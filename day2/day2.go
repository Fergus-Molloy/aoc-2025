package day2

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed 2.input
var inp string

func init() {
	inp = strings.TrimRight(inp, "\n")
}

func Pt1() (int, error) {
	var count int
	for ids := range strings.SplitSeq(inp, ",") {
		parts := strings.Split(ids, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		for i := range end + 1 - start {
			id := start + i
			if repeatsExactlyTwice(id) {
				count += id
			}
		}
	}

	return count, nil
}

func Pt2() (int, error) {
	var count int
	for ids := range strings.SplitSeq(inp, ",") {
		parts := strings.Split(ids, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, err
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return 0, err
		}

		for i := range end + 1 - start {
			id := start + i
			if repeatsAtLeastTwice(id) {
				count += id
			}
		}
	}

	return count, nil
}

func repeatsExactlyTwice(id int) bool {
	idStr := strconv.Itoa(id)
	if len(idStr)%2 != 0 {
		return false
	}
	return idStr[len(idStr)/2:] == idStr[:len(idStr)/2]
}

func repeatsAtLeastTwice(id int) bool {
	idStr := strconv.Itoa(id)

	for i := range len(idStr) / 2 {
		idx := i + 1
		if len(idStr)%idx != 0 {
			continue
		}
		subStr := idStr[:idx]
		fullStr := strings.Repeat(subStr, len(idStr)/idx)

		if fullStr == idStr {
			return true
		}
	}
	return false
}
