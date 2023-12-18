package main

import (
	"testing"
)

func BenchmarkSolvePart1(b *testing.B) {
	heatLosses, sideLen := parseInput("sample")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(heatLosses, sideLen, 0, 3)
	}
}

func BenchmarkSolvePart2(b *testing.B) {
	heatLosses, sideLen := parseInput("sample")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		solve(heatLosses, sideLen, 3, 10)
	}
}
