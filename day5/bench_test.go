package day5

import "testing"

func BenchmarkDay5Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay5Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}
