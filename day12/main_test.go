package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	if result != 21 {
		t.Errorf("Expected 21, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test case for another example scenario
}
