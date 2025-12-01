package day1

import "testing"

func BenchmarkDay1Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay1Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}
