package day2

import "testing"

func BenchmarkDay2Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay2Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}

func BenchmarkDay2Pt2Parallel(b *testing.B) {
	for b.Loop() {
		Pt2Parallel()
	}
}
