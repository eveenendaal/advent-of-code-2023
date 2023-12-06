package main

import "testing"

func TestPart1(t *testing.T) {
	result := Part1("test.txt")
	// assert the result is 288
	if result != 288 {
		t.Errorf("Expected 288, got %d", result)
	}
}

func TestPart2(t *testing.T) {
	result := Part2("test.txt")
	// assert the result is 288
	if result != 71503 {
		t.Errorf("Expected 71503, got %d", result)
	}
}
