package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	if result != 374 {
		t.Errorf("Result was incorrect, got: %d, want: %d.", result, 374)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test case
}
