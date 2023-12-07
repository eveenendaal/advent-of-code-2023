package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	// assert the result is 6440
	if result != 6440 {
		t.Errorf("Expected 6440, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	// TODO: Write test cases for Part2 function
}
