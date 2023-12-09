package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1(("test.txt"))
	if result != 114 {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", result, 114)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test cases for Part2 function
}
