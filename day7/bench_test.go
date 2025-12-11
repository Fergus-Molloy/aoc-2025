package day7

import "testing"

func BenchmarkDay7Pt1(b *testing.B) {
	for b.Loop() {
		Pt1()
	}
}

func BenchmarkDay7Pt2(b *testing.B) {
	for b.Loop() {
		Pt2()
	}
}
