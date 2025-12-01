package day1

import "testing"

func TestDay1Pt1(t *testing.T) {
	t.Parallel()

	v, err := Pt1()
	if err != nil {
		t.Fatalf("failed to run day 1 pt1: %v", err)
	}
	if v != 1026 {
		t.Fatalf("incorrect answer for day 1 pt 1\nwant: 1026\n got: %d", v)
	}
}

func TestDay1Pt2(t *testing.T) {
	t.Parallel()

	v, err := Pt2()
	if err != nil {
		t.Fatalf("failed to run day 1 pt2: %v", err)
	}
	if v != 5923 {
		t.Fatalf("incorrect answer for day 1 pt 2\nwant: 5923\n got: %d", v)
	}
}

func TestAddToDial(t *testing.T) {
	td := []struct {
		name                             string
		dial, value, wantDial, wantZeros int
	}{
		{
			name:      "R50 lands on zero",
			dial:      50,
			value:     50,
			wantDial:  0,
			wantZeros: 1,
		},
		{
			name:      "L50 lands on zero",
			dial:      50,
			value:     -50,
			wantDial:  0,
			wantZeros: 1,
		},
		{
			name:      "R50 passes zero",
			dial:      75,
			value:     50,
			wantDial:  25,
			wantZeros: 1,
		},
		{
			name:      "L50 passes zero",
			dial:      25,
			value:     -50,
			wantDial:  75,
			wantZeros: 1,
		},
		{
			name:      "R500 passes zero",
			dial:      25,
			value:     500,
			wantDial:  25,
			wantZeros: 5,
		},
		{
			name:      "L500 passes zero",
			dial:      25,
			value:     -500,
			wantDial:  25,
			wantZeros: 5,
		},
		{
			name:      "R550 lands on zero",
			dial:      50,
			value:     550,
			wantDial:  0,
			wantZeros: 6,
		},
		{
			name:      "L550 lands on zero",
			dial:      50,
			value:     -550,
			wantDial:  0,
			wantZeros: 6,
		},
		{
			name:      "R100 start and end on zero",
			dial:      0,
			value:     100,
			wantDial:  0,
			wantZeros: 1,
		},
		{
			name:      "L100 start and end on zero",
			dial:      0,
			value:     100,
			wantDial:  0,
			wantZeros: 1,
		},
		{
			name:      "L557 dial 0",
			dial:      0,
			value:     -557,
			wantDial:  43,
			wantZeros: 5,
		},
		{
			name:      "L39 dial 0",
			dial:      0,
			value:     -39,
			wantDial:  61,
			wantZeros: 0,
		},
	}

	for _, tc := range td {
		t.Run(tc.name, func(t *testing.T) {
			gotDial, gotZeros := addToDial(tc.dial, tc.value)

			if gotDial != tc.wantDial {
				t.Fatalf("incorrect value on dial\nwant: %d\n got: %d", tc.wantDial, gotDial)
			}

			if gotZeros != tc.wantZeros {
				t.Fatalf("incorrect zeros\nwant: %d\n got: %d", tc.wantZeros, gotZeros)
			}
		})
	}
}
