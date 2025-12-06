package day3

import "testing"

func TestDay3Pt1(t *testing.T) {
	t.Parallel()

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 3 pt1: %v", err)
	}
	if v != 17430 {
		t.Fatalf("incorrect answer for day 3 pt 1\nwant: 17430\n got: %d", v)
	}
}

func TestDay3Pt2(t *testing.T) {
	t.Parallel()

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 3 pt2: %v", err)
	}
	if v != 171975854269367 {
		t.Fatalf("incorrect answer for day 3 pt 2\nwant: 171975854269367\n got: %d", v)
	}
}
