package day2

import (
	"strconv"
	"testing"
)

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

func BenchmarkDay2Pt2Worker(b *testing.B) {
	workers := []int{8, 16, 24, 25, 50}
	for _, n := range workers {
		b.Run(strconv.Itoa(n), func(b *testing.B) {
			for b.Loop() {
				Pt2Worker(n)
			}
		})
	}
}
