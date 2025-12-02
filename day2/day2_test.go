package day2

import "testing"

func TestDay2Pt1(t *testing.T) {
	t.Parallel()

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 2 pt1: %v", err)
	}
	if v != 56660955519 {
		t.Fatalf("incorrect answer for day 2 pt 1\nwant: 56660955519\n got: %d", v)
	}
}

func TestDay2Pt2(t *testing.T) {
	t.Parallel()

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 2 pt2: %v", err)
	}
	if v != 79183223243 {
		t.Fatalf("incorrect answer for day 2 pt 2\nwant: 79183223243\n got: %d", v)
	}
}
