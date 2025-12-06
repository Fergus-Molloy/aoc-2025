package day4

import "testing"

func BenchmarkDay4Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay4Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}
