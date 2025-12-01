package day1

import "testing"

func TestDay1Pt1(t *testing.T) {
	v, err := Pt1()
	if err != nil {
		t.Fatalf("Failed to run Day 1 Pt1: %v", err)
	}
	if v != 1026 {
		t.Fatalf("Incorrect answer for Day 1 Pt 1\nwant: 1026\n got: %d", v)
	}
}

func TestDay1Pt2(t *testing.T) {
	v, err := Pt2()
	if err != nil {
		t.Fatalf("Failed to run Day 1 Pt2: %v", err)
	}
	if v != 5923 {
		t.Fatalf("Incorrect answer for Day 1 Pt 2\nwant: 5923\n got: %d", v)
	}
}
