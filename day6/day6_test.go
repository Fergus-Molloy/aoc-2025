package day6

import (
	"testing"
)

func TestDay6Pt1(t *testing.T) {
	t.Parallel()
	ans := 5060053676136

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 6 pt1: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 6 pt 1\nwant: %d\n got: %d", ans, v)
	}
}

func TestDay6Pt2(t *testing.T) {
	t.Parallel()
	ans := 9695042567249

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 6 pt2: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 6 pt 2\nwant: %d\n got: %d", ans, v)
	}
}
