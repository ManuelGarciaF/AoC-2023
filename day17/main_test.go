package main

import (
	"testing"
)

func TestSolvePart1(t *testing.T) {
	heatLosses, sideLen := parseInput("sample")
	want := 102
	got := solvePart1(heatLosses, sideLen)
	if got != want {
		t.Errorf("solvePart1() = %v, want %v", got, want)
	}
}
