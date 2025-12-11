package day6

import "testing"

func BenchmarkDay6Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay6Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}
