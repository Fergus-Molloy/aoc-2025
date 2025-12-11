package day5

import (
	_ "embed"
	"testing"
)

//go:embed 5_test.input
var testInp string

func TestDay5Pt1(t *testing.T) {
	t.Parallel()
	ans := 517

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 5 pt1: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 5 pt 1\nwant: %d\n got: %d", ans, v)
	}
}

func TestDay5Pt2(t *testing.T) {
	t.Parallel()
	ans := 336173027056994

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 5 pt2: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 5 pt 2\nwant: %d\n got: %d", ans, v)
	}
}

func TestDay5Pt2Small(t *testing.T) {
	t.Parallel()

	v, err := CountFreshIds(`1-5
200-500
15-20
25-50
3-10

`)
	ans := 10 + 6 + 26 + 301
	if err != nil {
		t.Fatalf("failed to run day 5 pt2 small: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 5 pt 2 small\nwant: %d\n got: %d", ans, v)
	}
}

func TestDay5Pt2Test(t *testing.T) {
	t.Parallel()

	v, err := CountFreshIds(testInp)
	ans := 14
	if err != nil {
		t.Fatalf("failed to run day 5 pt2 test: %v", err)
	}
	if v != ans {
		t.Fatalf("incorrect answer for day 5 pt 2 test\nwant: %d\n got: %d", ans, v)
	}
}

func TestParseRanges(t *testing.T) {
	t.Parallel()
	r, err := newFreshRange("2-3")
	ans := freshRange{2, 3}
	if err != nil {
		t.Fatalf("could not parse range 2-3, want %v got %v", nil, err)
	}
	if r != ans {
		t.Fatalf("could not parse range 2-3, want %v got %v", ans, r)
	}

	r, err = newFreshRange("3-2")
	if err != nil {
		t.Fatalf("could not parse range 2-3, want %v got %v", nil, err)
	}
	if r != ans {
		t.Fatalf("could not parse range 2-3, want %v got %v", ans, r)
	}
}

func TestMergeRanges(t *testing.T) {
	td := []struct {
		name    string
		a, b, o freshRange
		merged  bool
	}{
		{
			name:   "same range",
			a:      freshRange{1, 2},
			b:      freshRange{1, 2},
			o:      freshRange{1, 2},
			merged: true,
		},
		{
			name:   "extend range 1",
			a:      freshRange{1, 2},
			b:      freshRange{2, 3},
			o:      freshRange{1, 3},
			merged: true,
		},
		{
			name:   "extend range 2",
			a:      freshRange{1, 3},
			b:      freshRange{2, 4},
			o:      freshRange{1, 4},
			merged: true,
		},
		{
			name:   "extend range 3",
			a:      freshRange{100, 200},
			b:      freshRange{250, 500},
			o:      freshRange{},
			merged: false,
		},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			o, m := mergeRanges(tc.a, tc.b)
			if o != tc.o {
				t.Fatalf("want range %v, got %v", tc.o, o)
			}
			if m != tc.merged {
				t.Fatalf("want merged %v, got %v", tc.merged, m)
			}

			o, m = mergeRanges(tc.b, tc.a)
			if o != tc.o {
				t.Fatalf("reverse want range %v, got %v", tc.o, o)
			}
			if m != tc.merged {
				t.Fatalf("want merged %v, got %v", tc.merged, m)
			}
		})
	}
}
