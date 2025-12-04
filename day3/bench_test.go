package day3

import "testing"

func BenchmarkDay3Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay3Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}

func BenchmarkDay3Pt2Parallel(b *testing.B) {
	for b.Loop() {
		Pt2Parallel()
	}
}
