package day4

import "testing"

func TestDay4Pt1(t *testing.T) {
	t.Parallel()
	ans := 1395

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 4 pt1: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 4 pt 1\nwant: %d\n got: %d", ans, v)
	}
}

func TestDay4Pt2(t *testing.T) {
	t.Parallel()
	ans := 8451

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 4 pt2: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 4 pt 2\nwant: %d\n got: %d", ans, v)
	}
}
