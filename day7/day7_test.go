package day7

import (
	"testing"
)

func TestDay7Pt1(t *testing.T) {
	t.Parallel()
	ans := 1524

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 7 pt1: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 7 pt 1\nwant: %d\n got: %d", ans, v)
	}
}

func TestDay7Pt2(t *testing.T) {
	t.Parallel()
	ans := 32982105837605

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 7 pt2: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 7 pt 2\nwant: %d\n got: %d", ans, v)
	}
}
